package reflect

import (
	"bytes"
	"fmt"
	"strings"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/types"
)

// pushPropertyTo behaves like PushVariantTo, except that exprims types are
// reflected as userdata.
func pushPropertyTo(s rbxmk.State, v types.Value) (lv lua.LValue, err error) {
	rfl := s.Reflector(v.Type())
	if rfl.Name == "" {
		return nil, fmt.Errorf("unknown type %q", string(v.Type()))
	}
	if rfl.PushTo == nil {
		return nil, fmt.Errorf("unable to cast %s to Variant", rfl.Name)
	}
	if rfl.Flags&rbxmk.Exprim == 0 {
		return PushVariantTo(s.Context(), v)
	}
	u := s.UserDataOf(v, rfl.Name)
	return u, nil
}

func checkEnumDesc(desc *rtypes.RootDesc, name, class, prop string) (*rtypes.Enum, error) {
	var enumValue *rtypes.Enum
	if desc.EnumTypes != nil {
		enumValue = desc.EnumTypes.Enum(name)
	}
	if enumValue == nil {
		if desc.Enums[name] == nil {
			return nil, fmt.Errorf(
				"no enum descriptor %q for property descriptor %s.%s",
				name,
				class,
				prop,
			)
		}
		return nil, fmt.Errorf(
			"no enum value %q generated for property descriptor %s.%s",
			name,
			class,
			prop,
		)
	}
	return enumValue, nil
}

func checkClassDesc(desc *rtypes.RootDesc, name, class, prop string) (*rbxdump.Class, error) {
	classDesc := desc.Classes[name]
	if classDesc == nil {
		return nil, fmt.Errorf(
			"no class descriptor %q for property descriptor %s.%s",
			name,
			class,
			prop,
		)
	}
	return classDesc, nil
}

func defaultAttrConfig(inst *rtypes.Instance) rtypes.AttrConfig {
	attrcfg := inst.AttrConfig()
	if attrcfg != nil {
		return *attrcfg
	}
	return rtypes.AttrConfig{Property: "AttributesSerialize"}
}

func getAttributes(s rbxmk.State, inst *rtypes.Instance) rtypes.Dictionary {
	format := s.Format("rbxattr")
	if format.Name == "" {
		s.RaiseError("cannot decode attributes: format \"rbxattr\" not registered")
	}
	attrcfg := defaultAttrConfig(inst)
	v := inst.Get(attrcfg.Property)
	if v == nil {
		return rtypes.Dictionary{}
	}
	sv, ok := v.(types.Stringlike)
	if !ok {
		s.RaiseError("property %q is not string-like", attrcfg.Property)
		return nil
	}
	r := strings.NewReader(sv.Stringlike())
	dict, err := format.Decode(s.Global, nil, r)
	if err != nil {
		s.RaiseError("decode attributes from %q: %s", attrcfg.Property, err)
		return nil
	}
	return dict.(rtypes.Dictionary)
}

func setAttributes(s rbxmk.State, inst *rtypes.Instance, dict rtypes.Dictionary) {
	format := s.Format("rbxattr")
	if format.Name == "" {
		s.RaiseError("cannot encode attributes: format \"rbxattr\" not registered")
	}
	attrcfg := defaultAttrConfig(inst)
	var w bytes.Buffer
	if err := format.Encode(s.Global, nil, &w, dict); err != nil {
		s.RaiseError("encode attributes to %q: %s", attrcfg.Property, err)
		return
	}
	inst.Set(attrcfg.Property, types.BinaryString(w.Bytes()))
}

func reflectOne(s rbxmk.State, value types.Value) (lv lua.LValue, err error) {
	rfl := s.MustReflector(value.Type())
	lv, err = rfl.PushTo(s.Context(), value)
	if err != nil {
		return nil, err
	}
	if lv == nil {
		return lua.LNil, nil
	}
	return lv, nil
}

