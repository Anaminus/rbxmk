package reflect

import (
	"github.com/anaminus/rbxmk"
)

var registry []func() rbxmk.Reflector

func register(r func() rbxmk.Reflector) {
	registry = append(registry, r)
}

func All() []func() rbxmk.Reflector {
	return registry
}
