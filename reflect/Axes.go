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
		Name:     rtypes.T_Axes,
		PushTo:   rbxmk.PushTypeTo(rtypes.T_Axes),
		PullFrom: rbxmk.PullTypeFrom(rtypes.T_Axes),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.Axes:
				*p = v.(types.Axes)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Axes).(types.Axes)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Axes).(types.Axes)
				op := s.Pull(2, rtypes.T_Axes).(types.Axes)
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
						ValueType:   dt.Prim(rtypes.T_Bool),
						ReadOnly:    true,
						Summary:     "Types/Axes:Properties/X/Summary",
						Description: "Types/Axes:Properties/X/Description",
					}
				},
			},
			"Y": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).Y))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Bool),
						ReadOnly:    true,
						Summary:     "Types/Axes:Properties/Y/Summary",
						Description: "Types/Axes:Properties/Y/Description",
					}
				},
			},
			"Z": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).Z))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Bool),
						ReadOnly:    true,
						Summary:     "Types/Axes:Properties/Z/Summary",
						Description: "Types/Axes:Properties/Z/Description",
					}
				},
			},
			"Right": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).X))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Bool),
						ReadOnly:    true,
						Summary:     "Types/Axes:Properties/Right/Summary",
						Description: "Types/Axes:Properties/Right/Description",
					}
				},
			},
			"Top": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).Y))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Bool),
						ReadOnly:    true,
						Summary:     "Types/Axes:Properties/Top/Summary",
						Description: "Types/Axes:Properties/Top/Description",
					}
				},
			},
			"Back": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).Z))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Bool),
						ReadOnly:    true,
						Summary:     "Types/Axes:Properties/Back/Summary",
						Description: "Types/Axes:Properties/Back/Description",
					}
				},
			},
			"Left": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).X))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Bool),
						ReadOnly:    true,
						Summary:     "Types/Axes:Properties/Left/Summary",
						Description: "Types/Axes:Properties/Left/Description",
					}
				},
			},
			"Bottom": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).Y))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Bool),
						ReadOnly:    true,
						Summary:     "Types/Axes:Properties/Bottom/Summary",
						Description: "Types/Axes:Properties/Bottom/Description",
					}
				},
			},
			"Front": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Axes).Z))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Bool),
						ReadOnly:    true,
						Summary:     "Types/Axes:Properties/Front/Summary",
						Description: "Types/Axes:Properties/Front/Description",
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
						switch value := s.PullAnyOf(i, rtypes.T_EnumItem, rtypes.T_String, rtypes.T_Int).(type) {
						case *rtypes.EnumItem:
							if enum := value.Enum(); enum != nil {
								switch enum.Name() {
								case "Axis":
									setAxesFromAxisName(&v, value.Name())
								case "NormalId":
									setAxesFromNormalIdName(&v, value.Name())
								}
							}
						case types.Int:
							setAxesFromAxisValue(&v, int(value))
						case types.String:
							setAxesFromNormalIdName(&v, string(value))
						}
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "...", Type: dt.Or{
									dt.Prim(rtypes.T_EnumItem),
									dt.Prim(rtypes.T_String),
									dt.Prim(rtypes.T_Int),
								}},
							},
							Summary:     "Types/Axes:Constructors/new/Summary",
							Description: "Types/Axes:Constructors/new/Description",
						},
					}
				},
			},
			"fromComponents": {
				Func: func(s rbxmk.State) int {
					var v types.Axes
					switch s.Count() {
					case 3:
						v.X = bool(s.Pull(1, rtypes.T_Bool).(types.Bool))
						v.Y = bool(s.Pull(2, rtypes.T_Bool).(types.Bool))
						v.Z = bool(s.Pull(3, rtypes.T_Bool).(types.Bool))
					default:
						return s.RaiseError("expected 0 or 3 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(rtypes.T_Bool)},
								{Name: "y", Type: dt.Prim(rtypes.T_Bool)},
								{Name: "z", Type: dt.Prim(rtypes.T_Bool)},
							},
							Summary:     "Types/Axes:Constructors/fromComponents/Summary",
							Description: "Types/Axes:Constructors/fromComponents/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category: "roblox",
				Operators: &dump.Operators{
					Eq: &dump.Cmpop{
						Summary:     "Types/Axes:Operators/Eq/Summary",
						Description: "Types/Axes:Operators/Eq/Description",
					},
				},
				Summary:     "Types/Axes:Summary",
				Description: "Types/Axes:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Bool,
			Variant,
		},
	}
}
