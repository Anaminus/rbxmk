package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/rtypes"
)

func init() { register(Symbol) }
func Symbol() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "Symbol",
		PushTo:   rbxmk.PushPtrTypeTo("Symbol"),
		PullFrom: rbxmk.PullTypeFrom("Symbol"),
		Metatable: rbxmk.Metatable{
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "Symbol").(rtypes.Symbol)
				op := s.Pull(2, "Symbol").(rtypes.Symbol)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
