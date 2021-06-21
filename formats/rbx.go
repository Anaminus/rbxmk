package formats

import (
	"fmt"
	"io"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/rbxfile"
	"github.com/robloxapi/rbxfile/rbxl"
	"github.com/robloxapi/rbxfile/rbxlx"
	"github.com/robloxapi/types"
)

func init() { register(RBXL) }
func RBXL() rbxmk.Format {
	return rbxmk.Format{
		Name:        "rbxl",
		EncodeTypes: []string{"Instance", "Objects"},
		MediaTypes:  []string{"application/x-roblox-studio"},
		CanDecode: func(g rbxmk.Global, f rbxmk.FormatOptions, typeName string) bool {
			return typeName == "Instance"
		},
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			desc := descOf(f, "Desc", g, nil)
			mode, err := descModeOf(f, "DescMode")
			if err != nil {
				return nil, err
			}
			d := rbxDecoder{
				method: rbxl.DeserializePlace,
				r:      r,
				desc:   desc,
				mode:   mode,
			}
			return d.rbx()
		},
		Encode: func(g rbxmk.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			desc := descOf(f, "Desc", g, v)
			mode, err := descModeOf(f, "DescMode")
			if err != nil {
				return err
			}
			e := rbxEncoder{
				method: rbxl.SerializePlace,
				w:      w,
				desc:   desc,
				mode:   mode,
			}
			return e.rbx(v)
		},
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/rbxl:Summary",
				Description: "Formats/rbxl:Description",
			}
		},
	}
}

func init() { register(RBXM) }
func RBXM() rbxmk.Format {
	return rbxmk.Format{
		Name:        "rbxm",
		EncodeTypes: []string{"Instance", "Objects"},
		MediaTypes:  []string{"application/x-roblox-studio"},
		CanDecode: func(g rbxmk.Global, f rbxmk.FormatOptions, typeName string) bool {
			return typeName == "Instance"
		},
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			desc := descOf(f, "Desc", g, nil)
			mode, err := descModeOf(f, "DescMode")
			if err != nil {
				return nil, err
			}
			d := rbxDecoder{
				method: rbxl.DeserializeModel,
				r:      r,
				desc:   desc,
				mode:   mode,
			}
			return d.rbx()
		},
		Encode: func(g rbxmk.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			desc := descOf(f, "Desc", g, v)
			mode, err := descModeOf(f, "DescMode")
			if err != nil {
				return err
			}
			e := rbxEncoder{
				method: rbxl.SerializeModel,
				w:      w,
				desc:   desc,
				mode:   mode,
			}
			return e.rbx(v)
		},
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/rbxm:Summary",
				Description: "Formats/rbxm:Description",
			}
		},
	}
}

func init() { register(RBXLX) }
func RBXLX() rbxmk.Format {
	return rbxmk.Format{
		Name:        "rbxlx",
		EncodeTypes: []string{"Instance", "Objects"},
		MediaTypes:  []string{"application/x-roblox-studio", "application/xml", "text/plain"},
		CanDecode: func(g rbxmk.Global, f rbxmk.FormatOptions, typeName string) bool {
			return typeName == "Instance"
		},
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			desc := descOf(f, "Desc", g, nil)
			mode, err := descModeOf(f, "DescMode")
			if err != nil {
				return nil, err
			}
			d := rbxDecoder{
				method: rbxlx.Deserialize,
				r:      r,
				desc:   desc,
				mode:   mode,
			}
			return d.rbx()
		},
		Encode: func(g rbxmk.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			desc := descOf(f, "Desc", g, v)
			mode, err := descModeOf(f, "DescMode")
			if err != nil {
				return err
			}
			e := rbxEncoder{
				method: rbxlx.Serialize,
				w:      w,
				desc:   desc,
				mode:   mode,
			}
			return e.rbx(v)
		},
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/rbxlx:Summary",
				Description: "Formats/rbxlx:Description",
			}
		},
	}
}

func init() { register(RBXMX) }
func RBXMX() rbxmk.Format {
	return rbxmk.Format{
		Name:        "rbxmx",
		EncodeTypes: []string{"Instance", "Objects"},
		MediaTypes:  []string{"application/x-roblox-studio", "application/xml", "text/plain"},
		CanDecode: func(g rbxmk.Global, f rbxmk.FormatOptions, typeName string) bool {
			return typeName == "Instance"
		},
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			desc := descOf(f, "Desc", g, nil)
			mode, err := descModeOf(f, "DescMode")
			if err != nil {
				return nil, err
			}
			d := rbxDecoder{
				method: rbxlx.Deserialize,
				r:      r,
				desc:   desc,
				mode:   mode,
			}
			return d.rbx()
		},
		Encode: func(g rbxmk.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			desc := descOf(f, "Desc", g, v)
			mode, err := descModeOf(f, "DescMode")
			if err != nil {
				return err
			}
			e := rbxEncoder{
				method: rbxlx.Serialize,
				w:      w,
				desc:   desc,
				mode:   mode,
			}
			return e.rbx(v)
		},
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/rbxmx:Summary",
				Description: "Formats/rbxmx:Description",
			}
		},
	}
}

// descMode determines how deviations from a descriptor are handled.
type descMode byte

const (
	// NonStrict causes a deviation to be ignored. The deviation will be dropped
	// as if it never existed.
	modeNonStrict descMode = iota
	// Strict causes an error to be returned on the first deviation.
	modeStrict
	// Preserve tries to retain as much information as possible. Generally, a
	// deviation will be handled as if there was no descriptor specified.
	modePreserve
)

