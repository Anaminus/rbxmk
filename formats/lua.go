package formats

import (
	"io"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func decodeScript(r io.Reader, className string) (v types.Value, err error) {
	script := rtypes.NewInstance(className, nil)
	s, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	script.Set("Source", types.ProtectedString(s))
	return script, nil
}

func canDecodeInstance(g rbxmk.Global, f rbxmk.FormatOptions, typeName string) bool {
	return typeName == "Instance"
}

func encodeScript(g rbxmk.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
	s, ok := rtypes.Stringable(v)
	if !ok {
		return cannotEncode(v)
	}
	_, err := w.Write([]byte(s))
	return err
}

func init() { register(ModuleScriptLua) }
func ModuleScriptLua() rbxmk.Format {
	return rbxmk.Format{
		Name:       "modulescript.lua",
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "ModuleScript")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/modulescript.lua:Summary",
				Description: "Formats/modulescript.lua:Description",
			}
		},
	}
}

func init() { register(ScriptLua) }
func ScriptLua() rbxmk.Format {
	return rbxmk.Format{
		Name:       "script.lua",
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "Script")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/script.lua:Summary",
				Description: "Formats/script.lua:Description",
			}
		},
	}
}

func init() { register(LocalScriptLua) }
func LocalScriptLua() rbxmk.Format {
	return rbxmk.Format{
		Name:       "localscript.lua",
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "LocalScript")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/localscript.lua:Summary",
				Description: "Formats/localscript.lua:Description",
			}
		},
	}
}

func init() { register(Lua) }
func Lua() rbxmk.Format {
	return rbxmk.Format{
		Name:       "lua",
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "ModuleScript")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/lua:Summary",
				Description: "Formats/lua:Description",
			}
		},
	}
}

func init() { register(ServerLua) }
func ServerLua() rbxmk.Format {
	return rbxmk.Format{
		Name:       "server.lua",
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "Script")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/server.lua:Summary",
				Description: "Formats/server.lua:Description",
			}
		},
	}
}

func init() { register(ClientLua) }
func ClientLua() rbxmk.Format {
	return rbxmk.Format{
		Name:       "client.lua",
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "LocalScript")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/client.lua:Summary",
				Description: "Formats/client.lua:Description",
			}
		},
	}
}

////////////////////////////////////////////////////////////////////////////////
// Luau

func init() { register(ModuleScriptLuau) }
func ModuleScriptLuau() rbxmk.Format {
	return rbxmk.Format{
		Name:       "modulescript.luau",
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "ModuleScript")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/modulescript.luau:Summary",
				Description: "Formats/modulescript.luau:Description",
			}
		},
	}
}

func init() { register(ScriptLuau) }
func ScriptLuau() rbxmk.Format {
	return rbxmk.Format{
		Name:       "script.luau",
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "Script")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/script.luau:Summary",
				Description: "Formats/script.luau:Description",
			}
		},
	}
}

func init() { register(LocalScriptLuau) }
func LocalScriptLuau() rbxmk.Format {
	return rbxmk.Format{
		Name:       "localscript.luau",
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "LocalScript")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/localscript.luau:Summary",
				Description: "Formats/localscript.luau:Description",
			}
		},
	}
}

func init() { register(Luau) }
func Luau() rbxmk.Format {
	return rbxmk.Format{
		Name:       "lua",
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "ModuleScript")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/lua:Summary",
				Description: "Formats/lua:Description",
			}
		},
	}
}

func init() { register(ServerLuau) }
func ServerLuau() rbxmk.Format {
	return rbxmk.Format{
		Name:       "server.luau",
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "Script")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/server.luau:Summary",
				Description: "Formats/server.luau:Description",
			}
		},
	}
}

func init() { register(ClientLuau) }
func ClientLuau() rbxmk.Format {
	return rbxmk.Format{
		Name:       "client.luau",
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "LocalScript")
		},
		Encode: encodeScript,
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/client.luau:Summary",
				Description: "Formats/client.luau:Description",
			}
		},
	}
}
