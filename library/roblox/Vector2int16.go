package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

func init() { register(Vector2int16) }
func Vector2int16() Reflector {
	return Reflector{
		Name:     "Vector2int16",
		PushTo:   rbxmk.PushTypeTo("Vector2int16"),
		PullFrom: rbxmk.PullTypeFrom("Vector2int16"),
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "Vector2int16").(types.Vector2int16)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "Vector2int16").(types.Vector2int16)
				op := s.Pull(2, "Vector2int16").(types.Vector2int16)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
			"__add": func(s State) int {
				v := s.Pull(1, "Vector2int16").(types.Vector2int16)
				op := s.Pull(2, "Vector2int16").(types.Vector2int16)
				return s.Push(v.Add(op))
			},
			"__sub": func(s State) int {
				v := s.Pull(1, "Vector2int16").(types.Vector2int16)
				op := s.Pull(2, "Vector2int16").(types.Vector2int16)
				return s.Push(v.Sub(op))
			},
			"__mul": func(s State) int {
				v := s.Pull(1, "Vector2int16").(types.Vector2int16)
				switch op := s.PullAnyOf(2, "number", "Vector2int16").(type) {
				case types.Double:
					return s.Push(v.MulN(float64(op)))
				case types.Vector2int16:
					return s.Push(v.Mul(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector2int16 with an incompatible value type or nil")
					return 0
				}
			},
			"__div": func(s State) int {
				v := s.Pull(1, "Vector2int16").(types.Vector2int16)
				switch op := s.PullAnyOf(2, "number", "Vector2int16").(type) {
				case types.Double:
					return s.Push(v.DivN(float64(op)))
				case types.Vector2int16:
					return s.Push(v.Div(op))
				default:
					s.L.ArgError(2, "attempt to multiply a Vector2int16 with an incompatible value type or nil")
					return 0
				}
			},
			"__unm": func(s State) int {
				v := s.Pull(1, "Vector2int16").(types.Vector2int16)
				return s.Push(v.Neg())
			},
		},
		Members: map[string]Member{
			"X": {
				Get: func(s State, v types.Value) int {
					return s.Push(types.Int(v.(types.Vector2int16).X))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("int")} },
			},
			"Y": {
				Get: func(s State, v types.Value) int {
					return s.Push(types.Int(v.(types.Vector2int16).Y))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("int")} },
			},
		},
		Constructors: Constructors{
			"new": {
				Func: func(s State) int {
					var v types.Vector2int16
					switch s.Count() {
					case 0:
					case 2:
						v.X = int16(s.Pull(1, "int").(types.Int))
						v.Y = int16(s.Pull(2, "int").(types.Int))
					default:
						return s.RaiseError("expected 0 or 2 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{
						{
							Returns: dump.Parameters{
								{Type: dt.Prim("Vector2int16")},
							},
						},
						{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("int")},
								{Name: "y", Type: dt.Prim("int")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("Vector2int16")},
							},
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
