package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func PropertyDesc() Type {
	return Type{
		Name:     "PropertyDesc",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Members: Members{
			"Name": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
					return s.Push(types.String(desc.Name))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.PropertyDesc)
					desc.Name = string(s.Pull(3, "string").(types.String))
				},
			},
			"ValueType": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
					valueType := desc.ValueType
					return s.Push(rtypes.TypeDesc{Embedded: &valueType})
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.PropertyDesc)
					desc.ValueType = *s.Pull(3, "TypeDesc").(rtypes.TypeDesc).Embedded
				},
			},
			"ReadSecurity": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
					return s.Push(types.String(desc.ReadSecurity))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.PropertyDesc)
					desc.ReadSecurity = string(s.Pull(3, "string").(types.String))
				},
			},
			"WriteSecurity": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
					return s.Push(types.String(desc.WriteSecurity))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.PropertyDesc)
					desc.WriteSecurity = string(s.Pull(3, "string").(types.String))
				},
			},
			"CanLoad": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
					return s.Push(types.Bool(desc.CanLoad))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.PropertyDesc)
					desc.CanLoad = bool(s.Pull(3, "bool").(types.Bool))
				},
			},
			"CanSave": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
					return s.Push(types.Bool(desc.CanSave))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.PropertyDesc)
					desc.CanSave = bool(s.Pull(3, "bool").(types.Bool))
				},
			},
			"Tag": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.PropertyDesc)
				tag := string(s.Pull(2, "string").(types.String))
				return s.Push(types.Bool(desc.GetTag(tag)))
			}},
			"Tags": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.PropertyDesc)
				tags := desc.GetTags()
				array := make(rtypes.Array, len(tags))
				for i, tag := range tags {
					array[i] = types.String(tag)
				}
				return s.Push(array)
			}},
			"SetTag": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.PropertyDesc)
				tags := make([]string, s.Count()-1)
				for i := 2; i <= s.Count(); i++ {
					tags[i-2] = string(s.Pull(i, "string").(types.String))
				}
				desc.SetTag(tags...)
				return 0
			}},
			"UnsetTag": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.PropertyDesc)
				tags := make([]string, s.Count()-1)
				for i := 2; i <= s.Count(); i++ {
					tags[i-2] = string(s.Pull(i, "string").(types.String))
				}
				desc.UnsetTag(tags...)
				return 0
			}},
		},
	}
}
