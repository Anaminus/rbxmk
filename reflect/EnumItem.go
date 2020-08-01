package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(EnumItem) }
func EnumItem() Reflector {
	return Reflector{
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
				return s.RaiseError("%s is not a valid member", name)
			},
			"__newindex": func(s State) int {
				name := string(s.Pull(2, "string").(types.String))
				return s.RaiseError("%s cannot be assigned to", name)
			},
		},
	}
}
