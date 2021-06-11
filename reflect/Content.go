package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/robloxapi/types"
)

func init() { register(Content) }
func Content() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:  "Content",
		Flags: rbxmk.Exprim,
		PushTo: func(s rbxmk.Context, v types.Value) (lv lua.LValue, err error) {
			return lua.LString(v.(types.Content)), nil
		},
		PullFrom: func(s rbxmk.Context, lv lua.LValue) (v types.Value, err error) {
			switch v := lv.(type) {
			case lua.LString:
				return types.Content(v), nil
			case *lua.LUserData:
				if v.Metatable == s.GetTypeMetatable("Content") {
					if v, ok := v.Value().(types.Content); ok {
						return v, nil
					}
				}
			}
			return nil, rbxmk.TypeError{Want: "Content", Got: lv.Type().String()}
		},
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.Content:
				*p = v.(types.Content)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case types.Content:
				return v
			case types.Stringlike:
				return types.Content(v.Stringlike())
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/Content:Summary",
				Description: "Types/Content:Description",
			}
		},
	}
}
