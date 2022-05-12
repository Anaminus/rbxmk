package rtypes

import (
	"github.com/robloxapi/types"
)

const T_Color3uint8 = "Color3uint8"

// Color3uint8 wraps a Color3 value to be interpreted as the Color3uint8 Roblox
// type.
type Color3uint8 types.Color3

// Type returns a string indicating the type of the value.
func (Color3uint8) Type() string {
	return T_Color3uint8
}

// String returns a string representation of the value.
func (c Color3uint8) String() string {
	return types.Color3(c).String()
}

// Copy returns a copy of the value.
func (c Color3uint8) Copy() types.PropValue {
	return c
}

// Alias returns the underlying Color3 value.
func (c Color3uint8) Alias() types.Value {
	return types.Color3(c)
}
