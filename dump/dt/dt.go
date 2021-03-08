// The dt package describes the types of Lua API items.
package dt

import (
	"bytes"
	"encoding/json"
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

// Type describes an API type.
type Type interface {
	// String returns a readable representation of the type.
	String() string
	t()
}

// Prim is a Type that indicates the name of some defined type.
type Prim string

func (t Prim) t() {}
func (t Prim) String() string {
	return string(t)
}
func (t Prim) MarshalJSON() (b []byte, err error) {
	v := struct {
		Kind string
		Sig  string
		Type string
	}{"primitive", t.String(), string(t)}
	return marshal(v)
}

// Function is a Type that indicates the signature of a function type.
type Function struct {
	// Parameters are the values received by the function.
	Parameters Parameters `json:",omitempty"`
	// Returns are the values returned by the function.
	Returns Parameters `json:",omitempty"`
}

func (t Function) t() {}
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
	}{"function", t.String(), t.Parameters, t.Returns}
	return marshal(v)
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

// Array is a Type that indicates an array of elements of some type.
type Array struct {
	T Type
}

func (t Array) t() {}
func (t Array) String() string {
	return "{" + t.T.String() + "}"
}
func (t Array) MarshalJSON() (b []byte, err error) {
	v := struct {
		Kind string
		Sig  string
		Type Type
	}{"array", t.String(), t.T}
	return marshal(v)
}

// Or is a Type that indicates a union of two or more types.
type Or []Type

func (t Or) t() {}
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
		Types []Type
	}{"or", t.String(), t}
	return marshal(v)
}

// Optional is a Type that indicates a type of T or nil (shorthand for T | nil).
type Optional struct {
	T Type
}

func (t Optional) t() {}
func (t Optional) String() string {
	return t.T.String() + "?"
}
func (t Optional) MarshalJSON() (b []byte, err error) {
	v := struct {
		Kind string
		Sig  string
		Type Type
	}{"optional", t.String(), t.T}
	return marshal(v)
}

// Group is a Type that ensures the inner type is grouped unambiguously.
type Group struct {
	T Type
}

func (t Group) t() {}
func (t Group) String() string {
	return "(" + t.T.String() + ")"
}
func (t Group) MarshalJSON() (b []byte, err error) {
	v := struct {
		Kind string
		Sig  string
		Type Type
	}{"group", t.String(), t.T}
	return marshal(v)
}

// Struct is a Type that indicates a table with a number of named fields.
type Struct map[string]Type

func (t Struct) t() {}
func (t Struct) String() string {
	f := make([]string, 0, 16)
	for k := range t {
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
	s.WriteByte('}')
	return s.String()
}
func (t Struct) MarshalJSON() (b []byte, err error) {
	v := struct {
		Kind   string
		Sig    string
		Fields map[string]Type
	}{"struct", t.String(), t}
	return marshal(v)
}

// Map is a Type that indicates a table where each element maps a key to a
// value.
type Map struct {
	K Type
	V Type
}

func (t Map) t() {}
func (t Map) String() string {
	return "{[" + t.K.String() + "]: " + t.V.String() + "}"
}
func (t Map) MarshalJSON() (b []byte, err error) {
	v := struct {
		Kind  string
		Sig   string
		Key   Type
		Value Type
	}{"map", t.String(), t.K, t.V}
	return marshal(v)
}

// Dictionary is a Type that indicates a table where each element maps a string
// to a value.
type Dictionary struct {
	V Type
}

func (t Dictionary) t() {}
func (t Dictionary) String() string {
	return "{[string]: " + t.V.String() + "}"
}
func (t Dictionary) MarshalJSON() (b []byte, err error) {
	v := struct {
		Kind  string
		Sig   string
		Value Type
	}{"dictionary", t.String(), t.V}
	return marshal(v)
}

// Table is a Type that indicates a table with both a map part and a struct
// part.
type Table struct {
	Map    Map
	Struct Struct
}

func (t Table) t() {}
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
		Key    Type
		Value  Type
		Fields map[string]Type
	}{"table", t.String(), t.Map.K, t.Map.V, t.Struct}
	return marshal(v)
}

// MultiFunctionType is a Type that indicates a function with multiple
// signatures.
type MultiFunctionType struct{}

func (MultiFunctionType) t() {}
func (MultiFunctionType) String() string {
	return "(...) -> (...)"
}
func (t MultiFunctionType) MarshalJSON() (b []byte, err error) {
	v := struct {
		Kind string
		Sig  string
	}{"functions", t.String()}
	return marshal(v)
}
