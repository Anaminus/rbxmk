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
		PushTo: func(s rbxmk.State, v types.Value) (lv lua.LValue, err error) {
			objects, ok := v.(rtypes.Objects)
			if !ok {
				return nil, rbxmk.TypeError{Want: "Objects", Got: v.Type()}
			}
			instRfl := s.MustReflector("Instance")
			table := s.L.CreateTable(len(objects), 0)
			for i, v := range objects {
				lv, err := instRfl.PushTo(s, v)
				if err != nil {
					return nil, err
				}
				table.RawSetInt(i+1, lv)
			}
			return table, nil
		},
		PullFrom: func(s rbxmk.State, lv lua.LValue) (v types.Value, err error) {
			table, ok := lv.(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError{Want: "table", Got: lv.Type().String()}
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
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying:  dt.Array{T: dt.Prim("Instance")},
				Summary:     "Types/Objects:Summary",
				Description: "Types/Objects:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Instance,
		},
	}
}
