package reflect

import (
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/rbxdump/diff"
	"github.com/robloxapi/types"
)

func init() { register(DescAction) }
func DescAction() rbxmk.Reflector {
	return rbxmk.Reflector{
		Name:     "DescAction",
		PushTo:   rbxmk.PushPtrTypeTo("DescAction"),
		PullFrom: rbxmk.PullTypeFrom("DescAction"),
		Metatable: rbxmk.Metatable{
			"__tostring": func(s rbxmk.State) int {
				v := s.Pull(1, "DescAction").(*rtypes.DescAction)
				s.L.Push(lua.LString(v.String()))
				return 1
			},
		},
		Properties: rbxmk.Properties{
			"Type": {
				Get: func(s rbxmk.State, v types.Value) int {
					action := v.(*rtypes.DescAction)
					return s.Push(types.Int(action.Action.Type))
				},
				Set: func(s rbxmk.State, v types.Value) {
					action := v.(*rtypes.DescAction)
					action.Action.Type = diff.Type(s.Pull(3, "int").(types.Int))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("int"),
						Summary:     "Types/DescAction:Properties/Type/Summary",
						Description: "Types/DescAction:Properties/Type/Description",
					}
				},
			},
			"Element": {
				Get: func(s rbxmk.State, v types.Value) int {
					action := v.(*rtypes.DescAction)
					return s.Push(types.String(action.Action.Element.String()))
				},
				Set: func(s rbxmk.State, v types.Value) {
					action := v.(*rtypes.DescAction)
					switch e := string(s.Pull(3, "string").(types.String)); e {
					case "Class":
						action.Action.Element = diff.Class
					case "Property":
						action.Action.Element = diff.Property
					case "Function":
						action.Action.Element = diff.Function
					case "Event":
						action.Action.Element = diff.Event
					case "Callback":
						action.Action.Element = diff.Callback
					case "Enum":
						action.Action.Element = diff.Enum
					case "EnumItem":
						action.Action.Element = diff.EnumItem
					default:
						s.RaiseError("unexpected value %q", e)
					}
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						Summary:     "Types/DescAction:Properties/Element/Summary",
						Description: "Types/DescAction:Properties/Element/Description",
					}
				},
			},
			"Primary": {
				Get: func(s rbxmk.State, v types.Value) int {
					action := v.(*rtypes.DescAction)
					return s.Push(types.String(action.Action.Primary))
				},
				Set: func(s rbxmk.State, v types.Value) {
					action := v.(*rtypes.DescAction)
					action.Action.Primary = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						Summary:     "Types/DescAction:Properties/Primary/Summary",
						Description: "Types/DescAction:Properties/Primary/Description",
					}
				},
			},
			"Secondary": {
				Get: func(s rbxmk.State, v types.Value) int {
					action := v.(*rtypes.DescAction)
					return s.Push(types.String(action.Action.Secondary))
				},
				Set: func(s rbxmk.State, v types.Value) {
					action := v.(*rtypes.DescAction)
					action.Action.Secondary = string(s.Pull(3, "string").(types.String))
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("string"),
						Summary:     "Types/DescAction:Properties/Secondary/Summary",
						Description: "Types/DescAction:Properties/Secondary/Description",
					}
				},
			},
		},
		Methods: rbxmk.Methods{
			"Field": {
				Func: func(s rbxmk.State, v types.Value) int {
					action := v.(*rtypes.DescAction)
					name := string(s.Pull(3, "string").(types.String))
					fvalue, ok := action.Fields[name]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					value := pushDescActionField(fvalue)
					if value == nil {
						return s.Push(rtypes.Nil)
					}
					return s.Push(value)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "name", Type: dt.Prim("string")},
						},
						Returns: dump.Parameters{
							{Type: dt.Prim("any")},
						},
						Summary:     "Types/DescAction:Methods/Field/Summary",
						Description: "Types/DescAction:Methods/Field/Description",
					}
				},
			},
			"Fields": {
				Func: func(s rbxmk.State, v types.Value) int {
					action := v.(*rtypes.DescAction)
					dict := make(rtypes.Dictionary, len(action.Fields))
					for k, v := range action.Fields {
						value := pushDescActionField(v)
						if value == nil {
							//TODO: Emit a warning or something.
							continue
						}
						dict[k] = value
					}
					return s.Push(dict)
				},
				Dump: func() dump.Function {
					return dump.Function{
						Returns: dump.Parameters{
							{Type: dt.Dictionary{V: dt.Prim("any")}},
						},
						Summary:     "Types/DescAction:Methods/Fields/Summary",
						Description: "Types/DescAction:Methods/Fields/Description",
					}
				},
			},
			"SetField": {
				Func: func(s rbxmk.State, v types.Value) int {
					action := v.(*rtypes.DescAction)
					name := string(s.Pull(3, "string").(types.String))
					value := pullDescActionField(name, s.Pull(4, "Variant"))
					if value == nil {
						return s.RaiseError("unexpected type %s", v.Type())
					}
					action.Fields[name] = value
					return 0
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "name", Type: dt.Prim("string")},
							{Name: "value", Type: dt.Prim("any")},
						},
						Summary:     "Types/DescAction:Methods/SetField/Summary",
						Description: "Types/DescAction:Methods/SetField/Description",
					}
				},
			},
			"SetFields": {
				Func: func(s rbxmk.State, v types.Value) int {
					action := v.(*rtypes.DescAction)
					values := s.Pull(3, "Dictionary").(rtypes.Dictionary)
					fields := make(rbxdump.Fields, len(values))
					for k, v := range values {
						value := pullDescActionField(k, v)
						if value == nil {
							return s.RaiseError("field %q: unexpected type %s", k, v.Type())
						}
						fields[k] = value
					}
					action.Fields = fields
					return 0
				},
				Dump: func() dump.Function {
					return dump.Function{
						Parameters: dump.Parameters{
							{Name: "fields", Type: dt.Dictionary{V: dt.Prim("any")}},
						},
						Summary:     "Types/DescAction:Methods/SetFields/Summary",
						Description: "Types/DescAction:Methods/SetFields/Description",
					}
				},
			},
		},
		Dump: func() dump.TypeDef {
			return dump.TypeDef{
				Summary:     "Types/DescAction:Summary",
				Description: "Types/DescAction:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			Array,
			Bool,
			Dictionary,
			Int,
			ParameterDesc,
			String,
			TypeDesc,
		},
	}
}

