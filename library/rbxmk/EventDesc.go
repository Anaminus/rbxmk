package reflect

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/types"
)

func init() { register(EventDesc) }
func EventDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "EventDesc",
		PushTo:   rbxmk.PushPtrTypeTo("EventDesc"),
		PullFrom: rbxmk.PullTypeFrom("EventDesc"),
		Properties: rbxmk.Properties{
			"Name": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EventDesc)
					return s.Push(types.String(desc.Name))
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.EventDesc)
					desc.Name = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("string")} },
			},
			"Security": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EventDesc)
					return s.Push(types.String(desc.Security))
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.EventDesc)
					desc.Security = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Property { return dump.Property{ValueType: dt.Prim("string")} },
			},
		},
		Methods: rbxmk.Methods{
			"Parameters": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EventDesc)
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
					}
				},
			},
			"SetParameters": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EventDesc)
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
					}
				},
			},
			"Tag": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EventDesc)
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
					}
				},
			},
			"Tags": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EventDesc)
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
					}
				},
			},
			"SetTag": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EventDesc)
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
					}
				},
			},
			"UnsetTag": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EventDesc)
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
					}
				},
			},
		},
	}
}
