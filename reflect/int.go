package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/robloxapi/types"
)

func init() { register(Int) }
func Int() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:  "int",
		Flags: rbxmk.Exprim,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			return lua.LNumber(v.(types.Int)), nil
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			switch v := lv.(type) {
			case lua.LNumber:
				return types.Int(v), nil
			case *lua.LUserData:
				if v.Metatable == c.GetTypeMetatable("int") {
					if v, ok := v.Value().(types.Int); ok {
						return v, nil
					}
				}
			}
			return nil, rbxmk.TypeError{Want: "int", Got: lv.Type().String()}
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.Int:
				*p = v.(types.Int)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case types.Int:
				return v
			case types.Intlike:
				return types.Int(v.Intlike())
			case types.Numberlike:
				return types.Int(v.Numberlike())
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/int:Summary",
				Description: "Types/int:Description",
			}
		},
	}
}
