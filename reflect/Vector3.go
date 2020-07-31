package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Vector3() Type {
	return Type{
		Name:     "Vector3",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString(s.Pull(1, "Vector3").(types.Vector3).String()))
				return 1
			},
			"__eq": func(s State) int {
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push(types.Bool(s.Pull(1, "Vector3").(types.Vector3) == op))
			},
			"__add": func(s State) int {
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push(s.Pull(1, "Vector3").(types.Vector3).Add(op))
			},
			"__sub": func(s State) int {
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push(s.Pull(1, "Vector3").(types.Vector3).Sub(op))
			},
			"__mul": func(s State) int {
				switch op := s.PullAnyOf(2, "number", "Vector3").(type) {
				case types.Double:
					return s.Push(s.Pull(1, "Vector3").(types.Vector3).MulN(float64(op)))
				case types.Vector3:
					return s.Push(s.Pull(1, "Vector3").(types.Vector3).Mul(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector3 with an incompatible value type or nil")
					return 0
				}
			},
			"__div": func(s State) int {
				switch op := s.PullAnyOf(2, "number", "Vector3").(type) {
				case types.Double:
					return s.Push(s.Pull(1, "Vector3").(types.Vector3).DivN(float64(op)))
				case types.Vector3:
					return s.Push(s.Pull(1, "Vector3").(types.Vector3).Div(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector3 with an incompatible value type or nil")
					return 0
				}
			},
			"__unm": func(s State) int {
				return s.Push(s.Pull(1, "Vector3").(types.Vector3).Neg())
			},
		},
		Members: map[string]Member{
			"X": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.Vector3).X))
			}},
			"Y": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.Vector3).Y))
			}},
			"Z": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.Vector3).Z))
			}},
			"Magnitude": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.Vector3).Magnitude()))
			}},
			"Unit": {Get: func(s State, v types.Value) int {
				return s.Push(v.(types.Vector3).Unit())
			}},
			"Lerp": {Method: true, Get: func(s State, v types.Value) int {
				goal := s.Pull(2, "Vector3").(types.Vector3)
				alpha := float64(s.Pull(3, "number").(types.Double))
				return s.Push(v.(types.Vector3).Lerp(goal, alpha))
			}},
			"Dot": {Method: true, Get: func(s State, v types.Value) int {
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push(types.Float(v.(types.Vector3).Dot(op)))
			}},
			"Cross": {Method: true, Get: func(s State, v types.Value) int {
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push(v.(types.Vector3).Cross(op))
			}},
			"FuzzyEq": {Method: true, Get: func(s State, v types.Value) int {
				op := s.Pull(2, "Vector3").(types.Vector3)
				epsilon := float64(s.Pull(3, "number").(types.Double))
				return s.Push(types.Bool(v.(types.Vector3).FuzzyEq(op, epsilon)))
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.Vector3
				switch s.Count() {
				case 0:
				case 3:
					v.X = float32(s.Pull(1, "float").(types.Float))
					v.Y = float32(s.Pull(2, "float").(types.Float))
					v.Z = float32(s.Pull(3, "float").(types.Float))
				default:
					return s.RaiseError("expected 0 or 3 arguments")
				}
				return s.Push(v)
			},
		},
	}
}
