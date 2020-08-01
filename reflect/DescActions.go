package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func DescActions() Reflector {
	return Reflector{
		Name:     "DescActions",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "DescActions").(rtypes.DescActions)
				return s.Push(types.String(v.String()))
			},
		},
	}
}
