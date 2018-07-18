// Defines and handles custom fields for rbxmk.Option.Config.
package config

import (
	"github.com/anaminus/rbxauth"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxapi"
	"github.com/yuin/gopher-lua"
	"net/http"
)

// Indicies of config values.
const (
	// API specifies a default rbxapi.Root to be used by all functions.
	fieldAPI = iota

	// PPEnvs specifies a list of environment tables to be used by the
	// preprocessor.
	fieldPPEnvs

	// RobloxAuth is a table used to authenticate a request.
	fieldRobloxAuth

	// Host specifies the domain through which downloads and uploads will
	// occur.
	fieldHost

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

	opt.Config[fieldAPI] = (rbxapi.Root)(nil)

	envs := make([]*lua.LTable, PPEnvLen)
	for i := range envs {
		envs[i] = &lua.LTable{Metatable: lua.LNil}
	}
	opt.Config[fieldPPEnvs] = envs

	opt.Config[fieldRobloxAuth] = map[rbxauth.Cred][]*http.Cookie{}

	opt.Config[fieldHost] = "roblox.com"
}

// API gets the API config value.
func API(opt rbxmk.Options) rbxapi.Root {
	return opt.Config[fieldAPI].(rbxapi.Root)
}

// SetAPI sets the API config value.
func SetAPI(opt rbxmk.Options, api rbxapi.Root) {
	opt.Config[fieldAPI] = api
}

// PPEnvs gets the PPEnv config value.
func PPEnvs(opt rbxmk.Options) []*lua.LTable {
	return opt.Config[fieldPPEnvs].([]*lua.LTable)
}

// RobloxAuth gets the RobloxAuth config value.
func RobloxAuth(opt rbxmk.Options) map[rbxauth.Cred][]*http.Cookie {
	return opt.Config[fieldRobloxAuth].(map[rbxauth.Cred][]*http.Cookie)
}

// SetRobloxAuth sets the RobloxAuth config value.
func SetRobloxAuth(opt rbxmk.Options, users map[rbxauth.Cred][]*http.Cookie) {
	opt.Config[fieldRobloxAuth] = users
}

// Host gets the Host config value.
func Host(opt rbxmk.Options) string {
	return opt.Config[fieldHost].(string)
}

// SetHost sets the Host config value.
func SetHost(opt rbxmk.Options, host string) {
	opt.Config[fieldHost] = host
}
