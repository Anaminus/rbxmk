// The dt package describes the types of Lua API items.
package dt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

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

// T returns a new Type from a Kind.
func T(t Kind) Type { return Type{Kind: t} }

func (t Type) MarshalJSON() (b []byte, err error) {
	v := struct {
		Sig string
		K   string `json:"Kind"`
		Type
	}{Sig: t.String(), K: t.Kind.Kind(), Type: t}
	return marshal(v)
}

func (t *Type) UnmarshalJSON(b []byte) (err error) {
	var v struct{ Kind string }
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch v.Kind {
	case K_Primitive:
		return new(Prim).unmarshal(b, &t.Kind)
	case K_Function:
		return new(Function).unmarshal(b, &t.Kind)
	case K_Array:
		return new(Array).unmarshal(b, &t.Kind)
	case K_Or:
		return new(Or).unmarshal(b, &t.Kind)
	case K_Optional:
		return new(Optional).unmarshal(b, &t.Kind)
	case K_Group:
		return new(Group).unmarshal(b, &t.Kind)
	case K_Struct:
		return new(Struct).unmarshal(b, &t.Kind)
	case K_Map:
		return new(Map).unmarshal(b, &t.Kind)
	case K_Dictionary:
		return new(Dictionary).unmarshal(b, &t.Kind)
	case K_Table:
		return new(Table).unmarshal(b, &t.Kind)
	case K_Functions:
		return new(MultiFunctionType).unmarshal(b, &t.Kind)
	default:
		return fmt.Errorf("unknown type kind %q", v.Kind)
	}
}

// Prim is a Type that indicates the name of some defined type.
type Prim string

const K_Primitive = "primitive"

func (t Prim) k()           {}
func (t Prim) Kind() string { return K_Primitive }
func (t Prim) String() string {
	return string(t)
}
func (t Prim) MarshalJSON() (b []byte, err error) {
	v := struct {
		Kind string
		Sig  string
		Type string
	}{K_Primitive, t.String(), string(t)}
	return marshal(v)
}
func (k Prim) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}

// Function is a Type that indicates the signature of a function type.
type Function struct {
	// Parameters are the values received by the function.
	Parameters Parameters `json:",omitempty"`
	// Returns are the values returned by the function.
	Returns Parameters `json:",omitempty"`
}

const K_Function = "function"

func (t Function) k()           {}
func (t Function) Kind() string { return K_Function }
func (t Function) String() string {
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
func (t Function) MarshalJSON() (b []byte, err error) {
	v := struct {
		Kind       string
		Sig        string
		Parameters Parameters `json:",omitempty"`
		Returns    Parameters `json:",omitempty"`
	}{K_Function, t.String(), t.Parameters, t.Returns}
	return marshal(v)
}
func (k Function) unmarshal(b []byte, t *Kind) error {
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
	Type Kind
	// Default is the default value if the type is optional.
	Default string `json:",omitempty"`
	// Enum contains literal values that can be passed to the parameter.
	Enums Enums `json:",omitempty"`
}

// Parameters is a list of function parameters.
type Parameters = []Parameter

// Enums is a list of literal values that can be passed to a function parameter.
type Enums = []string

// Array is a Type that indicates an array of elements of some type.
type Array struct {
	T Kind
}

const K_Array = "array"

func (t Array) k()           {}
func (t Array) Kind() string { return K_Array }
func (t Array) String() string {
	return "{" + t.T.String() + "}"
}
func (t Array) MarshalJSON() (b []byte, err error) {
	v := struct {
		Kind string
		Sig  string
		Type Kind
	}{K_Array, t.String(), t.T}
	return marshal(v)
}
func (k Array) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}

// Or is a Type that indicates a union of two or more types.
type Or []Kind

const K_Or = "or"

