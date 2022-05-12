package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(EnumItem) }
func EnumItem() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     rtypes.T_EnumItem,
		PushTo:   rbxmk.PushPtrTypeTo(rtypes.T_EnumItem),
		PullFrom: rbxmk.PullTypeFrom(rtypes.T_EnumItem),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case **rtypes.EnumItem:
				*p = v.(*rtypes.EnumItem)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_EnumItem).(*rtypes.EnumItem)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"Name": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.String(v.(*rtypes.EnumItem).Name()))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_String),
						ReadOnly:    true,
						Summary:     "Types/EnumItem:Properties/Name/Summary",
						Description: "Types/EnumItem:Properties/Name/Description",
					}
				},
			},
			"Value": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Int(v.(*rtypes.EnumItem).Value()))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Int),
						ReadOnly:    true,
						Summary:     "Types/EnumItem:Properties/Value/Summary",
						Description: "Types/EnumItem:Properties/Value/Description",
					}
				},
			},
			"EnumType": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(*rtypes.EnumItem).Enum())
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Enum),
						ReadOnly:    true,
						Summary:     "Types/EnumItem:Properties/EnumType/Summary",
						Description: "Types/EnumItem:Properties/EnumType/Description",
					}
				},
			},
		},
		Methods: rbxmk.Methods{
			"IsA": {
				Func: func(s rbxmk.State, v types.Value) int {
					enumName := string(s.Pull(2, rtypes.T_String).(types.String))
					item := v.(*rtypes.EnumItem)
					return s.Push(types.Bool(item.Enum().Name() == enumName))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "enumName", Type: dt.Prim(rtypes.T_String)},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(rtypes.T_Bool)},
						},
						Summary:     "Types/EnumItem:Methods/IsA/Summary",
						Description: "Types/EnumItem:Methods/IsA/Description",
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/EnumItem:Summary",
				Description: "Types/EnumItem:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Enum,
			Int,
			String,
		},
	}
}
