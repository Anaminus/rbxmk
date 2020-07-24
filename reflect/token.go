package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Token() Type {
	return Type{
		Name: "token",
		ReflectTo: func(s State, t Type, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Token))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v types.Value, err error) {
			if n, ok := lvs[0].(lua.LNumber); ok {
				return types.Token(n), nil
			}
			return nil, TypeError(nil, 0, "token")
		},
	}
}
