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
		Name: rtypes.T_DescActions,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			actions, ok := v.(rtypes.DescActions)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_DescActions, Got: v.Type()}
			}
			actionRfl := c.MustReflector(rtypes.T_DescAction)
			table := c.CreateTable(len(actions), 0)
			for i, v := range actions {
				lv, err := actionRfl.PushTo(c, v)
				if err != nil {
					return nil, err
				}
				table.RawSetInt(i+1, lv)
			}
			return table, nil
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_Table, Got: lv.Type().String()}
			}
			actionRfl := c.MustReflector(rtypes.T_DescAction)
			n := table.Len()
			actions := make(rtypes.DescActions, n)
			for i := 1; i <= n; i++ {
				v, err := actionRfl.PullFrom(c, table.RawGetInt(i))
				if err != nil {
					return nil, err
				}
				actions[i-1] = v.(*rtypes.DescAction)
			}
			return actions, nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.DescActions:
				*p = v.(rtypes.DescActions)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category:    "rbxmk",
				Underlying:  dt.Array{T: dt.Prim(rtypes.T_DescAction)},
				Summary:     "Types/DescActions:Summary",
				Description: "Types/DescActions:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			DescAction,
		},
	}
}
