package rbxmk

type InputScheme struct {
	Handler InputSchemeHandler
}

// InputSchemeHandler is used to retrieve a Source from a location. ref is the
// first value of node.Reference after it has been parsed by the protocol
// detector, and excludes the scheme ("scheme://") portion of the string, if
// it was given. Returns the retrieved Source, as well as node.Reference after
// it has been processed.
//
// If InputSchemeHandler used a Format to retrieve the source, then it ensures
// that node.Format is set to the Format's extension.
type InputSchemeHandler func(opt *Options, node *InputNode, ref string) (ext string, src *Source, err error)

type OutputScheme struct {
	Handler   OutputSchemeHandler // Get current state of output source from location (if needed)
	Finalizer OutputFinalizer     // Write final source to location
}

// OutputSchemeHandler is used to retrieve a Source from a location. ref is
// the first value of node.Reference after it has been parsed by the protocol
// detector, and excludes the scheme ("scheme://") portion of the string, if
// it was given. Returns the retrieved Source, as well as node.Reference after
// it has been processed.
//
// If retrieving the current state of the location is not applicable, then an
// empty or nil Source may be returned.
type OutputSchemeHandler func(opt *Options, node *OutputNode, ref string) (ext string, src *Source, err error)

// OutputFinalizer is used to write a modified Source to a location. ref is
// the first value of node.Reference after it has been parsed by the protocol
// detector, and excludes the scheme ("scheme://") portion of the string, if
// it was given.
type OutputFinalizer func(opt *Options, node *OutputNode, ref, ext string, outsrc *Source) (err error)

type Schemes struct {
	input  map[string]*InputScheme
	output map[string]*OutputScheme
}

var DefaultSchemes = Schemes{
	input:  map[string]*InputScheme{},
	output: map[string]*OutputScheme{},
}

func (s Schemes) RegisterInput(name string, scheme InputScheme) {
	if scheme.Handler == nil {
		panic("input scheme must have handler")
	}
	if _, registered := s.input[name]; registered {
		panic("scheme already registered")
	}
	s.input[name] = &scheme
}

func (s Schemes) RegisterOutput(name string, scheme OutputScheme) {
	if scheme.Handler == nil {
		panic("output scheme must have handler")
	}
	if scheme.Finalizer == nil {
		panic("output scheme must have finalizer")
	}
	if _, registered := s.output[name]; registered {
		panic("scheme already registered")
	}
	s.output[name] = &scheme
}

func (s Schemes) Input(name string) *InputScheme {
	return s.input[name]
}

func (s Schemes) Output(name string) *OutputScheme {
	return s.output[name]
}
