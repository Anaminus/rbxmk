package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(CFrame) }
func CFrame() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "CFrame",
		PushTo:   rbxmk.PushTypeTo("CFrame"),
		PullFrom: rbxmk.PullTypeFrom("CFrame"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "CFrame").(types.CFrame)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "CFrame").(types.CFrame)
				op := s.Pull(2, "CFrame").(types.CFrame)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
			"__add": func(s rbxmk.State) int {
				v := s.Pull(1, "CFrame").(types.CFrame)
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push(v.AddV3(op))
			},
			"__sub": func(s rbxmk.State) int {
				v := s.Pull(1, "CFrame").(types.CFrame)
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push(v.SubV3(op))
			},
			"__mul": func(s rbxmk.State) int {
				v := s.Pull(1, "CFrame").(types.CFrame)
				switch op := s.PullAnyOf(2, "CFrame", "Vector3").(type) {
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
						ValueType:   dt.Prim("Vector3"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/CFrame:Properties/P/Summary",
						Description: "Libraries/roblox/Types/CFrame:Properties/P/Description",
					}
				},
			},
			"Position": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.CFrame).Position)
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("Vector3"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/CFrame:Properties/Position/Summary",
						Description: "Libraries/roblox/Types/CFrame:Properties/Position/Description",
					}
				},
			},
			"X": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.CFrame).X()))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("float"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/CFrame:Properties/X/Summary",
						Description: "Libraries/roblox/Types/CFrame:Properties/X/Description",
					}
				},
			},
			"Y": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.CFrame).Y()))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("float"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/CFrame:Properties/Y/Summary",
						Description: "Libraries/roblox/Types/CFrame:Properties/Y/Description",
					}
				},
			},
			"Z": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.CFrame).Z()))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("float"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/CFrame:Properties/Z/Summary",
						Description: "Libraries/roblox/Types/CFrame:Properties/Z/Description",
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
						ValueType:   dt.Prim("Vector3"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/CFrame:Properties/LookVector/Summary",
						Description: "Libraries/roblox/Types/CFrame:Properties/LookVector/Description",
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
						ValueType:   dt.Prim("Vector3"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/CFrame:Properties/RightVector/Summary",
						Description: "Libraries/roblox/Types/CFrame:Properties/RightVector/Description",
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
						ValueType:   dt.Prim("Vector3"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/CFrame:Properties/UpVector/Summary",
						Description: "Libraries/roblox/Types/CFrame:Properties/UpVector/Description",
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
						ValueType:   dt.Prim("Vector3"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/CFrame:Properties/XVector/Summary",
						Description: "Libraries/roblox/Types/CFrame:Properties/XVector/Description",
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
						ValueType:   dt.Prim("Vector3"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/CFrame:Properties/YVector/Summary",
						Description: "Libraries/roblox/Types/CFrame:Properties/YVector/Description",
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
						ValueType:   dt.Prim("Vector3"),
						ReadOnly:    true,
						Summary:     "Libraries/roblox/Types/CFrame:Properties/ZVector/Summary",
						Description: "Libraries/roblox/Types/CFrame:Properties/ZVector/Description",
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
							{Type: dt.Prim("CFrame")},
						},
						Summary:     "Libraries/roblox/Types/CFrame:Methods/Inverse/Summary",
						Description: "Libraries/roblox/Types/CFrame:Methods/Inverse/Description",
					}
				},
			},
			"Lerp": {
				Func: func(s rbxmk.State, v types.Value) int {
					goal := s.Pull(2, "CFrame").(types.CFrame)
					alpha := float64(s.Pull(3, "float").(types.Float))
					return s.Push(v.(types.CFrame).Lerp(goal, alpha))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "goal", Type: dt.Prim("CFrame")},
							{Name: "alpha", Type: dt.Prim("float")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("CFrame")},
						},
						Summary:     "Libraries/roblox/Types/CFrame:Methods/Lerp/Summary",
						Description: "Libraries/roblox/Types/CFrame:Methods/Lerp/Description",
					}
				},
			},
			"ToWorldSpace": {
				Func: func(s rbxmk.State, v types.Value) int {
					cf := s.Pull(2, "CFrame").(types.CFrame)
					return s.Push(v.(types.CFrame).ToWorldSpace(cf))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "cf", Type: dt.Prim("CFrame")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("CFrame")},
						},
						Summary:     "Libraries/roblox/Types/CFrame:Methods/ToWorldSpace/Summary",
						Description: "Libraries/roblox/Types/CFrame:Methods/ToWorldSpace/Description",
					}
				},
			},
			"ToObjectSpace": {
				Func: func(s rbxmk.State, v types.Value) int {
					cf := s.Pull(2, "CFrame").(types.CFrame)
					return s.Push(v.(types.CFrame).ToObjectSpace(cf))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "cf", Type: dt.Prim("CFrame")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("CFrame")},
						},
						Summary:     "Libraries/roblox/Types/CFrame:Methods/ToObjectSpace/Summary",
						Description: "Libraries/roblox/Types/CFrame:Methods/ToObjectSpace/Description",
					}
				},
			},
			"PointToWorldSpace": {
				Func: func(s rbxmk.State, v types.Value) int {
					v3 := s.Pull(2, "Vector3").(types.Vector3)
					return s.Push(v.(types.CFrame).PointToWorldSpace(v3))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "v", Type: dt.Prim("Vector3")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Vector3")},
						},
						Summary:     "Libraries/roblox/Types/CFrame:Methods/PointToWorldSpace/Summary",
						Description: "Libraries/roblox/Types/CFrame:Methods/PointToWorldSpace/Description",
					}
				},
			},
			"PointToObjectSpace": {
				Func: func(s rbxmk.State, v types.Value) int {
					v3 := s.Pull(2, "Vector3").(types.Vector3)
					return s.Push(v.(types.CFrame).PointToObjectSpace(v3))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "v", Type: dt.Prim("Vector3")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Vector3")},
						},
						Summary:     "Libraries/roblox/Types/CFrame:Methods/PointToObjectSpace/Summary",
						Description: "Libraries/roblox/Types/CFrame:Methods/PointToObjectSpace/Description",
					}
				},
			},
			"VectorToWorldSpace": {
				Func: func(s rbxmk.State, v types.Value) int {
					v3 := s.Pull(2, "Vector3").(types.Vector3)
					return s.Push(v.(types.CFrame).VectorToWorldSpace(v3))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "v", Type: dt.Prim("Vector3")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Vector3")},
						},
						Summary:     "Libraries/roblox/Types/CFrame:Methods/VectorToWorldSpace/Summary",
						Description: "Libraries/roblox/Types/CFrame:Methods/VectorToWorldSpace/Description",
					}
				},
			},
			"VectorToObjectSpace": {
				Func: func(s rbxmk.State, v types.Value) int {
					v3 := s.Pull(2, "Vector3").(types.Vector3)
					return s.Push(v.(types.CFrame).VectorToObjectSpace(v3))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "v", Type: dt.Prim("Vector3")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Vector3")},
						},
						Summary:     "Libraries/roblox/Types/CFrame:Methods/VectorToObjectSpace/Summary",
						Description: "Libraries/roblox/Types/CFrame:Methods/VectorToObjectSpace/Description",
					}
				},
			},
			"GetComponents": {
				Func: func(s rbxmk.State, v types.Value) int {
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
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Name: "x", Type: dt.Prim("float")},
							{Name: "y", Type: dt.Prim("float")},
							{Name: "z", Type: dt.Prim("float")},
							{Name: "r00", Type: dt.Prim("float")},
							{Name: "r01", Type: dt.Prim("float")},
							{Name: "r02", Type: dt.Prim("float")},
							{Name: "r10", Type: dt.Prim("float")},
							{Name: "r11", Type: dt.Prim("float")},
							{Name: "r12", Type: dt.Prim("float")},
							{Name: "r20", Type: dt.Prim("float")},
							{Name: "r21", Type: dt.Prim("float")},
							{Name: "r22", Type: dt.Prim("float")},
						},
						Summary:     "Libraries/roblox/Types/CFrame:Methods/GetComponents/Summary",
						Description: "Libraries/roblox/Types/CFrame:Methods/GetComponents/Description",
					}
				},
			},
			"ToEulerAnglesXYZ": {
				Func: func(s rbxmk.State, v types.Value) int {
					x, y, z := v.(types.CFrame).Angles()
					return s.Push(rtypes.Tuple{types.Float(x), types.Float(y), types.Float(z)})
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Name: "rx", Type: dt.Prim("float")},
							{Name: "ry", Type: dt.Prim("float")},
							{Name: "rz", Type: dt.Prim("float")},
						},
						Summary:     "Libraries/roblox/Types/CFrame:Methods/ToEulerAnglesXYZ/Summary",
						Description: "Libraries/roblox/Types/CFrame:Methods/ToEulerAnglesXYZ/Description",
					}
				},
			},
			"ToEulerAnglesYXZ": {
				Func: func(s rbxmk.State, v types.Value) int {
					x, y, z := v.(types.CFrame).Orientation()
					return s.Push(rtypes.Tuple{types.Float(x), types.Float(y), types.Float(z)})
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Name: "rx", Type: dt.Prim("float")},
							{Name: "ry", Type: dt.Prim("float")},
							{Name: "rz", Type: dt.Prim("float")},
						},
						Summary:     "Libraries/roblox/Types/CFrame:Methods/ToEulerAnglesYXZ/Summary",
						Description: "Libraries/roblox/Types/CFrame:Methods/ToEulerAnglesYXZ/Description",
					}
				},
			},
			"ToOrientation": {
				Func: func(s rbxmk.State, v types.Value) int {
					x, y, z := v.(types.CFrame).Orientation()
					return s.Push(rtypes.Tuple{types.Float(x), types.Float(y), types.Float(z)})
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Name: "rx", Type: dt.Prim("float")},
							{Name: "ry", Type: dt.Prim("float")},
							{Name: "rz", Type: dt.Prim("float")},
						},
						Summary:     "Libraries/roblox/Types/CFrame:Methods/ToOrientation/Summary",
						Description: "Libraries/roblox/Types/CFrame:Methods/ToOrientation/Description",
					}
				},
			},
			"ToAxisAngle": {
				Func: func(s rbxmk.State, v types.Value) int {
					axis, rotation := v.(types.CFrame).AxisAngle()
					return s.Push(rtypes.Tuple{axis, types.Float(rotation)})
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Name: "axis", Type: dt.Prim("Vector3")},
							{Name: "rotation", Type: dt.Prim("float")},
						},
						Summary:     "Libraries/roblox/Types/CFrame:Methods/ToAxisAngle/Summary",
						Description: "Libraries/roblox/Types/CFrame:Methods/ToAxisAngle/Description",
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
						pos := s.Pull(1, "Vector3").(types.Vector3)
						v = types.NewCFrameFromVector3(pos)
					case 2:
						pos := s.Pull(1, "Vector3").(types.Vector3)
						lookAt := s.Pull(2, "Vector3").(types.Vector3)
						v = types.NewCFrameFromLook(pos, lookAt)
					case 3:
						v = types.NewCFrameFromPosition(
							float64(s.Pull(1, "float").(types.Float)),
							float64(s.Pull(2, "float").(types.Float)),
							float64(s.Pull(3, "float").(types.Float)),
						)
					case 7:
						v = types.NewCFrameFromQuat(
							float64(s.Pull(1, "float").(types.Float)),
							float64(s.Pull(2, "float").(types.Float)),
							float64(s.Pull(3, "float").(types.Float)),
							float64(s.Pull(4, "float").(types.Float)),
							float64(s.Pull(5, "float").(types.Float)),
							float64(s.Pull(6, "float").(types.Float)),
							float64(s.Pull(7, "float").(types.Float)),
						)
					case 12:
						v = types.NewCFrameFromComponents(
							float64(s.Pull(1, "float").(types.Float)),
							float64(s.Pull(2, "float").(types.Float)),
							float64(s.Pull(3, "float").(types.Float)),
							float64(s.Pull(4, "float").(types.Float)),
							float64(s.Pull(5, "float").(types.Float)),
							float64(s.Pull(6, "float").(types.Float)),
							float64(s.Pull(7, "float").(types.Float)),
							float64(s.Pull(8, "float").(types.Float)),
							float64(s.Pull(9, "float").(types.Float)),
							float64(s.Pull(10, "float").(types.Float)),
							float64(s.Pull(11, "float").(types.Float)),
							float64(s.Pull(12, "float").(types.Float)),
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
								{Type: dt.Prim("CFrame")},
							},
							Summary:     "Libraries/roblox/Types/CFrame:Constructors/new/1/Summary",
							Description: "Libraries/roblox/Types/CFrame:Constructors/new/1/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "position", Type: dt.Prim("Vector3")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("CFrame")},
							},
							Summary:     "Libraries/roblox/Types/CFrame:Constructors/new/2/Summary",
							Description: "Libraries/roblox/Types/CFrame:Constructors/new/2/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "position", Type: dt.Prim("Vector3")},
								{Name: "lookAt", Type: dt.Prim("Vector3")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("CFrame")},
							},
							Summary:     "Libraries/roblox/Types/CFrame:Constructors/new/3/Summary",
							Description: "Libraries/roblox/Types/CFrame:Constructors/new/3/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("float")},
								{Name: "y", Type: dt.Prim("float")},
								{Name: "z", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("CFrame")},
							},
							Summary:     "Libraries/roblox/Types/CFrame:Constructors/new/4/Summary",
							Description: "Libraries/roblox/Types/CFrame:Constructors/new/4/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("float")},
								{Name: "y", Type: dt.Prim("float")},
								{Name: "z", Type: dt.Prim("float")},
								{Name: "qx", Type: dt.Prim("float")},
								{Name: "qy", Type: dt.Prim("float")},
								{Name: "qz", Type: dt.Prim("float")},
								{Name: "qw", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("CFrame")},
							},
							Summary:     "Libraries/roblox/Types/CFrame:Constructors/new/5/Summary",
							Description: "Libraries/roblox/Types/CFrame:Constructors/new/5/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "x", Type: dt.Prim("float")},
								{Name: "y", Type: dt.Prim("float")},
								{Name: "z", Type: dt.Prim("float")},
								{Name: "r00", Type: dt.Prim("float")},
								{Name: "r01", Type: dt.Prim("float")},
								{Name: "r02", Type: dt.Prim("float")},
								{Name: "r10", Type: dt.Prim("float")},
								{Name: "r11", Type: dt.Prim("float")},
								{Name: "r12", Type: dt.Prim("float")},
								{Name: "r20", Type: dt.Prim("float")},
								{Name: "r21", Type: dt.Prim("float")},
								{Name: "r22", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("CFrame")},
							},
							Summary:     "Libraries/roblox/Types/CFrame:Constructors/new/6/Summary",
							Description: "Libraries/roblox/Types/CFrame:Constructors/new/6/Description",
						},
					}
				},
			},
			"fromEulerAnglesXYZ": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewCFrameFromAngles(
						float64(s.Pull(1, "float").(types.Float)),
						float64(s.Pull(2, "float").(types.Float)),
						float64(s.Pull(3, "float").(types.Float)),
					))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "rx", Type: dt.Prim("float")},
								{Name: "ry", Type: dt.Prim("float")},
								{Name: "rz", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("CFrame")},
							},
							Summary:     "Libraries/roblox/Types/CFrame:Constructors/fromEulerAnglesXYZ/Summary",
							Description: "Libraries/roblox/Types/CFrame:Constructors/fromEulerAnglesXYZ/Description",
						},
					}
				},
			},
			"fromEulerAnglesYXZ": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewCFrameFromOrientation(
						float64(s.Pull(1, "float").(types.Float)),
						float64(s.Pull(2, "float").(types.Float)),
						float64(s.Pull(3, "float").(types.Float)),
					))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "rx", Type: dt.Prim("float")},
								{Name: "ry", Type: dt.Prim("float")},
								{Name: "rz", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("CFrame")},
							},
							Summary:     "Libraries/roblox/Types/CFrame:Constructors/fromEulerAnglesYXZ/Summary",
							Description: "Libraries/roblox/Types/CFrame:Constructors/fromEulerAnglesYXZ/Description",
						},
					}
				},
			},
			"Angles": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewCFrameFromAngles(
						float64(s.Pull(1, "float").(types.Float)),
						float64(s.Pull(2, "float").(types.Float)),
						float64(s.Pull(3, "float").(types.Float)),
					))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "rx", Type: dt.Prim("float")},
								{Name: "ry", Type: dt.Prim("float")},
								{Name: "rz", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("CFrame")},
							},
							Summary:     "Libraries/roblox/Types/CFrame:Constructors/Angles/Summary",
							Description: "Libraries/roblox/Types/CFrame:Constructors/Angles/Description",
						},
					}
				},
			},
			"fromOrientation": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewCFrameFromOrientation(
						float64(s.Pull(1, "float").(types.Float)),
						float64(s.Pull(2, "float").(types.Float)),
						float64(s.Pull(3, "float").(types.Float)),
					))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "rx", Type: dt.Prim("float")},
								{Name: "ry", Type: dt.Prim("float")},
								{Name: "rz", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("CFrame")},
							},
							Summary:     "Libraries/roblox/Types/CFrame:Constructors/fromOrientation/Summary",
							Description: "Libraries/roblox/Types/CFrame:Constructors/fromOrientation/Description",
						},
					}
				},
			},
			"fromAxisAngle": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewCFrameFromAxisAngle(
						s.Pull(1, "Vector3").(types.Vector3),
						float64(s.Pull(2, "float").(types.Float)),
					))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "axis", Type: dt.Prim("Vector3")},
								{Name: "rotation", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("CFrame")},
							},
							Summary:     "Libraries/roblox/Types/CFrame:Constructors/fromAxisAngle/Summary",
							Description: "Libraries/roblox/Types/CFrame:Constructors/fromAxisAngle/Description",
						},
					}
				},
			},
			"fromMatrix": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewCFrameFromMatrix(
						s.Pull(1, "Vector3").(types.Vector3),
						s.Pull(2, "Vector3").(types.Vector3),
						s.Pull(3, "Vector3").(types.Vector3),
						s.Pull(4, "Vector3").(types.Vector3),
					))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "position", Type: dt.Prim("Vector3")},
								{Name: "vx", Type: dt.Prim("Vector3")},
								{Name: "vy", Type: dt.Prim("Vector3")},
								{Name: "vz", Type: dt.Prim("Vector3")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("CFrame")},
							},
							Summary:     "Libraries/roblox/Types/CFrame:Constructors/fromMatrix/Summary",
							Description: "Libraries/roblox/Types/CFrame:Constructors/fromMatrix/Description",
						},
					}
				},
			},
			"lookAt": {
				Func: func(s rbxmk.State) int {
					return s.Push(types.NewCFrameFromLookAt(
						s.Pull(1, "Vector3").(types.Vector3),
						s.Pull(2, "Vector3").(types.Vector3),
						s.PullOpt(3, "Vector3", types.Vector3{X: 0, Y: 1, Z: 0}).(types.Vector3),
					))
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "position", Type: dt.Prim("Vector3")},
								{Name: "lookAt", Type: dt.Prim("Vector3")},
								{Name: "up", Type: dt.Optional{T: dt.Prim("Vector3")}, Default: `Vector3.new(0, 1, 0)`},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("CFrame")},
							},
							Summary:     "Libraries/roblox/Types/CFrame:Constructors/lookAt/Summary",
							Description: "Libraries/roblox/Types/CFrame:Constructors/lookAt/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Eq: true,
					Mul: []dump.Binop{
						{
							Operand:     dt.Prim("CFrame"),
							Result:      dt.Prim("CFrame"),
							Summary:     "Libraries/roblox/Types/CFrame:Operators/Mul/1/Summary",
							Description: "Libraries/roblox/Types/CFrame:Operators/Mul/1/Description",
						},
						{
							Operand:     dt.Prim("Vector3"),
							Result:      dt.Prim("Vector3"),
							Summary:     "Libraries/roblox/Types/CFrame:Operators/Mul/2/Summary",
							Description: "Libraries/roblox/Types/CFrame:Operators/Mul/2/Description",
						},
					},
					Add: []dump.Binop{
						{
							Operand:     dt.Prim("Vector3"),
							Result:      dt.Prim("CFrame"),
							Summary:     "Libraries/roblox/Types/CFrame:Operators/Add/Summary",
							Description: "Libraries/roblox/Types/CFrame:Operators/Add/Description",
						},
					},
					Sub: []dump.Binop{
						{
							Operand:     dt.Prim("Vector3"),
							Result:      dt.Prim("CFrame"),
							Summary:     "Libraries/roblox/Types/CFrame:Operators/Sub/Summary",
							Description: "Libraries/roblox/Types/CFrame:Operators/Sub/Description",
						},
					},
				},
				Summary:     "Libraries/roblox/Types/CFrame:Summary",
				Description: "Libraries/roblox/Types/CFrame:Description",
			}
		},
	}
}
