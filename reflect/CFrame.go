package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func CFrame() Type {
	return Type{
		Name:        "CFrame",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Serialize: func(s State, v Value) (sv rbxfile.Value, err error) {
			if v, ok := v.(types.CFrame); ok {
				return rbxfile.ValueCFrame{
					Position: rbxfile.ValueVector3(v.Position),
					Rotation: v.Rotation,
				}, nil
			}
			return nil, TypeError(nil, 0, "CFrame")
		},
		Deserialize: func(s State, sv rbxfile.Value) (v Value, err error) {
			if sv, ok := sv.(rbxfile.ValueCFrame); ok {
				return types.CFrame{
					Position: types.Vector3(sv.Position),
					Rotation: sv.Rotation,
				}, nil
			}
			return nil, TypeError(nil, 0, "CFrame")
		},
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				s.L.Push(lua.LString(v.(types.CFrame).String()))
				return 1
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "CFrame").(types.CFrame)
				return s.Push("bool", v.(types.CFrame) == op)
			},
			"__add": func(s State, v Value) int {
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push("CFrame", v.(types.CFrame).AddVec(op))
			},
			"__sub": func(s State, v Value) int {
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push("CFrame", v.(types.CFrame).SubVec(op))
			},
			"__mul": func(s State, v Value) int {
				switch op := s.PullAnyOf(2, "CFrame", "Vector3").(type) {
				case types.CFrame:
					return s.Push("CFrame", v.(types.CFrame).Mul(op))
				case types.Vector3:
					return s.Push("Vector3", v.(types.CFrame).MulVec(op))
				default:
					s.L.ArgError(2, "attempt to multiply a CFrame with an incompatible value type or nil")
					return 0
				}
			},
		},
		Members: map[string]Member{
			"P": {Get: func(s State, v Value) int {
				return s.Push("Vector3", v.(types.CFrame).Position)
			}},
			"Position": {Get: func(s State, v Value) int {
				return s.Push("Vector3", v.(types.CFrame).Position)
			}},
			"X": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.CFrame).X())
			}},
			"Y": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.CFrame).Y())
			}},
			"Z": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.CFrame).Z())
			}},
			"LookVector": {Get: func(s State, v Value) int {
				cf := v.(types.CFrame)
				return s.Push("Vector3", types.Vector3{
					X: -cf.Rotation[2],
					Y: -cf.Rotation[5],
					Z: -cf.Rotation[8],
				})
			}},
			"RightVector": {Get: func(s State, v Value) int {
				cf := v.(types.CFrame)
				return s.Push("Vector3", types.Vector3{
					X: -cf.Rotation[0],
					Y: -cf.Rotation[3],
					Z: -cf.Rotation[6],
				})
			}},
			"UpVector": {Get: func(s State, v Value) int {
				cf := v.(types.CFrame)
				return s.Push("Vector3", types.Vector3{
					X: -cf.Rotation[1],
					Y: -cf.Rotation[4],
					Z: -cf.Rotation[7],
				})
			}},
			"XVector": {Get: func(s State, v Value) int {
				cf := v.(types.CFrame)
				return s.Push("Vector3", types.Vector3{
					X: -cf.Rotation[0],
					Y: -cf.Rotation[3],
					Z: -cf.Rotation[6],
				})
			}},
			"YVector": {Get: func(s State, v Value) int {
				cf := v.(types.CFrame)
				return s.Push("Vector3", types.Vector3{
					X: -cf.Rotation[1],
					Y: -cf.Rotation[4],
					Z: -cf.Rotation[7],
				})
			}},
			"ZVector": {Get: func(s State, v Value) int {
				cf := v.(types.CFrame)
				return s.Push("Vector3", types.Vector3{
					X: -cf.Rotation[2],
					Y: -cf.Rotation[5],
					Z: -cf.Rotation[8],
				})
			}},
			"Inverse": {Method: true, Get: func(s State, v Value) int {
				return s.Push("float", v.(types.CFrame).Inverse())
			}},
			"Lerp": {Method: true, Get: func(s State, v Value) int {
				goal := s.Pull(2, "CFrame").(types.CFrame)
				alpha := s.Pull(3, "double").(float64)
				return s.Push("CFrame", v.(types.CFrame).Lerp(goal, alpha))
			}},
			"ToWorldSpace": {Method: true, Get: func(s State, v Value) int {
				cf := s.Pull(2, "CFrame").(types.CFrame)
				return s.Push("CFrame", v.(types.CFrame).ToWorldSpace(cf))
			}},
			"ToObjectSpace": {Method: true, Get: func(s State, v Value) int {
				cf := s.Pull(2, "CFrame").(types.CFrame)
				return s.Push("CFrame", v.(types.CFrame).ToObjectSpace(cf))
			}},
			"PointToWorldSpace": {Method: true, Get: func(s State, v Value) int {
				v3 := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push("Vector3", v.(types.CFrame).PointToWorldSpace(v3))
			}},
			"PointToObjectSpace": {Method: true, Get: func(s State, v Value) int {
				v3 := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push("Vector3", v.(types.CFrame).PointToObjectSpace(v3))
			}},
			"VectorToWorldSpace": {Method: true, Get: func(s State, v Value) int {
				v3 := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push("Vector3", v.(types.CFrame).VectorToWorldSpace(v3))
			}},
			"VectorToObjectSpace": {Method: true, Get: func(s State, v Value) int {
				v3 := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push("Vector3", v.(types.CFrame).VectorToObjectSpace(v3))
			}},
			"GetComponents": {Method: true, Get: func(s State, v Value) int {
				cf := v.(types.CFrame)
				return s.Push("Tuple", []Value{
					cf.Position.X,
					cf.Position.Y,
					cf.Position.Z,
					cf.Rotation[0],
					cf.Rotation[1],
					cf.Rotation[2],
					cf.Rotation[3],
					cf.Rotation[4],
					cf.Rotation[5],
					cf.Rotation[6],
					cf.Rotation[7],
					cf.Rotation[8],
				})
			}},
			"ToEulerAnglesXYZ": {Method: true, Get: func(s State, v Value) int {
				x, y, z := v.(types.CFrame).Angles()
				return s.Push("Tuple", []Value{x, y, z})
			}},
			"ToEulerAnglesYXZ": {Method: true, Get: func(s State, v Value) int {
				x, y, z := v.(types.CFrame).Orientation()
				return s.Push("Tuple", []Value{x, y, z})
			}},
			"ToOrientation": {Method: true, Get: func(s State, v Value) int {
				x, y, z := v.(types.CFrame).Orientation()
				return s.Push("Tuple", []Value{x, y, z})
			}},
			"ToAxisAngle": {Method: true, Get: func(s State, v Value) int {
				axis, rotation := v.(types.CFrame).AxisAngle()
				return s.Push("Tuple", []Value{TValue{Type: "Vector3", Value: axis}, rotation})
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.CFrame
				switch s.Count() {
				case 0:
					v = types.NewCFrame()
				case 1:
					pos := s.Pull(1, "Vector3").(types.Vector3)
					v = types.NewCFrameFromVector3(pos)
				case 2:
					pos := s.Pull(1, "Vector3").(types.Vector3)
					lookAt := s.Pull(2, "Vector3").(types.Vector3)
					v = types.NewCFrameFromLook(pos, lookAt)
				case 3:
					v = types.NewCFrameFromPosition(
						s.Pull(1, "number").(float64),
						s.Pull(2, "number").(float64),
						s.Pull(3, "number").(float64),
					)
				case 7:
					v = types.NewCFrameFromQuat(
						s.Pull(1, "number").(float64),
						s.Pull(2, "number").(float64),
						s.Pull(3, "number").(float64),
						s.Pull(4, "number").(float64),
						s.Pull(5, "number").(float64),
						s.Pull(6, "number").(float64),
						s.Pull(7, "number").(float64),
					)
				case 12:
					v = types.NewCFrameFromComponents(
						s.Pull(1, "number").(float64),
						s.Pull(2, "number").(float64),
						s.Pull(3, "number").(float64),
						s.Pull(4, "number").(float64),
						s.Pull(5, "number").(float64),
						s.Pull(6, "number").(float64),
						s.Pull(7, "number").(float64),
						s.Pull(8, "number").(float64),
						s.Pull(9, "number").(float64),
						s.Pull(10, "number").(float64),
						s.Pull(11, "number").(float64),
						s.Pull(12, "number").(float64),
					)
				default:
					s.L.RaiseError("unexpected number of arguments")
					return 0
				}
				return s.Push("CFrame", v)
			},
			"fromEulerAnglesXYZ": func(s State) int {
				return s.Push("CFrame", types.NewCFrameFromAngles(
					s.Pull(1, "number").(float64),
					s.Pull(2, "number").(float64),
					s.Pull(3, "number").(float64),
				))
			},
			"fromEulerAnglesYXZ": func(s State) int {
				return s.Push("CFrame", types.NewCFrameFromOrientation(
					s.Pull(1, "number").(float64),
					s.Pull(2, "number").(float64),
					s.Pull(3, "number").(float64),
				))
			},
			"Angles": func(s State) int {
				return s.Push("CFrame", types.NewCFrameFromAngles(
					s.Pull(1, "number").(float64),
					s.Pull(2, "number").(float64),
					s.Pull(3, "number").(float64),
				))
			},
			"fromOrientation": func(s State) int {
				return s.Push("CFrame", types.NewCFrameFromOrientation(
					s.Pull(1, "number").(float64),
					s.Pull(2, "number").(float64),
					s.Pull(3, "number").(float64),
				))
			},
			"fromAxisAngle": func(s State) int {
				return s.Push("CFrame", types.NewCFrameFromAxisAngle(
					s.Pull(1, "Vector3").(types.Vector3),
					s.Pull(2, "number").(float64),
				))
			},
			"fromMatrix": func(s State) int {
				return s.Push("CFrame", types.NewCFrameFromMatrix(
					s.Pull(1, "Vector3").(types.Vector3),
					s.Pull(2, "Vector3").(types.Vector3),
					s.Pull(3, "Vector3").(types.Vector3),
					s.Pull(4, "Vector3").(types.Vector3),
				))
			},
		},
	}
}
