package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Color3) }
func Color3() Reflector {
	return Reflector{
		Name:     "Color3",
		PushTo:   rbxmk.PushTypeTo("Color3"),
		PullFrom: rbxmk.PullTypeFrom("Color3"),
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "Color3").(types.Color3)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "Color3").(types.Color3)
				op := s.Pull(2, "Color3").(types.Color3)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: map[string]Member{
			"R": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.Color3).R))
			}},
			"G": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.Color3).G))
			}},
			"B": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.Color3).B))
			}},
			"Lerp": {Method: true, Get: func(s State, v types.Value) int {
				goal := s.Pull(2, "Color3").(types.Color3)
				alpha := float64(s.Pull(3, "number").(types.Double))
				return s.Push(v.(types.Color3).Lerp(goal, alpha))
			}},
			"ToHSV": {Method: true, Get: func(s State, v types.Value) int {
				hue, sat, val := v.(types.Color3).ToHSV()
				return s.Push(rtypes.Tuple{types.Double(hue), types.Double(sat), types.Double(val)})
			}},
		},
		Constructors: Constructors{
			"new": {
				Func: func(s State) int {
					var v types.Color3
					switch s.Count() {
					case 0:
					case 3:
						v.R = float32(s.Pull(1, "float").(types.Float))
						v.G = float32(s.Pull(2, "float").(types.Float))
						v.B = float32(s.Pull(3, "float").(types.Float))
					default:
						return s.RaiseError("expected 0 or 3 arguments")
					}
					return s.Push(v)
				},
			},
			"fromRGB": {
				Func: func(s State) int {
					return s.Push(types.NewColor3FromRGB(
						int(s.Pull(1, "int").(types.Int)),
						int(s.Pull(2, "int").(types.Int)),
						int(s.Pull(3, "int").(types.Int)),
					))
				},
			},
			"fromHSV": {
				Func: func(s State) int {
					return s.Push(types.NewColor3FromHSV(
						float64(s.Pull(1, "number").(types.Double)),
						float64(s.Pull(2, "number").(types.Double)),
						float64(s.Pull(3, "number").(types.Double)),
					))
				},
			},
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case types.Color3:
				return v
			case rtypes.Color3uint8:
				return types.Color3(v)
			}
			return nil
		},
	}
}
