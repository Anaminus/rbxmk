package reflect

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
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
			"__eq": func(s State) int {
				v := s.Pull(1, "Enums").(*rtypes.Enums)
				op := s.Pull(2, "Enums").(*rtypes.Enums)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
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
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Methods: dump.Methods{
					"GetEnums": dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Array{T: dt.Prim("Enum")}},
						},
					},
				},
				Operators: &dump.Operators{
					Eq: true,
					Index: dump.Function{
						Parameters: dump.Parameters{
							{Name: "name", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Enum")},
						},
					},
				},
			}
		},
	}
}
