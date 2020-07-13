package reflect

import (
	"strconv"

	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Region3() Type {
	return Type{
		Name:        "Region3",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				u := v.(types.Region3)
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
				op := s.Pull(2, "Region3").(types.Region3)
				return s.Push("bool", v.(types.Region3) == op)
			},
		},
		Members: map[string]Member{
			"CFrame": {Get: func(s State, v Value) int {
				return s.Push("CFrame", v.(types.Region3).CFrame())
			}},
			"Size": {Get: func(s State, v Value) int {
				return s.Push("Vector3", v.(types.Region3).Size())
			}},
			"ExpandToGrid": {Method: true, Get: func(s State, v Value) int {
				region := s.Pull(2, "int").(int)
				return s.Push("Region3", v.(types.Region3).ExpandToGrid(region))
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				return s.Push("Region3", types.Region3{
					Min: s.Pull(1, "Vector3").(types.Vector3),
					Max: s.Pull(2, "Vector3").(types.Vector3),
				})
			},
		},
	}
}
