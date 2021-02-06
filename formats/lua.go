package formats

import (
	"io"
	"io/ioutil"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func decodeScript(r io.Reader, className string) (v types.Value, err error) {
	script := rtypes.NewInstance(className, nil)
	s, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	script.Set("Source", types.ProtectedString(s))
	return script, nil
}

func canDecodeInstance(f rbxmk.FormatOptions, typeName string) bool {
	return typeName == "Instance"
}

func encodeScript(f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
	s := rtypes.Stringable{Value: v}
	if !s.IsStringable() {
		return cannotEncode(v)
	}
	_, err := w.Write([]byte(s.Stringable()))
	return err
}

func init() { register(ModuleScriptLua) }
func ModuleScriptLua() rbxmk.Format {
	return rbxmk.Format{
		Name:       "modulescript.lua",
		MediaTypes: []string{"application/lua", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
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
		Decode: func(f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
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
		Decode: func(f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
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
		Decode: func(f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
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
		Decode: func(f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
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
		Decode: func(f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			return decodeScript(r, "LocalScript")
		},
		Encode: encodeScript,
	}
}
