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
		Name: rtypes.T_Objects,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			objects, ok := v.(rtypes.Objects)
			if !ok {
				return nil, rbxmk.TypeError{Want: rtypes.T_Objects, Got: v.Type()}
			}
			instRfl := c.MustReflector(rtypes.T_Instance)
			table := c.CreateTable(len(objects), 0)
			for i, v := range objects {
				lv, err := instRfl.PushTo(c, v)
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
			instRfl := c.MustReflector(rtypes.T_Instance)
			n := table.Len()
			objects := make(rtypes.Objects, n)
			for i := 1; i <= n; i++ {
				v, err := instRfl.PullFrom(c, table.RawGetInt(i))
				if err != nil {
					return nil, err
				}
				objects[i-1] = v.(*rtypes.Instance)
			}
			return objects, nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.Objects:
				*p = v.(rtypes.Objects)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying:  dt.Array{T: dt.Prim(rtypes.T_Instance)},
				Summary:     "Types/Objects:Summary",
				Description: "Types/Objects:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Instance,
		},
	}
}
