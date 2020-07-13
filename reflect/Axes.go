package reflect

import (
	"strconv"

	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func Axes() Type {
	return Type{
		Name:        "Axes",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Serialize: func(s State, v Value) (sv rbxfile.Value, err error) {
			if v, ok := v.(types.Axes); ok {
				return rbxfile.ValueAxes(v), nil
			}
			return nil, TypeError(nil, 0, "Axes")
		},
		Deserialize: func(s State, sv rbxfile.Value) (v Value, err error) {
			if sv, ok := sv.(rbxfile.ValueAxes); ok {
				return types.Axes(sv), nil
			}
			return nil, TypeError(nil, 0, "Axes")
		},
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				u := v.(types.Axes)
				var b string
				b += "X:" + strconv.FormatBool(u.X) + ", "
				b += "Y:" + strconv.FormatBool(u.Y) + ", "
				b += "Z:" + strconv.FormatBool(u.Z)
				s.L.Push(lua.LString(b))
				return 1
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "Axes").(types.Axes)
				return s.Push("bool", v.(types.Axes) == op)
			},
		},
		Members: map[string]Member{
			"X": {Get: func(s State, v Value) int {
				return s.Push("bool", v.(types.Axes).X)
			}},
			"Y": {Get: func(s State, v Value) int {
				return s.Push("bool", v.(types.Axes).Y)
			}},
			"Z": {Get: func(s State, v Value) int {
				return s.Push("bool", v.(types.Axes).Z)
			}},
		},
		Constructors: Constructors{
			// TODO: match API.
			"new": func(s State) int {
				var v types.Axes
				switch s.Count() {
				case 3:
					v.X = s.Pull(1, "bool").(bool)
					v.Y = s.Pull(2, "bool").(bool)
					v.Z = s.Pull(3, "bool").(bool)
				default:
					s.L.RaiseError("expected 0 or 3 arguments")
					return 0
				}
				return s.Push("Axes", v)
			},
		},
	}
}
