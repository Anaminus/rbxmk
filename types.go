package rbxmk

import (
	"fmt"
	"github.com/robloxapi/rbxapi"
	"github.com/robloxapi/rbxapi/dump"
	"os"
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

func LoadAPI(path string) (api *rbxapi.API, err error) {
	if path != "" {
		file, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("failed to open API file: %s", err)
		}
		defer file.Close()
		if api, err = dump.Decode(file); err != nil {
			return nil, fmt.Errorf("failed to decode API file: %s", err)
		}
	}
	return api, nil
}
