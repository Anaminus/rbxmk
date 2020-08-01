package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	lua "github.com/yuin/gopher-lua"
)

func init() { register(Nil) }
func Nil() Reflector {
	return Reflector{
		Name: "nil",
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNil}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			if lvs[0] == lua.LNil {
				return nil, nil
			}
			return nil, TypeError(nil, 0, "nil")
		},
	}
}
