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

func init() { register(FunctionDesc) }
func FunctionDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "FunctionDesc",
		PushTo:   rbxmk.PushTypeTo("FunctionDesc"),
		PullFrom: rbxmk.PullTypeFrom("FunctionDesc"),
		Metatable: rbxmk.Metatable{
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "FunctionDesc").(rtypes.FunctionDesc)
				op := s.Pull(2, "FunctionDesc").(rtypes.FunctionDesc)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: rbxmk.Members{
			"Name": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.FunctionDesc)
					return s.Push(types.String(desc.Name))
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.FunctionDesc)
					desc.Name = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("string")} },
			},
			"Parameters": {Method: true,
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.FunctionDesc)
					array := make(rtypes.Array, len(desc.Parameters))
					for i, param := range desc.Parameters {
						p := param
						array[i] = rtypes.ParameterDesc{Parameter: p}
					}
					return s.Push(array)
				},
				Dump: func() dump.Value {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Array{T: dt.Prim("ParameterDesc")}},
						},
					}
				},
			},
			"SetParameters": {Method: true,
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.FunctionDesc)
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
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "params", Type: dt.Array{T: dt.Prim("ParameterDesc")}},
						},
					}
				},
			},
			"ReturnType": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.FunctionDesc)
					returnType := desc.ReturnType
					return s.Push(rtypes.TypeDesc{Embedded: returnType})
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.FunctionDesc)
					desc.ReturnType = s.Pull(3, "TypeDesc").(rtypes.TypeDesc).Embedded
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("TypeDesc")} },
			},
			"Security": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.FunctionDesc)
					return s.Push(types.String(desc.Security))
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.FunctionDesc)
					desc.Security = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("string")} },
			},
			"Tag": {Method: true,
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.FunctionDesc)
					tag := string(s.Pull(2, "string").(types.String))
					return s.Push(types.Bool(desc.GetTag(tag)))
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "name", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("bool")},
						},
					}
				},
			},
			"Tags": {Method: true,
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.FunctionDesc)
					tags := desc.GetTags()
					array := make(rtypes.Array, len(tags))
					for i, tag := range tags {
						array[i] = types.String(tag)
					}
					return s.Push(array)
				},
				Dump: func() dump.Value {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Array{T: dt.Prim("string")}},
						},
					}
				},
			},
			"SetTag": {Method: true,
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.FunctionDesc)
					tags := make([]string, s.Count()-1)
					for i := 2; i <= s.Count(); i++ {
						tags[i-2] = string(s.Pull(i, "string").(types.String))
					}
					desc.SetTag(tags...)
					return 0
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "...", Type: dt.Prim("string")},
						},
					}
				},
			},
			"UnsetTag": {Method: true,
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.FunctionDesc)
					tags := make([]string, s.Count()-1)
					for i := 2; i <= s.Count(); i++ {
						tags[i-2] = string(s.Pull(i, "string").(types.String))
					}
					desc.UnsetTag(tags...)
					return 0
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "...", Type: dt.Prim("string")},
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef { return dump.TypeDef{Operators: &dump.Operators{Eq: true}} },
	}
}
