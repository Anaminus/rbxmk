package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Color3() Type {
	return Type{
		Name:        "Color3",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Serialize: func(s State, v Value) (sv rbxfile.Value, err error) {
			if v, ok := v.(types.Color3); ok {
				return rbxfile.ValueColor3(v), nil
			}
			return nil, TypeError(nil, 0, "Color3")
		},
		Deserialize: func(s State, sv rbxfile.Value) (v Value, err error) {
			if sv, ok := sv.(rbxfile.ValueColor3); ok {
				return types.Color3(sv), nil
			}
			return nil, TypeError(nil, 0, "Color3")
		},
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				s.L.Push(lua.LString(v.(types.Color3).String()))
				return 1
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "Color3").(types.Color3)
				return s.Push("bool", v.(types.Color3) == op)
			},
		},
		Members: map[string]Member{
			"R": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.Color3).R)
			}},
			"G": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.Color3).G)
			}},
			"B": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.Color3).B)
			}},
			"Lerp": {Method: true, Get: func(s State, v Value) int {
				goal := s.Pull(2, "Color3").(types.Color3)
				alpha := s.Pull(3, "double").(float64)
				return s.Push("Color3", v.(types.Color3).Lerp(goal, alpha))
			}},
			"ToHSV": {Method: true, Get: func(s State, v Value) int {
				hue, sat, val := v.(types.Color3).ToHSV()
				return s.Push("Tuple", []Value{hue, sat, val})
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.Color3
				switch s.Count() {
				case 0:
				case 3:
					v.R = s.Pull(1, "float").(float32)
					v.G = s.Pull(2, "float").(float32)
					v.B = s.Pull(3, "float").(float32)
				default:
					s.L.RaiseError("expected 0 or 3 arguments")
					return 0
				}
				return s.Push("Color3", v)
			},
			"fromRGB": func(s State) int {
				return s.Push("Color3", types.NewColor3FromRGB(
					s.Pull(1, "int").(int),
					s.Pull(2, "int").(int),
					s.Pull(3, "int").(int),
				))
			},
			"fromHSV": func(s State) int {
				return s.Push("Color3", types.NewColor3FromHSV(
					s.Pull(1, "number").(float64),
					s.Pull(2, "number").(float64),
					s.Pull(3, "number").(float64),
				))
			},
		},
	}
}

// TODO: Serialization should be separated from reflection. Color3uint8 is not a
// reflected type, but a serialized type.

func Color3uint8() Type {
	t := Color3()
	t.Name = "Color3uint8"
	t.Serialize = func(s State, v Value) (sv rbxfile.Value, err error) {
		if v, ok := v.(types.Color3); ok {
			return rbxfile.ValueColor3uint8{
				R: byte(v.R * 255),
				G: byte(v.G * 255),
				B: byte(v.B * 255),
			}, nil
		}
		return nil, TypeError(nil, 0, "Color3uint8")
	}
	t.Deserialize = func(s State, sv rbxfile.Value) (v Value, err error) {
		if sv, ok := sv.(rbxfile.ValueColor3uint8); ok {
			return types.Color3{
				R: float32(sv.R) / 255,
				G: float32(sv.G) / 255,
				B: float32(sv.B) / 255,
			}, nil
		}
		return nil, TypeError(nil, 0, "Color3uint8")
	}
	return t
}
