package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/robloxapi/types"
)

func init() { register(Token) }
func Token() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:  "token",
		Flags: rbxmk.Exprim,
		PushTo: func(s rbxmk.State, v types.Value) (lv lua.LValue, err error) {
			return lua.LNumber(v.(types.Token)), nil
		},
		PullFrom: func(s rbxmk.State, lv lua.LValue) (v types.Value, err error) {
			switch v := lv.(type) {
			case lua.LNumber:
				return types.Token(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("token") {
					if v, ok := v.Value().(types.Token); ok {
						return v, nil
					}
				}
			}
			return nil, rbxmk.TypeError{Want: "token", Got: lv.Type().String()}
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/token:Summary",
				Description: "Types/token:Description",
			}
		},
	}
}
