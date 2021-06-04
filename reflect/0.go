package reflect

import (
	"github.com/anaminus/rbxmk"
)

// registry contains registered Reflectors.
var registry []func() rbxmk.Reflector

// register registers a Reflector to be returned by All.
func register(r func() rbxmk.Reflector) {
	registry = append(registry, r)
}

// All returns a list of Reflectors defined in the package.
func All() []func() rbxmk.Reflector {
	a := make([]func() rbxmk.Reflector, len(registry))
	copy(a, registry)
	return a
}
