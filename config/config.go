// Defines and handles custom fields for rbxmk.Option.Config.
package config

import (
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxapi"
	"github.com/yuin/gopher-lua"
)

// Indicies of config values.
const (
	// API specifies a default rbxapi.API to be used by all functions.
	fieldAPI = iota

	// PPEnvs specifies a list of environment tables to be used by the
	// preprocessor.
	fieldPPEnvs

	// Len is the number of configuration values.
	fieldLen
)

// Order of preprocessor variable environments.
const (
	PPEnvScript  = iota // Defined via script (rbxmk.configure).
	PPEnvCommand        // Defined via --define option.
	PPEnvLen            // Number of environments.
)

// Init initializes the configuration table of a rbxmk.Options.
func Init(opt *rbxmk.Options) {
	opt.Config = make([]interface{}, fieldLen)

	opt.Config[fieldAPI] = (*rbxapi.API)(nil)

	envs := make([]*lua.LTable, PPEnvLen)
	for i := range envs {
		envs[i] = &lua.LTable{Metatable: lua.LNil}
	}
	opt.Config[fieldPPEnvs] = envs
}

// API gets the API config value.
func API(opt rbxmk.Options) *rbxapi.API {
	return opt.Config[fieldAPI].(*rbxapi.API)
}

// SetAPI sets the API config value.
func SetAPI(opt rbxmk.Options, api *rbxapi.API) {
	opt.Config[fieldAPI] = api
}

// PPEnvs gets the PPEnv config value.
func PPEnvs(opt rbxmk.Options) []*lua.LTable {
	return opt.Config[fieldPPEnvs].([]*lua.LTable)
}
