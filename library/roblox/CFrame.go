package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func init() { register(CFrame) }
func CFrame() Reflector {
	return Reflector{
		Name:     "CFrame",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "CFrame").(types.CFrame)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "CFrame").(types.CFrame)
				op := s.Pull(2, "CFrame").(types.CFrame)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
			"__add": func(s State) int {
				v := s.Pull(1, "CFrame").(types.CFrame)
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push(v.AddVec(op))
			},
			"__sub": func(s State) int {
				v := s.Pull(1, "CFrame").(types.CFrame)
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push(v.SubVec(op))
			},
			"__mul": func(s State) int {
				v := s.Pull(1, "CFrame").(types.CFrame)
				switch op := s.PullAnyOf(2, "CFrame", "Vector3").(type) {
				case types.CFrame:
					return s.Push(v.Mul(op))
				case types.Vector3:
					return s.Push(v.MulVec(op))
				default:
					s.L.ArgError(2, "attempt to multiply a CFrame with an incompatible types.value type or nil")
					return 0
				}
			},
		},
		Members: map[string]Member{
			"P": {Get: func(s State, v types.Value) int {
				return s.Push(v.(types.CFrame).Position)
			}},
			"Position": {Get: func(s State, v types.Value) int {
				return s.Push(v.(types.CFrame).Position)
			}},
			"X": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.CFrame).X()))
			}},
			"Y": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.CFrame).Y()))
			}},
			"Z": {Get: func(s State, v types.Value) int {
				return s.Push(types.Float(v.(types.CFrame).Z()))
			}},
			"LookVector": {Get: func(s State, v types.Value) int {
				cf := v.(types.CFrame)
				return s.Push(types.Vector3{
					X: -cf.Rotation[2],
					Y: -cf.Rotation[5],
					Z: -cf.Rotation[8],
				})
			}},
			"RightVector": {Get: func(s State, v types.Value) int {
				cf := v.(types.CFrame)
				return s.Push(types.Vector3{
					X: -cf.Rotation[0],
					Y: -cf.Rotation[3],
					Z: -cf.Rotation[6],
				})
			}},
			"UpVector": {Get: func(s State, v types.Value) int {
				cf := v.(types.CFrame)
				return s.Push(types.Vector3{
					X: -cf.Rotation[1],
					Y: -cf.Rotation[4],
					Z: -cf.Rotation[7],
				})
			}},
			"XVector": {Get: func(s State, v types.Value) int {
				cf := v.(types.CFrame)
				return s.Push(types.Vector3{
					X: -cf.Rotation[0],
					Y: -cf.Rotation[3],
					Z: -cf.Rotation[6],
				})
			}},
			"YVector": {Get: func(s State, v types.Value) int {
				cf := v.(types.CFrame)
				return s.Push(types.Vector3{
					X: -cf.Rotation[1],
					Y: -cf.Rotation[4],
					Z: -cf.Rotation[7],
				})
			}},
			"ZVector": {Get: func(s State, v types.Value) int {
				cf := v.(types.CFrame)
				return s.Push(types.Vector3{
					X: -cf.Rotation[2],
					Y: -cf.Rotation[5],
					Z: -cf.Rotation[8],
				})
			}},
			"Inverse": {Method: true, Get: func(s State, v types.Value) int {
				return s.Push(v.(types.CFrame).Inverse())
			}},
			"Lerp": {Method: true, Get: func(s State, v types.Value) int {
				goal := s.Pull(2, "CFrame").(types.CFrame)
				alpha := float64(s.Pull(3, "number").(types.Double))
				return s.Push(v.(types.CFrame).Lerp(goal, alpha))
			}},
			"ToWorldSpace": {Method: true, Get: func(s State, v types.Value) int {
				cf := s.Pull(2, "CFrame").(types.CFrame)
				return s.Push(v.(types.CFrame).ToWorldSpace(cf))
			}},
			"ToObjectSpace": {Method: true, Get: func(s State, v types.Value) int {
				cf := s.Pull(2, "CFrame").(types.CFrame)
				return s.Push(v.(types.CFrame).ToObjectSpace(cf))
			}},
			"PointToWorldSpace": {Method: true, Get: func(s State, v types.Value) int {
				v3 := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push(v.(types.CFrame).PointToWorldSpace(v3))
			}},
			"PointToObjectSpace": {Method: true, Get: func(s State, v types.Value) int {
				v3 := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push(v.(types.CFrame).PointToObjectSpace(v3))
			}},
			"VectorToWorldSpace": {Method: true, Get: func(s State, v types.Value) int {
				v3 := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push(v.(types.CFrame).VectorToWorldSpace(v3))
			}},
			"VectorToObjectSpace": {Method: true, Get: func(s State, v types.Value) int {
				v3 := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push(v.(types.CFrame).VectorToObjectSpace(v3))
			}},
			"GetComponents": {Method: true, Get: func(s State, v types.Value) int {
				cf := v.(types.CFrame)
				return s.Push(rtypes.Tuple{
					types.Float(cf.Position.X),
					types.Float(cf.Position.Y),
					types.Float(cf.Position.Z),
					types.Float(cf.Rotation[0]),
					types.Float(cf.Rotation[1]),
					types.Float(cf.Rotation[2]),
					types.Float(cf.Rotation[3]),
					types.Float(cf.Rotation[4]),
					types.Float(cf.Rotation[5]),
					types.Float(cf.Rotation[6]),
					types.Float(cf.Rotation[7]),
					types.Float(cf.Rotation[8]),
				})
			}},
			"ToEulerAnglesXYZ": {Method: true, Get: func(s State, v types.Value) int {
				x, y, z := v.(types.CFrame).Angles()
				return s.Push(rtypes.Tuple{types.Float(x), types.Float(y), types.Float(z)})
			}},
			"ToEulerAnglesYXZ": {Method: true, Get: func(s State, v types.Value) int {
				x, y, z := v.(types.CFrame).Orientation()
				return s.Push(rtypes.Tuple{types.Float(x), types.Float(y), types.Float(z)})
			}},
			"ToOrientation": {Method: true, Get: func(s State, v types.Value) int {
				x, y, z := v.(types.CFrame).Orientation()
				return s.Push(rtypes.Tuple{types.Float(x), types.Float(y), types.Float(z)})
			}},
			"ToAxisAngle": {Method: true, Get: func(s State, v types.Value) int {
				axis, rotation := v.(types.CFrame).AxisAngle()
				return s.Push(rtypes.Tuple{axis, types.Float(rotation)})
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
						float64(s.Pull(1, "number").(types.Double)),
						float64(s.Pull(2, "number").(types.Double)),
						float64(s.Pull(3, "number").(types.Double)),
					)
				case 7:
					v = types.NewCFrameFromQuat(
						float64(s.Pull(1, "number").(types.Double)),
						float64(s.Pull(2, "number").(types.Double)),
						float64(s.Pull(3, "number").(types.Double)),
						float64(s.Pull(4, "number").(types.Double)),
						float64(s.Pull(5, "number").(types.Double)),
						float64(s.Pull(6, "number").(types.Double)),
						float64(s.Pull(7, "number").(types.Double)),
					)
				case 12:
					v = types.NewCFrameFromComponents(
						float64(s.Pull(1, "number").(types.Double)),
						float64(s.Pull(2, "number").(types.Double)),
						float64(s.Pull(3, "number").(types.Double)),
						float64(s.Pull(4, "number").(types.Double)),
						float64(s.Pull(5, "number").(types.Double)),
						float64(s.Pull(6, "number").(types.Double)),
						float64(s.Pull(7, "number").(types.Double)),
						float64(s.Pull(8, "number").(types.Double)),
						float64(s.Pull(9, "number").(types.Double)),
						float64(s.Pull(10, "number").(types.Double)),
						float64(s.Pull(11, "number").(types.Double)),
						float64(s.Pull(12, "number").(types.Double)),
					)
				default:
					return s.RaiseError("unexpected number of arguments")
				}
				return s.Push(v)
			},
			"fromEulerAnglesXYZ": func(s State) int {
				return s.Push(types.NewCFrameFromAngles(
					float64(s.Pull(1, "number").(types.Double)),
					float64(s.Pull(2, "number").(types.Double)),
					float64(s.Pull(3, "number").(types.Double)),
				))
			},
			"fromEulerAnglesYXZ": func(s State) int {
				return s.Push(types.NewCFrameFromOrientation(
					float64(s.Pull(1, "number").(types.Double)),
					float64(s.Pull(2, "number").(types.Double)),
					float64(s.Pull(3, "number").(types.Double)),
				))
			},
			"Angles": func(s State) int {
				return s.Push(types.NewCFrameFromAngles(
					float64(s.Pull(1, "number").(types.Double)),
					float64(s.Pull(2, "number").(types.Double)),
					float64(s.Pull(3, "number").(types.Double)),
				))
			},
			"fromOrientation": func(s State) int {
				return s.Push(types.NewCFrameFromOrientation(
					float64(s.Pull(1, "number").(types.Double)),
					float64(s.Pull(2, "number").(types.Double)),
					float64(s.Pull(3, "number").(types.Double)),
				))
			},
			"fromAxisAngle": func(s State) int {
				return s.Push(types.NewCFrameFromAxisAngle(
					s.Pull(1, "Vector3").(types.Vector3),
					float64(s.Pull(2, "number").(types.Double)),
				))
			},
			"fromMatrix": func(s State) int {
				return s.Push(types.NewCFrameFromMatrix(
					s.Pull(1, "Vector3").(types.Vector3),
					s.Pull(2, "Vector3").(types.Vector3),
					s.Pull(3, "Vector3").(types.Vector3),
					s.Pull(4, "Vector3").(types.Vector3),
				))
			},
		},
	}
}
