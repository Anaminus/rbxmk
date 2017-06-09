package rbxmk

import (
	"github.com/robloxapi/rbxapi"
)

type Options struct {
	Schemes *Schemes
	Formats *Formats
	Filters *Filters
	API     *rbxapi.API
}

func NewOptions() *Options {
	return &Options{
		Schemes: NewSchemes(),
		Formats: NewFormats(),
		Filters: NewFilters(),
		API:     nil,
	}
}
