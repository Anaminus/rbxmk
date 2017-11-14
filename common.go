package rbxmk

import (
	"fmt"
	"github.com/robloxapi/rbxapi"
	"github.com/robloxapi/rbxapi/dump"
	"github.com/yuin/gopher-lua"
	"os"
)

type Config struct {
	API             *rbxapi.API
	PreprocessorEnv *lua.LTable
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
		Config: Config{
			API:             nil,
			PreprocessorEnv: &lua.LTable{Metatable: lua.LNil},
		},
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

func checkStringVar(s string) bool {
	if !('A' <= s[0] && s[0] <= 'Z' ||
		'a' <= s[0] && s[0] <= 'z' ||
		s[0] == '_') {
		return false
	}
	for i := 1; i < len(s); i++ {
		if !('0' <= s[i] && s[i] <= '9' ||
			'A' <= s[i] && s[i] <= 'Z' ||
			'a' <= s[i] && s[i] <= 'z' ||
			s[i] == '_') {
			return false
		}
	}
	return true
}
