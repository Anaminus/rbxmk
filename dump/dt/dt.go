// The dt package describes the types of Lua API items.
package dt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func P(t Type) *Type { return &t }

func marshal(v interface{}) (b []byte, err error) {
	var buf bytes.Buffer
	j := json.NewEncoder(&buf)
	j.SetEscapeHTML(false)
	if err = j.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

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
	switch kind {
	case "":
		return fmt.Errorf("missing type kind")
	case K_Primitive:
		return new(KindPrim).unmarshal(b, &t.Kind)
	case K_Function:
		return new(KindFunction).unmarshal(b, &t.Kind)
	case K_Array:
		return new(KindArray).unmarshal(b, &t.Kind)
	case K_Or:
		return new(KindOr).unmarshal(b, &t.Kind)
	case K_Optional:
		return new(KindOptional).unmarshal(b, &t.Kind)
	case K_Group:
		return new(KindGroup).unmarshal(b, &t.Kind)
	case K_Struct:
		return new(KindStruct).unmarshal(b, &t.Kind)
	case K_Map:
		return new(KindMap).unmarshal(b, &t.Kind)
	case K_Dictionary:
		return new(KindDictionary).unmarshal(b, &t.Kind)
	case K_Table:
		return new(KindTable).unmarshal(b, &t.Kind)
	case K_Functions:
		return new(KindMultiFunctionType).unmarshal(b, &t.Kind)
	default:
		return fmt.Errorf("unknown type kind %q", kind)
	}
}

// KindPrim is a Type that indicates the name of some defined type.
type KindPrim string

const K_Primitive = "primitive"

func Prim(t string) Type { return Type{Kind: KindPrim(t)} }

func (t KindPrim) k()           {}
func (t KindPrim) Kind() string { return K_Primitive }
func (t KindPrim) String() string {
	return string(t)
}
func (k KindPrim) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}

// KindFunction is a Type that indicates the signature of a function type.
type KindFunction struct {
	// Parameters are the values received by the function.
	Parameters Parameters `json:",omitempty"`
	// Returns are the values returned by the function.
	Returns Parameters `json:",omitempty"`
}

const K_Function = "function"

func Function(k KindFunction) Type { return Type{Kind: k} }

func (t KindFunction) k()           {}
func (t KindFunction) Kind() string { return K_Function }
func (t KindFunction) String() string {
	var s strings.Builder
	s.WriteByte('(')
	for i, v := range t.Parameters {
		if i > 0 {
			s.WriteString(", ")
		}
		if v.Name != "" {
			s.WriteString(v.Name)
			s.WriteString(": ")
		}
		s.WriteString(v.Type.String())
	}
	s.WriteString(") -> (")
	for i, v := range t.Returns {
		if i > 0 {
			s.WriteString(", ")
		}
		if v.Name != "" {
			s.WriteString(v.Name)
			s.WriteString(": ")
		}
		s.WriteString(v.Type.String())
	}
	s.WriteByte(')')
	return s.String()
}
func (k KindFunction) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}

// Parameter describes a function parameter.
type Parameter struct {
	// Name is the name of the parameter.
	Name string `json:",omitempty"`
	// Type is the type of the parameter.
	Type Type
	// Default is the default value if the type is optional.
	Default string `json:",omitempty"`
	// Enum contains literal values that can be passed to the parameter.
	Enums Enums `json:",omitempty"`
}

// Parameters is a list of function parameters.
type Parameters = []Parameter

// Enums is a list of literal values that can be passed to a function parameter.
type Enums = []string

// KindArray is a Type that indicates an array of elements of some type.
type KindArray struct {
	Type
}

const K_Array = "array"

func Array(t Type) Type { return Type{Kind: KindArray{Type: t}} }

func (t KindArray) k()           {}
func (t KindArray) Kind() string { return K_Array }
func (t KindArray) String() string {
	return "{" + t.Type.String() + "}"
}
func (k KindArray) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}

// KindOr is a Type that indicates a union of two or more types.
type KindOr []Type

const K_Or = "or"

func Or(t ...Type) Type { return Type{Kind: KindOr(t)} }

func (t KindOr) k()           {}
func (t KindOr) Kind() string { return K_Or }
func (t KindOr) String() string {
	var s strings.Builder
	prim := true
	for _, v := range t {
		if _, ok := v.Kind.(KindPrim); !ok {
			prim = false
			break
		}
	}
	if prim {
		for i, v := range t {
			if i > 0 {
				s.WriteString(" | ")
			}
			s.WriteString(v.String())
		}
	} else {
		for i, v := range t {
			if i > 0 {
				s.WriteByte('|')
			}
			s.WriteString(v.String())
		}
	}
	return s.String()
}
func (k KindOr) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}

