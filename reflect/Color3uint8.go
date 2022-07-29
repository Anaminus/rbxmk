package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Color3uint8) }
func Color3uint8() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:  rtypes.T_Color3uint8,
		Flags: rbxmk.Exprim,
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			u := c.NewUserData(types.Color3(v.(rtypes.Color3uint8)))
			c.SetMetatable(u, c.GetTypeMetatable(rtypes.T_Color3))
			return u, nil
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			if u, ok := lv.(*lua.LUserData); ok {
				switch u.Metatable {
				case c.GetTypeMetatable(rtypes.T_Color3):
					if v, ok := u.Value().(types.Color3); ok {
						return rtypes.Color3uint8(v), nil
					}
				case c.GetTypeMetatable(rtypes.T_Color3uint8):
					if v, ok := u.Value().(rtypes.Color3uint8); ok {
						return v, nil
					}
				}
			}
			return nil, rbxmk.TypeError{Want: rtypes.T_Color3uint8, Got: lv.Type().String()}
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.Color3uint8:
				*p = v.(rtypes.Color3uint8)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case rtypes.Color3uint8:
				return v
			case types.Color3:
				return rtypes.Color3uint8(v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category:    "roblox",
				Underlying:  dt.P(dt.Prim(rtypes.T_Color3)),
				Summary:     "Types/Color3uint8:Summary",
				Description: "Types/Color3uint8:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Color3,
		},
	}
}
