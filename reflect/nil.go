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
		Name: "nil",
		PushTo: func(s rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			return lua.LNil, nil
		},
		PullFrom: func(s rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			if lv == lua.LNil {
				return rtypes.Nil, nil
			}
			return nil, rbxmk.TypeError{Want: "nil", Got: lv.Type().String()}
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/nil:Summary",
				Description: "Types/nil:Description",
			}
		},
	}
}
