package reflect

import (
	"sort"

	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/types"
	lua "github.com/yuin/gopher-lua"
)

func init() { register(EnumDesc) }
func EnumDesc() Reflector {
	return Reflector{
		Name:     "EnumDesc",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__eq": func(s State) int {
				v := s.Pull(1, "EnumDesc").(rtypes.EnumDesc)
				op := s.Pull(2, "EnumDesc").(rtypes.EnumDesc)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: Members{
			"Name": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.EnumDesc)
					return s.Push(types.String(desc.Name))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.EnumDesc)
					desc.Name = string(s.Pull(3, "string").(types.String))
				},
			},
			"Item": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.EnumDesc)
				name := string(s.Pull(2, "string").(types.String))
				item, ok := desc.Items[name]
				if !ok {
					return s.Push(rtypes.Nil)
				}
				return s.Push(rtypes.EnumItemDesc{EnumItem: item})
			}},
			"Items": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.EnumDesc)
				items := make(rtypes.Array, 0, len(desc.Items))
				for _, item := range desc.Items {
					items = append(items, rtypes.EnumItemDesc{EnumItem: item})
				}
				sort.Slice(items, func(i, j int) bool {
					return items[i].(rtypes.EnumItemDesc).Name < items[j].(rtypes.EnumItemDesc).Name
				})
				return s.Push(items)
			}},
			"AddItem": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.EnumDesc)
				item := s.Pull(2, "ClassDesc").(rtypes.EnumItemDesc)
				if _, ok := desc.Items[item.Name]; ok {
					return s.Push(types.False)
				}
				if desc.Items == nil {
					desc.Items = map[string]*rbxdump.EnumItem{}
				}
				desc.Items[item.Name] = item.EnumItem
				return s.Push(types.True)
			}},
			"RemoveItem": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.EnumDesc)
				name := string(s.Pull(2, "string").(types.String))
				if _, ok := desc.Items[name]; !ok {
					return s.Push(types.False)
				}
				delete(desc.Items, name)
				return s.Push(types.True)
			}},
			"Tag": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.EnumDesc)
				tag := string(s.Pull(2, "string").(types.String))
				return s.Push(types.Bool(desc.GetTag(tag)))
			}},
			"Tags": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.EnumDesc)
				tags := desc.GetTags()
				array := make(rtypes.Array, len(tags))
				for i, tag := range tags {
					array[i] = types.String(tag)
				}
				return s.Push(array)
			}},
			"SetTag": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.EnumDesc)
				tags := make([]string, s.Count()-1)
				for i := 2; i <= s.Count(); i++ {
					tags[i-2] = string(s.Pull(i, "string").(types.String))
				}
				desc.SetTag(tags...)
				return 0
			}},
			"UnsetTag": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.EnumDesc)
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
