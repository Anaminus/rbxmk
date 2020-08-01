package reflect

import (
	"sort"

	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/types"
)

func init() { register(ClassDesc) }
func ClassDesc() Reflector {
	return Reflector{
		Name:     "ClassDesc",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Members: Members{
			"Name": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.ClassDesc)
					return s.Push(types.String(desc.Name))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.ClassDesc)
					desc.Name = string(s.Pull(3, "string").(types.String))
				},
			},
			"Superclass": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.ClassDesc)
					return s.Push(types.String(desc.Superclass))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.ClassDesc)
					desc.Superclass = string(s.Pull(3, "string").(types.String))
				},
			},
			"MemoryCategory": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.ClassDesc)
					return s.Push(types.String(desc.MemoryCategory))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.ClassDesc)
					desc.MemoryCategory = string(s.Pull(3, "string").(types.String))
				},
			},
			"Member": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.ClassDesc)
				name := string(s.Pull(2, "string").(types.String))
				member, ok := desc.Members[name]
				if !ok {
					return s.Push(rtypes.Nil)
				}
				return s.Push(rtypes.NewMemberDesc(member))
			}},
			"Members": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.ClassDesc)
				members := make(rtypes.Array, 0, len(desc.Members))
				for _, member := range desc.Members {
					members = append(members, rtypes.NewMemberDesc(member))
				}
				sort.Slice(members, func(i, j int) bool {
					return members[i].(rtypes.ClassDesc).Name < members[j].(rtypes.ClassDesc).Name
				})
				return s.Push(members)
			}},
			"AddMember": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.ClassDesc)
				memberDesc := s.PullAnyOf(2,
					"PropertyDesc",
					"FunctionDesc",
					"EventDesc",
					"CallbackDesc",
				)
				var member rbxdump.Member
				switch m := memberDesc.(type) {
				case rtypes.PropertyDesc:
					member = m.Property
				case rtypes.FunctionDesc:
					member = m.Function
				case rtypes.EventDesc:
					member = m.Event
				case rtypes.CallbackDesc:
					member = m.Callback
				}
				if member == nil {
					return s.Push(types.False)
				}
				if _, ok := desc.Members[member.MemberName()]; ok {
					return s.Push(types.False)
				}
				if desc.Members == nil {
					desc.Members = map[string]rbxdump.Member{}
				}
				desc.Members[member.MemberName()] = member

				return s.Push(types.True)
			}},
			"RemoveMember": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.ClassDesc)
				name := string(s.Pull(2, "string").(types.String))
				if _, ok := desc.Members[name]; !ok {
					return s.Push(types.False)
				}
				delete(desc.Members, name)
				return s.Push(types.True)
			}},
			"Tag": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.ClassDesc)
				tag := string(s.Pull(2, "string").(types.String))
				return s.Push(types.Bool(desc.GetTag(tag)))
			}},
			"Tags": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.ClassDesc)
				tags := desc.GetTags()
				array := make(rtypes.Array, len(tags))
				for i, tag := range tags {
					array[i] = types.String(tag)
				}
				return s.Push(array)
			}},
			"SetTag": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.ClassDesc)
				tags := make([]string, s.Count()-1)
				for i := 2; i <= s.Count(); i++ {
					tags[i-2] = string(s.Pull(i, "string").(types.String))
				}
				desc.SetTag(tags...)
				return 0
			}},
			"UnsetTag": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.ClassDesc)
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
