package library

import (
	"sort"

	"github.com/anaminus/rbxmk"
)

type Libraries []rbxmk.Library

func (l Libraries) Len() int      { return len(l) }
func (l Libraries) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l Libraries) Less(i, j int) bool {
	if l[i].Priority == l[j].Priority {
		return l[i].Name < l[j].Name
	}
	return l[i].Priority < l[j].Priority
}

// registry contains registered Libraries.
var registry Libraries

// register registers a Library to be returned by All.
func register(library rbxmk.Library) {
	registry = append(registry, library)
}

// All returns a list of Libraries defined in the package, ordered by ascending
// priority.
func All() Libraries {
	libs := make(Libraries, len(registry))
	copy(libs, registry)
	sort.Sort(libs)
	return libs
}
