package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func TypeDesc() Type {
	return Type{
		Name:     "TypeDesc",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Members: Members{
			"Category": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.TypeDesc)
					return s.Push(types.String(desc.Embedded.Category))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.TypeDesc)
					desc.Embedded.Category = string(s.Pull(3, "string").(types.String))
				},
			},
			"Name": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.TypeDesc)
					return s.Push(types.String(desc.Embedded.Name))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.TypeDesc)
					desc.Embedded.Name = string(s.Pull(3, "string").(types.String))
				},
			},
		},
	}
}
