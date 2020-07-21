package formats

import (
	"bytes"
	"io"

	"github.com/anaminus/rbxmk"
	rtypes "github.com/anaminus/rbxmk/types"
	"github.com/robloxapi/rbxfile"
	"github.com/robloxapi/rbxfile/rbxl"
	"github.com/robloxapi/rbxfile/rbxlx"
	"github.com/robloxapi/types"
)

type decinst map[*rbxfile.Instance]*rtypes.Instance

type decprop struct {
	Instance *rtypes.Instance
	Property string
	Value    *rbxfile.Instance
}

func decodeValue(r rbxfile.Value, refs decinst, prefs *[]decprop) (t rbxmk.TValue, err error) {
	switch r := r.(type) {
	case rbxfile.ValueString:
		return rbxmk.TValue{Type: "string", Value: string(r)}, nil
	case rbxfile.ValueBinaryString:
		return rbxmk.TValue{Type: "BinaryString", Value: []byte(r)}, nil
	case rbxfile.ValueProtectedString:
		return rbxmk.TValue{Type: "ProtectedString", Value: string(r)}, nil
	case rbxfile.ValueContent:
		return rbxmk.TValue{Type: "Content", Value: string(r)}, nil
	case rbxfile.ValueBool:
		return rbxmk.TValue{Type: "bool", Value: bool(r)}, nil
	case rbxfile.ValueInt:
		return rbxmk.TValue{Type: "int", Value: int(r)}, nil
	case rbxfile.ValueFloat:
		return rbxmk.TValue{Type: "float", Value: float32(r)}, nil
	case rbxfile.ValueDouble:
		return rbxmk.TValue{Type: "double", Value: float64(r)}, nil
	case rbxfile.ValueUDim:
		return rbxmk.TValue{Type: "UDim", Value: types.UDim(r)}, nil
	case rbxfile.ValueUDim2:
		return rbxmk.TValue{Type: "UDim2", Value: types.UDim2{
			X: types.UDim(r.X),
			Y: types.UDim(r.Y),
		}}, nil
	case rbxfile.ValueRay:
		return rbxmk.TValue{Type: "Ray", Value: types.Ray{
			Origin:    types.Vector3(r.Origin),
			Direction: types.Vector3(r.Direction),
		}}, nil
	case rbxfile.ValueFaces:
		return rbxmk.TValue{Type: "Faces", Value: types.Faces(r)}, nil
	case rbxfile.ValueAxes:
		return rbxmk.TValue{Type: "Axes", Value: types.Axes(r)}, nil
	case rbxfile.ValueBrickColor:
		return rbxmk.TValue{Type: "BrickColor", Value: types.BrickColor(r)}, nil
	case rbxfile.ValueColor3:
		return rbxmk.TValue{Type: "Color3", Value: types.Color3(r)}, nil
	case rbxfile.ValueVector2:
		return rbxmk.TValue{Type: "Vector2", Value: types.Vector2(r)}, nil
	case rbxfile.ValueVector3:
		return rbxmk.TValue{Type: "Vector3", Value: types.Vector3(r)}, nil
	case rbxfile.ValueCFrame:
		return rbxmk.TValue{Type: "CFrame", Value: types.CFrame{
			Position: types.Vector3(r.Position),
			Rotation: r.Rotation,
		}}, nil
	case rbxfile.ValueToken:
		return rbxmk.TValue{Type: "token", Value: uint32(r)}, nil
	case rbxfile.ValueVector3int16:
		return rbxmk.TValue{Type: "Vector3int16", Value: types.Vector3int16(r)}, nil
	case rbxfile.ValueVector2int16:
		return rbxmk.TValue{Type: "Vector2int16", Value: types.Vector2int16(r)}, nil
	case rbxfile.ValueNumberSequence:
		t := make(types.NumberSequence, len(r))
		for i, k := range r {
			t[i] = types.NumberSequenceKeypoint(k)
		}
		return rbxmk.TValue{Type: "NumberSequence", Value: t}, nil
	case rbxfile.ValueColorSequence:
		t := make(types.ColorSequence, len(r))
		for i, k := range r {
			t[i] = types.ColorSequenceKeypoint{
				Time:     k.Time,
				Value:    types.Color3(k.Value),
				Envelope: k.Envelope,
			}
		}
		return rbxmk.TValue{Type: "ColorSequence", Value: t}, nil
	case rbxfile.ValueNumberRange:
		return rbxmk.TValue{Type: "NumberRange", Value: types.NumberRange(r)}, nil
	case rbxfile.ValueRect2D:
		return rbxmk.TValue{Type: "Rect", Value: types.Rect{
			Min: types.Vector2(r.Min),
			Max: types.Vector2(r.Max),
		}}, nil
	case rbxfile.ValuePhysicalProperties:
		return rbxmk.TValue{Type: "PhysicalProperties", Value: types.PhysicalProperties(r)}, nil
	case rbxfile.ValueColor3uint8:
		return rbxmk.TValue{Type: "Color3", Value: types.Color3{
			R: float32(r.R) / 255,
			G: float32(r.G) / 255,
			B: float32(r.B) / 255,
		}}, nil
	case rbxfile.ValueInt64:
		return rbxmk.TValue{Type: "int64", Value: int64(r)}, nil
	case rbxfile.ValueSharedString:
		return rbxmk.TValue{Type: "SharedString", Value: []byte(r)}, nil
	default:
		return rbxmk.TValue{}, cannotEncode(r, false)
	}
}

