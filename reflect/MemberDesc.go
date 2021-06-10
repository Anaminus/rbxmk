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
		Name: "MemberDesc",
		PushTo: func(s rbxmk.State, v types.Value) (lv lua.LValue, err error) {
			member, ok := v.(rtypes.MemberDesc)
			if !ok {
				return nil, rbxmk.TypeError{Want: "MemberDesc", Got: v.Type()}
			}
			if member.Member == nil {
				panic("member of MemberDesc is nil")
			}
			switch member := member.Member.(type) {
			case *rbxdump.Property:
				return s.MustReflector("PropertyDesc").PushTo(s, rtypes.PropertyDesc(*member))
			case *rbxdump.Function:
				return s.MustReflector("FunctionDesc").PushTo(s, rtypes.FunctionDesc(*member))
			case *rbxdump.Event:
				return s.MustReflector("EventDesc").PushTo(s, rtypes.EventDesc(*member))
			case *rbxdump.Callback:
				return s.MustReflector("CallbackDesc").PushTo(s, rtypes.CallbackDesc(*member))
			default:
				panic("MemberDesc has unknown member type")
			}
		},
		PullFrom: func(s rbxmk.State, lv lua.LValue) (v types.Value, err error) {
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: "table", Got: lv.Type().String()}
			}
			switch typ := s.PullFromDictionary(table, "MemberType", "string").(types.String); typ {
			case "Property":
				value, err := s.MustReflector("PropertyDesc").PullFrom(s, table)
				if err != nil {
					return nil, err
				}
				desc := rbxdump.Property(value.(rtypes.PropertyDesc))
				return rtypes.MemberDesc{Member: &desc}, nil
			case "Function":
				value, err := s.MustReflector("FunctionDesc").PullFrom(s, table)
				if err != nil {
					return nil, err
				}
				desc := rbxdump.Function(value.(rtypes.FunctionDesc))
				return rtypes.MemberDesc{Member: &desc}, nil
			case "Event":
				value, err := s.MustReflector("EventDesc").PullFrom(s, table)
				if err != nil {
					return nil, err
				}
				desc := rbxdump.Event(value.(rtypes.EventDesc))
				return rtypes.MemberDesc{Member: &desc}, nil
			case "Callback":
				value, err := s.MustReflector("CallbackDesc").PullFrom(s, table)
				if err != nil {
					return nil, err
				}
				desc := rbxdump.Callback(value.(rtypes.CallbackDesc))
				return rtypes.MemberDesc{Member: &desc}, nil
			default:
				return nil, fmt.Errorf("unexpected value %q for field MemberType", typ)
			}
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying: dt.Or{
					dt.Prim("PropertyDesc"),
					dt.Prim("FunctionDesc"),
					dt.Prim("EventDesc"),
					dt.Prim("CallbackDesc"),
				},
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
