package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(Bool) }
func Bool() Reflector {
	return Reflector{
		Name: "bool",
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LBool(v.(types.Bool))}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			if n, ok := lvs[0].(lua.LBool); ok {
				return types.Bool(n), nil
			}
			return nil, TypeError(nil, 0, "bool")
		},
	}
}
