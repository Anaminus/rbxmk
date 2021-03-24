package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/robloxapi/types"
)

func init() { register(Int64) }
func Int64() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:  "int64",
		Flags: rbxmk.Exprim,
		PushTo: func(s rbxmk.State, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Int64))}, nil
		},
		PullFrom: func(s rbxmk.State, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LNumber:
				return types.Int64(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("int64") {
					if v, ok := v.Value().(types.Int64); ok {
						return v, nil
					}
				}
			}
			return nil, rbxmk.TypeError{Want: "int64", Got: lvs[0].Type().String()}
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
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "$TODO",
				Description: "$TODO",
			}
		},
	}
}
