package filter

import (
	"github.com/anaminus/rbxmk"
)

var registry = map[string]rbxmk.Filter{}

func register(name string, filter rbxmk.Filter) {
	registry[name] = filter
}

func Register(filters *rbxmk.Filters) {
	for name, filter := range registry {
		filters.Register(name, filter)
	}
}
