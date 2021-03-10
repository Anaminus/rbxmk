package reflect

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Array) }
func Array() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "Array",
		PushTo: func(s rbxmk.State, v types.Value) (lvs []lua.LValue, err error) {
			if s.CycleGuard() {
				defer s.CycleClear()
			}
			array, ok := v.(rtypes.Array)
			if !ok {
				return nil, rbxmk.TypeError{Want: "Array", Got: v.Type()}
			}
			if s.CycleMark(&array) {
				return nil, fmt.Errorf("arrays cannot be cyclic")
			}
			variantRfl := s.MustReflector("Variant")
			table := s.L.CreateTable(len(array), 0)
			for i, v := range array {
				lv, err := variantRfl.PushTo(s, v)
				if err != nil {
					return nil, err
				}
				table.RawSetInt(i+1, lv[0])
			}
			return []lua.LValue{table}, nil
		},
		PullFrom: func(s rbxmk.State, lvs ...lua.LValue) (v types.Value, err error) {
			if s.CycleGuard() {
				defer s.CycleClear()
			}
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: "table", Got: lvs[0].Type().String()}
			}
			if s.CycleMark(table) {
				return nil, fmt.Errorf("tables cannot be cyclic")
			}
			variantRfl := s.MustReflector("Variant")
			n := table.Len()
			array := make(rtypes.Array, n)
			for i := 1; i <= n; i++ {
				if array[i-1], err = variantRfl.PullFrom(s, table.RawGetInt(i)); err != nil {
					return nil, err
				}
			}
			return array, nil
		},
	}
}
