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
		Members: map[string]rbxmk.Member{
			"X": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).X))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("bool"), ReadOnly: true} },
			},
			"Y": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).Y))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("bool"), ReadOnly: true} },
			},
			"Z": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).Z))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("bool"), ReadOnly: true} },
			},
			"Right": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).X))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("bool"), ReadOnly: true} },
			},
			"Top": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).Y))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("bool"), ReadOnly: true} },
			},
			"Back": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).Z))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("bool"), ReadOnly: true} },
			},
			"Left": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).X))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("bool"), ReadOnly: true} },
			},
			"Bottom": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).Y))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("bool"), ReadOnly: true} },
			},
			"Front": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).Z))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("bool"), ReadOnly: true} },
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
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
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Parameters: dump.Parameters{
							{Name: "...", Type: dt.Optional{T: dt.Prim("any")}},
						},
					}}
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
					return []dump.Function{{
						Parameters: dump.Parameters{
							{Name: "x", Type: dt.Prim("bool")},
							{Name: "y", Type: dt.Prim("bool")},
							{Name: "z", Type: dt.Prim("bool")},
						},
					}}
				},
			},
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
