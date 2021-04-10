package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Enum) }
func Enum() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "Enum",
		PushTo:   rbxmk.PushPtrTypeTo("Enum"),
		PullFrom: rbxmk.PullTypeFrom("Enum"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "Enum").(*rtypes.Enum)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__index": func(s rbxmk.State) int {
				enum := s.Pull(1, "Enum").(*rtypes.Enum)
				name := string(s.Pull(2, "string").(types.String))
				item := enum.Item(name)
				if item == nil {
					return s.RaiseError("%s is not a valid EnumItem", name)
				}
				return s.Push(item)
			},
			"__newindex": func(s rbxmk.State) int {
				name := string(s.Pull(2, "string").(types.String))
				return s.RaiseError("%s cannot be assigned to", name)
			},
		},
		Methods: rbxmk.Methods{
			"GetEnumItems": {
				Func: func(s rbxmk.State, v types.Value) int {
					items := v.(*rtypes.Enum).Items()
					array := make(rtypes.Array, len(items))
					for i, item := range items {
						array[i] = item
					}
					return s.Push(array)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Array{T: dt.Prim("EnumItem")}},
						},
						Summary:     "Libraries/roblox/Types/Enum:Methods/GetEnumItems/Summary",
						Description: "Libraries/roblox/Types/Enum:Methods/GetEnumItems/Description",
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Index: &dump.Function{
						Parameters: dump.Parameters{
							{Name: "name", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("EnumItem")},
						},
						CanError:    true,
						Summary:     "Libraries/roblox/Types/Enum:Operators/Index/Summary",
						Description: "Libraries/roblox/Types/Enum:Operators/Index/Description",
					},
				},
				Summary:     "Libraries/roblox/Types/Enum:Summary",
				Description: "Libraries/roblox/Types/Enum:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Array,
			EnumItem,
			String,
		},
	}
}
