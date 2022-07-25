package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/types"
)

func init() { register(TypeDesc) }
func TypeDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_TypeDesc,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			typ, ok := v.(rtypes.TypeDesc)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_TypeDesc, Got: v.Type()}
			}
			table := c.CreateTable(0, 2)
			if err := c.PushToDictionary(table, "Category", types.String(typ.Embedded.Category)); err != nil {
				return nil, err
			}
			if err := c.PushToDictionary(table, "Name", types.String(typ.Embedded.Name)); err != nil {
				return nil, err
			}
			return table, nil
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_Table, Got: lv.Type().String()}
			}
			category, err := c.PullFromDictionary(table, "Category", rtypes.T_String)
			if err != nil {
				return nil, err
			}
			name, err := c.PullFromDictionary(table, "Name", rtypes.T_String)
			if err != nil {
				return nil, err
			}
			typ := rtypes.TypeDesc{
				Embedded: rbxdump.Type{
					Category: string(category.(types.String)),
					Name:     string(name.(types.String)),
				},
			}
			return typ, nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.TypeDesc:
				*p = v.(rtypes.TypeDesc)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category: "rbxmk",
				Underlying: dt.Struct{
					"Category": dt.Prim(rtypes.T_String),
					"Name":     dt.Prim(rtypes.T_String),
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
