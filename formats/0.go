package formats

import (
	"fmt"

	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

// registry contains registered Formats.
var registry []func() rbxmk.Format

// register registers a Format to be returned by All.
func register(f func() rbxmk.Format) {
	registry = append(registry, f)
}

// All returns a list of Formats defined in the package.
func All() []func() rbxmk.Format {
	return registry
}

// cannotEncode returns an error indicating that v cannot be encoded.
func cannotEncode(v interface{}) error {
	if v, ok := v.(types.Value); ok {
		return fmt.Errorf("cannot encode %s", v.Type())
	}
	return fmt.Errorf("cannot encode %T", v)
}
