package formats

import (
	"fmt"

	"github.com/anaminus/rbxmk"
)

var registry []func() rbxmk.Format

func register(f func() rbxmk.Format) {
	registry = append(registry, f)
}

func All() []func() rbxmk.Format {
	return registry
}

func cannotEncode(v interface{}) error {
	return fmt.Errorf("cannot encode %T", v)
}
