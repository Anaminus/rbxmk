package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(EnumItem) }
func EnumItem() Reflector {
	return Reflector{
		Name:     "EnumItem",
		PushTo:   rbxmk.PushTypeTo("EnumItem"),
		PullFrom: rbxmk.PullTypeFrom("EnumItem"),
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "EnumItem").(*rtypes.EnumItem)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "EnumItem").(*rtypes.EnumItem)
				op := s.Pull(2, "EnumItem").(*rtypes.EnumItem)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: rbxmk.Members{
			"Name": Member{
				Get: func(s State, v types.Value) int {
					return s.Push(types.String(v.(*rtypes.EnumItem).Name()))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("string"), ReadOnly: true} },
			},
			"Value": Member{
				Get: func(s State, v types.Value) int {
					return s.Push(types.Int(v.(*rtypes.EnumItem).Value()))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("int"), ReadOnly: true} },
			},
			"EnumType": Member{
				Get: func(s State, v types.Value) int {
					return s.Push(v.(*rtypes.EnumItem).Enum())
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("Enum"), ReadOnly: true} },
			},
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
