package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Optional) }
func Optional() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:   "Optional",
		Flags:  rbxmk.Exprim,
		PushTo: rbxmk.PushTypeTo("Optional"),
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			if u, ok := lv.(*lua.LUserData); ok {
				if u.Metatable == c.GetTypeMetatable("Optional") {
					if v, ok := u.Value().(rtypes.Optional); ok {
						return v, nil
					}
				}
			}
			if v, err = PullVariantFrom(c, lv); err != nil {
				//TODO: Better error?
				return nil, err
			}
			return rtypes.Some(v), nil
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.Optional:
				//TODO: Is it within the scope of SetTo to set the content of
				// the optional?
				*p = v.(rtypes.Optional)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case rtypes.Optional:
				return v
			default:
				return rtypes.Some(v)
			}
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Underlying:  dt.Prim("Optional"),
				Summary:     "Types/Optional:Summary",
				Description: "Types/Optional:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Variant,
		},
	}
}
