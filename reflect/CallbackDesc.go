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

func init() { register(CallbackDesc) }
func CallbackDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_CallbackDesc,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			desc, ok := v.(rtypes.CallbackDesc)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_CallbackDesc, Got: v.Type()}
			}
			member := rbxdump.Callback(desc)
			fields := member.Fields()
			fields["MemberType"] = member.MemberType()
			return c.MustReflector(rtypes.T_DescFields).PushTo(c, rtypes.DescFields(fields))
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			fields, err := c.MustReflector(rtypes.T_DescFields).PullFrom(c, lv)
			if err != nil {
				return nil, err
			}
			member := rbxdump.Callback{}
			member.SetFields(rbxdump.Fields(fields.(rtypes.DescFields)))
			return rtypes.CallbackDesc(member), nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.CallbackDesc:
				*p = v.(rtypes.CallbackDesc)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying: dt.Struct{
					"MemberType": dt.Prim(rtypes.T_String),
					"Name":       dt.Prim(rtypes.T_String),
					"ReturnType": dt.Prim(rtypes.T_TypeDesc),
					"Security":   dt.Prim(rtypes.T_String),
					"Parameters": dt.Array{T: dt.Prim(rtypes.T_ParameterDesc)},
					"Tags":       dt.Array{T: dt.Prim(rtypes.T_String)},
				},
				Summary:     "Types/CallbackDesc:Summary",
				Description: "Types/CallbackDesc:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			DescFields,
		},
	}
}
