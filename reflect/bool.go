package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/yuin/gopher-lua"
)

func Bool() Type {
	return Type{
		Name: "bool",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LBool(v.(bool))}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if n, ok := lvs[0].(lua.LBool); ok {
				return bool(n), nil
			}
			return nil, TypeError(nil, 0, "bool")
		},
	}
}
