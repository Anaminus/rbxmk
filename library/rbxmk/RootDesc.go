package reflect

import (
	"sort"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/types"
)

func init() { register(RootDesc) }
func RootDesc() Reflector {
	return Reflector{
		Name:     "RootDesc",
		PushTo:   rbxmk.PushTypeTo("RootDesc"),
		PullFrom: rbxmk.PullTypeFrom("RootDesc"),
		Metatable: Metatable{
			"__eq": func(s State) int {
				v := s.Pull(1, "RootDesc").(*rtypes.RootDesc)
				op := s.Pull(2, "RootDesc").(*rtypes.RootDesc)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: Members{
			"Class": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(*rtypes.RootDesc)
				name := string(s.Pull(2, "string").(types.String))
				class, ok := desc.Classes[name]
				if !ok {
					return s.Push(rtypes.Nil)
				}
				return s.Push(rtypes.ClassDesc{Class: class})
			}},
			"Classes": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(*rtypes.RootDesc)
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
				desc := v.(*rtypes.RootDesc)
				class := s.Pull(2, "ClassDesc").(rtypes.ClassDesc)
				if _, ok := desc.Classes[class.Name]; ok {
					return s.Push(types.False)
				}
				if desc.Classes == nil {
					desc.Classes = map[string]*rbxdump.Class{}
				}
				desc.Classes[class.Name] = class.Class
				return s.Push(types.True)
			}},
			"RemoveClass": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(*rtypes.RootDesc)
				name := string(s.Pull(2, "string").(types.String))
				if _, ok := desc.Classes[name]; !ok {
					return s.Push(types.False)
				}
				delete(desc.Classes, name)
				return s.Push(types.True)
			}},
			"Enum": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(*rtypes.RootDesc)
				name := string(s.Pull(2, "string").(types.String))
				enum, ok := desc.Enums[name]
				if !ok {
					return s.Push(rtypes.Nil)
				}
				return s.Push(rtypes.EnumDesc{Enum: enum})
			}},
			"Enums": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(*rtypes.RootDesc)
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
				desc := v.(*rtypes.RootDesc)
				enum := s.Pull(2, "EnumDesc").(rtypes.EnumDesc)
				if _, ok := desc.Enums[enum.Name]; ok {
					return s.Push(types.False)
				}
				if desc.Enums == nil {
					desc.Enums = map[string]*rbxdump.Enum{}
				}
				desc.Enums[enum.Name] = enum.Enum
				return s.Push(types.True)
			}},
			"RemoveEnum": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(*rtypes.RootDesc)
				name := string(s.Pull(2, "string").(types.String))
				if _, ok := desc.Enums[name]; !ok {
					return s.Push(types.False)
				}
				delete(desc.Enums, name)
				return s.Push(types.True)
			}},
			"EnumTypes": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(*rtypes.RootDesc)
				desc.GenerateEnumTypes()
				return s.Push(desc.EnumTypes)
			}},
		},
	}
}
