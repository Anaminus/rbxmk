package library

import (
	"sort"

	"github.com/anaminus/rbxmk"
)

// registry contains registered Libraries.
var registry rbxmk.Libraries

// register registers a Library to be returned by All.
func register(library rbxmk.Library) {
	registry = append(registry, library)
}

// All returns a list of Libraries defined in the package, ordered by ascending
// priority.
func All() rbxmk.Libraries {
	libs := make(rbxmk.Libraries, len(registry))
	copy(libs, registry)
	sort.Sort(libs)
	return libs
}
