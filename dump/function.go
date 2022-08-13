package dump

import "github.com/anaminus/rbxmk/dump/dt"

// Methods maps a name to a method.
type Methods = map[string]Function

// Constructors maps a name to a number of constructor functions.
type Constructors = map[string]MultiFunction

// Function describes the API of a function.
type Function struct {
	// CanError returns whether the function may throw an error, excluding type
	// errors from received arguments.
	CanError bool `json:",omitempty"`

	// Summary is a fragment reference pointing to a short summary of the
	// function.
	Summary string `json:",omitempty"`
	// Description is a fragment reference pointing to a detailed description of
	// the function.
	Description string `json:",omitempty"`

	// Parameters are the values received by the function.
	Parameters Parameters `json:",omitempty"`
	// Returns are the values returned by the function.
	Returns Parameters `json:",omitempty"`
}

const V_Function = "Function"

func (v Function) v() {}

func (v Function) Kind() string { return V_Function }

// Type implements Value by returning a dt.Function with the parameters and
// returns of the value.
func (v Function) Type() dt.Type {
	fn := dt.KindFunction{
		Parameters: make(Parameters, len(v.Parameters)),
		Returns:    make(Parameters, len(v.Returns)),
	}
	copy(fn.Parameters, v.Parameters)
	copy(fn.Returns, v.Returns)
	return dt.Function(fn)
}

func (v Function) Index(path []string, name string) ([]string, Value) { return nil, nil }

func (v Function) Indices() []string { return nil }

// MultiFunction describes a Function with multiple signatures.
type MultiFunction []Function

const V_MultiFunction = "MultiFunction"

func (v MultiFunction) v() {}

func (v MultiFunction) Kind() string { return V_MultiFunction }

// Type implements Value by returning dt.MultiFunctionType.
func (MultiFunction) Type() dt.Type {
	return dt.Functions()
}

func (v MultiFunction) Index(path []string, name string) ([]string, Value) { return nil, nil }

func (v MultiFunction) Indices() []string { return nil }

// Parameter describes a function parameter.
type Parameter = dt.Parameter

// Parameters is a list of function parameters.
type Parameters = []Parameter
