package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Number) }
func Number() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_Number,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			return lua.LNumber(v.(types.Double)), nil
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			if n, ok := lv.(lua.LNumber); ok {
				return types.Double(n), nil
			}
			return nil, rbxmk.TypeError{Want: rtypes.T_Number, Got: lv.Type().String()}
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
