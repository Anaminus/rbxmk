package reflect

import (
	"math/rand"

	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func BrickColor() Type {
	return Type{
		Name:        "BrickColor",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State, v types.Value) int {
				s.L.Push(lua.LString(v.(types.BrickColor).String()))
				return 1
			},
			"__eq": func(s State, v types.Value) int {
				op := s.Pull(2, "BrickColor").(types.BrickColor)
				return s.Push("bool", types.Bool(v.(types.BrickColor) == op))
			},
		},
		Members: map[string]Member{
			"Name": {Get: func(s State, v types.Value) int {
				return s.Push("string", types.String(v.(types.BrickColor).Name()))
			}},
			"Number": {Get: func(s State, v types.Value) int {
				return s.Push("int", types.Int(v.(types.BrickColor).Number()))
			}},
			"R": {Get: func(s State, v types.Value) int {
				return s.Push("number", types.Double(v.(types.BrickColor).R()))
			}},
			"G": {Get: func(s State, v types.Value) int {
				return s.Push("number", types.Double(v.(types.BrickColor).G()))
			}},
			"B": {Get: func(s State, v types.Value) int {
				return s.Push("number", types.Double(v.(types.BrickColor).B()))
			}},
			"Color": {Get: func(s State, v types.Value) int {
				return s.Push("Color3", v.(types.BrickColor).Color())
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
					s.L.RaiseError("expected 1 or 3 arguments")
					return 0
				}
				return s.Push("BrickColor", v)
			},
			"palette": func(s State) int {
				index := int(s.Pull(1, "int").(types.Int))
				return s.Push("BrickColor", types.NewBrickColorFromPalette(index))
			},
			"random": func(s State) int {
				index := rand.Intn(types.BrickColorIndexSize)
				return s.Push("BrickColor", types.NewBrickColorFromIndex(index))
			},
			"White": func(s State) int {
				return s.Push("BrickColor", types.NewBrickColorFromName("White"))
			},
			"Gray": func(s State) int {
				return s.Push("BrickColor", types.NewBrickColorFromName("Medium stone grey"))
			},
			"DarkGray": func(s State) int {
				return s.Push("BrickColor", types.NewBrickColorFromName("Dark stone grey"))
			},
			"Black": func(s State) int {
				return s.Push("BrickColor", types.NewBrickColorFromName("Black"))
			},
			"Red": func(s State) int {
				return s.Push("BrickColor", types.NewBrickColorFromName("Bright red"))
			},
			"Yellow": func(s State) int {
				return s.Push("BrickColor", types.NewBrickColorFromName("Bright yellow"))
			},
			"Green": func(s State) int {
				return s.Push("BrickColor", types.NewBrickColorFromName("Dark green"))
			},
			"Blue": func(s State) int {
				return s.Push("BrickColor", types.NewBrickColorFromName("Bright blue"))
			},
		},
	}
}
