package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(UDim2) }
func UDim2() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     rtypes.T_UDim2,
		PushTo:   rbxmk.PushTypeTo(rtypes.T_UDim2),
		PullFrom: rbxmk.PullTypeFrom(rtypes.T_UDim2),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.UDim2:
				*p = v.(types.UDim2)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_UDim2).(types.UDim2)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_UDim2).(types.UDim2)
				op := s.Pull(2, rtypes.T_UDim2).(types.UDim2)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
			"__add": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_UDim2).(types.UDim2)
				op := s.Pull(2, rtypes.T_UDim2).(types.UDim2)
				return s.Push(v.Add(op))
			},
			"__sub": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_UDim2).(types.UDim2)
				op := s.Pull(2, rtypes.T_UDim2).(types.UDim2)
				return s.Push(v.Sub(op))
			},
			"__unm": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_UDim2).(types.UDim2)
				return s.Push(v.Neg())
			},
		},
		Properties: rbxmk.Properties{
			"X": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.UDim2).X)
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_UDim),
						ReadOnly:    true,
						Summary:     "Types/UDim2:Properties/X/Summary",
						Description: "Types/UDim2:Properties/X/Description",
					}
				},
			},
			"Y": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.UDim2).Y)
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_UDim),
						ReadOnly:    true,
						Summary:     "Types/UDim2:Properties/Y/Summary",
						Description: "Types/UDim2:Properties/Y/Description",
					}
				},
			},
			"Width": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.UDim2).X)
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_UDim),
						ReadOnly:    true,
						Summary:     "Types/UDim2:Properties/Width/Summary",
						Description: "Types/UDim2:Properties/Width/Description",
					}
				},
			},
			"Height": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.UDim2).Y)
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_UDim),
						ReadOnly:    true,
						Summary:     "Types/UDim2:Properties/Height/Summary",
						Description: "Types/UDim2:Properties/Height/Description",
					}
				},
			},
		},
		Methods: rbxmk.Methods{
			"Lerp": {
				Func: func(s rbxmk.State, v types.Value) int {
					goal := s.Pull(2, rtypes.T_UDim2).(types.UDim2)
					alpha := float64(s.Pull(3, rtypes.T_Float).(types.Float))
					return s.Push(v.(types.UDim2).Lerp(goal, alpha))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "goal", Type: dt.Prim(rtypes.T_UDim2)},
							{Name: "alpha", Type: dt.Prim(rtypes.T_Float)},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(rtypes.T_UDim2)},
						},
						Summary:     "Types/UDim2:Methods/Lerp/Summary",
						Description: "Types/UDim2:Methods/Lerp/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.UDim2
					switch s.Count() {
					case 2:
						v.X = s.Pull(1, rtypes.T_UDim).(types.UDim)
						v.Y = s.Pull(2, rtypes.T_UDim).(types.UDim)
					case 4:
						v.X.Scale = float32(s.Pull(1, rtypes.T_Float).(types.Float))
						v.X.Offset = int32(s.Pull(2, rtypes.T_Int).(types.Int))
						v.Y.Scale = float32(s.Pull(3, rtypes.T_Float).(types.Float))
						v.Y.Offset = int32(s.Pull(4, rtypes.T_Int).(types.Int))
					default:
						return s.RaiseError("expected 2 or 4 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "xScale", Type: dt.Prim(rtypes.T_Float)},
								{Name: "xOffset", Type: dt.Prim(rtypes.T_Int)},
								{Name: "yScale", Type: dt.Prim(rtypes.T_Float)},
								{Name: "yOffset", Type: dt.Prim(rtypes.T_Int)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_UDim2)},
							},
							Summary:     "Types/UDim2:Constructors/new/Components/Summary",
							Description: "Types/UDim2:Constructors/new/Components/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_UDim)},
								{Name: "y", Type: dt.Prim(rtypes.T_UDim)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_UDim2)},
							},
							Summary:     "Types/UDim2:Constructors/new/UDim/Summary",
							Description: "Types/UDim2:Constructors/new/UDim/Description",
						},
					}
				},
			},
			"fromScale": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.UDim2{
						X: types.UDim{Scale: float32(s.Pull(1, rtypes.T_Float).(types.Float))},
						Y: types.UDim{Scale: float32(s.Pull(2, rtypes.T_Float).(types.Float))},
					})
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_Float)},
								{Name: "y", Type: dt.Prim(rtypes.T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_UDim2)},
							},
							Summary:     "Types/UDim2:Constructors/fromScale/Summary",
							Description: "Types/UDim2:Constructors/fromScale/Description",
						},
					}
				},
			},
			"fromOffset": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.UDim2{
						X: types.UDim{Offset: int32(s.Pull(1, rtypes.T_Int).(types.Int))},
						Y: types.UDim{Offset: int32(s.Pull(2, rtypes.T_Int).(types.Int))},
					})
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_Int)},
								{Name: "y", Type: dt.Prim(rtypes.T_Int)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_UDim2)},
							},
							Summary:     "Types/UDim2:Constructors/fromOffset/Summary",
							Description: "Types/UDim2:Constructors/fromOffset/Description",
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
						Summary:     "Types/UDim2:Operators/Eq/Summary",
						Description: "Types/UDim2:Operators/Eq/Description",
					},
					Add: []dump.Binop{
						{
							Operand:     dt.Prim(rtypes.T_UDim2),
							Result:      dt.Prim(rtypes.T_UDim2),
							Summary:     "Types/UDim2:Operators/Add/Summary",
							Description: "Types/UDim2:Operators/Add/Description",
						},
					},
					Sub: []dump.Binop{
						{
							Operand:     dt.Prim(rtypes.T_UDim2),
							Result:      dt.Prim(rtypes.T_UDim2),
							Summary:     "Types/UDim2:Operators/Sub/Summary",
							Description: "Types/UDim2:Operators/Sub/Description",
						},
					},
					Unm: &dump.Unop{
						Result:      dt.Prim(rtypes.T_UDim2),
						Summary:     "Types/UDim2:Operators/Unm/Summary",
						Description: "Types/UDim2:Operators/Unm/Description",
					},
				},
				Summary:     "Types/UDim2:Summary",
				Description: "Types/UDim2:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Float,
			Int,
			UDim,
		},
	}
}
