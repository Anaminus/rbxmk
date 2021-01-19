package reflect

import (
	"math/rand"

	lua "github.com/anaminus/gopher-lua"
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(BrickColor) }
func BrickColor() Reflector {
	return Reflector{
		Name:     "BrickColor",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "BrickColor").(types.BrickColor)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "BrickColor").(types.BrickColor)
				op := s.Pull(2, "BrickColor").(types.BrickColor)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: map[string]Member{
			"Name": {Get: func(s State, v types.Value) int {
				return s.Push(types.String(v.(types.BrickColor).Name()))
			}},
			"Number": {Get: func(s State, v types.Value) int {
				return s.Push(types.Int(v.(types.BrickColor).Number()))
			}},
			"R": {Get: func(s State, v types.Value) int {
				return s.Push(types.Double(v.(types.BrickColor).R()))
			}},
			"G": {Get: func(s State, v types.Value) int {
				return s.Push(types.Double(v.(types.BrickColor).G()))
			}},
			"B": {Get: func(s State, v types.Value) int {
				return s.Push(types.Double(v.(types.BrickColor).B()))
			}},
			"Color": {Get: func(s State, v types.Value) int {
				return s.Push(v.(types.BrickColor).Color())
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.BrickColor
				switch s.Count() {
				case 1:
					switch arg := s.PullAnyOf(1, "int", "string", "Color3").(type) {
					case types.Int:
						v = types.NewBrickColor(int(arg))
					case types.String:
						v = types.NewBrickColorFromName(string(arg))
					case types.Color3:
						v = types.NewBrickColorFromColor3(arg)
					}
				case 3:
					v = types.NewBrickColorFromColor(
						float64(s.Pull(1, "number").(types.Double)),
						float64(s.Pull(2, "number").(types.Double)),
						float64(s.Pull(3, "number").(types.Double)),
					)
				default:
					return s.RaiseError("expected 1 or 3 arguments")
				}
				return s.Push(v)
			},
			"palette": func(s State) int {
				index := int(s.Pull(1, "int").(types.Int))
				return s.Push(types.NewBrickColorFromPalette(index))
			},
			"random": func(s State) int {
				index := rand.Intn(types.BrickColorIndexSize)
				return s.Push(types.NewBrickColorFromIndex(index))
			},
			"White": func(s State) int {
				return s.Push(types.NewBrickColorFromName("White"))
			},
			"Gray": func(s State) int {
				return s.Push(types.NewBrickColorFromName("Medium stone grey"))
			},
			"DarkGray": func(s State) int {
				return s.Push(types.NewBrickColorFromName("Dark stone grey"))
			},
			"Black": func(s State) int {
				return s.Push(types.NewBrickColorFromName("Black"))
			},
			"Red": func(s State) int {
				return s.Push(types.NewBrickColorFromName("Bright red"))
			},
			"Yellow": func(s State) int {
				return s.Push(types.NewBrickColorFromName("Bright yellow"))
			},
			"Green": func(s State) int {
				return s.Push(types.NewBrickColorFromName("Dark green"))
			},
			"Blue": func(s State) int {
				return s.Push(types.NewBrickColorFromName("Bright blue"))
			},
		},
	}
}
