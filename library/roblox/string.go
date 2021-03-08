package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(String) }
func String() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "string",
		PushTo: func(s rbxmk.State, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(types.String))}, nil
		},
		PullFrom: func(s rbxmk.State, lvs ...lua.LValue) (v types.Value, err error) {
			if n, ok := lvs[0].(lua.LString); ok {
				return types.String(n), nil
			}
			return nil, rbxmk.TypeError(nil, 0, "string")
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case types.String:
				return v
			case types.Stringlike:
				return types.String(v.Stringlike())
			}
			return nil
		},
	}
}
