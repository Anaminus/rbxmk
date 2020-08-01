package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/yuin/gopher-lua"
)

func DescActions() Reflector {
	return Reflector{
		Name:     "DescActions",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "DescActions").(rtypes.DescActions)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
		},
	}
}
