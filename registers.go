package main

// Receives a Node and a reference string. ref is n.Reference after it has been
// parsed by the protocol detector, and excludes the scheme ("scheme://")
// portion of the string, if it was given.
type InputSchemeHandler func(opt *Options, node *InputNode, ref string) (src *Source, err error)

var registeredInputSchemes = map[string]InputSchemeHandler{}

func RegisterInputScheme(name string, handler InputSchemeHandler) {
	if handler == nil {
		panic("cannot register nil scheme handler")
	}
	if _, registered := registeredInputSchemes[name]; registered {
		panic("scheme already registered")
	}
	registeredInputSchemes[name] = handler
}

// Receives a Node and a reference string. ref is n.Reference after it has been
// parsed by the protocol detector, and excludes the scheme ("scheme://")
// portion of the string, if it was given. Also receives an input source.
type OutputSchemeHandler func(opt *Options, node *OutputNode, ref string, src *Source) (err error)

var registeredOutputSchemes = map[string]OutputSchemeHandler{}

func RegisterOutputScheme(name string, handler OutputSchemeHandler) {
	if handler == nil {
		panic("cannot register nil scheme handler")
	}
	if _, registered := registeredOutputSchemes[name]; registered {
		panic("scheme already registered")
	}
	registeredOutputSchemes[name] = handler
}

var registeredFilters = map[string]Filter{}

func RegisterFilter(name string, filter Filter) {
	if filter == nil {
		panic("cannot register nil filter")
	}
	if _, registered := registeredFilters[name]; registered {
		panic("filter already registered")
	}
	registeredFilters[name] = filter
}
