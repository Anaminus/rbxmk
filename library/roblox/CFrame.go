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
func CFrame() Reflector {
	return Reflector{
		Name:     "CFrame",
		PushTo:   rbxmk.PushTypeTo("CFrame"),
		PullFrom: rbxmk.PullTypeFrom("CFrame"),
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
				return s.Push(v.AddV3(op))
			},
			"__sub": func(s State) int {
				v := s.Pull(1, "CFrame").(types.CFrame)
				op := s.Pull(2, "Vector3").(types.Vector3)
				return s.Push(v.SubV3(op))
			},
			"__mul": func(s State) int {
				v := s.Pull(1, "CFrame").(types.CFrame)
				switch op := s.PullAnyOf(2, "CFrame", "Vector3").(type) {
				case types.CFrame:
					return s.Push(v.Mul(op))
				case types.Vector3:
					return s.Push(v.MulV3(op))
				default:
					s.L.ArgError(2, "attempt to multiply a CFrame with an incompatible types.value type or nil")
					return 0
				}
			},
		},
		Members: map[string]Member{
			"P": {
				Get: func(s State, v types.Value) int {
					return s.Push(v.(types.CFrame).Position)
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("Vector3")} },
			},
			"Position": {
				Get: func(s State, v types.Value) int {
					return s.Push(v.(types.CFrame).Position)
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("Vector3")} },
			},
			"X": {
				Get: func(s State, v types.Value) int {
					return s.Push(types.Float(v.(types.CFrame).X()))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("float")} },
			},
			"Y": {
				Get: func(s State, v types.Value) int {
					return s.Push(types.Float(v.(types.CFrame).Y()))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("float")} },
			},
			"Z": {
				Get: func(s State, v types.Value) int {
					return s.Push(types.Float(v.(types.CFrame).Z()))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("float")} },
			},
			"LookVector": {
				Get: func(s State, v types.Value) int {
					cf := v.(types.CFrame)
					return s.Push(types.Vector3{
						X: -cf.Rotation[2],
						Y: -cf.Rotation[5],
						Z: -cf.Rotation[8],
					})
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("Vector3")} },
			},
			"RightVector": {
				Get: func(s State, v types.Value) int {
					cf := v.(types.CFrame)
					return s.Push(types.Vector3{
						X: -cf.Rotation[0],
						Y: -cf.Rotation[3],
						Z: -cf.Rotation[6],
					})
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("Vector3")} },
			},
			"UpVector": {
				Get: func(s State, v types.Value) int {
					cf := v.(types.CFrame)
					return s.Push(types.Vector3{
						X: -cf.Rotation[1],
						Y: -cf.Rotation[4],
						Z: -cf.Rotation[7],
					})
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("Vector3")} },
			},
			"XVector": {
				Get: func(s State, v types.Value) int {
					cf := v.(types.CFrame)
					return s.Push(types.Vector3{
						X: -cf.Rotation[0],
						Y: -cf.Rotation[3],
						Z: -cf.Rotation[6],
					})
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("Vector3")} },
			},
			"YVector": {
				Get: func(s State, v types.Value) int {
					cf := v.(types.CFrame)
					return s.Push(types.Vector3{
						X: -cf.Rotation[1],
						Y: -cf.Rotation[4],
						Z: -cf.Rotation[7],
					})
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("Vector3")} },
			},
			"ZVector": {
				Get: func(s State, v types.Value) int {
					cf := v.(types.CFrame)
					return s.Push(types.Vector3{
						X: -cf.Rotation[2],
						Y: -cf.Rotation[5],
						Z: -cf.Rotation[8],
					})
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("Vector3")} },
			},
			"Inverse": {Method: true,
				Get: func(s State, v types.Value) int {
					return s.Push(v.(types.CFrame).Inverse())
				},
				Dump: func() dump.Value {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Prim("CFrame")},
						},
					}
				},
			},
			"Lerp": {Method: true,
				Get: func(s State, v types.Value) int {
					goal := s.Pull(2, "CFrame").(types.CFrame)
					alpha := float64(s.Pull(3, "number").(types.Double))
					return s.Push(v.(types.CFrame).Lerp(goal, alpha))
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "goal", Type: dt.Prim("CFrame")},
							{Name: "alpha", Type: dt.Prim("float")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("CFrame")},
						},
					}
				},
			},
			"ToWorldSpace": {Method: true,
				Get: func(s State, v types.Value) int {
					cf := s.Pull(2, "CFrame").(types.CFrame)
					return s.Push(v.(types.CFrame).ToWorldSpace(cf))
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "cf", Type: dt.Prim("CFrame")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("CFrame")},
						},
					}
				},
			},
			"ToObjectSpace": {Method: true,
				Get: func(s State, v types.Value) int {
					cf := s.Pull(2, "CFrame").(types.CFrame)
					return s.Push(v.(types.CFrame).ToObjectSpace(cf))
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "cf", Type: dt.Prim("CFrame")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("CFrame")},
						},
					}
				},
			},
			"PointToWorldSpace": {Method: true,
				Get: func(s State, v types.Value) int {
					v3 := s.Pull(2, "Vector3").(types.Vector3)
					return s.Push(v.(types.CFrame).PointToWorldSpace(v3))
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "v", Type: dt.Prim("Vector3")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Vector3")},
						},
					}
				},
			},
			"PointToObjectSpace": {Method: true,
				Get: func(s State, v types.Value) int {
					v3 := s.Pull(2, "Vector3").(types.Vector3)
					return s.Push(v.(types.CFrame).PointToObjectSpace(v3))
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "v", Type: dt.Prim("Vector3")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Vector3")},
						},
					}
				},
			},
			"VectorToWorldSpace": {Method: true,
				Get: func(s State, v types.Value) int {
					v3 := s.Pull(2, "Vector3").(types.Vector3)
					return s.Push(v.(types.CFrame).VectorToWorldSpace(v3))
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "v", Type: dt.Prim("Vector3")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Vector3")},
						},
					}
				},
			},
			"VectorToObjectSpace": {Method: true,
				Get: func(s State, v types.Value) int {
					v3 := s.Pull(2, "Vector3").(types.Vector3)
					return s.Push(v.(types.CFrame).VectorToObjectSpace(v3))
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "v", Type: dt.Prim("Vector3")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Vector3")},
						},
					}
				},
			},
			"GetComponents": {Method: true,
				Get: func(s State, v types.Value) int {
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
				Dump: func() dump.Value {
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
					}
				},
			},
			"ToEulerAnglesXYZ": {Method: true,
				Get: func(s State, v types.Value) int {
					x, y, z := v.(types.CFrame).Angles()
					return s.Push(rtypes.Tuple{types.Float(x), types.Float(y), types.Float(z)})
				},
				Dump: func() dump.Value {
					return dump.Function{
						Returns: dump.Parameters{
							{Name: "rx", Type: dt.Prim("float")},
							{Name: "ry", Type: dt.Prim("float")},
							{Name: "rz", Type: dt.Prim("float")},
						},
					}
				},
			},
			"ToEulerAnglesYXZ": {Method: true,
				Get: func(s State, v types.Value) int {
					x, y, z := v.(types.CFrame).Orientation()
					return s.Push(rtypes.Tuple{types.Float(x), types.Float(y), types.Float(z)})
				},
				Dump: func() dump.Value {
					return dump.Function{
						Returns: dump.Parameters{
							{Name: "rx", Type: dt.Prim("float")},
							{Name: "ry", Type: dt.Prim("float")},
							{Name: "rz", Type: dt.Prim("float")},
						},
					}
				},
			},
			"ToOrientation": {Method: true,
				Get: func(s State, v types.Value) int {
					x, y, z := v.(types.CFrame).Orientation()
					return s.Push(rtypes.Tuple{types.Float(x), types.Float(y), types.Float(z)})
				},
				Dump: func() dump.Value {
					return dump.Function{
						Returns: dump.Parameters{
							{Name: "rx", Type: dt.Prim("float")},
							{Name: "ry", Type: dt.Prim("float")},
							{Name: "rz", Type: dt.Prim("float")},
						},
					}
				},
			},
			"ToAxisAngle": {Method: true,
				Get: func(s State, v types.Value) int {
					axis, rotation := v.(types.CFrame).AxisAngle()
					return s.Push(rtypes.Tuple{axis, types.Float(rotation)})
				},
				Dump: func() dump.Value {
					return dump.Function{
						Returns: dump.Parameters{
							{Name: "axis", Type: dt.Prim("Vector3")},
							{Name: "rotation", Type: dt.Prim("float")},
						},
					}
				},
			},
		},
		Constructors: Constructors{
			"new": {
				Func: func(s State) int {
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
				Dump: func() dump.MultiFunction {
					return []dump.Function{
						{
							Parameters: dump.Parameters{
								{Name: "", Type: dt.Prim("")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("CFrame")},
							},
						},
						{
							Parameters: dump.Parameters{
								{Name: "position", Type: dt.Prim("Vector3")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("CFrame")},
							},
						},
						{
							Parameters: dump.Parameters{
								{Name: "position", Type: dt.Prim("Vector3")},
								{Name: "lookAt", Type: dt.Prim("Vector3")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("CFrame")},
							},
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
						},
					}
				},
			},
			"fromEulerAnglesXYZ": {
				Func: func(s State) int {
					return s.Push(types.NewCFrameFromAngles(
						float64(s.Pull(1, "number").(types.Double)),
						float64(s.Pull(2, "number").(types.Double)),
						float64(s.Pull(3, "number").(types.Double)),
					))
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Parameters: dump.Parameters{
							{Name: "rx", Type: dt.Prim("float")},
							{Name: "ry", Type: dt.Prim("float")},
							{Name: "rz", Type: dt.Prim("float")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("CFrame")},
						},
					}}
				},
			},
			"fromEulerAnglesYXZ": {
				Func: func(s State) int {
					return s.Push(types.NewCFrameFromOrientation(
						float64(s.Pull(1, "number").(types.Double)),
						float64(s.Pull(2, "number").(types.Double)),
						float64(s.Pull(3, "number").(types.Double)),
					))
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Parameters: dump.Parameters{
							{Name: "rx", Type: dt.Prim("float")},
							{Name: "ry", Type: dt.Prim("float")},
							{Name: "rz", Type: dt.Prim("float")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("CFrame")},
						},
					}}
				},
			},
			"Angles": {
				Func: func(s State) int {
					return s.Push(types.NewCFrameFromAngles(
						float64(s.Pull(1, "number").(types.Double)),
						float64(s.Pull(2, "number").(types.Double)),
						float64(s.Pull(3, "number").(types.Double)),
					))
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Parameters: dump.Parameters{
							{Name: "rx", Type: dt.Prim("float")},
							{Name: "ry", Type: dt.Prim("float")},
							{Name: "rz", Type: dt.Prim("float")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("CFrame")},
						},
					}}
				},
			},
			"fromOrientation": {
				Func: func(s State) int {
					return s.Push(types.NewCFrameFromOrientation(
						float64(s.Pull(1, "number").(types.Double)),
						float64(s.Pull(2, "number").(types.Double)),
						float64(s.Pull(3, "number").(types.Double)),
					))
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Parameters: dump.Parameters{
							{Name: "rx", Type: dt.Prim("float")},
							{Name: "ry", Type: dt.Prim("float")},
							{Name: "rz", Type: dt.Prim("float")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("CFrame")},
						},
					}}
				},
			},
			"fromAxisAngle": {
				Func: func(s State) int {
					return s.Push(types.NewCFrameFromAxisAngle(
						s.Pull(1, "Vector3").(types.Vector3),
						float64(s.Pull(2, "number").(types.Double)),
					))
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Parameters: dump.Parameters{
							{Name: "axis", Type: dt.Prim("Vector3")},
							{Name: "rotation", Type: dt.Prim("float")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("CFrame")},
						},
					}}
				},
			},
			"fromMatrix": {
				Func: func(s State) int {
					return s.Push(types.NewCFrameFromMatrix(
						s.Pull(1, "Vector3").(types.Vector3),
						s.Pull(2, "Vector3").(types.Vector3),
						s.Pull(3, "Vector3").(types.Vector3),
						s.Pull(4, "Vector3").(types.Vector3),
					))
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Parameters: dump.Parameters{
							{Name: "position", Type: dt.Prim("Vector3")},
							{Name: "vx", Type: dt.Prim("Vector3")},
							{Name: "vy", Type: dt.Prim("Vector3")},
							{Name: "vz", Type: dt.Prim("Vector3")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("CFrame")},
						},
					}}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{Operators: &dump.Operators{
				Eq: true,
				Mul: []dump.Binop{
					{Operand: dt.Prim("CFrame"), Result: dt.Prim("CFrame")},
					{Operand: dt.Prim("Vector3"), Result: dt.Prim("Vector3")},
				},
				Add: []dump.Binop{{Operand: dt.Prim("Vector3"), Result: dt.Prim("CFrame")}},
				Sub: []dump.Binop{{Operand: dt.Prim("Vector3"), Result: dt.Prim("CFrame")}},
			}}
		},
	}
}
