package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/robloxapi/types"
)

func init() { register(Double) }
func Double() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "double",
		PushTo: func(s rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			return lua.LNumber(v.(types.Double)), nil
		},
		PullFrom: func(s rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			switch v := lv.(type) {
			case lua.LNumber:
				return types.Double(v), nil
			case *lua.LUserData:
				if v.Metatable == s.GetTypeMetatable("double") {
					if v, ok := v.Value().(types.Double); ok {
						return v, nil
					}
				}
			}
			return nil, rbxmk.TypeError{Want: "double", Got: lv.Type().String()}
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.Double:
				*p = v.(types.Double)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case types.Double:
				return v
			case types.Numberlike:
				return types.Double(v.Numberlike())
			case types.Intlike:
				return types.Double(v.Intlike())
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/double:Summary",
				Description: "Types/double:Description",
			}
		},
	}
}
