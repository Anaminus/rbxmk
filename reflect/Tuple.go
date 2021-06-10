package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/robloxapi/types"
)

func init() { register(Tuple) }
func Tuple() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "Tuple",
		PushTo: func(s rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			panic("incorrect use of Tuple reflection (use State.PushTuple)")
		},
		PullFrom: func(s rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			panic("incorrect use of Tuple reflection (use State.PullTuple)")
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
