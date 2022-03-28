package reflect

import (
	"sort"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/rbxdump/diff"
	"github.com/robloxapi/types"
)

func init() { register(RootDesc) }
func RootDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "RootDesc",
		PushTo:   rbxmk.PushPtrTypeTo("RootDesc"),
		PullFrom: rbxmk.PullTypeFrom("RootDesc"),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case **rtypes.RootDesc:
				*p = v.(*rtypes.RootDesc)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Methods: rbxmk.Methods{
			"Copy": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.RootDesc)
					return s.Push(root.Copy())
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Prim("RootDesc")},
						},
						Summary:     "Types/RootDesc:Methods/Copy/Summary",
						Description: "Types/RootDesc:Methods/Copy/Description",
					}
				},
			},
			"Class": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.RootDesc)
					class := string(s.Pull(2, "string").(types.String))
					desc, ok := root.Classes[class]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					return s.Push(rtypes.ClassDesc(*desc))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "class", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("ClassDesc")}},
						},
						Summary:     "Types/RootDesc:Methods/Class/Summary",
						Description: "Types/RootDesc:Methods/Class/Description",
					}
				},
			},
			"Classes": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.RootDesc)
					classes := make(rtypes.Array, 0, len(root.Classes))
					for _, class := range root.Classes {
						classes = append(classes, types.String(class.Name))
					}
					sort.Slice(classes, func(i, j int) bool {
						return classes[i].(types.String) < classes[j].(types.String)
					})
					return s.Push(classes)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Array{T: dt.Prim("string")}},
						},
						Summary:     "Types/RootDesc:Methods/Classes/Summary",
						Description: "Types/RootDesc:Methods/Classes/Description",
					}
				},
			},
			"SetClass": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.RootDesc)
					class := string(s.Pull(2, "string").(types.String))
					desc := s.PullOpt(3, rtypes.Nil, "ClassDesc")
					switch desc := desc.(type) {
					case rtypes.NilType:
						if _, ok := root.Classes[class]; ok {
							delete(root.Classes, class)
							return s.Push(types.True)
						}
						return s.Push(types.False)
					case rtypes.ClassDesc:
						desc.Name = class
						if root.Classes == nil {
							root.Classes = map[string]*rbxdump.Class{}
						}
						root.Classes[class] = (*rbxdump.Class)(&desc)
					default:
						return s.ReflectorError(3)
					}
					return s.Push(types.True)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "class", Type: dt.Prim("string")},
							{Name: "desc", Type: dt.Optional{T: dt.Prim("ClassDesc")}},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("bool")},
						},
						Summary:     "Types/RootDesc:Methods/SetClass/Summary",
						Description: "Types/RootDesc:Methods/SetClass/Description",
					}
				},
			},
			"ClassTag": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.RootDesc)
					class := string(s.Pull(2, "string").(types.String))
					tag := string(s.Pull(3, "string").(types.String))
					desc, ok := root.Classes[class]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					return s.Push(types.Bool(desc.GetTag(tag)))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "class", Type: dt.Prim("string")},
							{Name: "tag", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("bool")}},
						},
						Summary:     "Types/RootDesc:Methods/ClassTag/Summary",
						Description: "Types/RootDesc:Methods/ClassTag/Description",
					}
				},
			},
			"Member": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.RootDesc)
					class := string(s.Pull(2, "string").(types.String))
					member := string(s.Pull(3, "string").(types.String))
					classDesc, ok := root.Classes[class]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					memberDesc, ok := classDesc.Members[member]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					return s.Push(rtypes.MemberDesc{Member: memberDesc})
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "class", Type: dt.Prim("string")},
							{Name: "member", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("MemberDesc")}},
						},
						Summary:     "Types/RootDesc:Methods/Member/Summary",
						Description: "Types/RootDesc:Methods/Member/Description",
					}
				},
			},
			"Members": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.RootDesc)
					class := string(s.Pull(2, "string").(types.String))
					classDesc, ok := root.Classes[class]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					members := make(rtypes.Array, 0, len(classDesc.Members))
					for member := range classDesc.Members {
						members = append(members, types.String(member))
					}
					sort.Slice(members, func(i, j int) bool {
						return members[i].(types.String) < members[j].(types.String)
					})
					return s.Push(members)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Array{T: dt.Prim("string")}}},
						},
						Summary:     "Types/RootDesc:Methods/Members/Summary",
						Description: "Types/RootDesc:Methods/Members/Description",
					}
				},
			},
			"SetMember": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.RootDesc)
					class := string(s.Pull(2, "string").(types.String))
					member := string(s.Pull(3, "string").(types.String))
					memberDesc := s.PullOpt(4, rtypes.Nil, "MemberDesc")
					classDesc, ok := root.Classes[class]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					switch memberDesc := memberDesc.(type) {
					case rtypes.NilType:
						if _, ok := classDesc.Members[member]; ok {
							delete(classDesc.Members, member)
							return s.Push(types.True)
						}
						return s.Push(types.False)
					case rtypes.MemberDesc:
						switch m := memberDesc.Member.(type) {
						case *rbxdump.Property:
							m.Name = member
						case *rbxdump.Function:
							m.Name = member
						case *rbxdump.Event:
							m.Name = member
						case *rbxdump.Callback:
							m.Name = member
						}
						if classDesc.Members == nil {
							classDesc.Members = map[string]rbxdump.Member{}
						}
						classDesc.Members[member] = memberDesc.Member
					default:
						return s.ReflectorError(4)
					}
					return s.Push(types.True)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "class", Type: dt.Prim("string")},
							{Name: "member", Type: dt.Prim("string")},
							{Name: "desc", Type: dt.Optional{T: dt.Prim("ClassDesc")}},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("bool")}},
						},
						Summary:     "Types/RootDesc:Methods/SetMember/Summary",
						Description: "Types/RootDesc:Methods/SetMember/Description",
					}
				},
			},
			"MemberTag": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.RootDesc)
					class := string(s.Pull(2, "string").(types.String))
					member := string(s.Pull(3, "string").(types.String))
					tag := string(s.Pull(4, "string").(types.String))
					classDesc, ok := root.Classes[class]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					memberDesc, ok := classDesc.Members[member]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					return s.Push(types.Bool(memberDesc.GetTag(tag)))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "class", Type: dt.Prim("string")},
							{Name: "member", Type: dt.Prim("string")},
							{Name: "tag", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("bool")}},
						},
						Summary:     "Types/RootDesc:Methods/MemberTag/Summary",
						Description: "Types/RootDesc:Methods/MemberTag/Description",
					}
				},
			},
			"Enum": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.RootDesc)
					enum := string(s.Pull(2, "string").(types.String))
					desc, ok := root.Enums[enum]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					return s.Push(rtypes.EnumDesc(*desc))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "enum", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("EnumDesc")}},
						},
						Summary:     "Types/RootDesc:Methods/Enum/Summary",
						Description: "Types/RootDesc:Methods/Enum/Description",
					}
				},
			},
			"Enums": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.RootDesc)
					enums := make(rtypes.Array, 0, len(root.Enums))
					for _, enum := range root.Enums {
						enums = append(enums, types.String(enum.Name))
					}
					sort.Slice(enums, func(i, j int) bool {
						return enums[i].(types.String) < enums[j].(types.String)
					})
					return s.Push(enums)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Array{T: dt.Prim("string")}},
						},
						Summary:     "Types/RootDesc:Methods/Enums/Summary",
						Description: "Types/RootDesc:Methods/Enums/Description",
					}
				},
			},
			"SetEnum": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.RootDesc)
					enum := string(s.Pull(2, "string").(types.String))
					desc := s.PullOpt(3, rtypes.Nil, "EnumDesc")
					switch desc := desc.(type) {
					case rtypes.NilType:
						if _, ok := root.Enums[enum]; ok {
							delete(root.Enums, enum)
							return s.Push(types.True)
						}
						return s.Push(types.False)
					case rtypes.EnumDesc:
						desc.Name = enum
						if root.Enums == nil {
							root.Enums = map[string]*rbxdump.Enum{}
						}
						root.Enums[enum] = (*rbxdump.Enum)(&desc)
					default:
						return s.ReflectorError(3)
					}
					return s.Push(types.True)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "enum", Type: dt.Prim("string")},
							{Name: "desc", Type: dt.Optional{T: dt.Prim("EnumDesc")}},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("bool")},
						},
						Summary:     "Types/RootDesc:Methods/SetEnum/Summary",
						Description: "Types/RootDesc:Methods/SetEnum/Description",
					}
				},
			},
			"EnumTag": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.RootDesc)
					enum := string(s.Pull(2, "string").(types.String))
					tag := string(s.Pull(3, "string").(types.String))
					desc, ok := root.Enums[enum]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					return s.Push(types.Bool(desc.GetTag(tag)))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "enum", Type: dt.Prim("string")},
							{Name: "tag", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("bool")}},
						},
						Summary:     "Types/RootDesc:Methods/EnumTag/Summary",
						Description: "Types/RootDesc:Methods/EnumTag/Description",
					}
				},
			},
			"EnumItem": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.RootDesc)
					enum := string(s.Pull(2, "string").(types.String))
					item := string(s.Pull(3, "string").(types.String))
					enumDesc, ok := root.Enums[enum]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					itemDesc, ok := enumDesc.Items[item]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					return s.Push(rtypes.EnumItemDesc(*itemDesc))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "enum", Type: dt.Prim("string")},
							{Name: "item", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("EnumItemDesc")}},
						},
						Summary:     "Types/RootDesc:Methods/EnumItem/Summary",
						Description: "Types/RootDesc:Methods/EnumItem/Description",
					}
				},
			},
			"EnumItems": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.RootDesc)
					enum := string(s.Pull(2, "string").(types.String))
					enumDesc, ok := root.Enums[enum]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					items := make(rtypes.Array, 0, len(enumDesc.Items))
					for item := range enumDesc.Items {
						items = append(items, types.String(item))
					}
					sort.Slice(items, func(i, j int) bool {
						return items[i].(types.String) < items[j].(types.String)
					})
					return s.Push(items)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Array{T: dt.Prim("string")}}},
						},
						Summary:     "Types/RootDesc:Methods/EnumItems/Summary",
						Description: "Types/RootDesc:Methods/EnumItems/Description",
					}
				},
			},
			"SetEnumItem": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.RootDesc)
					enum := string(s.Pull(2, "string").(types.String))
					item := string(s.Pull(3, "string").(types.String))
					itemDesc := s.PullOpt(4, rtypes.Nil, "EnumItemDesc")
					enumDesc, ok := root.Enums[enum]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					switch itemDesc := itemDesc.(type) {
					case rtypes.NilType:
						if _, ok := enumDesc.Items[item]; ok {
							delete(enumDesc.Items, item)
							return s.Push(types.True)
						}
						return s.Push(types.False)
					case rtypes.EnumItemDesc:
						itemDesc.Name = item
						if enumDesc.Items == nil {
							enumDesc.Items = map[string]*rbxdump.EnumItem{}
						}
						enumDesc.Items[item] = (*rbxdump.EnumItem)(&itemDesc)
					default:
						return s.ReflectorError(4)
					}
					return s.Push(types.True)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "enum", Type: dt.Prim("string")},
							{Name: "item", Type: dt.Prim("string")},
							{Name: "desc", Type: dt.Optional{T: dt.Prim("EnumItemDesc")}},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("bool")}},
						},
						Summary:     "Types/RootDesc:Methods/SetEnumItem/Summary",
						Description: "Types/RootDesc:Methods/SetEnumItem/Description",
					}
				},
			},
			"EnumItemTag": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.RootDesc)
					enum := string(s.Pull(2, "string").(types.String))
					item := string(s.Pull(3, "string").(types.String))
					tag := string(s.Pull(4, "string").(types.String))
					enumDesc, ok := root.Enums[enum]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					itemDesc, ok := enumDesc.Items[item]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					return s.Push(types.Bool(itemDesc.GetTag(tag)))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "enum", Type: dt.Prim("string")},
							{Name: "item", Type: dt.Prim("string")},
							{Name: "tag", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional{T: dt.Prim("bool")}},
						},
						Summary:     "Types/RootDesc:Methods/EnumItemTag/Summary",
						Description: "Types/RootDesc:Methods/EnumItemTag/Description",
					}
				},
			},
			"EnumTypes": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.RootDesc)
					root.GenerateEnumTypes()
					return s.Push(root.EnumTypes)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Prim("Enums")},
						},
						Summary:     "Types/RootDesc:Methods/EnumTypes/Summary",
						Description: "Types/RootDesc:Methods/EnumTypes/Description",
					}
				},
			},
			"Diff": {
				Func: func(s rbxmk.State, v types.Value) int {
					prev := v.(*rtypes.RootDesc).Root
					var next *rbxdump.Root
					switch v := s.PullAnyOf(2, "RootDesc", "nil").(type) {
					case rtypes.NilType:
					case *rtypes.RootDesc:
						next = v.Root
					default:
						return s.ReflectorError(1)
					}
					actions := diff.Diff{Prev: prev, Next: next}.Diff()
					descActions := make(rtypes.DescActions, len(actions))
					for i, action := range actions {
						descAction := rtypes.DescAction{Action: action}
						descActions[i] = &descAction
					}
					return s.Push(descActions)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "next", Type: dt.Optional{T: dt.Prim("RootDesc")}},
						},
						Returns: dump.Parameters{
							{Name: "diff", Type: dt.Prim("DescActions")},
						},
						Summary:     "Types/RootDesc:Methods/Diff/Summary",
						Description: "Types/RootDesc:Methods/Diff/Description",
					}
				},
			},
			"Patch": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.RootDesc).Root
					descActions := s.Pull(2, "DescActions").(rtypes.DescActions)
					actions := make([]diff.Action, len(descActions))
					for i, action := range descActions {
						actions[i] = action.Action
					}
					diff.Patch{Root: root}.Patch(actions)
					return 0
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "actions", Type: dt.Prim("DescActions")},
						},
						Summary:     "Types/RootDesc:Methods/Patch/Summary",
						Description: "Types/RootDesc:Methods/Patch/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": rbxmk.Constructor{
				Func: func(s rbxmk.State) int {
					return s.Push(&rtypes.RootDesc{Root: &rbxdump.Root{
						Classes: make(map[string]*rbxdump.Class),
						Enums:   make(map[string]*rbxdump.Enum),
					}})
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						dump.Function{
							Returns: dump.Parameters{
								{Type: dt.Prim("RootDesc")},
							},
							Summary:     "Types/RootDesc:Constructors/new/Summary",
							Description: "Types/RootDesc:Constructors/new/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/RootDesc:Summary",
				Description: "Types/RootDesc:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Array,
			Bool,
			ClassDesc,
			DescActions,
			EnumDesc,
			EnumItemDesc,
			Enums,
			MemberDesc,
			Nil,
			String,
		},
	}
}
