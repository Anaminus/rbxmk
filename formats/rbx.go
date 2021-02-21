package formats

import (
	"io"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxfile"
	"github.com/robloxapi/rbxfile/rbxl"
	"github.com/robloxapi/rbxfile/rbxlx"
	"github.com/robloxapi/types"
)

func init() { register(RBXL) }
func RBXL() rbxmk.Format {
	return rbxmk.Format{
		Name:       "rbxl",
		MediaTypes: []string{"application/x-roblox-studio"},
		CanDecode: func(g rbxmk.Global, f rbxmk.FormatOptions, typeName string) bool {
			return typeName == "Instance"
		},
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeRBX(rbxl.DeserializePlace, r)
		},
		Encode: func(g rbxmk.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			return encodeRBX(rbxl.SerializePlace, w, v)
		},
	}
}

func init() { register(RBXM) }
func RBXM() rbxmk.Format {
	return rbxmk.Format{
		Name:       "rbxm",
		MediaTypes: []string{"application/x-roblox-studio"},
		CanDecode: func(g rbxmk.Global, f rbxmk.FormatOptions, typeName string) bool {
			return typeName == "Instance"
		},
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeRBX(rbxl.DeserializeModel, r)
		},
		Encode: func(g rbxmk.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			return encodeRBX(rbxl.SerializeModel, w, v)
		},
	}
}

func init() { register(RBXLX) }
func RBXLX() rbxmk.Format {
	return rbxmk.Format{
		Name:       "rbxlx",
		MediaTypes: []string{"application/x-roblox-studio", "application/xml", "text/plain"},
		CanDecode: func(g rbxmk.Global, f rbxmk.FormatOptions, typeName string) bool {
			return typeName == "Instance"
		},
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeRBX(rbxlx.Deserialize, r)
		},
		Encode: func(g rbxmk.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			return encodeRBX(rbxlx.Serialize, w, v)
		},
	}
}

func init() { register(RBXMX) }
func RBXMX() rbxmk.Format {
	return rbxmk.Format{
		Name:       "rbxmx",
		MediaTypes: []string{"application/x-roblox-studio", "application/xml", "text/plain"},
		CanDecode: func(g rbxmk.Global, f rbxmk.FormatOptions, typeName string) bool {
			return typeName == "Instance"
		},
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			root, err := decodeRBX(rbxlx.Deserialize, r)
			if err != nil {
				return nil, err
			}
			return root, nil
		},
		Encode: func(g rbxmk.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			return encodeRBX(rbxlx.Serialize, w, v)
		},
	}
}

func decodeRBX(method func(r io.Reader) (root *rbxfile.Root, err error), r io.Reader) (v types.Value, err error) {
	root, err := method(r)
	if err != nil {
		return nil, err
	}
	t, err := decodeDataModel(root)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func decodeDataModel(r *rbxfile.Root) (t *rtypes.Instance, err error) {
	t = rtypes.NewDataModel()
	meta := t.Metadata()
	for k, v := range r.Metadata {
		meta[k] = v
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
			pref.Instance.Set(pref.Property, t)
		}
	}
	return t, nil
}

type decinst map[*rbxfile.Instance]*rtypes.Instance

type decprop struct {
	Instance *rtypes.Instance
	Property string
	Value    *rbxfile.Instance
}