// descModeOf gets the descMode from a given field.
func descModeOf(f rbxmk.FormatOptions, field string) (mode descMode, err error) {
	if v, ok := stringOf(f, field); ok {
		switch v {
		case "NonStrict":
			return modeNonStrict, nil
		case "Strict":
			return modeStrict, nil
		case "Preserve":
			return modePreserve, nil
		}
		return mode, fmt.Errorf("option %s: invalid value %q (expected Strict, NonStrict, or Preserve)", field, v)
	}
	return mode, nil
}

// descOf gets a descriptor from a given field. A RootDesc field returns the
// RootDesc. A false bool returns nil. Otherwise, if v is an Instance, returns
// the descriptor according to g.Desc.Of(v). Otherwise, returns g.Desc.
//
// If v is an Objects, then no descriptor can be reasonably selected, so g.Desc
// is returned instead.
func descOf(f rbxmk.FormatOptions, field string, g rbxmk.Global, v types.Value) (desc *rtypes.RootDesc) {
	if f != nil {
		switch v := f.ValueOf("Desc").(type) {
		case *rtypes.RootDesc:
			return v
		case types.Bool:
			if !v {
				return nil
			}
		}
	}
	switch v := v.(type) {
	case *rtypes.Instance:
		return g.Desc.Of(v)
	case rtypes.Objects:
		// Ambiguous.
	}
	return g.Desc
}

// decinst maps instances for decoding.
type decinst map[*rbxfile.Instance]*rtypes.Instance

// decprop holds a reference property value to be resolved later.
type decprop struct {
	Instance *rtypes.Instance
	Property string
	Value    *rbxfile.Instance
}

// rbxDecoder decodes an rbxfile structure into an rbxmk data model.
type rbxDecoder struct {
	method func(r io.Reader) (root *rbxfile.Root, err error)
	r      io.Reader
	desc   *rtypes.RootDesc
	mode   descMode
	refs   decinst
	prefs  []decprop
}

