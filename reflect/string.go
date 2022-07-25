package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(String) }
func String() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_String,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			return lua.LString(v.(types.String)), nil
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			if n, ok := lv.(lua.LString); ok {
				return types.String(n), nil
			}
			return nil, rbxmk.TypeError{Want: rtypes.T_String, Got: lv.Type().String()}
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.String:
				*p = v.(types.String)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case types.String:
				return v
			case types.Stringlike:
				return types.String(v.Stringlike())
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category:    "primitive",
				Summary:     "Types/string:Summary",
				Description: "Types/string:Description",
			}
		},
	}
}
