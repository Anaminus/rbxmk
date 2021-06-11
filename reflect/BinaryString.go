package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/robloxapi/types"
)

func init() { register(BinaryString) }
func BinaryString() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:  "BinaryString",
		Flags: rbxmk.Exprim,
		PushTo: func(s rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			return lua.LString(v.(types.BinaryString)), nil
		},
		PullFrom: func(s rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			switch v := lv.(type) {
			case lua.LString:
				return types.BinaryString(v), nil
			case *lua.LUserData:
				if v.Metatable == s.GetTypeMetatable("BinaryString") {
					if v, ok := v.Value().(types.BinaryString); ok {
						return v, nil
					}
				}
			}
			return nil, rbxmk.TypeError{Want: "BinaryString", Got: lv.Type().String()}
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
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/BinaryString:Summary",
				Description: "Types/BinaryString:Description",
			}
		},
	}
}
