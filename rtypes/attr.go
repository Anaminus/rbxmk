package rtypes

import (
	"fmt"
	"io"
	"sort"

	"github.com/robloxapi/rbxattr"
	"github.com/robloxapi/types"
)

const T_AttrConfig = "AttrConfig"

// AttrConfig configures an Instance's attributes API.
type AttrConfig struct {
	// Property is the name of the property to which attributes will be
	// serialized. An empty string defaults to "AttributesSerialize".
	Property string
}

// Type returns a string identifying the type of the value.
func (*AttrConfig) Type() string {
	return T_AttrConfig
}

// String returns a string representation of the value.
func (*AttrConfig) String() string {
	return "Attr"
}

// Of returns the AttrConfig of an instance. If inst is nil, a is returned.
func (a *AttrConfig) Of(inst *Instance) *AttrConfig {
	if inst != nil {
		if attrcfg := inst.AttrConfig(); attrcfg != nil {
			return attrcfg
		}
	}
	return a
}

// cannotEncode returns an error indicating that v cannot be encoded.
func cannotEncode(v interface{}) error {
	if v, ok := v.(types.Value); ok {
		return fmt.Errorf("cannot encode %s", v.Type())
	}
	return fmt.Errorf("cannot encode %T", v)
}

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
	case *rbxattr.ValueCFrame:
		return types.CFrame{
			Position: types.Vector3(a.Position),
			Rotation: a.Rotation,
		}, nil
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
	case types.CFrame:
		a := rbxattr.ValueCFrame{
			Position: rbxattr.ValueVector3(t.Position),
			Rotation: t.Rotation,
		}
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
	case NilType:
		return nil, nil
	default:
		return nil, cannotEncode(t)
	}
}

func DecodeAttributes(r io.Reader) (v types.Value, err error) {
	var model rbxattr.Model
	if _, err = model.ReadFrom(r); err != nil {
		return nil, fmt.Errorf("decode attributes: %w", err)
	}
	dict := make(Dictionary, len(model.Value))
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
}
func EncodeAttributes(w io.Writer, v types.Value) error {
	dict, ok := v.(Dictionary)
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
}
