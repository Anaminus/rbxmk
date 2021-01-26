package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(Float) }
func Float() Reflector {
	return Reflector{
		Name:  "float",
		Flags: Exprim,
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Float))}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LNumber:
				return types.Float(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("float") {
					if v, ok := v.Value.(types.Float); ok {
						return v, nil
					}
				}
			}
			return nil, TypeError(nil, 0, "float")
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case types.Float:
				return v
			case types.Numberlike:
				return types.Float(v.Numberlike())
			case types.Intlike:
				return types.Float(v.Intlike())
			}
			return nil
		},
	}
}
