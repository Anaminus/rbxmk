package rbxmk

type Options struct {
	Schemes *Schemes
	Formats *Formats
	Filters *Filters
	Config  []interface{}
}

func NewOptions() Options {
	return Options{
		Schemes: NewSchemes(),
		Formats: NewFormats(),
		Filters: NewFilters(),
	}
}
