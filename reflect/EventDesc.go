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

const T_EventDesc = "EventDesc"

func init() { register(EventDesc) }
func EventDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: T_EventDesc,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			desc, ok := v.(rtypes.EventDesc)
			if !ok {
				return nil, rbxmk.TypeError{Want: T_EventDesc, Got: v.Type()}
			}
			member := rbxdump.Event(desc)
			fields := member.Fields()
			fields["MemberType"] = member.MemberType()
			return c.MustReflector(T_DescFields).PushTo(c, rtypes.DescFields(fields))
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			fields, err := c.MustReflector(T_DescFields).PullFrom(c, lv)
			if err != nil {
				return nil, err
			}
			member := rbxdump.Event{}
			member.SetFields(rbxdump.Fields(fields.(rtypes.DescFields)))
			return rtypes.EventDesc(member), nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.EventDesc:
				*p = v.(rtypes.EventDesc)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying: dt.Struct{
					"MemberType": dt.Prim(T_String),
					"Name":       dt.Prim(T_String),
					"Security":   dt.Prim(T_String),
					"Parameters": dt.Array{T: dt.Prim(T_ParameterDesc)},
					"Tags":       dt.Array{T: dt.Prim(T_String)},
				},
				Summary:     "Types/EventDesc:Summary",
				Description: "Types/EventDesc:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Array,
			ParameterDesc,
			String,
		},
	}
}
