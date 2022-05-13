package dumpformats

import (
	"bufio"
	"encoding/json"
	"io"

	"github.com/anaminus/rbxmk/dump"
)

const F_JSON = "json"

func init() { register(JSON) }

var JSON = Format{
	Name: F_JSON,
	Options: Options{
		"indent": "\t",
	},
	Func: func(w io.Writer, root dump.Root, opts Options) error {
		indent := opts["indent"].(*string)
		return dumpJSON(w, root, *indent)
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
