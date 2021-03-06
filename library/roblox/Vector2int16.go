package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

func init() { register(Vector2int16) }
func Vector2int16() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "Vector2int16",
		PushTo:   rbxmk.PushTypeTo("Vector2int16"),
		PullFrom: rbxmk.PullTypeFrom("Vector2int16"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector2int16").(types.Vector2int16)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector2int16").(types.Vector2int16)
				op := s.Pull(2, "Vector2int16").(types.Vector2int16)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
			"__add": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector2int16").(types.Vector2int16)
				op := s.Pull(2, "Vector2int16").(types.Vector2int16)
				return s.Push(v.Add(op))
			},
			"__sub": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector2int16").(types.Vector2int16)
				op := s.Pull(2, "Vector2int16").(types.Vector2int16)
				return s.Push(v.Sub(op))
			},
			"__mul": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector2int16").(types.Vector2int16)
				switch op := s.PullAnyOf(2, "number", "Vector2int16").(type) {
				case types.Double:
					return s.Push(v.MulN(float64(op)))
				case types.Vector2int16:
					return s.Push(v.Mul(op))
				default:
					return s.ReflectorError(2)
				}
			},
			"__div": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector2int16").(types.Vector2int16)
				switch op := s.PullAnyOf(2, "number", "Vector2int16").(type) {
				case types.Double:
					return s.Push(v.DivN(float64(op)))
				case types.Vector2int16:
					return s.Push(v.Div(op))
				default:
					return s.ReflectorError(2)
				}
			},
			"__unm": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector2int16").(types.Vector2int16)
				return s.Push(v.Neg())
			},
		},
		Properties: rbxmk.Properties{
			"X": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Int(v.(types.Vector2int16).X))
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("int"), ReadOnly: true} },
			},
			"Y": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Int(v.(types.Vector2int16).Y))
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("int"), ReadOnly: true} },
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
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
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Eq:  true,
					Add: []dump.Binop{{Operand: dt.Prim("Vector2int16"), Result: dt.Prim("Vector2int16")}},
					Sub: []dump.Binop{{Operand: dt.Prim("Vector2int16"), Result: dt.Prim("Vector2int16")}},
					Mul: []dump.Binop{
						{Operand: dt.Prim("Vector2int16"), Result: dt.Prim("Vector2int16")},
						{Operand: dt.Prim("number"), Result: dt.Prim("Vector2int16")},
					},
					Div: []dump.Binop{
						{Operand: dt.Prim("Vector2int16"), Result: dt.Prim("Vector2int16")},
						{Operand: dt.Prim("number"), Result: dt.Prim("Vector2int16")},
					},
					Unm: &dump.Unop{Result: dt.Prim("Vector2int16")},
				},
			}
		},
	}
}
