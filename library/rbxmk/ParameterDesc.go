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
func ParameterDesc() Reflector {
	return Reflector{
		Name:     "ParameterDesc",
		PushTo:   rbxmk.PushTypeTo("ParameterDesc"),
		PullFrom: rbxmk.PullTypeFrom("ParameterDesc"),
		Metatable: Metatable{
			"__eq": func(s State) int {
				v := s.Pull(1, "ParameterDesc").(rtypes.ParameterDesc)
				op := s.Pull(2, "ParameterDesc").(rtypes.ParameterDesc)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: Members{
			"Type": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.ParameterDesc)
					typ := desc.Parameter.Type
					return s.Push(rtypes.TypeDesc{Embedded: typ})
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("TypeDesc"), ReadOnly: true} },
			},
			"Name": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.ParameterDesc)
					return s.Push(types.String(desc.Name))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("string"), ReadOnly: true} },
			},
			"Default": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.ParameterDesc)
					if !desc.Optional {
						return s.Push(rtypes.Nil)
					}
					return s.Push(types.String(desc.Default))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Optional{T: dt.Prim("string")}, ReadOnly: true} },
			},
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
