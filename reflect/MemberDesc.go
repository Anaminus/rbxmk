package reflect

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/types"
)

func init() { register(MemberDesc) }
func MemberDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_MemberDesc,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			member, ok := v.(rtypes.MemberDesc)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_MemberDesc, Got: v.Type()}
			}
			if member.Member == nil {
				panic("member of MemberDesc is nil")
			}
			switch member := member.Member.(type) {
			case *rbxdump.Property:
				return c.MustReflector(rtypes.T_PropertyDesc).PushTo(c, rtypes.PropertyDesc(*member))
			case *rbxdump.Function:
				return c.MustReflector(rtypes.T_FunctionDesc).PushTo(c, rtypes.FunctionDesc(*member))
			case *rbxdump.Event:
				return c.MustReflector(rtypes.T_EventDesc).PushTo(c, rtypes.EventDesc(*member))
			case *rbxdump.Callback:
				return c.MustReflector(rtypes.T_CallbackDesc).PushTo(c, rtypes.CallbackDesc(*member))
			default:
				panic("MemberDesc has unknown member type")
			}
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_Table, Got: lv.Type().String()}
			}
			typ, err := c.PullFromDictionary(table, "MemberType", rtypes.T_String)
			if err != nil {
				return nil, err
			}
			switch typ.(types.String) {
			case "Property":
				value, err := c.MustReflector(rtypes.T_PropertyDesc).PullFrom(c, table)
				if err != nil {
					return nil, err
				}
				desc := rbxdump.Property(value.(rtypes.PropertyDesc))
				return rtypes.MemberDesc{Member: &desc}, nil
			case "Function":
				value, err := c.MustReflector(rtypes.T_FunctionDesc).PullFrom(c, table)
				if err != nil {
					return nil, err
				}
				desc := rbxdump.Function(value.(rtypes.FunctionDesc))
				return rtypes.MemberDesc{Member: &desc}, nil
			case "Event":
				value, err := c.MustReflector(rtypes.T_EventDesc).PullFrom(c, table)
				if err != nil {
					return nil, err
				}
				desc := rbxdump.Event(value.(rtypes.EventDesc))
				return rtypes.MemberDesc{Member: &desc}, nil
			case "Callback":
				value, err := c.MustReflector(rtypes.T_CallbackDesc).PullFrom(c, table)
				if err != nil {
					return nil, err
				}
				desc := rbxdump.Callback(value.(rtypes.CallbackDesc))
				return rtypes.MemberDesc{Member: &desc}, nil
			default:
				return nil, fmt.Errorf("unexpected value %q for field MemberType", typ)
			}
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.MemberDesc:
				*p = v.(rtypes.MemberDesc)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category: "rbxmk",
				Underlying: dt.P(dt.Or(
					dt.Prim(rtypes.T_PropertyDesc),
					dt.Prim(rtypes.T_FunctionDesc),
					dt.Prim(rtypes.T_EventDesc),
					dt.Prim(rtypes.T_CallbackDesc),
				)),
				Summary:     "Types/MemberDesc:Summary",
				Description: "Types/MemberDesc:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			PropertyDesc,
			FunctionDesc,
			EventDesc,
			CallbackDesc,
		},
	}
}
