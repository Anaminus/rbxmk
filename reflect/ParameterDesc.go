package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/types"
)

func init() { register(ParameterDesc) }
func ParameterDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "ParameterDesc",
		PushTo:   rbxmk.PushTypeTo("ParameterDesc"),
		PullFrom: rbxmk.PullTypeFrom("ParameterDesc"),
		Metatable: rbxmk.Metatable{
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "ParameterDesc").(rtypes.ParameterDesc)
				op := s.Pull(2, "ParameterDesc").(rtypes.ParameterDesc)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"Type": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.ParameterDesc)
					typ := desc.Parameter.Type
					return s.Push(rtypes.TypeDesc{Embedded: typ})
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("TypeDesc"),
						ReadOnly:    true,
						Summary:     "Types/ParameterDesc:Properties/Type/Summary",
						Description: "Types/ParameterDesc:Properties/Type/Description",
					}
				},
			},
			"Name": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.ParameterDesc)
					return s.Push(types.String(desc.Name))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						ReadOnly:    true,
						Summary:     "Types/ParameterDesc:Properties/Name/Summary",
						Description: "Types/ParameterDesc:Properties/Name/Description",
					}
				},
			},
			"Default": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.ParameterDesc)
					if !desc.Optional {
						return s.Push(rtypes.Nil)
					}
					return s.Push(types.String(desc.Default))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Optional{T: dt.Prim("string")},
						ReadOnly:    true,
						Summary:     "Types/ParameterDesc:Properties/Default/Summary",
						Description: "Types/ParameterDesc:Properties/Default/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": rbxmk.Constructor{
				Func: func(s rbxmk.State) int {
					var param rbxdump.Parameter
					param.Type = s.PullOpt(1, "TypeDesc", rtypes.TypeDesc{}).(rtypes.TypeDesc).Embedded
					param.Name = string(s.PullOpt(2, "string", types.String("")).(types.String))
					switch def := s.PullOpt(3, "string", rtypes.Nil).(type) {
					case rtypes.NilType:
						param.Optional = false
					case types.String:
						param.Optional = true
						param.Default = string(def)
					}
					return s.Push(rtypes.ParameterDesc{Parameter: param})
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						dump.Function{
							Parameters: dump.Parameters{
								{Name: "type", Type: dt.Optional{T: dt.Prim("TypeDesc")}},
								{Name: "name", Type: dt.Optional{T: dt.Prim("string")}},
								{Name: "default", Type: dt.Optional{T: dt.Prim("string")}},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("ParameterDesc")},
							},
							Summary:     "Types/ParameterDesc:Constructors/new/Summary",
							Description: "Types/ParameterDesc:Constructors/new/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Operators: &dump.Operators{
					Eq: &dump.Cmpop{
						Summary:     "Types/ParameterDesc:Operators/Eq/Summary",
						Description: "Types/ParameterDesc:Operators/Eq/Description",
					},
				},
				Summary:     "Types/ParameterDesc:Summary",
				Description: "Types/ParameterDesc:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Nil,
			String,
			TypeDesc,
		},
	}
}
