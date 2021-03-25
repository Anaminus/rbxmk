package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(DescActions) }
func DescActions() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "DescActions",
		PushTo: func(s rbxmk.State, v types.Value) (lvs []lua.LValue, err error) {
			actions, ok := v.(rtypes.DescActions)
			if !ok {
				return nil, rbxmk.TypeError{Want: "DescActions", Got: v.Type()}
			}
			actionRfl := s.MustReflector("DescAction")
			table := s.L.CreateTable(len(actions), 0)
			for i, v := range actions {
				lv, err := actionRfl.PushTo(s, v)
				if err != nil {
					return nil, err
				}
				table.RawSetInt(i+1, lv[0])
			}
			return []lua.LValue{table}, nil
		},
		PullFrom: func(s rbxmk.State, lvs ...lua.LValue) (v types.Value, err error) {
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: "table", Got: lvs[0].Type().String()}
			}
			actionRfl := s.MustReflector("DescAction")
			n := table.Len()
			actions := make(rtypes.DescActions, n)
			for i := 1; i <= n; i++ {
				v, err := actionRfl.PullFrom(s, table.RawGetInt(i))
				if err != nil {
					return nil, err
				}
				actions[i-1] = v.(*rtypes.DescAction)
			}
			return actions, nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying:  dt.Array{T: dt.Prim("DescAction")},
				Summary:     "Libraries/rbxmk/Types/DescActions:Summary",
				Description: "Libraries/rbxmk/Types/DescActions:Description",
			}
		},
	}
}