// rbx decodes d.r according to d.method, then converts the result.
func (d *rbxDecoder) rbx() (v types.Value, err error) {
	root, err := d.method(d.r)
	if err != nil {
		return nil, err
	}
	t, err := d.dataModel(root)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// dataModel converts a root into a data model.
func (d *rbxDecoder) dataModel(r *rbxfile.Root) (t *rtypes.Instance, err error) {
	t = rtypes.NewDataModel()
	meta := t.Metadata()
	for k, v := range r.Metadata {
		meta[k] = v
	}
	d.refs = decinst{}
	d.prefs = []decprop{}
	for _, rc := range r.Instances {
		tc, err := d.instance(rc)
		if err != nil {
			switch d.mode {
			case modeNonStrict:
				continue
			case modeStrict:
				return nil, err
			case modePreserve:
			}
		}
		if tc != nil {
			t.AddChild(tc)
		}
	}
	for _, pref := range d.prefs {
		pref.Instance.Set(pref.Property, d.refs[pref.Value])
	}
	return t, nil
}

// instance converts an instance.
func (d *rbxDecoder) instance(r *rbxfile.Instance) (t *rtypes.Instance, err error) {
	if t, ok := d.refs[r]; ok {
		return t, nil
	}
	if d.desc != nil && d.desc.Class(r.ClassName) == nil {
		switch d.mode {
		case modeNonStrict:
			return nil, nil
		case modeStrict:
			return nil, fmt.Errorf("decode instance: unknown class %q", r.ClassName)
		case modePreserve:
		}
	}
	t = rtypes.NewInstance(r.ClassName, nil)
	t.IsService = r.IsService
	t.Reference = r.Reference
	d.refs[r] = t
	for prop, value := range r.Properties {
		v, err := d.value(t, prop, value)
		if err != nil {
			switch d.mode {
			case modeNonStrict:
				continue
			case modeStrict:
				return nil, fmt.Errorf("property %s.%s: %w", t.ClassName, prop, err)
			case modePreserve:
			}
		}
		if v != nil {
			t.Set(prop, v)
		}
	}
	for _, rc := range r.Children {
		tc, err := d.instance(rc)
		if err != nil {
			switch d.mode {
			case modeNonStrict:
				continue
			case modeStrict:
				return nil, err
			case modePreserve:
			}
		}
		if tc != nil {
			t.AddChild(tc)
		}
	}
	return t, nil
}

// value converts a property value.
func (d *rbxDecoder) value(inst *rtypes.Instance, prop string, r rbxfile.Value) (t types.PropValue, err error) {
	if d.desc != nil {
		if propDesc := d.desc.Property(inst.ClassName, prop); propDesc == nil {
			switch d.mode {
			case modeNonStrict:
			case modeStrict:
				return nil, fmt.Errorf("unknown property")
			case modePreserve:
			}
		} else {
			return d.convertType(inst, prop, r, &propDesc.ValueType)
		}
	}
	return d.convertType(inst, prop, r, nil)
}

// convertType converts a value according to descType, acquired from a
// descriptor. If descType is nil, then the type is determined by the value
// instead.
func (d *rbxDecoder) convertType(inst *rtypes.Instance, prop string, r rbxfile.Value, descType *rbxdump.Type) (t types.PropValue, err error) {
	var typ rbxdump.Type
	if descType != nil {
		typ = *descType
	} else {
		typ = descTypeFromFileValue(r)
	}
	switch {
	case typ == rbxdump.Type{}:
		switch d.mode {
		case modeNonStrict:
			return nil, nil
		case modeStrict:
			return nil, fmt.Errorf("unknown type %q", r.Type())
		case modePreserve:
			// No knowledge of the value type; the best we can do is drop it.
			return nil, nil
		}
	case typ.Category == "Class":
		switch r := r.(type) {
		case rbxfile.ValueReference:
			// Having non-empty Name implies a descriptor.
			if typ.Name != "" {
				if d.desc.Class(typ.Name) == nil {
					// This is a problem with the descriptor, not the encoding.
					// Force the error through by setting to strict mode.
					d.mode = modeStrict
					return nil, fmt.Errorf("descriptor has no definition for class %q", typ.Name)
				}
				if r.Instance != nil && !d.desc.ClassIsA(r.ClassName, typ.Name) {
					switch d.mode {
					case modeNonStrict:
						return nil, nil
					case modeStrict:
						return nil, fmt.Errorf("Reference expected class of %s, got %s", typ.Name, r.ClassName)
					case modePreserve:
						// Decode it anyway.
					}
				}
			}
			d.prefs = append(d.prefs, decprop{
				Instance: inst,
				Property: prop,
				Value:    r.Instance,
			})
			return nil, nil
		default:
			switch d.mode {
			case modeNonStrict:
				return nil, nil
			case modeStrict:
				return nil, fmt.Errorf("cannot decode type %s into Reference", r.Type())
			case modePreserve:
				return d.convertType(inst, prop, r, nil)
			}
		}
	case typ.Category == "Enum":
		if typ.Name != "" {
			switch r := r.(type) {
			case rbxfile.ValueInt:
				t, err = d.enumItemValue(typ.Name, int(r))
			case rbxfile.ValueFloat:
				t, err = d.enumItemValue(typ.Name, int(r))
			case rbxfile.ValueDouble:
				t, err = d.enumItemValue(typ.Name, int(r))
			case rbxfile.ValueBrickColor:
				t, err = d.enumItemValue(typ.Name, int(r))
			case rbxfile.ValueToken:
				t, err = d.enumItemValue(typ.Name, int(r))
			case rbxfile.ValueInt64:
				t, err = d.enumItemValue(typ.Name, int(r))
			case rbxfile.ValueString:
				t, err = d.enumItemName(typ.Name, string(r))
			case rbxfile.ValueBinaryString:
				t, err = d.enumItemName(typ.Name, string(r))
			case rbxfile.ValueProtectedString:
				t, err = d.enumItemName(typ.Name, string(r))
			case rbxfile.ValueContent:
				t, err = d.enumItemName(typ.Name, string(r))
			case rbxfile.ValueSharedString:
				t, err = d.enumItemName(typ.Name, string(r))
			default:
				switch d.mode {
				case modeNonStrict:
					return nil, nil
				case modeStrict:
					return nil, fmt.Errorf("cannot decode type %s into Enum", r.Type())
				case modePreserve:
					return d.convertType(inst, prop, r, nil)
				}
			}
			if err != nil {
				switch d.mode {
				case modeNonStrict:
					return nil, nil
				case modeStrict:
					return nil, err
				case modePreserve:
					return d.convertType(inst, prop, r, nil)
				}
			}
			return t, nil
		}
		if t, ok := d.int(r); ok {
			return types.Token(t), nil
		}
	case typ.Name == "string":
		if t, ok := d.string(r); ok {
			return types.String(t), nil
		}
	case typ.Name == "BinaryString":
		if t, ok := d.string(r); ok {
			return types.BinaryString(t), nil
		}
	case typ.Name == "ProtectedString":
		if t, ok := d.string(r); ok {
			return types.ProtectedString(t), nil
		}
	case typ.Name == "Content":
		if t, ok := d.string(r); ok {
			return types.Content(t), nil
		}
	case typ.Name == "bool":
		switch r := r.(type) {
		case rbxfile.ValueBool:
			return types.Bool(r), nil
		}
	case typ.Name == "int":
		if t, ok := d.int(r); ok {
			return types.Int(t), nil
		}
	case typ.Name == "float":
		if t, ok := d.float(r); ok {
			return types.Float(t), nil
		}
	case typ.Name == "double":
		if t, ok := d.float(r); ok {
			return types.Double(t), nil
		}
	case typ.Name == "UDim":
		switch r := r.(type) {
		case rbxfile.ValueUDim:
			return types.UDim(r), nil
		}
	case typ.Name == "UDim2":
		switch r := r.(type) {
		case rbxfile.ValueUDim2:
			return types.UDim2{
				X: types.UDim(r.X),
				Y: types.UDim(r.Y),
			}, nil
		}
	case typ.Name == "Ray":
		switch r := r.(type) {
		case rbxfile.ValueRay:
			return types.Ray{
				Origin:    types.Vector3(r.Origin),
				Direction: types.Vector3(r.Direction),
			}, nil
		}
	case typ.Name == "Faces":
		switch r := r.(type) {
		case rbxfile.ValueFaces:
			return types.Faces(r), nil
		}
	case typ.Name == "Axes":
		switch r := r.(type) {
		case rbxfile.ValueAxes:
			return types.Axes(r), nil
		}
	case typ.Name == "BrickColor":
		if t, ok := d.int(r); ok {
			return types.BrickColor(t), nil
		}
	case typ.Name == "Color3":
		switch r := r.(type) {
		case rbxfile.ValueColor3:
			return types.Color3(r), nil
		}
	case typ.Name == "Vector2":
		switch r := r.(type) {
		case rbxfile.ValueVector2:
			return types.Vector2(r), nil
		}
	case typ.Name == "Vector3":
		switch r := r.(type) {
		case rbxfile.ValueVector3:
			return types.Vector3(r), nil
		}
	case typ.Name == "CFrame":
		switch r := r.(type) {
		case rbxfile.ValueCFrame:
			return types.CFrame{
				Position: types.Vector3(r.Position),
				Rotation: r.Rotation,
			}, nil
		}
	case typ.Name == "Vector3int16":
		switch r := r.(type) {
		case rbxfile.ValueVector3int16:
			return types.Vector3int16(r), nil
		}
	case typ.Name == "Vector2int16":
		switch r := r.(type) {
		case rbxfile.ValueVector2int16:
			return types.Vector2int16(r), nil
		}
	case typ.Name == "NumberSequence":
		switch r := r.(type) {
		case rbxfile.ValueNumberSequence:
			t := make(types.NumberSequence, len(r))
			for i, k := range r {
				t[i] = types.NumberSequenceKeypoint(k)
			}
			return t, nil
		}
	case typ.Name == "ColorSequence":
		switch r := r.(type) {
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
		}
	case typ.Name == "NumberRange":
		switch r := r.(type) {
		case rbxfile.ValueNumberRange:
			return types.NumberRange(r), nil
		}
	case typ.Name == "Rect":
		switch r := r.(type) {
		case rbxfile.ValueRect:
			return types.Rect{
				Min: types.Vector2(r.Min),
				Max: types.Vector2(r.Max),
			}, nil
		}
	case typ.Name == "PhysicalProperties":
		switch r := r.(type) {
		case rbxfile.ValuePhysicalProperties:
			return types.PhysicalProperties(r), nil
		}
	case typ.Name == "Color3uint8":
		switch r := r.(type) {
		case rbxfile.ValueColor3uint8:
			return rtypes.Color3uint8{
				R: float32(r.R) / 255,
				G: float32(r.G) / 255,
				B: float32(r.B) / 255,
			}, nil
		}
	case typ.Name == "int64":
		if t, ok := d.int(r); ok {
			return types.Int64(t), nil
		}
	case typ.Name == "SharedString":
		if t, ok := d.string(r); ok {
			return types.SharedString(t), nil
		}
	}
	switch d.mode {
	case modeNonStrict:
	case modeStrict:
		return nil, fmt.Errorf("cannot decode type %s into %s", r.Type(), typ.Name)
	case modePreserve:
		if descType != nil {
			// Retry with type derived from value instead of descriptor.
			return d.convertType(inst, prop, r, nil)
		}
		// Value somehow didn't match type through itself.
		panic("unreachable")
	}
	return nil, nil
}

// descTypeFromFileValue returns a descriptor type derived from the given
// rbxfile value. Returns the zero type if the value type is not known.
func descTypeFromFileValue(r rbxfile.Value) rbxdump.Type {
	switch r.(type) {
	case rbxfile.ValueString:
		return rbxdump.Type{Name: "string"}
	case rbxfile.ValueBinaryString:
		return rbxdump.Type{Name: "BinaryString"}
	case rbxfile.ValueProtectedString:
		return rbxdump.Type{Name: "ProtectedString"}
	case rbxfile.ValueContent:
		return rbxdump.Type{Name: "Content"}
	case rbxfile.ValueBool:
		return rbxdump.Type{Name: "bool"}
	case rbxfile.ValueInt:
		return rbxdump.Type{Name: "int"}
	case rbxfile.ValueFloat:
		return rbxdump.Type{Name: "float"}
	case rbxfile.ValueDouble:
		return rbxdump.Type{Name: "double"}
	case rbxfile.ValueUDim:
		return rbxdump.Type{Name: "UDim"}
	case rbxfile.ValueUDim2:
		return rbxdump.Type{Name: "UDim2"}
	case rbxfile.ValueRay:
		return rbxdump.Type{Name: "Ray"}
	case rbxfile.ValueFaces:
		return rbxdump.Type{Name: "Faces"}
	case rbxfile.ValueAxes:
		return rbxdump.Type{Name: "Axes"}
	case rbxfile.ValueBrickColor:
		return rbxdump.Type{Name: "BrickColor"}
	case rbxfile.ValueColor3:
		return rbxdump.Type{Name: "Color3"}
	case rbxfile.ValueVector2:
		return rbxdump.Type{Name: "Vector2"}
	case rbxfile.ValueVector3:
		return rbxdump.Type{Name: "Vector3"}
	case rbxfile.ValueCFrame:
		return rbxdump.Type{Name: "CFrame"}
	case rbxfile.ValueToken:
		return rbxdump.Type{Category: "Enum"}
	case rbxfile.ValueReference:
		return rbxdump.Type{Category: "Class"}
	case rbxfile.ValueVector3int16:
		return rbxdump.Type{Name: "Vector3int16"}
	case rbxfile.ValueVector2int16:
		return rbxdump.Type{Name: "Vector2int16"}
	case rbxfile.ValueNumberSequence:
		return rbxdump.Type{Name: "NumberSequence"}
	case rbxfile.ValueColorSequence:
		return rbxdump.Type{Name: "ColorSequence"}
	case rbxfile.ValueNumberRange:
		return rbxdump.Type{Name: "NumberRange"}
	case rbxfile.ValueRect:
		return rbxdump.Type{Name: "Rect"}
	case rbxfile.ValuePhysicalProperties:
		return rbxdump.Type{Name: "PhysicalProperties"}
	case rbxfile.ValueColor3uint8:
		return rbxdump.Type{Name: "Color3uint8"}
	case rbxfile.ValueInt64:
		return rbxdump.Type{Name: "int64"}
	case rbxfile.ValueSharedString:
		return rbxdump.Type{Name: "SharedString"}
	default:
		return rbxdump.Type{}
	}
}

// string converts common string types to a string.
func (d *rbxDecoder) string(r rbxfile.Value) (t string, ok bool) {
	switch r := r.(type) {
	case rbxfile.ValueString:
		return string(r), true
	case rbxfile.ValueBinaryString:
		return string(r), true
	case rbxfile.ValueProtectedString:
		return string(r), true
	case rbxfile.ValueContent:
		return string(r), true
	case rbxfile.ValueSharedString:
		return string(r), true
	default:
		return "", false
	}
}

// int converts common numeric types to an int.
func (d *rbxDecoder) int(r rbxfile.Value) (t int64, ok bool) {
	switch r := r.(type) {
	case rbxfile.ValueInt:
		return int64(r), true
	case rbxfile.ValueFloat:
		return int64(r), true
	case rbxfile.ValueDouble:
		return int64(r), true
	case rbxfile.ValueBrickColor:
		return int64(r), true
	case rbxfile.ValueToken:
		return int64(r), true
	case rbxfile.ValueInt64:
		return int64(r), true
	default:
		return 0, false
	}
}

// float converts common numeric types to a float.
func (d *rbxDecoder) float(r rbxfile.Value) (t float64, ok bool) {
	switch r := r.(type) {
	case rbxfile.ValueInt:
		return float64(r), true
	case rbxfile.ValueFloat:
		return float64(r), true
	case rbxfile.ValueDouble:
		return float64(r), true
	case rbxfile.ValueBrickColor:
		return float64(r), true
	case rbxfile.ValueToken:
		return float64(r), true
	case rbxfile.ValueInt64:
		return float64(r), true
	default:
		return 0, false
	}
}

// enumItemValue attempts to convert to a token by enum item value.
func (d *rbxDecoder) enumItemValue(enum string, value int) (t types.PropValue, err error) {
	enumDesc := d.desc.Enum(enum)
	if enumDesc == nil {
		// This is a problem with the descriptor, not the encoding. Force the
		// error through by setting to strict mode.
		d.mode = modeStrict
		return nil, fmt.Errorf("descriptor has no definition for enum %q", enum)
	}
	for _, item := range enumDesc.Items {
		if item.Value == value {
			return types.Token(value), nil
		}
	}
	return nil, fmt.Errorf("invalid item value %d for enum %s", value, enum)
}

// enumItemName attempts to convert to a token by enum item name.
func (d *rbxDecoder) enumItemName(enum string, name string) (t types.PropValue, err error) {
	enumDesc := d.desc.Enum(enum)
	if enumDesc == nil {
		// This is a problem with the descriptor, not the encoding. Force the
		// error through by setting to strict mode.
		d.mode = modeStrict
		return nil, fmt.Errorf("descriptor has no definition for enum %q", enum)
	}
	if itemDesc := enumDesc.Items[name]; itemDesc != nil {
		return types.Token(itemDesc.Value), nil
	}
	return nil, fmt.Errorf("invalid item name %s for enum %s", name, enum)
}

////////////////////////////////////////////////////////////////////////////////

type encinst map[*rtypes.Instance]*rbxfile.Instance

type encprop struct {
	Instance *rbxfile.Instance
	Property string
	Value    *rtypes.Instance
}

type rbxEncoder struct {
	method func(w io.Writer, root *rbxfile.Root) (err error)
	w      io.Writer
	desc   *rtypes.RootDesc
	mode   descMode
	refs   encinst
	prefs  []encprop
}

// rbx converts v, then encodes the result to e.w according to e.method.
func (e *rbxEncoder) rbx(v types.Value) (err error) {
	var r *rbxfile.Root
	switch v := v.(type) {
	case *rtypes.Instance:
		if !v.IsDataModel() {
			r, err = e.rootInstance(v)
			break
		}
		r, err = e.dataModel(v)
	case rtypes.Objects:
		r, err = e.objects(v)
	default:
		return cannotEncode(v)
	}
	if err != nil {
		return err
	}
	return e.method(e.w, r)
}

// dataModel converts a data model into a root.
func (e *rbxEncoder) dataModel(t *rtypes.Instance) (r *rbxfile.Root, err error) {
	r = rbxfile.NewRoot()
	meta := t.Metadata()
	r.Metadata = make(map[string]string, len(meta))
	for k, v := range meta {
		r.Metadata[k] = v
	}
	e.refs = encinst{}
	e.prefs = []encprop{}
	for _, tc := range t.Children() {
		rc, err := e.instance(tc)
		if err != nil {
			switch e.mode {
			case modeNonStrict:
				continue
			case modeStrict:
				return nil, err
			case modePreserve:
			}
		}
		if rc != nil {
			r.Instances = append(r.Instances, rc)
		}
	}
	for _, pref := range e.prefs {
		pref.Instance.Properties[pref.Property] = rbxfile.ValueReference{Instance: e.refs[pref.Value]}
	}
	return r, nil
}

// rootInstance converts a single instance into a root.
func (e *rbxEncoder) rootInstance(t *rtypes.Instance) (r *rbxfile.Root, err error) {
	r = rbxfile.NewRoot()
	e.refs = encinst{}
	e.prefs = []encprop{}
	rc, err := e.instance(t)
	if err != nil {
		switch e.mode {
		case modeNonStrict:
			return r, nil
		case modeStrict:
			return nil, err
		case modePreserve:
		}
	}
	if rc != nil {
		r.Instances = append(r.Instances, rc)
	}
	for _, pref := range e.prefs {
		pref.Instance.Properties[pref.Property] = rbxfile.ValueReference{Instance: e.refs[pref.Value]}
	}
	return r, nil
}

// objects converts a list of instances into a root.
func (e *rbxEncoder) objects(t rtypes.Objects) (r *rbxfile.Root, err error) {
	r = rbxfile.NewRoot()
	e.refs = encinst{}
	e.prefs = []encprop{}
	for _, tc := range t {
		rc, err := e.instance(tc)
		if err != nil {
			switch e.mode {
			case modeNonStrict:
				continue
			case modeStrict:
				return nil, err
			case modePreserve:
			}
		}
		if rc != nil {
			r.Instances = append(r.Instances, rc)
		}
	}
	for _, pref := range e.prefs {
		pref.Instance.Properties[pref.Property] = rbxfile.ValueReference{Instance: e.refs[pref.Value]}
	}
	return r, nil
}

// instance converts an instance.
func (e *rbxEncoder) instance(t *rtypes.Instance) (r *rbxfile.Instance, err error) {
	if r, ok := e.refs[t]; ok {
		return r, nil
	}
	if e.desc != nil && e.desc.Class(t.ClassName) == nil {
		switch e.mode {
		case modeNonStrict:
			return nil, nil
		case modeStrict:
			return nil, fmt.Errorf("encode instance: unknown class %q", t.ClassName)
		case modePreserve:
		}
	}
	r = rbxfile.NewInstance(t.ClassName)
	r.IsService = t.IsService
	r.Reference = t.Reference
	e.refs[t] = r
	for prop, value := range t.Properties() {
		v, err := e.value(r, prop, value)
		if err != nil {
			switch e.mode {
			case modeNonStrict:
				continue
			case modeStrict:
				return nil, fmt.Errorf("property %s.%s: %w", t.ClassName, prop, err)
			case modePreserve:
			}
		}
		if v != nil {
			r.Properties[prop] = v
		}
	}
	for _, tc := range t.Children() {
		rc, err := e.instance(tc)
		if err != nil {
			switch e.mode {
			case modeNonStrict:
				continue
			case modeStrict:
				return nil, err
			case modePreserve:
			}
		}
		if rc != nil {
			r.Children = append(r.Children, rc)
		}
	}
	return r, nil
}

// value converts a property value.
func (e *rbxEncoder) value(inst *rbxfile.Instance, prop string, t types.PropValue) (r rbxfile.Value, err error) {
	if e.desc != nil {
		if propDesc := e.desc.Property(inst.ClassName, prop); propDesc == nil {
			switch e.mode {
			case modeNonStrict:
			case modeStrict:
				return nil, fmt.Errorf("unknown property")
			case modePreserve:
			}
		} else {
			return e.convertType(inst, prop, t, &propDesc.ValueType)
		}
	}
	return e.convertType(inst, prop, t, nil)
}

// convertType converts a value according to descType, acquired from a
// descriptor. If descType is nil, then the type is determined by the value
// instead.
func (e *rbxEncoder) convertType(inst *rbxfile.Instance, prop string, t types.PropValue, descType *rbxdump.Type) (rqq rbxfile.Value, err error) {
	var typ rbxdump.Type
	if descType != nil {
		typ = *descType
	} else {
		typ = descTypeFromValue(t)
	}
	switch {
	case typ == rbxdump.Type{}:
		switch e.mode {
		case modeNonStrict:
			return nil, nil
		case modeStrict:
			return nil, fmt.Errorf("unknown type %q", t.Type())
		case modePreserve:
			// No knowledge of the value type; the best we can do is drop it.
			return nil, nil
		}
	case typ.Category == "Class":
		switch t := t.(type) {
		case *rtypes.Instance:
			// Having non-empty Name implies a descriptor.
			if typ.Name != "" {
				if e.desc.Class(typ.Name) == nil {
					// This is a problem with the descriptor, not the encoding.
					// Force the error through by setting to strict mode.
					e.mode = modeStrict
					return nil, fmt.Errorf("descriptor has no definition for class %q", typ.Name)
				}
				if t != nil && !e.desc.ClassIsA(t.ClassName, typ.Name) {
					switch e.mode {
					case modeNonStrict:
						return nil, nil
					case modeStrict:
						return nil, fmt.Errorf("Reference expected class of %s, got %s", typ.Name, t.ClassName)
					case modePreserve:
						// Encode it anyway.
					}
				}
			}
			e.prefs = append(e.prefs, encprop{
				Instance: inst,
				Property: prop,
				Value:    t,
			})
			return nil, nil
		default:
			switch e.mode {
			case modeNonStrict:
				return nil, nil
			case modeStrict:
				return nil, fmt.Errorf("cannot encode type %s into Reference", t.Type())
			case modePreserve:
				return e.convertType(inst, prop, t, nil)
			}
		}
	case typ.Category == "Enum":
		if typ.Name != "" {
			switch t := t.(type) {
			case types.Int:
				rqq, err = e.enumItemValue(typ.Name, int(t))
			case types.Float:
				rqq, err = e.enumItemValue(typ.Name, int(t))
			case types.Double:
				rqq, err = e.enumItemValue(typ.Name, int(t))
			case types.BrickColor:
				rqq, err = e.enumItemValue(typ.Name, int(t))
			case types.Token:
				rqq, err = e.enumItemValue(typ.Name, int(t))
			case types.Int64:
				rqq, err = e.enumItemValue(typ.Name, int(t))
			case types.String:
				rqq, err = e.enumItemName(typ.Name, string(t))
			case types.BinaryString:
				rqq, err = e.enumItemName(typ.Name, string(t))
			case types.ProtectedString:
				rqq, err = e.enumItemName(typ.Name, string(t))
			case types.Content:
				rqq, err = e.enumItemName(typ.Name, string(t))
			case types.SharedString:
				rqq, err = e.enumItemName(typ.Name, string(t))
			default:
				switch e.mode {
				case modeNonStrict:
					return nil, nil
				case modeStrict:
					return nil, fmt.Errorf("cannot encode type %s into Enum", t.Type())
				case modePreserve:
					return e.convertType(inst, prop, t, nil)
				}
			}
			if err != nil {
				switch e.mode {
				case modeNonStrict:
					return nil, nil
				case modeStrict:
					return nil, err
				case modePreserve:
					return e.convertType(inst, prop, t, nil)
				}
			}
			return rqq, nil
		}
		if t, ok := e.int(t); ok {
			return rbxfile.ValueToken(t), nil
		}
	case typ.Name == "string":
		if t, ok := e.string(t); ok {
			return rbxfile.ValueString(t), nil
		}
	case typ.Name == "BinaryString":
		if t, ok := e.string(t); ok {
			return rbxfile.ValueBinaryString(t), nil
		}
	case typ.Name == "ProtectedString":
		if t, ok := e.string(t); ok {
			return rbxfile.ValueProtectedString(t), nil
		}
	case typ.Name == "Content":
		if t, ok := e.string(t); ok {
			return rbxfile.ValueContent(t), nil
		}
	case typ.Name == "bool":
		switch t := t.(type) {
		case types.Bool:
			return rbxfile.ValueBool(t), nil
		}
	case typ.Name == "int":
		if t, ok := e.int(t); ok {
			return rbxfile.ValueInt(t), nil
		}
	case typ.Name == "float":
		if t, ok := e.float(t); ok {
			return rbxfile.ValueFloat(t), nil
		}
	case typ.Name == "double":
		if t, ok := e.float(t); ok {
			return rbxfile.ValueDouble(t), nil
		}
	case typ.Name == "UDim":
		switch t := t.(type) {
		case types.UDim:
			return rbxfile.ValueUDim(t), nil
		}
	case typ.Name == "UDim2":
		switch t := t.(type) {
		case types.UDim2:
			return rbxfile.ValueUDim2{
				X: rbxfile.ValueUDim(t.X),
				Y: rbxfile.ValueUDim(t.Y),
			}, nil
		}
	case typ.Name == "Ray":
		switch t := t.(type) {
		case types.Ray:
			return rbxfile.ValueRay{
				Origin:    rbxfile.ValueVector3(t.Origin),
				Direction: rbxfile.ValueVector3(t.Direction),
			}, nil
		}
	case typ.Name == "Faces":
		switch t := t.(type) {
		case types.Faces:
			return rbxfile.ValueFaces(t), nil
		}
	case typ.Name == "Axes":
		switch t := t.(type) {
		case types.Axes:
			return rbxfile.ValueAxes(t), nil
		}
	case typ.Name == "BrickColor":
		if t, ok := e.int(t); ok {
			return rbxfile.ValueBrickColor(t), nil
		}
	case typ.Name == "Color3":
		switch t := t.(type) {
		case types.Color3:
			return rbxfile.ValueColor3(t), nil
		}
	case typ.Name == "Vector2":
		switch t := t.(type) {
		case types.Vector2:
			return rbxfile.ValueVector2(t), nil
		}
	case typ.Name == "Vector3":
		switch t := t.(type) {
		case types.Vector3:
			return rbxfile.ValueVector3(t), nil
		}
	case typ.Name == "CFrame":
		switch t := t.(type) {
		case types.CFrame:
			return rbxfile.ValueCFrame{
				Position: rbxfile.ValueVector3(t.Position),
				Rotation: t.Rotation,
			}, nil
		}
	case typ.Name == "Vector3int16":
		switch t := t.(type) {
		case types.Vector3int16:
			return rbxfile.ValueVector3int16(t), nil
		}
	case typ.Name == "Vector2int16":
		switch t := t.(type) {
		case types.Vector2int16:
			return rbxfile.ValueVector2int16(t), nil
		}
	case typ.Name == "NumberSequence":
		switch t := t.(type) {
		case types.NumberSequence:
			r := make(rbxfile.ValueNumberSequence, len(t))
			for i, k := range t {
				r[i] = rbxfile.ValueNumberSequenceKeypoint(k)
			}
			return r, nil
		}
	case typ.Name == "ColorSequence":
		switch t := t.(type) {
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
		}
	case typ.Name == "NumberRange":
		switch t := t.(type) {
		case types.NumberRange:
			return rbxfile.ValueNumberRange(t), nil
		}
	case typ.Name == "Rect":
		switch t := t.(type) {
		case types.Rect:
			return rbxfile.ValueRect{
				Min: rbxfile.ValueVector2(t.Min),
				Max: rbxfile.ValueVector2(t.Max),
			}, nil
		}
	case typ.Name == "PhysicalProperties":
		switch t := t.(type) {
		case types.PhysicalProperties:
			return rbxfile.ValuePhysicalProperties(t), nil
		}
	case typ.Name == "Color3uint8":
		switch t := t.(type) {
		case rtypes.Color3uint8:
			return rbxfile.ValueColor3uint8{
				R: uint8(t.R * 255),
				G: uint8(t.G * 255),
				B: uint8(t.B * 255),
			}, nil
		}
	case typ.Name == "int64":
		if t, ok := e.int(t); ok {
			return rbxfile.ValueInt64(t), nil
		}
	case typ.Name == "SharedString":
		if t, ok := e.string(t); ok {
			return rbxfile.ValueSharedString(t), nil
		}
	}
	switch e.mode {
	case modeNonStrict:
	case modeStrict:
		return nil, fmt.Errorf("cannot encode type %s into %s", t.Type(), typ.Name)
	case modePreserve:
		if descType != nil {
			// Retry with type derived from value instead of descriptor.
			return e.convertType(inst, prop, t, nil)
		}
		// Value somehow didn't match type through itself.
		panic("unreachable")
	}
	return nil, nil
}

// descTypeFromValue returns a descriptor type derived from the given value.
// Returns the zero type if the value type is not known.
func descTypeFromValue(t types.PropValue) rbxdump.Type {
	switch t.(type) {
	case types.String:
		return rbxdump.Type{Name: "string"}
	case types.BinaryString:
		return rbxdump.Type{Name: "BinaryString"}
	case types.ProtectedString:
		return rbxdump.Type{Name: "ProtectedString"}
	case types.Content:
		return rbxdump.Type{Name: "Content"}
	case types.Bool:
		return rbxdump.Type{Name: "bool"}
	case types.Int:
		return rbxdump.Type{Name: "int"}
	case types.Float:
		return rbxdump.Type{Name: "float"}
	case types.Double:
		return rbxdump.Type{Name: "double"}
	case types.UDim:
		return rbxdump.Type{Name: "UDim"}
	case types.UDim2:
		return rbxdump.Type{Name: "UDim2"}
	case types.Ray:
		return rbxdump.Type{Name: "Ray"}
	case types.Faces:
		return rbxdump.Type{Name: "Faces"}
	case types.Axes:
		return rbxdump.Type{Name: "Axes"}
	case types.BrickColor:
		return rbxdump.Type{Name: "BrickColor"}
	case types.Color3:
		return rbxdump.Type{Name: "Color3"}
	case types.Vector2:
		return rbxdump.Type{Name: "Vector2"}
	case types.Vector3:
		return rbxdump.Type{Name: "Vector3"}
	case types.CFrame:
		return rbxdump.Type{Name: "CFrame"}
	case types.Token:
		return rbxdump.Type{Category: "Enum"}
	case *rtypes.Instance:
		return rbxdump.Type{Category: "Class"}
	case types.Vector3int16:
		return rbxdump.Type{Name: "Vector3int16"}
	case types.Vector2int16:
		return rbxdump.Type{Name: "Vector2int16"}
	case types.NumberSequence:
		return rbxdump.Type{Name: "NumberSequence"}
	case types.ColorSequence:
		return rbxdump.Type{Name: "ColorSequence"}
	case types.NumberRange:
		return rbxdump.Type{Name: "NumberRange"}
	case types.Rect:
		return rbxdump.Type{Name: "Rect"}
	case types.PhysicalProperties:
		return rbxdump.Type{Name: "PhysicalProperties"}
	case rtypes.Color3uint8:
		return rbxdump.Type{Name: "Color3uint8"}
	case types.Int64:
		return rbxdump.Type{Name: "int64"}
	case types.SharedString:
		return rbxdump.Type{Name: "SharedString"}
	default:
		return rbxdump.Type{}
	}
}

// string converts common string types to a string.
func (e *rbxEncoder) string(t types.PropValue) (r string, ok bool) {
	switch t := t.(type) {
	case types.String:
		return string(t), true
	case types.BinaryString:
		return string(t), true
	case types.ProtectedString:
		return string(t), true
	case types.Content:
		return string(t), true
	case types.SharedString:
		return string(t), true
	default:
		return "", false
	}
}

// int converts common numeric types to an int.
func (e *rbxEncoder) int(t types.PropValue) (r int64, ok bool) {
	switch t := t.(type) {
	case types.Int:
		return int64(t), true
	case types.Float:
		return int64(t), true
	case types.Double:
		return int64(t), true
	case types.BrickColor:
		return int64(t), true
	case types.Token:
		return int64(t), true
	case types.Int64:
		return int64(t), true
	default:
		return 0, false
	}
}

// float converts common numeric types to a float.
func (e *rbxEncoder) float(t types.PropValue) (r float64, ok bool) {
	switch t := t.(type) {
	case types.Int:
		return float64(t), true
	case types.Float:
		return float64(t), true
	case types.Double:
		return float64(t), true
	case types.BrickColor:
		return float64(t), true
	case types.Token:
		return float64(t), true
	case types.Int64:
		return float64(t), true
	default:
		return 0, false
	}
}

// enumItemValue attempts to convert to a token by enum item value.
func (e *rbxEncoder) enumItemValue(enum string, value int) (r rbxfile.Value, err error) {
	enumDesc := e.desc.Enum(enum)
	if enumDesc == nil {
		// This is a problem with the descriptor, not the encoding. Force the
		// error through by setting to strict mode.
		e.mode = modeStrict
		return nil, fmt.Errorf("descriptor has no definition for enum %q", enum)
	}
	for _, item := range enumDesc.Items {
		if item.Value == value {
			return rbxfile.ValueToken(value), nil
		}
	}
	return nil, fmt.Errorf("invalid item value %d for enum %s", value, enum)
}

// enumItemName attempts to convert to a token by enum item name.
func (e *rbxEncoder) enumItemName(enum string, name string) (r rbxfile.Value, err error) {
	enumDesc := e.desc.Enum(enum)
	if enumDesc == nil {
		// This is a problem with the descriptor, not the encoding. Force the
		// error through by setting to strict mode.
		e.mode = modeStrict
		return nil, fmt.Errorf("descriptor has no definition for enum %q", enum)
	}
	if itemDesc := enumDesc.Items[name]; itemDesc != nil {
		return rbxfile.ValueToken(itemDesc.Value), nil
	}
	return nil, fmt.Errorf("invalid item name %s for enum %s", name, enum)
}
