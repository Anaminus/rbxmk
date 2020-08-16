package library

import (
	"sort"

	"github.com/anaminus/rbxmk"
)

// entry contains a Library registered with a priority.
type entry struct {
	library  rbxmk.Library
	priority int
}

// registry contains registered Libraries.
var registry []entry

// register registers a Library to be returned by All.
func register(library rbxmk.Library, priority int) {
	registry = append(registry, entry{
		library:  library,
		priority: priority,
	})
}

// All returns a list of Libraries defined in the package, ordered by ascending
// priority.
func All() []rbxmk.Library {
	sort.SliceStable(registry, func(i, j int) bool {
		return registry[i].priority < registry[j].priority
	})
	libs := make([]rbxmk.Library, len(registry))
	for i, e := range registry {
		libs[i] = e.library
	}
	return libs
}
