package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Vector3int16) }
func Vector3int16() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     rtypes.T_Vector3int16,
		PushTo:   rbxmk.PushTypeTo(rtypes.T_Vector3int16),
		PullFrom: rbxmk.PullTypeFrom(rtypes.T_Vector3int16),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.Vector3int16:
				*p = v.(types.Vector3int16)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Vector3int16).(types.Vector3int16)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Vector3int16).(types.Vector3int16)
				op := s.Pull(2, rtypes.T_Vector3int16).(types.Vector3int16)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
			"__add": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Vector3int16).(types.Vector3int16)
				op := s.Pull(2, rtypes.T_Vector3int16).(types.Vector3int16)
				return s.Push(v.Add(op))
			},
			"__sub": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Vector3int16).(types.Vector3int16)
				op := s.Pull(2, rtypes.T_Vector3int16).(types.Vector3int16)
				return s.Push(v.Sub(op))
			},
			"__mul": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Vector3int16).(types.Vector3int16)
				switch op := s.PullAnyOf(2, rtypes.T_Number, rtypes.T_Vector3int16).(type) {
				case types.Double:
					return s.Push(v.MulN(float64(op)))
				case types.Vector3int16:
					return s.Push(v.Mul(op))
				default:
					return s.ReflectorError(2)
				}
			},
			"__div": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Vector3int16).(types.Vector3int16)
				switch op := s.PullAnyOf(2, rtypes.T_Number, rtypes.T_Vector3int16).(type) {
				case types.Double:
					return s.Push(v.DivN(float64(op)))
				case types.Vector3int16:
					return s.Push(v.Div(op))
				default:
					return s.ReflectorError(2)
				}
			},
			"__unm": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Vector3int16).(types.Vector3int16)
				return s.Push(v.Neg())
			},
		},
		Properties: rbxmk.Properties{
			"X": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Int(v.(types.Vector3int16).X))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Int),
						ReadOnly:    true,
						Summary:     "Types/Vector3int16:Properties/X/Summary",
						Description: "Types/Vector3int16:Properties/X/Description",
					}
				},
			},
			"Y": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Int(v.(types.Vector3int16).Y))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Int),
						ReadOnly:    true,
						Summary:     "Types/Vector3int16:Properties/Y/Summary",
						Description: "Types/Vector3int16:Properties/Y/Description",
					}
				},
			},
			"Z": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Int(v.(types.Vector3int16).Z))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Int),
						ReadOnly:    true,
						Summary:     "Types/Vector3int16:Properties/Z/Summary",
						Description: "Types/Vector3int16:Properties/Z/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.Vector3int16
					switch s.Count() {
					case 0:
					case 3:
						v.X = int16(s.Pull(1, rtypes.T_Int).(types.Int))
						v.Y = int16(s.Pull(2, rtypes.T_Int).(types.Int))
						v.Z = int16(s.Pull(3, rtypes.T_Int).(types.Int))
					default:
						return s.RaiseError("expected 0 or 3 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_Vector3int16)},
							},
							Summary:     "Types/Vector3int16:Constructors/new/Zero/Summary",
							Description: "Types/Vector3int16:Constructors/new/Zero/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_Int)},
								{Name: "y", Type: dt.Prim(rtypes.T_Int)},
								{Name: "z", Type: dt.Prim(rtypes.T_Int)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_Vector3int16)},
							},
							Summary:     "Types/Vector3int16:Constructors/new/Components/Summary",
							Description: "Types/Vector3int16:Constructors/new/Components/Description",
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
						Summary:     "Types/Vector3int16:Operators/Eq/Summary",
						Description: "Types/Vector3int16:Operators/Eq/Description",
					},
					Add: []dump.Binop{
						{
							Operand:     dt.Prim(rtypes.T_Vector3int16),
							Result:      dt.Prim(rtypes.T_Vector3int16),
							Summary:     "Types/Vector3int16:Operators/Add/Summary",
							Description: "Types/Vector3int16:Operators/Add/Description",
						},
					},
					Sub: []dump.Binop{
						{
							Operand:     dt.Prim(rtypes.T_Vector3int16),
							Result:      dt.Prim(rtypes.T_Vector3int16),
							Summary:     "Types/Vector3int16:Operators/Sub/Summary",
							Description: "Types/Vector3int16:Operators/Sub/Description",
						},
					},
					Mul: []dump.Binop{
						{
							Operand:     dt.Prim(rtypes.T_Vector3int16),
							Result:      dt.Prim(rtypes.T_Vector3int16),
							Summary:     "Types/Vector3int16:Operators/Mul/Vector3int16/Summary",
							Description: "Types/Vector3int16:Operators/Mul/Vector3int16/Description",
						},
						{
							Operand:     dt.Prim(rtypes.T_Number),
							Result:      dt.Prim(rtypes.T_Vector3int16),
							Summary:     "Types/Vector3int16:Operators/Mul/Number/Summary",
							Description: "Types/Vector3int16:Operators/Mul/Number/Description",
						},
					},
					Div: []dump.Binop{
						{
							Operand:     dt.Prim(rtypes.T_Vector3int16),
							Result:      dt.Prim(rtypes.T_Vector3int16),
							Summary:     "Types/Vector3int16:Operators/Div/Vector3int16/Summary",
							Description: "Types/Vector3int16:Operators/Div/Vector3int16/Description",
						},
						{
							Operand:     dt.Prim(rtypes.T_Number),
							Result:      dt.Prim(rtypes.T_Vector3int16),
							Summary:     "Types/Vector3int16:Operators/Div/Number/Summary",
							Description: "Types/Vector3int16:Operators/Div/Number/Description",
						},
					},
					Unm: &dump.Unop{
						Result:      dt.Prim(rtypes.T_Vector3int16),
						Summary:     "Types/Vector3int16:Operators/Unm/Summary",
						Description: "Types/Vector3int16:Operators/Unm/Description",
					},
				},
				Summary:     "Types/Vector3int16:Summary",
				Description: "Types/Vector3int16:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Int,
			Number,
		},
	}
}
