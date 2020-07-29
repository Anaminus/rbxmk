package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func ParameterDesc() Type {
	return Type{
		Name:     "ParameterDesc",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Members: Members{
			"Type": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.ParameterDesc)
					return s.Push(rtypes.TypeDesc{Embedded: desc.Parameter.Type})
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.ParameterDesc)
					desc.Parameter.Type = s.Pull(3, "TypeDesc").(rtypes.TypeDesc).Embedded
				},
			},
			"Name": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.ParameterDesc)
					return s.Push(types.String(desc.Name))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.ParameterDesc)
					desc.Name = string(s.Pull(3, "string").(types.String))
				},
			},
			"Default": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.ParameterDesc)
					if !desc.Optional {
						return s.Push(rtypes.Nil)
					}
					return s.Push(types.String(desc.Default))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.ParameterDesc)
					switch value := s.PullAnyOf(3, "string", "nil").(type) {
					case rtypes.NilType:
						desc.Optional = false
						desc.Default = ""
					case types.String:
						desc.Optional = true
						desc.Default = string(value)
					}
				},
			},
		},
	}
}
