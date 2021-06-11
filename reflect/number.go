package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/robloxapi/types"
)

func init() { register(Number) }
func Number() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "number",
		PushTo: func(s rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			return lua.LNumber(v.(types.Double)), nil
		},
		PullFrom: func(s rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			if n, ok := lv.(lua.LNumber); ok {
				return types.Double(n), nil
			}
			return nil, rbxmk.TypeError{Want: "number", Got: lv.Type().String()}
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
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/number:Summary",
				Description: "Types/number:Description",
			}
		},
	}
}
