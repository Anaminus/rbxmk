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
		Name:   "Color3uint8",
		Flags:  rbxmk.Exprim,
		PushTo: rbxmk.PushTypeTo("Color3uint8"),
		PullFrom: func(s rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			if u, ok := lv.(*lua.LUserData); ok {
				switch u.Metatable {
				case s.GetTypeMetatable("Color3"):
					if v, ok := u.Value().(types.Color3); ok {
						return rtypes.Color3uint8(v), nil
					}
				case s.GetTypeMetatable("Color3uint8"):
					if v, ok := u.Value().(rtypes.Color3uint8); ok {
						return v, nil
					}
				}
			}
			return nil, rbxmk.TypeError{Want: "Color3uint8", Got: lv.Type().String()}
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
				Underlying:  dt.Prim("Color3"),
				Summary:     "Types/Color3uint8:Summary",
				Description: "Types/Color3uint8:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Color3,
		},
	}
}
