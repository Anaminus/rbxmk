package reflect

import (
	"fmt"

	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func setUserdata(s State, t string) int {
	u := s.L.NewUserData()
	u.Value = s.Pull(1, t)
	s.L.SetMetatable(u, s.L.GetTypeMetatable(t))
	s.L.Push(u)
	return 1
}

func typesInt(s State) int {
	return setUserdata(s, "int")
}

// reflectPropertyTo behaves like ReflectVariantTo, except that exprims types
// are reflected as userdata.
func reflectPropertyTo(s State, v types.Value) (lv lua.LValue, err error) {
	switch v.(type) {
	case types.Int,
		types.Int64,
		types.Float,
		types.Double,
		types.Token,
		types.BinaryString,
		types.ProtectedString,
		types.Content,
		types.SharedString:
	default:
		return ReflectVariantTo(s, v)
	}
	typ := s.Type(v.Type())
	if typ.Name == "" {
		return nil, fmt.Errorf("unknown type %q", string(v.Type()))
	}
	if typ.ReflectTo == nil {
		return nil, fmt.Errorf("unable to cast %s to Variant", typ.Name)
	}
	u := s.L.NewUserData()
	u.Value = v
	s.L.SetMetatable(u, s.L.GetTypeMetatable(typ.Name))
	return u, nil
}

func Instance() Type {
	return Type{
		Name:        "Instance",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString(s.Pull(1, "Instance").(*rtypes.Instance).String()))
				return 1
			},
			"__eq": func(s State) int {
				op := s.Pull(2, "Instance").(*rtypes.Instance)
				return s.Push(types.Bool(s.Pull(1, "Instance").(*rtypes.Instance) == op))
			},
			"__index": func(s State) int {
				inst := s.Pull(1, "Instance").(*rtypes.Instance)

				// Try symbol.
				if typ := s.Type("Symbol"); typ.Name != "" {
					if sym, err := typ.ReflectFrom(s, typ, s.L.CheckAny(2)); err == nil {
						switch name := sym.(SymbolType).Name; name {
						case "Reference":
							return s.Push(types.String(inst.Reference))
						case "IsService":
							return s.Push(types.Bool(inst.IsService))
						default:
							s.L.RaiseError("symbol %s is not a valid member", name)
							return 0
						}
					}
				}

				name := string(s.Pull(2, "string").(types.String))
				// Try GetService.
				if inst.IsDataModel() && name == "GetService" {
					s.L.Push(s.L.NewFunction(func(l *lua.LState) int {
						u := l.CheckUserData(1)
						if u.Metatable != l.GetTypeMetatable("Instance") {
							TypeError(l, 1, "Instance")
							return 0
						}
						inst, ok := u.Value.(*rtypes.Instance)
						if !ok {
							TypeError(l, 1, "Instance")
							return 0
						}
						s := State{World: s.World, L: l}
						className := string(s.Pull(2, "string").(types.String))
						service := inst.FindFirstChildOfClass(className, false)
						if service == nil {
							service = rtypes.NewInstance(className)
							service.IsService = true
							service.SetName(className)
							service.SetParent(inst)
						}
						return s.Push(service)
					}))
					return 1
				}
				// Try property.
				value := inst.Get(name)
				if value == nil {
					// s.L.RaiseError("%s is not a valid member", name)
					return s.Push(rtypes.Nil)
				}
				lv, err := reflectPropertyTo(s, value)
				if err != nil {
					s.L.RaiseError(err.Error())
					return 0
				}
				s.L.Push(lv)
				return 1
			},
			"__newindex": func(s State) int {
				inst := s.Pull(1, "Instance").(*rtypes.Instance)

				// Try symbol.
				if typ := s.Type("Symbol"); typ.Name != "" {
					if sym, err := typ.ReflectFrom(s, typ, s.L.CheckAny(2)); err == nil {
						switch name := sym.(SymbolType).Name; name {
						case "Reference":
							value := string(s.Pull(3, "string").(types.String))
							inst.Reference = value
							return 0
						case "IsService":
							value := bool(s.Pull(3, "bool").(types.Bool))
							inst.IsService = value
							return 0
						default:
							s.L.RaiseError("symbol %s is not a valid member", name)
							return 0
						}
					}
				}

				name := string(s.Pull(2, "string").(types.String))
				// Try GetService.
				if inst.IsDataModel() && name == "GetService" {
					s.L.RaiseError("%s cannot be assigned to", name)
					return 0
				}
				// Try property.
				value := PullVariant(s, 3)
				prop, ok := value.(types.PropValue)
				if !ok {
					s.L.RaiseError("cannot assign %s as property", value.Type())
					return 0
				}
				inst.Set(name, prop)
				return 0
			},
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
						s.L.RaiseError("%s cannot be assigned to", "ClassName")
						return
					}
					inst.ClassName = string(s.Pull(3, "string").(types.String))
				},
			},
			"Name": Member{
				Get: func(s State, v types.Value) int {
					return s.Push(types.String(v.(*rtypes.Instance).Name()))
				},
				Set: func(s State, v types.Value) {
					v.(*rtypes.Instance).SetName(string(s.Pull(3, "string").(types.String)))
				},
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
						s.L.RaiseError(err.Error())
					}
				},
			},
			"ClearAllChildren": Member{Method: true, Get: func(s State, v types.Value) int {
				v.(*rtypes.Instance).RemoveAll()
				return 0
			}},
			"Clone": Member{Method: true, Get: func(s State, v types.Value) int {
				return s.Push(v.(*rtypes.Instance).Clone())
			}},
			"Destroy": Member{Method: true, Get: func(s State, v types.Value) int {
				v.(*rtypes.Instance).SetParent(nil)
				return 0
			}},
			"FindFirstAncestor": Member{Method: true, Get: func(s State, v types.Value) int {
				name := string(s.Pull(2, "string").(types.String))
				if ancestor := v.(*rtypes.Instance).FindFirstAncestorOfClass(name); ancestor != nil {
					return s.Push(ancestor)
				}
				return s.Push(rtypes.Nil)
			}},
			"FindFirstAncestorOfClass": Member{Method: true, Get: func(s State, v types.Value) int {
				className := string(s.Pull(2, "string").(types.String))
				if ancestor := v.(*rtypes.Instance).FindFirstAncestorOfClass(className); ancestor != nil {
					return s.Push(ancestor)
				}
				return s.Push(rtypes.Nil)
			}},
			"FindFirstChild": Member{Method: true, Get: func(s State, v types.Value) int {
				name := string(s.Pull(2, "string").(types.String))
				recurse := bool(s.PullOpt(3, "bool", types.False).(types.Bool))
				if child := v.(*rtypes.Instance).FindFirstChild(name, recurse); child != nil {
					return s.Push(child)
				}
				return s.Push(rtypes.Nil)
			}},
			"FindFirstChildOfClass": Member{Method: true, Get: func(s State, v types.Value) int {
				className := string(s.Pull(2, "string").(types.String))
				recurse := bool(s.PullOpt(3, "bool", types.False).(types.Bool))
				if child := v.(*rtypes.Instance).FindFirstChildOfClass(className, recurse); child != nil {
					return s.Push(child)
				}
				return s.Push(rtypes.Nil)
			}},
			"GetChildren": Member{Method: true, Get: func(s State, v types.Value) int {
				t := v.(*rtypes.Instance).Children()
				return s.Push(rtypes.Instances(t))
			}},
			"GetDescendants": Member{Method: true, Get: func(s State, v types.Value) int {
				return s.Push(rtypes.Instances(v.(*rtypes.Instance).Descendants()))
			}},
			"GetFullName": Member{Method: true, Get: func(s State, v types.Value) int {
				return s.Push(types.String(v.(*rtypes.Instance).GetFullName()))
			}},
			"IsAncestorOf": Member{Method: true, Get: func(s State, v types.Value) int {
				descendant := s.Pull(2, "Instance").(*rtypes.Instance)
				return s.Push(types.Bool(v.(*rtypes.Instance).IsAncestorOf(descendant)))
			}},
			"IsDescendantOf": Member{Method: true, Get: func(s State, v types.Value) int {
				ancestor := s.Pull(2, "Instance").(*rtypes.Instance)
				return s.Push(types.Bool(v.(*rtypes.Instance).IsDescendantOf(ancestor)))
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				className := string(s.Pull(1, "string").(types.String))
				inst := rtypes.NewInstance(className)
				return s.Push(inst)
			},
		},
		Environment: func(s State) {
			t := s.L.CreateTable(0, 1)
			t.RawSetString("new", s.L.NewFunction(func(l *lua.LState) int {
				dataModel := rtypes.NewDataModel()
				return s.Push(dataModel)
			}))
			s.L.SetGlobal("DataModel", t)
		},
	}
}
