package reflect

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
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
						Summary:     "Types/PropertyDesc:Properties/Name/Summary",
						Description: "Types/PropertyDesc:Properties/Name/Description",
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
						Summary:     "Types/PropertyDesc:Properties/ValueType/Summary",
						Description: "Types/PropertyDesc:Properties/ValueType/Description",
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
						Summary:     "Types/PropertyDesc:Properties/ReadSecurity/Summary",
						Description: "Types/PropertyDesc:Properties/ReadSecurity/Description",
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
						Summary:     "Types/PropertyDesc:Properties/WriteSecurity/Summary",
						Description: "Types/PropertyDesc:Properties/WriteSecurity/Description",
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
						Summary:     "Types/PropertyDesc:Properties/CanLoad/Summary",
						Description: "Types/PropertyDesc:Properties/CanLoad/Description",
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
						Summary:     "Types/PropertyDesc:Properties/CanSave/Summary",
						Description: "Types/PropertyDesc:Properties/CanSave/Description",
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
						Summary:     "Types/PropertyDesc:Methods/Tag/Summary",
						Description: "Types/PropertyDesc:Methods/Tag/Description",
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
						Summary:     "Types/PropertyDesc:Methods/Tags/Summary",
						Description: "Types/PropertyDesc:Methods/Tags/Description",
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
						Summary:     "Types/PropertyDesc:Methods/SetTag/Summary",
						Description: "Types/PropertyDesc:Methods/SetTag/Description",
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
						Summary:     "Types/PropertyDesc:Methods/UnsetTag/Summary",
						Description: "Types/PropertyDesc:Methods/UnsetTag/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": rbxmk.Constructor{
				Func: func(s rbxmk.State) int {
					return s.Push(rtypes.PropertyDesc{Property: &rbxdump.Property{
						ReadSecurity:  "None",
						WriteSecurity: "None",
					}})
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						dump.Function{
							Returns: dump.Parameters{
								{Type: dt.Prim("PropertyDesc")},
							},
							Summary:     "Types/PropertyDesc:Constructors/new/Summary",
							Description: "Types/PropertyDesc:Constructors/new/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/PropertyDesc:Summary",
				Description: "Types/PropertyDesc:Description",
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
