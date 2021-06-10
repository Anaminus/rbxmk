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
		Name: "EnumItemDesc",
		PushTo: func(s rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			desc, ok := v.(rtypes.EnumItemDesc)
			if !ok {
				return nil, rbxmk.TypeError{Want: "EnumItemDesc", Got: v.Type()}
			}
			item := rbxdump.EnumItem(desc)
			return s.MustReflector("DescFields").PushTo(s, rtypes.DescFields(item.Fields()))
		},
		PullFrom: func(s rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			fields, err := s.MustReflector("DescFields").PullFrom(s, lv)
			if err != nil {
				return nil, err
			}
			item := rbxdump.EnumItem{}
			item.SetFields(rbxdump.Fields(fields.(rtypes.DescFields)))
			return rtypes.EnumItemDesc(item), nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying: dt.Struct{
					"Name":  dt.Prim("string"),
					"Value": dt.Prim("int"),
					"Index": dt.Optional{T: dt.Prim("int")},
					"Tags":  dt.Array{T: dt.Prim("string")},
				},
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
