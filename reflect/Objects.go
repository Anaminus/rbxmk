package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/types"
	"github.com/yuin/gopher-lua"
)

func Objects() Type {
	return Type{
		Name: "Objects",
		ReflectTo: func(s State, t Type, v Value) (lvs []lua.LValue, err error) {
			objects, ok := v.([]*types.Instance)
			if !ok {
				return nil, TypeError(nil, 0, "[]*types.Instance")
			}
			instType := s.Type("Instance")
			table := s.L.CreateTable(len(objects), 0)
			for i, v := range objects {
				lv, err := instType.ReflectTo(s, instType, v)
				if err != nil {
					return nil, err
				}
				table.RawSetInt(i+1, lv[0])
			}
			return []lua.LValue{table}, nil
		},
		ReflectFrom: func(s State, t Type, lvs ...lua.LValue) (v Value, err error) {
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, TypeError(nil, 0, "table")
			}
			instType := s.Type("Instance")
			n := table.Len()
			objects := make([]Value, n)
			for i := 1; i <= n; i++ {
				if objects[i-1], err = instType.ReflectFrom(s, instType, table.RawGetInt(i)); err != nil {
					return nil, err
				}
			}
			return objects, nil
		},
	}
}
