package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(Double) }
func Double() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "double",
		PushTo: func(s rbxmk.State, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Double))}, nil
		},
		PullFrom: func(s rbxmk.State, lvs ...lua.LValue) (v types.Value, err error) {
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
			return nil, rbxmk.TypeError("double", lvs[0].Type().String())
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case types.Double:
				return v
			case types.Numberlike:
				return types.Double(v.Numberlike())
			case types.Intlike:
				return types.Double(v.Intlike())
			}
			return nil
		},
	}
}
