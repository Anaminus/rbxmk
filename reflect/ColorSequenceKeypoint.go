package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(ColorSequenceKeypoint) }
func ColorSequenceKeypoint() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     rtypes.T_ColorSequenceKeypoint,
		PushTo:   rbxmk.PushTypeTo(rtypes.T_ColorSequenceKeypoint),
		PullFrom: rbxmk.PullTypeFrom(rtypes.T_ColorSequenceKeypoint),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.ColorSequenceKeypoint:
				*p = v.(types.ColorSequenceKeypoint)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_ColorSequenceKeypoint).(types.ColorSequenceKeypoint)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_ColorSequenceKeypoint).(types.ColorSequenceKeypoint)
				op := s.Pull(2, rtypes.T_ColorSequenceKeypoint).(types.ColorSequenceKeypoint)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"Time": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.ColorSequenceKeypoint).Time))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Float),
						ReadOnly:    true,
						Summary:     "Types/ColorSequenceKeypoint:Properties/Time/Summary",
						Description: "Types/ColorSequenceKeypoint:Properties/Time/Description",
					}
				},
			},
			"Value": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.ColorSequenceKeypoint).Value)
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Color3),
						ReadOnly:    true,
						Summary:     "Types/ColorSequenceKeypoint:Properties/Value/Summary",
						Description: "Types/ColorSequenceKeypoint:Properties/Value/Description",
					}
				},
			},
			"Envelope": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.ColorSequenceKeypoint).Envelope))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Float),
						ReadOnly:    true,
						Summary:     "Types/ColorSequenceKeypoint:Properties/Envelope/Summary",
						Description: "Types/ColorSequenceKeypoint:Properties/Envelope/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.ColorSequenceKeypoint
					switch s.Count() {
					case 2:
						v.Time = float32(s.Pull(1, rtypes.T_Float).(types.Float))
						v.Value = s.Pull(2, rtypes.T_Color3).(types.Color3)
					case 3:
						v.Time = float32(s.Pull(1, rtypes.T_Float).(types.Float))
						v.Value = s.Pull(2, rtypes.T_Color3).(types.Color3)
						v.Envelope = float32(s.Pull(3, rtypes.T_Float).(types.Float))
					default:
						return s.RaiseError("expected 2 or 3 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "time", Type: dt.Prim(rtypes.T_Float)},
								{Name: "color", Type: dt.Prim(rtypes.T_Color3)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_ColorSequenceKeypoint)},
							},
							Summary:     "Types/ColorSequenceKeypoint:Constructors/new/Components/Summary",
							Description: "Types/ColorSequenceKeypoint:Constructors/new/Components/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "time", Type: dt.Prim(rtypes.T_Float)},
								{Name: "color", Type: dt.Prim(rtypes.T_Color3)},
								{Name: "envelope", Type: dt.Prim(rtypes.T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_ColorSequenceKeypoint)},
							},
							Summary:     "Types/ColorSequenceKeypoint:Constructors/new/Envelope/Summary",
							Description: "Types/ColorSequenceKeypoint:Constructors/new/Envelope/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Eq: &dump.Cmpop{
						Summary:     "Types/ColorSequenceKeypoint:Operators/Eq/Summary",
						Description: "Types/ColorSequenceKeypoint:Operators/Eq/Description",
					},
				},
				Summary:     "Types/ColorSequenceKeypoint:Summary",
				Description: "Types/ColorSequenceKeypoint:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Color3,
			Float,
		},
	}
}
