package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(Double) }
func Double() Reflector {
	return Reflector{
		Name: "double",
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Double))}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LNumber:
				return types.Double(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("double") {
					if v, ok := v.Value.(types.Double); ok {
						return v, nil
					}
				}
			}
			return nil, TypeError(nil, 0, "double")
		},
	}
}
