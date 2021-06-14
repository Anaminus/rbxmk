package dumpformats

import (
	"io"

	"github.com/anaminus/rbxmk/dump"
)

func init() { register(JSONMin) }

var JSONMin = Format{
	Name: "json-min",
	Func: func(w io.Writer, root dump.Root) error {
		return dumpJSON(w, root, "")
	},
}
