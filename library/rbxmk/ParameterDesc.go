package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(ParameterDesc) }
func ParameterDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "ParameterDesc",
		PushTo:   rbxmk.PushTypeTo("ParameterDesc"),
		PullFrom: rbxmk.PullTypeFrom("ParameterDesc"),
		Metatable: rbxmk.Metatable{
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "ParameterDesc").(rtypes.ParameterDesc)
				op := s.Pull(2, "ParameterDesc").(rtypes.ParameterDesc)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"Type": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.ParameterDesc)
					typ := desc.Parameter.Type
					return s.Push(rtypes.TypeDesc{Embedded: typ})
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("TypeDesc"), ReadOnly: true} },
			},
			"Name": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.ParameterDesc)
					return s.Push(types.String(desc.Name))
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("string"), ReadOnly: true} },
			},
			"Default": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.ParameterDesc)
					if !desc.Optional {
						return s.Push(rtypes.Nil)
					}
					return s.Push(types.String(desc.Default))
				},
				Dump: func() dump.Property {
					return dump.Property{ValueType: dt.Optional{T: dt.Prim("string")}, ReadOnly: true}
				},
			},
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
