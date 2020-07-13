package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/types"
)

func Instance() Type {
	return Type{
		Name:        "Instance",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				return 1
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "Instance").(*types.Instance)
				return s.Push("bool", v.(*types.Instance) == op)
			},
			"__index": func(s State, v Value) int {
				inst := v.(*types.Instance)
				switch name := s.Pull(2, "string").(string); name {
				case "ClassName":
					return s.Push("string", inst.ClassName)
				case "Name":
					return s.Push("string", inst.Name())
				case "Parent":
					if parent := inst.Parent(); parent != nil {
						return s.Push("Instance", parent)
					}
					return s.Push("nil", nil)
				case "ClearAllChildren":
					inst.RemoveAll()
					return 0
				case "Clone":
					return s.Push("Instance", inst.Clone())
				case "Destroy":
					inst.SetParent(nil)
					return 0
				case "FindFirstAncestor":
					name := s.Pull(2, "string").(string)
					if ancestor := inst.FindFirstAncestorOfClass(name); ancestor != nil {
						return s.Push("Instance", ancestor)
					}
					return s.Push("nil", nil)
				case "FindFirstAncestorOfClass":
					className := s.Pull(2, "string").(string)
					if ancestor := inst.FindFirstAncestorOfClass(className); ancestor != nil {
						return s.Push("Instance", ancestor)
					}
					return s.Push("nil", nil)
				case "FindFirstChild":
					name := s.Pull(2, "string").(string)
					recurse := s.PullOpt(3, "bool", false).(bool)
					if child := inst.FindFirstChild(name, recurse); child != nil {
						return s.Push("Instance", child)
					}
					return s.Push("nil", nil)
				case "FindFirstChildOfClass":
					className := s.Pull(2, "string").(string)
					recurse := s.PullOpt(3, "bool", false).(bool)
					if child := inst.FindFirstChildOfClass(className, recurse); child != nil {
						return s.Push("Instance", child)
					}
					return s.Push("nil", nil)
				case "GetChildren":
					return s.Push("Objects", inst.Children())
				case "GetDescendants":
					return s.Push("Objects", inst.Descendants())
				case "GetFullName":
					return s.Push("string", inst.GetFullName())
				case "IsAncestorOf":
					descendant := s.Pull(2, "Instance").(*types.Instance)
					return s.Push("bool", inst.IsAncestorOf(descendant))
				case "IsDescendantOf":
					ancestor := s.Pull(2, "Instance").(*types.Instance)
					return s.Push("bool", inst.IsDescendantOf(ancestor))
				default:
					v := inst.Get(name)
					if v.Type == "" {
						// s.L.RaiseError("%s is not a valid member", name)
						return s.Push("nil", nil)
					}
					return s.Push(v.Type, v.Value)
				}
			},
			"__newindex": func(s State, v Value) int {
				inst := v.(*types.Instance)
				switch name := s.Pull(2, "string").(string); name {
				case "ClassName":
					inst.ClassName = s.Pull(3, "string").(string)
				case "Name":
					inst.SetName(s.Pull(3, "string").(string))
				case "Parent":
					switch parent := s.PullAnyOf(3, "Instance", "nil").(type) {
					case *types.Instance:
						inst.SetParent(parent)
					case nil:
						inst.SetParent(nil)
					}
				case "ClearAllChildren",
					"Clone",
					"Destroy",
					"FindFirstAncestor",
					"FindFirstAncestorOfClass",
					"FindFirstChild",
					"FindFirstChildOfClass",
					"GetChildren",
					"GetDescendants",
					"GetFullName",
					"IsAncestorOf",
					"IsDescendantOf":
					s.L.RaiseError("%s is not a valid member", name)
				default:
					value, typ := PullVariant(s, 3)
					inst.Set(name, TValue{Type: typ.Name, Value: value})
				}
				return 0
			},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				className := s.Pull(1, "string").(string)
				inst := types.NewInstance(className)
				return s.Push("Instance", inst)
			},
		},
	}
}
