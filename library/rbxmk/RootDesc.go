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

func init() { register(RootDesc) }
func RootDesc() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "RootDesc",
		PushTo:   rbxmk.PushPtrTypeTo("RootDesc"),
		PullFrom: rbxmk.PullTypeFrom("RootDesc"),
		Methods: rbxmk.Methods{
			"Class": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(*rtypes.RootDesc)
					name := string(s.Pull(2, "string").(types.String))
					class, ok := desc.Classes[name]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					return s.Push(rtypes.ClassDesc{Class: class})
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "name", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("ClassDesc")},
						},
					}
				},
			},
			"Classes": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(*rtypes.RootDesc)
					classes := make(rtypes.Array, 0, len(desc.Classes))
					for _, class := range desc.Classes {
						classes = append(classes, rtypes.ClassDesc{Class: class})
					}
					sort.Slice(classes, func(i, j int) bool {
						return classes[i].(rtypes.ClassDesc).Name < classes[j].(rtypes.ClassDesc).Name
					})
					return s.Push(classes)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Array{T: dt.Prim("ClassDesc")}},
						},
					}
				},
			},
			"AddClass": {
				Func: func(s rbxmk.State, v types.Value) int {
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
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "class", Type: dt.Prim("ClassDesc")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("bool")},
						},
					}
				},
			},
			"RemoveClass": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(*rtypes.RootDesc)
					name := string(s.Pull(2, "string").(types.String))
					if _, ok := desc.Classes[name]; !ok {
						return s.Push(types.False)
					}
					delete(desc.Classes, name)
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
					}
				},
			},
			"Enum": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(*rtypes.RootDesc)
					name := string(s.Pull(2, "string").(types.String))
					enum, ok := desc.Enums[name]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					return s.Push(rtypes.EnumDesc{Enum: enum})
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "name", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("EnumDesc")},
						},
					}
				},
			},
			"Enums": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(*rtypes.RootDesc)
					enums := make(rtypes.Array, 0, len(desc.Enums))
					for _, enum := range desc.Enums {
						enums = append(enums, rtypes.EnumDesc{Enum: enum})
					}
					sort.Slice(enums, func(i, j int) bool {
						return enums[i].(rtypes.EnumDesc).Name < enums[j].(rtypes.EnumDesc).Name
					})
					return s.Push(enums)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Array{T: dt.Prim("EnumDesc")}},
						},
					}
				},
			},
			"AddEnum": {
				Func: func(s rbxmk.State, v types.Value) int {
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
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "enum", Type: dt.Prim("EnumDesc")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("bool")},
						},
					}
				},
			},
			"RemoveEnum": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(*rtypes.RootDesc)
					name := string(s.Pull(2, "string").(types.String))
					if _, ok := desc.Enums[name]; !ok {
						return s.Push(types.False)
					}
					delete(desc.Enums, name)
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
					}
				},
			},
			"EnumTypes": {
				Func: func(s rbxmk.State, v types.Value) int {
					desc := v.(*rtypes.RootDesc)
					desc.GenerateEnumTypes()
					return s.Push(desc.EnumTypes)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Prim("Enums")},
						},
					}
				},
			},
		},
	}
}
