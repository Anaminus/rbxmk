package rtypes

import (
	"github.com/robloxapi/types"
)

type NilType struct{}

var Nil NilType

func (NilType) Type() string {
	return "nil"
}

func (NilType) String() string {
	return "nil"
}

func (n NilType) Copy() types.PropValue {
	return n
}
