package reflect

import (
	"fmt"

	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
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

func setPtrErr(p interface{}, v types.Value) error {
	return fmt.Errorf("cannot set %s to %T", v.Type(), p)
}
