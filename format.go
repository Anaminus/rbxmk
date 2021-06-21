package rbxmk

import (
	"io"

	"github.com/anaminus/rbxmk/dump"
	"github.com/robloxapi/types"
)

// Format defines a format for encoding between a sequence of bytes and a
// types.Value. The format can be registered with a World.
type Format struct {
	// Name is the name that identifies the format. The name matches a file
	// extension.
	Name string

	// EncodeTypes is an optional list of types that Encode can receive. These
	// are called with State.PullAnyOf to reflect the value to a type known by
	// the encoder. If empty, then the value is pulled as Variant.
	EncodeTypes []string

	// MediaTypes is a list of media types that are associated with the format,
	// to be used by sources as needed.
	MediaTypes []string

	// Options maps a field name to a value type. A FormatOptions received by
	// Encode or Decode will have only these fields. The value of a field, if it
	// exists, will be of the specified type.
	Options map[string]string

	// CanDecode returns whether the format decodes into the given type.
	CanDecode func(g Global, opt FormatOptions, typeName string) bool

	// Encode receives a value of one of a number of types and encodes it as a
	// sequence of bytes written to w.
	Encode func(g Global, opt FormatOptions, w io.Writer, v types.Value) error

	// Decode receives a sequence of bytes read from r, and decodes it into a
	// value of a single type.
	Decode func(g Global, opt FormatOptions, r io.Reader) (types.Value, error)

	// Dump returns a description of the format.
	Dump func() dump.Format
}

// FormatOptions contains options to be passed to a Format.
type FormatOptions interface {
	// ValueOf returns the value of field. Returns nil if the value does not
	// exist.
	ValueOf(field string) types.Value
}
