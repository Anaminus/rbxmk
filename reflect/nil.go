package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Nil) }
func Nil() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_Nil,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			return lua.LNil, nil
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			if lv == lua.LNil {
				return rtypes.Nil, nil
			}
			return nil, rbxmk.TypeError{Want: rtypes.T_Nil, Got: lv.Type().String()}
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.NilType:
				*p = v.(rtypes.NilType)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category:    "primitive",
				Summary:     "Types/nil:Summary",
				Description: "Types/nil:Description",
			}
		},
	}
}
