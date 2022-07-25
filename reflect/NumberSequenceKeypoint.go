package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(NumberSequenceKeypoint) }
func NumberSequenceKeypoint() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     rtypes.T_NumberSequenceKeypoint,
		PushTo:   rbxmk.PushTypeTo(rtypes.T_NumberSequenceKeypoint),
		PullFrom: rbxmk.PullTypeFrom(rtypes.T_NumberSequenceKeypoint),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.NumberSequenceKeypoint:
				*p = v.(types.NumberSequenceKeypoint)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_NumberSequenceKeypoint).(types.NumberSequenceKeypoint)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_NumberSequenceKeypoint).(types.NumberSequenceKeypoint)
				op := s.Pull(2, rtypes.T_NumberSequenceKeypoint).(types.NumberSequenceKeypoint)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"Time": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.NumberSequenceKeypoint).Time))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Float),
						ReadOnly:    true,
						Summary:     "Types/NumberSequenceKeypoint:Properties/Time/Summary",
						Description: "Types/NumberSequenceKeypoint:Properties/Time/Description",
					}
				},
			},
			"Value": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.NumberSequenceKeypoint).Value))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Float),
						ReadOnly:    true,
						Summary:     "Types/NumberSequenceKeypoint:Properties/Value/Summary",
						Description: "Types/NumberSequenceKeypoint:Properties/Value/Description",
					}
				},
			},
			"Envelope": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.NumberSequenceKeypoint).Envelope))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Float),
						ReadOnly:    true,
						Summary:     "Types/NumberSequenceKeypoint:Properties/Envelope/Summary",
						Description: "Types/NumberSequenceKeypoint:Properties/Envelope/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.NumberSequenceKeypoint
					switch s.Count() {
					case 2:
						v.Time = float32(s.Pull(1, rtypes.T_Float).(types.Float))
						v.Value = float32(s.Pull(2, rtypes.T_Float).(types.Float))
					case 3:
						v.Time = float32(s.Pull(1, rtypes.T_Float).(types.Float))
						v.Value = float32(s.Pull(2, rtypes.T_Float).(types.Float))
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
								{Name: "value", Type: dt.Prim(rtypes.T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_NumberSequenceKeypoint)},
							},
							Summary:     "Types/NumberSequenceKeypoint:Constructors/new/Components/Summary",
							Description: "Types/NumberSequenceKeypoint:Constructors/new/Components/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "time", Type: dt.Prim(rtypes.T_Float)},
								{Name: "value", Type: dt.Prim(rtypes.T_Float)},
								{Name: "envelope", Type: dt.Prim(rtypes.T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_NumberSequenceKeypoint)},
							},
							Summary:     "Types/NumberSequenceKeypoint:Constructors/new/Envelope/Summary",
							Description: "Types/NumberSequenceKeypoint:Constructors/new/Envelope/Description",
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
						Summary:     "Types/NumberSequenceKeypoint:Operators/Eq/Summary",
						Description: "Types/NumberSequenceKeypoint:Operators/Eq/Description",
					},
				},
				Summary:     "Types/NumberSequenceKeypoint:Summary",
				Description: "Types/NumberSequenceKeypoint:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Float,
		},
	}
}
