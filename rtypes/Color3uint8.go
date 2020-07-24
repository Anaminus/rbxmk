package rtypes

import (
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

type Color3uint8 struct{ R, G, B uint8 }

func NewColor3uint8(c types.Color3) Color3uint8 {
	return Color3uint8{
		R: uint8(c.R * 255),
		G: uint8(c.G * 255),
		B: uint8(c.B * 255),
	}
}

func (Color3uint8) Type() string            { return "Color3uint8" }
func (c Color3uint8) String() string        { return c.Color3().String() }
func (c Color3uint8) Copy() rbxmk.PropValue { return c }
func (c Color3uint8) Color3() types.Color3 {
	return types.Color3{
		R: float32(c.R) / 255,
		G: float32(c.G) / 255,
		B: float32(c.B) / 255,
	}
}
