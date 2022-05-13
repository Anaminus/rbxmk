package formats

import (
	"io"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/reflect"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func luaTypes() []func() rbxmk.Reflector {
	return []func() rbxmk.Reflector{
		reflect.Instance,
		reflect.ProtectedString,
	}
}

func decodeScript(r io.Reader, className string) (v types.Value, err error) {
	script := rtypes.NewInstance(className, nil)
	s, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	script.Set("Source", types.ProtectedString(s))
	return script, nil
}

func canDecodeInstance(g rtypes.Global, f rbxmk.FormatOptions, typeName string) bool {
	return typeName == "Instance"
}

func encodeScript(g rtypes.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
	s, ok := rtypes.Stringable(v)
	if !ok {
		return cannotEncode(v)
	}
	_, err := w.Write([]byte(s))
	return err
}

const F_ModuleScriptLua = "modulescript.lua"

func init() { register(ModuleScriptLua) }
func ModuleScriptLua() rbxmk.Format {
	return rbxmk.Format{
		Name:       F_ModuleScriptLua,
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rtypes.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "ModuleScript")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/modulescript.lua:Summary",
				Description: "Formats/modulescript.lua:Description",
			}
		},
		Types: luaTypes(),
	}
}

const F_ScriptLua = "script.lua"

func init() { register(ScriptLua) }
func ScriptLua() rbxmk.Format {
	return rbxmk.Format{
		Name:       F_ScriptLua,
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rtypes.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "Script")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/script.lua:Summary",
				Description: "Formats/script.lua:Description",
			}
		},
		Types: luaTypes(),
	}
}

const F_LocalScriptLua = "localscript.lua"

func init() { register(LocalScriptLua) }
func LocalScriptLua() rbxmk.Format {
	return rbxmk.Format{
		Name:       F_LocalScriptLua,
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rtypes.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "LocalScript")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/localscript.lua:Summary",
				Description: "Formats/localscript.lua:Description",
			}
		},
		Types: luaTypes(),
	}
}

const F_Lua = "lua"

func init() { register(Lua) }
func Lua() rbxmk.Format {
	return rbxmk.Format{
		Name:       F_Lua,
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rtypes.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "ModuleScript")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/lua:Summary",
				Description: "Formats/lua:Description",
			}
		},
		Types: luaTypes(),
	}
}

const F_ServerLua = "server.lua"

func init() { register(ServerLua) }
func ServerLua() rbxmk.Format {
	return rbxmk.Format{
		Name:       F_ServerLua,
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rtypes.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "Script")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/server.lua:Summary",
				Description: "Formats/server.lua:Description",
			}
		},
		Types: luaTypes(),
	}
}

const F_ClientLua = "client.lua"

func init() { register(ClientLua) }
func ClientLua() rbxmk.Format {
	return rbxmk.Format{
		Name:       F_ClientLua,
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rtypes.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "LocalScript")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/client.lua:Summary",
				Description: "Formats/client.lua:Description",
			}
		},
		Types: luaTypes(),
	}
}

////////////////////////////////////////////////////////////////////////////////
// Luau

const F_ModuleScriptLuau = "modulescript.luau"

func init() { register(ModuleScriptLuau) }
func ModuleScriptLuau() rbxmk.Format {
	return rbxmk.Format{
		Name:       F_ModuleScriptLuau,
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rtypes.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "ModuleScript")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/modulescript.luau:Summary",
				Description: "Formats/modulescript.luau:Description",
			}
		},
		Types: luaTypes(),
	}
}

const F_ScriptLuau = "script.luau"

func init() { register(ScriptLuau) }
func ScriptLuau() rbxmk.Format {
	return rbxmk.Format{
		Name:       F_ScriptLuau,
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rtypes.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "Script")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/script.luau:Summary",
				Description: "Formats/script.luau:Description",
			}
		},
		Types: luaTypes(),
	}
}

const F_LocalScriptLuau = "localscript.luau"

func init() { register(LocalScriptLuau) }
func LocalScriptLuau() rbxmk.Format {
	return rbxmk.Format{
		Name:       F_LocalScriptLuau,
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rtypes.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "LocalScript")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/localscript.luau:Summary",
				Description: "Formats/localscript.luau:Description",
			}
		},
		Types: luaTypes(),
	}
}

const F_Luau = "luau"

func init() { register(Luau) }
func Luau() rbxmk.Format {
	return rbxmk.Format{
		Name:       F_Luau,
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rtypes.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "ModuleScript")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/lua:Summary",
				Description: "Formats/lua:Description",
			}
		},
		Types: luaTypes(),
	}
}

const F_ServerLuau = "server.luau"

func init() { register(ServerLuau) }
func ServerLuau() rbxmk.Format {
	return rbxmk.Format{
		Name:       F_ServerLuau,
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rtypes.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "Script")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/server.luau:Summary",
				Description: "Formats/server.luau:Description",
			}
		},
		Types: luaTypes(),
	}
}

const F_ClientLuau = "client.luau"

func init() { register(ClientLuau) }
func ClientLuau() rbxmk.Format {
	return rbxmk.Format{
		Name:       F_ClientLuau,
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rtypes.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "LocalScript")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/client.luau:Summary",
				Description: "Formats/client.luau:Description",
			}
		},
		Types: luaTypes(),
	}
}
