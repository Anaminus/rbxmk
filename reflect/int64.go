package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/robloxapi/types"
)

func init() { register(Int64) }
func Int64() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:  "int64",
		Flags: rbxmk.Exprim,
		PushTo: func(s rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			return lua.LNumber(v.(types.Int64)), nil
		},
		PullFrom: func(s rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			switch v := lv.(type) {
			case lua.LNumber:
				return types.Int64(v), nil
			case *lua.LUserData:
				if v.Metatable == s.GetTypeMetatable("int64") {
					if v, ok := v.Value().(types.Int64); ok {
						return v, nil
					}
				}
			}
			return nil, rbxmk.TypeError{Want: "int64", Got: lv.Type().String()}
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.Int64:
				*p = v.(types.Int64)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case types.Int64:
				return v
			case types.Intlike:
				return types.Int64(v.Intlike())
			case types.Numberlike:
				return types.Int64(v.Numberlike())
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/int64:Summary",
				Description: "Types/int64:Description",
			}
		},
	}
}
