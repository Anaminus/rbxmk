package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	lua "github.com/yuin/gopher-lua"
)

func init() { register(Symbol) }
func Symbol() Reflector {
	return Reflector{
		Name:     "Symbol",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__eq": func(s State) int {
				v := s.Pull(1, "Symbol").(rtypes.Symbol)
				op := s.Pull(2, "Symbol").(rtypes.Symbol)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
	}
}
