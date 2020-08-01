package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	lua "github.com/yuin/gopher-lua"
)

func init() { register(Vector3) }
func Vector3() Reflector {
	return Reflector{
		Name:     "Vector3",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "Vector3").(types.Vector3)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "Vector3").(types.Vector3)
				op := s.Pull(2, "Vector3").(types.Vector3)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
			"__add": func(s State) int {
				v := s.Pull(1, "Vector3").(types.Vector3)
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push(v.Add(op))
			},
			"__sub": func(s State) int {
				v := s.Pull(1, "Vector3").(types.Vector3)
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push(v.Sub(op))
			},
			"__mul": func(s State) int {
				v := s.Pull(1, "Vector3").(types.Vector3)
				switch op := s.PullAnyOf(2, "number", "Vector3").(type) {
				case types.Double:
					return s.Push(v.MulN(float64(op)))
				case types.Vector3:
					return s.Push(v.Mul(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector3 with an incompatible value type or nil")
					return 0
				}
			},
			"__div": func(s State) int {
				v := s.Pull(1, "Vector3").(types.Vector3)
				switch op := s.PullAnyOf(2, "number", "Vector3").(type) {
				case types.Double:
					return s.Push(v.DivN(float64(op)))
				case types.Vector3:
					return s.Push(v.Div(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector3 with an incompatible value type or nil")
					return 0
				}
			},
			"__unm": func(s State) int {
				v := s.Pull(1, "Vector3").(types.Vector3)
				return s.Push(v.Neg())
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
