package reflect

import (
	"strconv"

	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func NumberRange() Type {
	return Type{
		Name:        "NumberRange",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				u := v.(types.NumberRange)
				var b string
				b += strconv.FormatFloat(float64(u.Min), 'g', -1, 32) + ", "
				b += strconv.FormatFloat(float64(u.Max), 'g', -1, 32)
				s.L.Push(lua.LString(b))
				return 1
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "NumberRange").(types.NumberRange)
				return s.Push("bool", v.(types.NumberRange) == op)
			},
		},
		Members: map[string]Member{
			"Min": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.NumberRange).Min)
			}},
			"Max": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.NumberRange).Max)
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.NumberRange
				switch s.Count() {
				case 1:
					v.Min = s.Pull(1, "float").(float32)
					v.Max = v.Min
				case 2:
					v.Min = s.Pull(1, "float").(float32)
					v.Max = s.Pull(2, "float").(float32)
				default:
					s.L.RaiseError("expected 1 or 2 arguments")
					return 0
				}
				return s.Push("NumberRange", v)
			},
		},
	}
}
