package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	lua "github.com/yuin/gopher-lua"
)

func init() { register(DescAction) }
func DescAction() Reflector {
	return Reflector{
		Name:     "DescAction",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
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
	}
}
