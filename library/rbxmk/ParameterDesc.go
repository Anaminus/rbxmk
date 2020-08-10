package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
	lua "github.com/yuin/gopher-lua"
)

func init() { register(ParameterDesc) }
func ParameterDesc() Reflector {
	return Reflector{
		Name:     "ParameterDesc",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__eq": func(s State) int {
				v := s.Pull(1, "ParameterDesc").(rtypes.ParameterDesc)
				op := s.Pull(2, "ParameterDesc").(rtypes.ParameterDesc)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: Members{
			"Type": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(*rtypes.ParameterDesc)
					typ := desc.Parameter.Type
					return s.Push(rtypes.TypeDesc{Embedded: typ})
				},
				Set: func(s State, v types.Value) {
					desc := v.(*rtypes.ParameterDesc)
					desc.Parameter.Type = s.Pull(3, "TypeDesc").(rtypes.TypeDesc).Embedded
				},
			},
			"Name": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(*rtypes.ParameterDesc)
					return s.Push(types.String(desc.Name))
				},
				Set: func(s State, v types.Value) {
					desc := v.(*rtypes.ParameterDesc)
					desc.Name = string(s.Pull(3, "string").(types.String))
				},
			},
			"Default": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(*rtypes.ParameterDesc)
					if !desc.Optional {
						return s.Push(rtypes.Nil)
					}
					return s.Push(types.String(desc.Default))
				},
				Set: func(s State, v types.Value) {
					desc := v.(*rtypes.ParameterDesc)
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
