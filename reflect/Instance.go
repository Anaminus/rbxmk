package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/types"
	"github.com/yuin/gopher-lua"
)

func Instance() Type {
	return Type{
		Name:        "Instance",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				s.L.Push(lua.LString(v.(*types.Instance).String()))
				return 1
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "Instance").(*types.Instance)
				return s.Push("bool", v.(*types.Instance) == op)
			},
			"__index": func(s State, v Value) int {
				inst := v.(*types.Instance)

				// Try symbol.
				if typ := s.Type("Symbol"); typ.Name != "" {
					if sym, err := typ.ReflectFrom(s, typ, s.L.CheckAny(2)); err == nil {
						switch name := sym.(SymbolType).Name; name {
						case "Reference":
							return s.Push("string", inst.Reference)
						case "IsService":
							return s.Push("bool", inst.IsService)
						default:
							s.L.RaiseError("symbol %s is not a valid member", name)
							return 0
						}
					}
				}

				name := s.Pull(2, "string").(string)
				// Try GetService.
				if inst.IsDataModel() && name == "GetService" {
					s.L.Push(s.L.NewFunction(func(l *lua.LState) int {
						u := l.CheckUserData(1)
						if u.Metatable != l.GetTypeMetatable("Instance") {
							TypeError(l, 1, "Instance")
							return 0
						}
						inst, ok := u.Value.(*types.Instance)
						if !ok {
							TypeError(l, 1, "Instance")
							return 0
						}
						s := State{World: s.World, L: l}
						className := s.Pull(2, "string").(string)
						service := inst.FindFirstChildOfClass(className, false)
						if service == nil {
							service = types.NewInstance(className)
							service.IsService = true
							service.SetName(className)
							service.SetParent(inst)
						}
						return s.Push("Instance", service)
					}))
					return 1
				}
				// Try property.
				value := inst.Get(name)
				if value.Type == "" {
					// s.L.RaiseError("%s is not a valid member", name)
					return s.Push("nil", nil)
				}
				return s.Push(value.Type, value.Value)
			},
			"__newindex": func(s State, v Value) int {
				inst := v.(*types.Instance)

				// Try symbol.
				if typ := s.Type("Symbol"); typ.Name != "" {
					if sym, err := typ.ReflectFrom(s, typ, s.L.CheckAny(2)); err == nil {
						switch name := sym.(SymbolType).Name; name {
						case "Reference":
							value := s.Pull(3, "string").(string)
							inst.Reference = value
							return 0
						case "IsService":
							value := s.Pull(3, "bool").(bool)
							inst.IsService = value
							return 0
						default:
							s.L.RaiseError("symbol %s is not a valid member", name)
							return 0
						}
					}
				}

				name := s.Pull(2, "string").(string)
				// Try GetService.
				if inst.IsDataModel() && name == "GetService" {
					s.L.RaiseError("%s cannot be assigned to", name)
					return 0
				}
				// Try property.
				value, typ := PullVariant(s, 3)
				inst.Set(name, TValue{Type: typ.Name, Value: value})
				return 0
			},
		},
		Members: Members{
			"ClassName": Member{
				Get: func(s State, v Value) int {
					return s.Push("string", v.(*types.Instance).ClassName)
				},
				// Allowed to be set for convenience.
				Set: func(s State, v Value) {
					v.(*types.Instance).ClassName = s.Pull(3, "string").(string)
				},
			},
			"Name": Member{
				Get: func(s State, v Value) int {
					return s.Push("string", v.(*types.Instance).Name())
				},
				Set: func(s State, v Value) {
					v.(*types.Instance).SetName(s.Pull(3, "string").(string))
				},
			},
			"Parent": Member{
				Get: func(s State, v Value) int {
					if parent := v.(*types.Instance).Parent(); parent != nil {
						return s.Push("Instance", parent)
					}
					return s.Push("nil", nil)
				},
				Set: func(s State, v Value) {
					var err error
					switch parent := s.PullAnyOf(3, "Instance", "nil").(type) {
					case *types.Instance:
						err = v.(*types.Instance).SetParent(parent)
					case nil:
						err = v.(*types.Instance).SetParent(nil)
					}
					if err != nil {
						s.L.RaiseError(err.Error())
					}
				},
			},
			"ClearAllChildren": Member{Method: true, Get: func(s State, v Value) int {
				v.(*types.Instance).RemoveAll()
				return 0
			}},
			"Clone": Member{Method: true, Get: func(s State, v Value) int {
				return s.Push("Instance", v.(*types.Instance).Clone())
			}},
			"Destroy": Member{Method: true, Get: func(s State, v Value) int {
				v.(*types.Instance).SetParent(nil)
				return 0
			}},
			"FindFirstAncestor": Member{Method: true, Get: func(s State, v Value) int {
				name := s.Pull(2, "string").(string)
				if ancestor := v.(*types.Instance).FindFirstAncestorOfClass(name); ancestor != nil {
					return s.Push("Instance", ancestor)
				}
				return s.Push("nil", nil)
			}},
			"FindFirstAncestorOfClass": Member{Method: true, Get: func(s State, v Value) int {
				className := s.Pull(2, "string").(string)
				if ancestor := v.(*types.Instance).FindFirstAncestorOfClass(className); ancestor != nil {
					return s.Push("Instance", ancestor)
				}
				return s.Push("nil", nil)
			}},
			"FindFirstChild": Member{Method: true, Get: func(s State, v Value) int {
				name := s.Pull(2, "string").(string)
				recurse := s.PullOpt(3, "bool", false).(bool)
				if child := v.(*types.Instance).FindFirstChild(name, recurse); child != nil {
					return s.Push("Instance", child)
				}
				return s.Push("nil", nil)
			}},
			"FindFirstChildOfClass": Member{Method: true, Get: func(s State, v Value) int {
				className := s.Pull(2, "string").(string)
				recurse := s.PullOpt(3, "bool", false).(bool)
				if child := v.(*types.Instance).FindFirstChildOfClass(className, recurse); child != nil {
					return s.Push("Instance", child)
				}
				return s.Push("nil", nil)
			}},
			"GetChildren": Member{Method: true, Get: func(s State, v Value) int {
				t := v.(*types.Instance).Children()
				return s.Push("Objects", t)
			}},
			"GetDescendants": Member{Method: true, Get: func(s State, v Value) int {
				return s.Push("Objects", v.(*types.Instance).Descendants())
			}},
			"GetFullName": Member{Method: true, Get: func(s State, v Value) int {
				return s.Push("string", v.(*types.Instance).GetFullName())
			}},
			"IsAncestorOf": Member{Method: true, Get: func(s State, v Value) int {
				descendant := s.Pull(2, "Instance").(*types.Instance)
				return s.Push("bool", v.(*types.Instance).IsAncestorOf(descendant))
			}},
			"IsDescendantOf": Member{Method: true, Get: func(s State, v Value) int {
				ancestor := s.Pull(2, "Instance").(*types.Instance)
				return s.Push("bool", v.(*types.Instance).IsDescendantOf(ancestor))
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				className := s.Pull(1, "string").(string)
				inst := types.NewInstance(className)
				return s.Push("Instance", inst)
			},
		},
		Environment: func(s State) {
			t := s.L.CreateTable(0, 1)
			t.RawSetString("new", s.L.NewFunction(func(l *lua.LState) int {
				dataModel := types.NewDataModel()
				return s.Push("Instance", dataModel)
			}))
			s.L.SetGlobal("DataModel", t)
		},
	}
}
