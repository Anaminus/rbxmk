// The dump package describes Lua APIs.
package dump

import (
	"bytes"
	"encoding/json"

	"github.com/anaminus/rbxmk/dump/dt"
)

func marshal(v interface{}) (b []byte, err error) {
	var buf bytes.Buffer
	j := json.NewEncoder(&buf)
	j.SetEscapeHTML(false)
	if err = j.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Root describes an entire API.
type Root struct {
	// Libraries contains libraries defined in the API.
	Libraries Libraries
	// Types contains types defined by the API.
	Types TypeDefs `json:",omitempty"`
}

// Libraries is a list of libraries.
type Libraries = []Library

// Library describes the API of a library.
type Library struct {
	// Name is the name of the library.
	Name string
	// ImportedAs is the name that the library is imported as. Empty indicates
	// that the contents of the library are merged into the global environment.
	ImportedAs string
	// Struct contains the items of the library.
	Struct Struct `json:",omitempty"`
	// Types contains types defined by the library.
	Types TypeDefs `json:",omitempty"`
}

// Fields maps a name to a value.
type Fields map[string]Value

func (f Fields) MarshalJSON() (b []byte, err error) {
	type field struct {
		Kind  string
		Value Value
	}
	m := make(map[string]field, len(f))
	for k, v := range f {
		f := field{Kind: "", Value: v}
		switch v.(type) {
		case Property:
			f.Kind = "Property"
		case Struct:
			f.Kind = "Struct"
		case Function:
			f.Kind = "Function"
		case MultiFunction:
			f.Kind = "MultiFunction"
		default:
			continue
		}
		m[k] = f
	}
	return marshal(m)
}

// TypeDefs maps a name to a type definition.
type TypeDefs = map[string]TypeDef

// Value is a value that has a Type.
type Value interface {
	Type() dt.Type
}

// Property describes the API of a property.
type Property struct {
	// ValueType is the type of the property's value.
	ValueType dt.Type
	// ReadOnly indicates whether the property can be written to.
	ReadOnly bool `json:",omitempty"`
	// Description is a detailed description of the property.
	Description string `json:",omitempty"`
}

// Type implements Value by returning v.ValueType.
func (v Property) Type() dt.Type {
	return v.ValueType
}

// Struct describes the API of a table with a number of fields.
type Struct struct {
	// Fields are the fields of the structure.
	Fields Fields
	// Description is a detailed description of the structure.
	Description string `json:",omitempty"`
}

// Type implements Value by returning a dt.Struct that maps each field name the
// type of the field's value.
func (v Struct) Type() dt.Type {
	t := make(dt.Struct, len(v.Fields))
	for name, value := range v.Fields {
		t[name] = value.Type()
	}
	return t
}

// TypeDef describes the definition of a type.
type TypeDef struct {
	// Underlying indicates that the type has an underlying type.
	Underlying dt.Type `json:",omitempty"`
	// Operators describes the operators defined on the type.
	Operators *Operators `json:",omitempty"`
	// Operators describes the properties defined on the type.
	Properties Properties `json:",omitempty"`
	// Operators describes the methods defined on the type.
	Methods Methods `json:",omitempty"`
	// Operators describes constructor functions that create the type.
	Constructors Constructors `json:",omitempty"`
	// Description is a detailed description of the type definition.
	Description string `json:",omitempty"`
}

// Properties maps a name to a Property.
type Properties = map[string]Property

// Methods maps a name to a method.
type Methods = map[string]Function

// Constructors maps a name to a number of constructor functions.
type Constructors = map[string]MultiFunction

// Function describes the API of a function.
type Function struct {
	// Parameters are the values received by the function.
	Parameters Parameters `json:",omitempty"`
	// Returns are the values returned by the function.
	Returns Parameters `json:",omitempty"`
	// CanError returns whether the function may throw an error, excluding type
	// errors from received arguments.
	CanError bool `json:",omitempty"`
	// Description is a detailed description of the function.
	Description string `json:",omitempty"`
}

// Type implements Value by returning a dt.Function with the parameters and
// returns of the value.
func (v Function) Type() dt.Type {
	fn := dt.Function{
		Parameters: make(Parameters, len(v.Parameters)),
		Returns:    make(Parameters, len(v.Returns)),
	}
	copy(fn.Parameters, v.Parameters)
	copy(fn.Returns, v.Returns)
	return fn
}

// MultiFunction describes a Function with multiple signatures.
type MultiFunction []Function

// Type implements Value by returning dt.MultiFunctionType.
func (MultiFunction) Type() dt.Type {
	return dt.MultiFunctionType{}
}

// Parameter describes a function parameter.
type Parameter = dt.Parameter

// Parameters is a list of function parameters.
type Parameters = []Parameter

// Operators describes the operators of a type.
type Operators struct {
	// Add describes a number of signatures for the __add operator.
	Add []Binop `json:"__add,omitempty"`
	// Add describes a number of signatures for the __sub operator.
	Sub []Binop `json:"__sub,omitempty"`
	// Add describes a number of signatures for the __mul operator.
	Mul []Binop `json:"__mul,omitempty"`
	// Add describes a number of signatures for the __div operator.
	Div []Binop `json:"__div,omitempty"`
	// Add describes a number of signatures for the __mod operator.
	Mod []Binop `json:"__mod,omitempty"`
	// Add describes a number of signatures for the __pow operator.
	Pow []Binop `json:"__pow,omitempty"`
	// Add describes a number of signatures for the __concat operator.
	Concat []Binop `json:"__concat,omitempty"`

	// Eq indicates whether the type defines a __eq operator.
	Eq bool `json:"__eq,omitempty"`
	// Eq indicates whether the type defines a __le operator.
	Le bool `json:"__le,omitempty"`
	// Eq indicates whether the type defines a __lt operator.
	Lt bool `json:"__lt,omitempty"`

	// Len describes the signature for the __len operator, if defined.
	Len *Unop `json:"__len,omitempty"`
	// Len describes the signature for the __unm operator, if defined.
	Unm *Unop `json:"__unm,omitempty"`

	// Call describes the function signature for the __call operator, if
	// defined.
	Call *Function `json:"__call,omitempty"`

	Index    Value `json:"__index,omitempty"`
	Newindex Value `json:"__newindex,omitempty"`
}

// Binop describes a binary operator. The left operand is assumed to be of an
// outer type definition.
type Binop struct {
	// Operand is the type of the right operand.
	Operand dt.Type
	// Result is the type of the result of the operation.
	Result dt.Type
	// Description is a detailed description of the operator.
	Description string `json:",omitempty"`
}

// Unop describes a unary operator. The operand is assumed to be of an outer
// type definition.
type Unop struct {
	// Result is the type of the result of the operation.
	Result dt.Type
	// Description is a detailed description of the operator.
	Description string `json:",omitempty"`
}
