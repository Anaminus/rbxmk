package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Tuple) }
func Tuple() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "Tuple",
		PushTo: func(c rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			panic("incorrect use of Tuple reflection (use State.PushTuple)")
		},
		PullFrom: func(c rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			panic("incorrect use of Tuple reflection (use State.PullTuple)")
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.Tuple:
				*p = v.(rtypes.Tuple)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/Tuple:Summary",
				Description: "Types/Tuple:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Variant,
		},
	}
}
