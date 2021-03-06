package reflect

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Enum) }
func Enum() Reflector {
	return Reflector{
		Name:     "Enum",
		PushTo:   rbxmk.PushTypeTo("Enum"),
		PullFrom: rbxmk.PullTypeFrom("Enum"),
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
					return s.RaiseError("%s is not a valid EnumItem", name)
				}
				return s.Push(item)
			},
			"__newindex": func(s State) int {
				name := string(s.Pull(2, "string").(types.String))
				return s.RaiseError("%s cannot be assigned to", name)
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Methods: dump.Methods{
					"GetEnumItems": dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Array{T: dt.Prim("EnumItem")}},
						},
					},
				},
				Operators: &dump.Operators{
					Index: dump.Function{
						Parameters: dump.Parameters{
							{Name: "name", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("EnumItem")},
						},
					},
				},
			}
		},
	}
}
