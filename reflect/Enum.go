package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func Enum() Type {
	return Type{
		Name:     "Enum",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__index": func(s State) int {
				name := string(s.Pull(2, "string").(types.String))

				if name == "GetEnumItems" {
					s.L.Push(s.WrapFunc(func(s State) int {
						enum := s.Pull(1, "Enum").(*rtypes.Enum)
						items := enum.Items()
						array := make(rtypes.Array, len(items))
						for i, item := range items {
							array[i] = item
						}
						return s.Push(array)
					}))
					return 1
				}

				enum := s.Pull(1, "Enum").(*rtypes.Enum)
				item := enum.Item(name)
				if item == nil {
					s.L.RaiseError("%s is not a valid EnumItem", name)
					return 0
				}
				return s.Push(item)
			},
			"__newindex": func(s State) int {
				name := string(s.Pull(2, "string").(types.String))
				s.L.RaiseError("%s cannot be assigned to", name)
				return 0
			},
		},
	}
}
