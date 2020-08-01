package library

import (
	"github.com/anaminus/rbxmk"
)

var registry []rbxmk.Library

func register(f rbxmk.Library) {
	registry = append(registry, f)
}

func All() []rbxmk.Library {
	return registry
}
