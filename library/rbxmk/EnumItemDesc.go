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
		Members: rbxmk.Members{
			"Name": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EnumItemDesc)
					return s.Push(types.String(desc.Name))
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.EnumItemDesc)
					desc.Name = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("string")} },
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
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("int")} },
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
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("int")} },
			},
			"Tag": {Method: true,
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EnumItemDesc)
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
					desc := v.(rtypes.EnumItemDesc)
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
					desc := v.(rtypes.EnumItemDesc)
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
					desc := v.(rtypes.EnumItemDesc)
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
	}
}
