package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func NumberRange() Reflector {
	return Reflector{
		Name:     "NumberRange",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString(s.Pull(1, "NumberRange").(types.NumberRange).String()))
				return 1
			},
			"__eq": func(s State) int {
				op := s.Pull(2, "NumberRange").(types.NumberRange)
				return s.Push(types.Bool(s.Pull(1, "NumberRange").(types.NumberRange) == op))
			},
		},
		Members: map[string]Member{
			"Min": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.NumberRange).Min))
			}},
			"Max": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.NumberRange).Max))
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.NumberRange
				switch s.Count() {
				case 1:
					v.Min = float32(s.Pull(1, "float").(types.Float))
					v.Max = v.Min
				case 2:
					v.Min = float32(s.Pull(1, "float").(types.Float))
					v.Max = float32(s.Pull(2, "float").(types.Float))
				default:
					return s.RaiseError("expected 1 or 2 arguments")
				}
				return s.Push(v)
			},
		},
	}
}
