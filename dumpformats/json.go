package dumpformats

import (
	"bufio"
	"encoding/json"
	"io"

	"github.com/anaminus/rbxmk/dump"
)

func init() { register(JSON) }

var JSON = Format{
	Name: "json",
	Func: func(w io.Writer, root dump.Root) error {
		return dumpJSON(w, root, "\t")
	},
}

func dumpJSON(w io.Writer, root dump.Root, indent string) error {
	buf := bufio.NewWriter(w)
	j := json.NewEncoder(buf)
	j.SetEscapeHTML(false)
	j.SetIndent("", indent)
	var vroot struct {
		Version int
		dump.Root
	}
	vroot.Version = 0
	vroot.Root = root
	if err := j.Encode(vroot); err != nil {
		return err
	}
	return buf.Flush()
}
