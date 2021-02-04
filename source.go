package rbxmk

// Source defines an external source from which a sequence of bytes can be read
// from or written to. A Source can be registered with a World.
type Source struct {
	// Name is the name of the source.
	Name string

	// Library is a library that provides access to the source. The library is
	// set as a global according Library.Name. If the name is empty, then
	// Source.Name is used instead.
	Library Library
}
