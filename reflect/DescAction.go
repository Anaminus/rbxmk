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
		SetTo: func(p interface{}, v types.Value) error {
			switch p := p.(type) {
			case **rtypes.DescAction:
				*p = v.(*rtypes.DescAction)
			default:
				return setPtrErr(p, v)
			}
			return nil
		},
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
					enum := s.MustEnum("DescActionType")
					value := int(action.Action.Type)
					item := enum.Value(value)
					if item == nil {
						s.RaiseError("invalid value %d for %s", value, enum.Name())
					}
					return s.Push(item)
				},
				Set: func(s rbxmk.State, v types.Value) {
					action := v.(*rtypes.DescAction)
					enum := s.MustEnum("DescActionType")
					value := s.L.Get(3)
					item := enum.Pull(value)
					if item == nil {
						s.RaiseError("invalid value %s for %s", value, enum.Name())
						return
					}
					action.Action.Type = diff.Type(item.Value())
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("Enum.DescActionType"),
						Summary:     "Types/DescAction:Properties/Type/Summary",
						Description: "Types/DescAction:Properties/Type/Description",
					}
				},
			},
			"Element": {
				Get: func(s rbxmk.State, v types.Value) int {
					action := v.(*rtypes.DescAction)
					enum := s.MustEnum("DescActionElement")
					value := int(action.Element)
					item := enum.Value(value)
					if item == nil {
						s.RaiseError("invalid value %d for %s", value, enum.Name())
					}
					return s.Push(item)
				},
				Set: func(s rbxmk.State, v types.Value) {
					action := v.(*rtypes.DescAction)
					enum := s.MustEnum("DescActionElement")
					value := s.L.Get(3)
					item := enum.Pull(value)
					if item == nil {
						s.RaiseError("invalid value %s for %s", value, enum.Name())
						return
					}
					action.Element = diff.Element(item.Value())
				},
				Dump: func() dump.Property {
					return dump.Property{
						ValueType:   dt.Prim("Enum.DescActionElement"),
						Summary:     "Types/DescAction:Properties/Element/Summary",
						Description: "Types/DescAction:Properties/Element/Description",
					}
				},
			},
			"Primary": {
				Get: func(s rbxmk.State, v types.Value) int {
					action := v.(*rtypes.DescAction)
					return s.Push(types.String(action.Primary))
				},
				Set: func(s rbxmk.State, v types.Value) {
					action := v.(*rtypes.DescAction)
					action.Primary = string(s.Pull(3, "string").(types.String))
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
					return s.Push(types.String(action.Secondary))
				},
				Set: func(s rbxmk.State, v types.Value) {
					action := v.(*rtypes.DescAction)
					action.Secondary = string(s.Pull(3, "string").(types.String))
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
					name := string(s.Pull(2, "string").(types.String))
					fvalue, ok := action.Fields[name]
					if !ok {
						return s.Push(rtypes.Nil)
					}
					value := pushDescField(s.Context(), fvalue)
					if value == nil {
						return s.Push(rtypes.Nil)
					}
					s.L.Push(value)
					return 1
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
					return s.Push(rtypes.DescFields(action.Fields))
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
					name := string(s.Pull(2, "string").(types.String))
					fvalue := s.L.Get(3)
					if fvalue == lua.LNil {
						delete(action.Fields, name)
						return 0
					}
					value, err := pullDescField(s.Context(), name, fvalue)
					if err != nil {
						return s.RaiseError(err.Error())
					}
					if action.Fields == nil {
						action.Fields = rbxdump.Fields{}
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
					fields := s.Pull(2, "DescFields").(rtypes.DescFields)
					action.Fields = rbxdump.Fields(fields)
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
		Constructors: rbxmk.Constructors{
			"new": rbxmk.Constructor{
				Func: func(s rbxmk.State) int {
					var v rtypes.DescAction

					typeEnum := s.MustEnum("DescActionType")
					actionType := typeEnum.Pull(s.L.Get(1))
					if actionType == nil {
						return s.ArgError(1, "invalid value %s for %s", actionType, typeEnum.Name())
					}
					v.Action.Type = diff.Type(actionType.Value())

					elementEnum := s.MustEnum("DescActionElement")
					actionElement := elementEnum.Pull(s.L.Get(2))
					if actionElement == nil {
						return s.ArgError(2, "invalid value %s for %s", actionElement, elementEnum.Name())
					}
					v.Element = diff.Element(actionElement.Value())

					return s.Push(&v)
				},
				Dump: func() dump.MultiFunction {
					return dump.MultiFunction{
						dump.Function{
							Parameters: dump.Parameters{
								{Name: "type", Type: dt.Prim("Enum.DescActionType")},
								{Name: "element", Type: dt.Prim("Enum.DescActionElement")},
							},
							Returns: dump.Parameters{
								{Type: dt.Prim("DescAction")},
							},
							Summary:     "Types/DescAction:Constructors/new/Summary",
							Description: "Types/DescAction:Constructors/new/Description",
						},
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
		Enums: []*rtypes.Enum{
			rtypes.NewEnum("DescActionType",
				rtypes.NewItem{Name: diff.Remove.String(), Value: int(diff.Remove)},
				rtypes.NewItem{Name: diff.Change.String(), Value: int(diff.Change)},
				rtypes.NewItem{Name: diff.Add.String(), Value: int(diff.Add)},
			),
			rtypes.NewEnum("DescActionElement",
				rtypes.NewItem{Name: diff.Class.String(), Value: int(diff.Class)},
				rtypes.NewItem{Name: diff.Property.String(), Value: int(diff.Property)},
				rtypes.NewItem{Name: diff.Function.String(), Value: int(diff.Function)},
				rtypes.NewItem{Name: diff.Event.String(), Value: int(diff.Event)},
				rtypes.NewItem{Name: diff.Callback.String(), Value: int(diff.Callback)},
				rtypes.NewItem{Name: diff.Enum.String(), Value: int(diff.Enum)},
				rtypes.NewItem{Name: diff.EnumItem.String(), Value: int(diff.EnumItem)},
			),
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
