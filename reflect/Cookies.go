package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Cookies) }
func Cookies() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "Cookies",
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			cookies, ok := v.(rtypes.Cookies)
			if !ok {
				return nil, rbxmk.TypeError{Want: "Cookies", Got: v.Type()}
			}
			cookieRfl := c.MustReflector("Cookie")
			table := c.CreateTable(len(cookies), 0)
			for i, v := range cookies {
				lv, err := cookieRfl.PushTo(c, v)
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
				return nil, rbxmk.TypeError{Want: "table", Got: lv.Type().String()}
			}
			cookieRfl := c.MustReflector("Cookie")
			n := table.Len()
			cookies := make(rtypes.Cookies, n)
			for i := 1; i <= n; i++ {
				v, err := cookieRfl.PullFrom(c, table.RawGetInt(i))
				if err != nil {
					return nil, err
				}
				cookies[i-1] = v.(rtypes.Cookie)
			}
			return cookies, nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.Cookies:
				*p = v.(rtypes.Cookies)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying:  dt.Array{T: dt.Prim("Cookie")},
				Summary:     "Types/Cookies:Summary",
				Description: "Types/Cookies:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Cookie,
		},
	}
}
