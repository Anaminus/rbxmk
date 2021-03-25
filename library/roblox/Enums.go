package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Enums) }
func Enums() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "Enums",
		PushTo:   rbxmk.PushPtrTypeTo("Enums"),
		PullFrom: rbxmk.PullTypeFrom("Enums"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "Enums").(*rtypes.Enums)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__index": func(s rbxmk.State) int {
				enums := s.Pull(1, "Enums").(*rtypes.Enums)
				name := string(s.Pull(2, "string").(types.String))
				enum := enums.Enum(name)
				if enum == nil {
					return s.RaiseError("%s is not a valid Enum", name)
				}
				return s.Push(enum)
			},
			"__newindex": func(s rbxmk.State) int {
				name := string(s.Pull(2, "string").(types.String))
				return s.RaiseError("%s cannot be assigned to", name)
			},
		},
		Methods: rbxmk.Methods{
			"GetEnums": {
				Func: func(s rbxmk.State, v types.Value) int {
					enums := v.(*rtypes.Enums).Enums()
					array := make(rtypes.Array, len(enums))
					for i, enum := range enums {
						array[i] = enum
					}
					return s.Push(array)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Array{T: dt.Prim("Enum")}},
						},
						Summary:     "libraries/roblox/types/Enums:Methods/GetEnums/Summary",
						Description: "libraries/roblox/types/Enums:Methods/GetEnums/Description",
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
							{Type: dt.Prim("Enum")},
						},
						CanError:    true,
						Summary:     "libraries/roblox/types/Enums:Operators/Index/Summary",
						Description: "libraries/roblox/types/Enums:Operators/Index/Description",
					},
				},
				Summary:     "libraries/roblox/types/Enums:Summary",
				Description: "libraries/roblox/types/Enums:Description",
			}
		},
	}
}
