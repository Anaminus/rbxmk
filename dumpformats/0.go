package dumpformats

import (
	"io"
	"sort"

	"github.com/anaminus/rbxmk/dump"
)

// registry contains registered Formats.
var registry = map[string]Format{}

// register registers a Format to be returned by All.
func register(format Format) {
	if _, ok := registry[format.Name]; ok {
		panic(format.Name + " already registered")
	}
	registry[format.Name] = format
}

// All returns a list of Formats defined in the package.
func All() Formats {
	formats := make(Formats, 0, len(registry))
	for _, format := range registry {
		formats = append(formats, format)
	}
	sort.Sort(formats)
	return formats
}

type Options map[string]interface{}

// Format specifies how to format an API dump.
type Format struct {
	Name    string
	Func    func(io.Writer, dump.Root, Options) error
	Options Options
}

// Formats is a list of Format values.
type Formats []Format

func (f Formats) Len() int {
	return len(f)
}

func (f Formats) Less(i, j int) bool {
	return f[i].Name < f[j].Name
}

func (f Formats) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// Get returns the format corresponding to the given name.
func (f Formats) Get(name string) (format Format, ok bool) {
	for _, format := range f {
		if format.Name == name {
			return format, true
		}
	}
	return format, false
}

func isName(s string) bool {
	if s == "" {
		return false
	}
	if r := s[0]; !('A' <= r && r <= 'Z' || 'a' <= r && r <= 'z' || r == '_') {
		return false
	}
	for i := 1; i < len(s); i++ {
		if r := s[i]; !('0' <= r && r <= '9' || 'A' <= r && r <= 'Z' || 'a' <= r && r <= 'z' || r == '_') {
			return false
		}
	}
	return true
}

func sortFormats(m dump.Formats, cb func(k string, v dump.Format)) {
	if len(m) == 0 {
		return
	}
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		cb(key, m[key])
	}
}

func sortTypeDefs(m dump.TypeDefs, cb func(k string, v dump.TypeDef)) {
	if len(m) == 0 {
		return
	}
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		cb(key, m[key])
	}
}

func sortProperties(m dump.Properties, cb func(k string, v dump.Property)) {
	if len(m) == 0 {
		return
	}
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		cb(key, m[key])
	}
}

func sortMethods(m dump.Methods, cb func(k string, v dump.Function)) {
	if len(m) == 0 {
		return
	}
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		cb(key, m[key])
	}
}

func sortFields(m dump.Fields, cb func(k string, v dump.Value)) {
	if len(m) == 0 {
		return
	}
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		cb(key, m[key])
	}
}

func sortConstructors(m dump.Constructors, cb func(k string, v dump.MultiFunction)) {
	if len(m) == 0 {
		return
	}
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		cb(key, m[key])
	}
}
