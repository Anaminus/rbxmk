package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(Rect) }
func Rect() Reflector {
	return Reflector{
		Name:     "Rect",
		PushTo:   rbxmk.PushTypeTo("Rect"),
		PullFrom: rbxmk.PullTypeFrom("Rect"),
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "Rect").(types.Rect)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "Rect").(types.Rect)
				op := s.Pull(2, "Rect").(types.Rect)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: map[string]Member{
			"Min": {Get: func(s State, v types.Value) int {
				return s.Push(v.(types.Rect).Min)
			}},
			"Max": {Get: func(s State, v types.Value) int {
				return s.Push(v.(types.Rect).Max)
			}},
			"Width": {Get: func(s State, v types.Value) int {
				return s.Push(types.Double(v.(types.Rect).Width()))
			}},
			"Height": {Get: func(s State, v types.Value) int {
				return s.Push(types.Double(v.(types.Rect).Height()))
			}},
		},
		Constructors: Constructors{
			"new": {
				Func: func(s State) int {
					var v types.Rect
					switch s.Count() {
					case 2:
						v.Min = s.Pull(1, "Vector2").(types.Vector2)
						v.Max = s.Pull(2, "Vector2").(types.Vector2)
					case 4:
						v.Min.X = float32(s.Pull(1, "float").(types.Float))
						v.Min.Y = float32(s.Pull(2, "float").(types.Float))
						v.Max.Y = float32(s.Pull(3, "float").(types.Float))
						v.Max.Y = float32(s.Pull(4, "float").(types.Float))
					default:
						return s.RaiseError("expected 2 or 4 arguments")
					}
					return s.Push(v)
				},
			},
		},
	}
}
