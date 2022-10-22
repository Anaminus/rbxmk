package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(Font) }
func Font() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     rtypes.T_Font,
		PushTo:   rbxmk.PushTypeTo(rtypes.T_Font),
		PullFrom: rbxmk.PullTypeFrom(rtypes.T_Font),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case *rtypes.Font:
				*p = v.(rtypes.Font)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Font).(rtypes.Font)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, rtypes.T_Font).(rtypes.Font)
				op := s.Pull(2, rtypes.T_Font).(rtypes.Font)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"Family": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Content(v.(rtypes.Font).Family))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Content),
						ReadOnly:    true,
						Summary:     "Types/Font:Properties/Family/Summary",
						Description: "Types/Font:Properties/Family/Description",
					}
				},
			},
			"Weight": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Int(v.(rtypes.Font).Weight))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Int),
						ReadOnly:    true,
						Summary:     "Types/Font:Properties/Weight/Summary",
						Description: "Types/Font:Properties/Weight/Description",
					}
				},
			},
			"Style": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Int(v.(rtypes.Font).Style))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Int),
						ReadOnly:    true,
						Summary:     "Types/Font:Properties/Style/Summary",
						Description: "Types/Font:Properties/Style/Description",
					}
				},
			},
			"CachedFaceId": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Content(v.(rtypes.Font).CachedFaceId))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim(rtypes.T_Content),
						ReadOnly:    true,
						Summary:     "Types/Font:Properties/CachedFaceId/Summary",
						Description: "Types/Font:Properties/CachedFaceId/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v rtypes.Font
					v.Family = string(s.Pull(1, rtypes.T_Content).(types.Content))
					v.Weight = int(s.PullOpt(2, types.Int(400), rtypes.T_Int).(types.Int))
					v.Style = int(s.PullOpt(3, types.Int(0), rtypes.T_Int).(types.Int))
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "family", Type: dt.Prim(rtypes.T_Content)},
								{Name: "weight", Type: dt.Prim(rtypes.T_Int)},
								{Name: "style", Type: dt.Prim(rtypes.T_Int)},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_Font)},
							},
							Summary:     "Types/Font:Constructors/new/Summary",
							Description: "Types/Font:Constructors/new/Description",
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
						Summary:     "Types/Font:Operators/Eq/Summary",
						Description: "Types/Font:Operators/Eq/Description",
					},
				},
				Summary:     "Types/Font:Summary",
				Description: "Types/Font:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Content,
			Int,
		},
	}
}
