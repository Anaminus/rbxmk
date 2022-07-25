package reflect

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Array) }
func Array() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: rtypes.T_Array,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			if c.CycleGuard() {
				defer c.CycleClear()
			}
			array, ok := v.(rtypes.Array)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_Array, Got: v.Type()}
			}
			if c.CycleMark(&array) {
				return nil, fmt.Errorf("arrays cannot be cyclic")
			}
			variantRfl := c.MustReflector(rtypes.T_Variant)
			table := c.CreateTable(len(array), 0)
			for i, v := range array {
				lv, err := variantRfl.PushTo(c, v)
				if err != nil {
					return nil, err
				}
				table.RawSetInt(i+1, lv)
			}
			return table, nil
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			if c.CycleGuard() {
				defer c.CycleClear()
			}
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_Table, Got: lv.Type().String()}
			}
			if c.CycleMark(table) {
				return nil, fmt.Errorf("tables cannot be cyclic")
			}
			variantRfl := c.MustReflector(rtypes.T_Variant)
			n := table.Len()
			array := make(rtypes.Array, n)
			for i := 1; i <= n; i++ {
				if array[i-1], err = variantRfl.PullFrom(c, table.RawGetInt(i)); err != nil {
					return nil, err
				}
			}
			return array, nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.Array:
				*p = v.(rtypes.Array)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category:    "group",
				Summary:     "Types/Array:Summary",
				Description: "Types/Array:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Variant,
		},
	}
}
