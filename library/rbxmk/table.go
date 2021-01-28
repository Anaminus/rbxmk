package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Table) }
func Table() Reflector {
	return Reflector{
		Name: "table",
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			table, ok := v.(rtypes.Table)
			if !ok {
				return nil, rbxmk.TypeError(nil, 0, "*lua.LTable")
			}
			return []lua.LValue{table.LTable}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError(nil, 0, "table")
			}
			return rtypes.Table{LTable: table}, nil
		},
	}
}
