package rbxmk

type Source struct {
	// Name is the name of the source.
	Name string

	// Read reads p from the source.
	Read func(options ...interface{}) (p []byte, err error)

	// Write writes p to the source.
	Write func(p []byte, options ...interface{}) (err error)
}
