package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Color3uint8() Type {
	return Type{
		Name:   "Color3uint8",
		Flags:  Exprim,
		PushTo: PushTypeTo,
		PullFrom: func(s State, t Type, lvs ...lua.LValue) (v types.Value, err error) {
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
			return nil, TypeError(nil, 0, "Color3uint8")
		},
	}
}
