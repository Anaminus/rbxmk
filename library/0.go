package library

import (
	"sort"

	"github.com/anaminus/rbxmk"
)

type entry struct {
	library  rbxmk.Library
	priority int
}

var registry []entry

func register(library rbxmk.Library, priority int) {
	registry = append(registry, entry{
		library:  library,
		priority: priority,
	})
}

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
