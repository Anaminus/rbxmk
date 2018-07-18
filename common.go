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
	Config  []interface{}
}

// Copy returns a deep copy of the options. Changes to the copy will not
// affect the original.
func (opt *Options) Copy() *Options {
	cfg := make([]interface{}, len(opt.Config))
	copy(cfg, opt.Config)
	return &Options{
		Schemes: opt.Schemes.Copy(),
		Formats: opt.Formats.Copy(),
		Filters: opt.Filters.Copy(),
		Config:  cfg,
	}
}

// NewOptions initializes and returns a new Options.
func NewOptions() *Options {
	return &Options{
		Schemes: NewSchemes(),
		Formats: NewFormats(),
		Filters: NewFilters(),
	}
}
