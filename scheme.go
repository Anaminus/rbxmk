package rbxmk

import (
	"fmt"
	"sort"
)

type Scheme struct {
	Name   string
	Input  *InputScheme
	Output *OutputScheme
}

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
type InputSchemeHandler func(opt Options, node *InputNode, inref []string) (ext string, outref []string, data Data, err error)

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
type OutputSchemeHandler func(opt Options, node *OutputNode, inref []string) (ext string, outref []string, data Data, err error)

// OutputFinalizer is used to write a modified Source to a location. ref is
// the first value of node.Reference after it has been parsed by the protocol
// detector, and excludes the scheme ("scheme://") portion of the string, if
// it was given.
type OutputFinalizer func(opt Options, node *OutputNode, inref []string, ext string, outdata Data) (err error)

type Schemes struct {
	input  map[string]*InputScheme
	output map[string]*OutputScheme
}

func NewSchemes() *Schemes {
	return &Schemes{
		input:  map[string]*InputScheme{},
		output: map[string]*OutputScheme{},
	}
}

func (s *Schemes) Register(schemes ...Scheme) error {
	for _, scheme := range schemes {
		if scheme.Input == nil && scheme.Output == nil {
			return fmt.Errorf("cannot register empty scheme \"%s\"", scheme.Name)
		}
		if scheme.Input != nil {
			if _, registered := s.input[scheme.Name]; registered {
				return fmt.Errorf("input scheme \"%s\" is already registered", scheme.Name)
			}
			if scheme.Input.Handler == nil {
				return fmt.Errorf("input scheme \"%s\" must have Handler function", scheme.Name)
			}
		}
		if scheme.Output != nil {
			if _, registered := s.output[scheme.Name]; registered {
				return fmt.Errorf("output scheme \"%s\" is already registered", scheme.Name)
			}
			if scheme.Output.Handler == nil {
				return fmt.Errorf("output scheme \"%s\" must have Handler function", scheme.Name)
			}
			if scheme.Output.Finalizer == nil {
				return fmt.Errorf("output scheme \"%s\" must have Finalizer function", scheme.Name)
			}
		}
	}
	for _, scheme := range schemes {
		if scheme.Input != nil {
			input := *scheme.Input
			s.input[scheme.Name] = &input
		}
		if scheme.Output != nil {
			output := *scheme.Output
			s.output[scheme.Name] = &output
		}
	}

	return nil
}

func (s *Schemes) List() []Scheme {
	var schemes map[string]Scheme
	if len(s.input) > len(s.output) {
		schemes = make(map[string]Scheme, len(s.input))
	} else {
		schemes = make(map[string]Scheme, len(s.output))
	}
	for name, inp := range s.input {
		scheme := schemes[name]
		scheme.Name = name
		input := *inp
		scheme.Input = &input
		schemes[name] = scheme
	}
	for name, out := range s.output {
		scheme := schemes[name]
		scheme.Name = name
		output := *out
		scheme.Output = &output
		schemes[name] = scheme
	}

	l := make([]Scheme, len(schemes))
	i := 0
	for _, scheme := range schemes {
		l[i] = scheme
		i++
	}
	sort.Slice(l, func(i, j int) bool {
		return l[i].Name < l[j].Name
	})
	return l
}

func (s *Schemes) Input(name string) *InputScheme {
	return s.input[name]
}

func (s *Schemes) Output(name string) *OutputScheme {
	return s.output[name]
}
