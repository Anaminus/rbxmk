package format

import (
	"github.com/anaminus/rbxmk"
)

var registry []rbxmk.FormatInfo

func register(format rbxmk.FormatInfo) {
	registry = append(registry, format)
}

func Register(formats *rbxmk.Formats) {
	for _, format := range registry {
		formats.Register(format)
	}
}
