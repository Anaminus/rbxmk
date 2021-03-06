package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/rtypes"
)

func init() { register(DescAction) }
func DescAction() Reflector {
	return Reflector{
		Name:     "DescAction",
		PushTo:   rbxmk.PushTypeTo("DescAction"),
		PullFrom: rbxmk.PullTypeFrom("DescAction"),
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "DescAction").(*rtypes.DescAction)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "DescAction").(*rtypes.DescAction)
				op := s.Pull(2, "DescAction").(*rtypes.DescAction)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
