package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(TypeDesc) }
func TypeDesc() Reflector {
	return Reflector{
		Name:     "TypeDesc",
		PushTo:   rbxmk.PushTypeTo("TypeDesc"),
		PullFrom: rbxmk.PullTypeFrom("TypeDesc"),
		Metatable: Metatable{
			"__eq": func(s State) int {
				v := s.Pull(1, "TypeDesc").(rtypes.TypeDesc)
				op := s.Pull(2, "TypeDesc").(rtypes.TypeDesc)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: Members{
			"Category": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.TypeDesc)
					return s.Push(types.String(desc.Embedded.Category))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("string"), ReadOnly: true} },
			},
			"Name": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.TypeDesc)
					return s.Push(types.String(desc.Embedded.Name))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("string"), ReadOnly: true} },
			},
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
