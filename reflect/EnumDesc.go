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

func init() { register(EnumDesc) }
func EnumDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "EnumDesc",
		PushTo: func(s rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			desc, ok := v.(rtypes.EnumDesc)
			if !ok {
				return nil, rbxmk.TypeError{Want: "EnumDesc", Got: v.Type()}
			}
			enum := rbxdump.Enum(desc)
			return s.MustReflector("DescFields").PushTo(s, rtypes.DescFields(enum.Fields()))
		},
		PullFrom: func(s rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			fields, err := s.MustReflector("DescFields").PullFrom(s, lv)
			if err != nil {
				return nil, err
			}
			enum := rbxdump.Enum{}
			enum.SetFields(rbxdump.Fields(fields.(rtypes.DescFields)))
			return rtypes.EnumDesc(enum), nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying: dt.Struct{
					"Name": dt.Prim("string"),
					"Tags": dt.Array{T: dt.Prim("string")},
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
