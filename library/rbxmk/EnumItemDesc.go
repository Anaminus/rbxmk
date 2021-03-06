package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(EnumItemDesc) }
func EnumItemDesc() Reflector {
	return Reflector{
		Name:     "EnumItemDesc",
		PushTo:   rbxmk.PushTypeTo("EnumItemDesc"),
		PullFrom: rbxmk.PullTypeFrom("EnumItemDesc"),
		Metatable: Metatable{
			"__eq": func(s State) int {
				v := s.Pull(1, "EnumItemDesc").(rtypes.EnumItemDesc)
				op := s.Pull(2, "EnumItemDesc").(rtypes.EnumItemDesc)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: Members{
			"Name": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.EnumItemDesc)
					return s.Push(types.String(desc.Name))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.EnumItemDesc)
					desc.Name = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("string")} },
			},
			"Value": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.EnumItemDesc)
					return s.Push(types.Int(desc.Value))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.EnumItemDesc)
					desc.Value = int(s.Pull(3, "int").(types.Int))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("int")} },
			},
			"Index": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.EnumItemDesc)
					return s.Push(types.Int(desc.Index))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.EnumItemDesc)
					desc.Index = int(s.Pull(3, "int").(types.Int))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("int")} },
			},
			"Tag": Member{Method: true,
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.EnumItemDesc)
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
			"Tags": Member{Method: true,
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.EnumItemDesc)
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
			"SetTag": Member{Method: true,
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.EnumItemDesc)
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
			"UnsetTag": Member{Method: true,
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.EnumItemDesc)
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
