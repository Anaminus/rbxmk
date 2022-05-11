package reflect

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

const T_AttrConfig = "AttrConfig"

func init() { register(AttrConfig) }
func AttrConfig() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     T_AttrConfig,
		PushTo:   rbxmk.PushPtrTypeTo(T_AttrConfig),
		PullFrom: rbxmk.PullTypeFrom(T_AttrConfig),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case **rtypes.AttrConfig:
				*p = v.(*rtypes.AttrConfig)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Properties: rbxmk.Properties{
			"Property": {
				Get: func(s rbxmk.State, v types.Value) int {
					attrConfig := v.(*rtypes.AttrConfig)
					return s.Push(types.String(attrConfig.Property))
				},
				Set: func(s rbxmk.State, v types.Value) {
					attrConfig := v.(*rtypes.AttrConfig)
					attrConfig.Property = string(s.Pull(3, T_String).(types.String))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_String),
						Summary:     "Types/AttrConfig:Properties/Property/Summary",
						Description: "Types/AttrConfig:Properties/Property/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": rbxmk.Constructor{
				Func: func(s rbxmk.State) int {
					var v rtypes.AttrConfig
					v.Property = string(s.PullOpt(1, types.String(""), T_String).(types.String))
					return s.Push(&v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						dump.Function{
							Parameters: dump.Parameters{
								{Name: "property", Type: dt.Prim(T_String)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_AttrConfig)},
							},
							Summary:     "Types/AttrConfig:Constructors/new/Summary",
							Description: "Types/AttrConfig:Constructors/new/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/AttrConfig:Summary",
				Description: "Types/AttrConfig:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			String,
		},
	}
}
