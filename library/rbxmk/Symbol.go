package reflect

import (
	"github.com/anaminus/rbxmk"
)

func init() { register(Symbol) }
func Symbol() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "Symbol",
		PushTo:   rbxmk.PushPtrTypeTo("Symbol"),
		PullFrom: rbxmk.PullTypeFrom("Symbol"),
	}
}
