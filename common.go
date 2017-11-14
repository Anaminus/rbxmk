package rbxmk

import (
	"github.com/robloxapi/rbxapi"
	"github.com/yuin/gopher-lua"
)

type Config struct {
	API              *rbxapi.API
	PreprocessorEnvs []*lua.LTable
}

type Options struct {
	Schemes *Schemes
	Formats *Formats
	Filters *Filters
	Config  Config
}

func NewOptions() Options {
	return Options{
		Schemes: NewSchemes(),
		Formats: NewFormats(),
		Filters: NewFilters(),
	}
}
