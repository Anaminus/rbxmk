package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(NumberSequenceKeypoint) }
func NumberSequenceKeypoint() Reflector {
	return Reflector{
		Name:     "NumberSequenceKeypoint",
		PushTo:   rbxmk.PushTypeTo("NumberSequenceKeypoint"),
		PullFrom: rbxmk.PullTypeFrom("NumberSequenceKeypoint"),
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "NumberSequenceKeypoint").(types.NumberSequenceKeypoint)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "NumberSequenceKeypoint").(types.NumberSequenceKeypoint)
				op := s.Pull(2, "NumberSequenceKeypoint").(types.NumberSequenceKeypoint)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: map[string]Member{
			"Time": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.NumberSequenceKeypoint).Time))
			}},
			"Value": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.NumberSequenceKeypoint).Value))
			}},
			"Envelope": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.NumberSequenceKeypoint).Envelope))
			}},
		},
		Constructors: Constructors{
			"new": {
				Func: func(s State) int {
					var v types.NumberSequenceKeypoint
					switch s.Count() {
					case 2:
						v.Time = float32(s.Pull(1, "float").(types.Float))
						v.Value = float32(s.Pull(2, "float").(types.Float))
					case 3:
						v.Time = float32(s.Pull(1, "float").(types.Float))
						v.Value = float32(s.Pull(2, "float").(types.Float))
						v.Envelope = float32(s.Pull(3, "float").(types.Float))
					default:
						return s.RaiseError("expected 2 or 3 arguments")
					}
					return s.Push(v)
				},
			},
		},
	}
}
