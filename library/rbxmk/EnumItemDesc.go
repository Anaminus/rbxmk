package reflect

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(EnumItemDesc) }
func EnumItemDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "EnumItemDesc",
		PushTo:   rbxmk.PushPtrTypeTo("EnumItemDesc"),
		PullFrom: rbxmk.PullTypeFrom("EnumItemDesc"),
		Properties: rbxmk.Properties{
			"Name": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EnumItemDesc)
					return s.Push(types.String(desc.Name))
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.EnumItemDesc)
					desc.Name = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						Summary:     "libraries/rbxmk/types/EnumItemDesc:Properties/Name/Summary",
						Description: "libraries/rbxmk/types/EnumItemDesc:Properties/Name/Description",
					}
				},
			},
			"Value": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EnumItemDesc)
					return s.Push(types.Int(desc.Value))
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.EnumItemDesc)
					desc.Value = int(s.Pull(3, "int").(types.Int))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("int"),
						Summary:     "libraries/rbxmk/types/EnumItemDesc:Properties/Value/Summary",
						Description: "libraries/rbxmk/types/EnumItemDesc:Properties/Value/Description",
					}
				},
			},
			"Index": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EnumItemDesc)
					return s.Push(types.Int(desc.Index))
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.EnumItemDesc)
					desc.Index = int(s.Pull(3, "int").(types.Int))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("int"),
						Summary:     "libraries/rbxmk/types/EnumItemDesc:Properties/Index/Summary",
						Description: "libraries/rbxmk/types/EnumItemDesc:Properties/Index/Description",
					}
				},
			},
		},
		Methods: rbxmk.Methods{
			"Tag": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EnumItemDesc)
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
					desc := v.(rtypes.EnumItemDesc)
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
					desc := v.(rtypes.EnumItemDesc)
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
					desc := v.(rtypes.EnumItemDesc)
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
