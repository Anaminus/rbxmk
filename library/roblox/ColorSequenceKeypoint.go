package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(ColorSequenceKeypoint) }
func ColorSequenceKeypoint() Reflector {
	return Reflector{
		Name:     "ColorSequenceKeypoint",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "ColorSequenceKeypoint").(types.ColorSequenceKeypoint)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "ColorSequenceKeypoint").(types.ColorSequenceKeypoint)
				op := s.Pull(2, "ColorSequenceKeypoint").(types.ColorSequenceKeypoint)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: map[string]Member{
			"Time": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.ColorSequenceKeypoint).Time))
			}},
			"Value": {Get: func(s State, v types.Value) int {
				return s.Push(v.(types.ColorSequenceKeypoint).Value)
			}},
			"Envelope": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.ColorSequenceKeypoint).Envelope))
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.ColorSequenceKeypoint
				switch s.Count() {
				case 2:
					v.Time = float32(s.Pull(1, "float").(types.Float))
					v.Value = s.Pull(2, "Color3").(types.Color3)
				case 3:
					v.Time = float32(s.Pull(1, "float").(types.Float))
					v.Value = s.Pull(2, "Color3").(types.Color3)
					v.Envelope = float32(s.Pull(3, "float").(types.Float))
				default:
					return s.RaiseError("expected 2 or 3 arguments")
				}
				return s.Push(v)
			},
		},
	}
}
