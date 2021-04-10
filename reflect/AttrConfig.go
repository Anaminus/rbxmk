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
