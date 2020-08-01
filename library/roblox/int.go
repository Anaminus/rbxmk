package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	lua "github.com/yuin/gopher-lua"
)

func init() { register(Int) }
func Int() Reflector {
	return Reflector{
		Name:  "int",
		Flags: Exprim,
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Int))}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LNumber:
				return types.Int(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("int") {
					if v, ok := v.Value.(types.Int); ok {
						return v, nil
					}
				}
			}
			return nil, TypeError(nil, 0, "int")
		},
	}
}