func getProperty(s rbxmk.State, inst *rtypes.Instance, desc *rtypes.RootDesc, name string) (lv lua.LValue, err error) {
	var classDesc *rbxdump.Class
	if desc != nil {
		classDesc = desc.Classes[inst.ClassName]
	}
	value := inst.Get(name)
	if classDesc != nil {
		propDesc := desc.Property(classDesc.Name, name)
		if propDesc == nil {
			return nil, fmt.Errorf("%s is not a valid member", name)
		}
		if value == nil {
			return nil, fmt.Errorf("property %s not initialized", name)
		}
		switch propDesc.ValueType.Category {
		case "Class":
			switch value := value.(type) {
			case *rtypes.Instance:
				if value == nil {
					return lua.LNil, nil
				}
				class, err := checkClassDesc(desc, propDesc.ValueType.Name, classDesc.Name, propDesc.Name)
				if err != nil {
					return nil, fmt.Errorf("%s", err)
				}
				if !value.WithDescIsA(desc, class.Name) {
					return nil, fmt.Errorf("instance of class %s expected, got %s", class.Name, value.ClassName)
				}
				return reflectOne(s, value)
			default:
				return nil, fmt.Errorf("stored value type %s is not an instance", value.Type())
			}
		case "Enum":
			enum, err := checkEnumDesc(desc, propDesc.ValueType.Name, classDesc.Name, propDesc.Name)
			if err != nil {
				return nil, err
			}
			token, ok := value.(types.Token)
			if !ok {
				return nil, fmt.Errorf("stored value type %s is not a token", value.Type())
			}
			item := enum.Value(int(token))
			if item == nil {
				return nil, fmt.Errorf("invalid stored value %d for enum %s", value, enum.Name())
			}
			return reflectOne(s, item)
		default:
			pt := propDesc.ValueType.Name
			opt := strings.HasSuffix(pt, "?")
			if opt {
				pt = strings.TrimSuffix(pt, "?")
				switch v := value.(type) {
				case rtypes.Optional:
					switch inner := v.Value().(type) {
					case nil:
						return reflectOne(s, rtypes.Nil)
					case types.PropValue:
						value = inner
					}
				}
			}
			rfl := s.Reflector(pt)
			if rfl.Name != "" && rfl.ConvertFrom != nil {
				if v := rfl.ConvertFrom(value); v != nil {
					return reflectOne(s, v)
				}
			}
			if vt := value.Type(); vt != pt {
				return nil, fmt.Errorf("stored value type %s does not match property type %s", vt, pt)
			}
		}
		// Push without converting exprims.
		return PushVariantTo(s.Context(), value)
	}
	if value == nil {
		// Fallback to nil.
		return lua.LNil, nil
	}
	return pushPropertyTo(s, value)
}

func canSetProperty(s rbxmk.State, inst *rtypes.Instance, desc *rtypes.RootDesc, name string, value types.Value) (types.PropValue, error) {
	var classDesc *rbxdump.Class
	if desc != nil {
		classDesc = desc.Classes[inst.ClassName]
	}
	if classDesc != nil {
		propDesc := desc.Property(classDesc.Name, name)
		if propDesc == nil {
			return nil, fmt.Errorf("%s is not a valid member", name)
		}
		switch propDesc.ValueType.Category {
		case "Class":
			switch value := value.(type) {
			case rtypes.NilType:
				return (*rtypes.Instance)(nil), nil
			case *rtypes.Instance:
				if value == nil {
					return value, nil
				}
				class, err := checkClassDesc(desc, propDesc.ValueType.Name, classDesc.Name, propDesc.Name)
				if err != nil {
					return nil, err
				}
				if !value.WithDescIsA(desc, class.Name) {
					return nil, fmt.Errorf("instance of class %s expected, got %s", class.Name, inst.ClassName)
				}
				return value, nil
			default:
				return nil, fmt.Errorf("Instance expected, got %s", value.Type())
			}
		case "Enum":
			enum, err := checkEnumDesc(desc, propDesc.ValueType.Name, classDesc.Name, propDesc.Name)
			if err != nil {
				return nil, err
			}
			switch value := value.(type) {
			case types.Token:
				item := enum.Value(int(value))
				if item == nil {
					return nil, fmt.Errorf("invalid value %d for enum %s", value, enum.Name())
				}
				return value, nil
			case *rtypes.EnumItem:
				item := enum.Value(value.Value())
				if item == nil {
					return nil, fmt.Errorf(
						"invalid value %s (%d) for enum %s",
						value.String(),
						value.Value(),
						enum.String(),
					)
				}
				if a, b := enum.Name(), value.Enum().Name(); a != b {
					return nil, fmt.Errorf("expected enum %s, got %s", a, b)
				}
				if a, b := item.Name(), value.Name(); a != b {
					return nil, fmt.Errorf("expected enum item %s, got %s", a, b)
				}
				return types.Token(item.Value()), nil
			case types.Intlike:
				v := int(value.Intlike())
				item := enum.Value(v)
				if item == nil {
					return nil, fmt.Errorf("invalid value %d for enum %s", v, enum.Name())
				}
				return types.Token(item.Value()), nil
			case types.Numberlike:
				v := int(value.Numberlike())
				item := enum.Value(v)
				if item == nil {
					return nil, fmt.Errorf("invalid value %d for enum %s", v, enum.Name())
				}
				return types.Token(item.Value()), nil
			case types.Stringlike:
				v := value.Stringlike()
				item := enum.Item(v)
				if item == nil {
					return nil, fmt.Errorf("invalid value %s for enum %s", v, enum.Name())
				}
				return types.Token(item.Value()), nil
			default:
				return nil, fmt.Errorf("invalid value for enum %s", enum.Name())
			}
		default:
			pt := propDesc.ValueType.Name
			vt := value.Type()
			opt := strings.HasSuffix(pt, "?")
			if opt {
				pt = strings.TrimSuffix(pt, "?")
				switch v := value.(type) {
				case rtypes.NilType:
					// Convert nil to None of property type.
					return rtypes.None(pt), nil
				case rtypes.Optional:
					inner := v.Value()
					if inner == nil {
						// Returning rtypes.None(pt) here would have the effect
						// of converting None of any type to None of property
						// type.

						// Attempt to convert value as-is. Set opt to false to
						// prevent reboxing.
						opt = false
					} else {
						value = inner
					}
					vt = v.ValueType()
				}
				// value, pt, and vt are unboxed; can be inspected as usual.
			}
			if vt != pt {
				// Attempt to convert value type to property type.
				rfl := s.Reflector(pt)
				if rfl.Name == "" || rfl.ConvertFrom == nil {
					return nil, fmt.Errorf("%s expected, got %s", pt, vt)
				}
				if value = rfl.ConvertFrom(value); value == nil {
					return nil, fmt.Errorf("%s expected, got %s", pt, vt)
				}
			}
			if opt {
				// Rebox value into optional.
				value = rtypes.Some(value)
			}
		}
	}
	if _, ok := value.(rtypes.NilType); ok {
		return nil, nil
	}
	prop, ok := value.(types.PropValue)
	if !ok {
		return nil, fmt.Errorf("cannot assign %s as property", value.Type())
	}
	return prop, nil
}

