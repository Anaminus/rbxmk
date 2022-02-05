package formats

import (
	"fmt"
	"io"
	"sort"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxattr"
	"github.com/robloxapi/types"
)

func decodeAttributeValue(a rbxattr.Value) (t types.Value, err error) {
	switch a := a.(type) {
	case *rbxattr.ValueString:
		return types.String(*a), nil
	case *rbxattr.ValueBool:
		return types.Bool(*a), nil
	case *rbxattr.ValueFloat:
		return types.Float(*a), nil
	case *rbxattr.ValueDouble:
		return types.Double(*a), nil
	case *rbxattr.ValueUDim:
		return types.UDim(*a), nil
	case *rbxattr.ValueUDim2:
		return types.UDim2{
			X: types.UDim(a.X),
			Y: types.UDim(a.Y),
		}, nil
	case *rbxattr.ValueBrickColor:
		return types.BrickColor(*a), nil
	case *rbxattr.ValueColor3:
		return types.Color3(*a), nil
	case *rbxattr.ValueVector2:
		return types.Vector2(*a), nil
	case *rbxattr.ValueVector3:
		return types.Vector3(*a), nil
	case *rbxattr.ValueNumberSequence:
		t := make(types.NumberSequence, len(*a))
		for i, k := range *a {
			t[i] = types.NumberSequenceKeypoint{
				Time:     k.Time,
				Value:    k.Value,
				Envelope: k.Envelope,
			}
		}
		return t, nil
	case *rbxattr.ValueColorSequence:
		t := make(types.ColorSequence, len(*a))
		for i, k := range *a {
			t[i] = types.ColorSequenceKeypoint{
				Time:     k.Time,
				Value:    types.Color3(k.Value),
				Envelope: k.Envelope,
			}
		}
		return t, nil
	case *rbxattr.ValueNumberRange:
		return types.NumberRange(*a), nil
	case *rbxattr.ValueRect:
		return types.Rect{
			Min: types.Vector2(a.Min),
			Max: types.Vector2(a.Max),
		}, nil
	default:
		return nil, cannotEncode(a)
	}
}

func encodeAttributeValue(t types.Value) (a rbxattr.Value, err error) {
	switch t := t.(type) {
	case types.String:
		a := rbxattr.ValueString(t)
		return &a, nil
	case types.Bool:
		a := rbxattr.ValueBool(t)
		return &a, nil
	case types.Float:
		a := rbxattr.ValueFloat(t)
		return &a, nil
	case types.Double:
		a := rbxattr.ValueDouble(t)
		return &a, nil
	case types.UDim:
		a := rbxattr.ValueUDim(t)
		return &a, nil
	case types.UDim2:
		a := rbxattr.ValueUDim2{
			X: rbxattr.ValueUDim(t.X),
			Y: rbxattr.ValueUDim(t.Y),
		}
		return &a, nil
	case types.BrickColor:
		a := rbxattr.ValueBrickColor(t)
		return &a, nil
	case types.Color3:
		a := rbxattr.ValueColor3(t)
		return &a, nil
	case types.Vector2:
		a := rbxattr.ValueVector2(t)
		return &a, nil
	case types.Vector3:
		a := rbxattr.ValueVector3(t)
		return &a, nil
	case types.NumberSequence:
		a := make(rbxattr.ValueNumberSequence, len(t))
		for i, k := range t {
			a[i] = rbxattr.ValueNumberSequenceKeypoint{
				Envelope: k.Envelope,
				Time:     k.Time,
				Value:    k.Value,
			}
		}
		return &a, nil
	case types.ColorSequence:
		a := make(rbxattr.ValueColorSequence, len(t))
		for i, k := range t {
			a[i] = rbxattr.ValueColorSequenceKeypoint{
				Envelope: k.Envelope,
				Time:     k.Time,
				Value:    rbxattr.ValueColor3(k.Value),
			}
		}
		return &a, nil
	case types.NumberRange:
		a := rbxattr.ValueNumberRange(t)
		return &a, nil
	case types.Rect:
		a := rbxattr.ValueRect{
			Min: rbxattr.ValueVector2(t.Min),
			Max: rbxattr.ValueVector2(t.Max),
		}
		return &a, nil
	case types.Stringlike:
		a := rbxattr.ValueString(t.Stringlike())
		return &a, nil
	case types.Numberlike:
		a := rbxattr.ValueDouble(t.Numberlike())
		return &a, nil
	case rtypes.NilType:
		return nil, nil
	default:
		return nil, cannotEncode(t)
	}
}

func init() { register(RBXAttr) }
func RBXAttr() rbxmk.Format {
	return rbxmk.Format{
		Name:       "rbxattr",
		MediaTypes: []string{"application/octet-stream"},
		CanDecode: func(g rbxmk.Global, f rbxmk.FormatOptions, typeName string) bool {
			return typeName == "Instance"
		},
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			var model rbxattr.Model
			if _, err = model.ReadFrom(r); err != nil {
				return nil, fmt.Errorf("decode attributes: %w", err)
			}
			dict := make(rtypes.Dictionary, len(model.Value))
			for _, entry := range model.Value {
				if _, ok := dict[entry.Key]; ok {
					continue
				}
				dict[entry.Key], err = decodeAttributeValue(entry.Value)
				if err != nil {
					return nil, fmt.Errorf("decode %q: %w", entry.Key, err)
				}
			}
			return dict, nil
		},
		Encode: func(g rbxmk.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			dict, ok := v.(rtypes.Dictionary)
			if !ok {
				return fmt.Errorf("Dictionary expected, got %s", v.Type())
			}

			// Roblox's implementation encodes using reverse insertion order. To
			// match this would require some sort of internal ordered dictionary
			// type. Instead, we'll just sort ascending.
			keys := make([]string, 0, len(dict))
			for key := range dict {
				keys = append(keys, key)
			}
			sort.Strings(keys)

			var model rbxattr.Model
			model.Value = make(rbxattr.ValueDictionary, 0, len(dict))
			for _, key := range keys {
				value, err := encodeAttributeValue(dict[key])
				if err != nil {
					return err
				}
				if value == nil {
					continue
				}
				model.Value = append(model.Value, rbxattr.Entry{
					Key:   key,
					Value: value,
				})
			}

			if _, err := model.WriteTo(w); err != nil {
				return fmt.Errorf("encode attributes: %w", err)
			}
			return nil
		},
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/rbxattr:Summary",
				Description: "Formats/rbxattr:Description",
			}
		},
	}
}
