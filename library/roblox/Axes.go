package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
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
func Axes() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "Axes",
		PushTo:   rbxmk.PushTypeTo("Axes"),
		PullFrom: rbxmk.PullTypeFrom("Axes"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "Axes").(types.Axes)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "Axes").(types.Axes)
				op := s.Pull(2, "Axes").(types.Axes)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"X": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).X))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("bool"),
						ReadOnly:    true,
						Summary:     "libraries/roblox/types/Axes:Properties/X/Summary",
						Description: "libraries/roblox/types/Axes:Properties/X/Description",
					}
				},
			},
			"Y": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).Y))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("bool"),
						ReadOnly:    true,
						Summary:     "libraries/roblox/types/Axes:Properties/Y/Summary",
						Description: "libraries/roblox/types/Axes:Properties/Y/Description",
					}
				},
			},
			"Z": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).Z))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("bool"),
						ReadOnly:    true,
						Summary:     "libraries/roblox/types/Axes:Properties/Z/Summary",
						Description: "libraries/roblox/types/Axes:Properties/Z/Description",
					}
				},
			},
			"Right": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).X))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("bool"),
						ReadOnly:    true,
						Summary:     "libraries/roblox/types/Axes:Properties/Right/Summary",
						Description: "libraries/roblox/types/Axes:Properties/Right/Description",
					}
				},
			},
			"Top": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).Y))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("bool"),
						ReadOnly:    true,
						Summary:     "libraries/roblox/types/Axes:Properties/Top/Summary",
						Description: "libraries/roblox/types/Axes:Properties/Top/Description",
					}
				},
			},
			"Back": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).Z))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("bool"),
						ReadOnly:    true,
						Summary:     "libraries/roblox/types/Axes:Properties/Back/Summary",
						Description: "libraries/roblox/types/Axes:Properties/Back/Description",
					}
				},
			},
			"Left": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).X))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("bool"),
						ReadOnly:    true,
						Summary:     "libraries/roblox/types/Axes:Properties/Left/Summary",
						Description: "libraries/roblox/types/Axes:Properties/Left/Description",
					}
				},
			},
			"Bottom": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).Y))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("bool"),
						ReadOnly:    true,
						Summary:     "libraries/roblox/types/Axes:Properties/Bottom/Summary",
						Description: "libraries/roblox/types/Axes:Properties/Bottom/Description",
					}
				},
			},
			"Front": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).Z))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("bool"),
						ReadOnly:    true,
						Summary:     "libraries/roblox/types/Axes:Properties/Front/Summary",
						Description: "libraries/roblox/types/Axes:Properties/Front/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.Axes
					n := s.Count()
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
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "...", Type: dt.Optional{T: dt.Prim("any")}},
							},
							Summary:     "libraries/roblox/types/Axes:Constructors/new/Summary",
							Description: "libraries/roblox/types/Axes:Constructors/new/Description",
						},
					}
				},
			},
			"fromComponents": {
				Func: func(s rbxmk.State) int {
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
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("bool")},
								{Name: "y", Type: dt.Prim("bool")},
								{Name: "z", Type: dt.Prim("bool")},
							},
							Summary:     "libraries/roblox/types/Axes:Constructors/fromComponents/Summary",
							Description: "libraries/roblox/types/Axes:Constructors/fromComponents/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators:   &dump.Operators{Eq: true},
				Summary:     "libraries/roblox/types/Axes:Summary",
				Description: "libraries/roblox/types/Axes:Description",
			}
		},
	}
}
