package reflect

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(PropertyDesc) }
func PropertyDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "PropertyDesc",
		PushTo:   rbxmk.PushPtrTypeTo("PropertyDesc"),
		PullFrom: rbxmk.PullTypeFrom("PropertyDesc"),
		Properties: rbxmk.Properties{
			"Name": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
					return s.Push(types.String(desc.Name))
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.PropertyDesc)
					desc.Name = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						Summary:     "libraries/rbxmk/types/PropertyDesc:Properties/Name/Summary",
						Description: "libraries/rbxmk/types/PropertyDesc:Properties/Name/Description",
					}
				},
			},
			"ValueType": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
					valueType := desc.ValueType
					return s.Push(rtypes.TypeDesc{Embedded: valueType})
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.PropertyDesc)
					desc.ValueType = s.Pull(3, "TypeDesc").(rtypes.TypeDesc).Embedded
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("TypeDesc"),
						Summary:     "libraries/rbxmk/types/PropertyDesc:Properties/ValueType/Summary",
						Description: "libraries/rbxmk/types/PropertyDesc:Properties/ValueType/Description",
					}
				},
			},
			"ReadSecurity": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
					return s.Push(types.String(desc.ReadSecurity))
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.PropertyDesc)
					desc.ReadSecurity = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						Summary:     "libraries/rbxmk/types/PropertyDesc:Properties/ReadSecurity/Summary",
						Description: "libraries/rbxmk/types/PropertyDesc:Properties/ReadSecurity/Description",
					}
				},
			},
			"WriteSecurity": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
					return s.Push(types.String(desc.WriteSecurity))
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.PropertyDesc)
					desc.WriteSecurity = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						Summary:     "libraries/rbxmk/types/PropertyDesc:Properties/WriteSecurity/Summary",
						Description: "libraries/rbxmk/types/PropertyDesc:Properties/WriteSecurity/Description",
					}
				},
			},
			"CanLoad": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
					return s.Push(types.Bool(desc.CanLoad))
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.PropertyDesc)
					desc.CanLoad = bool(s.Pull(3, "bool").(types.Bool))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("bool"),
						Summary:     "libraries/rbxmk/types/PropertyDesc:Properties/CanLoad/Summary",
						Description: "libraries/rbxmk/types/PropertyDesc:Properties/CanLoad/Description",
					}
				},
			},
			"CanSave": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
					return s.Push(types.Bool(desc.CanSave))
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.PropertyDesc)
					desc.CanSave = bool(s.Pull(3, "bool").(types.Bool))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("bool"),
						Summary:     "libraries/rbxmk/types/PropertyDesc:Properties/CanSave/Summary",
						Description: "libraries/rbxmk/types/PropertyDesc:Properties/CanSave/Description",
					}
				},
			},
		},
		Methods: rbxmk.Methods{
			"Tag": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
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
						Summary:     "$TODO",
						Description: "$TODO",
					}
				},
			},
			"Tags": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
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
						Summary:     "$TODO",
						Description: "$TODO",
					}
				},
			},
			"SetTag": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
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
						Summary:     "$TODO",
						Description: "$TODO",
					}
				},
			},
			"UnsetTag": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
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
						Summary:     "$TODO",
						Description: "$TODO",
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "$TODO",
				Description: "$TODO",
			}
		},
	}
}
