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
		Name:     rtypes.T_Enum,
		PushTo:   rbxmk.PushPtrTypeTo(rtypes.T_Enum),
		PullFrom: rbxmk.PullTypeFrom(rtypes.T_Enum),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case **rtypes.Enum:
				*p = v.(*rtypes.Enum)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Enum).(*rtypes.Enum)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__index": func(s rbxmk.State) int {
				enum := s.Pull(1, rtypes.T_Enum).(*rtypes.Enum)
				name := string(s.Pull(2, rtypes.T_String).(types.String))
				item := enum.Item(name)
				if item == nil {
					return s.RaiseError("%s is not a valid EnumItem", name)
				}
				return s.Push(item)
			},
			"__newindex": func(s rbxmk.State) int {
				name := string(s.Pull(2, rtypes.T_String).(types.String))
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
							{Type: dt.Array(dt.Prim(rtypes.T_EnumItem))},
						},
						Summary:     "Types/Enum:Methods/GetEnumItems/Summary",
						Description: "Types/Enum:Methods/GetEnumItems/Description",
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
							{Type: dt.Prim(rtypes.T_EnumItem)},
						},
						CanError:    true,
						Summary:     "Types/Enum:Operators/Index/Summary",
						Description: "Types/Enum:Operators/Index/Description",
					},
				},
				Summary:     "Types/Enum:Summary",
				Description: "Types/Enum:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Array,
			EnumItem,
			String,
		},
	}
}
