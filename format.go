package rbxmk

type Format struct {
	// Name is the name that identifies the format. The name matches a file
	// extension.
	Name string

	// Encode receives a value of one of a number of types and encodes it as a
	// sequence of bytes.
	Encode func(v Value) ([]byte, error)

	// Decode receives a sequence of bytes an decodes it into a value of a
	// single type.
	Decode func(b []byte) (Value, error)
}
