package reflect

import (
	"fmt"

	. "github.com/anaminus/rbxmk"
	"github.com/yuin/gopher-lua"
)

func Array() Type {
	return Type{
		Name: "Array",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			if s.Cycle == nil {
				s.Cycle = &Cycle{}
				defer func() { s.Cycle = nil }()
			}
			array, ok := v.([]Value)
			if !ok {
				return nil, TypeError(nil, 0, "[]Value")
			}
			if s.Cycle.Has(&array) {
				return nil, fmt.Errorf("arrays cannot be cyclic")
			}
			s.Cycle.Put(&array)
			variantType := s.Type("Variant")
			table := s.L.CreateTable(len(array), 0)
			for i, v := range array {
				lv, err := variantType.ReflectTo(s, variantType, v)
				if err != nil {
					return nil, err
				}
				table.RawSetInt(i+1, lv[0])
			}
			return []lua.LValue{table}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			if s.Cycle == nil {
				s.Cycle = &Cycle{}
				defer func() { s.Cycle = nil }()
			}
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, TypeError(nil, 0, "table")
			}
			if s.Cycle.Has(table) {
				return nil, fmt.Errorf("tables cannot be cyclic")
			}
			s.Cycle.Put(table)
			variantType := s.Type("Variant")
			n := table.Len()
			array := make([]Value, n)
			for i := 1; i <= n; i++ {
				if array[i-1], err = variantType.ReflectFrom(s, variantType, table.RawGetInt(i)); err != nil {
					return nil, err
				}
			}
			return array, nil
		},
	}
}
