package dump

import "github.com/anaminus/rbxmk/dump/dt"

// TypeDefs maps a name to a type definition.
type TypeDefs map[string]TypeDef

// Resolve implements Node.
func (t TypeDefs) Resolve(path ...string) any {
	if len(path) == 0 {
		return t
	}
	if v, ok := t[path[0]]; ok {
		return v.Resolve(path[1:]...)
	}
	return nil
}

// TypeDef describes the definition of a type.
type TypeDef struct {
	// Category describes a category for the type.
	Category string `json:",omitempty"`
	// Underlying indicates that the type has an underlying type.
	Underlying *dt.Type `json:",omitempty"`
	// Requires is a list of names of types that the type depends on.
	Requires []string

	// Summary is a fragment reference pointing to a short summary of the type.
	Summary string `json:",omitempty"`
	// Description is a fragment reference pointing to a detailed description of
	// the type.
	Description string `json:",omitempty"`

	// Constructors describes constructor functions that create the type.
	Constructors Constructors `json:",omitempty"`
	// Properties describes the properties defined on the type.
	Properties Properties `json:",omitempty"`
	// Symbols describes the symbols defined on the type.
	Symbols Symbols `json:",omitempty"`
	// Methods describes the methods defined on the type.
	Methods Methods `json:",omitempty"`
	// Operators describes the operators defined on the type.
	Operators *Operators `json:",omitempty"`
	// Enums describes enums related to the type.
	Enums Enums `json:",omitempty"`
}

// Resolve implements Node.
func (t TypeDef) Resolve(path ...string) any {
	if len(path) == 0 {
		return t
	}
	switch name, path := path[0], path[1:]; name {
	case "Constructors":
		return t.Constructors.Resolve(path...)
	case "Properties":
		return t.Properties.Resolve(path...)
	case "Symbols":
		return t.Symbols.Resolve(path...)
	case "Methods":
		return t.Methods.Resolve(path...)
	case "Operators":
		return resolveValue(path, t.Operators)
	case "Enums":
		return t.Enums.Resolve(path...)
	}
	return nil
}
