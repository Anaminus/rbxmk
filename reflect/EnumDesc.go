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

const T_EnumDesc = "EnumDesc"

func init() { register(EnumDesc) }
func EnumDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: T_EnumDesc,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			desc, ok := v.(rtypes.EnumDesc)
			if !ok {
				return nil, rbxmk.TypeError{Want: T_EnumDesc, Got: v.Type()}
			}
			enum := rbxdump.Enum(desc)
			return c.MustReflector(T_DescFields).PushTo(c, rtypes.DescFields(enum.Fields()))
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			fields, err := c.MustReflector(T_DescFields).PullFrom(c, lv)
			if err != nil {
				return nil, err
			}
			enum := rbxdump.Enum{}
			enum.SetFields(rbxdump.Fields(fields.(rtypes.DescFields)))
			return rtypes.EnumDesc(enum), nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.EnumDesc:
				*p = v.(rtypes.EnumDesc)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying: dt.Struct{
					"Name": dt.Prim(T_String),
					"Tags": dt.Array{T: dt.Prim(T_String)},
				},
				Summary:     "Types/EnumDesc:Summary",
				Description: "Types/EnumDesc:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Array,
			String,
		},
	}
}
