package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

const T_CFrame = "CFrame"

func init() { register(CFrame) }
func CFrame() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     T_CFrame,
		PushTo:   rbxmk.PushTypeTo(T_CFrame),
		PullFrom: rbxmk.PullTypeFrom(T_CFrame),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.CFrame:
				*p = v.(types.CFrame)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, T_CFrame).(types.CFrame)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, T_CFrame).(types.CFrame)
				op := s.Pull(2, T_CFrame).(types.CFrame)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
			"__add": func(s rbxmk.State) int {
				v := s.Pull(1, T_CFrame).(types.CFrame)
				op := s.Pull(2, T_Vector3).(types.Vector3)
				return s.Push(v.AddV3(op))
			},
			"__sub": func(s rbxmk.State) int {
				v := s.Pull(1, T_CFrame).(types.CFrame)
				op := s.Pull(2, T_Vector3).(types.Vector3)
				return s.Push(v.SubV3(op))
			},
			"__mul": func(s rbxmk.State) int {
				v := s.Pull(1, T_CFrame).(types.CFrame)
				switch op := s.PullAnyOf(2, T_CFrame, T_Vector3).(type) {
				case types.CFrame:
					return s.Push(v.Mul(op))
				case types.Vector3:
					return s.Push(v.MulV3(op))
				default:
					return s.ReflectorError(2)
				}
			},
		},
		Properties: map[string]rbxmk.Property{
			"P": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.CFrame).Position)
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Vector3),
						ReadOnly:    true,
						Summary:     "Types/CFrame:Properties/P/Summary",
						Description: "Types/CFrame:Properties/P/Description",
					}
				},
			},
			"Position": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.CFrame).Position)
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Vector3),
						ReadOnly:    true,
						Summary:     "Types/CFrame:Properties/Position/Summary",
						Description: "Types/CFrame:Properties/Position/Description",
					}
				},
			},
			"X": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.CFrame).X()))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Float),
						ReadOnly:    true,
						Summary:     "Types/CFrame:Properties/X/Summary",
						Description: "Types/CFrame:Properties/X/Description",
					}
				},
			},
			"Y": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.CFrame).Y()))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Float),
						ReadOnly:    true,
						Summary:     "Types/CFrame:Properties/Y/Summary",
						Description: "Types/CFrame:Properties/Y/Description",
					}
				},
			},
			"Z": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.CFrame).Z()))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Float),
						ReadOnly:    true,
						Summary:     "Types/CFrame:Properties/Z/Summary",
						Description: "Types/CFrame:Properties/Z/Description",
					}
				},
			},
			"LookVector": {
				Get: func(s rbxmk.State, v types.Value) int {
					cf := v.(types.CFrame)
					return s.Push(types.Vector3{
						X: -cf.Rotation[2],
						Y: -cf.Rotation[5],
						Z: -cf.Rotation[8],
					})
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Vector3),
						ReadOnly:    true,
						Summary:     "Types/CFrame:Properties/LookVector/Summary",
						Description: "Types/CFrame:Properties/LookVector/Description",
					}
				},
			},
			"RightVector": {
				Get: func(s rbxmk.State, v types.Value) int {
					cf := v.(types.CFrame)
					return s.Push(types.Vector3{
						X: -cf.Rotation[0],
						Y: -cf.Rotation[3],
						Z: -cf.Rotation[6],
					})
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Vector3),
						ReadOnly:    true,
						Summary:     "Types/CFrame:Properties/RightVector/Summary",
						Description: "Types/CFrame:Properties/RightVector/Description",
					}
				},
			},
			"UpVector": {
				Get: func(s rbxmk.State, v types.Value) int {
					cf := v.(types.CFrame)
					return s.Push(types.Vector3{
						X: -cf.Rotation[1],
						Y: -cf.Rotation[4],
						Z: -cf.Rotation[7],
					})
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Vector3),
						ReadOnly:    true,
						Summary:     "Types/CFrame:Properties/UpVector/Summary",
						Description: "Types/CFrame:Properties/UpVector/Description",
					}
				},
			},
			"XVector": {
				Get: func(s rbxmk.State, v types.Value) int {
					cf := v.(types.CFrame)
					return s.Push(types.Vector3{
						X: -cf.Rotation[0],
						Y: -cf.Rotation[3],
						Z: -cf.Rotation[6],
					})
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Vector3),
						ReadOnly:    true,
						Summary:     "Types/CFrame:Properties/XVector/Summary",
						Description: "Types/CFrame:Properties/XVector/Description",
					}
				},
			},
			"YVector": {
				Get: func(s rbxmk.State, v types.Value) int {
					cf := v.(types.CFrame)
					return s.Push(types.Vector3{
						X: -cf.Rotation[1],
						Y: -cf.Rotation[4],
						Z: -cf.Rotation[7],
					})
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Vector3),
						ReadOnly:    true,
						Summary:     "Types/CFrame:Properties/YVector/Summary",
						Description: "Types/CFrame:Properties/YVector/Description",
					}
				},
			},
			"ZVector": {
				Get: func(s rbxmk.State, v types.Value) int {
					cf := v.(types.CFrame)
					return s.Push(types.Vector3{
						X: -cf.Rotation[2],
						Y: -cf.Rotation[5],
						Z: -cf.Rotation[8],
					})
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(T_Vector3),
						ReadOnly:    true,
						Summary:     "Types/CFrame:Properties/ZVector/Summary",
						Description: "Types/CFrame:Properties/ZVector/Description",
					}
				},
			},
		},
		Methods: rbxmk.Methods{
			"Inverse": {
				Func: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.CFrame).Inverse())
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Prim(T_CFrame)},
						},
						Summary:     "Types/CFrame:Methods/Inverse/Summary",
						Description: "Types/CFrame:Methods/Inverse/Description",
					}
				},
			},
			"Lerp": {
				Func: func(s rbxmk.State, v types.Value) int {
					goal := s.Pull(2, T_CFrame).(types.CFrame)
					alpha := float64(s.Pull(3, T_Float).(types.Float))
					return s.Push(v.(types.CFrame).Lerp(goal, alpha))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "goal", Type: dt.Prim(T_CFrame)},
							{Name: "alpha", Type: dt.Prim(T_Float)},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(T_CFrame)},
						},
						Summary:     "Types/CFrame:Methods/Lerp/Summary",
						Description: "Types/CFrame:Methods/Lerp/Description",
					}
				},
			},
			"ToWorldSpace": {
				Func: func(s rbxmk.State, v types.Value) int {
					cf := s.Pull(2, T_CFrame).(types.CFrame)
					return s.Push(v.(types.CFrame).ToWorldSpace(cf))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "cf", Type: dt.Prim(T_CFrame)},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(T_CFrame)},
						},
						Summary:     "Types/CFrame:Methods/ToWorldSpace/Summary",
						Description: "Types/CFrame:Methods/ToWorldSpace/Description",
					}
				},
			},
			"ToObjectSpace": {
				Func: func(s rbxmk.State, v types.Value) int {
					cf := s.Pull(2, T_CFrame).(types.CFrame)
					return s.Push(v.(types.CFrame).ToObjectSpace(cf))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "cf", Type: dt.Prim(T_CFrame)},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(T_CFrame)},
						},
						Summary:     "Types/CFrame:Methods/ToObjectSpace/Summary",
						Description: "Types/CFrame:Methods/ToObjectSpace/Description",
					}
				},
			},
			"PointToWorldSpace": {
				Func: func(s rbxmk.State, v types.Value) int {
					v3 := s.Pull(2, T_Vector3).(types.Vector3)
					return s.Push(v.(types.CFrame).PointToWorldSpace(v3))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "v", Type: dt.Prim(T_Vector3)},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(T_Vector3)},
						},
						Summary:     "Types/CFrame:Methods/PointToWorldSpace/Summary",
						Description: "Types/CFrame:Methods/PointToWorldSpace/Description",
					}
				},
			},
			"PointToObjectSpace": {
				Func: func(s rbxmk.State, v types.Value) int {
					v3 := s.Pull(2, T_Vector3).(types.Vector3)
					return s.Push(v.(types.CFrame).PointToObjectSpace(v3))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "v", Type: dt.Prim(T_Vector3)},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(T_Vector3)},
						},
						Summary:     "Types/CFrame:Methods/PointToObjectSpace/Summary",
						Description: "Types/CFrame:Methods/PointToObjectSpace/Description",
					}
				},
			},
			"VectorToWorldSpace": {
				Func: func(s rbxmk.State, v types.Value) int {
					v3 := s.Pull(2, T_Vector3).(types.Vector3)
					return s.Push(v.(types.CFrame).VectorToWorldSpace(v3))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "v", Type: dt.Prim(T_Vector3)},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(T_Vector3)},
						},
						Summary:     "Types/CFrame:Methods/VectorToWorldSpace/Summary",
						Description: "Types/CFrame:Methods/VectorToWorldSpace/Description",
					}
				},
			},
			"VectorToObjectSpace": {
				Func: func(s rbxmk.State, v types.Value) int {
					v3 := s.Pull(2, T_Vector3).(types.Vector3)
					return s.Push(v.(types.CFrame).VectorToObjectSpace(v3))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "v", Type: dt.Prim(T_Vector3)},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(T_Vector3)},
						},
						Summary:     "Types/CFrame:Methods/VectorToObjectSpace/Summary",
						Description: "Types/CFrame:Methods/VectorToObjectSpace/Description",
					}
				},
			},
			"GetComponents": {
				Func: func(s rbxmk.State, v types.Value) int {
					cf := v.(types.CFrame)
					return s.PushTuple(
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
					)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Name: "x", Type: dt.Prim(T_Float)},
							{Name: "y", Type: dt.Prim(T_Float)},
							{Name: "z", Type: dt.Prim(T_Float)},
							{Name: "r00", Type: dt.Prim(T_Float)},
							{Name: "r01", Type: dt.Prim(T_Float)},
							{Name: "r02", Type: dt.Prim(T_Float)},
							{Name: "r10", Type: dt.Prim(T_Float)},
							{Name: "r11", Type: dt.Prim(T_Float)},
							{Name: "r12", Type: dt.Prim(T_Float)},
							{Name: "r20", Type: dt.Prim(T_Float)},
							{Name: "r21", Type: dt.Prim(T_Float)},
							{Name: "r22", Type: dt.Prim(T_Float)},
						},
						Summary:     "Types/CFrame:Methods/GetComponents/Summary",
						Description: "Types/CFrame:Methods/GetComponents/Description",
					}
				},
			},
			"ToEulerAnglesXYZ": {
				Func: func(s rbxmk.State, v types.Value) int {
					x, y, z := v.(types.CFrame).Angles()
					return s.PushTuple(types.Float(x), types.Float(y), types.Float(z))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Name: "rx", Type: dt.Prim(T_Float)},
							{Name: "ry", Type: dt.Prim(T_Float)},
							{Name: "rz", Type: dt.Prim(T_Float)},
						},
						Summary:     "Types/CFrame:Methods/ToEulerAnglesXYZ/Summary",
						Description: "Types/CFrame:Methods/ToEulerAnglesXYZ/Description",
					}
				},
			},
			"ToEulerAnglesYXZ": {
				Func: func(s rbxmk.State, v types.Value) int {
					x, y, z := v.(types.CFrame).Orientation()
					return s.PushTuple(types.Float(x), types.Float(y), types.Float(z))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Name: "rx", Type: dt.Prim(T_Float)},
							{Name: "ry", Type: dt.Prim(T_Float)},
							{Name: "rz", Type: dt.Prim(T_Float)},
						},
						Summary:     "Types/CFrame:Methods/ToEulerAnglesYXZ/Summary",
						Description: "Types/CFrame:Methods/ToEulerAnglesYXZ/Description",
					}
				},
			},
			"ToOrientation": {
				Func: func(s rbxmk.State, v types.Value) int {
					x, y, z := v.(types.CFrame).Orientation()
					return s.PushTuple(types.Float(x), types.Float(y), types.Float(z))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Name: "rx", Type: dt.Prim(T_Float)},
							{Name: "ry", Type: dt.Prim(T_Float)},
							{Name: "rz", Type: dt.Prim(T_Float)},
						},
						Summary:     "Types/CFrame:Methods/ToOrientation/Summary",
						Description: "Types/CFrame:Methods/ToOrientation/Description",
					}
				},
			},
			"ToAxisAngle": {
				Func: func(s rbxmk.State, v types.Value) int {
					axis, rotation := v.(types.CFrame).AxisAngle()
					return s.PushTuple(axis, types.Float(rotation))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Name: "axis", Type: dt.Prim(T_Vector3)},
							{Name: "rotation", Type: dt.Prim(T_Float)},
						},
						Summary:     "Types/CFrame:Methods/ToAxisAngle/Summary",
						Description: "Types/CFrame:Methods/ToAxisAngle/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.CFrame
					switch s.Count() {
					case 0:
						v = types.NewCFrame()
					case 1:
						pos := s.Pull(1, T_Vector3).(types.Vector3)
						v = types.NewCFrameFromVector3(pos)
					case 2:
						pos := s.Pull(1, T_Vector3).(types.Vector3)
						lookAt := s.Pull(2, T_Vector3).(types.Vector3)
						v = types.NewCFrameFromLook(pos, lookAt)
					case 3:
						v = types.NewCFrameFromPosition(
							float64(s.Pull(1, T_Float).(types.Float)),
							float64(s.Pull(2, T_Float).(types.Float)),
							float64(s.Pull(3, T_Float).(types.Float)),
						)
					case 7:
						v = types.NewCFrameFromQuat(
							float64(s.Pull(1, T_Float).(types.Float)),
							float64(s.Pull(2, T_Float).(types.Float)),
							float64(s.Pull(3, T_Float).(types.Float)),
							float64(s.Pull(4, T_Float).(types.Float)),
							float64(s.Pull(5, T_Float).(types.Float)),
							float64(s.Pull(6, T_Float).(types.Float)),
							float64(s.Pull(7, T_Float).(types.Float)),
						)
					case 12:
						v = types.NewCFrameFromComponents(
							float64(s.Pull(1, T_Float).(types.Float)),
							float64(s.Pull(2, T_Float).(types.Float)),
							float64(s.Pull(3, T_Float).(types.Float)),
							float64(s.Pull(4, T_Float).(types.Float)),
							float64(s.Pull(5, T_Float).(types.Float)),
							float64(s.Pull(6, T_Float).(types.Float)),
							float64(s.Pull(7, T_Float).(types.Float)),
							float64(s.Pull(8, T_Float).(types.Float)),
							float64(s.Pull(9, T_Float).(types.Float)),
							float64(s.Pull(10, T_Float).(types.Float)),
							float64(s.Pull(11, T_Float).(types.Float)),
							float64(s.Pull(12, T_Float).(types.Float)),
						)
					default:
						return s.RaiseError("unexpected number of arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Returns: dump.Parameters{
								{Type: dt.Prim(T_CFrame)},
							},
							Summary:     "Types/CFrame:Constructors/new/Identity/Summary",
							Description: "Types/CFrame:Constructors/new/Identity/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "position", Type: dt.Prim(T_Vector3)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_CFrame)},
							},
							Summary:     "Types/CFrame:Constructors/new/Position/Summary",
							Description: "Types/CFrame:Constructors/new/Position/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "position", Type: dt.Prim(T_Vector3)},
								{Name: "lookAt", Type: dt.Prim(T_Vector3)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_CFrame)},
							},
							Summary:     "Types/CFrame:Constructors/new/LookAt/Summary",
							Description: "Types/CFrame:Constructors/new/LookAt/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(T_Float)},
								{Name: "y", Type: dt.Prim(T_Float)},
								{Name: "z", Type: dt.Prim(T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_CFrame)},
							},
							Summary:     "Types/CFrame:Constructors/new/Position components/Summary",
							Description: "Types/CFrame:Constructors/new/Position components/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(T_Float)},
								{Name: "y", Type: dt.Prim(T_Float)},
								{Name: "z", Type: dt.Prim(T_Float)},
								{Name: "qx", Type: dt.Prim(T_Float)},
								{Name: "qy", Type: dt.Prim(T_Float)},
								{Name: "qz", Type: dt.Prim(T_Float)},
								{Name: "qw", Type: dt.Prim(T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_CFrame)},
							},
							Summary:     "Types/CFrame:Constructors/new/Quaternion/Summary",
							Description: "Types/CFrame:Constructors/new/Quaternion/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim(T_Float)},
								{Name: "y", Type: dt.Prim(T_Float)},
								{Name: "z", Type: dt.Prim(T_Float)},
								{Name: "r00", Type: dt.Prim(T_Float)},
								{Name: "r01", Type: dt.Prim(T_Float)},
								{Name: "r02", Type: dt.Prim(T_Float)},
								{Name: "r10", Type: dt.Prim(T_Float)},
								{Name: "r11", Type: dt.Prim(T_Float)},
								{Name: "r12", Type: dt.Prim(T_Float)},
								{Name: "r20", Type: dt.Prim(T_Float)},
								{Name: "r21", Type: dt.Prim(T_Float)},
								{Name: "r22", Type: dt.Prim(T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_CFrame)},
							},
							Summary:     "Types/CFrame:Constructors/new/Components/Summary",
							Description: "Types/CFrame:Constructors/new/Components/Description",
						},
					}
				},
			},
			"fromEulerAnglesXYZ": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewCFrameFromAngles(
						float64(s.Pull(1, T_Float).(types.Float)),
						float64(s.Pull(2, T_Float).(types.Float)),
						float64(s.Pull(3, T_Float).(types.Float)),
					))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "rx", Type: dt.Prim(T_Float)},
								{Name: "ry", Type: dt.Prim(T_Float)},
								{Name: "rz", Type: dt.Prim(T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_CFrame)},
							},
							Summary:     "Types/CFrame:Constructors/fromEulerAnglesXYZ/Summary",
							Description: "Types/CFrame:Constructors/fromEulerAnglesXYZ/Description",
						},
					}
				},
			},
			"fromEulerAnglesYXZ": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewCFrameFromOrientation(
						float64(s.Pull(1, T_Float).(types.Float)),
						float64(s.Pull(2, T_Float).(types.Float)),
						float64(s.Pull(3, T_Float).(types.Float)),
					))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "rx", Type: dt.Prim(T_Float)},
								{Name: "ry", Type: dt.Prim(T_Float)},
								{Name: "rz", Type: dt.Prim(T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_CFrame)},
							},
							Summary:     "Types/CFrame:Constructors/fromEulerAnglesYXZ/Summary",
							Description: "Types/CFrame:Constructors/fromEulerAnglesYXZ/Description",
						},
					}
				},
			},
			"Angles": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewCFrameFromAngles(
						float64(s.Pull(1, T_Float).(types.Float)),
						float64(s.Pull(2, T_Float).(types.Float)),
						float64(s.Pull(3, T_Float).(types.Float)),
					))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "rx", Type: dt.Prim(T_Float)},
								{Name: "ry", Type: dt.Prim(T_Float)},
								{Name: "rz", Type: dt.Prim(T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_CFrame)},
							},
							Summary:     "Types/CFrame:Constructors/Angles/Summary",
							Description: "Types/CFrame:Constructors/Angles/Description",
						},
					}
				},
			},
			"fromOrientation": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewCFrameFromOrientation(
						float64(s.Pull(1, T_Float).(types.Float)),
						float64(s.Pull(2, T_Float).(types.Float)),
						float64(s.Pull(3, T_Float).(types.Float)),
					))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "rx", Type: dt.Prim(T_Float)},
								{Name: "ry", Type: dt.Prim(T_Float)},
								{Name: "rz", Type: dt.Prim(T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_CFrame)},
							},
							Summary:     "Types/CFrame:Constructors/fromOrientation/Summary",
							Description: "Types/CFrame:Constructors/fromOrientation/Description",
						},
					}
				},
			},
			"fromAxisAngle": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewCFrameFromAxisAngle(
						s.Pull(1, T_Vector3).(types.Vector3),
						float64(s.Pull(2, T_Float).(types.Float)),
					))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "axis", Type: dt.Prim(T_Vector3)},
								{Name: "rotation", Type: dt.Prim(T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_CFrame)},
							},
							Summary:     "Types/CFrame:Constructors/fromAxisAngle/Summary",
							Description: "Types/CFrame:Constructors/fromAxisAngle/Description",
						},
					}
				},
			},
			"fromMatrix": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewCFrameFromMatrix(
						s.Pull(1, T_Vector3).(types.Vector3),
						s.Pull(2, T_Vector3).(types.Vector3),
						s.Pull(3, T_Vector3).(types.Vector3),
						s.Pull(4, T_Vector3).(types.Vector3),
					))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "position", Type: dt.Prim(T_Vector3)},
								{Name: "vx", Type: dt.Prim(T_Vector3)},
								{Name: "vy", Type: dt.Prim(T_Vector3)},
								{Name: "vz", Type: dt.Prim(T_Vector3)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_CFrame)},
							},
							Summary:     "Types/CFrame:Constructors/fromMatrix/Summary",
							Description: "Types/CFrame:Constructors/fromMatrix/Description",
						},
					}
				},
			},
			"lookAt": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewCFrameFromLookAt(
						s.Pull(1, T_Vector3).(types.Vector3),
						s.Pull(2, T_Vector3).(types.Vector3),
						s.PullOpt(3, types.Vector3{X: 0, Y: 1, Z: 0}, T_Vector3).(types.Vector3),
					))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "position", Type: dt.Prim(T_Vector3)},
								{Name: "lookAt", Type: dt.Prim(T_Vector3)},
								{Name: "up", Type: dt.Optional{T: dt.Prim(T_Vector3)}, Default: `Vector3.new(0, 1, 0)`},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(T_CFrame)},
							},
							Summary:     "Types/CFrame:Constructors/lookAt/Summary",
							Description: "Types/CFrame:Constructors/lookAt/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Eq: &dump.Cmpop{
						Summary:     "Types/CFrame:Operators/Eq/Summary",
						Description: "Types/CFrame:Operators/Eq/Description",
					},
					Mul: []dump.Binop{
						{
							Operand:     dt.Prim(T_CFrame),
							Result:      dt.Prim(T_CFrame),
							Summary:     "Types/CFrame:Operators/Mul/CFrame/Summary",
							Description: "Types/CFrame:Operators/Mul/CFrame/Description",
						},
						{
							Operand:     dt.Prim(T_Vector3),
							Result:      dt.Prim(T_Vector3),
							Summary:     "Types/CFrame:Operators/Mul/Vector3/Summary",
							Description: "Types/CFrame:Operators/Mul/Vector3/Description",
						},
					},
					Add: []dump.Binop{
						{
							Operand:     dt.Prim(T_Vector3),
							Result:      dt.Prim(T_CFrame),
							Summary:     "Types/CFrame:Operators/Add/Summary",
							Description: "Types/CFrame:Operators/Add/Description",
						},
					},
					Sub: []dump.Binop{
						{
							Operand:     dt.Prim(T_Vector3),
							Result:      dt.Prim(T_CFrame),
							Summary:     "Types/CFrame:Operators/Sub/Summary",
							Description: "Types/CFrame:Operators/Sub/Description",
						},
					},
				},
				Summary:     "Types/CFrame:Summary",
				Description: "Types/CFrame:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Float,
			String,
			Vector3,
		},
	}
}
