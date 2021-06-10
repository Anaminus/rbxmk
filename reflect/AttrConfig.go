package reflect

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(AttrConfig) }
func AttrConfig() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "AttrConfig",
		PushTo:   rbxmk.PushPtrTypeTo("AttrConfig"),
		PullFrom: rbxmk.PullTypeFrom("AttrConfig"),
		Properties: rbxmk.Properties{
			"Property": {
				Get: func(s rbxmk.State, v types.Value) int {
					attrConfig := v.(*rtypes.AttrConfig)
					return s.Push(types.String(attrConfig.Property))
				},
				Set: func(s rbxmk.State, v types.Value) {
					attrConfig := v.(*rtypes.AttrConfig)
					attrConfig.Property = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
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
					v.Property = string(s.PullOpt(1, types.String(""), "string").(types.String))
					return s.Push(&v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						dump.Function{
							Parameters: dump.Parameters{
								{Name: "property", Type: dt.Prim("string")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("AttrConfig")},
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
