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
	// Enums contains enums defined by the API.
	Enums Enums `json:",omitempty"`
	// Formats contains formats registered by a world.
	Formats Formats
	// Program contains the root command created by the program.
	Program Command
	// Fragments contains fragment references requested by the program.
	Fragments []string
	// Description is a fragment reference pointing to a general description of
	// the API.
	Description string `json:",omitempty"`
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
	// Priority determines the order in which the library is loaded.
	Priority int
	// Struct contains the items of the library.
	Struct Struct `json:",omitempty"`
	// Types contains types defined by the library.
	Types TypeDefs `json:",omitempty"`
	// Enums contains enums defined by the library.
	Enums Enums `json:",omitempty"`
}

// Formats maps a name to a format.
type Formats map[string]Format

// Format describes a format.
type Format struct {
	// Summary is a fragment reference pointing to a short summary of the
	// format.
	Summary string `json:",omitempty"`
	// Description is a fragment reference pointing to a detailed description of
	// the format.
	Description string `json:",omitempty"`
}

// Commands maps a name to a command.
type Commands map[string]Command

// Command describes a program command.
type Command struct {
	// Aliases lists available aliases for the command.
	Aliases []string `json:",omitempty"`
	// Hidden indicates whether the command is hidden.
	Hidden bool `json:",omitempty"`
	// Arguments is a fragment reference pointing to a definition of the
	// command's arguments.
	Arguments string `json:",omitempty"`
	// Summary is a fragment reference pointing to a short summary of the
	// command.
	Summary string `json:",omitempty"`
	// Description is a fragment reference pointing to a detailed description of
	// the command.
	Description string `json:",omitempty"`
	// Deprecated is a fragment reference pointing to a message detailing the
	// deprecation of the command.
	Deprecated string `json:",omitempty"`
	// Flags contains the flags defined on the command.
	Flags Flags `json:",omitempty"`
	// Commands contains subcommands defined on the command.
	Commands Commands `json:",omitempty"`
}

// Flags maps a name to a flag.
type Flags map[string]Flag

// Flag describes a command flag.
type Flag struct {
	// Type indicates the value type of the flag.
	Type string
	// Default indicates the default value for the flag.
	Default string `json:",omitempty"`
	// Deprecated indicates whether the flag is deprecated, and if so, a
	// fragment reference pointing to a message describing the deprecation.
	Deprecated string `json:",omitempty"`
	// Hidden indicates whether the flag is hidden.
	Hidden bool `json:",omitempty"`
	// Shorthand indicates a one-letter abbreviation for the flag.
	Shorthand string `json:",omitempty"`
	// ShorthandDeprecated indicates whether the shorthand of the flag is
	// deprecated, and if so, a fragment reference pointing to a message
	// describing the deprecation.
	ShorthandDeprecated string `json:",omitempty"`
	// Description is a fragment reference pointing to a description of the
	// flag.
	Description string `json:",omitempty"`
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
		case Enum:
			f.Kind = "Enum"
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
	// Summary is a fragment reference pointing to a short summary of the
	// property.
	Summary string `json:",omitempty"`
	// Description is a fragment reference pointing to a detailed description of
	// the property.
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
	// Summary is a fragment reference pointing to a short summary of the
	// struct.
	Summary string `json:",omitempty"`
	// Description is a fragment reference pointing to a detailed description of
	// the struct.
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
	// Properties describes the properties defined on the type.
	Properties Properties `json:",omitempty"`
	// Symbols describes the symbols defined on the type.
	Symbols Symbols `json:",omitempty"`
	// Methods describes the methods defined on the type.
	Methods Methods `json:",omitempty"`
	// Constructors describes constructor functions that create the type.
	Constructors Constructors `json:",omitempty"`
	// Enums describes enums related to the type.
	Enums Enums `json:",omitempty"`
	// Summary is a fragment reference pointing to a short summary of the type.
	Summary string `json:",omitempty"`
	// Description is a fragment reference pointing to a detailed description of
	// the type.
	Description string `json:",omitempty"`
	// Requires is a list of names of types that the type depends on.
	Requires []string
}

// Properties maps a name to a Property.
type Properties = map[string]Property

// Symbols maps a name to a Property.
type Symbols = map[string]Property

// Methods maps a name to a method.
type Methods = map[string]Function

