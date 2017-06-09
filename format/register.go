package format

import (
	"github.com/anaminus/rbxmk"
)

var registry []rbxmk.FormatInfo

func register(format rbxmk.FormatInfo) {
	registry = append(registry, format)
}

// Register registers the formats implemented by this package to a given
// rbxmk.Formats.
func Register(formats *rbxmk.Formats) {
	for _, format := range registry {
		formats.Register(format)
	}
}
