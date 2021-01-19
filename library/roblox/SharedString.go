package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(SharedString) }
func SharedString() Reflector {
	return Reflector{
		Name:  "SharedString",
		Flags: Exprim,
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(types.SharedString))}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LString:
				return types.SharedString(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("SharedString") {
					if v, ok := v.Value.(types.SharedString); ok {
						return v, nil
					}
				}
			}
			return nil, TypeError(nil, 0, "SharedString")
		},
	}
}
