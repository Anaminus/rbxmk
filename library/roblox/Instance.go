package reflect

import (
	"bytes"
	"fmt"
	"strings"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/formats"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/types"
)

// pushPropertyTo behaves like PushVariantTo, except that exprims types are
// reflected as userdata.
func pushPropertyTo(s State, v types.Value) (lv lua.LValue, err error) {
	rfl := s.Reflector(v.Type())
	if rfl.Name == "" {
		return nil, fmt.Errorf("unknown type %q", string(v.Type()))
	}
	if rfl.PushTo == nil {
		return nil, fmt.Errorf("unable to cast %s to Variant", rfl.Name)
	}
	if rfl.Flags&rbxmk.Exprim == 0 {
		return PushVariantTo(s, v)
	}
	u := s.UserDataOf(v, rfl.Name)
	return u, nil
}

func checkEnumDesc(s State, desc *rtypes.RootDesc, name, class, prop string) *rtypes.Enum {
	var enumValue *rtypes.Enum
	if desc.EnumTypes != nil {
		enumValue = desc.EnumTypes.Enum(name)
	}
	if enumValue == nil {
		if desc.Enums[name] == nil {
			s.RaiseError(
				"no enum descriptor %q for property descriptor %s.%s",
				name,
				class,
				prop,
			)
			return nil
		}
		s.RaiseError(
			"no enum value %q generated for property descriptor %s.%s",
			name,
			class,
			prop,
		)
		return nil
	}
	return enumValue
}

func checkClassDesc(s State, desc *rtypes.RootDesc, name, class, prop string) *rbxdump.Class {
	classDesc := desc.Classes[name]
	if classDesc == nil {
		s.RaiseError(
			"no class descriptor %q for property descriptor %s.%s",
			name,
			class,
			prop,
		)
		return nil
	}
	return classDesc
}

func defaultAttrConfig(inst *rtypes.Instance) rtypes.AttrConfig {
	attrcfg := inst.AttrConfig()
	if attrcfg != nil {
		return *attrcfg
	}
	return rtypes.AttrConfig{Property: "AttributesSerialize"}
}

func getAttributes(s State, inst *rtypes.Instance) rtypes.Dictionary {
	attrcfg := defaultAttrConfig(inst)
	sv, ok := inst.Get(attrcfg.Property).(types.Stringlike)
	if ok {
		s.RaiseError("property %q is not string-like", attrcfg.Property)
		return nil
	}
	r := strings.NewReader(sv.Stringlike())
	dict, err := formats.RBXAttr().Decode(s.Global, nil, r)
	if err != nil {
		s.RaiseError("decode attributes from %q: %s", attrcfg.Property, err)
		return nil
	}
	return dict.(rtypes.Dictionary)
}

func setAttributes(s State, inst *rtypes.Instance, dict rtypes.Dictionary) {
	attrcfg := defaultAttrConfig(inst)
	var w bytes.Buffer
	if err := formats.RBXAttr().Encode(s.Global, nil, &w, dict); err != nil {
		s.RaiseError("encode attributes to %q: %s", attrcfg.Property, err)
		return
	}
	inst.Set(attrcfg.Property, types.BinaryString(w.Bytes()))
}

