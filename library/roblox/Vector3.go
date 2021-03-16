package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

func init() { register(Vector3) }
func Vector3() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "Vector3",
		PushTo:   rbxmk.PushTypeTo("Vector3"),
		PullFrom: rbxmk.PullTypeFrom("Vector3"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector3").(types.Vector3)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector3").(types.Vector3)
				op := s.Pull(2, "Vector3").(types.Vector3)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
			"__add": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector3").(types.Vector3)
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push(v.Add(op))
			},
			"__sub": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector3").(types.Vector3)
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push(v.Sub(op))
			},
			"__mul": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector3").(types.Vector3)
				switch op := s.PullAnyOf(2, "number", "Vector3").(type) {
				case types.Double:
					return s.Push(v.MulN(float64(op)))
				case types.Vector3:
					return s.Push(v.Mul(op))
				default:
					return s.ReflectorError(2)
				}
			},
			"__div": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector3").(types.Vector3)
				switch op := s.PullAnyOf(2, "number", "Vector3").(type) {
				case types.Double:
					return s.Push(v.DivN(float64(op)))
				case types.Vector3:
					return s.Push(v.Div(op))
				default:
					return s.ReflectorError(2)
				}
			},
			"__unm": func(s rbxmk.State) int {
				v := s.Pull(1, "Vector3").(types.Vector3)
				return s.Push(v.Neg())
			},
		},
		Properties: rbxmk.Properties{
			"X": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Vector3).X))
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
			"Y": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Vector3).Y))
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
			"Z": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Vector3).Z))
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
			"Magnitude": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Vector3).Magnitude()))
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
			"Unit": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.Vector3).Unit())
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("Vector3"), ReadOnly: true} },
			},
		},
		Methods: rbxmk.Methods{
			"Lerp": {
				Func: func(s rbxmk.State, v types.Value) int {
					goal := s.Pull(2, "Vector3").(types.Vector3)
					alpha := float64(s.Pull(3, "float").(types.Float))
					return s.Push(v.(types.Vector3).Lerp(goal, alpha))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "goal", Type: dt.Prim("Vector3")},
							{Name: "alpha", Type: dt.Prim("float")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Vector3")},
						},
					}
				},
			},
			"Dot": {
				Func: func(s rbxmk.State, v types.Value) int {
					op := s.Pull(2, "Vector3").(types.Vector3)
					return s.Push(types.Float(v.(types.Vector3).Dot(op)))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "op", Type: dt.Prim("Vector3")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("float")},
						},
					}
				},
			},
			"Cross": {
				Func: func(s rbxmk.State, v types.Value) int {
					op := s.Pull(2, "Vector3").(types.Vector3)
					return s.Push(v.(types.Vector3).Cross(op))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "op", Type: dt.Prim("Vector3")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Vector3")},
						},
					}
				},
			},
			"FuzzyEq": {
				Func: func(s rbxmk.State, v types.Value) int {
					op := s.Pull(2, "Vector3").(types.Vector3)
					epsilon := float64(s.Pull(3, "float").(types.Float))
					return s.Push(types.Bool(v.(types.Vector3).FuzzyEq(op, epsilon)))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "op", Type: dt.Prim("Vector3")},
							{Name: "epsilon", Type: dt.Prim("float")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("bool")},
						},
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.Vector3
					switch s.Count() {
					case 0:
					case 3:
						v.X = float32(s.Pull(1, "float").(types.Float))
						v.Y = float32(s.Pull(2, "float").(types.Float))
						v.Z = float32(s.Pull(3, "float").(types.Float))
					default:
						return s.RaiseError("expected 0 or 3 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{
						{
							Returns: dump.Parameters{
								{Type: dt.Prim("Vector3")},
							},
						},
						{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("float")},
								{Name: "y", Type: dt.Prim("float")},
								{Name: "z", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("Vector3")},
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
					Add: []dump.Binop{{Operand: dt.Prim("Vector3"), Result: dt.Prim("Vector3")}},
					Sub: []dump.Binop{{Operand: dt.Prim("Vector3"), Result: dt.Prim("Vector3")}},
					Mul: []dump.Binop{
						{Operand: dt.Prim("Vector3"), Result: dt.Prim("Vector3")},
						{Operand: dt.Prim("number"), Result: dt.Prim("Vector3")},
					},
					Div: []dump.Binop{
						{Operand: dt.Prim("Vector3"), Result: dt.Prim("Vector3")},
						{Operand: dt.Prim("number"), Result: dt.Prim("Vector3")},
					},
					Unm: &dump.Unop{Result: dt.Prim("Vector3")},
				},
			}
		},
	}
}
