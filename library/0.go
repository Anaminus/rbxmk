package library

import (
	"sort"

	"github.com/anaminus/rbxmk"
)

// registry contains registered Libraries.
var registry []rbxmk.Library

// register registers a Library to be returned by All.
func register(library rbxmk.Library) {
	registry = append(registry, library)
}

// All returns a list of Libraries defined in the package, ordered by ascending
// priority.
func All() []rbxmk.Library {
	libs := make([]rbxmk.Library, len(registry))
	copy(libs, registry)
	sort.SliceStable(libs, func(i, j int) bool {
		if libs[i].Priority == libs[j].Priority {
			return libs[i].Name < libs[j].Name
		}
		return libs[i].Priority < libs[j].Priority
	})
	return libs
}
