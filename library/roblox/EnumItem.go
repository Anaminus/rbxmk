package reflect

import (
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
			"__eq": func(s State) int {
				v := s.Pull(1, "EnumItem").(*rtypes.EnumItem)
				op := s.Pull(2, "EnumItem").(*rtypes.EnumItem)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
			"__index": func(s State) int {
				item := s.Pull(1, "EnumItem").(*rtypes.EnumItem)
				name := string(s.Pull(2, "string").(types.String))
				switch name {
				case "Name":
					return s.Push(types.String(item.Name()))
				case "Value":
					return s.Push(types.Int(item.Value()))
				case "EnumType":
					return s.Push(item.Enum())
				}
				return s.RaiseError("%s is not a valid member", name)
			},
			"__newindex": func(s State) int {
				name := string(s.Pull(2, "string").(types.String))
				return s.RaiseError("%s cannot be assigned to", name)
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Properties: dump.Properties{
					"Name":     dump.Property{ValueType: dt.Prim("string"), ReadOnly: true},
					"Value":    dump.Property{ValueType: dt.Prim("int"), ReadOnly: true},
					"EnumType": dump.Property{ValueType: dt.Prim("Enum"), ReadOnly: true},
				},
				Operators: &dump.Operators{Eq: true},
			}
		},
	}
}
