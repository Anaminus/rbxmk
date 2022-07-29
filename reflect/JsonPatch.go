package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(JsonPatch) }
func JsonPatch() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_JsonPatch,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			actions, ok := v.(rtypes.JsonPatch)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_JsonPatch, Got: v.Type()}
			}
			opRfl := c.MustReflector(rtypes.T_JsonOperation)
			table := c.CreateTable(len(actions), 0)
			for i, v := range actions {
				lv, err := opRfl.PushTo(c, v)
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
			opRfl := c.MustReflector(rtypes.T_JsonOperation)
			n := table.Len()
			actions := make(rtypes.JsonPatch, n)
			for i := 1; i <= n; i++ {
				v, err := opRfl.PullFrom(c, table.RawGetInt(i))
				if err != nil {
					return nil, err
				}
				actions[i-1] = v.(rtypes.JsonOperation)
			}
			return actions, nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.JsonPatch:
				*p = v.(rtypes.JsonPatch)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category:    "rbxmk",
				Underlying:  dt.P(dt.Array(dt.Prim(rtypes.T_JsonOperation))),
				Summary:     "Types/JsonPatch:Summary",
				Description: "Types/JsonPatch:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			JsonOperation,
		},
	}
}