func decodeInstance(r *rbxfile.Instance, refs decinst, prefs *[]decprop) (t *rtypes.Instance, err error) {
	if t, ok := refs[r]; ok {
		return t, nil
	}
	t = rtypes.NewInstance(r.ClassName, nil)
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
		v, err := decodeValue(value)
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

func decodeValue(r rbxfile.Value) (t types.PropValue, err error) {
	switch r := r.(type) {
	case rbxfile.ValueString:
		return types.String(r), nil
	case rbxfile.ValueBinaryString:
		return types.BinaryString(r), nil
	case rbxfile.ValueProtectedString:
		return types.ProtectedString(r), nil
	case rbxfile.ValueContent:
		return types.Content(r), nil
	case rbxfile.ValueBool:
		return types.Bool(r), nil
	case rbxfile.ValueInt:
		return types.Int(r), nil
	case rbxfile.ValueFloat:
		return types.Float(r), nil
	case rbxfile.ValueDouble:
		return types.Double(r), nil
	case rbxfile.ValueUDim:
		return types.UDim(r), nil
	case rbxfile.ValueUDim2:
		return types.UDim2{
			X: types.UDim(r.X),
			Y: types.UDim(r.Y),
		}, nil
	case rbxfile.ValueRay:
		return types.Ray{
			Origin:    types.Vector3(r.Origin),
			Direction: types.Vector3(r.Direction),
		}, nil
	case rbxfile.ValueFaces:
		return types.Faces(r), nil
	case rbxfile.ValueAxes:
		return types.Axes(r), nil
	case rbxfile.ValueBrickColor:
		return types.BrickColor(r), nil
	case rbxfile.ValueColor3:
		return types.Color3(r), nil
	case rbxfile.ValueVector2:
		return types.Vector2(r), nil
	case rbxfile.ValueVector3:
		return types.Vector3(r), nil
	case rbxfile.ValueCFrame:
		return types.CFrame{
			Position: types.Vector3(r.Position),
			Rotation: r.Rotation,
		}, nil
	case rbxfile.ValueToken:
		return types.Token(r), nil
	case rbxfile.ValueVector3int16:
		return types.Vector3int16(r), nil
	case rbxfile.ValueVector2int16:
		return types.Vector2int16(r), nil
	case rbxfile.ValueNumberSequence:
		t := make(types.NumberSequence, len(r))
		for i, k := range r {
			t[i] = types.NumberSequenceKeypoint(k)
		}
		return t, nil
	case rbxfile.ValueColorSequence:
		t := make(types.ColorSequence, len(r))
		for i, k := range r {
			t[i] = types.ColorSequenceKeypoint{
				Time:     k.Time,
				Value:    types.Color3(k.Value),
				Envelope: k.Envelope,
			}
		}
		return t, nil
	case rbxfile.ValueNumberRange:
		return types.NumberRange(r), nil
	case rbxfile.ValueRect2D:
		return types.Rect{
			Min: types.Vector2(r.Min),
			Max: types.Vector2(r.Max),
		}, nil
	case rbxfile.ValuePhysicalProperties:
		return types.PhysicalProperties(r), nil
	case rbxfile.ValueColor3uint8:
		return rtypes.Color3uint8{
			R: float32(r.R) / 255,
			G: float32(r.G) / 255,
			B: float32(r.B) / 255,
		}, nil
	case rbxfile.ValueInt64:
		return types.Int64(r), nil
	case rbxfile.ValueSharedString:
		return types.SharedString(r), nil
	default:
		return nil, cannotEncode(r)
	}
}

////////////////////////////////////////////////////////////////////////////////

func encodeRBX(method func(w io.Writer, root *rbxfile.Root) (err error), w io.Writer, v types.Value) error {
	var t *rtypes.Instance
	switch v := v.(type) {
	case *rtypes.Instance:
		if !v.IsDataModel() {
			t = rtypes.NewDataModel()
			t.AddChild(v)
			break
		}
		t = v
	case rtypes.Objects:
		t = rtypes.NewDataModel()
		for _, inst := range v {
			t.AddChild(inst)
		}
	default:
		return cannotEncode(v)
	}
	r, err := encodeDataModel(t)
	if err != nil {
		return err
	}
	return method(w, r)
}

type encinst map[*rtypes.Instance]*rbxfile.Instance

type encprop struct {
	Instance *rbxfile.Instance
	Property string
	Value    *rtypes.Instance
}

func encodeDataModel(t *rtypes.Instance) (r *rbxfile.Root, err error) {
	r = rbxfile.NewRoot()
	meta := t.Metadata()
	r.Metadata = make(map[string]string, len(meta))
	for k, v := range meta {
		r.Metadata[k] = v
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

func encodeInstance(t *rtypes.Instance, refs encinst, prefs *[]encprop) (r *rbxfile.Instance, err error) {
	if r, ok := refs[t]; ok {
		return r, nil
	}
	r = rbxfile.NewInstance(t.ClassName)
	r.IsService = t.IsService
	r.Reference = t.Reference
	refs[t] = r
	for prop, value := range t.Properties() {
		if v, ok := value.(*rtypes.Instance); ok {
			*prefs = append(*prefs, encprop{
				Instance: r,
				Property: prop,
				Value:    v,
			})
			continue
		}
		v, err := encodeValue(value)
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

func encodeValue(t types.PropValue) (r rbxfile.Value, err error) {
	switch t := t.(type) {
	case types.String:
		return rbxfile.ValueString(t), nil
	case types.BinaryString:
		return rbxfile.ValueBinaryString(t), nil
	case types.ProtectedString:
		return rbxfile.ValueProtectedString(t), nil
	case types.Content:
		return rbxfile.ValueContent(t), nil
	case types.Bool:
		return rbxfile.ValueBool(t), nil
	case types.Int:
		return rbxfile.ValueInt(t), nil
	case types.Float:
		return rbxfile.ValueFloat(t), nil
	case types.Double:
		return rbxfile.ValueDouble(t), nil
	case types.UDim:
		return rbxfile.ValueUDim(t), nil
	case types.UDim2:
		return rbxfile.ValueUDim2{
			X: rbxfile.ValueUDim(t.X),
			Y: rbxfile.ValueUDim(t.Y),
		}, nil
	case types.Ray:
		return rbxfile.ValueRay{
			Origin:    rbxfile.ValueVector3(t.Origin),
			Direction: rbxfile.ValueVector3(t.Direction),
		}, nil
	case types.Faces:
		return rbxfile.ValueFaces(t), nil
	case types.Axes:
		return rbxfile.ValueAxes(t), nil
	case types.BrickColor:
		return rbxfile.ValueBrickColor(t), nil
	case types.Color3:
		return rbxfile.ValueColor3(t), nil
	case types.Vector2:
		return rbxfile.ValueVector2(t), nil
	case types.Vector3:
		return rbxfile.ValueVector3(t), nil
	case types.CFrame:
		return rbxfile.ValueCFrame{
			Position: rbxfile.ValueVector3(t.Position),
			Rotation: t.Rotation,
		}, nil
	case types.Token:
		return rbxfile.ValueToken(t), nil
	case types.Vector3int16:
		return rbxfile.ValueVector3int16(t), nil
	case types.Vector2int16:
		return rbxfile.ValueVector2int16(t), nil
	case types.NumberSequence:
		r := make(rbxfile.ValueNumberSequence, len(t))
		for i, k := range t {
			r[i] = rbxfile.ValueNumberSequenceKeypoint(k)
		}
		return r, nil
	case types.ColorSequence:
		r := make(rbxfile.ValueColorSequence, len(t))
		for i, k := range t {
			r[i] = rbxfile.ValueColorSequenceKeypoint{
				Time:     k.Time,
				Value:    rbxfile.ValueColor3(k.Value),
				Envelope: k.Envelope,
			}
		}
		return r, nil
	case types.NumberRange:
		return rbxfile.ValueNumberRange(t), nil
	case types.Rect:
		return rbxfile.ValueRect2D{
			Min: rbxfile.ValueVector2(t.Min),
			Max: rbxfile.ValueVector2(t.Max),
		}, nil
	case types.PhysicalProperties:
		return rbxfile.ValuePhysicalProperties(t), nil
	case rtypes.Color3uint8:
		return rbxfile.ValueColor3uint8{
			R: uint8(t.R * 255),
			G: uint8(t.G * 255),
			B: uint8(t.B * 255),
		}, nil
	case types.Int64:
		return rbxfile.ValueInt64(t), nil
	case types.SharedString:
		return rbxfile.ValueSharedString(t), nil
	default:
		return nil, cannotEncode(t)
	}
}
