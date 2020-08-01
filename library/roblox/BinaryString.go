package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func init() { register(BinaryString) }
func BinaryString() Reflector {
	return Reflector{
		Name:  "BinaryString",
		Flags: Exprim,
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LString(v.(types.BinaryString))}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			switch v := lvs[0].(type) {
			case lua.LString:
				return types.BinaryString(v), nil
			case *lua.LUserData:
				if v.Metatable == s.L.GetTypeMetatable("BinaryString") {
					if v, ok := v.Value.(types.BinaryString); ok {
						return v, nil
					}
				}
			}
			return nil, TypeError(nil, 0, "BinaryString")
		},
	}
}
