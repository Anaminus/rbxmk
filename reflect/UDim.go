package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

const T_UDim = "UDim"

func init() { register(UDim) }
func UDim() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     T_UDim,
		PushTo:   rbxmk.PushTypeTo(T_UDim),
		PullFrom: rbxmk.PullTypeFrom(T_UDim),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.UDim:
				*p = v.(types.UDim)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, T_UDim).(types.UDim)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, T_UDim).(types.UDim)
				op := s.Pull(2, T_UDim).(types.UDim)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
			"__add": func(s rbxmk.State) int {
				v := s.Pull(1, T_UDim).(types.UDim)
				op := s.Pull(2, T_UDim).(types.UDim)
				return s.Push(v.Add(op))
			},
			"__sub": func(s rbxmk.State) int {
				v := s.Pull(1, T_UDim).(types.UDim)
				op := s.Pull(2, T_UDim).(types.UDim)
				return s.Push(v.Sub(op))
			},
			"__unm": func(s rbxmk.State) int {
				v := s.Pull(1, T_UDim).(types.UDim)
				return s.Push(v.Neg())
			},
		},
		Properties: rbxmk.Properties{
			"Scale": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.UDim).Scale))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Float),
						ReadOnly:    true,
						Summary:     "Types/UDim:Properties/Scale/Summary",
						Description: "Types/UDim:Properties/Scale/Description",
					}
				},
			},
			"Offset": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Int(v.(types.UDim).Offset))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Int),
						ReadOnly:    true,
						Summary:     "Types/UDim:Properties/Offset/Summary",
						Description: "Types/UDim:Properties/Offset/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.UDim{
						Scale:  float32(s.Pull(1, T_Float).(types.Float)),
						Offset: int32(s.Pull(2, T_Int).(types.Int)),
					})
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "scale", Type: dt.Prim(T_Float)},
								{Name: "offset", Type: dt.Prim(T_Int)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_UDim)},
							},
							Summary:     "Types/UDim:Constructors/new/Summary",
							Description: "Types/UDim:Constructors/new/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Eq: &dump.Cmpop{
						Summary:     "Types/UDim:Operators/Eq/Summary",
						Description: "Types/UDim:Operators/Eq/Description",
					},
					Add: []dump.Binop{
						{
							Operand:     dt.Prim(T_UDim),
							Result:      dt.Prim(T_UDim),
							Summary:     "Types/UDim:Operators/Add/Summary",
							Description: "Types/UDim:Operators/Add/Description",
						},
					},
					Sub: []dump.Binop{
						{
							Operand:     dt.Prim(T_UDim),
							Result:      dt.Prim(T_UDim),
							Summary:     "Types/UDim:Operators/Sub/Summary",
							Description: "Types/UDim:Operators/Sub/Description",
						},
					},
					Unm: &dump.Unop{
						Result:      dt.Prim(T_UDim),
						Summary:     "Types/UDim:Operators/Unm/Summary",
						Description: "Types/UDim:Operators/Unm/Description",
					},
				},
				Summary:     "Types/UDim:Summary",
				Description: "Types/UDim:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Float,
			Int,
		},
	}
}
