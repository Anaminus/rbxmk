package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Rect) }
func Rect() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     rtypes.T_Rect,
		PushTo:   rbxmk.PushTypeTo(rtypes.T_Rect),
		PullFrom: rbxmk.PullTypeFrom(rtypes.T_Rect),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *types.Rect:
				*p = v.(types.Rect)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Rect).(types.Rect)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Rect).(types.Rect)
				op := s.Pull(2, rtypes.T_Rect).(types.Rect)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"Min": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.Rect).Min)
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Vector2),
						ReadOnly:    true,
						Summary:     "Types/Rect:Properties/Min/Summary",
						Description: "Types/Rect:Properties/Min/Description",
					}
				},
			},
			"Max": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(v.(types.Rect).Max)
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Vector2),
						ReadOnly:    true,
						Summary:     "Types/Rect:Properties/Max/Summary",
						Description: "Types/Rect:Properties/Max/Description",
					}
				},
			},
			"Width": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Rect).Width()))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Float),
						ReadOnly:    true,
						Summary:     "Types/Rect:Properties/Width/Summary",
						Description: "Types/Rect:Properties/Width/Description",
					}
				},
			},
			"Height": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.Rect).Height()))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Float),
						ReadOnly:    true,
						Summary:     "Types/Rect:Properties/Height/Summary",
						Description: "Types/Rect:Properties/Height/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.Rect
					switch s.Count() {
					case 2:
						v.Min = s.Pull(1, rtypes.T_Vector2).(types.Vector2)
						v.Max = s.Pull(2, rtypes.T_Vector2).(types.Vector2)
					case 4:
						v.Min.X = float32(s.Pull(1, rtypes.T_Float).(types.Float))
						v.Min.Y = float32(s.Pull(2, rtypes.T_Float).(types.Float))
						v.Max.Y = float32(s.Pull(3, rtypes.T_Float).(types.Float))
						v.Max.Y = float32(s.Pull(4, rtypes.T_Float).(types.Float))
					default:
						return s.RaiseError("expected 2 or 4 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "min", Type: dt.Prim(rtypes.T_Vector2)},
								{Name: "max", Type: dt.Prim(rtypes.T_Vector2)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_Rect)},
							},
							Summary:     "Types/Rect:Constructors/new/Vector2/Summary",
							Description: "Types/Rect:Constructors/new/Vector2/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "minX", Type: dt.Prim(rtypes.T_Float)},
								{Name: "minY", Type: dt.Prim(rtypes.T_Float)},
								{Name: "maxX", Type: dt.Prim(rtypes.T_Float)},
								{Name: "maxY", Type: dt.Prim(rtypes.T_Float)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_Rect)},
							},
							Summary:     "Types/Rect:Constructors/new/Components/Summary",
							Description: "Types/Rect:Constructors/new/Components/Description",
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
						Summary:     "Types/Rect:Operators/Eq/Summary",
						Description: "Types/Rect:Operators/Eq/Description",
					},
				},
				Summary:     "Types/Rect:Summary",
				Description: "Types/Rect:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Vector2,
			Float,
		},
	}
}
