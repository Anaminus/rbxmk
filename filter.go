package rbxmk

import (
	"fmt"
	"runtime"
	"sort"
)

// Filter represents a rbxmk filter.
type Filter struct {
	Name string
	Func FilterFunc
}

// FilterFunc defines a signature for a rbxmk filter. A filter can receive an
// arbitrary number of arguments with arbitrary types. To simplify the process
// of checking arguments, FilterArgs is used. While checking arguments, any
// panics that occur will be recovered and returned as an error. After
// f.ProcessedArgs is called, panics go back to being handled as usual.
//
//     func SomeFilter(f FilterArgs, opt *Options, arguments []interface{}) (results []interface{}, err error) {
//         index := arguments[0].(int)
//         value := arguments[1].(string)
//         f.ProcessedArgs()
//         ...
//
type FilterFunc func(f FilterArgs, opt *Options, arguments []interface{}) (results []interface{}, err error)

// FilterArgs is used by a FilterFunc to indicate that it is done processing
// its arguments.
type FilterArgs interface {
	ProcessedArgs()
}

type filterArgs bool

func (f *filterArgs) ProcessedArgs() {
	*f = true
}

// CallFilter invokes a FilterFunc, allowing arguments to be processed without
// panicking.
func CallFilter(filter FilterFunc, opt *Options, arguments ...interface{}) (results []interface{}, err error) {
	var argsProcessed filterArgs
	defer func() {
		if !argsProcessed {
			if e := recover().(runtime.Error); e != nil {
				err = e
			}
		}
	}()
	results, err = filter(&argsProcessed, opt, arguments)
	return
}

// Filters is a container of rbxmk filters.
type Filters struct {
	f map[string]FilterFunc
}

// NewFilters creates and initializes a new Filters container.
func NewFilters() *Filters {
	return &Filters{f: map[string]FilterFunc{}}
}

// Copy returns a copy of Filters. Changes to the copy will not affect the
// original.
func (fs *Filters) Copy() *Filters {
	c := Filters{
		f: make(map[string]FilterFunc, len(fs.f)),
	}
	for k, v := range fs.f {
		c.f[k] = v
	}
	return &c
}

// Register registers a number of rbxmk filters with the container. An error
// is returned if a filter of the same name is already registered.
func (fs *Filters) Register(filters ...Filter) error {
	for _, f := range filters {
		if _, registered := fs.f[f.Name]; registered {
			return fmt.Errorf("filter \"%s\" is already registered", f.Name)
		}
		if f.Func == nil {
			return fmt.Errorf("filter \"%s\" must have Func function", f.Name)
		}
	}
	for _, f := range filters {
		fs.f[f.Name] = f.Func
	}
	return nil
}

// List returns a list of rbxmk filters registered with the container. The
// list is sorted by name.
func (fs *Filters) List() []Filter {
	l := make([]Filter, len(fs.f))
	i := 0
	for name, fn := range fs.f {
		l[i] = Filter{Name: name, Func: fn}
		i++
	}
	sort.Slice(l, func(i, j int) bool {
		return l[i].Name < l[j].Name
	})
	return l
}

// Filter returns the FilterFunc of a registered filter with the given name.
// Returns nil if the name is not registered.
func (fs *Filters) Filter(name string) FilterFunc {
	return fs.f[name]
}