func decodeInstance(r *rbxfile.Instance, refs decinst, prefs *[]decprop) (t *rtypes.Instance, err error) {
	if t, ok := refs[r]; ok {
		return t, nil
	}
	t = rtypes.NewInstance(r.ClassName)
	t.IsService = r.IsService
	t.Reference = r.Reference
	refs[r] = t
	for prop, value := range r.Properties {
		if v, ok := value.(rbxfile.ValueReference); ok {
			*prefs = append(*prefs, decprop{
				Instance: t,
				Property: prop,
				Value:    v.Instance,
			})
			continue
		}
		v, err := decodeValue(value, refs, prefs)
		if err != nil {
			return nil, err
		}
		t.Set(prop, v)
	}
	for _, rc := range r.Children {
		tc, err := decodeInstance(rc, refs, prefs)
		if err != nil {
			return nil, err
		}
		t.AddChild(tc)
	}
	return t, nil
}

func decodeDataModel(r *rbxfile.Root) (t *rtypes.Instance, err error) {
	t = rtypes.NewDataModel()
	for k, v := range r.Metadata {
		t.Set(k, rbxmk.TValue{Type: "string", Value: v})
	}
	refs := decinst{}
	prefs := []decprop{}
	for _, rc := range r.Instances {
		tc, err := decodeInstance(rc, refs, &prefs)
		if err != nil {
			return nil, err
		}
		t.AddChild(tc)
	}
	for _, pref := range prefs {
		if t, ok := refs[pref.Value]; ok {
			pref.Instance.Set(pref.Property, rbxmk.TValue{Type: "Instance", Value: t})
		}
	}
	return t, nil
}

type encinst map[*rtypes.Instance]*rbxfile.Instance

type encprop struct {
	Instance *rbxfile.Instance
	Property string
	Value    *rtypes.Instance
}

