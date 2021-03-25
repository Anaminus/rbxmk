package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/robloxapi/types"
)

func init() { register(NumberRange) }
func NumberRange() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "NumberRange",
		PushTo:   rbxmk.PushTypeTo("NumberRange"),
		PullFrom: rbxmk.PullTypeFrom("NumberRange"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "NumberRange").(types.NumberRange)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "NumberRange").(types.NumberRange)
				op := s.Pull(2, "NumberRange").(types.NumberRange)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"Min": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.NumberRange).Min))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("float"),
						ReadOnly:    true,
						Summary:     "libraries/roblox/types/NumberRange:Properties/Min/Summary",
						Description: "libraries/roblox/types/NumberRange:Properties/Min/Description",
					}
				},
			},
			"Max": {
				Get: func(s rbxmk.State, v types.Value) int {
					return s.Push(types.Float(v.(types.NumberRange).Max))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("float"),
						ReadOnly:    true,
						Summary:     "libraries/roblox/types/NumberRange:Properties/Max/Summary",
						Description: "libraries/roblox/types/NumberRange:Properties/Max/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": {
				Func: func(s rbxmk.State) int {
					var v types.NumberRange
					switch s.Count() {
					case 1:
						v.Min = float32(s.Pull(1, "float").(types.Float))
						v.Max = v.Min
					case 2:
						v.Min = float32(s.Pull(1, "float").(types.Float))
						v.Max = float32(s.Pull(2, "float").(types.Float))
					default:
						return s.RaiseError("expected 1 or 2 arguments")
					}
					return s.Push(v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						{
							Parameters: dump.Parameters{
								{Name: "value", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("NumberRange")},
							},
							Summary:     "libraries/roblox/types/NumberRange:Constructors/new/1/Summary",
							Description: "libraries/roblox/types/NumberRange:Constructors/new/1/Description",
						},
						{
							Parameters: dump.Parameters{
								{Name: "minimum", Type: dt.Prim("float")},
								{Name: "maxmimum", Type: dt.Prim("float")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("NumberRange")},
							},
							Summary:     "libraries/roblox/types/NumberRange:Constructors/new/2/Summary",
							Description: "libraries/roblox/types/NumberRange:Constructors/new/2/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators:   &dump.Operators{Eq: true},
				Summary:     "libraries/roblox/types/NumberRange:Summary",
				Description: "libraries/roblox/types/NumberRange:Description",
			}
		},
	}
}
