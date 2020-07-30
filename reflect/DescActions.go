package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/yuin/gopher-lua"
)

func DescActions() Type {
	return Type{
		Name:     "DescActions",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString(s.Pull(1, "DescActions").(rtypes.DescActions).String()))
				return 1
			},
		},
	}
}