func init() { register(Instance) }
func Instance() Reflector {
	return Reflector{
		Name: "Instance",
		PushTo: func(s State, v types.Value) (lvs []lua.LValue, err error) {
			if inst, ok := v.(*rtypes.Instance); ok && inst == nil {
				return append(lvs, lua.LNil), nil
			}
			u := s.UserDataOf(v, "Instance")
			return append(lvs, u), nil
		},
		PullFrom: func(s State, lvs ...lua.LValue) (v types.Value, err error) {
			switch lv := lvs[0].(type) {
			case *lua.LNilType:
				return (*rtypes.Instance)(nil), nil
			case *lua.LUserData:
				if lv.Metatable != s.L.GetTypeMetatable("Instance") {
					return nil, rbxmk.TypeError(nil, 0, "Instance")
				}
				v, ok := lv.Value.(types.Value)
				if !ok {
					return nil, rbxmk.TypeError(nil, 0, "Instance")
				}
				return v, nil
			default:
				return nil, rbxmk.TypeError(nil, 0, "Instance")
			}
		},
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "Instance").(*rtypes.Instance)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "Instance").(*rtypes.Instance)
				op := s.Pull(2, "Instance").(*rtypes.Instance)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
			"__index": func(s State) int {
				inst := s.Pull(1, "Instance").(*rtypes.Instance)

				// Try symbol.
				if typ := s.Reflector("Symbol"); typ.Name != "" {
					if sym, err := typ.PullFrom(s, s.L.CheckAny(2)); err == nil {
						name := sym.(rtypes.Symbol).Name
						switch name {
						case "Reference":
							return s.Push(types.String(inst.Reference))
						case "IsService":
							return s.Push(types.Bool(inst.IsService))
						case "Desc":
							desc := inst.Desc()
							if desc == nil {
								return s.Push(rtypes.Nil)
							}
							return s.Push(desc)
						case "RawDesc":
							desc, blocked := inst.RawDesc()
							if blocked {
								return s.Push(types.False)
							}
							if desc == nil {
								return s.Push(rtypes.Nil)
							}
							return s.Push(desc)
						case "AttrConfig":
							attrcfg := inst.AttrConfig()
							if attrcfg == nil {
								return s.Push(rtypes.Nil)
							}
							return s.Push(attrcfg)
						case "RawAttrConfig":
							attrcfg, blocked := inst.RawAttrConfig()
							if blocked {
								return s.Push(types.False)
							}
							if attrcfg == nil {
								return s.Push(rtypes.Nil)
							}
							return s.Push(attrcfg)
						case "Metadata":
							if meta := inst.Metadata(); meta != nil {
								dict := make(rtypes.Dictionary, len(meta))
								for k, v := range meta {
									dict[k] = types.String(v)
								}
								return s.Push(dict)
							}
						}
						s.L.RaiseError("symbol %s is not a valid member", name)
						return 0
					}
				}

				name := string(s.Pull(2, "string").(types.String))
				desc := s.Desc.Of(inst)
				var classDesc *rbxdump.Class
				if desc != nil {
					classDesc = desc.Classes[inst.ClassName]
				}

				// Try GetService.
				if inst.IsDataModel() && name == "GetService" {
					s.L.Push(s.L.NewFunction(func(l *lua.LState) int {
						u := l.CheckUserData(1)
						if u.Metatable != l.GetTypeMetatable("Instance") {
							rbxmk.TypeError(l, 1, "Instance")
							return 0
						}
						inst, ok := u.Value.(*rtypes.Instance)
						if !ok {
							rbxmk.TypeError(l, 1, "Instance")
							return 0
						}
						s := State{World: s.World, L: l}
						className := string(s.Pull(2, "string").(types.String))
						if desc != nil {
							classDesc := desc.Classes[className]
							if classDesc == nil || !classDesc.GetTag("Service") {
								return s.RaiseError("%q is not a valid service", className)
							}
						}
						service := inst.FindFirstChildOfClass(className, false)
						if service == nil {
							service = rtypes.NewInstance(className, nil)
							service.IsService = true
							service.SetName(className)
							service.SetParent(inst)
						}
						return s.Push(service)
					}))
					return 1
				}

				// Try property.
				var lv lua.LValue
				var err error
				value := inst.Get(name)
				if classDesc != nil {
					propDesc := desc.Property(classDesc.Name, name)
					if propDesc == nil {
						return s.RaiseError("%s is not a valid member", name)
					}
					if value == nil {
						return s.RaiseError("property %s not initialized", name)
					}
					switch propDesc.ValueType.Category {
					case "Class":
						switch value := value.(type) {
						case *rtypes.Instance:
							if value == nil {
								return s.Push(rtypes.Nil)
							}
							class := checkClassDesc(s, desc, propDesc.ValueType.Name, classDesc.Name, propDesc.Name)
							if class == nil {
								return 0
							}
							if !value.WithDescIsA(desc, class.Name) {
								return s.RaiseError("instance of class %s expected, got %s", class.Name, value.ClassName)
							}
							return s.Push(value)
						default:
							return s.RaiseError("stored value type %s is not an instance", value.Type())
						}
					case "Enum":
						enum := checkEnumDesc(s, desc, propDesc.ValueType.Name, classDesc.Name, propDesc.Name)
						if enum == nil {
							return 0
						}
						token, ok := value.(types.Token)
						if !ok {
							return s.RaiseError("stored value type %s is not a token", value.Type())
						}
						item := enum.Value(int(token))
						if item == nil {
							return s.RaiseError("invalid stored value %d for enum %s", value, enum.Name())
						}
						return s.Push(item)
					default:
						if a, b := value.Type(), propDesc.ValueType.Name; a != b {
							return s.RaiseError("stored value type %s does not match property type %s", a, b)
						}
					}
					// Push without converting exprims.
					lv, err = PushVariantTo(s, value)
				} else {
					if value == nil {
						// Fallback to nil.
						return s.Push(rtypes.Nil)
					}
					lv, err = pushPropertyTo(s, value)
				}
				if err != nil {
					return s.RaiseError("%s", err)
				}
				s.L.Push(lv)
				return 1
			},
			"__newindex": func(s State) int {
				inst := s.Pull(1, "Instance").(*rtypes.Instance)

				// Try symbol.
				if typ := s.Reflector("Symbol"); typ.Name != "" {
					if sym, err := typ.PullFrom(s, s.L.CheckAny(2)); err == nil {
						name := sym.(rtypes.Symbol).Name
						switch name {
						case "Reference":
							inst.Reference = string(s.Pull(3, "string").(types.String))
							return 0
						case "IsService":
							inst.IsService = bool(s.Pull(3, "bool").(types.Bool))
							return 0
						case "Desc", "RawDesc":
							switch v := s.PullAnyOf(3, "RootDesc", "bool", "nil").(type) {
							case *rtypes.RootDesc:
								inst.SetDesc(v, false)
							case types.Bool:
								if v {
									return s.RaiseError("descriptor cannot be true")
								}
								inst.SetDesc(nil, true)
							case rtypes.NilType:
								inst.SetDesc(nil, false)
							}
							return 0
						case "AttrConfig", "RawAttrConfig":
							switch v := s.PullAnyOf(3, "Attr", "bool", "nil").(type) {
							case *rtypes.AttrConfig:
								inst.SetAttrConfig(v, false)
							case types.Bool:
								if v {
									return s.RaiseError("AttrConfig cannot be true")
								}
								inst.SetAttrConfig(nil, true)
							case rtypes.NilType:
								inst.SetAttrConfig(nil, false)
							}
							return 0
						case "Metadata":
							if meta := inst.Metadata(); meta != nil {
								dict := s.Pull(3, "Dictionary").(rtypes.Dictionary)
								for k := range meta {
									delete(meta, k)
								}
								for k, v := range dict {
									w, ok := v.(types.String)
									if !ok {
										return s.RaiseError("field %q: string expected, got %s (%T)", k, v.Type(), v)
									}
									meta[k] = string(w)
								}
								return 0
							}
						}
						s.L.RaiseError("symbol %s is not a valid member", name)
						return 0
					}
				}

				name := string(s.Pull(2, "string").(types.String))

				// Try GetService.
				if inst.IsDataModel() && name == "GetService" {
					return s.RaiseError("%s cannot be assigned to", name)
				}

				// Try property.
				value := PullVariant(s, 3)

				desc := s.Desc.Of(inst)
				var classDesc *rbxdump.Class
				if desc != nil {
					classDesc = desc.Classes[inst.ClassName]
				}
				if classDesc != nil {
					propDesc := desc.Property(classDesc.Name, name)
					if propDesc == nil {
						return s.RaiseError("%s is not a valid member", name)
					}
					switch propDesc.ValueType.Category {
					case "Class":
						switch value := value.(type) {
						case rtypes.NilType:
							inst.Set(name, (*rtypes.Instance)(nil))
							return 0
						case *rtypes.Instance:
							if value == nil {
								inst.Set(name, value)
								return 0
							}
							class := checkClassDesc(s, desc, propDesc.ValueType.Name, classDesc.Name, propDesc.Name)
							if class == nil {
								return 0
							}
							if !value.WithDescIsA(desc, class.Name) {
								return s.RaiseError("instance of class %s expected, got %s", class.Name, inst.ClassName)
							}
							inst.Set(name, value)
							return 0
						default:
							return s.RaiseError("Instance expected, got %s", value.Type())
						}
					case "Enum":
						enum := checkEnumDesc(s, desc, propDesc.ValueType.Name, classDesc.Name, propDesc.Name)
						if enum == nil {
							return 0
						}
						switch value := value.(type) {
						case types.Token:
							item := enum.Value(int(value))
							if item == nil {
								return s.RaiseError("invalid value %d for enum %s", value, enum.Name())
							}
							inst.Set(name, value)
							return 0
						case *rtypes.EnumItem:
							item := enum.Value(value.Value())
							if item == nil {
								return s.RaiseError(
									"invalid value %s (%d) for enum %s",
									value.String(),
									value.Value(),
									enum.String(),
								)
							}
							if a, b := enum.Name(), value.Enum().Name(); a != b {
								return s.RaiseError("expected enum %s, got %s", a, b)
							}
							if a, b := item.Name(), value.Name(); a != b {
								return s.RaiseError("expected enum item %s, got %s", a, b)
							}
							inst.Set(name, types.Token(item.Value()))
							return 0
						case types.Intlike:
							v := int(value.Intlike())
							item := enum.Value(v)
							if item == nil {
								return s.RaiseError("invalid value %d for enum %s", v, enum.Name())
							}
							inst.Set(name, types.Token(item.Value()))
							return 0
						case types.Numberlike:
							v := int(value.Numberlike())
							item := enum.Value(v)
							if item == nil {
								return s.RaiseError("invalid value %d for enum %s", v, enum.Name())
							}
							inst.Set(name, types.Token(item.Value()))
							return 0
						case types.Stringlike:
							v := value.Stringlike()
							item := enum.Item(v)
							if item == nil {
								return s.RaiseError("invalid value %s for enum %s", v, enum.Name())
							}
							inst.Set(name, types.Token(item.Value()))
							return 0
						default:
							return s.RaiseError("invalid value for enum %s", enum.Name())
						}
					default:
						if pt, vt := propDesc.ValueType.Name, value.Type(); vt != pt {
							// Attempt to convert value type to property type.
							rfl := s.Reflector(pt)
							if rfl.Name == "" || rfl.ConvertFrom == nil {
								return s.RaiseError("%s expected, got %s", pt, vt)
							}
							if value = rfl.ConvertFrom(value); value == nil {
								return s.RaiseError("%s expected, got %s", pt, vt)
							}
						}
					}
				}
				if _, ok := value.(rtypes.NilType); ok {
					inst.Set(name, nil)
					return 0
				}
				prop, ok := value.(types.PropValue)
				if !ok {
					return s.RaiseError("cannot assign %s as property", value.Type())
				}
				inst.Set(name, prop)
				return 0
			},
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case rtypes.NilType:
				return (*rtypes.Instance)(nil)
			case *rtypes.Instance:
				return v
			}
			return nil
		},
		Members: Members{
			"ClassName": Member{
				Get: func(s State, v types.Value) int {
					return s.Push(types.String(v.(*rtypes.Instance).ClassName))
				},
				// Allowed to be set for convenience.
				Set: func(s State, v types.Value) {
					inst := v.(*rtypes.Instance)
					if inst.IsDataModel() {
						s.RaiseError("%s cannot be assigned to", "ClassName")
						return
					}
					inst.ClassName = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("string")} },
			},
			"Name": Member{
				Get: func(s State, v types.Value) int {
					return s.Push(types.String(v.(*rtypes.Instance).Name()))
				},
				Set: func(s State, v types.Value) {
					v.(*rtypes.Instance).SetName(string(s.Pull(3, "string").(types.String)))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("string")} },
			},
			"Parent": Member{
				Get: func(s State, v types.Value) int {
					if parent := v.(*rtypes.Instance).Parent(); parent != nil {
						return s.Push(parent)
					}
					return s.Push(rtypes.Nil)
				},
				Set: func(s State, v types.Value) {
					var err error
					switch parent := s.PullAnyOf(3, "Instance", "nil").(type) {
					case *rtypes.Instance:
						err = v.(*rtypes.Instance).SetParent(parent)
					case nil:
						err = v.(*rtypes.Instance).SetParent(nil)
					}
					if err != nil {
						s.RaiseError("%s", err)
					}
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Optional{T: dt.Prim("Instance")}} },
			},
			"Descend": Member{Method: true,
				Get: func(s State, v types.Value) int {
					n := s.L.GetTop()
					if n-1 <= 0 {
						return s.RaiseError("expected at least 1 string")
					}
					names := make([]string, n-1)
					for i := 2; i <= n; i++ {
						names[i-2] = s.L.CheckString(i)
					}
					if child := v.(*rtypes.Instance).Descend(names...); child != nil {
						return s.Push(child)
					}
					return s.Push(rtypes.Nil)
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "...", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("Instance")}},
						},
					}
				},
			},
			"ClearAllChildren": Member{Method: true,
				Get: func(s State, v types.Value) int {
					v.(*rtypes.Instance).RemoveAll()
					return 0
				},
				Dump: func() dump.Value { return dump.Function{} },
			},
			"Clone": Member{Method: true,
				Get: func(s State, v types.Value) int {
					return s.Push(v.(*rtypes.Instance).Clone())
				},
				Dump: func() dump.Value {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Prim("Instance")},
						},
					}
				},
			},
			"Destroy": Member{Method: true,
				Get: func(s State, v types.Value) int {
					v.(*rtypes.Instance).SetParent(nil)
					return 0
				},
				Dump: func() dump.Value { return dump.Function{} },
			},
			"FindFirstAncestor": Member{Method: true,
				Get: func(s State, v types.Value) int {
					name := string(s.Pull(2, "string").(types.String))
					if ancestor := v.(*rtypes.Instance).FindFirstAncestor(name); ancestor != nil {
						return s.Push(ancestor)
					}
					return s.Push(rtypes.Nil)
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "name", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("Instance")}},
						},
					}
				},
			},
			"FindFirstAncestorOfClass": Member{Method: true,
				Get: func(s State, v types.Value) int {
					className := string(s.Pull(2, "string").(types.String))
					if ancestor := v.(*rtypes.Instance).FindFirstAncestorOfClass(className); ancestor != nil {
						return s.Push(ancestor)
					}
					return s.Push(rtypes.Nil)
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "className", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("Instance")}},
						},
					}
				},
			},
			"FindFirstAncestorWhichIsA": Member{Method: true,
				Get: func(s State, v types.Value) int {
					className := string(s.Pull(2, "string").(types.String))
					if ancestor := v.(*rtypes.Instance).FindFirstAncestorWhichIsA(className); ancestor != nil {
						return s.Push(ancestor)
					}
					return s.Push(rtypes.Nil)
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "className", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("Instance")}},
						},
					}
				},
			},
			"FindFirstChild": Member{Method: true,
				Get: func(s State, v types.Value) int {
					name := string(s.Pull(2, "string").(types.String))
					recurse := bool(s.PullOpt(3, "bool", types.False).(types.Bool))
					if child := v.(*rtypes.Instance).FindFirstChild(name, recurse); child != nil {
						return s.Push(child)
					}
					return s.Push(rtypes.Nil)
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "name", Type: dt.Prim("string")},
							{Name: "recurse", Type: dt.Optional{T: dt.Prim("bool")}},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("Instance")}},
						},
					}
				},
			},
			"FindFirstChildOfClass": Member{Method: true,
				Get: func(s State, v types.Value) int {
					className := string(s.Pull(2, "string").(types.String))
					recurse := bool(s.PullOpt(3, "bool", types.False).(types.Bool))
					if child := v.(*rtypes.Instance).FindFirstChildOfClass(className, recurse); child != nil {
						return s.Push(child)
					}
					return s.Push(rtypes.Nil)
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "className", Type: dt.Prim("string")},
							{Name: "recurse", Type: dt.Optional{T: dt.Prim("bool")}},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("Instance")}},
						},
					}
				},
			},
			"FindFirstChildWhichIsA": Member{Method: true,
				Get: func(s State, v types.Value) int {
					className := string(s.Pull(2, "string").(types.String))
					recurse := bool(s.PullOpt(3, "bool", types.False).(types.Bool))
					if child := v.(*rtypes.Instance).FindFirstChildWhichIsA(className, recurse); child != nil {
						return s.Push(child)
					}
					return s.Push(rtypes.Nil)
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "className", Type: dt.Prim("string")},
							{Name: "recurse", Type: dt.Optional{T: dt.Prim("bool")}},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("Instance")}},
						},
					}
				},
			},
			"GetAttribute": Member{Method: true,
				Get: func(s State, v types.Value) int {
					attribute := string(s.Pull(2, "string").(types.String))
					dict := getAttributes(s, v.(*rtypes.Instance))
					if v, ok := dict[attribute]; ok {
						return s.Push(v)
					}
					return s.Push(rtypes.Nil)
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "attribute", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("any")}},
						},
					}
				},
			},
			"GetAttributes": Member{Method: true,
				Get: func(s State, v types.Value) int {
					return s.Push(getAttributes(s, v.(*rtypes.Instance)))
				},
				Dump: func() dump.Value {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Prim("Dictionary")},
						},
					}
				},
			},
			"GetChildren": Member{Method: true,
				Get: func(s State, v types.Value) int {
					t := v.(*rtypes.Instance).Children()
					return s.Push(rtypes.Objects(t))
				},
				Dump: func() dump.Value {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Prim("Objects")},
						},
					}
				},
			},
			"GetDescendants": Member{Method: true,
				Get: func(s State, v types.Value) int {
					return s.Push(rtypes.Objects(v.(*rtypes.Instance).Descendants()))
				},
				Dump: func() dump.Value {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Prim("Objects")},
						},
					}
				},
			},
			"GetFullName": Member{Method: true,
				Get: func(s State, v types.Value) int {
					return s.Push(types.String(v.(*rtypes.Instance).GetFullName()))
				},
				Dump: func() dump.Value {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Prim("string")},
						},
					}
				},
			},
			"IsA": Member{Method: true,
				Get: func(s State, v types.Value) int {
					className := string(s.Pull(2, "string").(types.String))
					return s.Push(types.Bool(v.(*rtypes.Instance).IsA(className)))
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "className", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("bool")},
						},
					}
				},
			},
			"IsAncestorOf": Member{Method: true,
				Get: func(s State, v types.Value) int {
					descendant := s.Pull(2, "Instance").(*rtypes.Instance)
					return s.Push(types.Bool(v.(*rtypes.Instance).IsAncestorOf(descendant)))
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "descendant", Type: dt.Optional{T: dt.Prim("Instance")}},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("bool")},
						},
					}
				},
			},
			"IsDescendantOf": Member{Method: true,
				Get: func(s State, v types.Value) int {
					ancestor := s.Pull(2, "Instance").(*rtypes.Instance)
					return s.Push(types.Bool(v.(*rtypes.Instance).IsDescendantOf(ancestor)))
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "ancestor", Type: dt.Optional{T: dt.Prim("Instance")}},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("bool")},
						},
					}
				},
			},
			"SetAttribute": Member{Method: true,
				Get: func(s State, v types.Value) int {
					inst := v.(*rtypes.Instance)
					attribute := string(s.Pull(2, "string").(types.String))
					value := s.Pull(3, "Variant")
					dict := getAttributes(s, inst)
					dict[attribute] = value
					setAttributes(s, inst, dict)
					return 0
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "attribute", Type: dt.Prim("string")},
							{Name: "value", Type: dt.Optional{T: dt.Prim("any")}},
						},
					}
				},
			},
			"SetAttributes": Member{Method: true,
				Get: func(s State, v types.Value) int {
					dict := s.Pull(3, "Dictionary").(rtypes.Dictionary)
					setAttributes(s, v.(*rtypes.Instance), dict)
					return 0
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "attributes", Type: dt.Prim("Dictionary")},
						},
					}
				},
			},
		},
		Constructors: Constructors{
			"new": {
				Func: func(s State) int {
					className := string(s.Pull(1, "string").(types.String))
					parent, _ := s.PullOpt(2, "Instance", nil).(*rtypes.Instance)
					var desc *rtypes.RootDesc
					var blocked bool
					if s.L.GetTop() >= 3 {
						switch v := s.PullAnyOf(3, "RootDesc", "bool", "nil").(type) {
						case rtypes.NilType:
						case types.Bool:
							if v {
								return s.RaiseError("descriptor cannot be true")
							}
							blocked = true
						case *rtypes.RootDesc:
							desc = v
						}
					}
					if !blocked {
						checkDesc := desc
						if checkDesc == nil {
							// Use global descriptor, if available.
							checkDesc = s.Desc.Of(nil)
						}
						if checkDesc != nil {
							class := checkDesc.Classes[className]
							if class == nil {
								return s.RaiseError("unable to create instance of type %q", className)
							}
						}
					}
					inst := rtypes.NewInstance(className, parent)
					inst.SetDesc(desc, blocked)
					return s.Push(inst)
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Parameters: dump.Parameters{
							{Name: "className", Type: dt.Prim("string")},
							{Name: "parent", Type: dt.Optional{T: dt.Prim("Instance")}},
							{Name: "descriptor", Type: dt.Optional{T: dt.Group{T: dt.Or{dt.Prim("RootDesc"), dt.Prim("bool")}}}},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Instance")},
						},
					}}
				},
			},
		},
		Environment: func(s State, env *lua.LTable) {
			t := s.L.CreateTable(0, 1)
			t.RawSetString("new", s.L.NewFunction(func(l *lua.LState) int {
				var desc *rtypes.RootDesc
				var blocked bool
				if s.L.GetTop() >= 3 {
					switch v := s.PullAnyOf(3, "RootDesc", "bool", "nil").(type) {
					case rtypes.NilType:
					case types.Bool:
						if v {
							return s.RaiseError("descriptor cannot be true")
						}
						blocked = true
					case *rtypes.RootDesc:
						desc = v
					}
				}
				dataModel := rtypes.NewDataModel()
				dataModel.SetDesc(desc, blocked)
				return s.Push(dataModel)
			}))
			env.RawSetString("DataModel", t)
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Eq: true,
					Index: dump.Function{
						Parameters: dump.Parameters{
							{Name: "member", Type: dt.Or{dt.Prim("string"), dt.Prim("Symbol")}},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("any")}},
						},
					},
					Newindex: dump.Function{
						Parameters: dump.Parameters{
							{Name: "member", Type: dt.Or{dt.Prim("string"), dt.Prim("Symbol")}},
							{Name: "value", Type: dt.Optional{T: dt.Prim("any")}},
						},
					},
				},
			}
		},
	}
}
