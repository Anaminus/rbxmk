package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(Int) }
func Int() Reflector {
	return Reflector{
		Name:  "int",
		Flags: rbxmk.Exprim,
		PushTo: func(s State, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Int))}, nil
		},
		PullFrom: func(s State, lvs ...lua.LValue) (v types.Value, err error) {
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
			return nil, rbxmk.TypeError(nil, 0, "int")
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case types.Int:
				return v
			case types.Intlike:
				return types.Int(v.Intlike())
			case types.Numberlike:
				return types.Int(v.Numberlike())
			}
			return nil
		},
	}
}
