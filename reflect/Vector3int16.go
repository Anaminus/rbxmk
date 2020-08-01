package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func Vector3int16() Reflector {
	return Reflector{
		Name:     "Vector3int16",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "Vector3int16").(types.Vector3int16)
				return s.Push(types.String(v.String()))
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "Vector3int16").(types.Vector3int16)
				op := s.Pull(2, "Vector3int16").(types.Vector3int16)
				return s.Push(types.Bool(v == op))
			},
			"__add": func(s State) int {
				v := s.Pull(1, "Vector3int16").(types.Vector3int16)
				op := s.Pull(2, "Vector3int16").(types.Vector3int16)
				return s.Push(v.Add(op))
			},
			"__sub": func(s State) int {
				v := s.Pull(1, "Vector3int16").(types.Vector3int16)
				op := s.Pull(2, "Vector3int16").(types.Vector3int16)
				return s.Push(v.Sub(op))
			},
			"__mul": func(s State) int {
				v := s.Pull(1, "Vector3int16").(types.Vector3int16)
				switch op := s.PullAnyOf(2, "number", "Vector3int16").(type) {
				case types.Double:
					return s.Push(v.MulN(float64(op)))
				case types.Vector3int16:
					return s.Push(v.Mul(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector3int16 with an incompatible value type or nil")
					return 0
				}
			},
			"__div": func(s State) int {
				v := s.Pull(1, "Vector3int16").(types.Vector3int16)
				switch op := s.PullAnyOf(2, "number", "Vector3int16").(type) {
				case types.Double:
					return s.Push(v.DivN(float64(op)))
				case types.Vector3int16:
					return s.Push(v.Div(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector3int16 with an incompatible value type or nil")
					return 0
				}
			},
			"__unm": func(s State) int {
				v := s.Pull(1, "Vector3int16").(types.Vector3int16)
				return s.Push(v.Neg())
			},
		},
		Members: map[string]Member{
			"X": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.Vector3int16).X))
			}},
			"Y": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.Vector3int16).Y))
			}},
			"Z": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.Vector3int16).Z))
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.Vector3int16
				switch s.Count() {
				case 0:
				case 3:
					v.X = int16(s.Pull(1, "int").(types.Int))
					v.Y = int16(s.Pull(2, "int").(types.Int))
					v.Z = int16(s.Pull(3, "int").(types.Int))
				default:
					return s.RaiseError("expected 0 or 3 arguments")
				}
				return s.Push(v)
			},
		},
	}
}
