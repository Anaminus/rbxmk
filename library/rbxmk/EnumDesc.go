package reflect

import (
	"sort"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/types"
)

func init() { register(EnumDesc) }
func EnumDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "EnumDesc",
		PushTo:   rbxmk.PushPtrTypeTo("EnumDesc"),
		PullFrom: rbxmk.PullTypeFrom("EnumDesc"),
		Metatable: rbxmk.Metatable{
			"__eq": func(s rbxmk.State) int {
				v := s.Pull(1, "EnumDesc").(rtypes.EnumDesc)
				op := s.Pull(2, "EnumDesc").(rtypes.EnumDesc)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: rbxmk.Members{
			"Name": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EnumDesc)
					return s.Push(types.String(desc.Name))
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.EnumDesc)
					desc.Name = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("string")} },
			},
			"Item": {Method: true,
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EnumDesc)
					name := string(s.Pull(2, "string").(types.String))
					item, ok := desc.Items[name]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					return s.Push(rtypes.EnumItemDesc{EnumItem: item})
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "name", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("EnumItemDesc")},
						},
					}
				},
			},
			"Items": {Method: true,
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EnumDesc)
					items := make(rtypes.Array, 0, len(desc.Items))
					for _, item := range desc.Items {
						items = append(items, rtypes.EnumItemDesc{EnumItem: item})
					}
					sort.Slice(items, func(i, j int) bool {
						return items[i].(rtypes.EnumItemDesc).Name < items[j].(rtypes.EnumItemDesc).Name
					})
					return s.Push(items)
				},
				Dump: func() dump.Value {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Array{T: dt.Prim("EnumItemDesc")}},
						},
					}
				},
			},
			"AddItem": {Method: true,
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EnumDesc)
					item := s.Pull(2, "EnumItemDesc").(rtypes.EnumItemDesc)
					if _, ok := desc.Items[item.Name]; ok {
						return s.Push(types.False)
					}
					if desc.Items == nil {
						desc.Items = map[string]*rbxdump.EnumItem{}
					}
					desc.Items[item.Name] = item.EnumItem
					return s.Push(types.True)
				},
				Dump: func() dump.Value {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "item", Type: dt.Prim("EnumItemDesc")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("bool")},
						},
					}
				},
			},
			"RemoveItem": {Method: true,
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EnumDesc)
					name := string(s.Pull(2, "string").(types.String))
					if _, ok := desc.Items[name]; !ok {
						return s.Push(types.False)
					}
					delete(desc.Items, name)
					return s.Push(types.True)
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
			"Tag": {Method: true,
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EnumDesc)
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
					desc := v.(rtypes.EnumDesc)
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
					desc := v.(rtypes.EnumDesc)
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
					desc := v.(rtypes.EnumDesc)
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
