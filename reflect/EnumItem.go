package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func EnumItem() Type {
	return Type{
		Name:     "EnumItem",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
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
				s.L.RaiseError("%s is not a valid member", name)
				return 0
			},
			"__newindex": func(s State) int {
				name := string(s.Pull(2, "string").(types.String))
				s.L.RaiseError("%s cannot be assigned to", name)
				return 0
			},
		},
	}
}
