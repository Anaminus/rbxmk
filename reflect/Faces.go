package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func setFacesFromNormalIdName(faces *types.Faces, name string) {
	switch name {
	case "Right":
		faces.Right = true
	case "Left":
		faces.Left = true
	case "Top":
		faces.Top = true
	case "Bottom":
		faces.Bottom = true
	case "Back":
		faces.Back = true
	case "Front":
		faces.Front = true
	}
}

func setFacesFromNormalIdValue(faces *types.Faces, value int) {
	switch value {
	case 0:
		faces.Right = true
	case 1:
		faces.Left = true
	case 2:
		faces.Top = true
	case 3:
		faces.Bottom = true
	case 4:
		faces.Back = true
	case 5:
		faces.Front = true
	}
}

func setFacesFromAxisName(faces *types.Faces, name string) {
	switch name {
	case "X":
		faces.Right = true
		faces.Left = true
	case "Y":
		faces.Top = true
		faces.Bottom = true
	case "Z":
		faces.Back = true
		faces.Front = true
	}
}

func init() { register(Faces) }
func Faces() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     rtypes.T_Faces,
		PushTo:   rbxmk.PushTypeTo(rtypes.T_Faces),
		PullFrom: rbxmk.PullTypeFrom(rtypes.T_Faces),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.Faces:
				*p = v.(types.Faces)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Faces).(types.Faces)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Faces).(types.Faces)
				op := s.Pull(2, rtypes.T_Faces).(types.Faces)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"Right": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Faces).Right))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Bool),
						ReadOnly:    true,
						Summary:     "Types/Faces:Properties/Right/Summary",
						Description: "Types/Faces:Properties/Right/Description",
					}
				},
			},
			"Top": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Faces).Top))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Bool),
						ReadOnly:    true,
						Summary:     "Types/Faces:Properties/Top/Summary",
						Description: "Types/Faces:Properties/Top/Description",
					}
				},
			},
			"Back": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Faces).Back))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Bool),
						ReadOnly:    true,
						Summary:     "Types/Faces:Properties/Back/Summary",
						Description: "Types/Faces:Properties/Back/Description",
					}
				},
			},
			"Left": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Faces).Left))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Bool),
						ReadOnly:    true,
						Summary:     "Types/Faces:Properties/Left/Summary",
						Description: "Types/Faces:Properties/Left/Description",
					}
				},
			},
			"Bottom": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Faces).Bottom))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Bool),
						ReadOnly:    true,
						Summary:     "Types/Faces:Properties/Bottom/Summary",
						Description: "Types/Faces:Properties/Bottom/Description",
					}
				},
			},
			"Front": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Bool(v.(types.Faces).Front))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Bool),
						ReadOnly:    true,
						Summary:     "Types/Faces:Properties/Front/Summary",
						Description: "Types/Faces:Properties/Front/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.Faces
					n := s.Count()
					for i := 1; i <= n; i++ {
						switch value := s.PullAnyOf(i, rtypes.T_EnumItem, rtypes.T_String, rtypes.T_Int).(type) {
						case *rtypes.EnumItem:
							if enum := value.Enum(); enum != nil {
								switch enum.Name() {
								case "NormalId":
									setFacesFromNormalIdName(&v, value.Name())
								case "Axis":
									setFacesFromAxisName(&v, value.Name())
								}
							}
						case types.Int:
							setFacesFromNormalIdValue(&v, int(value))
						case types.String:
							setFacesFromNormalIdName(&v, string(value))
						}
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "...", Type: dt.Or(
									dt.Prim(rtypes.T_EnumItem),
									dt.Prim(rtypes.T_String),
									dt.Prim(rtypes.T_Int),
								)},
							},
							Summary:     "Types/Faces:Constructors/new/Summary",
							Description: "Types/Faces:Constructors/new/Description",
						},
					}
				},
			},
			"fromComponents": {
				Func: func(s rbxmk.State) int {
					var v types.Faces
					switch s.Count() {
					case 6:
						v.Right = bool(s.Pull(1, rtypes.T_Bool).(types.Bool))
						v.Top = bool(s.Pull(2, rtypes.T_Bool).(types.Bool))
						v.Back = bool(s.Pull(3, rtypes.T_Bool).(types.Bool))
						v.Left = bool(s.Pull(4, rtypes.T_Bool).(types.Bool))
						v.Bottom = bool(s.Pull(5, rtypes.T_Bool).(types.Bool))
						v.Front = bool(s.Pull(6, rtypes.T_Bool).(types.Bool))
					default:
						return s.RaiseError("expected 0 or 6 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "right", Type: dt.Prim(rtypes.T_Bool)},
								{Name: "top", Type: dt.Prim(rtypes.T_Bool)},
								{Name: "back", Type: dt.Prim(rtypes.T_Bool)},
								{Name: "left", Type: dt.Prim(rtypes.T_Bool)},
								{Name: "bottom", Type: dt.Prim(rtypes.T_Bool)},
								{Name: "front", Type: dt.Prim(rtypes.T_Bool)},
							},
							Summary:     "Types/Faces:Constructors/fromComponents/Summary",
							Description: "Types/Faces:Constructors/fromComponents/Description",
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
						Summary:     "Types/Faces:Operators/Eq/Summary",
						Description: "Types/Faces:Operators/Eq/Description",
					},
				},
				Summary:     "Types/Faces:Summary",
				Description: "Types/Faces:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Bool,
			Variant,
		},
	}
}
