package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(PropertyDesc) }
func PropertyDesc() Reflector {
	return Reflector{
		Name:     "PropertyDesc",
		PushTo:   rbxmk.PushTypeTo("PropertyDesc"),
		PullFrom: rbxmk.PullTypeFrom("PropertyDesc"),
		Metatable: Metatable{
			"__eq": func(s State) int {
				v := s.Pull(1, "PropertyDesc").(rtypes.PropertyDesc)
				op := s.Pull(2, "PropertyDesc").(rtypes.PropertyDesc)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: Members{
			"Name": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
					return s.Push(types.String(desc.Name))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.PropertyDesc)
					desc.Name = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("string")} },
			},
			"ValueType": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
					valueType := desc.ValueType
					return s.Push(rtypes.TypeDesc{Embedded: valueType})
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.PropertyDesc)
					desc.ValueType = s.Pull(3, "TypeDesc").(rtypes.TypeDesc).Embedded
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("TypeDesc")} },
			},
			"ReadSecurity": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
					return s.Push(types.String(desc.ReadSecurity))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.PropertyDesc)
					desc.ReadSecurity = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("string")} },
			},
			"WriteSecurity": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
					return s.Push(types.String(desc.WriteSecurity))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.PropertyDesc)
					desc.WriteSecurity = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("string")} },
			},
			"CanLoad": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
					return s.Push(types.Bool(desc.CanLoad))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.PropertyDesc)
					desc.CanLoad = bool(s.Pull(3, "bool").(types.Bool))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("bool")} },
			},
			"CanSave": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
					return s.Push(types.Bool(desc.CanSave))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.PropertyDesc)
					desc.CanSave = bool(s.Pull(3, "bool").(types.Bool))
				},
				Dump: func() dump.Value { return dump.Property{ValueType: dt.Prim("bool")} },
			},
			"Tag": Member{Method: true,
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.PropertyDesc)
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
					desc := v.(rtypes.PropertyDesc)
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
					desc := v.(rtypes.PropertyDesc)
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
					desc := v.(rtypes.PropertyDesc)
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
