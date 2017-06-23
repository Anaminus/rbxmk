package rbxmk

import (
	"runtime"
)

type Filter func(f FilterArgs, opt Options, arguments []interface{}) (results []interface{})

// FilterArgs is used by a Filter to indicate that it is done processing its
// arguments.
type FilterArgs interface {
	ProcessedArgs()
}

type filterArgs bool

func (f *filterArgs) ProcessedArgs() {
	*f = true
}

func CallFilter(filter Filter, opt Options, arguments ...interface{}) (results []interface{}, err error) {
	var argsProcessed filterArgs
	defer func() {
		if !argsProcessed {
			err = recover().(runtime.Error)
		}
	}()
	results = filter(&argsProcessed, opt, arguments)
	return
}

type Filters struct {
	f map[string]Filter
}

func NewFilters() *Filters {
	return &Filters{f: map[string]Filter{}}
}

func (fs *Filters) Register(name string, filter Filter) {
	if filter == nil {
		panic("cannot register nil filter")
	}
	if _, registered := fs.f[name]; registered {
		panic("filter already registered")
	}
	fs.f[name] = filter
}

func (fs *Filters) Filter(name string) Filter {
	return fs.f[name]
}
