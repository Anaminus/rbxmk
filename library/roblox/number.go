package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/robloxapi/types"
)

func init() { register(Number) }
func Number() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "number",
		PushTo: func(s rbxmk.State, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Double))}, nil
		},
		PullFrom: func(s rbxmk.State, lvs ...lua.LValue) (v types.Value, err error) {
			if n, ok := lvs[0].(lua.LNumber); ok {
				return types.Double(n), nil
			}
			return nil, rbxmk.TypeError{Want: "number", Got: lvs[0].Type().String()}
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Libraries/roblox/Types/number:Summary",
				Description: "Libraries/roblox/Types/number:Description",
			}
		},
	}
}