func encodeValue(t rbxmk.TValue, refs encinst, prefs *[]encprop) (r rbxfile.Value, err error) {
	switch t.Type {
	case "string":
		return rbxfile.ValueString(t.Value.(string)), nil
	case "BinaryString":
		return rbxfile.ValueBinaryString(t.Value.([]byte)), nil
	case "ProtectedString":
		return rbxfile.ValueProtectedString(t.Value.(string)), nil
	case "Content":
		return rbxfile.ValueContent(t.Value.(string)), nil
	case "bool":
		return rbxfile.ValueBool(t.Value.(bool)), nil
	case "int":
		return rbxfile.ValueInt(t.Value.(int)), nil
	case "float":
		return rbxfile.ValueFloat(t.Value.(float32)), nil
	case "double":
		return rbxfile.ValueDouble(t.Value.(float64)), nil
	case "UDim":
		return rbxfile.ValueUDim(t.Value.(types.UDim)), nil
	case "UDim2":
		return rbxfile.ValueUDim2{
			X: rbxfile.ValueUDim(t.Value.(types.UDim2).X),
			Y: rbxfile.ValueUDim(t.Value.(types.UDim2).Y),
		}, nil
	case "Ray":
		return rbxfile.ValueRay{
			Origin:    rbxfile.ValueVector3(t.Value.(types.Ray).Origin),
			Direction: rbxfile.ValueVector3(t.Value.(types.Ray).Direction),
		}, nil
	case "Faces":
		return rbxfile.ValueFaces(t.Value.(types.Faces)), nil
	case "Axes":
		return rbxfile.ValueAxes(t.Value.(types.Axes)), nil
	case "BrickColor":
		return rbxfile.ValueBrickColor(t.Value.(types.BrickColor)), nil
	case "Color3":
		return rbxfile.ValueColor3(t.Value.(types.Color3)), nil
	case "Vector2":
		return rbxfile.ValueVector2(t.Value.(types.Vector2)), nil
	case "Vector3":
		return rbxfile.ValueVector3(t.Value.(types.Vector3)), nil
	case "CFrame":
		return rbxfile.ValueCFrame{
			Position: rbxfile.ValueVector3(t.Value.(types.CFrame).Position),
			Rotation: t.Value.(types.CFrame).Rotation,
		}, nil
	case "token":
		return rbxfile.ValueToken(t.Value.(uint32)), nil
	case "Vector3int16":
		return rbxfile.ValueVector3int16(t.Value.(types.Vector3int16)), nil
	case "Vector2int16":
		return rbxfile.ValueVector2int16(t.Value.(types.Vector2int16)), nil
	case "NumberSequence":
		t := t.Value.(rbxfile.ValueNumberSequence)
		r := make(rbxfile.ValueNumberSequence, len(t))
		for i, k := range t {
			r[i] = rbxfile.ValueNumberSequenceKeypoint(k)
		}
		return r, nil
	case "ColorSequence":
		t := t.Value.(rbxfile.ValueColorSequence)
		r := make(rbxfile.ValueColorSequence, len(t))
		for i, k := range t {
			r[i] = rbxfile.ValueColorSequenceKeypoint{
				Time:     k.Time,
				Value:    rbxfile.ValueColor3(k.Value),
				Envelope: k.Envelope,
			}
		}
		return r, nil
	case "NumberRange":
		return rbxfile.ValueNumberRange(t.Value.(types.NumberRange)), nil
	case "Rect":
		return rbxfile.ValueRect2D{
			Min: rbxfile.ValueVector2(t.Value.(types.Rect).Min),
			Max: rbxfile.ValueVector2(t.Value.(types.Rect).Max),
		}, nil
	case "PhysicalProperties":
		return rbxfile.ValuePhysicalProperties(t.Value.(types.PhysicalProperties)), nil
	case "Color3uint8":
		return rbxfile.ValueColor3uint8{
			R: byte(t.Value.(types.Color3).R * 255),
			G: byte(t.Value.(types.Color3).G * 255),
			B: byte(t.Value.(types.Color3).B * 255),
		}, nil
	case "int64":
		return rbxfile.ValueInt64(t.Value.(int64)), nil
	case "SharedString":
		return rbxfile.ValueSharedString(t.Value.([]byte)), nil
	default:
		return nil, cannotEncode(t.Type, true)
	}
}

