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
func AttrConfig() Reflector {
	return Reflector{
		Name:     "AttrConfig",
		PushTo:   rbxmk.PushTypeTo("AttrConfig"),
		PullFrom: rbxmk.PullTypeFrom("AttrConfig"),
		Metatable: Metatable{
			"__eq": func(s State) int {
				v := s.Pull(1, "AttrConfig").(*rtypes.AttrConfig)
				op := s.Pull(2, "AttrConfig").(*rtypes.AttrConfig)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: Members{
			"Property": Member{
				Get: func(s State, v types.Value) int {
					attrConfig := v.(*rtypes.AttrConfig)
					return s.Push(types.String(attrConfig.Property))
				},
				Set: func(s State, v types.Value) {
					attrConfig := v.(*rtypes.AttrConfig)
					attrConfig.Property = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("string")} },
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s State) int {
					var v rtypes.AttrConfig
					v.Property = string(s.PullOpt(1, "string", types.String("")).(types.String))
					return s.Push(&v)
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Parameters: dump.Parameters{
							{Name: "property", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("AttrConfig")},
						},
					}}
				},
			},
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
