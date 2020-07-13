package reflect

import (
	"strconv"

	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Rect() Type {
	return Type{
		Name:        "Rect",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				u := v.(types.Rect)
				var b string
				b += strconv.FormatFloat(float64(u.Min.X), 'g', -1, 32) + ", "
				b += strconv.FormatFloat(float64(u.Min.Y), 'g', -1, 32) + ", "
				b += strconv.FormatFloat(float64(u.Max.X), 'g', -1, 32) + ", "
				b += strconv.FormatFloat(float64(u.Max.Y), 'g', -1, 32)
				s.L.Push(lua.LString(b))
				return 1
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "Rect").(types.Rect)
				return s.Push("bool", v.(types.Rect) == op)
			},
		},
		Members: map[string]Member{
			"Min": {Get: func(s State, v Value) int {
				return s.Push("Vector2", v.(types.Rect).Min)
			}},
			"Max": {Get: func(s State, v Value) int {
				return s.Push("Vector2", v.(types.Rect).Max)
			}},
			"Width": {Get: func(s State, v Value) int {
				return s.Push("number", v.(types.Rect).Width())
			}},
			"Height": {Get: func(s State, v Value) int {
				return s.Push("number", v.(types.Rect).Height())
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.Rect
				switch s.Count() {
				case 2:
					v.Min = s.Pull(1, "Vector2").(types.Vector2)
					v.Max = s.Pull(2, "Vector2").(types.Vector2)
				case 4:
					v.Min.X = s.Pull(1, "float").(float32)
					v.Min.Y = s.Pull(2, "float").(float32)
					v.Max.Y = s.Pull(3, "float").(float32)
					v.Max.Y = s.Pull(4, "float").(float32)
				default:
					s.L.RaiseError("expected 2 or 4 arguments")
					return 0
				}
				return s.Push("Rect", v)
			},
		},
	}
}
