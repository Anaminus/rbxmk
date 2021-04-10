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
		PushTo: func(s rbxmk.State, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNil}, nil
		},
		PullFrom: func(s rbxmk.State, lvs ...lua.LValue) (v types.Value, err error) {
			if lvs[0] == lua.LNil {
				return rtypes.Nil, nil
			}
			return nil, rbxmk.TypeError{Want: "nil", Got: lvs[0].Type().String()}
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/nil:Summary",
				Description: "Types/nil:Description",
			}
		},
	}
}
