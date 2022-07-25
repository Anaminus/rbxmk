package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Vector2) }
func Vector2() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     rtypes.T_Vector2,
		PushTo:   rbxmk.PushTypeTo(rtypes.T_Vector2),
		PullFrom: rbxmk.PullTypeFrom(rtypes.T_Vector2),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.Vector2:
				*p = v.(types.Vector2)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Vector2).(types.Vector2)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Vector2).(types.Vector2)
				op := s.Pull(2, rtypes.T_Vector2).(types.Vector2)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
			"__add": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Vector2).(types.Vector2)
				op := s.Pull(2, rtypes.T_Vector2).(types.Vector2)
				return s.Push(v.Add(op))
			},
			"__sub": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Vector2).(types.Vector2)
				op := s.Pull(2, rtypes.T_Vector2).(types.Vector2)
				return s.Push(v.Sub(op))
			},
			"__mul": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Vector2).(types.Vector2)
				switch op := s.PullAnyOf(2, rtypes.T_Number, rtypes.T_Vector2).(type) {
				case types.Double:
					return s.Push(v.MulN(float64(op)))
				case types.Vector2:
					return s.Push(v.Mul(op))
				default:
					return s.ReflectorError(2)
				}
			},
			"__div": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Vector2).(types.Vector2)
				switch op := s.PullAnyOf(2, rtypes.T_Number, rtypes.T_Vector2).(type) {
				case types.Double:
					return s.Push(v.DivN(float64(op)))
				case types.Vector2:
					return s.Push(v.Div(op))
				default:
					return s.ReflectorError(2)
				}
			},
			"__unm": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Vector2).(types.Vector2)
				return s.Push(v.Neg())
			},
		},
		Properties: rbxmk.Properties{
			"X": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Vector2).X))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Float),
						ReadOnly:    true,
						Summary:     "Types/Vector2:Properties/X/Summary",
						Description: "Types/Vector2:Properties/X/Description",
					}
				},
			},
			"Y": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Vector2).Y))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Float),
						ReadOnly:    true,
						Summary:     "Types/Vector2:Properties/Y/Summary",
						Description: "Types/Vector2:Properties/Y/Description",
					}
				},
			},
			"Magnitude": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Vector2).Magnitude()))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Float),
						ReadOnly:    true,
						Summary:     "Types/Vector2:Properties/Magnitude/Summary",
						Description: "Types/Vector2:Properties/Magnitude/Description",
					}
				},
			},
			"Unit": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.Vector2).Unit())
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Vector2),
						ReadOnly:    true,
						Summary:     "Types/Vector2:Properties/Unit/Summary",
						Description: "Types/Vector2:Properties/Unit/Description",
					}
				},
			},
		},
		Methods: rbxmk.Methods{
			"Lerp": {
				Func: func(s rbxmk.State, v types.Value) int {
					goal := s.Pull(2, rtypes.T_Vector2).(types.Vector2)
					alpha := float64(s.Pull(3, rtypes.T_Float).(types.Float))
					return s.Push(v.(types.Vector2).Lerp(goal, alpha))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "goal", Type: dt.Prim(rtypes.T_Vector2)},
							{Name: "alpha", Type: dt.Prim(rtypes.T_Float)},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(rtypes.T_Vector2)},
						},
						Summary:     "Types/Vector2:Methods/Lerp/Summary",
						Description: "Types/Vector2:Methods/Lerp/Description",
					}
				},
			},
			"Dot": {
				Func: func(s rbxmk.State, v types.Value) int {
					op := s.Pull(2, rtypes.T_Vector2).(types.Vector2)
					return s.Push(types.Double(v.(types.Vector2).Dot(op)))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "op", Type: dt.Prim(rtypes.T_Vector2)},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(rtypes.T_Float)},
						},
						Summary:     "Types/Vector2:Methods/Dot/Summary",
						Description: "Types/Vector2:Methods/Dot/Description",
					}
				},
			},
			"Cross": {
				Func: func(s rbxmk.State, v types.Value) int {
					op := s.Pull(2, rtypes.T_Vector2).(types.Vector2)
					return s.Push(types.Double(v.(types.Vector2).Cross(op)))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "op", Type: dt.Prim(rtypes.T_Vector2)},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(rtypes.T_Float)},
						},
						Summary:     "Types/Vector2:Methods/Cross/Summary",
						Description: "Types/Vector2:Methods/Cross/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.Vector2
					switch s.Count() {
					case 0:
					case 2:
						v.X = float32(s.Pull(1, rtypes.T_Float).(types.Float))
						v.Y = float32(s.Pull(2, rtypes.T_Float).(types.Float))
					default:
						return s.RaiseError("expected 0 or 2 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_Vector2)},
							},
							Summary:     "Types/Vector2:Constructors/new/Zero/Summary",
							Description: "Types/Vector2:Constructors/new/Zero/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_Float)},
								{Name: "y", Type: dt.Prim(rtypes.T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_Vector2)},
							},
							Summary:     "Types/Vector2:Constructors/new/Components/Summary",
							Description: "Types/Vector2:Constructors/new/Components/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category: "roblox",
				Operators: &dump.Operators{
					Eq: &dump.Cmpop{
						Summary:     "Types/Vector2:Operators/Eq/Summary",
						Description: "Types/Vector2:Operators/Eq/Description",
					},
					Add: []dump.Binop{
						{
							Operand:     dt.Prim(rtypes.T_Vector2),
							Result:      dt.Prim(rtypes.T_Vector2),
							Summary:     "Types/Vector2:Operators/Add/Summary",
							Description: "Types/Vector2:Operators/Add/Description",
						},
					},
					Sub: []dump.Binop{
						{
							Operand:     dt.Prim(rtypes.T_Vector2),
							Result:      dt.Prim(rtypes.T_Vector2),
							Summary:     "Types/Vector2:Operators/Sub/Summary",
							Description: "Types/Vector2:Operators/Sub/Description",
						},
					},
					Mul: []dump.Binop{
						{
							Operand:     dt.Prim(rtypes.T_Vector2),
							Result:      dt.Prim(rtypes.T_Vector2),
							Summary:     "Types/Vector2:Operators/Mul/Vector2/Summary",
							Description: "Types/Vector2:Operators/Mul/Vector2/Description",
						},
						{
							Operand:     dt.Prim(rtypes.T_Number),
							Result:      dt.Prim(rtypes.T_Vector2),
							Summary:     "Types/Vector2:Operators/Mul/Number/Summary",
							Description: "Types/Vector2:Operators/Mul/Number/Description",
						},
					},
					Div: []dump.Binop{
						{
							Operand:     dt.Prim(rtypes.T_Vector2),
							Result:      dt.Prim(rtypes.T_Vector2),
							Summary:     "Types/Vector2:Operators/Div/Vector2/Summary",
							Description: "Types/Vector2:Operators/Div/Vector2/Description",
						},
						{
							Operand:     dt.Prim(rtypes.T_Number),
							Result:      dt.Prim(rtypes.T_Vector2),
							Summary:     "Types/Vector2:Operators/Div/Number/Summary",
							Description: "Types/Vector2:Operators/Div/Number/Description",
						},
					},
					Unm: &dump.Unop{
						Result:      dt.Prim(rtypes.T_Vector2),
						Summary:     "Types/Vector2:Operators/Unm/Summary",
						Description: "Types/Vector2:Operators/Unm/Description",
					},
				},
				Summary:     "Types/Vector2:Summary",
				Description: "Types/Vector2:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Float,
			Number,
		},
	}
}
