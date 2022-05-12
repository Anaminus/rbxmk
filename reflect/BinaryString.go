package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(BinaryString) }
func BinaryString() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:  rtypes.T_BinaryString,
		Flags: rbxmk.Exprim,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			return lua.LString(v.(types.BinaryString)), nil
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			switch v := lv.(type) {
			case lua.LString:
				return types.BinaryString(v), nil
			case *lua.LUserData:
				if v.Metatable == c.GetTypeMetatable(rtypes.T_BinaryString) {
					if v, ok := v.Value().(types.BinaryString); ok {
						return v, nil
					}
				}
			}
			return nil, rbxmk.TypeError{Want: rtypes.T_BinaryString, Got: lv.Type().String()}
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.BinaryString:
				*p = v.(types.BinaryString)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case types.BinaryString:
				return v
			case types.Stringlike:
				return types.BinaryString(v.Stringlike())
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/BinaryString:Summary",
				Description: "Types/BinaryString:Description",
			}
		},
	}
}
