package sources

import (
	"github.com/anaminus/rbxmk"
)

// registry contains registered Sources.
var registry []func() rbxmk.Source

// register registers a Source to be returned by All.
func register(f func() rbxmk.Source) {
	registry = append(registry, f)
}

// All returns a list of Sources defined in the package.
func All() []func() rbxmk.Source {
	return registry
}
