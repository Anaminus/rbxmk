package reflect

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Enums) }
func Enums() Reflector {
	return Reflector{
		Name:     "Enums",
		PushTo:   rbxmk.PushTypeTo("Enums"),
		PullFrom: rbxmk.PullTypeFrom("Enums"),
		Metatable: Metatable{
			"__index": func(s State) int {
				name := string(s.Pull(2, "string").(types.String))

				if name == "GetEnums" {
					s.L.Push(s.WrapFunc(func(s State) int {
						enums := s.Pull(1, "Enums").(*rtypes.Enums)
						es := enums.Enums()
						array := make(rtypes.Array, len(es))
						for i, enum := range es {
							array[i] = enum
						}
						return s.Push(array)
					}))
					return 1
				}

				enums := s.Pull(1, "Enums").(*rtypes.Enums)
				enum := enums.Enum(name)
				if enum == nil {
					return s.RaiseError("%s is not a valid Enum", name)
				}
				return s.Push(enum)
			},
			"__newindex": func(s State) int {
				name := string(s.Pull(2, "string").(types.String))
				return s.RaiseError("%s cannot be assigned to", name)
			},
		},
	}
}
