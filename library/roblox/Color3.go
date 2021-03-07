package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Color3) }
func Color3() Reflector {
	return Reflector{
		Name:     "Color3",
		PushTo:   rbxmk.PushTypeTo("Color3"),
		PullFrom: rbxmk.PullTypeFrom("Color3"),
		Metatable: Metatable{
			"__tostring": func(s State) int {
				v := s.Pull(1, "Color3").(types.Color3)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s State) int {
				v := s.Pull(1, "Color3").(types.Color3)
				op := s.Pull(2, "Color3").(types.Color3)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: map[string]Member{
			"R": {
				Get: func(s State, v types.Value) int {
					return s.Push(types.Float(v.(types.Color3).R))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
			"G": {
				Get: func(s State, v types.Value) int {
					return s.Push(types.Float(v.(types.Color3).G))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
			"B": {
				Get: func(s State, v types.Value) int {
					return s.Push(types.Float(v.(types.Color3).B))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
			"Lerp": {Method: true,
				Get: func(s State, v types.Value) int {
					goal := s.Pull(2, "Color3").(types.Color3)
					alpha := float64(s.Pull(3, "float").(types.Float))
					return s.Push(v.(types.Color3).Lerp(goal, alpha))
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "goal", Type: dt.Prim("Color3")},
							{Name: "alpha", Type: dt.Prim("float")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Color3")},
						},
					}
				},
			},
			"ToHSV": {Method: true,
				Get: func(s State, v types.Value) int {
					hue, sat, val := v.(types.Color3).ToHSV()
					return s.Push(rtypes.Tuple{types.Float(hue), types.Float(sat), types.Float(val)})
				},
				Dump: func() dump.Value {
					return dump.Function{
						Returns: dump.Parameters{
							{Name: "h", Type: dt.Prim("float")},
							{Name: "s", Type: dt.Prim("float")},
							{Name: "v", Type: dt.Prim("float")},
						},
					}
				},
			},
		},
		Constructors: Constructors{
			"new": {
				Func: func(s State) int {
					var v types.Color3
					switch s.Count() {
					case 0:
					case 3:
						v.R = float32(s.Pull(1, "float").(types.Float))
						v.G = float32(s.Pull(2, "float").(types.Float))
						v.B = float32(s.Pull(3, "float").(types.Float))
					default:
						return s.RaiseError("expected 0 or 3 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{
						{
							Returns: dump.Parameters{
								{Type: dt.Prim("Color3")},
							},
						},
						{
							Parameters: dump.Parameters{
								{Name: "r", Type: dt.Prim("float")},
								{Name: "g", Type: dt.Prim("float")},
								{Name: "b", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("Color3")},
							},
						},
					}
				},
			},
			"fromRGB": {
				Func: func(s State) int {
					return s.Push(types.NewColor3FromRGB(
						int(s.Pull(1, "int").(types.Int)),
						int(s.Pull(2, "int").(types.Int)),
						int(s.Pull(3, "int").(types.Int)),
					))
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Parameters: dump.Parameters{
							{Name: "r", Type: dt.Prim("int")},
							{Name: "g", Type: dt.Prim("int")},
							{Name: "b", Type: dt.Prim("int")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Color3")},
						},
					}}
				},
			},
			"fromHSV": {
				Func: func(s State) int {
					return s.Push(types.NewColor3FromHSV(
						float64(s.Pull(1, "float").(types.Float)),
						float64(s.Pull(2, "float").(types.Float)),
						float64(s.Pull(3, "float").(types.Float)),
					))
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{{
						Parameters: dump.Parameters{
							{Name: "h", Type: dt.Prim("float")},
							{Name: "s", Type: dt.Prim("float")},
							{Name: "v", Type: dt.Prim("float")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("Color3")},
						},
					}}
				},
			},
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case types.Color3:
				return v
			case rtypes.Color3uint8:
				return types.Color3(v)
			}
			return nil
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
