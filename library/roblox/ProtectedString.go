package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(ProtectedString) }
func ProtectedString() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:  "ProtectedString",
		Flags: rbxmk.Exprim,
		PushTo: func(s rbxmk.State, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(types.ProtectedString))}, nil
		},
		PullFrom: func(s rbxmk.State, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LString:
				return types.ProtectedString(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("ProtectedString") {
					if v, ok := v.Value().(types.ProtectedString); ok {
						return v, nil
					}
				}
			}
			return nil, rbxmk.TypeError{Want: "ProtectedString", Got: lvs[0].Type().String()}
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case types.ProtectedString:
				return v
			case types.Stringlike:
				return types.ProtectedString(v.Stringlike())
			}
			return nil
		},
	}
}
