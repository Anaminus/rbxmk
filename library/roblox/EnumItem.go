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
		Properties: rbxmk.Properties{
			"Name": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.String(v.(*rtypes.EnumItem).Name()))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						ReadOnly:    true,
						Summary:     "libraries/roblox/types/EnumItem:Properties/Name/Summary",
						Description: "libraries/roblox/types/EnumItem:Properties/Name/Description",
					}
				},
			},
			"Value": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Int(v.(*rtypes.EnumItem).Value()))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("int"),
						ReadOnly:    true,
						Summary:     "libraries/roblox/types/EnumItem:Properties/Value/Summary",
						Description: "libraries/roblox/types/EnumItem:Properties/Value/Description",
					}
				},
			},
			"EnumType": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(*rtypes.EnumItem).Enum())
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("Enum"),
						ReadOnly:    true,
						Summary:     "libraries/roblox/types/EnumItem:Properties/EnumType/Summary",
						Description: "libraries/roblox/types/EnumItem:Properties/EnumType/Description",
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "$TODO",
				Description: "$TODO",
			}
		},
	}
}
