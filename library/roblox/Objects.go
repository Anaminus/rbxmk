package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Objects) }
func Objects() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "Objects",
		PushTo: func(s rbxmk.State, v types.Value) (lvs []lua.LValue, err error) {
			objects, ok := v.(rtypes.Objects)
			if !ok {
				return nil, rbxmk.TypeError(nil, 0, "Objects")
			}
			instRfl := s.MustReflector("Instance")
			table := s.L.CreateTable(len(objects), 0)
			for i, v := range objects {
				lv, err := instRfl.PushTo(s, v)
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
				return nil, rbxmk.TypeError(nil, 0, "table")
			}
			instRfl := s.MustReflector("Instance")
			n := table.Len()
			objects := make(rtypes.Objects, n)
			for i := 1; i <= n; i++ {
				v, err := instRfl.PullFrom(s, table.RawGetInt(i))
				if err != nil {
					return nil, err
				}
				objects[i-1] = v.(*rtypes.Instance)
			}
			return objects, nil
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Underlying: dt.Array{T: dt.Prim("Instance")}} },
	}
}
