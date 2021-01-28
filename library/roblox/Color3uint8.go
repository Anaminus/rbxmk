package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Color3uint8) }
func Color3uint8() Reflector {
	return Reflector{
		Name:   "Color3uint8",
		Flags:  rbxmk.Exprim,
		PushTo: rbxmk.PushTypeTo,
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			if u, ok := lvs[0].(*lua.LUserData); ok {
				switch u.Metatable {
				case s.L.GetTypeMetatable("Color3"):
					if v, ok := u.Value.(types.Color3); ok {
						return rtypes.Color3uint8(v), nil
					}
				case s.L.GetTypeMetatable("Color3uint8"):
					if v, ok := u.Value.(rtypes.Color3uint8); ok {
						return v, nil
					}
				}
			}
			return nil, rbxmk.TypeError(nil, 0, "Color3uint8")
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
	}
}
