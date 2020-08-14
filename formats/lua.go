package formats

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func decodeScript(b []byte, className string) (v types.Value, err error) {
	script := rtypes.NewInstance(className, nil)
	script.Set("Source", types.ProtectedString(b))
	return script, nil
}

func encodeScript(f rbxmk.FormatOptions, v types.Value) (b []byte, err error) {
	s := rtypes.Stringlike{Value: v}
	if !s.IsStringlike() {
		return nil, cannotEncode(v)
	}
	return []byte(s.Stringlike()), nil
}

func init() { register(ModuleScriptLua) }
func ModuleScriptLua() rbxmk.Format {
	return rbxmk.Format{
		Name: "modulescript.lua",
		Decode: func(f rbxmk.FormatOptions, b []byte) (v types.Value, err error) {
			return decodeScript(b, "ModuleScript")
		},
		Encode: encodeScript,
	}
}

func init() { register(ScriptLua) }
func ScriptLua() rbxmk.Format {
	return rbxmk.Format{
		Name: "script.lua",
		Decode: func(f rbxmk.FormatOptions, b []byte) (v types.Value, err error) {
			return decodeScript(b, "Script")
		},
		Encode: encodeScript,
	}
}

func init() { register(LocalScriptLua) }
func LocalScriptLua() rbxmk.Format {
	return rbxmk.Format{
		Name: "localscript.lua",
		Decode: func(f rbxmk.FormatOptions, b []byte) (v types.Value, err error) {
			return decodeScript(b, "LocalScript")
		},
		Encode: encodeScript,
	}
}

func init() { register(Lua) }
func Lua() rbxmk.Format {
	return rbxmk.Format{
		Name: "lua",
		Decode: func(f rbxmk.FormatOptions, b []byte) (v types.Value, err error) {
			return decodeScript(b, "ModuleScript")
		},
		Encode: encodeScript,
	}
}

func init() { register(ServerLua) }
func ServerLua() rbxmk.Format {
	return rbxmk.Format{
		Name: "server.lua",
		Decode: func(f rbxmk.FormatOptions, b []byte) (v types.Value, err error) {
			return decodeScript(b, "Script")
		},
		Encode: encodeScript,
	}
}

func init() { register(ClientLua) }
func ClientLua() rbxmk.Format {
	return rbxmk.Format{
		Name: "client.lua",
		Decode: func(f rbxmk.FormatOptions, b []byte) (v types.Value, err error) {
			return decodeScript(b, "LocalScript")
		},
		Encode: encodeScript,
	}
}
