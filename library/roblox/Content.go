package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(Content) }
func Content() Reflector {
	return Reflector{
		Name:  "Content",
		Flags: rbxmk.Exprim,
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(types.Content))}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LString:
				return types.Content(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("Content") {
					if v, ok := v.Value.(types.Content); ok {
						return v, nil
					}
				}
			}
			return nil, rbxmk.TypeError(nil, 0, "Content")
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case types.Content:
				return v
			case types.Stringlike:
				return types.Content(v.Stringlike())
			}
			return nil
		},
	}
}
