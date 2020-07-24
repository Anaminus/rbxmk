package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func ColorSequenceKeypoint() Type {
	return Type{
		Name:        "ColorSequenceKeypoint",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString(s.Pull(1, "ColorSequenceKeypoint").(types.ColorSequenceKeypoint).String()))
				return 1
			},
			"__eq": func(s State) int {
				op := s.Pull(2, "ColorSequenceKeypoint").(types.ColorSequenceKeypoint)
				return s.Push(types.Bool(s.Pull(1, "ColorSequenceKeypoint").(types.ColorSequenceKeypoint) == op))
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
					s.L.RaiseError("expected 2 or 3 arguments")
					return 0
				}
				return s.Push(v)
			},
		},
	}
}
