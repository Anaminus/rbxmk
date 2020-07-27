package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Instances() Type {
	return Type{
		Name: "Instances",
		PushTo: func(s State, t Type, v types.Value) (lvs []lua.LValue, err error) {
			instances, ok := v.(rtypes.Instances)
			if !ok {
				return nil, TypeError(nil, 0, "Instances")
			}
			instType := s.Type("Instance")
			table := s.L.CreateTable(len(instances), 0)
			for i, v := range instances {
				lv, err := instType.PushTo(s, instType, v)
				if err != nil {
					return nil, err
				}
				table.RawSetInt(i+1, lv[0])
			}
			return []lua.LValue{table}, nil
		},
		PullFrom: func(s State, t Type, lvs ...lua.LValue) (v types.Value, err error) {
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, TypeError(nil, 0, "table")
			}
			instType := s.Type("Instance")
			n := table.Len()
			instances := make(rtypes.Instances, n)
			for i := 1; i <= n; i++ {
				v, err := instType.PullFrom(s, instType, table.RawGetInt(i))
				if err != nil {
					return nil, err
				}
				instances[i-1] = v.(*rtypes.Instance)
			}
			return instances, nil
		},
	}
}
