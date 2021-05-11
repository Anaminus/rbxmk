package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(TypeDesc) }
func TypeDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "TypeDesc",
		PushTo:   rbxmk.PushTypeTo("TypeDesc"),
		PullFrom: rbxmk.PullTypeFrom("TypeDesc"),
		Metatable: rbxmk.Metatable{
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "TypeDesc").(rtypes.TypeDesc)
				op := s.Pull(2, "TypeDesc").(rtypes.TypeDesc)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"Category": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.TypeDesc)
					return s.Push(types.String(desc.Embedded.Category))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						ReadOnly:    true,
						Summary:     "Types/TypeDesc:Properties/Category/Summary",
						Description: "Types/TypeDesc:Properties/Category/Description",
					}
				},
			},
			"Name": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.TypeDesc)
					return s.Push(types.String(desc.Embedded.Name))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						ReadOnly:    true,
						Summary:     "Types/TypeDesc:Properties/Name/Summary",
						Description: "Types/TypeDesc:Properties/Name/Description",
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Eq: &dump.Cmpop{
						Summary:     "Types/TypeDesc:Operators/Eq/Summary",
						Description: "Types/TypeDesc:Operators/Eq/Description",
					},
				},
				Summary:     "Types/TypeDesc:Summary",
				Description: "Types/TypeDesc:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			String,
		},
	}
}