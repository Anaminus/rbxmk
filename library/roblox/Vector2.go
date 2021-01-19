package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(Vector2) }
func Vector2() Reflector {
	return Reflector{
		Name:     "Vector2",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "Vector2").(types.Vector2)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "Vector2").(types.Vector2)
				op := s.Pull(2, "Vector2").(types.Vector2)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
			"__add": func(s State) int {
				v := s.Pull(1, "Vector2").(types.Vector2)
				op := s.Pull(2, "Vector2").(types.Vector2)
				return s.Push(v.Add(op))
			},
			"__sub": func(s State) int {
				v := s.Pull(1, "Vector2").(types.Vector2)
				op := s.Pull(2, "Vector2").(types.Vector2)
				return s.Push(v.Sub(op))
			},
			"__mul": func(s State) int {
				v := s.Pull(1, "Vector2").(types.Vector2)
				switch op := s.PullAnyOf(2, "number", "Vector2").(type) {
				case types.Double:
					return s.Push(v.MulN(float64(op)))
				case types.Vector2:
					return s.Push(v.Mul(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector2 with an incompatible value type or nil")
					return 0
				}
			},
			"__div": func(s State) int {
				v := s.Pull(1, "Vector2").(types.Vector2)
				switch op := s.PullAnyOf(2, "number", "Vector2").(type) {
				case types.Double:
					return s.Push(v.DivN(float64(op)))
				case types.Vector2:
					return s.Push(v.Div(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector2 with an incompatible value type or nil")
					return 0
				}
			},
			"__unm": func(s State) int {
				v := s.Pull(1, "Vector2").(types.Vector2)
				return s.Push(v.Neg())
			},
		},
		Members: map[string]Member{
			"X": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.Vector2).X))
			}},
			"Y": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.Vector2).Y))
			}},
			"Magnitude": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.Vector2).Magnitude()))
			}},
			"Unit": {Get: func(s State, v types.Value) int {
				return s.Push(v.(types.Vector2).Unit())
			}},
			"Lerp": {Method: true, Get: func(s State, v types.Value) int {
				goal := s.Pull(2, "Vector2").(types.Vector2)
				alpha := float64(s.Pull(3, "number").(types.Double))
				return s.Push(v.(types.Vector2).Lerp(goal, alpha))
			}},
			"Dot": {Method: true, Get: func(s State, v types.Value) int {
				op := s.Pull(2, "Vector2").(types.Vector2)
				return s.Push(types.Double(v.(types.Vector2).Dot(op)))
			}},
			"Cross": {Method: true, Get: func(s State, v types.Value) int {
				op := s.Pull(2, "Vector2").(types.Vector2)
				return s.Push(types.Double(v.(types.Vector2).Cross(op)))
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.Vector2
				switch s.Count() {
				case 0:
				case 2:
					v.X = float32(s.Pull(1, "float").(types.Float))
					v.Y = float32(s.Pull(2, "float").(types.Float))
				default:
					return s.RaiseError("expected 0 or 2 arguments")
				}
				return s.Push(v)
			},
		},
	}
}
