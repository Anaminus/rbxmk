package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Bool() Type {
	return Type{
		Name: "bool",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LBool(v.(types.Bool))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LBool); ok {
				return types.Bool(n), nil
			}
			return nil, TypeError(nil, 0, "bool")
		},
	}
}
