package reflect

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/types"
)

func init() { register(CallbackDesc) }
func CallbackDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "CallbackDesc",
		PushTo:   rbxmk.PushPtrTypeTo("CallbackDesc"),
		PullFrom: rbxmk.PullTypeFrom("CallbackDesc"),
		Properties: rbxmk.Properties{
			"Name": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.CallbackDesc)
					return s.Push(types.String(desc.Name))
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.CallbackDesc)
					desc.Name = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						Summary:     "Types/CallbackDesc:Properties/Name/Summary",
						Description: "Types/CallbackDesc:Properties/Name/Description",
					}
				},
			},
			"ReturnType": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.CallbackDesc)
					returnType := desc.ReturnType
					return s.Push(rtypes.TypeDesc{Embedded: returnType})
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.CallbackDesc)
					desc.ReturnType = s.Pull(3, "TypeDesc").(rtypes.TypeDesc).Embedded
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("TypeDesc"),
						Summary:     "Types/CallbackDesc:Properties/ReturnType/Summary",
						Description: "Types/CallbackDesc:Properties/ReturnType/Description",
					}
				},
			},
			"Security": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.CallbackDesc)
					return s.Push(types.String(desc.Security))
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.CallbackDesc)
					desc.Security = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						Summary:     "Types/CallbackDesc:Properties/Security/Summary",
						Description: "Types/CallbackDesc:Properties/Security/Description",
					}
				},
			},
		},
		Methods: rbxmk.Methods{
			"Parameters": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.CallbackDesc)
					array := make(rtypes.Array, len(desc.Parameters))
					for i, param := range desc.Parameters {
						p := param
						array[i] = rtypes.ParameterDesc{Parameter: p}
					}
					return s.Push(array)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Array{T: dt.Prim("ParameterDesc")}},
						},
						Summary:     "Types/CallbackDesc:Methods/Parameters/Summary",
						Description: "Types/CallbackDesc:Methods/Parameters/Description",
					}
				},
			},
			"SetParameters": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.CallbackDesc)
					array := s.Pull(2, "Array").(rtypes.Array)
					params := make([]rbxdump.Parameter, len(array))
					for i, paramDesc := range array {
						param, ok := paramDesc.(rtypes.ParameterDesc)
						if !ok {
							err := rbxmk.TypeError{Want: param.Type(), Got: paramDesc.Type()}
							return s.ArgError(2, "Array[%d]: %s", i, err)
						}
						params[i] = param.Parameter
					}
					desc.Parameters = params
					return 0
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "params", Type: dt.Array{T: dt.Prim("ParameterDesc")}},
						},
						Summary:     "Types/CallbackDesc:Methods/SetParameters/Summary",
						Description: "Types/CallbackDesc:Methods/SetParameters/Description",
					}
				},
			},
			"Tag": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.CallbackDesc)
					tag := string(s.Pull(2, "string").(types.String))
					return s.Push(types.Bool(desc.GetTag(tag)))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "name", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("bool")},
						},
						Summary:     "Types/CallbackDesc:Methods/Tag/Summary",
						Description: "Types/CallbackDesc:Methods/Tag/Description",
					}
				},
			},
			"Tags": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.CallbackDesc)
					tags := desc.GetTags()
					array := make(rtypes.Array, len(tags))
					for i, tag := range tags {
						array[i] = types.String(tag)
					}
					return s.Push(array)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Array{T: dt.Prim("string")}},
						},
						Summary:     "Types/CallbackDesc:Methods/Tags/Summary",
						Description: "Types/CallbackDesc:Methods/Tags/Description",
					}
				},
			},
			"SetTag": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.CallbackDesc)
					tags := make([]string, s.Count()-1)
					for i := 2; i <= s.Count(); i++ {
						tags[i-2] = string(s.Pull(i, "string").(types.String))
					}
					desc.SetTag(tags...)
					return 0
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "...", Type: dt.Prim("string")},
						},
						Summary:     "Types/CallbackDesc:Methods/SetTag/Summary",
						Description: "Types/CallbackDesc:Methods/SetTag/Description",
					}
				},
			},
			"UnsetTag": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.CallbackDesc)
					tags := make([]string, s.Count()-1)
					for i := 2; i <= s.Count(); i++ {
						tags[i-2] = string(s.Pull(i, "string").(types.String))
					}
					desc.UnsetTag(tags...)
					return 0
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "...", Type: dt.Prim("string")},
						},
						Summary:     "Types/CallbackDesc:Methods/UnsetTag/Summary",
						Description: "Types/CallbackDesc:Methods/UnsetTag/Description",
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/CallbackDesc:Summary",
				Description: "Types/CallbackDesc:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Array,
			Bool,
			ParameterDesc,
			String,
			TypeDesc,
		},
	}
}
