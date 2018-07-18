package rbxmk

// Options is used to pass data though many functions in the rbxmk package. It
// provides a context of registered schemes, formats, and filters.
//
// Also included is the Config field, a generic container of values. This may
// be used by other packages to pass extra data along.
type Options struct {
	Schemes *Schemes
	Formats *Formats
	Filters *Filters
	Config  map[string]interface{}
}

// Copy returns a deep copy of the options. Changes to the copy will not
// affect the original.
func (opt *Options) Copy() *Options {
	c := Options{
		Schemes: opt.Schemes.Copy(),
		Formats: opt.Formats.Copy(),
		Filters: opt.Filters.Copy(),
		Config:  make(map[string]interface{}, len(opt.Config)),
	}
	for k, v := range opt.Config {
		c.Config[k] = v
	}
	return &c
}

// NewOptions initializes and returns a new Options.
func NewOptions() *Options {
	return &Options{
		Schemes: NewSchemes(),
		Formats: NewFormats(),
		Filters: NewFilters(),
		Config:  make(map[string]interface{}),
	}
}
