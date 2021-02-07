package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(NumberRange) }
func NumberRange() Reflector {
	return Reflector{
		Name:     "NumberRange",
		PushTo:   rbxmk.PushTypeTo("NumberRange"),
		PullFrom: rbxmk.PullTypeFrom("NumberRange"),
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "NumberRange").(types.NumberRange)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "NumberRange").(types.NumberRange)
				op := s.Pull(2, "NumberRange").(types.NumberRange)
				s.L.Push(lua.LBool(v == op))
				return 1
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
