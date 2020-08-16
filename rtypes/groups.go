package rtypes

import (
	"github.com/robloxapi/types"
)

// Array is a list of types.Values that itself implements types.Value. It
// corresponds to the Array type in Roblox.
type Array []types.Value

// Type returns a string identifying the type of the value.
func (Array) Type() string {
	return "Array"
}

// Dictionary is a collection of strings mapping to types.Values, that itself
// implements a types.Value. It corresponds to the Dictionary type in Roblox.
type Dictionary map[string]types.Value

// Type returns a string identifying the type of the value.
func (Dictionary) Type() string {
	return "Dictionary"
}

// Tuple is a sequence of types.Values that itself implements types.Value. It
// corresponds to the Tuple type in Roblox.
type Tuple []types.Value

// Type returns a string identifying the type of the value.
func (Tuple) Type() string {
	return "Tuple"
}

// Objects is a list of Instances that implements types.Value. It corresponds to
// the Objects type in Roblox.
type Objects []*Instance

// Type returns a string identifying the type of the value.
func (Objects) Type() string {
	return "Objects"
}
