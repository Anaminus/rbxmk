package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func PhysicalProperties() Type {
	return Type{
		Name:        "PhysicalProperties",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				s.L.Push(lua.LString(v.(types.PhysicalProperties).String()))
				return 1
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "PhysicalProperties").(types.PhysicalProperties)
				return s.Push("bool", v.(types.PhysicalProperties) == op)
			},
		},
		Members: map[string]Member{
			"Density": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.PhysicalProperties).Density)
			}},
			"Friction": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.PhysicalProperties).Friction)
			}},
			"Elasticity": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.PhysicalProperties).Elasticity)
			}},
			"FrictionWeight": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.PhysicalProperties).FrictionWeight)
			}},
			"ElasticityWeight": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.PhysicalProperties).ElasticityWeight)
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.PhysicalProperties
				switch s.Count() {
				case 3:
					v.Density = s.Pull(1, "float").(float32)
					v.Friction = s.Pull(2, "float").(float32)
					v.Elasticity = s.Pull(3, "float").(float32)
				case 5:
					v.Density = s.Pull(1, "float").(float32)
					v.Friction = s.Pull(2, "float").(float32)
					v.Elasticity = s.Pull(3, "float").(float32)
					v.FrictionWeight = s.Pull(4, "float").(float32)
					v.ElasticityWeight = s.Pull(5, "float").(float32)
				default:
					s.L.RaiseError("expected 3 or 5 arguments")
					return 0
				}
				return s.Push("PhysicalProperties", v)
			},
		},
	}
}
