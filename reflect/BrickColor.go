package reflect

import (
	"math/rand"

	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func BrickColor() Type {
	return Type{
		Name:        "BrickColor",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Serialize: func(s State, v Value) (sv rbxfile.Value, err error) {
			if v, ok := v.(types.BrickColor); ok {
				return rbxfile.ValueBrickColor(v), nil
			}
			return nil, TypeError(nil, 0, "BrickColor")
		},
		Deserialize: func(s State, sv rbxfile.Value) (v Value, err error) {
			if sv, ok := sv.(rbxfile.ValueBrickColor); ok {
				return types.BrickColor(sv), nil
			}
			return nil, TypeError(nil, 0, "BrickColor")
		},
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				s.L.Push(lua.LString(v.(types.BrickColor).String()))
				return 1
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "BrickColor").(types.BrickColor)
				return s.Push("bool", v.(types.BrickColor) == op)
			},
		},
		Members: map[string]Member{
			"Name": {Get: func(s State, v Value) int {
				return s.Push("string", v.(types.BrickColor).Name())
			}},
			"Number": {Get: func(s State, v Value) int {
				return s.Push("int", v.(types.BrickColor).Number())
			}},
			"R": {Get: func(s State, v Value) int {
				return s.Push("number", v.(types.BrickColor).R())
			}},
			"G": {Get: func(s State, v Value) int {
				return s.Push("number", v.(types.BrickColor).G())
			}},
			"B": {Get: func(s State, v Value) int {
				return s.Push("number", v.(types.BrickColor).B())
			}},
			"Color": {Get: func(s State, v Value) int {
				return s.Push("Color3", v.(types.BrickColor).Color())
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.BrickColor
				switch s.Count() {
				case 1:
					switch arg := s.PullAnyOf(1, "int", "string", "Color3").(type) {
					case int:
						v = types.NewBrickColor(arg)
					case string:
						v = types.NewBrickColorFromName(arg)
					case types.Color3:
						v = types.NewBrickColorFromColor3(arg)
					}
				case 3:
					v = types.NewBrickColorFromColor(
						s.Pull(1, "double").(float64),
						s.Pull(2, "double").(float64),
						s.Pull(3, "double").(float64),
					)
				default:
					s.L.RaiseError("expected 1 or 3 arguments")
					return 0
				}
				return s.Push("BrickColor", v)
			},
			"palette": func(s State) int {
				index := s.Pull(1, "int").(int)
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
