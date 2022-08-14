package dump

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/anaminus/rbxmk/dump/dt"
)

// Struct describes the API of a table with a number of fields.
type Struct struct {
	// Summary is a fragment reference pointing to a short summary of the
	// struct.
	Summary string `json:",omitempty"`
	// Description is a fragment reference pointing to a detailed description of
	// the struct.
	Description string `json:",omitempty"`

	// Fields are the fields of the structure.
	Fields Fields
}

const V_Struct = "Struct"

func (v Struct) v() {}

func (v Struct) Kind() string { return V_Struct }

// Type implements Value by returning a dt.Struct that maps each field name the
// type of the field's value.
func (v Struct) Type() dt.Type {
	k := make(dt.KindStruct, len(v.Fields))
	for name, value := range v.Fields {
		k[name] = value.Type()
	}
	return dt.Type{Kind: k}
}

func (v Struct) Index(path []string, name string) ([]string, Value) {
	return append(path, "Struct", "Fields", name), v.Fields[name]
}

func (v Struct) Indices() []string {
	l := make([]string, 0, len(v.Fields))
	for k := range v.Fields {
		l = append(l, k)
	}
	sort.Strings(l)
	return l
}

// Resolve implements Node.
func (s Struct) Resolve(path ...string) any {
	if len(path) == 0 {
		return s
	}
	switch name, path := path[0], path[1:]; name {
	case "Fields":
		return s.Fields.Resolve(path...)
	}
	return nil
}

// Fields maps a name to a value.
type Fields map[string]Value

// Resolve implements Node.
func (f Fields) Resolve(path ...string) any {
	if len(path) == 0 {
		return f
	}
	if v, ok := f[path[0]]; ok {
		if n, ok := v.(Node); ok {
			return n.Resolve(path[1:]...)
		} else {
			return resolveValue(path[1:], v)
		}
	}
	return nil
}

func marshal(v interface{}) (b []byte, err error) {
	var buf bytes.Buffer
	j := json.NewEncoder(&buf)
	j.SetEscapeHTML(false)
	if err = j.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (f Fields) MarshalJSON() (b []byte, err error) {
	type field map[string]Value
	m := make(map[string]field, len(f))
	for k, v := range f {
		f := make(field, 1)
		switch v := v.(type) {
		case Property:
			f[V_Property] = v
		case Struct:
			f[V_Struct] = v
		case Function:
			f[V_Function] = v
		case MultiFunction:
			f[V_MultiFunction] = v
		case Enum:
			f[V_Enum] = v
		default:
			continue
		}
		m[k] = f
	}
	return marshal(m)
}

// Unmarshal b as V, and set to f[k] on success.
func unmarshalValue[V Value](b []byte, f *Fields, k string) error {
	var v V
	if err := json.Unmarshal(b, &v); err != nil {
		return fmt.Errorf("decode value type %s: %w", v.Kind(), err)
	}
	if *f == nil {
		*f = Fields{}
	}
	(*f)[k] = v
	return nil
}

func (f *Fields) UnmarshalJSON(b []byte) (err error) {
	type field map[string]json.RawMessage
	var m map[string]field
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	for k, r := range m {
		var typ string
		for t := range r {
			typ = t
			break
		}
		var unmarshal func(b []byte, f *Fields, k string) error
		switch typ {
		case V_Property:
			unmarshal = unmarshalValue[Property]
		case V_Struct:
			unmarshal = unmarshalValue[Struct]
		case V_Function:
			unmarshal = unmarshalValue[Function]
		case V_MultiFunction:
			unmarshal = unmarshalValue[MultiFunction]
		case V_Enum:
			unmarshal = unmarshalValue[Enum]
		default:
			return fmt.Errorf("field %q: unknown type %q", k, typ)
		}
		if err := unmarshal(r[typ], f, k); err != nil {
			return fmt.Errorf("field %q: %w", k, err)
		}
	}
	return nil
}
