package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(BinaryString) }
func BinaryString() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:  "BinaryString",
		Flags: rbxmk.Exprim,
		PushTo: func(s rbxmk.State, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(types.BinaryString))}, nil
		},
		PullFrom: func(s rbxmk.State, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LString:
				return types.BinaryString(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("BinaryString") {
					if v, ok := v.Value().(types.BinaryString); ok {
						return v, nil
					}
				}
			}
			return nil, rbxmk.TypeError{Want: "BinaryString", Got: lvs[0].Type().String()}
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case types.BinaryString:
				return v
			case types.Stringlike:
				return types.BinaryString(v.Stringlike())
			}
			return nil
		},
	}
}
