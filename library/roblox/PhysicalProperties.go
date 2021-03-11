package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(PhysicalProperties) }
func PhysicalProperties() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name: "PhysicalProperties",
		PushTo: func(s rbxmk.State, v types.Value) (lvs []lua.LValue, err error) {
			if pp, ok := v.(types.PhysicalProperties); ok && !pp.CustomPhysics {
				return append(lvs, lua.LNil), nil
			}
			u := s.L.NewUserData(v)
			s.L.SetMetatable(u, s.L.GetTypeMetatable("PhysicalProperties"))
			return append(lvs, u), nil
		},
		PullFrom: func(s rbxmk.State, lvs ...lua.LValue) (v types.Value, err error) {
			switch lv := lvs[0].(type) {
			case *lua.LNilType:
				return types.PhysicalProperties{}, nil
			case *lua.LUserData:
				if lv.Metatable != s.L.GetTypeMetatable("PhysicalProperties") {
					return nil, rbxmk.TypeError{Want: "PhysicalProperties", Got: lvs[0].Type().String()}
				}
				v, ok := lv.Value().(types.Value)
				if !ok {
					return nil, rbxmk.TypeError{Want: "PhysicalProperties", Got: lvs[0].Type().String()}
				}
				return v, nil
			default:
				return nil, rbxmk.TypeError{Want: "PhysicalProperties", Got: lvs[0].Type().String()}
			}
		},
		ConvertFrom: func(v types.Value) types.Value {
			switch v := v.(type) {
			case rtypes.NilType:
				return types.PhysicalProperties{}
			case types.PhysicalProperties:
				return v
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "PhysicalProperties").(types.PhysicalProperties)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
		},
		Members: map[string]rbxmk.Member{
			"Density": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.PhysicalProperties).Density))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
			"Friction": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.PhysicalProperties).Friction))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
			"Elasticity": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.PhysicalProperties).Elasticity))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
			"FrictionWeight": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.PhysicalProperties).FrictionWeight))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
			"ElasticityWeight": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.PhysicalProperties).ElasticityWeight))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("float"), ReadOnly: true} },
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.PhysicalProperties
					switch s.Count() {
					case 3:
						v.Density = float32(s.Pull(1, "float").(types.Float))
						v.Friction = float32(s.Pull(2, "float").(types.Float))
						v.Elasticity = float32(s.Pull(3, "float").(types.Float))
					case 5:
						v.Density = float32(s.Pull(1, "float").(types.Float))
						v.Friction = float32(s.Pull(2, "float").(types.Float))
						v.Elasticity = float32(s.Pull(3, "float").(types.Float))
						v.FrictionWeight = float32(s.Pull(4, "float").(types.Float))
						v.ElasticityWeight = float32(s.Pull(5, "float").(types.Float))
					default:
						return s.RaiseError("expected 3 or 5 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return []dump.Function{
						{
							Parameters: dump.Parameters{
								{Name: "material", Type: dt.Prim("EnumItem")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("PhysicalProperties")},
							},
							Description: "Not supported.",
						},
						{
							Parameters: dump.Parameters{
								{Name: "density", Type: dt.Prim("float")},
								{Name: "friction", Type: dt.Prim("float")},
								{Name: "elasticity", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("PhysicalProperties")},
							},
						},
						{
							Parameters: dump.Parameters{
								{Name: "density", Type: dt.Prim("float")},
								{Name: "friction", Type: dt.Prim("float")},
								{Name: "elasticity", Type: dt.Prim("float")},
								{Name: "frictionWeight", Type: dt.Prim("float")},
								{Name: "elasticityWeight", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("PhysicalProperties")},
							},
						},
					}
				},
			},
		},
	}
}
