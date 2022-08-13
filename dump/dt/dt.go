// The dt package describes the types of Lua API items.
package dt

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func P(t Type) *Type { return &t }

// Kind is implemented by one of several kinds of types.
type Kind interface {
	// Kind returns a string representing the kind of the type.
	Kind() string
	// String returns a readable representation of the type.
	String() string

	k()
}

// Type represents an API type.
type Type struct {
	Kind
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

func (t Type) MarshalJSON() (b []byte, err error) {
	if t.Kind == nil {
		return []byte("null"), nil
	}
	v := map[string]any{
		"Sig":         t.Kind.String(),
		t.Kind.Kind(): t.Kind,
	}
	return marshal(v)
}

func unmarshalKind[K Kind](b []byte, t *Kind) error {
	var k K
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}

func (t *Type) UnmarshalJSON(b []byte) (err error) {
	var v map[string]json.RawMessage
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var kind string
	for k := range v {
		if k == "Sig" {
			continue
		}
		kind = k
		break
	}
	var unmarshal func(b []byte, t *Kind) error
	switch kind {
	case "":
		return fmt.Errorf("missing type kind")
	case K_Primitive:
		unmarshal = unmarshalKind[KindPrim]
	case K_Function:
		unmarshal = unmarshalKind[KindFunction]
	case K_Array:
		unmarshal = unmarshalKind[KindArray]
	case K_Or:
		unmarshal = unmarshalKind[KindOr]
	case K_Optional:
		unmarshal = unmarshalKind[KindOptional]
	case K_Group:
		unmarshal = unmarshalKind[KindGroup]
	case K_Struct:
		unmarshal = unmarshalKind[KindStruct]
	case K_Map:
		unmarshal = unmarshalKind[KindMap]
	case K_Dictionary:
		unmarshal = unmarshalKind[KindDictionary]
	case K_Table:
		unmarshal = unmarshalKind[KindTable]
	case K_Functions:
		unmarshal = unmarshalKind[KindMultiFunctionType]
	default:
		return fmt.Errorf("unknown type kind %q", kind)
	}
	if err := unmarshal(v[kind], &t.Kind); err != nil {
		return fmt.Errorf("kind %s: %w", kind, err)
	}
	return nil
}
