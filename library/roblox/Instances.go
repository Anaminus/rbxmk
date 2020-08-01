package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
	lua "github.com/yuin/gopher-lua"
)

func init() { register(Instances) }
func Instances() Reflector {
	return Reflector{
		Name: "Instances",
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			instances, ok := v.(rtypes.Instances)
			if !ok {
				return nil, TypeError(nil, 0, "Instances")
			}
			instRfl := s.Reflector("Instance")
			table := s.L.CreateTable(len(instances), 0)
			for i, v := range instances {
				lv, err := instRfl.PushTo(s, instRfl, v)
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
				return nil, TypeError(nil, 0, "table")
			}
			instRfl := s.Reflector("Instance")
			n := table.Len()
			instances := make(rtypes.Instances, n)
			for i := 1; i <= n; i++ {
				v, err := instRfl.PullFrom(s, instRfl, table.RawGetInt(i))
				if err != nil {
					return nil, err
				}
				instances[i-1] = v.(*rtypes.Instance)
			}
			return instances, nil
		},
	}
}
