package rbxmk

import (
	"fmt"
	"io"

	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

// Format defines a format for encoding between a sequence of bytes and a
// types.Value. The format can be registered with a World.
type Format struct {
	// Name is the name that identifies the format. The name matches a file
	// extension.
	Name string

	// MediaTypes is a list of media types that are associated with the format,
	// to be used by sources as needed.
	MediaTypes []string

	// Options maps a field name to a value type. A FormatOptions received by
	// Encode or Decode will have only these fields. The value of a field, if it
	// exists, will be of the specified type.
	Options map[string]string

	// CanDecode returns whether the format decodes into the given type.
	CanDecode func(opt FormatOptions, typeName string) bool

	// Encode receives a value of one of a number of types and encodes it as a
	// sequence of bytes written to w.
	Encode func(opt FormatOptions, w io.Writer, v types.Value) error

	// Decode receives a sequence of bytes read from r, and decodes it into a
	// value of a single type.
	Decode func(opt FormatOptions, r io.Reader) (types.Value, error)
}

// FormatOptions contains options to be passed to a Format.
type FormatOptions interface {
	// ValueOf returns the value of field. Returns nil if the value does not
	// exist.
	ValueOf(field string) types.Value
}

// FormatSelector selects a format and provides options for configuring the
// format.
type FormatSelector struct {
	Format  Format
	Options rtypes.Dictionary
}

// Type returns a string identifying the type of the value.
func (FormatSelector) Type() string {
	return "FormatSelector"
}

// ValueOf returns the value of field. Returns nil if the value does not exist.
func (f FormatSelector) ValueOf(field string) types.Value {
	return f.Options[field]
}

// CanDecode calls f.Format.CanDecode with f as the options.
func (f FormatSelector) CanDecode(typeName string) bool {
	if f.Format.CanDecode == nil {
		return false
	}
	return f.Format.CanDecode(f, typeName)
}

// Decode calls f.Format.Decode with f as the options.
func (f FormatSelector) Decode(r io.Reader) (v types.Value, err error) {
	if f.Format.Decode == nil {
		return nil, fmt.Errorf("cannot decode with format %s", f.Format.Name)
	}
	return f.Format.Decode(f, r)
}

// Encode calls f.Format.Encode with f as the options.
func (f FormatSelector) Encode(w io.Writer, v types.Value) error {
	if f.Format.Encode == nil {
		return fmt.Errorf("cannot encode with format %s", f.Format.Name)
	}
	return f.Format.Encode(f, w, v)
}
