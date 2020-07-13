package reflect

import (
	"strconv"

	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Vector3int16() Type {
	return Type{
		Name:        "Vector3int16",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				u := v.(types.Vector3int16)
				var b string
				b += strconv.FormatInt(int64(u.X), 10) + ", "
				b += strconv.FormatInt(int64(u.Y), 10) + ", "
				b += strconv.FormatInt(int64(u.Z), 10)
				s.L.Push(lua.LString(b))
				return 1
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "Vector3int16").(types.Vector3int16)
				return s.Push("bool", v.(types.Vector3int16) == op)
			},
			"__add": func(s State, v Value) int {
				op := s.Pull(2, "Vector3int16").(types.Vector3int16)
				return s.Push("Vector3int16", v.(types.Vector3int16).Add(op))
			},
			"__sub": func(s State, v Value) int {
				op := s.Pull(2, "Vector3int16").(types.Vector3int16)
				return s.Push("Vector3int16", v.(types.Vector3int16).Sub(op))
			},
			"__mul": func(s State, v Value) int {
				switch op := s.PullAnyOf(2, "number", "Vector3int16").(type) {
				case float64:
					return s.Push("Vector3int16", v.(types.Vector3int16).MulN(op))
				case types.Vector3int16:
					return s.Push("Vector3int16", v.(types.Vector3int16).Mul(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector3int16 with an incompatible value type or nil")
					return 0
				}
			},
			"__div": func(s State, v Value) int {
				switch op := s.PullAnyOf(2, "number", "Vector3int16").(type) {
				case float64:
					return s.Push("Vector3int16", v.(types.Vector3int16).DivN(op))
				case types.Vector3int16:
					return s.Push("Vector3int16", v.(types.Vector3int16).Div(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector3int16 with an incompatible value type or nil")
					return 0
				}
			},
			"__unm": func(s State, v Value) int {
				return s.Push("Vector3int16", v.(types.Vector3int16).Neg())
			},
		},
		Members: map[string]Member{
			"X": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.Vector3int16).X)
			}},
			"Y": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.Vector3int16).Y)
			}},
			"Z": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.Vector3int16).Z)
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.Vector3int16
				switch s.Count() {
				case 0:
				case 3:
					v.X = int16(s.Pull(1, "int").(int))
					v.Y = int16(s.Pull(2, "int").(int))
					v.Z = int16(s.Pull(3, "int").(int))
				default:
					s.L.RaiseError("expected 0 or 3 arguments")
					return 0
				}
				return s.Push("Vector3int16", v)
			},
		},
	}
}
