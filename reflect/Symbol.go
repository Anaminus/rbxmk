package reflect

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Symbol) }
func Symbol() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     rtypes.T_Symbol,
		PushTo:   rbxmk.PushPtrTypeTo(rtypes.T_Symbol),
		PullFrom: rbxmk.PullTypeFrom(rtypes.T_Symbol),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case **rtypes.Symbol:
				*p = v.(*rtypes.Symbol)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/Symbol:Summary",
				Description: "Types/Symbol:Description",
			}
		},
	}
}