func setProperty(s rbxmk.State, inst *rtypes.Instance, desc *rtypes.RootDesc, name string, value types.Value) error {
	prop, err := canSetProperty(s, inst, desc, name, value)
	if err != nil {
		return err
	}
	inst.Set(name, prop)
	return nil
}

func getService(s rbxmk.State) int {
	inst := s.Pull(1, "Instance").(*rtypes.Instance)
	className := string(s.Pull(2, "string").(types.String))
	if desc := s.Desc.Of(inst); desc != nil {
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
}

func init() { register(Instance) }
func Instance() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "Instance",
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			if inst, ok := v.(*rtypes.Instance); ok && inst == nil {
				return lua.LNil, nil
			}
			u := c.UserDataOf(v, "Instance")
			return u, nil
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			switch lv := lv.(type) {
			case *lua.LNilType:
				return (*rtypes.Instance)(nil), nil
			case *lua.LUserData:
				if lv.Metatable != c.GetTypeMetatable("Instance") {
					return nil, rbxmk.TypeError{Want: "Instance", Got: lv.Type().String()}
				}
				v, ok := lv.Value().(types.Value)
				if !ok {
					return nil, rbxmk.TypeError{Want: "Instance", Got: lv.Type().String()}
				}
				return v, nil
			default:
				return nil, rbxmk.TypeError{Want: "Instance", Got: lv.Type().String()}
			}
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case **rtypes.Instance:
				*p = v.(*rtypes.Instance)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "Instance").(*rtypes.Instance)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__index": func(s rbxmk.State) int {
				inst := s.Pull(1, "Instance").(*rtypes.Instance)
				name := string(s.Pull(2, "string").(types.String))

				// Try GetService.
				if inst.IsDataModel() && name == "GetService" {
					s.L.Push(s.WrapMethod(getService))
					return 1
				}

				// Try property.
				lv, err := getProperty(s, inst, s.Desc.Of(inst), name)
				if err != nil {
					return s.RaiseError("%s", err)
				}
				s.L.Push(lv)
				return 1
			},
			"__newindex": func(s rbxmk.State) int {
				inst := s.Pull(1, "Instance").(*rtypes.Instance)
				name := string(s.Pull(2, "string").(types.String))

				// Try GetService.
				if inst.IsDataModel() && name == "GetService" {
					return s.RaiseError("%s cannot be assigned to", name)
				}

				// Try property.
				value := PullVariant(s, 3)
				if err := setProperty(s, inst, s.Desc.Of(inst), name, value); err != nil {
					return s.RaiseError("%s", err)
				}
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
		Properties: rbxmk.Properties{
			"ClassName": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.String(v.(*rtypes.Instance).ClassName))
				},
				// Allowed to be set for convenience.
				Set: func(s rbxmk.State, v types.Value) {
					inst := v.(*rtypes.Instance)
					if inst.IsDataModel() {
						s.RaiseError("%s cannot be assigned to", "ClassName")
						return
					}
					className := string(s.Pull(3, "string").(types.String))
					if className == "DataModel" {
						s.RaiseError("cannot set ClassName to DataModel")
						return
					}
					inst.ClassName = className
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						Summary:     "Types/Instance:Properties/ClassName/Summary",
						Description: "Types/Instance:Properties/ClassName/Description",
					}
				},
			},
			"Name": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.String(v.(*rtypes.Instance).Name()))
				},
				Set: func(s rbxmk.State, v types.Value) {
					v.(*rtypes.Instance).SetName(string(s.Pull(3, "string").(types.String)))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						Summary:     "Types/Instance:Properties/Name/Summary",
						Description: "Types/Instance:Properties/Name/Description",
					}
				},
			},
			"Parent": {
				Get: func(s rbxmk.State, v types.Value) int {
					if parent := v.(*rtypes.Instance).Parent(); parent != nil {
						return s.Push(parent)
					}
					return s.Push(rtypes.Nil)
				},
				Set: func(s rbxmk.State, v types.Value) {
					var err error
					switch parent := s.PullAnyOf(3, "Instance", "nil").(type) {
					case *rtypes.Instance:
						err = v.(*rtypes.Instance).SetParent(parent)
					case nil:
						err = v.(*rtypes.Instance).SetParent(nil)
					default:
						s.ReflectorError(3)
					}
					if err != nil {
						s.RaiseError("%s", err)
					}
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Optional{T: dt.Prim("Instance")},
						Summary:     "Types/Instance:Properties/Parent/Summary",
						Description: "Types/Instance:Properties/Parent/Description",
					}
				},
			},
		},
		Symbols: rbxmk.Symbols{
			rtypes.Symbol{Name: "Reference"}: {
				Get: func(s rbxmk.State, v types.Value) int {
					inst := v.(*rtypes.Instance)
					return s.Push(types.String(inst.Reference))
				},
				Set: func(s rbxmk.State, v types.Value) {
					inst := v.(*rtypes.Instance)
					inst.Reference = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						Summary:     "Types/Instance:Symbols/Reference/Summary",
						Description: "Types/Instance:Symbols/Reference/Description",
					}
				},
			},
			rtypes.Symbol{Name: "IsService"}: {
				Get: func(s rbxmk.State, v types.Value) int {
					inst := v.(*rtypes.Instance)
					return s.Push(types.Bool(inst.IsService))
				},
				Set: func(s rbxmk.State, v types.Value) {
					inst := v.(*rtypes.Instance)
					inst.IsService = bool(s.Pull(3, "bool").(types.Bool))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("bool"),
						Summary:     "Types/Instance:Symbols/IsService/Summary",
						Description: "Types/Instance:Symbols/IsService/Description",
					}
				},
			},
			rtypes.Symbol{Name: "Desc"}: {
				Get: func(s rbxmk.State, v types.Value) int {
					inst := v.(*rtypes.Instance)
					desc := inst.Desc()
					if desc == nil {
						return s.Push(rtypes.Nil)
					}
					return s.Push(desc)
				},
				Set: func(s rbxmk.State, v types.Value) {
					inst := v.(*rtypes.Instance)
					switch v := s.PullAnyOf(3, "RootDesc", "bool", "nil").(type) {
					case *rtypes.RootDesc:
						inst.SetDesc(v, false)
					case types.Bool:
						if v {
							s.RaiseError("descriptor cannot be true")
							return
						}
						inst.SetDesc(nil, true)
					case rtypes.NilType:
						inst.SetDesc(nil, false)
					default:
						s.ReflectorError(3)
					}
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType: dt.Or{
							dt.Prim("RootDesc"),
							dt.Prim("bool"),
							dt.Prim("nil"),
						},
						Summary:     "Types/Instance:Symbols/Desc/Summary",
						Description: "Types/Instance:Symbols/Desc/Description",
					}
				},
			},
			rtypes.Symbol{Name: "RawDesc"}: {
				Get: func(s rbxmk.State, v types.Value) int {
					inst := v.(*rtypes.Instance)
					desc, blocked := inst.RawDesc()
					if blocked {
						return s.Push(types.False)
					}
					if desc == nil {
						return s.Push(rtypes.Nil)
					}
					return s.Push(desc)
				},
				Set: func(s rbxmk.State, v types.Value) {
					inst := v.(*rtypes.Instance)
					switch v := s.PullAnyOf(3, "RootDesc", "bool", "nil").(type) {
					case *rtypes.RootDesc:
						inst.SetDesc(v, false)
					case types.Bool:
						if v {
							s.RaiseError("descriptor cannot be true")
							return
						}
						inst.SetDesc(nil, true)
					case rtypes.NilType:
						inst.SetDesc(nil, false)
					default:
						s.ReflectorError(3)
					}
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType: dt.Or{
							dt.Prim("RootDesc"),
							dt.Prim("bool"),
							dt.Prim("nil"),
						},
						Summary:     "Types/Instance:Symbols/RawDesc/Summary",
						Description: "Types/Instance:Symbols/RawDesc/Description",
					}
				},
			},
			rtypes.Symbol{Name: "AttrConfig"}: {
				Get: func(s rbxmk.State, v types.Value) int {
					inst := v.(*rtypes.Instance)
					attrcfg := inst.AttrConfig()
					if attrcfg == nil {
						return s.Push(rtypes.Nil)
					}
					return s.Push(attrcfg)
				},
				Set: func(s rbxmk.State, v types.Value) {
					inst := v.(*rtypes.Instance)
					switch v := s.PullAnyOf(3, "AttrConfig", "bool", "nil").(type) {
					case *rtypes.AttrConfig:
						inst.SetAttrConfig(v, false)
					case types.Bool:
						if v {
							s.RaiseError("AttrConfig cannot be true")
							return
						}
						inst.SetAttrConfig(nil, true)
					case rtypes.NilType:
						inst.SetAttrConfig(nil, false)
					default:
						s.ReflectorError(3)
					}
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType: dt.Or{
							dt.Prim("AttrConfig"),
							dt.Prim("bool"),
							dt.Prim("nil"),
						},
						Summary:     "Types/Instance:Symbols/AttrConfig/Summary",
						Description: "Types/Instance:Symbols/AttrConfig/Description",
					}
				},
			},
			rtypes.Symbol{Name: "RawAttrConfig"}: {
				Get: func(s rbxmk.State, v types.Value) int {
					inst := v.(*rtypes.Instance)
					attrcfg, blocked := inst.RawAttrConfig()
					if blocked {
						return s.Push(types.False)
					}
					if attrcfg == nil {
						return s.Push(rtypes.Nil)
					}
					return s.Push(attrcfg)
				},
				Set: func(s rbxmk.State, v types.Value) {
					inst := v.(*rtypes.Instance)
					switch v := s.PullAnyOf(3, "AttrConfig", "bool", "nil").(type) {
					case *rtypes.AttrConfig:
						inst.SetAttrConfig(v, false)
					case types.Bool:
						if v {
							s.RaiseError("AttrConfig cannot be true")
							return
						}
						inst.SetAttrConfig(nil, true)
					case rtypes.NilType:
						inst.SetAttrConfig(nil, false)
					default:
						s.ReflectorError(3)
					}
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType: dt.Or{
							dt.Prim("AttrConfig"),
							dt.Prim("bool"),
							dt.Prim("nil"),
						},
						Summary:     "Types/Instance:Symbols/RawAttrConfig/Summary",
						Description: "Types/Instance:Symbols/RawAttrConfig/Description",
					}
				},
			},
			rtypes.Symbol{Name: "Properties"}: {
				Get: func(s rbxmk.State, v types.Value) int {
					inst := v.(*rtypes.Instance)
					desc := s.Desc.Of(inst)
					props := inst.PropertyNames()
					dict := s.L.CreateTable(0, len(props))
					for _, name := range props {
						if value, err := getProperty(s, inst, desc, name); err == nil {
							dict.RawSetString(name, value)
						}
					}
					s.L.Push(dict)
					return 1
				},
				Set: func(s rbxmk.State, v types.Value) {
					inst := v.(*rtypes.Instance)
					desc := s.Desc.Of(inst)
					dict := s.PullAnyOf(3, "Dictionary").(rtypes.Dictionary)
					props := make(map[string]types.PropValue, len(dict))
					for name, value := range dict {
						prop, err := canSetProperty(s, inst, desc, name, value)
						if err != nil {
							s.RaiseError("%s", err)
							return
						}
						props[name] = prop
					}
					inst.SetProperties(props, true)
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("Dictionary"),
						Summary:     "Types/Instance:Symbols/Properties/Summary",
						Description: "Types/Instance:Symbols/Properties/Description",
					}
				},
			},
			rtypes.Symbol{Name: "Metadata"}: {
				Get: func(s rbxmk.State, v types.Value) int {
					inst := v.(*rtypes.Instance)
					if meta := inst.Metadata(); meta != nil {
						dict := make(rtypes.Dictionary, len(meta))
						for k, v := range meta {
							dict[k] = types.String(v)
						}
						return s.Push(dict)
					}
					return s.RaiseError("symbol Metadata is not a valid member of Instance")
				},
				Set: func(s rbxmk.State, v types.Value) {
					inst := v.(*rtypes.Instance)
					if meta := inst.Metadata(); meta != nil {
						dict := s.Pull(3, "Dictionary").(rtypes.Dictionary)
						for k := range meta {
							delete(meta, k)
						}
						for k, v := range dict {
							w, ok := v.(types.String)
							if !ok {
								s.RaiseError("field %q: string expected, got %s (%T)", k, v.Type(), v)
								return
							}
							meta[k] = string(w)
						}
						return
					}
					s.RaiseError("symbol Metadata is not a valid member of Instance")
				},
			},
		},
		Methods: rbxmk.Methods{
			"Descend": {
				Func: func(s rbxmk.State, v types.Value) int {
					n := s.Count()
					if n-1 <= 0 {
						return s.RaiseError("expected at least 1 string")
					}
					names := make([]string, n-1)
					for i := 2; i <= n; i++ {
						names[i-2] = s.CheckString(i)
					}
					if child := v.(*rtypes.Instance).Descend(names...); child != nil {
						return s.Push(child)
					}
					return s.Push(rtypes.Nil)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "...", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("Instance")}},
						},
						Summary:     "Types/Instance:Methods/Descend/Summary",
						Description: "Types/Instance:Methods/Descend/Description",
					}
				},
			},
			"ClearAllChildren": {
				Func: func(s rbxmk.State, v types.Value) int {
					v.(*rtypes.Instance).RemoveAll()
					return 0
				},
				Dump: func() dump.Function {
					return dump.Function{
						Summary:     "Types/Instance:Methods/ClearAllChildren/Summary",
						Description: "Types/Instance:Methods/ClearAllChildren/Description",
					}
				},
			},
			"Clone": {
				Func: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(*rtypes.Instance).Clone())
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Prim("Instance")},
						},
						Summary:     "Types/Instance:Methods/Clone/Summary",
						Description: "Types/Instance:Methods/Clone/Description",
					}
				},
			},
			"Destroy": {
				Func: func(s rbxmk.State, v types.Value) int {
					v.(*rtypes.Instance).SetParent(nil)
					return 0
				},
				Dump: func() dump.Function {
					return dump.Function{
						Summary:     "Types/Instance:Methods/Destroy/Summary",
						Description: "Types/Instance:Methods/Destroy/Description",
					}
				},
			},
			"FindFirstAncestor": {
				Func: func(s rbxmk.State, v types.Value) int {
					name := string(s.Pull(2, "string").(types.String))
					if ancestor := v.(*rtypes.Instance).FindFirstAncestor(name); ancestor != nil {
						return s.Push(ancestor)
					}
					return s.Push(rtypes.Nil)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "name", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("Instance")}},
						},
						Summary:     "Types/Instance:Methods/FindFirstAncestor/Summary",
						Description: "Types/Instance:Methods/FindFirstAncestor/Description",
					}
				},
			},
			"FindFirstAncestorOfClass": {
				Func: func(s rbxmk.State, v types.Value) int {
					className := string(s.Pull(2, "string").(types.String))
					if ancestor := v.(*rtypes.Instance).FindFirstAncestorOfClass(className); ancestor != nil {
						return s.Push(ancestor)
					}
					return s.Push(rtypes.Nil)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "className", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("Instance")}},
						},
						Summary:     "Types/Instance:Methods/FindFirstAncestorOfClass/Summary",
						Description: "Types/Instance:Methods/FindFirstAncestorOfClass/Description",
					}
				},
			},
			"FindFirstAncestorWhichIsA": {
				Func: func(s rbxmk.State, v types.Value) int {
					className := string(s.Pull(2, "string").(types.String))
					if ancestor := v.(*rtypes.Instance).FindFirstAncestorWhichIsA(className); ancestor != nil {
						return s.Push(ancestor)
					}
					return s.Push(rtypes.Nil)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "className", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("Instance")}},
						},
						Summary:     "Types/Instance:Methods/FindFirstAncestorWhichIsA/Summary",
						Description: "Types/Instance:Methods/FindFirstAncestorWhichIsA/Description",
					}
				},
			},
			"FindFirstChild": {
				Func: func(s rbxmk.State, v types.Value) int {
					name := string(s.Pull(2, "string").(types.String))
					recurse := bool(s.PullOpt(3, types.False, "bool").(types.Bool))
					if child := v.(*rtypes.Instance).FindFirstChild(name, recurse); child != nil {
						return s.Push(child)
					}
					return s.Push(rtypes.Nil)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "name", Type: dt.Prim("string")},
							{Name: "recurse", Type: dt.Optional{T: dt.Prim("bool")}},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("Instance")}},
						},
						Summary:     "Types/Instance:Methods/FindFirstChild/Summary",
						Description: "Types/Instance:Methods/FindFirstChild/Description",
					}
				},
			},
			"FindFirstChildOfClass": {
				Func: func(s rbxmk.State, v types.Value) int {
					className := string(s.Pull(2, "string").(types.String))
					recurse := bool(s.PullOpt(3, types.False, "bool").(types.Bool))
					if child := v.(*rtypes.Instance).FindFirstChildOfClass(className, recurse); child != nil {
						return s.Push(child)
					}
					return s.Push(rtypes.Nil)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "className", Type: dt.Prim("string")},
							{Name: "recurse", Type: dt.Optional{T: dt.Prim("bool")}},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("Instance")}},
						},
						Summary:     "Types/Instance:Methods/FindFirstChildOfClass/Summary",
						Description: "Types/Instance:Methods/FindFirstChildOfClass/Description",
					}
				},
			},
			"FindFirstChildWhichIsA": {
				Func: func(s rbxmk.State, v types.Value) int {
					className := string(s.Pull(2, "string").(types.String))
					recurse := bool(s.PullOpt(3, types.False, "bool").(types.Bool))
					if child := v.(*rtypes.Instance).FindFirstChildWhichIsA(className, recurse); child != nil {
						return s.Push(child)
					}
					return s.Push(rtypes.Nil)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "className", Type: dt.Prim("string")},
							{Name: "recurse", Type: dt.Optional{T: dt.Prim("bool")}},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("Instance")}},
						},
						Summary:     "Types/Instance:Methods/FindFirstChildWhichIsA/Summary",
						Description: "Types/Instance:Methods/FindFirstChildWhichIsA/Description",
					}
				},
			},
			"GetAttribute": {
				Func: func(s rbxmk.State, v types.Value) int {
					attribute := string(s.Pull(2, "string").(types.String))
					dict := getAttributes(s, v.(*rtypes.Instance))
					if v, ok := dict[attribute]; ok {
						return s.Push(v)
					}
					return s.Push(rtypes.Nil)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "attribute", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("any")}},
						},
						Summary:     "Types/Instance:Methods/GetAttribute/Summary",
						Description: "Types/Instance:Methods/GetAttribute/Description",
					}
				},
			},
			"GetAttributes": {
				Func: func(s rbxmk.State, v types.Value) int {
					return s.Push(getAttributes(s, v.(*rtypes.Instance)))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Prim("Dictionary")},
						},
						Summary:     "Types/Instance:Methods/GetAttributes/Summary",
						Description: "Types/Instance:Methods/GetAttributes/Description",
					}
				},
			},
			"GetChildren": {
				Func: func(s rbxmk.State, v types.Value) int {
					t := v.(*rtypes.Instance).Children()
					return s.Push(rtypes.Objects(t))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Prim("Objects")},
						},
						Summary:     "Types/Instance:Methods/GetChildren/Summary",
						Description: "Types/Instance:Methods/GetChildren/Description",
					}
				},
			},
			"GetDescendants": {
				Func: func(s rbxmk.State, v types.Value) int {
					return s.Push(rtypes.Objects(v.(*rtypes.Instance).Descendants()))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Prim("Objects")},
						},
						Summary:     "Types/Instance:Methods/GetDescendants/Summary",
						Description: "Types/Instance:Methods/GetDescendants/Description",
					}
				},
			},
			"GetFullName": {
				Func: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.String(v.(*rtypes.Instance).GetFullName()))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Prim("string")},
						},
						Summary:     "Types/Instance:Methods/GetFullName/Summary",
						Description: "Types/Instance:Methods/GetFullName/Description",
					}
				},
			},
			"IsA": {
				Func: func(s rbxmk.State, v types.Value) int {
					className := string(s.Pull(2, "string").(types.String))
					return s.Push(types.Bool(v.(*rtypes.Instance).IsA(className)))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "className", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("bool")},
						},
						Summary:     "Types/Instance:Methods/IsA/Summary",
						Description: "Types/Instance:Methods/IsA/Description",
					}
				},
			},
			"IsAncestorOf": {
				Func: func(s rbxmk.State, v types.Value) int {
					descendant := s.Pull(2, "Instance").(*rtypes.Instance)
					return s.Push(types.Bool(v.(*rtypes.Instance).IsAncestorOf(descendant)))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "descendant", Type: dt.Optional{T: dt.Prim("Instance")}},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("bool")},
						},
						Summary:     "Types/Instance:Methods/IsAncestorOf/Summary",
						Description: "Types/Instance:Methods/IsAncestorOf/Description",
					}
				},
			},
			"IsDescendantOf": {
				Func: func(s rbxmk.State, v types.Value) int {
					ancestor := s.Pull(2, "Instance").(*rtypes.Instance)
					return s.Push(types.Bool(v.(*rtypes.Instance).IsDescendantOf(ancestor)))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "ancestor", Type: dt.Optional{T: dt.Prim("Instance")}},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("bool")},
						},
						Summary:     "Types/Instance:Methods/IsDescendantOf/Summary",
						Description: "Types/Instance:Methods/IsDescendantOf/Description",
					}
				},
			},
			"SetAttribute": {
				Func: func(s rbxmk.State, v types.Value) int {
					inst := v.(*rtypes.Instance)
					attribute := string(s.Pull(2, "string").(types.String))
					value := s.Pull(3, "Variant")
					dict := getAttributes(s, inst)
					if value == nil {
						delete(dict, attribute)
					} else {
						dict[attribute] = value
					}
					setAttributes(s, inst, dict)
					return 0
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "attribute", Type: dt.Prim("string")},
							{Name: "value", Type: dt.Optional{T: dt.Prim("any")}},
						},
						Summary:     "Types/Instance:Methods/SetAttribute/Summary",
						Description: "Types/Instance:Methods/SetAttribute/Description",
					}
				},
			},
			"SetAttributes": {
				Func: func(s rbxmk.State, v types.Value) int {
					dict := s.Pull(2, "Dictionary").(rtypes.Dictionary)
					setAttributes(s, v.(*rtypes.Instance), dict)
					return 0
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "attributes", Type: dt.Prim("Dictionary")},
						},
						Summary:     "Types/Instance:Methods/SetAttributes/Summary",
						Description: "Types/Instance:Methods/SetAttributes/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					className := string(s.Pull(1, "string").(types.String))
					parent, _ := s.PullOpt(2, nil, "Instance").(*rtypes.Instance)
					if className == "DataModel" && parent != nil {
						return s.RaiseError("DataModel Parent must be nil")
					}

					var desc *rtypes.RootDesc
					var blocked bool
					if s.Count() >= 3 {
						switch v := s.PullAnyOf(3, "RootDesc", "bool", "nil").(type) {
						case rtypes.NilType:
						case types.Bool:
							if v {
								return s.RaiseError("descriptor cannot be true")
							}
							blocked = true
						case *rtypes.RootDesc:
							desc = v
						default:
							return s.ReflectorError(3)
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
					var inst *rtypes.Instance
					if className == "DataModel" {
						inst = rtypes.NewDataModel()
					} else {
						inst = rtypes.NewInstance(className, parent)
					}
					inst.SetDesc(desc, blocked)
					return s.Push(inst)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "className", Type: dt.Prim("string")},
								{Name: "parent", Type: dt.Optional{T: dt.Prim("Instance")}},
								{Name: "descriptor", Type: dt.Optional{T: dt.Group{T: dt.Or{dt.Prim("RootDesc"), dt.Prim("bool")}}}},
							},
							Returns: dump.Parameters{
								{Type: dt.Or{dt.Prim("Instance"), dt.Prim("DataModel")}},
							},
							CanError:    true,
							Summary:     "Types/Instance:Constructors/new/Summary",
							Description: "Types/Instance:Constructors/new/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Index: &dump.Function{
						Parameters: dump.Parameters{
							{Name: "property", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("any")}},
						},
						CanError:    true,
						Summary:     "Types/Instance:Operators/Index/Summary",
						Description: "Types/Instance:Operators/Index/Description",
					},
					Newindex: &dump.Function{
						Parameters: dump.Parameters{
							{Name: "property", Type: dt.Prim("string")},
							{Name: "value", Type: dt.Optional{T: dt.Prim("any")}},
						},
						CanError:    true,
						Summary:     "Types/Instance:Operators/Newindex/Summary",
						Description: "Types/Instance:Operators/Newindex/Description",
					},
				},
				Summary:     "Types/Instance:Summary",
				Description: "Types/Instance:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			AttrConfig,
			Bool,
			Dictionary,
			Nil,
			Objects,
			Optional,
			RootDesc,
			String,
			Symbol,
			Variant,
		},
	}
}
