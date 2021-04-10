package reflect

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
)

func init() { register(Symbol) }
func Symbol() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "Symbol",
		PushTo:   rbxmk.PushPtrTypeTo("Symbol"),
		PullFrom: rbxmk.PullTypeFrom("Symbol"),
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/Symbol:Summary",
				Description: "Types/Symbol:Description",
			}
		},
	}
}
