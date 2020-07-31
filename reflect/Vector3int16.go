package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Vector3int16() Reflector {
	return Reflector{
		Name:     "Vector3int16",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString(s.Pull(1, "Vector3int16").(types.Vector3int16).String()))
				return 1
			},
			"__eq": func(s State) int {
				op := s.Pull(2, "Vector3int16").(types.Vector3int16)
				return s.Push(types.Bool(s.Pull(1, "Vector3int16").(types.Vector3int16) == op))
			},
			"__add": func(s State) int {
				op := s.Pull(2, "Vector3int16").(types.Vector3int16)
				return s.Push(s.Pull(1, "Vector3int16").(types.Vector3int16).Add(op))
			},
			"__sub": func(s State) int {
				op := s.Pull(2, "Vector3int16").(types.Vector3int16)
				return s.Push(s.Pull(1, "Vector3int16").(types.Vector3int16).Sub(op))
			},
			"__mul": func(s State) int {
				switch op := s.PullAnyOf(2, "number", "Vector3int16").(type) {
				case types.Double:
					return s.Push(s.Pull(1, "Vector3int16").(types.Vector3int16).MulN(float64(op)))
				case types.Vector3int16:
					return s.Push(s.Pull(1, "Vector3int16").(types.Vector3int16).Mul(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector3int16 with an incompatible value type or nil")
					return 0
				}
			},
			"__div": func(s State) int {
				switch op := s.PullAnyOf(2, "number", "Vector3int16").(type) {
				case types.Double:
					return s.Push(s.Pull(1, "Vector3int16").(types.Vector3int16).DivN(float64(op)))
				case types.Vector3int16:
					return s.Push(s.Pull(1, "Vector3int16").(types.Vector3int16).Div(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector3int16 with an incompatible value type or nil")
					return 0
				}
			},
			"__unm": func(s State) int {
				return s.Push(s.Pull(1, "Vector3int16").(types.Vector3int16).Neg())
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
