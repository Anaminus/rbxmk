package dump

import "github.com/anaminus/rbxmk/dump/dt"

// Formats maps a name to a format.
type Formats map[string]Format

// Format describes a format.
type Format struct {
	// Summary is a fragment reference pointing to a short summary of the
	// format.
	Summary string `json:",omitempty"`
	// Description is a fragment reference pointing to a detailed description of
	// the format.
	Description string `json:",omitempty"`

	// Options describes the options of the format.
	Options FormatOptions `json:",omitempty"`
}

// FormatOptions maps a name to a format option.
type FormatOptions map[string]FormatOption

type FormatOption struct {
	// Type describes the expected types of the option.
	Type dt.Type
	// Default is a string describing the default value for the option.
	Default string

	// Description is a fragment reference pointing to a detailed description of
	// the option.
	Description string `json:",omitempty"`
}
