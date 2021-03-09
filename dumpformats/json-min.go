package dumpformats

import (
	"io"

	"github.com/anaminus/rbxmk/dump"
)

func init() { register(JSONMin) }

var JSONMin = Format{
	Name:        "json-min",
	Description: `Minified JSON format.`,
	Func: func(w io.Writer, root dump.Root) error {
		return dumpJSON(w, root, "")
	},
}
