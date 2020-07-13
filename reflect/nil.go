package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/yuin/gopher-lua"
)

func Nil() Type {
	return Type{
		Name: "nil",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNil}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if lvs[0] == lua.LNil {
				return nil, nil
			}
			return nil, TypeError(nil, 0, "number")
		},
	}
}
