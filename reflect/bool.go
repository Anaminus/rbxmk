package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Bool) }
func Bool() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_Bool,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			return lua.LBool(v.(types.Bool)), nil
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			if n, ok := lv.(lua.LBool); ok {
				return types.Bool(n), nil
			}
			return nil, rbxmk.TypeError{Want: rtypes.T_Bool, Got: lv.Type().String()}
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.Bool:
				*p = v.(types.Bool)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category:    "primitive",
				Underlying:  dt.P(dt.Prim(rtypes.T_Bool)),
				Summary:     "Types/bool:Summary",
				Description: "Types/bool:Description",
			}
		},
	}
}
