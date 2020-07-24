package reflect

import (
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
			"__tostring": func(s State, v types.Value) int {
				s.L.Push(lua.LString(v.(types.Vector2).String()))
				return 1
			},
			"__eq": func(s State, v types.Value) int {
				op := s.Pull(2, "Vector2").(types.Vector2)
				return s.Push("bool", types.Bool(v.(types.Vector2) == op))
			},
			"__add": func(s State, v types.Value) int {
				op := s.Pull(2, "Vector2").(types.Vector2)
				return s.Push("Vector2", v.(types.Vector2).Add(op))
			},
			"__sub": func(s State, v types.Value) int {
				op := s.Pull(2, "Vector2").(types.Vector2)
				return s.Push("Vector2", v.(types.Vector2).Sub(op))
			},
			"__mul": func(s State, v types.Value) int {
				switch op := s.PullAnyOf(2, "number", "Vector2").(type) {
				case types.Double:
					return s.Push("Vector2", v.(types.Vector2).MulN(float64(op)))
				case types.Vector2:
					return s.Push("Vector2", v.(types.Vector2).Mul(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector2 with an incompatible value type or nil")
					return 0
				}
			},
			"__div": func(s State, v types.Value) int {
				switch op := s.PullAnyOf(2, "number", "Vector2").(type) {
				case types.Double:
					return s.Push("Vector2", v.(types.Vector2).DivN(float64(op)))
				case types.Vector2:
					return s.Push("Vector2", v.(types.Vector2).Div(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector2 with an incompatible value type or nil")
					return 0
				}
			},
			"__unm": func(s State, v types.Value) int {
				return s.Push("Vector2", v.(types.Vector2).Neg())
			},
		},
		Members: map[string]Member{
			"X": {Get: func(s State, v types.Value) int {
				return s.Push("float", types.Float(v.(types.Vector2).X))
			}},
			"Y": {Get: func(s State, v types.Value) int {
				return s.Push("float", types.Float(v.(types.Vector2).Y))
			}},
			"Magnitude": {Get: func(s State, v types.Value) int {
				return s.Push("float", types.Float(v.(types.Vector2).Magnitude()))
			}},
			"Unit": {Get: func(s State, v types.Value) int {
				return s.Push("Vector2", v.(types.Vector2).Unit())
			}},
			"Lerp": {Method: true, Get: func(s State, v types.Value) int {
				goal := s.Pull(2, "Vector2").(types.Vector2)
				alpha := float64(s.Pull(3, "number").(types.Double))
				return s.Push("Vector2", v.(types.Vector2).Lerp(goal, alpha))
			}},
			"Dot": {Method: true, Get: func(s State, v types.Value) int {
				op := s.Pull(2, "Vector2").(types.Vector2)
				return s.Push("number", types.Double(v.(types.Vector2).Dot(op)))
			}},
			"Cross": {Method: true, Get: func(s State, v types.Value) int {
				op := s.Pull(2, "Vector2").(types.Vector2)
				return s.Push("number", types.Double(v.(types.Vector2).Cross(op)))
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
					s.L.RaiseError("expected 0 or 2 arguments")
					return 0
				}
				return s.Push("Vector2", v)
			},
		},
	}
}
