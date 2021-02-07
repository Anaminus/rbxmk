package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Cookies) }
func Cookies() Reflector {
	return Reflector{
		Name: "Cookies",
		PushTo: func(s State, v types.Value) (lvs []lua.LValue, err error) {
			cookies, ok := v.(rtypes.Cookies)
			if !ok {
				return nil, rbxmk.TypeError(nil, 0, "Cookies")
			}
			cookieRfl := s.Reflector("Cookie")
			table := s.L.CreateTable(len(cookies), 0)
			for i, v := range cookies {
				lv, err := cookieRfl.PushTo(s, v)
				if err != nil {
					return nil, err
				}
				table.RawSetInt(i+1, lv[0])
			}
			return []lua.LValue{table}, nil
		},
		PullFrom: func(s State, lvs ...lua.LValue) (v types.Value, err error) {
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, rbxmk.TypeError(nil, 0, "table")
			}
			cookieRfl := s.Reflector("Cookie")
			n := table.Len()
			cookies := make(rtypes.Cookies, n)
			for i := 1; i <= n; i++ {
				v, err := cookieRfl.PullFrom(s, table.RawGetInt(i))
				if err != nil {
					return nil, err
				}
				cookies[i-1] = v.(rtypes.Cookie)
			}
			return cookies, nil
		},
	}
}
