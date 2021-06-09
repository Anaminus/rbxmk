package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Table) }
func Table() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "table",
		PushTo: func(s rbxmk.State, v types.Value) (lv lua.LValue, err error) {
			table, ok := v.(rtypes.Table)
			if !ok {
				return nil, rbxmk.TypeError{Want: "*lua.LTable", Got: v.Type()}
			}
			return table.LTable, nil
		},
		PullFrom: func(s rbxmk.State, lv lua.LValue) (v types.Value, err error) {
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: "table", Got: lv.Type().String()}
			}
			return rtypes.Table{LTable: table}, nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/table:Summary",
				Description: "Types/table:Description",
			}
		},
	}
}
