package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func Enums() Type {
	return Type{
		Name:     "Enums",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__index": func(s State) int {
				name := string(s.Pull(2, "string").(types.String))

				if name == "GetEnums" {
					s.L.Push(s.WrapFunc(func(s State) int {
						enums := s.Pull(1, "Enums").(rtypes.Enums)
						es := enums.Enums()
						array := make(rtypes.Array, len(es))
						for i, enum := range es {
							array[i] = enum
						}
						return s.Push(array)
					}))
					return 1
				}

				enums := s.Pull(1, "Enums").(rtypes.Enums)
				enum := enums.Enum(name)
				if enum == nil {
					s.L.RaiseError("%s is not a valid Enum", name)
					return 0
				}
				return s.Push(enum)
			},
			"__newindex": func(s State) int {
				name := string(s.Pull(2, "string").(types.String))
				s.L.RaiseError("%s cannot be assigned to", name)
				return 0
			},
		},
	}
}
