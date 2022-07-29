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

func init() { register(Desc) }
func Desc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     rtypes.T_Desc,
		PushTo:   rbxmk.PushPtrTypeTo(rtypes.T_Desc),
		PullFrom: rbxmk.PullTypeFrom(rtypes.T_Desc),
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case **rtypes.Desc:
				*p = v.(*rtypes.Desc)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
		Methods: rbxmk.Methods{
			"Copy": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.Desc)
					return s.Push(root.Copy())
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Prim(rtypes.T_Desc)},
						},
						Summary:     "Types/Desc:Methods/Copy/Summary",
						Description: "Types/Desc:Methods/Copy/Description",
					}
				},
			},
			"Class": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.Desc)
					class := string(s.Pull(2, rtypes.T_String).(types.String))
					desc, ok := root.Classes[class]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					return s.Push(rtypes.ClassDesc(*desc))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "class", Type: dt.Prim(rtypes.T_String)},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional(dt.Prim(rtypes.T_ClassDesc))},
						},
						Summary:     "Types/Desc:Methods/Class/Summary",
						Description: "Types/Desc:Methods/Class/Description",
					}
				},
			},
			"Classes": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.Desc)
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
							{Type: dt.Array(dt.Prim(rtypes.T_String))},
						},
						Summary:     "Types/Desc:Methods/Classes/Summary",
						Description: "Types/Desc:Methods/Classes/Description",
					}
				},
			},
			"SetClass": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.Desc)
					class := string(s.Pull(2, rtypes.T_String).(types.String))
					desc := s.PullOpt(3, rtypes.Nil, rtypes.T_ClassDesc)
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
							{Name: "class", Type: dt.Prim(rtypes.T_String)},
							{Name: "desc", Type: dt.Optional(dt.Prim(rtypes.T_ClassDesc))},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(rtypes.T_Bool)},
						},
						Summary:     "Types/Desc:Methods/SetClass/Summary",
						Description: "Types/Desc:Methods/SetClass/Description",
					}
				},
			},
			"ClassTag": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.Desc)
					class := string(s.Pull(2, rtypes.T_String).(types.String))
					tag := string(s.Pull(3, rtypes.T_String).(types.String))
					desc, ok := root.Classes[class]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					return s.Push(types.Bool(desc.GetTag(tag)))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "class", Type: dt.Prim(rtypes.T_String)},
							{Name: "tag", Type: dt.Prim(rtypes.T_String)},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional(dt.Prim(rtypes.T_Bool))},
						},
						Summary:     "Types/Desc:Methods/ClassTag/Summary",
						Description: "Types/Desc:Methods/ClassTag/Description",
					}
				},
			},
			"Member": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.Desc)
					class := string(s.Pull(2, rtypes.T_String).(types.String))
					member := string(s.Pull(3, rtypes.T_String).(types.String))
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
							{Name: "class", Type: dt.Prim(rtypes.T_String)},
							{Name: "member", Type: dt.Prim(rtypes.T_String)},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional(dt.Prim(rtypes.T_MemberDesc))},
						},
						Summary:     "Types/Desc:Methods/Member/Summary",
						Description: "Types/Desc:Methods/Member/Description",
					}
				},
			},
			"Members": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.Desc)
					class := string(s.Pull(2, rtypes.T_String).(types.String))
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
							{Type: dt.Optional(dt.Array(dt.Prim(rtypes.T_String)))},
						},
						Summary:     "Types/Desc:Methods/Members/Summary",
						Description: "Types/Desc:Methods/Members/Description",
					}
				},
			},
			"SetMember": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.Desc)
					class := string(s.Pull(2, rtypes.T_String).(types.String))
					member := string(s.Pull(3, rtypes.T_String).(types.String))
					memberDesc := s.PullOpt(4, rtypes.Nil, rtypes.T_MemberDesc)
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
							{Name: "class", Type: dt.Prim(rtypes.T_String)},
							{Name: "member", Type: dt.Prim(rtypes.T_String)},
							{Name: "desc", Type: dt.Optional(dt.Prim(rtypes.T_ClassDesc))},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional(dt.Prim(rtypes.T_Bool))},
						},
						Summary:     "Types/Desc:Methods/SetMember/Summary",
						Description: "Types/Desc:Methods/SetMember/Description",
					}
				},
			},
			"MemberTag": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.Desc)
					class := string(s.Pull(2, rtypes.T_String).(types.String))
					member := string(s.Pull(3, rtypes.T_String).(types.String))
					tag := string(s.Pull(4, rtypes.T_String).(types.String))
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
							{Name: "class", Type: dt.Prim(rtypes.T_String)},
							{Name: "member", Type: dt.Prim(rtypes.T_String)},
							{Name: "tag", Type: dt.Prim(rtypes.T_String)},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional(dt.Prim(rtypes.T_Bool))},
						},
						Summary:     "Types/Desc:Methods/MemberTag/Summary",
						Description: "Types/Desc:Methods/MemberTag/Description",
					}
				},
			},
			"Enum": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.Desc)
					enum := string(s.Pull(2, rtypes.T_String).(types.String))
					desc, ok := root.Enums[enum]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					return s.Push(rtypes.EnumDesc(*desc))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "enum", Type: dt.Prim(rtypes.T_String)},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional(dt.Prim(rtypes.T_EnumDesc))},
						},
						Summary:     "Types/Desc:Methods/Enum/Summary",
						Description: "Types/Desc:Methods/Enum/Description",
					}
				},
			},
			"Enums": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.Desc)
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
							{Type: dt.Array(dt.Prim(rtypes.T_String))},
						},
						Summary:     "Types/Desc:Methods/Enums/Summary",
						Description: "Types/Desc:Methods/Enums/Description",
					}
				},
			},
			"SetEnum": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.Desc)
					enum := string(s.Pull(2, rtypes.T_String).(types.String))
					desc := s.PullOpt(3, rtypes.Nil, rtypes.T_EnumDesc)
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
							{Name: "enum", Type: dt.Prim(rtypes.T_String)},
							{Name: "desc", Type: dt.Optional(dt.Prim(rtypes.T_EnumDesc))},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim(rtypes.T_Bool)},
						},
						Summary:     "Types/Desc:Methods/SetEnum/Summary",
						Description: "Types/Desc:Methods/SetEnum/Description",
					}
				},
			},
			"EnumTag": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.Desc)
					enum := string(s.Pull(2, rtypes.T_String).(types.String))
					tag := string(s.Pull(3, rtypes.T_String).(types.String))
					desc, ok := root.Enums[enum]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					return s.Push(types.Bool(desc.GetTag(tag)))
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "enum", Type: dt.Prim(rtypes.T_String)},
							{Name: "tag", Type: dt.Prim(rtypes.T_String)},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional(dt.Prim(rtypes.T_Bool))},
						},
						Summary:     "Types/Desc:Methods/EnumTag/Summary",
						Description: "Types/Desc:Methods/EnumTag/Description",
					}
				},
			},
			"EnumItem": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.Desc)
					enum := string(s.Pull(2, rtypes.T_String).(types.String))
					item := string(s.Pull(3, rtypes.T_String).(types.String))
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
							{Name: "enum", Type: dt.Prim(rtypes.T_String)},
							{Name: "item", Type: dt.Prim(rtypes.T_String)},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional(dt.Prim(rtypes.T_EnumItemDesc))},
						},
						Summary:     "Types/Desc:Methods/EnumItem/Summary",
						Description: "Types/Desc:Methods/EnumItem/Description",
					}
				},
			},
			"EnumItems": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.Desc)
					enum := string(s.Pull(2, rtypes.T_String).(types.String))
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
							{Type: dt.Optional(dt.Array(dt.Prim(rtypes.T_String)))},
						},
						Summary:     "Types/Desc:Methods/EnumItems/Summary",
						Description: "Types/Desc:Methods/EnumItems/Description",
					}
				},
			},
			"SetEnumItem": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.Desc)
					enum := string(s.Pull(2, rtypes.T_String).(types.String))
					item := string(s.Pull(3, rtypes.T_String).(types.String))
					itemDesc := s.PullOpt(4, rtypes.Nil, rtypes.T_EnumItemDesc)
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
							{Name: "enum", Type: dt.Prim(rtypes.T_String)},
							{Name: "item", Type: dt.Prim(rtypes.T_String)},
							{Name: "desc", Type: dt.Optional(dt.Prim(rtypes.T_EnumItemDesc))},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional(dt.Prim(rtypes.T_Bool))},
						},
						Summary:     "Types/Desc:Methods/SetEnumItem/Summary",
						Description: "Types/Desc:Methods/SetEnumItem/Description",
					}
				},
			},
			"EnumItemTag": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.Desc)
					enum := string(s.Pull(2, rtypes.T_String).(types.String))
					item := string(s.Pull(3, rtypes.T_String).(types.String))
					tag := string(s.Pull(4, rtypes.T_String).(types.String))
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
							{Name: "enum", Type: dt.Prim(rtypes.T_String)},
							{Name: "item", Type: dt.Prim(rtypes.T_String)},
							{Name: "tag", Type: dt.Prim(rtypes.T_String)},
						},
						Returns: dump.Parameters{
							{Type: dt.Optional(dt.Prim(rtypes.T_Bool))},
						},
						Summary:     "Types/Desc:Methods/EnumItemTag/Summary",
						Description: "Types/Desc:Methods/EnumItemTag/Description",
					}
				},
			},
			"EnumTypes": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.Desc)
					root.GenerateEnumTypes()
					return s.Push(root.EnumTypes)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Prim(rtypes.T_Enums)},
						},
						Summary:     "Types/Desc:Methods/EnumTypes/Summary",
						Description: "Types/Desc:Methods/EnumTypes/Description",
					}
				},
			},
			"Diff": {
				Func: func(s rbxmk.State, v types.Value) int {
					prev := v.(*rtypes.Desc).Root
					var next *rbxdump.Root
					switch v := s.PullAnyOf(2, rtypes.T_Desc, rtypes.T_Nil).(type) {
					case rtypes.NilType:
					case *rtypes.Desc:
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
							{Name: "next", Type: dt.Optional(dt.Prim(rtypes.T_Desc))},
						},
						Returns: dump.Parameters{
							{Name: "diff", Type: dt.Prim(rtypes.T_DescActions)},
						},
						Summary:     "Types/Desc:Methods/Diff/Summary",
						Description: "Types/Desc:Methods/Diff/Description",
					}
				},
			},
			"Patch": {
				Func: func(s rbxmk.State, v types.Value) int {
					root := v.(*rtypes.Desc).Root
					descActions := s.Pull(2, rtypes.T_DescActions).(rtypes.DescActions)
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
							{Name: "actions", Type: dt.Prim(rtypes.T_DescActions)},
						},
						Summary:     "Types/Desc:Methods/Patch/Summary",
						Description: "Types/Desc:Methods/Patch/Description",
					}
				},
			},
		},
		Constructors: rbxmk.Constructors{
			"new": rbxmk.Constructor{
				Func: func(s rbxmk.State) int {
					return s.Push(&rtypes.Desc{Root: &rbxdump.Root{
						Classes: make(map[string]*rbxdump.Class),
						Enums:   make(map[string]*rbxdump.Enum),
					}})
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						dump.Function{
							Returns: dump.Parameters{
								{Type: dt.Prim(rtypes.T_Desc)},
							},
							Summary:     "Types/Desc:Constructors/new/Summary",
							Description: "Types/Desc:Constructors/new/Description",
						},
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Category:    "rbxmk",
				Summary:     "Types/Desc:Summary",
				Description: "Types/Desc:Description",
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
