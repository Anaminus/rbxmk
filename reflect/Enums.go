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
		Name:     rtypes.T_Enums,
		PushTo:   rbxmk.PushPtrTypeTo(rtypes.T_Enums),
		PullFrom: rbxmk.PullTypeFrom(rtypes.T_Enums),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case **rtypes.Enums:
				*p = v.(*rtypes.Enums)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Enums).(*rtypes.Enums)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__index": func(s rbxmk.State) int {
				enums := s.Pull(1, rtypes.T_Enums).(*rtypes.Enums)
				name := string(s.Pull(2, rtypes.T_String).(types.String))
				enum := enums.Enum(name)
				if enum == nil {
					return s.RaiseError("%s is not a valid Enum", name)
				}
				return s.Push(enum)
			},
			"__newindex": func(s rbxmk.State) int {
				name := string(s.Pull(2, rtypes.T_String).(types.String))
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
							{Type: dt.Array{T: dt.Prim(rtypes.T_Enum)}},
						},
						Summary:     "Types/Enums:Methods/GetEnums/Summary",
						Description: "Types/Enums:Methods/GetEnums/Description",
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category: "roblox",
				Operators: &dump.Operators{
					Index: &dump.Function{
						Parameters: dump.Parameters{
							{Name: "name", Type: dt.Prim(rtypes.T_String)},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(rtypes.T_Enum)},
						},
						CanError:    true,
						Summary:     "Types/Enums:Operators/Index/Summary",
						Description: "Types/Enums:Operators/Index/Description",
					},
				},
				Summary:     "Types/Enums:Summary",
				Description: "Types/Enums:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Array,
			Enum,
			String,
		},
	}
}
