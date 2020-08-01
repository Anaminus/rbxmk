package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(EnumItemDesc) }
func EnumItemDesc() Reflector {
	return Reflector{
		Name:     "EnumItemDesc",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Members: Members{
			"Name": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.EnumItemDesc)
					return s.Push(types.String(desc.Name))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.EnumItemDesc)
					desc.Name = string(s.Pull(3, "string").(types.String))
				},
			},
			"Value": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.EnumItemDesc)
					return s.Push(types.Int(desc.Value))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.EnumItemDesc)
					desc.Value = int(s.Pull(3, "int").(types.Int))
				},
			},
			"Index": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.EnumItemDesc)
					return s.Push(types.Int(desc.Index))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.EnumItemDesc)
					desc.Index = int(s.Pull(3, "int").(types.Int))
				},
			},
			"Tag": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.EnumItemDesc)
				tag := string(s.Pull(2, "string").(types.String))
				return s.Push(types.Bool(desc.GetTag(tag)))
			}},
			"Tags": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.EnumItemDesc)
				tags := desc.GetTags()
				array := make(rtypes.Array, len(tags))
				for i, tag := range tags {
					array[i] = types.String(tag)
				}
				return s.Push(array)
			}},
			"SetTag": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.EnumItemDesc)
				tags := make([]string, s.Count()-1)
				for i := 2; i <= s.Count(); i++ {
					tags[i-2] = string(s.Pull(i, "string").(types.String))
				}
				desc.SetTag(tags...)
				return 0
			}},
			"UnsetTag": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.EnumItemDesc)
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
