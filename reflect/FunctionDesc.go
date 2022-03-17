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

func init() { register(FunctionDesc) }
func FunctionDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "FunctionDesc",
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			desc, ok := v.(rtypes.FunctionDesc)
			if !ok {
				return nil, rbxmk.TypeError{Want: "FunctionDesc", Got: v.Type()}
			}
			member := rbxdump.Function(desc)
			fields := member.Fields()
			fields["MemberType"] = member.MemberType()
			return c.MustReflector("DescFields").PushTo(c, rtypes.DescFields(fields))
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			fields, err := c.MustReflector("DescFields").PullFrom(c, lv)
			if err != nil {
				return nil, err
			}
			member := rbxdump.Function{}
			member.SetFields(rbxdump.Fields(fields.(rtypes.DescFields)))
			return rtypes.FunctionDesc(member), nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.FunctionDesc:
				*p = v.(rtypes.FunctionDesc)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying: dt.Struct{
					"MemberType": dt.Prim("string"),
					"Name":       dt.Prim("string"),
					"ReturnType": dt.Prim("TypeDesc"),
					"Security":   dt.Prim("string"),
					"Parameters": dt.Array{T: dt.Prim("ParameterDesc")},
					"Tags":       dt.Array{T: dt.Prim("string")},
				},
				Summary:     "Types/FunctionDesc:Summary",
				Description: "Types/FunctionDesc:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Array,
			ParameterDesc,
			String,
			TypeDesc,
		},
	}
}
