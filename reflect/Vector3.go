package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

const T_Vector3 = "Vector3"

func init() { register(Vector3) }
func Vector3() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     T_Vector3,
		PushTo:   rbxmk.PushTypeTo(T_Vector3),
		PullFrom: rbxmk.PullTypeFrom(T_Vector3),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.Vector3:
				*p = v.(types.Vector3)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, T_Vector3).(types.Vector3)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, T_Vector3).(types.Vector3)
				op := s.Pull(2, T_Vector3).(types.Vector3)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
			"__add": func(s rbxmk.State) int {
				v := s.Pull(1, T_Vector3).(types.Vector3)
				op := s.Pull(2, T_Vector3).(types.Vector3)
				return s.Push(v.Add(op))
			},
			"__sub": func(s rbxmk.State) int {
				v := s.Pull(1, T_Vector3).(types.Vector3)
				op := s.Pull(2, T_Vector3).(types.Vector3)
				return s.Push(v.Sub(op))
			},
			"__mul": func(s rbxmk.State) int {
				v := s.Pull(1, T_Vector3).(types.Vector3)
				switch op := s.PullAnyOf(2, T_Number, T_Vector3).(type) {
				case types.Double:
					return s.Push(v.MulN(float64(op)))
				case types.Vector3:
					return s.Push(v.Mul(op))
				default:
					return s.ReflectorError(2)
				}
			},
			"__div": func(s rbxmk.State) int {
				v := s.Pull(1, T_Vector3).(types.Vector3)
				switch op := s.PullAnyOf(2, T_Number, T_Vector3).(type) {
				case types.Double:
					return s.Push(v.DivN(float64(op)))
				case types.Vector3:
					return s.Push(v.Div(op))
				default:
					return s.ReflectorError(2)
				}
			},
			"__unm": func(s rbxmk.State) int {
				v := s.Pull(1, T_Vector3).(types.Vector3)
				return s.Push(v.Neg())
			},
		},
		Properties: rbxmk.Properties{
			"X": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Vector3).X))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Float),
						ReadOnly:    true,
						Summary:     "Types/Vector3:Properties/X/Summary",
						Description: "Types/Vector3:Properties/X/Description",
					}
				},
			},
			"Y": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Vector3).Y))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Float),
						ReadOnly:    true,
						Summary:     "Types/Vector3:Properties/Y/Summary",
						Description: "Types/Vector3:Properties/Y/Description",
					}
				},
			},
			"Z": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Vector3).Z))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Float),
						ReadOnly:    true,
						Summary:     "Types/Vector3:Properties/Z/Summary",
						Description: "Types/Vector3:Properties/Z/Description",
					}
				},
			},
			"Magnitude": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Vector3).Magnitude()))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Float),
						ReadOnly:    true,
						Summary:     "Types/Vector3:Properties/Magnitude/Summary",
						Description: "Types/Vector3:Properties/Magnitude/Description",
					}
				},
			},
			"Unit": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.Vector3).Unit())
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Vector3),
						ReadOnly:    true,
						Summary:     "Types/Vector3:Properties/Unit/Summary",
						Description: "Types/Vector3:Properties/Unit/Description",
					}
				},
			},
		},
		Methods: rbxmk.Methods{
			"Lerp": {
				Func: func(s rbxmk.State, v types.Value) int {
					goal := s.Pull(2, T_Vector3).(types.Vector3)
					alpha := float64(s.Pull(3, T_Float).(types.Float))
					return s.Push(v.(types.Vector3).Lerp(goal, alpha))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "goal", Type: dt.Prim(T_Vector3)},
							{Name: "alpha", Type: dt.Prim(T_Float)},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(T_Vector3)},
						},
						Summary:     "Types/Vector3:Methods/Lerp/Summary",
						Description: "Types/Vector3:Methods/Lerp/Description",
					}
				},
			},
			"Dot": {
				Func: func(s rbxmk.State, v types.Value) int {
					op := s.Pull(2, T_Vector3).(types.Vector3)
					return s.Push(types.Float(v.(types.Vector3).Dot(op)))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "op", Type: dt.Prim(T_Vector3)},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(T_Float)},
						},
						Summary:     "Types/Vector3:Methods/Dot/Summary",
						Description: "Types/Vector3:Methods/Dot/Description",
					}
				},
			},
			"Cross": {
				Func: func(s rbxmk.State, v types.Value) int {
					op := s.Pull(2, T_Vector3).(types.Vector3)
					return s.Push(v.(types.Vector3).Cross(op))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "op", Type: dt.Prim(T_Vector3)},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(T_Vector3)},
						},
						Summary:     "Types/Vector3:Methods/Cross/Summary",
						Description: "Types/Vector3:Methods/Cross/Description",
					}
				},
			},
			"FuzzyEq": {
				Func: func(s rbxmk.State, v types.Value) int {
					op := s.Pull(2, T_Vector3).(types.Vector3)
					epsilon := float64(s.Pull(3, T_Float).(types.Float))
					return s.Push(types.Bool(v.(types.Vector3).FuzzyEq(op, epsilon)))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "op", Type: dt.Prim(T_Vector3)},
							{Name: "epsilon", Type: dt.Prim(T_Float)},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(T_Bool)},
						},
						Summary:     "Types/Vector3:Methods/FuzzyEq/Summary",
						Description: "Types/Vector3:Methods/FuzzyEq/Description",
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
						v.X = float32(s.Pull(1, T_Float).(types.Float))
						v.Y = float32(s.Pull(2, T_Float).(types.Float))
						v.Z = float32(s.Pull(3, T_Float).(types.Float))
					default:
						return s.RaiseError("expected 0 or 3 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Returns: dump.Parameters{
								{Type: dt.Prim(T_Vector3)},
							},
							Summary:     "Types/Vector3:Constructors/new/Zero/Summary",
							Description: "Types/Vector3:Constructors/new/Zero/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(T_Float)},
								{Name: "y", Type: dt.Prim(T_Float)},
								{Name: "z", Type: dt.Prim(T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_Vector3)},
							},
							Summary:     "Types/Vector3:Constructors/new/Components/Summary",
							Description: "Types/Vector3:Constructors/new/Components/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Eq: &dump.Cmpop{
						Summary:     "Types/Vector3:Operators/Eq/Summary",
						Description: "Types/Vector3:Operators/Eq/Description",
					},
					Add: []dump.Binop{
						{
							Operand:     dt.Prim(T_Vector3),
							Result:      dt.Prim(T_Vector3),
							Summary:     "Types/Vector3:Operators/Add/Summary",
							Description: "Types/Vector3:Operators/Add/Description",
						},
					},
					Sub: []dump.Binop{
						{
							Operand:     dt.Prim(T_Vector3),
							Result:      dt.Prim(T_Vector3),
							Summary:     "Types/Vector3:Operators/Sub/Summary",
							Description: "Types/Vector3:Operators/Sub/Description",
						},
					},
					Mul: []dump.Binop{
						{
							Operand:     dt.Prim(T_Vector3),
							Result:      dt.Prim(T_Vector3),
							Summary:     "Types/Vector3:Operators/Mul/Vector3/Summary",
							Description: "Types/Vector3:Operators/Mul/Vector3/Description",
						},
						{
							Operand:     dt.Prim(T_Number),
							Result:      dt.Prim(T_Vector3),
							Summary:     "Types/Vector3:Operators/Mul/Number/Summary",
							Description: "Types/Vector3:Operators/Mul/Number/Description",
						},
					},
					Div: []dump.Binop{
						{
							Operand:     dt.Prim(T_Vector3),
							Result:      dt.Prim(T_Vector3),
							Summary:     "Types/Vector3:Operators/Div/Vector3/Summary",
							Description: "Types/Vector3:Operators/Div/Vector3/Description",
						},
						{
							Operand:     dt.Prim(T_Number),
							Result:      dt.Prim(T_Vector3),
							Summary:     "Types/Vector3:Operators/Div/Number/Summary",
							Description: "Types/Vector3:Operators/Div/Number/Description",
						},
					},
					Unm: &dump.Unop{
						Result:      dt.Prim(T_Vector3),
						Summary:     "Types/Vector3:Operators/Unm/Summary",
						Description: "Types/Vector3:Operators/Unm/Description",
					},
				},
				Summary:     "Types/Vector3:Summary",
				Description: "Types/Vector3:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Float,
			Number,
		},
	}
}
