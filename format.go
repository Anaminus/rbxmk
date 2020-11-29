package rbxmk

import (
	"github.com/robloxapi/types"
)

// Format defines a format for encoding between a sequence of bytes and a
// types.Value. The format can be registered with a World.
type Format struct {
	// Name is the name that identifies the format. The name matches a file
	// extension.
	Name string

	// CanDecode returns whether the format decodes into the given type.
	CanDecode func(typeName string) bool

	// Encode receives a value of one of a number of types and encodes it as a
	// sequence of bytes.
	Encode func(opt FormatOptions, v types.Value) ([]byte, error)

	// Decode receives a sequence of bytes an decodes it into a value of a
	// single type.
	Decode func(opt FormatOptions, b []byte) (types.Value, error)
}

// FormatOptions contains options to be passed to Format.Encode and
// Format.Decode.
type FormatOptions struct {
}
