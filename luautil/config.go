package luautil

import (
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxapi"
	"github.com/yuin/gopher-lua"
)

// Indicies of config values.
const (
	configAPI = iota
	configPPEnvs
	configLen
)

// Order of preprocessor variable environments.
const (
	PPEnvScript  = iota // Defined via script (rbxmk.configure).
	PPEnvCommand        // Defined via --define option.
	PPEnvLen            // Number of environments.
)

func InitConfig(opt *rbxmk.Options) {
	opt.Config = make([]interface{}, configLen)

	opt.Config[configAPI] = (*rbxapi.API)(nil)

	envs := make([]*lua.LTable, PPEnvLen)
	for i := range envs {
		envs[i] = &lua.LTable{Metatable: lua.LNil}
	}
	opt.Config[configPPEnvs] = envs
}

func ConfigAPI(opt rbxmk.Options) *rbxapi.API {
	return opt.Config[configAPI].(*rbxapi.API)
}

func ConfigSetAPI(opt rbxmk.Options, api *rbxapi.API) {
	opt.Config[configAPI] = api
}

func ConfigPPEnvs(opt rbxmk.Options) []*lua.LTable {
	return opt.Config[configPPEnvs].([]*lua.LTable)
}
