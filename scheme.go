package rbxmk

import (
	"fmt"
	"sort"
)

// Scheme represents a rbxmk scheme, either an input scheme, output scheme, or
// both.
type Scheme struct {
	Name   string
	Input  *InputScheme
	Output *OutputScheme
}

// InputScheme represents a rbxmk input scheme.
type InputScheme struct {
	Handler InputSchemeHandler
}

// InputSchemeHandler is used to retrieve a Data from a location. ref is the
// first value of node.Reference after it has been parsed by the protocol
// detector, and excludes the scheme ("scheme://") portion of the string, if
// it was given. Returns the retrieved Data, as well as node.Reference after
// it has been processed.
type InputSchemeHandler func(opt *Options, node *InputNode, inref []string) (outref []string, data Data, err error)

// OutputScheme represents a rbxmk output scheme.
type OutputScheme struct {
	Handler   OutputSchemeHandler // Get current state of output source from location (if needed)
	Finalizer OutputFinalizer     // Write final source to location
}

// OutputSchemeHandler is used to retrieve a Data from a location. ref is the
// first value of node.Reference after it has been parsed by the protocol
// detector, and excludes the scheme ("scheme://") portion of the string, if
// it was given. Returns the retrieved Data, as well as node.Reference after
// it has been processed.
//
// If retrieving the current state of the location is not applicable, then an
// empty or nil Data may be returned.
type OutputSchemeHandler func(opt *Options, node *OutputNode, inref []string) (ext string, outref []string, data Data, err error)

// OutputFinalizer is used to write a modified Data to a location. ref is the
// first value of node.Reference after it has been parsed by the protocol
// detector, and excludes the scheme ("scheme://") portion of the string, if
// it was given.
type OutputFinalizer func(opt *Options, node *OutputNode, inref []string, ext string, outdata Data) (err error)

// Schemes is a container of rbxmk schemes.
type Schemes struct {
	input  map[string]*InputScheme
	output map[string]*OutputScheme
}

// NewSchemes creates and initializes a new Schemes container.
func NewSchemes() *Schemes {
	return &Schemes{
		input:  map[string]*InputScheme{},
		output: map[string]*OutputScheme{},
	}
}

// Copy returns a copy of Schemes. Changes to the copy will not affect the
// original.
func (s *Schemes) Copy() *Schemes {
	c := Schemes{
		input:  make(map[string]*InputScheme, len(s.input)),
		output: make(map[string]*OutputScheme, len(s.output)),
	}
	for k, v := range s.input {
		c.input[k] = v
	}
	for k, v := range s.output {
		c.output[k] = v
	}
	return &c
}

// Register registers a number of rbxmk schemes with the container. Input and
// output schemes may be registered independently. An error is returned if a
// scheme of the same name is already registered.
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

// List returns a list of rbxmk schemes registered with the container. The
// list is sorted by name.
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

// Input returns the InputScheme of a registered scheme with the given name.
// Returns nil if an input scheme is not registered with the name.
func (s *Schemes) Input(name string) *InputScheme {
	return s.input[name]
}

// Output returns the OutputScheme of a registered scheme with the given name.
// Returns nil if an output scheme is not registered with the name.
func (s *Schemes) Output(name string) *OutputScheme {
	return s.output[name]
}
