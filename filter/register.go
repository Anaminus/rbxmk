package filter

import (
	"github.com/anaminus/rbxmk"
)

var registry = map[string]rbxmk.Filter{}

func register(name string, filter rbxmk.Filter) {
	registry[name] = filter
}

// Register registers the filters implemented by this package to a given
// rbxmk.Filters.
func Register(filters *rbxmk.Filters) {
	for name, filter := range registry {
		filters.Register(name, filter)
	}
}
