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
func Enums() Reflector {
	return Reflector{
		Name:     "Enums",
		PushTo:   rbxmk.PushTypeTo("Enums"),
		PullFrom: rbxmk.PullTypeFrom("Enums"),
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "Enums").(*rtypes.Enums)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "Enums").(*rtypes.Enums)
				op := s.Pull(2, "Enums").(*rtypes.Enums)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
			"__index": func(s State) int {
				enums := s.Pull(1, "Enums").(*rtypes.Enums)
				name := string(s.Pull(2, "string").(types.String))
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
		Members: rbxmk.Members{
			"GetEnums": {Method: true,
				Get: func(s State, v types.Value) int {
					enums := v.(*rtypes.Enums).Enums()
					array := make(rtypes.Array, len(enums))
					for i, enum := range enums {
						array[i] = enum
					}
					return s.Push(array)
				},
				Dump: func() dump.Value {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Array{T: dt.Prim("Enum")}},
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
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
