package reflect

import (
	lua "github.com/anaminus/gopher-lua"
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
		Metatable: rbxmk.Metatable{
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "AttrConfig").(*rtypes.AttrConfig)
				op := s.Pull(2, "AttrConfig").(*rtypes.AttrConfig)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: rbxmk.Members{
			"Property": {
				Get: func(s rbxmk.State, v types.Value) int {
					attrConfig := v.(*rtypes.AttrConfig)
					return s.Push(types.String(attrConfig.Property))
				},
				Set: func(s rbxmk.State, v types.Value) {
					attrConfig := v.(*rtypes.AttrConfig)
					attrConfig.Property = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("string")} },
			},
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
