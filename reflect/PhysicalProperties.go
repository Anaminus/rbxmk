package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func PhysicalProperties() Reflector {
	return Reflector{
		Name:     "PhysicalProperties",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "PhysicalProperties").(types.PhysicalProperties)
				return s.Push(types.String(v.String()))
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "PhysicalProperties").(types.PhysicalProperties)
				op := s.Pull(2, "PhysicalProperties").(types.PhysicalProperties)
				return s.Push(types.Bool(v == op))
			},
		},
		Members: map[string]Member{
			"Density": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.PhysicalProperties).Density))
			}},
			"Friction": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.PhysicalProperties).Friction))
			}},
			"Elasticity": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.PhysicalProperties).Elasticity))
			}},
			"FrictionWeight": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.PhysicalProperties).FrictionWeight))
			}},
			"ElasticityWeight": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.PhysicalProperties).ElasticityWeight))
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.PhysicalProperties
				switch s.Count() {
				case 3:
					v.Density = float32(s.Pull(1, "float").(types.Float))
					v.Friction = float32(s.Pull(2, "float").(types.Float))
					v.Elasticity = float32(s.Pull(3, "float").(types.Float))
				case 5:
					v.Density = float32(s.Pull(1, "float").(types.Float))
					v.Friction = float32(s.Pull(2, "float").(types.Float))
					v.Elasticity = float32(s.Pull(3, "float").(types.Float))
					v.FrictionWeight = float32(s.Pull(4, "float").(types.Float))
					v.ElasticityWeight = float32(s.Pull(5, "float").(types.Float))
				default:
					return s.RaiseError("expected 3 or 5 arguments")
				}
				return s.Push(v)
			},
		},
	}
}
