package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Nil() Type {
	return Type{
		Name: "nil",
		PushTo: func(s State, t Type, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNil}, nil
		},
		PullFrom: func(s State, t Type, lvs ...lua.LValue) (v types.Value, err error) {
			if lvs[0] == lua.LNil {
				return nil, nil
			}
			return nil, TypeError(nil, 0, "nil")
		},
	}
}
