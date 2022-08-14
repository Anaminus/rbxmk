package dump

import (
	"sort"
	"strings"
)

// Libraries is a list of libraries.
type Libraries map[string]Library

// Resolve implements Node.
func (l Libraries) Resolve(path ...string) any {
	if len(path) == 0 {
		return l
	}
	if v, ok := l[path[0]]; ok {
		return v.Resolve(path[1:]...)
	}
	return nil
}

type libEntry struct {
	name     string
	priority int
}

type libEntries []libEntry

func (l libEntries) Len() int      { return len(l) }
func (l libEntries) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l libEntries) Less(i, j int) bool {
	if l[i].priority == l[j].priority {
		return l[i].name < l[j].name
	}
	return l[i].priority < l[j].priority
}

// ForEach visits each library, ordered by ascending priority, then by name.
func (l Libraries) ForEach(visit func(name string, library Library) bool) {
	if visit == nil {
		return
	}
	a := make(libEntries, 0, len(l))
	for name, lib := range l {
		a = append(a, libEntry{name: name, priority: lib.Priority})
	}
	sort.Sort(a)
	for _, entry := range a {
		if !visit(entry.name, l[entry.name]) {
			break
		}
	}
}

// Library describes the API of a library.
type Library struct {
	// Import is a path of indices to where the table returned by Open will be
	// merged, starting at the global table. If empty, the table is merged
	// directly into the global table.
	Import []string
	// Priority determines the order in which the library is loaded.
	Priority int
	// Types contains types defined by the library.
	Types TypeDefs `json:",omitempty"`
	// Enums contains enums defined by the library.
	Enums Enums `json:",omitempty"`
	// Struct contains the items of the library.
	Struct Struct `json:",omitempty"`
}

// Resolve implements Node.
func (l Library) Resolve(path ...string) any {
	if len(path) == 0 {
		return l
	}
	switch path[0] {
	case "Types":
		return l.Types.Resolve(path[1:]...)
	case "Enums":
		return l.Enums.Resolve(path[1:]...)
	case "Struct":
		return l.Struct.Resolve(path...)
	}
	return nil
}

func (l Library) ImportString() string {
	return strings.Join(l.Import, ".")
}
