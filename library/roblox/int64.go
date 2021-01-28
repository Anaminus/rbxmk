package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(Int64) }
func Int64() Reflector {
	return Reflector{
		Name:  "int64",
		Flags: rbxmk.Exprim,
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Int64))}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LNumber:
				return types.Int64(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("int64") {
					if v, ok := v.Value.(types.Int64); ok {
						return v, nil
					}
				}
			}
			return nil, rbxmk.TypeError(nil, 0, "int64")
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case types.Int64:
				return v
			case types.Intlike:
				return types.Int64(v.Intlike())
			case types.Numberlike:
				return types.Int64(v.Numberlike())
			}
			return nil
		},
	}
}
