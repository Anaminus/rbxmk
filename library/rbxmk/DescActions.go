package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(DescActions) }
func DescActions() Reflector {
	return Reflector{
		Name: "DescActions",
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			actions, ok := v.(rtypes.DescActions)
			if !ok {
				return nil, rbxmk.TypeError(nil, 0, "DescActions")
			}
			actionRfl := s.Reflector("DescAction")
			table := s.L.CreateTable(len(actions), 0)
			for i, v := range actions {
				lv, err := actionRfl.PushTo(s, actionRfl, v)
				if err != nil {
					return nil, err
				}
				table.RawSetInt(i+1, lv[0])
			}
			return []lua.LValue{table}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError(nil, 0, "table")
			}
			actionRfl := s.Reflector("DescAction")
			n := table.Len()
			actions := make(rtypes.DescActions, n)
			for i := 1; i <= n; i++ {
				v, err := actionRfl.PullFrom(s, actionRfl, table.RawGetInt(i))
				if err != nil {
					return nil, err
				}
				actions[i-1] = v.(*rtypes.DescAction)
			}
			return actions, nil
		},
	}
}
