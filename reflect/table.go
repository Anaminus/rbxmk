package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Table() Type {
	return Type{
		Name: "table",
		PushTo: func(s State, t Type, v types.Value) (lvs []lua.LValue, err error) {
			table, ok := v.(rtypes.Table)
			if !ok {
				return nil, TypeError(nil, 0, "*lua.LTable")
			}
			return []lua.LValue{table.LTable}, nil
		},
		PullFrom: func(s State, t Type, lvs ...lua.LValue) (v types.Value, err error) {
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, TypeError(nil, 0, "table")
			}
			return rtypes.Table{LTable: table}, nil
		},
	}
}
