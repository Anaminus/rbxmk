package reflect

import (
	"strconv"

	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func NumberSequenceKeypoint() Type {
	return Type{
		Name:        "NumberSequenceKeypoint",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				u := v.(types.NumberSequenceKeypoint)
				var b string
				b += strconv.FormatFloat(float64(u.Time), 'g', -1, 32) + ", "
				b += strconv.FormatFloat(float64(u.Value), 'g', -1, 32) + ", "
				b += strconv.FormatFloat(float64(u.Envelope), 'g', -1, 32)
				s.L.Push(lua.LString(b))
				return 1
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "NumberSequenceKeypoint").(types.NumberSequenceKeypoint)
				return s.Push("bool", v.(types.NumberSequenceKeypoint) == op)
			},
		},
		Members: map[string]Member{
			"Time": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.NumberSequenceKeypoint).Time)
			}},
			"Value": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.NumberSequenceKeypoint).Value)
			}},
			"Envelope": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.NumberSequenceKeypoint).Envelope)
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.NumberSequenceKeypoint
				switch s.Count() {
				case 2:
					v.Time = s.Pull(1, "float").(float32)
					v.Value = s.Pull(2, "float").(float32)
				case 3:
					v.Time = s.Pull(1, "float").(float32)
					v.Value = s.Pull(2, "float").(float32)
					v.Envelope = s.Pull(3, "float").(float32)
				default:
					s.L.RaiseError("expected 2 or 3 arguments")
					return 0
				}
				return s.Push("NumberSequenceKeypoint", v)
			},
		},
	}
}
