package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
)

func init() { register(DescAction) }
func DescAction() Reflector {
	return Reflector{
		Name:     "DescAction",
		PushTo:   rbxmk.PushTypeTo,
		PullFrom: rbxmk.PullTypeFrom,
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
