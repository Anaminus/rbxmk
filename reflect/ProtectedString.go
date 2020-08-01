package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func init() { register(ProtectedString) }
func ProtectedString() Reflector {
	return Reflector{
		Name:  "ProtectedString",
		Flags: Exprim,
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(types.ProtectedString))}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LString:
				return types.ProtectedString(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("ProtectedString") {
					if v, ok := v.Value.(types.ProtectedString); ok {
						return v, nil
					}
				}
			}
			return nil, TypeError(nil, 0, "ProtectedString")
		},
	}
}
