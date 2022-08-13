package dump

import "github.com/anaminus/rbxmk/dump/dt"

// TypeDefs maps a name to a type definition.
type TypeDefs = map[string]TypeDef

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
