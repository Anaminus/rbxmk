package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func setAxesFromAxisName(axes *types.Axes, name string) {
	switch name {
	case "X":
		axes.X = true
	case "Y":
		axes.Y = true
	case "Z":
		axes.Z = true
	}
}

func setAxesFromAxisValue(axes *types.Axes, value int) {
	switch value {
	case 0:
		axes.X = true
	case 1:
		axes.Y = true
	case 2:
		axes.Z = true
	}
}

func setAxesFromNormalIdName(axes *types.Axes, name string) {
	switch name {
	case "Right", "Left":
		axes.X = true
	case "Top", "Bottom":
		axes.Y = true
	case "Back", "Front":
		axes.Z = true
	}
}

func init() { register(Axes) }
func Axes() Reflector {
	return Reflector{
		Name:     "Axes",
		PushTo:   rbxmk.PushTypeTo("Axes"),
		PullFrom: rbxmk.PullTypeFrom("Axes"),
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "Axes").(types.Axes)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "Axes").(types.Axes)
				op := s.Pull(2, "Axes").(types.Axes)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: map[string]Member{
			"X": {Get: func(s State, v types.Value) int {
				return s.Push(types.Bool(v.(types.Axes).X))
			}},
			"Y": {Get: func(s State, v types.Value) int {
				return s.Push(types.Bool(v.(types.Axes).Y))
			}},
			"Z": {Get: func(s State, v types.Value) int {
				return s.Push(types.Bool(v.(types.Axes).Z))
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.Axes
				n := s.L.GetTop()
				for i := 1; i <= n; i++ {
					switch value := PullVariant(s, i).(type) {
					case *rtypes.EnumItem:
						if enum := value.Enum(); enum != nil {
							switch enum.Name() {
							case "Axis":
								setAxesFromAxisName(&v, value.Name())
							case "NormalId":
								setAxesFromNormalIdName(&v, value.Name())
							}
						}
					case types.Intlike:
						setAxesFromAxisValue(&v, int(value.Intlike()))
					case types.Numberlike:
						setAxesFromAxisValue(&v, int(value.Numberlike()))
					case types.Stringlike:
						setAxesFromNormalIdName(&v, value.Stringlike())
					}
				}
				return s.Push(v)
			},
			"fromComponents": func(s State) int {
				var v types.Axes
				switch s.Count() {
				case 3:
					v.X = bool(s.Pull(1, "bool").(types.Bool))
					v.Y = bool(s.Pull(2, "bool").(types.Bool))
					v.Z = bool(s.Pull(3, "bool").(types.Bool))
				default:
					return s.RaiseError("expected 0 or 3 arguments")
				}
				return s.Push(v)
			},
		},
	}
}
