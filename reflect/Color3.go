package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Color3) }
func Color3() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     rtypes.T_Color3,
		PushTo:   rbxmk.PushTypeTo(rtypes.T_Color3),
		PullFrom: rbxmk.PullTypeFrom(rtypes.T_Color3),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.Color3:
				*p = v.(types.Color3)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Color3).(types.Color3)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Color3).(types.Color3)
				op := s.Pull(2, rtypes.T_Color3).(types.Color3)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"R": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Color3).R))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Float),
						ReadOnly:    true,
						Summary:     "Types/Color3:Properties/R/Summary",
						Description: "Types/Color3:Properties/R/Description",
					}
				},
			},
			"G": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Color3).G))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Float),
						ReadOnly:    true,
						Summary:     "Types/Color3:Properties/G/Summary",
						Description: "Types/Color3:Properties/G/Description",
					}
				},
			},
			"B": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Color3).B))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Float),
						ReadOnly:    true,
						Summary:     "Types/Color3:Properties/B/Summary",
						Description: "Types/Color3:Properties/B/Description",
					}
				},
			},
		},
		Methods: rbxmk.Methods{
			"Lerp": {
				Func: func(s rbxmk.State, v types.Value) int {
					goal := s.Pull(2, rtypes.T_Color3).(types.Color3)
					alpha := float64(s.Pull(3, rtypes.T_Float).(types.Float))
					return s.Push(v.(types.Color3).Lerp(goal, alpha))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "goal", Type: dt.Prim(rtypes.T_Color3)},
							{Name: "alpha", Type: dt.Prim(rtypes.T_Float)},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(rtypes.T_Color3)},
						},
						Summary:     "Types/Color3:Methods/Lerp/Summary",
						Description: "Types/Color3:Methods/Lerp/Description",
					}
				},
			},
			"ToHSV": {
				Func: func(s rbxmk.State, v types.Value) int {
					hue, sat, val := v.(types.Color3).ToHSV()
					return s.PushTuple(types.Float(hue), types.Float(sat), types.Float(val))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Name: "h", Type: dt.Prim(rtypes.T_Float)},
							{Name: "s", Type: dt.Prim(rtypes.T_Float)},
							{Name: "v", Type: dt.Prim(rtypes.T_Float)},
						},
						Summary:     "Types/Color3:Methods/ToHSV/Summary",
						Description: "Types/Color3:Methods/ToHSV/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.Color3
					switch s.Count() {
					case 0:
					case 3:
						v.R = float32(s.Pull(1, rtypes.T_Float).(types.Float))
						v.G = float32(s.Pull(2, rtypes.T_Float).(types.Float))
						v.B = float32(s.Pull(3, rtypes.T_Float).(types.Float))
					default:
						return s.RaiseError("expected 0 or 3 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_Color3)},
							},
							Summary:     "Types/Color3:Constructors/new/Zero/Summary",
							Description: "Types/Color3:Constructors/new/Zero/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "r", Type: dt.Prim(rtypes.T_Float)},
								{Name: "g", Type: dt.Prim(rtypes.T_Float)},
								{Name: "b", Type: dt.Prim(rtypes.T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_Color3)},
							},
							Summary:     "Types/Color3:Constructors/new/Components/Summary",
							Description: "Types/Color3:Constructors/new/Components/Description",
						},
					}
				},
			},
			"fromRGB": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewColor3FromRGB(
						int(s.Pull(1, rtypes.T_Int).(types.Int)),
						int(s.Pull(2, rtypes.T_Int).(types.Int)),
						int(s.Pull(3, rtypes.T_Int).(types.Int)),
					))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "r", Type: dt.Prim(rtypes.T_Int)},
								{Name: "g", Type: dt.Prim(rtypes.T_Int)},
								{Name: "b", Type: dt.Prim(rtypes.T_Int)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_Color3)},
							},
							Summary:     "Types/Color3:Constructors/fromRGB/Summary",
							Description: "Types/Color3:Constructors/fromRGB/Description",
						},
					}
				},
			},
			"fromHSV": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewColor3FromHSV(
						float64(s.Pull(1, rtypes.T_Float).(types.Float)),
						float64(s.Pull(2, rtypes.T_Float).(types.Float)),
						float64(s.Pull(3, rtypes.T_Float).(types.Float)),
					))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "h", Type: dt.Prim(rtypes.T_Float)},
								{Name: "s", Type: dt.Prim(rtypes.T_Float)},
								{Name: "v", Type: dt.Prim(rtypes.T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_Color3)},
							},
							Summary:     "Types/Color3:Constructors/fromHSV/Summary",
							Description: "Types/Color3:Constructors/fromHSV/Description",
						},
					}
				},
			},
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case types.Color3:
				return v
			case rtypes.Color3uint8:
				return types.Color3(v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Eq: &dump.Cmpop{
						Summary:     "Types/Color3:Operators/Eq/Summary",
						Description: "Types/Color3:Operators/Eq/Description",
					},
				},
				Summary:     "Types/Color3:Summary",
				Description: "Types/Color3:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Float,
			Int,
		},
	}
}
