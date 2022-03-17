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
		Name: "PropertyDesc",
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			desc, ok := v.(rtypes.PropertyDesc)
			if !ok {
				return nil, rbxmk.TypeError{Want: "PropertyDesc", Got: v.Type()}
			}
			member := rbxdump.Property(desc)
			fields := member.Fields()
			fields["MemberType"] = member.MemberType()
			return c.MustReflector("DescFields").PushTo(c, rtypes.DescFields(fields))
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			fields, err := c.MustReflector("DescFields").PullFrom(c, lv)
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
				Underlying: dt.Struct{
					"MemberType":    dt.Prim("string"),
					"Name":          dt.Prim("string"),
					"ValueType":     dt.Prim("TypeDesc"),
					"Category":      dt.Prim("string"),
					"ReadSecurity":  dt.Prim("string"),
					"WriteSecurity": dt.Prim("string"),
					"CanLoad":       dt.Prim("bool"),
					"CanSave":       dt.Prim("bool"),
					"Tags":          dt.Array{T: dt.Prim("string")},
				},
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
