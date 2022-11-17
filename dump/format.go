package dump

import "github.com/anaminus/rbxmk/dump/dt"

// Formats maps a name to a format.
type Formats map[string]Format

// Resolve implements Node.
func (f Formats) Resolve(path ...string) any {
	if len(path) == 0 {
		return f
	}
	if v, ok := f[path[0]]; ok {
		return v.Resolve(path[1:]...)
	}
	return nil
}

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

	// Hidden sets whether the format should be hidden from the public API.
	Hidden bool `json:",omitempty"`
}

// Resolve implements Node.
func (f Format) Resolve(path ...string) any {
	if len(path) == 0 {
		return f
	}
	switch name, path := path[0], path[1:]; name {
	case "Options":
		return f.Options.Resolve(path...)
	}
	return nil
}

// FormatOptions maps a name to a format option.
type FormatOptions map[string]FormatOption

// Resolve implements Node.
func (f FormatOptions) Resolve(path ...string) any {
	if len(path) == 0 {
		return f
	}
	if v, ok := f[path[0]]; ok {
		return resolveValue(path[1:], v)
	}
	return nil
}

type FormatOption struct {
	// Type describes the expected types of the option.
	Type dt.Type
	// Default is a string describing the default value for the option.
	Default string

	// Description is a fragment reference pointing to a detailed description of
	// the option.
	Description string `json:",omitempty"`

	// Hidden sets whether the option should be hidden from the public API.
	Hidden bool `json:",omitempty"`
}