// Constructors maps a name to a number of constructor functions.
type Constructors = map[string]MultiFunction

// Enums maps a name to an enum.
type Enums map[string]Enum

// Enum describes the API of an enum.
type Enum struct {
	// Items are the items that exist on the enum.
	Items EnumItems
	// Summary is a fragment reference pointing to a short summary of the enum.
	Summary string `json:",omitempty"`
	// Description is a fragment reference pointing to a detailed description of
	// the enum.
	Description string `json:",omitempty"`
}

// Type implements Value by returning the Enum primitive.
func (v Enum) Type() dt.Type {
	return dt.Prim("Enum")
}

// EnumItems maps a name to an enum.
type EnumItems map[string]EnumItem

// EnumItem describes the API of an enum item.
type EnumItem struct {
	// Value is the value of the item.
	Value int
	// Summary is a fragment reference pointing to a short summary of the enum
	// item.
	Summary string `json:",omitempty"`
	// Description is a fragment reference pointing to a detailed description of
	// the enum item.
	Description string `json:",omitempty"`
}

// Function describes the API of a function.
type Function struct {
	// Parameters are the values received by the function.
	Parameters Parameters `json:",omitempty"`
	// Returns are the values returned by the function.
	Returns Parameters `json:",omitempty"`
	// CanError returns whether the function may throw an error, excluding type
	// errors from received arguments.
	CanError bool `json:",omitempty"`
	// Summary is a fragment reference pointing to a short summary of the
	// function.
	Summary string `json:",omitempty"`
	// Description is a fragment reference pointing to a detailed description of
	// the function.
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
	// Sub describes a number of signatures for the __sub operator.
	Sub []Binop `json:"__sub,omitempty"`
	// Mul describes a number of signatures for the __mul operator.
	Mul []Binop `json:"__mul,omitempty"`
	// Div describes a number of signatures for the __div operator.
	Div []Binop `json:"__div,omitempty"`
	// Mod describes a number of signatures for the __mod operator.
	Mod []Binop `json:"__mod,omitempty"`
	// Pow describes a number of signatures for the __pow operator.
	Pow []Binop `json:"__pow,omitempty"`
	// Concat describes a number of signatures for the __concat operator.
	Concat []Binop `json:"__concat,omitempty"`

	// Eq describes the signature for the __eq operator, if defined.
	Eq *Cmpop `json:"__eq,omitempty"`
	// Le describes the signature for the __le operator, if defined.
	Le *Cmpop `json:"__le,omitempty"`
	// Lt describes the signature for the __lt operator, if defined.
	Lt *Cmpop `json:"__lt,omitempty"`

	// Len describes the signature for the __len operator, if defined.
	Len *Unop `json:"__len,omitempty"`
	// Unm describes the signature for the __unm operator, if defined.
	Unm *Unop `json:"__unm,omitempty"`

	// Call describes the function signature for the __call operator, if
	// defined.
	Call *Function `json:"__call,omitempty"`

	Index    *Function `json:"__index,omitempty"`
	Newindex *Function `json:"__newindex,omitempty"`
}

// Binop describes a binary operator. The left operand is assumed to be of an
// outer type definition.
type Binop struct {
	// Operand is the type of the right operand.
	Operand dt.Type
	// Result is the type of the result of the operation.
	Result dt.Type
	// Summary is a fragment reference pointing to a short summary of the
	// operator.
	Summary string `json:",omitempty"`
	// Description is a fragment reference pointing to a detailed description of
	// the operator.
	Description string `json:",omitempty"`
}

// Cmpop describes a comparison operator. The left and right operands are
// assumed to be of the outer type definition, and a boolean is always returned.
type Cmpop struct {
	// Summary is a fragment reference pointing to a short summary of the
	// operator.
	Summary string `json:",omitempty"`
	// Description is a fragment reference pointing to a detailed description of
	// the operator.
	Description string `json:",omitempty"`
}

// Unop describes a unary operator. The operand is assumed to be of an outer
// type definition.
type Unop struct {
	// Result is the type of the result of the operation.
	Result dt.Type
	// Summary is a fragment reference pointing to a short summary of the
	// operator.
	Summary string `json:",omitempty"`
	// Description is a fragment reference pointing to a detailed description of
	// the operator.
	Description string `json:",omitempty"`
}
