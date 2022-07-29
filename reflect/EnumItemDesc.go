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

func init() { register(EnumItemDesc) }
func EnumItemDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_EnumItemDesc,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			desc, ok := v.(rtypes.EnumItemDesc)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_EnumItemDesc, Got: v.Type()}
			}
			item := rbxdump.EnumItem(desc)
			return c.MustReflector(rtypes.T_DescFields).PushTo(c, rtypes.DescFields(item.Fields()))
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			fields, err := c.MustReflector(rtypes.T_DescFields).PullFrom(c, lv)
			if err != nil {
				return nil, err
			}
			item := rbxdump.EnumItem{}
			item.SetFields(rbxdump.Fields(fields.(rtypes.DescFields)))
			return rtypes.EnumItemDesc(item), nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.EnumItemDesc:
				*p = v.(rtypes.EnumItemDesc)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category: "rbxmk",
				Underlying: dt.P(dt.Struct(dt.KindStruct{
					"Name":  dt.Prim(rtypes.T_String),
					"Value": dt.Prim(rtypes.T_Int),
					"Index": dt.Optional(dt.Prim(rtypes.T_Int)),
					"Tags":  dt.Array(dt.Prim(rtypes.T_String)),
				})),
				Summary:     "Types/EnumItemDesc:Summary",
				Description: "Types/EnumItemDesc:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Array,
			Int,
			String,
		},
	}
}
