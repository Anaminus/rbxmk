package rtypes

import (
	"github.com/robloxapi/rbxfile"
	"github.com/robloxapi/types"
)

const T_Font = "Font"

// Font represents a font face.
type Font struct {
	Family       string
	Weight       int
	Style        int
	CachedFaceId string
}

// Type returns a string indicating the type of the value.
func (Font) Type() string {
	return T_Font
}

// String returns a string representation of the value.
func (f Font) String() string {
	return rbxfile.ValueFont{
		Family:       rbxfile.ValueContent(f.Family),
		Weight:       rbxfile.FontWeight(f.Weight),
		Style:        rbxfile.FontStyle(f.Style),
		CachedFaceId: rbxfile.ValueContent(f.CachedFaceId),
	}.String()
}

// Copy returns a copy of the value.
func (f Font) Copy() types.PropValue {
	return f
}
