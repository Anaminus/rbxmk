package rtypes

import (
	"github.com/robloxapi/types"
)

type Color3uint8 types.Color3

func (Color3uint8) Type() string {
	return "Color3uint8"
}

func (c Color3uint8) String() string {
	return types.Color3(c).String()
}

func (c Color3uint8) Copy() types.PropValue {
	return c
}

func (c Color3uint8) Alias() types.Value {
	return types.Color3(c)
}
