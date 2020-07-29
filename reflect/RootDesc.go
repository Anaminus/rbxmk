package reflect

import (
	"sort"

	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func RootDesc() Type {
	return Type{
		Name:     "RootDesc",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Members: Members{
			"Class": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.RootDesc)
				name := string(s.Pull(2, "string").(types.String))
				class, ok := desc.Classes[name]
				if !ok {
					return s.Push(rtypes.Nil)
				}
				return s.Push(rtypes.ClassDesc{Class: class})
			}},
			"Classes": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.RootDesc)
				classes := make(rtypes.Array, 0, len(desc.Classes))
				for _, class := range desc.Classes {
					classes = append(classes, rtypes.ClassDesc{Class: class})
				}
				sort.Slice(classes, func(i, j int) bool {
					return classes[i].(rtypes.ClassDesc).Name < classes[j].(rtypes.ClassDesc).Name
				})
				return s.Push(classes)
			}},
			"AddClass": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.RootDesc)
				class := s.Pull(2, "ClassDesc").(rtypes.ClassDesc)
				if _, ok := desc.Classes[class.Name]; ok {
					return s.Push(types.False)
				}
				desc.Classes[class.Name] = class.Class
				return s.Push(types.True)
			}},
			"RemoveClass": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.RootDesc)
				name := string(s.Pull(2, "string").(types.String))
				if _, ok := desc.Classes[name]; !ok {
					return s.Push(types.False)
				}
				delete(desc.Classes, name)
				return s.Push(types.True)
			}},
			"Enum": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.RootDesc)
				name := string(s.Pull(2, "string").(types.String))
				enum, ok := desc.Enums[name]
				if !ok {
					return s.Push(rtypes.Nil)
				}
				return s.Push(rtypes.EnumDesc{Enum: enum})
			}},
			"Enums": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.RootDesc)
				enums := make(rtypes.Array, 0, len(desc.Enums))
				for _, enum := range desc.Enums {
					enums = append(enums, rtypes.EnumDesc{Enum: enum})
				}
				sort.Slice(enums, func(i, j int) bool {
					return enums[i].(rtypes.EnumDesc).Name < enums[j].(rtypes.EnumDesc).Name
				})
				return s.Push(enums)
			}},
			"AddEnum": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.RootDesc)
				enum := s.Pull(2, "ClassDesc").(rtypes.EnumDesc)
				if _, ok := desc.Enums[enum.Name]; ok {
					return s.Push(types.False)
				}
				desc.Enums[enum.Name] = enum.Enum
				return s.Push(types.True)
			}},
			"RemoveEnum": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.RootDesc)
				name := string(s.Pull(2, "string").(types.String))
				if _, ok := desc.Enums[name]; !ok {
					return s.Push(types.False)
				}
				delete(desc.Enums, name)
				return s.Push(types.True)
			}},
		},
	}
}
