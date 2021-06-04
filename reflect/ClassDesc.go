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

func init() { register(ClassDesc) }
func ClassDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "ClassDesc",
		PushTo:   rbxmk.PushPtrTypeTo("ClassDesc"),
		PullFrom: rbxmk.PullTypeFrom("ClassDesc"),
		Properties: rbxmk.Properties{
			"Name": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.ClassDesc)
					return s.Push(types.String(desc.Name))
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.ClassDesc)
					desc.Name = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						Summary:     "Types/ClassDesc:Properties/Name/Summary",
						Description: "Types/ClassDesc:Properties/Name/Description",
					}
				},
			},
			"Superclass": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.ClassDesc)
					return s.Push(types.String(desc.Superclass))
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.ClassDesc)
					desc.Superclass = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						Summary:     "Types/ClassDesc:Properties/Superclass/Summary",
						Description: "Types/ClassDesc:Properties/Superclass/Description",
					}
				},
			},
			"MemoryCategory": {
				Get: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.ClassDesc)
					return s.Push(types.String(desc.MemoryCategory))
				},
				Set: func(s rbxmk.State, v types.Value) {
					desc := v.(rtypes.ClassDesc)
					desc.MemoryCategory = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						Summary:     "Types/ClassDesc:Properties/MemoryCategory/Summary",
						Description: "Types/ClassDesc:Properties/MemoryCategory/Description",
					}
				},
			},
		},
		Methods: rbxmk.Methods{
			"Member": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.ClassDesc)
					name := string(s.Pull(2, "string").(types.String))
					member, ok := desc.Members[name]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					return s.Push(rtypes.NewMemberDesc(member))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "name", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("MemberDesc")},
						},
						Summary:     "Types/ClassDesc:Methods/Member/Summary",
						Description: "Types/ClassDesc:Methods/Member/Description",
					}
				},
			},
			"Members": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.ClassDesc)
					members := make(rtypes.Array, 0, len(desc.Members))
					for _, member := range desc.Members {
						members = append(members, rtypes.NewMemberDesc(member))
					}
					sort.Slice(members, func(i, j int) bool {
						return members[i].(rbxdump.Member).MemberName() < members[j].(rbxdump.Member).MemberName()
					})
					return s.Push(members)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Array{T: dt.Prim("MemberDesc")}},
						},
						Summary:     "Types/ClassDesc:Methods/Members/Summary",
						Description: "Types/ClassDesc:Methods/Members/Description",
					}
				},
			},
			"AddMember": {
				Func: func(s rbxmk.State, v types.Value) int {
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
					default:
						return s.ReflectorError(2)
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
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "member", Type: dt.Prim("MemberDesc")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("bool")},
						},
						Summary:     "Types/ClassDesc:Methods/AddMember/Summary",
						Description: "Types/ClassDesc:Methods/AddMember/Description",
					}
				},
			},
			"RemoveMember": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.ClassDesc)
					name := string(s.Pull(2, "string").(types.String))
					if _, ok := desc.Members[name]; !ok {
						return s.Push(types.False)
					}
					delete(desc.Members, name)
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
						Summary:     "Types/ClassDesc:Methods/RemoveMember/Summary",
						Description: "Types/ClassDesc:Methods/RemoveMember/Description",
					}
				},
			},
			"Tag": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.ClassDesc)
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
						Summary:     "Types/ClassDesc:Methods/Tag/Summary",
						Description: "Types/ClassDesc:Methods/Tag/Description",
					}
				},
			},
			"Tags": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.ClassDesc)
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
						Summary:     "Types/ClassDesc:Methods/Tags/Summary",
						Description: "Types/ClassDesc:Methods/Tags/Description",
					}
				},
			},
			"SetTag": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.ClassDesc)
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
						Summary:     "Types/ClassDesc:Methods/SetTag/Summary",
						Description: "Types/ClassDesc:Methods/SetTag/Description",
					}
				},
			},
			"UnsetTag": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(rtypes.ClassDesc)
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
						Summary:     "Types/ClassDesc:Methods/UnsetTag/Summary",
						Description: "Types/ClassDesc:Methods/UnsetTag/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": rbxmk.Constructor{
				Func: func(s rbxmk.State) int {
					return s.Push(rtypes.ClassDesc{Class: &rbxdump.Class{
						Members: make(map[string]rbxdump.Member),
					}})
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						dump.Function{
							Returns: dump.Parameters{
								{Type: dt.Prim("ClassDesc")},
							},
							Summary:     "Types/ClassDesc:Constructors/new/Summary",
							Description: "Types/ClassDesc:Constructors/new/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/ClassDesc:Summary",
				Description: "Types/ClassDesc:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Array,
			Bool,
			CallbackDesc,
			EventDesc,
			FunctionDesc,
			PropertyDesc,
			String,
		},
	}
}
