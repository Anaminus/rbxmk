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
func EnumItem() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "EnumItem",
		PushTo:   rbxmk.PushPtrTypeTo("EnumItem"),
		PullFrom: rbxmk.PullTypeFrom("EnumItem"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "EnumItem").(*rtypes.EnumItem)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
		},
		Members: rbxmk.Members{
			"Name": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.String(v.(*rtypes.EnumItem).Name()))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("string"), ReadOnly: true} },
			},
			"Value": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Int(v.(*rtypes.EnumItem).Value()))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("int"), ReadOnly: true} },
			},
			"EnumType": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(*rtypes.EnumItem).Enum())
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("Enum"), ReadOnly: true} },
			},
		},
	}
}
