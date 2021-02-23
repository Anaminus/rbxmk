package formats

import (
	"io"

	"github.com/anaminus/rbxmk"
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
	}
}
