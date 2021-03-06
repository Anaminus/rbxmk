// The dt package describes the types of Lua API items.
package dt

import (
	"sort"
	"strings"
)

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

// Function is a Type that indicates the signature of a function type.
type Function struct {
	// Parameters are the values received by the function.
	Parameters Parameters
	// Returns are the values returned by the function.
	Returns Parameters
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

// Parameter describes a function parameter.
type Parameter struct {
	// Name is the name of the parameter.
	Name string
	// Type is the type of the parameter.
	Type Type
	// Default is the default value if the type is optional.
	Default string
	// Enum contains literal values that can be passed to the parameter.
	Enums Enums
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

// Optional is a Type that indicates a type of T or nil (shorthand for T | nil).
type Optional struct {
	T Type
}

func (t Optional) t() {}
func (t Optional) String() string {
	return t.String() + "?"
}

// Group is a Type that ensures the inner type is grouped unambiguously.
type Group struct {
	T Type
}

func (t Group) t() {}
func (t Group) String() string {
	return "(" + t.String() + ")"
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

// Dictionary is a Type that indicates a table where each element maps a string
// to a value.
type Dictionary struct {
	V Type
}

func (t Dictionary) t() {}
func (t Dictionary) String() string {
	return "{[string]: " + t.V.String() + "}"
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

// MultiFunctionType is a Type that indicates a function with multiple
// signatures.
type MultiFunctionType struct{}

func (MultiFunctionType) t() {}
func (MultiFunctionType) String() string {
	return "(...) -> (...)"
}
