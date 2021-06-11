package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/robloxapi/types"
)

func init() { register(SharedString) }
func SharedString() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:  "SharedString",
		Flags: rbxmk.Exprim,
		PushTo: func(s rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			return lua.LString(v.(types.SharedString)), nil
		},
		PullFrom: func(s rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			switch v := lv.(type) {
			case lua.LString:
				return types.SharedString(v), nil
			case *lua.LUserData:
				if v.Metatable == s.GetTypeMetatable("SharedString") {
					if v, ok := v.Value().(types.SharedString); ok {
						return v, nil
					}
				}
			}
			return nil, rbxmk.TypeError{Want: "SharedString", Got: lv.Type().String()}
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case types.SharedString:
				return v
			case types.Stringlike:
				return types.SharedString(v.Stringlike())
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/SharedString:Summary",
				Description: "Types/SharedString:Description",
			}
		},
	}
}
