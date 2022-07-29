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

func init() { register(ClassDesc) }
func ClassDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_ClassDesc,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			desc, ok := v.(rtypes.ClassDesc)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_ClassDesc, Got: v.Type()}
			}
			class := rbxdump.Class(desc)
			return c.MustReflector(rtypes.T_DescFields).PushTo(c, rtypes.DescFields(class.Fields()))
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			fields, err := c.MustReflector(rtypes.T_DescFields).PullFrom(c, lv)
			if err != nil {
				return nil, err
			}
			class := rbxdump.Class{}
			class.SetFields(rbxdump.Fields(fields.(rtypes.DescFields)))
			return rtypes.ClassDesc(class), nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.ClassDesc:
				*p = v.(rtypes.ClassDesc)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category: "rbxmk",
				Underlying: dt.P(dt.Struct(dt.KindStruct{
					"Name":           dt.Prim(rtypes.T_String),
					"Superclass":     dt.Prim(rtypes.T_String),
					"MemoryCategory": dt.Prim(rtypes.T_String),
					"Tags":           dt.Array(dt.Prim(rtypes.T_String)),
				})),
				Summary:     "Types/ClassDesc:Summary",
				Description: "Types/ClassDesc:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Array,
			String,
		},
	}
}