// KindOptional is a Type that indicates a type of T or nil (shorthand for T | nil).
type KindOptional struct {
	Type
}

const K_Optional = "optional"

func Optional(t Type) Type { return Type{Kind: KindOptional{Type: t}} }

func (t KindOptional) k()           {}
func (t KindOptional) Kind() string { return K_Optional }
func (t KindOptional) String() string {
	return t.Type.String() + "?"
}
func (k KindOptional) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}

// KindGroup is a Type that ensures the inner type is grouped unambiguously.
type KindGroup struct {
	Type
}

const K_Group = "group"

func Group(t Type) Type { return Type{Kind: KindGroup{Type: t}} }

func (t KindGroup) k()           {}
func (t KindGroup) Kind() string { return K_Group }
func (t KindGroup) String() string {
	return "(" + t.Type.String() + ")"
}
func (k KindGroup) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}

// KindStruct is a Type that indicates a table with a number of named fields.
type KindStruct map[string]Type

const K_Struct = "struct"

func Struct(k KindStruct) Type { return Type{Kind: k} }

func (t KindStruct) k()           {}
func (t KindStruct) Kind() string { return K_Struct }
func (t KindStruct) String() string {
	f := make([]string, 0, 16)
	var variadic Type
	for k, v := range t {
		if k == "..." {
			w := v
			variadic = w
			continue
		}
		f = append(f, k)
	}
	var s strings.Builder
	s.WriteByte('{')
	sort.Strings(f)
	for i, k := range f {
		if i > 0 {
			s.WriteString(", ")
		}
		v := t[k]
		s.WriteString(k)
		s.WriteString(": ")
		s.WriteString(v.String())
	}
	if variadic.Kind != nil {
		if len(f) > 0 {
			s.WriteString(", ")
		}
		s.WriteString("...: ")
		s.WriteString(variadic.String())
	}
	s.WriteByte('}')
	return s.String()
}
func (k KindStruct) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}

// KindMap is a Type that indicates a table where each element maps a key to a
// value.
type KindMap struct {
	K Type
	V Type
}

const K_Map = "map"

func Map(k, v Type) Type { return Type{Kind: KindMap{K: k, V: v}} }

func (t KindMap) k()           {}
func (t KindMap) Kind() string { return K_Map }
func (t KindMap) String() string {
	return "{[" + t.K.String() + "]: " + t.V.String() + "}"
}
func (k KindMap) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}

// KindDictionary is a Type that indicates a table where each element maps a
// string to a value.
type KindDictionary struct {
	Type
}

const K_Dictionary = "dictionary"

func Dictionary(v Type) Type { return Type{Kind: KindDictionary{Type: v}} }

func (t KindDictionary) k()           {}
func (t KindDictionary) Kind() string { return K_Dictionary }
func (t KindDictionary) String() string {
	return "{[string]: " + t.Type.String() + "}"
}
func (k KindDictionary) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}

// KindTable is a Type that indicates a table with both a map part and a struct
// part.
type KindTable struct {
	Key    Type
	Value  Type
	Fields KindStruct
}

const K_Table = "table"

func Table(k KindTable) Type { return Type{Kind: k} }

func (t KindTable) k()           {}
func (t KindTable) Kind() string { return K_Table }
func (t KindTable) String() string {
	f := make([]string, 0, 16)
	for k := range t.Fields {
		f = append(f, k)
	}
	sort.Strings(f)
	var s strings.Builder
	s.WriteString("{[")
	s.WriteString(t.Key.String())
	s.WriteString("]: ")
	s.WriteString(t.Value.String())
	for _, k := range f {
		s.WriteString(", ")
		v := t.Fields[k]
		s.WriteString(k)
		s.WriteString(": ")
		s.WriteString(v.String())
	}
	s.WriteByte('}')
	return s.String()
}
func (k KindTable) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}

// KindMultiFunctionType is a Type that indicates a function with multiple
// signatures.
type KindMultiFunctionType struct{}

const K_Functions = "functions"

func Functions() Type { return Type{Kind: KindMultiFunctionType{}} }

func (KindMultiFunctionType) k()           {}
func (KindMultiFunctionType) Kind() string { return K_Functions }
func (KindMultiFunctionType) String() string {
	return "(...) -> (...)"
}
func (k KindMultiFunctionType) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}
