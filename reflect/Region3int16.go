package reflect

import (
	"strconv"

	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Region3int16() Type {
	return Type{
		Name:        "Region3int16",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				u := v.(types.Region3int16)
				var b string
				b += strconv.FormatFloat(float64(u.Min.X), 'g', -1, 32) + ", "
				b += strconv.FormatFloat(float64(u.Min.Y), 'g', -1, 32) + ", "
				b += strconv.FormatFloat(float64(u.Min.Z), 'g', -1, 32) + ", "
				b += strconv.FormatFloat(float64(u.Max.X), 'g', -1, 32) + ", "
				b += strconv.FormatFloat(float64(u.Max.Y), 'g', -1, 32) + ", "
				b += strconv.FormatFloat(float64(u.Max.Z), 'g', -1, 32)
				s.L.Push(lua.LString(b))
				return 1
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "Region3int16").(types.Region3int16)
				return s.Push("bool", v.(types.Region3int16) == op)
			},
		},
		Members: map[string]Member{
			"Min": {Get: func(s State, v Value) int {
				return s.Push("Vector3int16", v.(types.Region3int16).Min)
			}},
			"Max": {Get: func(s State, v Value) int {
				return s.Push("Vector3int16", v.(types.Region3int16).Max)
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				return s.Push("Region3int16", types.Region3int16{
					Min: s.Pull(1, "Vector3int16").(types.Vector3int16),
					Max: s.Pull(2, "Vector3int16").(types.Vector3int16),
				})
			},
		},
	}
}
