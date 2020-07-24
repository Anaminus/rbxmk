package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Color3() Type {
	return Type{
		Name:        "Color3",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				s.L.Push(lua.LString(s.Pull(1, "Color3").(types.Color3).String()))
				return 1
			},
			"__eq": func(s State) int {
				op := s.Pull(2, "Color3").(types.Color3)
				return s.Push("bool", types.Bool(s.Pull(1, "Color3").(types.Color3) == op))
			},
		},
		Members: map[string]Member{
			"R": {Get: func(s State, v types.Value) int {
				return s.Push("float", types.Float(v.(types.Color3).R))
			}},
			"G": {Get: func(s State, v types.Value) int {
				return s.Push("float", types.Float(v.(types.Color3).G))
			}},
			"B": {Get: func(s State, v types.Value) int {
				return s.Push("float", types.Float(v.(types.Color3).B))
			}},
			"Lerp": {Method: true, Get: func(s State, v types.Value) int {
				goal := s.Pull(2, "Color3").(types.Color3)
				alpha := float64(s.Pull(3, "number").(types.Double))
				return s.Push("Color3", v.(types.Color3).Lerp(goal, alpha))
			}},
			"ToHSV": {Method: true, Get: func(s State, v types.Value) int {
				hue, sat, val := v.(types.Color3).ToHSV()
				return s.Push("Tuple", rtypes.Tuple{types.Double(hue), types.Double(sat), types.Double(val)})
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.Color3
				switch s.Count() {
				case 0:
				case 3:
					v.R = float32(s.Pull(1, "float").(types.Float))
					v.G = float32(s.Pull(2, "float").(types.Float))
					v.B = float32(s.Pull(3, "float").(types.Float))
				default:
					s.L.RaiseError("expected 0 or 3 arguments")
					return 0
				}
				return s.Push("Color3", v)
			},
			"fromRGB": func(s State) int {
				return s.Push("Color3", types.NewColor3FromRGB(
					int(s.Pull(1, "int").(types.Int)),
					int(s.Pull(2, "int").(types.Int)),
					int(s.Pull(3, "int").(types.Int)),
				))
			},
			"fromHSV": func(s State) int {
				return s.Push("Color3", types.NewColor3FromHSV(
					float64(s.Pull(1, "number").(types.Double)),
					float64(s.Pull(2, "number").(types.Double)),
					float64(s.Pull(3, "number").(types.Double)),
				))
			},
		},
	}
}

func Color3uint8() Type {
	t := Color3()
	t.Name = "Color3uint8"
	return t
}
