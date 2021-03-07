package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

func init() { register(Bool) }
func Bool() Reflector {
	return Reflector{
		Name: "bool",
		PushTo: func(s State, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LBool(v.(types.Bool))}, nil
		},
		PullFrom: func(s State, lvs ...lua.LValue) (v types.Value, err error) {
			if n, ok := lvs[0].(lua.LBool); ok {
				return types.Bool(n), nil
			}
			return nil, rbxmk.TypeError(nil, 0, "bool")
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Underlying: dt.Prim("boolean")} },
	}
}
