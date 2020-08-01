package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func init() { register(String) }
func String() Reflector {
	return Reflector{
		Name: "string",
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(types.String))}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			if n, ok := lvs[0].(lua.LString); ok {
				return types.String(n), nil
			}
			return nil, TypeError(nil, 0, "string")
		},
	}
}
