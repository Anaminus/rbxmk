package reflect

import (
	"strconv"

	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Vector2() Type {
	return Type{
		Name:        "Vector2",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				u := v.(types.Vector2)
				var b string
				b += strconv.FormatFloat(float64(u.X), 'g', -1, 32) + ", "
				b += strconv.FormatFloat(float64(u.Y), 'g', -1, 32)
				s.L.Push(lua.LString(b))
				return 1
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "Vector2").(types.Vector2)
				return s.Push("bool", v.(types.Vector2) == op)
			},
			"__add": func(s State, v Value) int {
				op := s.Pull(2, "Vector2").(types.Vector2)
				return s.Push("Vector2", v.(types.Vector2).Add(op))
			},
			"__sub": func(s State, v Value) int {
				op := s.Pull(2, "Vector2").(types.Vector2)
				return s.Push("Vector2", v.(types.Vector2).Sub(op))
			},
			"__mul": func(s State, v Value) int {
				switch op := s.PullAnyOf(2, "number", "Vector2").(type) {
				case float64:
					return s.Push("Vector2", v.(types.Vector2).MulN(op))
				case types.Vector2:
					return s.Push("Vector2", v.(types.Vector2).Mul(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector2 with an incompatible value type or nil")
					return 0
				}
			},
			"__div": func(s State, v Value) int {
				switch op := s.PullAnyOf(2, "number", "Vector2").(type) {
				case float64:
					return s.Push("Vector2", v.(types.Vector2).DivN(op))
				case types.Vector2:
					return s.Push("Vector2", v.(types.Vector2).Div(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector2 with an incompatible value type or nil")
					return 0
				}
			},
			"__unm": func(s State, v Value) int {
				return s.Push("Vector2", v.(types.Vector2).Neg())
			},
		},
		Members: map[string]Member{
			"X": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.Vector2).X)
			}},
			"Y": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.Vector2).Y)
			}},
			"Magnitude": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.Vector2).Magnitude())
			}},
			"Unit": {Get: func(s State, v Value) int {
				return s.Push("Vector2", v.(types.Vector2).Unit())
			}},
			"Lerp": {Method: true, Get: func(s State, v Value) int {
				goal := s.Pull(2, "Vector2").(types.Vector2)
				alpha := s.Pull(3, "double").(float64)
				return s.Push("Vector2", v.(types.Vector2).Lerp(goal, alpha))
			}},
			"Dot": {Method: true, Get: func(s State, v Value) int {
				op := s.Pull(2, "Vector2").(types.Vector2)
				return s.Push("Vector2", v.(types.Vector2).Dot(op))
			}},
			"Cross": {Method: true, Get: func(s State, v Value) int {
				op := s.Pull(2, "Vector2").(types.Vector2)
				return s.Push("number", v.(types.Vector2).Cross(op))
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.Vector2
				switch s.Count() {
				case 0:
				case 2:
					v.X = s.Pull(1, "float").(float32)
					v.Y = s.Pull(2, "float").(float32)
				default:
					s.L.RaiseError("expected 0 or 2 arguments")
					return 0
				}
				return s.Push("Vector2", v)
			},
		},
	}
}
