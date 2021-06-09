package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/robloxapi/types"
)

func init() { register(Float) }
func Float() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:  "float",
		Flags: rbxmk.Exprim,
		PushTo: func(s rbxmk.State, v types.Value) (lv lua.LValue, err error) {
			return lua.LNumber(v.(types.Float)), nil
		},
		PullFrom: func(s rbxmk.State, lv lua.LValue) (v types.Value, err error) {
			switch v := lv.(type) {
			case lua.LNumber:
				return types.Float(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("float") {
					if v, ok := v.Value().(types.Float); ok {
						return v, nil
					}
				}
			}
			return nil, rbxmk.TypeError{Want: "float", Got: lv.Type().String()}
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
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/float:Summary",
				Description: "Types/float:Description",
			}
		},
	}
}