func (t Or) k()           {}
func (t Or) Kind() string { return K_Or }
func (t Or) String() string {
	var s strings.Builder
	prim := true
	for _, v := range t {
		if _, ok := v.(Prim); !ok {
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
func (t Or) MarshalJSON() (b []byte, err error) {
	v := struct {
		Kind  string
		Sig   string
		Types []Kind
	}{K_Or, t.String(), t}
	return marshal(v)
}
func (k Or) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}

// Optional is a Type that indicates a type of T or nil (shorthand for T | nil).
type Optional struct {
	T Kind
}

const K_Optional = "optional"

func (t Optional) k()           {}
func (t Optional) Kind() string { return K_Optional }
func (t Optional) String() string {
	return t.T.String() + "?"
}
func (t Optional) MarshalJSON() (b []byte, err error) {
	v := struct {
		Kind string
		Sig  string
		Type Kind
	}{K_Optional, t.String(), t.T}
	return marshal(v)
}
func (k Optional) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}

// Group is a Type that ensures the inner type is grouped unambiguously.
type Group struct {
	T Kind
}

const K_Group = "group"

func (t Group) k()           {}
func (t Group) Kind() string { return K_Group }
func (t Group) String() string {
	return "(" + t.T.String() + ")"
}
func (t Group) MarshalJSON() (b []byte, err error) {
	v := struct {
		Kind string
		Sig  string
		Type Kind
	}{K_Group, t.String(), t.T}
	return marshal(v)
}
func (k Group) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}

// Struct is a Type that indicates a table with a number of named fields.
type Struct map[string]Kind

const K_Struct = "struct"

func (t Struct) k()           {}
func (t Struct) Kind() string { return K_Struct }
func (t Struct) String() string {
	f := make([]string, 0, 16)
	var variadic Kind
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
	if variadic != nil {
		if len(f) > 0 {
			s.WriteString(", ")
		}
		s.WriteString("...: ")
		s.WriteString(variadic.String())
	}
	s.WriteByte('}')
	return s.String()
}
func (t Struct) MarshalJSON() (b []byte, err error) {
	v := struct {
		Kind   string
		Sig    string
		Fields map[string]Kind
	}{K_Struct, t.String(), t}
	return marshal(v)
}
func (k Struct) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}

// Map is a Type that indicates a table where each element maps a key to a
// value.
type Map struct {
	K Kind
	V Kind
}

const K_Map = "map"

func (t Map) k()           {}
func (t Map) Kind() string { return K_Map }
func (t Map) String() string {
	return "{[" + t.K.String() + "]: " + t.V.String() + "}"
}
func (t Map) MarshalJSON() (b []byte, err error) {
	v := struct {
		Kind  string
		Sig   string
		Key   Kind
		Value Kind
	}{K_Map, t.String(), t.K, t.V}
	return marshal(v)
}
func (k Map) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}

// Dictionary is a Type that indicates a table where each element maps a string
// to a value.
type Dictionary struct {
	V Kind
}

const K_Dictionary = "dictionary"

func (t Dictionary) k()           {}
func (t Dictionary) Kind() string { return K_Dictionary }
func (t Dictionary) String() string {
	return "{[string]: " + t.V.String() + "}"
}
func (t Dictionary) MarshalJSON() (b []byte, err error) {
	v := struct {
		Kind  string
		Sig   string
		Value Kind
	}{K_Dictionary, t.String(), t.V}
	return marshal(v)
}
func (k Dictionary) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}

// Table is a Type that indicates a table with both a map part and a struct
// part.
type Table struct {
	Map    Map
	Struct Struct
}

const K_Table = "table"

func (t Table) k()           {}
func (t Table) Kind() string { return K_Table }
func (t Table) String() string {
	f := make([]string, 0, 16)
	for k := range t.Struct {
		f = append(f, k)
	}
	sort.Strings(f)
	var s strings.Builder
	s.WriteString("{[")
	s.WriteString(t.Map.K.String())
	s.WriteString("]: ")
	s.WriteString(t.Map.V.String())
	for _, k := range f {
		s.WriteString(", ")
		v := t.Struct[k]
		s.WriteString(k)
		s.WriteString(": ")
		s.WriteString(v.String())
	}
	s.WriteByte('}')
	return s.String()
}
func (t Table) MarshalJSON() (b []byte, err error) {
	v := struct {
		Kind   string
		Sig    string
		Key    Kind
		Value  Kind
		Fields map[string]Kind
	}{K_Table, t.String(), t.Map.K, t.Map.V, t.Struct}
	return marshal(v)
}
func (k Table) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}

// MultiFunctionType is a Type that indicates a function with multiple
// signatures.
type MultiFunctionType struct{}

const K_Functions = "functions"

func (MultiFunctionType) k()           {}
func (MultiFunctionType) Kind() string { return K_Functions }
func (MultiFunctionType) String() string {
	return "(...) -> (...)"
}
func (t MultiFunctionType) MarshalJSON() (b []byte, err error) {
	v := struct {
		Kind string
		Sig  string
	}{K_Functions, t.String()}
	return marshal(v)
}
func (k MultiFunctionType) unmarshal(b []byte, t *Kind) error {
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}
	*t = k
	return nil
}