func pushDescActionField(v interface{}) types.Value {
	switch v := v.(type) {
	case bool:
		return types.Bool(v)
	case int:
		return types.Int(v)
	case string:
		return types.String(v)
	case rbxdump.Tags:
		a := make(rtypes.Array, len(v))
		for i, v := range v {
			a[i] = types.String(v)
		}
		return a
	case rbxdump.Type:
		return rtypes.TypeDesc{Embedded: v}
	case []rbxdump.Parameter:
		a := make(rtypes.Array, len(v))
		for i, v := range v {
			a[i] = rtypes.ParameterDesc{Parameter: v}
		}
		return a
	default:
		return nil
	}
}

func pullDescActionField(k string, v types.Value) interface{} {
	switch v := v.(type) {
	case types.Bool:
		return bool(v)
	case types.Int:
		return int(v)
	case types.String:
		return string(v)
	case rtypes.TypeDesc:
		return v.Embedded
	case rtypes.Array:
		switch k {
		case "Parameters":
			a := make([]rbxdump.Parameter, len(v))
			for i, v := range v {
				v, ok := v.(rtypes.ParameterDesc)
				if !ok {
					return nil
				}
				a[i] = v.Parameter
			}
			return a
		case "Tags":
			a := make([]string, len(v))
			for i, v := range v {
				v, ok := v.(types.Stringlike)
				if !ok {
					return nil
				}
				a[i] = v.Stringlike()
			}
			return a
		}
		return nil
	default:
		return nil
	}
}
