package dump

import "github.com/anaminus/rbxmk/dump/dt"

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
