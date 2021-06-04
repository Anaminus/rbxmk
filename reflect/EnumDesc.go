package reflect

import (
	"sort"

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
		Properties: rbxmk.Properties{
			"Name": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EnumDesc)
					return s.Push(types.String(desc.Name))
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.EnumDesc)
					desc.Name = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						Summary:     "Types/EnumDesc:Properties/Name/Summary",
						Description: "Types/EnumDesc:Properties/Name/Description",
					}
				},
			},
		},
		Methods: rbxmk.Methods{
			"Item": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EnumDesc)
					name := string(s.Pull(2, "string").(types.String))
					item, ok := desc.Items[name]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					return s.Push(rtypes.EnumItemDesc{EnumItem: item})
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "name", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("EnumItemDesc")},
						},
						Summary:     "Types/EnumDesc:Methods/Item/Summary",
						Description: "Types/EnumDesc:Methods/Item/Description",
					}
				},
			},
			"Items": {
				Func: func(s rbxmk.State, v types.Value) int {
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
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Array{T: dt.Prim("EnumItemDesc")}},
						},
						Summary:     "Types/EnumDesc:Methods/Items/Summary",
						Description: "Types/EnumDesc:Methods/Items/Description",
					}
				},
			},
			"AddItem": {
				Func: func(s rbxmk.State, v types.Value) int {
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
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "item", Type: dt.Prim("EnumItemDesc")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("bool")},
						},
						Summary:     "Types/EnumDesc:Methods/AddItem/Summary",
						Description: "Types/EnumDesc:Methods/AddItem/Description",
					}
				},
			},
			"RemoveItem": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EnumDesc)
					name := string(s.Pull(2, "string").(types.String))
					if _, ok := desc.Items[name]; !ok {
						return s.Push(types.False)
					}
					delete(desc.Items, name)
					return s.Push(types.True)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "name", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("bool")},
						},
						Summary:     "Types/EnumDesc:Methods/RemoveItem/Summary",
						Description: "Types/EnumDesc:Methods/RemoveItem/Description",
					}
				},
			},
			"Tag": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EnumDesc)
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
						Summary:     "Types/EnumDesc:Methods/Tag/Summary",
						Description: "Types/EnumDesc:Methods/Tag/Description",
					}
				},
			},
			"Tags": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EnumDesc)
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
						Summary:     "Types/EnumDesc:Methods/Tags/Summary",
						Description: "Types/EnumDesc:Methods/Tags/Description",
					}
				},
			},
			"SetTag": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EnumDesc)
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
						Summary:     "Types/EnumDesc:Methods/SetTag/Summary",
						Description: "Types/EnumDesc:Methods/SetTag/Description",
					}
				},
			},
			"UnsetTag": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.EnumDesc)
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
						Summary:     "Types/EnumDesc:Methods/UnsetTag/Summary",
						Description: "Types/EnumDesc:Methods/UnsetTag/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": rbxmk.Constructor{
				Func: func(s rbxmk.State) int {
					return s.Push(rtypes.EnumDesc{Enum: &rbxdump.Enum{
						Items: make(map[string]*rbxdump.EnumItem),
					}})
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						dump.Function{
							Returns: dump.Parameters{
								{Type: dt.Prim("EnumDesc")},
							},
							Summary:     "Types/EnumDesc:Constructors/new/Summary",
							Description: "Types/EnumDesc:Constructors/new/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/EnumDesc:Summary",
				Description: "Types/EnumDesc:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Array,
			Bool,
			EnumItemDesc,
			String,
		},
	}
}
