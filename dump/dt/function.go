package dt

import "strings"

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

// Parameters is a list of function parameters.
type Parameters = []Parameter

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

// Enums is a list of literal values that can be passed to a function parameter.
type Enums = []string
