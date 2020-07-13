package reflect

import (
	"strconv"

	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Vector3() Type {
	return Type{
		Name:        "Vector3",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Serialize: func(s State, v Value) (sv rbxfile.Value, err error) {
			if v, ok := v.(types.Vector3); ok {
				return rbxfile.ValueVector3(v), nil
			}
			return nil, TypeError(nil, 0, "Vector3")
		},
		Deserialize: func(s State, sv rbxfile.Value) (v Value, err error) {
			if sv, ok := sv.(rbxfile.ValueVector3); ok {
				return types.Vector3(sv), nil
			}
			return nil, TypeError(nil, 0, "Vector3")
		},
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				u := v.(types.Vector3)
				var b string
				b += strconv.FormatFloat(float64(u.X), 'g', -1, 32) + ", "
				b += strconv.FormatFloat(float64(u.Y), 'g', -1, 32) + ", "
				b += strconv.FormatFloat(float64(u.Z), 'g', -1, 32)
				s.L.Push(lua.LString(b))
				return 1
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push("bool", v.(types.Vector3) == op)
			},
			"__add": func(s State, v Value) int {
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push("Vector3", v.(types.Vector3).Add(op))
			},
			"__sub": func(s State, v Value) int {
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push("Vector3", v.(types.Vector3).Sub(op))
			},
			"__mul": func(s State, v Value) int {
				switch op := s.PullAnyOf(2, "number", "Vector3").(type) {
				case float64:
					return s.Push("Vector3", v.(types.Vector3).MulN(op))
				case types.Vector3:
					return s.Push("Vector3", v.(types.Vector3).Mul(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector3 with an incompatible value type or nil")
					return 0
				}
			},
			"__div": func(s State, v Value) int {
				switch op := s.PullAnyOf(2, "number", "Vector3").(type) {
				case float64:
					return s.Push("Vector3", v.(types.Vector3).DivN(op))
				case types.Vector3:
					return s.Push("Vector3", v.(types.Vector3).Div(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector3 with an incompatible value type or nil")
					return 0
				}
			},
			"__unm": func(s State, v Value) int {
				return s.Push("Vector3", v.(types.Vector3).Neg())
			},
		},
		Members: map[string]Member{
			"X": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.Vector3).X)
			}},
			"Y": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.Vector3).Y)
			}},
			"Z": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.Vector3).Z)
			}},
			"Magnitude": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.Vector3).Magnitude())
			}},
			"Unit": {Get: func(s State, v Value) int {
				return s.Push("Vector3", v.(types.Vector3).Unit())
			}},
			"Lerp": {Method: true, Get: func(s State, v Value) int {
				goal := s.Pull(2, "Vector3").(types.Vector3)
				alpha := s.Pull(3, "double").(float64)
				return s.Push("Vector3", v.(types.Vector3).Lerp(goal, alpha))
			}},
			"Dot": {Method: true, Get: func(s State, v Value) int {
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push("Vector3", v.(types.Vector3).Dot(op))
			}},
			"Cross": {Method: true, Get: func(s State, v Value) int {
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push("Vector3", v.(types.Vector3).Cross(op))
			}},
			"FuzzyEq": {Method: true, Get: func(s State, v Value) int {
				op := s.Pull(2, "Vector3").(types.Vector3)
				epsilon := s.Pull(3, "double").(float64)
				return s.Push("Vector3", v.(types.Vector3).FuzzyEq(op, epsilon))
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.Vector3
				switch s.Count() {
				case 0:
				case 3:
					v.X = s.Pull(1, "float").(float32)
					v.Y = s.Pull(2, "float").(float32)
					v.Z = s.Pull(3, "float").(float32)
				default:
					s.L.RaiseError("expected 0 or 3 arguments")
					return 0
				}
				return s.Push("Vector3", v)
			},
		},
	}
}
