package main

var registeredInputSchemes = map[string]InputScheme{}

func RegisterInputScheme(name string, scheme InputScheme) {
	if scheme.Handler == nil {
		panic("input scheme must have handler")
	}
	if _, registered := registeredInputSchemes[name]; registered {
		panic("scheme already registered")
	}
	registeredInputSchemes[name] = scheme
}

var registeredOutputSchemes = map[string]OutputScheme{}

func RegisterOutputScheme(name string, scheme OutputScheme) {
	if scheme.Handler == nil {
		panic("output scheme must have handler")
	}
	if scheme.Finalizer == nil {
		panic("output scheme must have finalizer")
	}
	if _, registered := registeredOutputSchemes[name]; registered {
		panic("scheme already registered")
	}
	registeredOutputSchemes[name] = scheme
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
