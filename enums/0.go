package enums

import (
	"github.com/anaminus/rbxmk/rtypes"
)

// registry contains registered Reflectors.
var registry []*rtypes.Enum

// register registers a Reflector to be returned by All.
func register(r *rtypes.Enum) {
	registry = append(registry, r)
}

// All returns a list of Reflectors defined in the package.
func All() []*rtypes.Enum {
	a := make([]*rtypes.Enum, len(registry))
	copy(a, registry)
	return a

}
