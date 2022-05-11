package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/robloxapi/types"
)

const T_Token = "token"

func init() { register(Token) }
func Token() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:  T_Token,
		Flags: rbxmk.Exprim,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			return lua.LNumber(v.(types.Token)), nil
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			switch v := lv.(type) {
			case lua.LNumber:
				return types.Token(v), nil
			case *lua.LUserData:
				if v.Metatable == c.GetTypeMetatable(T_Token) {
					if v, ok := v.Value().(types.Token); ok {
						return v, nil
					}
				}
			}
			return nil, rbxmk.TypeError{Want: T_Token, Got: lv.Type().String()}
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.Token:
				*p = v.(types.Token)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/token:Summary",
				Description: "Types/token:Description",
			}
		},
	}
}
