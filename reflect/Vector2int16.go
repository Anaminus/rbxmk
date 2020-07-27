package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Vector2int16() Type {
	return Type{
		Name:     "Vector2int16",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString(s.Pull(1, "Vector2int16").(types.Vector2int16).String()))
				return 1
			},
			"__eq": func(s State) int {
				op := s.Pull(2, "Vector2int16").(types.Vector2int16)
				return s.Push(types.Bool(s.Pull(1, "Vector2int16").(types.Vector2int16) == op))
			},
			"__add": func(s State) int {
				op := s.Pull(2, "Vector2int16").(types.Vector2int16)
				return s.Push(s.Pull(1, "Vector2int16").(types.Vector2int16).Add(op))
			},
			"__sub": func(s State) int {
				op := s.Pull(2, "Vector2int16").(types.Vector2int16)
				return s.Push(s.Pull(1, "Vector2int16").(types.Vector2int16).Sub(op))
			},
			"__mul": func(s State) int {
				switch op := s.PullAnyOf(2, "number", "Vector2int16").(type) {
				case types.Double:
					return s.Push(s.Pull(1, "Vector2int16").(types.Vector2int16).MulN(float64(op)))
				case types.Vector2int16:
					return s.Push(s.Pull(1, "Vector2int16").(types.Vector2int16).Mul(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector2int16 with an incompatible value type or nil")
					return 0
				}
			},
			"__div": func(s State) int {
				switch op := s.PullAnyOf(2, "number", "Vector2int16").(type) {
				case types.Double:
					return s.Push(s.Pull(1, "Vector2int16").(types.Vector2int16).DivN(float64(op)))
				case types.Vector2int16:
					return s.Push(s.Pull(1, "Vector2int16").(types.Vector2int16).Div(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector2int16 with an incompatible value type or nil")
					return 0
				}
			},
			"__unm": func(s State) int {
				return s.Push(s.Pull(1, "Vector2int16").(types.Vector2int16).Neg())
			},
		},
		Members: map[string]Member{
			"X": {Get: func(s State, v types.Value) int {
				return s.Push(types.Int(v.(types.Vector2int16).X))
			}},
			"Y": {Get: func(s State, v types.Value) int {
				return s.Push(types.Int(v.(types.Vector2int16).Y))
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.Vector2int16
				switch s.Count() {
				case 0:
				case 2:
					v.X = int16(s.Pull(1, "int").(types.Int))
					v.Y = int16(s.Pull(2, "int").(types.Int))
				default:
					s.L.RaiseError("expected 0 or 2 arguments")
					return 0
				}
				return s.Push(v)
			},
		},
	}
}
