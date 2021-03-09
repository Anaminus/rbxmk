package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(Float) }
func Float() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:  "float",
		Flags: rbxmk.Exprim,
		PushTo: func(s rbxmk.State, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Float))}, nil
		},
		PullFrom: func(s rbxmk.State, lvs ...lua.LValue) (v types.Value, err error) {
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
			return nil, rbxmk.TypeError("float", lvs[0].Type().String())
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
