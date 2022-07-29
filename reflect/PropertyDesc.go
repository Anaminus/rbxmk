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

func init() { register(PropertyDesc) }
func PropertyDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_PropertyDesc,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			desc, ok := v.(rtypes.PropertyDesc)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_PropertyDesc, Got: v.Type()}
			}
			member := rbxdump.Property(desc)
			fields := member.Fields()
			fields["MemberType"] = member.MemberType()
			return c.MustReflector(rtypes.T_DescFields).PushTo(c, rtypes.DescFields(fields))
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			fields, err := c.MustReflector(rtypes.T_DescFields).PullFrom(c, lv)
			if err != nil {
				return nil, err
			}
			member := rbxdump.Property{}
			member.SetFields(rbxdump.Fields(fields.(rtypes.DescFields)))
			return rtypes.PropertyDesc(member), nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.PropertyDesc:
				*p = v.(rtypes.PropertyDesc)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category: "rbxmk",
				Underlying: dt.P(dt.Struct(dt.KindStruct{
					"MemberType":    dt.Prim(rtypes.T_String),
					"Name":          dt.Prim(rtypes.T_String),
					"ValueType":     dt.Prim(rtypes.T_TypeDesc),
					"Category":      dt.Prim(rtypes.T_String),
					"ReadSecurity":  dt.Prim(rtypes.T_String),
					"WriteSecurity": dt.Prim(rtypes.T_String),
					"CanLoad":       dt.Prim(rtypes.T_Bool),
					"CanSave":       dt.Prim(rtypes.T_Bool),
					"Tags":          dt.Array(dt.Prim(rtypes.T_String)),
				})),
				Summary:     "Types/PropertyDesc:Summary",
				Description: "Types/PropertyDesc:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Array,
			Bool,
			String,
			TypeDesc,
		},
	}
}
