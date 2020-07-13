package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/yuin/gopher-lua"
)

func Table() Type {
	return Type{
		Name: "table",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			table, ok := v.(*lua.LTable)
			if !ok {
				return nil, TypeError(nil, 0, "*lua.LTable")
			}
			return []lua.LValue{table}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, TypeError(nil, 0, "table")
			}
			return table, nil
		},
	}
}