func encodeInstance(t *rtypes.Instance, refs encinst, prefs *[]encprop) (r *rbxfile.Instance, err error) {
	if r, ok := refs[t]; ok {
		return r, nil
	}
	r = rbxfile.NewInstance(t.ClassName)
	r.IsService = t.IsService
	r.Reference = t.Reference
	refs[t] = r
	for prop, value := range t.Properties() {
		if v, ok := value.Value.(*rtypes.Instance); ok {
			*prefs = append(*prefs, encprop{
				Instance: r,
				Property: prop,
				Value:    v,
			})
			continue
		}
		v, err := encodeValue(value, refs, prefs)
		if err != nil {
			return nil, err
		}
		r.Properties[prop] = v
	}
	for _, tc := range t.Children() {
		rc, err := encodeInstance(tc, refs, prefs)
		if err != nil {
			return nil, err
		}
		r.Children = append(r.Children, rc)
	}
	return r, nil
}

func encodeDataModel(t *rtypes.Instance) (r *rbxfile.Root, err error) {
	r = rbxfile.NewRoot()
	for prop, value := range t.Properties() {
		if s, ok := (rtypes.Stringlike{Value: value.Value}.Stringlike()); ok {
			r.Metadata[prop] = string(s)
		}
	}
	refs := encinst{}
	prefs := []encprop{}
	for _, tc := range t.Children() {
		rc, err := encodeInstance(tc, refs, &prefs)
		if err != nil {
			return nil, err
		}
		r.Instances = append(r.Instances, rc)
	}
	for _, pref := range prefs {
		if r, ok := refs[pref.Value]; ok {
			pref.Instance.Properties[pref.Property] = rbxfile.ValueReference{Instance: r}
		}
	}
	return
}

func decodeRBX(method func(r io.Reader) (root *rbxfile.Root, err error), b []byte) (v rbxmk.Value, err error) {
	root, err := method(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return decodeDataModel(root)
}

func encodeRBX(method func(w io.Writer, root *rbxfile.Root) (err error), v rbxmk.Value) (b []byte, err error) {
	var t *rtypes.Instance
	switch v := v.(type) {
	case *rtypes.Instance:
		if !v.IsDataModel() {
			t = rtypes.NewDataModel()
			t.AddChild(v)
			break
		}
		t = v
	case []*rtypes.Instance:
		t = rtypes.NewDataModel()
		for _, inst := range v {
			t.AddChild(inst)
		}
	default:
		return nil, cannotEncode(v, false)
	}
	r, err := encodeDataModel(t)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := method(&buf, r); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func RBXL() rbxmk.Format {
	return rbxmk.Format{
		Name: "rbxl",
		Decode: func(b []byte) (v rbxmk.Value, err error) {
			return decodeRBX(rbxl.DeserializePlace, b)
		},
		Encode: func(v rbxmk.Value) (b []byte, err error) {
			return encodeRBX(rbxl.SerializePlace, v)
		},
	}
}

func RBXM() rbxmk.Format {
	return rbxmk.Format{
		Name: "rbxm",
		Decode: func(b []byte) (v rbxmk.Value, err error) {
			return decodeRBX(rbxl.DeserializeModel, b)
		},
		Encode: func(v rbxmk.Value) (b []byte, err error) {
			return encodeRBX(rbxl.SerializeModel, v)
		},
	}
}

func RBXLX() rbxmk.Format {
	return rbxmk.Format{
		Name: "rbxlx",
		Decode: func(b []byte) (v rbxmk.Value, err error) {
			return decodeRBX(rbxlx.Deserialize, b)
		},
		Encode: func(v rbxmk.Value) (b []byte, err error) {
			return encodeRBX(rbxlx.Serialize, v)
		},
	}
}

func RBXMX() rbxmk.Format {
	return rbxmk.Format{
		Name: "rbxmx",
		Decode: func(b []byte) (v rbxmk.Value, err error) {
			return decodeRBX(rbxlx.Deserialize, b)
		},
		Encode: func(v rbxmk.Value) (b []byte, err error) {
			return encodeRBX(rbxlx.Serialize, v)
		},
	}
}
