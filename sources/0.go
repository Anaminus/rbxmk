package sources

import (
	"github.com/anaminus/rbxmk"
)

var registry []func() rbxmk.Source

func register(f func() rbxmk.Source) {
	registry = append(registry, f)
}

func All() []func() rbxmk.Source {
	return registry
}
