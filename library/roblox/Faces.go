package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
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
func Faces() Reflector {
	return Reflector{
		Name:     "Faces",
		PushTo:   rbxmk.PushTypeTo("Faces"),
		PullFrom: rbxmk.PullTypeFrom("Faces"),
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "Faces").(types.Faces)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "Faces").(types.Faces)
				op := s.Pull(2, "Faces").(types.Faces)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: map[string]Member{
			"Right": {Get: func(s State, v types.Value) int {
				return s.Push(types.Bool(v.(types.Faces).Right))
			}},
			"Top": {Get: func(s State, v types.Value) int {
				return s.Push(types.Bool(v.(types.Faces).Top))
			}},
			"Back": {Get: func(s State, v types.Value) int {
				return s.Push(types.Bool(v.(types.Faces).Back))
			}},
			"Left": {Get: func(s State, v types.Value) int {
				return s.Push(types.Bool(v.(types.Faces).Left))
			}},
			"Bottom": {Get: func(s State, v types.Value) int {
				return s.Push(types.Bool(v.(types.Faces).Bottom))
			}},
			"Front": {Get: func(s State, v types.Value) int {
				return s.Push(types.Bool(v.(types.Faces).Front))
			}},
		},
		Constructors: Constructors{
			"new": {
				Func: func(s State) int {
					var v types.Faces
					n := s.L.GetTop()
					for i := 1; i <= n; i++ {
						switch value := PullVariant(s, i).(type) {
						case *rtypes.EnumItem:
							if enum := value.Enum(); enum != nil {
								switch enum.Name() {
								case "NormalId":
									setFacesFromNormalIdName(&v, value.Name())
								case "Axis":
									setFacesFromAxisName(&v, value.Name())
								}
							}
						case types.Intlike:
							setFacesFromNormalIdValue(&v, int(value.Intlike()))
						case types.Numberlike:
							setFacesFromNormalIdValue(&v, int(value.Numberlike()))
						case types.Stringlike:
							setFacesFromNormalIdName(&v, value.Stringlike())
						}
					}
					return s.Push(v)
				},
			},
			"fromComponents": {
				Func: func(s State) int {
					var v types.Faces
					switch s.Count() {
					case 6:
						v.Right = bool(s.Pull(1, "bool").(types.Bool))
						v.Top = bool(s.Pull(2, "bool").(types.Bool))
						v.Back = bool(s.Pull(3, "bool").(types.Bool))
						v.Left = bool(s.Pull(4, "bool").(types.Bool))
						v.Bottom = bool(s.Pull(5, "bool").(types.Bool))
						v.Front = bool(s.Pull(6, "bool").(types.Bool))
					default:
						return s.RaiseError("expected 0 or 6 arguments")
					}
					return s.Push(v)
				},
			},
		},
	}
}
