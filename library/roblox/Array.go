package reflect

import (
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Array) }
func Array() Reflector {
	return Reflector{
		Name: "Array",
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			if s.Cycle == nil {
				s.Cycle = &Cycle{}
				defer func() { s.Cycle = nil }()
			}
			array, ok := v.(rtypes.Array)
			if !ok {
				return nil, TypeError(nil, 0, "Array")
			}
			if s.Cycle.Mark(&array) {
				return nil, fmt.Errorf("arrays cannot be cyclic")
			}
			variantRfl := s.Reflector("Variant")
			table := s.L.CreateTable(len(array), 0)
			for i, v := range array {
				lv, err := variantRfl.PushTo(s, variantRfl, v)
				if err != nil {
					return nil, err
				}
				table.RawSetInt(i+1, lv[0])
			}
			return []lua.LValue{table}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			if s.Cycle == nil {
				s.Cycle = &Cycle{}
				defer func() { s.Cycle = nil }()
			}
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, TypeError(nil, 0, "table")
			}
			if s.Cycle.Mark(table) {
				return nil, fmt.Errorf("tables cannot be cyclic")
			}
			variantRfl := s.Reflector("Variant")
			n := table.Len()
			array := make(rtypes.Array, n)
			for i := 1; i <= n; i++ {
				if array[i-1], err = variantRfl.PullFrom(s, variantRfl, table.RawGetInt(i)); err != nil {
					return nil, err
				}
			}
			return array, nil
		},
	}
}
