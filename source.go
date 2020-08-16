package rbxmk

// Source defines an external source from which a sequence of bytes can be read
// from or written to. A Source can be registered with a World.
type Source struct {
	// Name is the name of the source.
	Name string

	// Read implements rbxmk.readSource for the source. The first argument is
	// automatically pulled as the source name. For the implementation of Read,
	// additional arguments can be pulled from s starting at 1.
	Read func(s State) (b []byte, err error)

	// Write implements rbxmk.writeSource for the source. The first and second
	// arguments are automatically pulled as the source name and the bytes to
	// write, respectively. For the implementation of Write, additional
	// arguments can be pulled from s starting at 1.
	Write func(s State, b []byte) (err error)

	// Library is a library that provides access to the source. The library is
	// set as a global according Library.Name. If the name is empty, then
	// Source.Name is used instead.
	Library Library
}
