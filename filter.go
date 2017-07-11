package rbxmk

import (
	"fmt"
	"runtime"
	"sort"
)

type Filter struct {
	Name string
	Func FilterFunc
}

type FilterFunc func(f FilterArgs, opt Options, arguments []interface{}) (results []interface{}, err error)

// FilterArgs is used by a FilterFunc to indicate that it is done processing
// its arguments.
type FilterArgs interface {
	ProcessedArgs()
}

type filterArgs bool

func (f *filterArgs) ProcessedArgs() {
	*f = true
}

func CallFilter(filter FilterFunc, opt Options, arguments ...interface{}) (results []interface{}, err error) {
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

type Filters struct {
	f map[string]FilterFunc
}

func NewFilters() *Filters {
	return &Filters{f: map[string]FilterFunc{}}
}

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

func (fs *Filters) Filter(name string) FilterFunc {
	return fs.f[name]
}
