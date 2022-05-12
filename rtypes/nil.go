package rtypes

const T_Nil = "nil"

// NilType represents a nil value that implements types.Value.
type NilType struct{}

// Nil is a value of NilType.
var Nil NilType

// Type returns a string identifying the type of the value.
func (NilType) Type() string {
	return T_Nil
}

// String returns a string representation of the value.
func (NilType) String() string {
	return "nil"
}
