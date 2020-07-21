package formats

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/types"
)

func decodeScript(b []byte, className string) (v rbxmk.Value, err error) {
	script := types.NewInstance(className)
	script.Set("Source", rbxmk.TValue{Type: "ProtectedString", Value: string(b)})
	return script, nil
}

func encodeScript(v rbxmk.Value) (b []byte, err error) {
	b, ok := types.Stringlike{Value: v}.Stringlike()
	if !ok {
		return nil, cannotEncode(v, false)
	}
	return b, nil
}

func ModuleScriptLua() rbxmk.Format {
	return rbxmk.Format{
		Name: "modulescript.lua",
		Decode: func(b []byte) (v rbxmk.Value, err error) {
			return decodeScript(b, "ModuleScript")
		},
		Encode: encodeScript,
	}
}

func ScriptLua() rbxmk.Format {
	return rbxmk.Format{
		Name: "script.lua",
		Decode: func(b []byte) (v rbxmk.Value, err error) {
			return decodeScript(b, "Script")
		},
		Encode: encodeScript,
	}
}

func LocalScriptLua() rbxmk.Format {
	return rbxmk.Format{
		Name: "localscript.lua",
		Decode: func(b []byte) (v rbxmk.Value, err error) {
			return decodeScript(b, "LocalScript")
		},
		Encode: encodeScript,
	}
}

func Lua() rbxmk.Format {
	return rbxmk.Format{
		Name: "lua",
		Decode: func(b []byte) (v rbxmk.Value, err error) {
			return decodeScript(b, "ModuleScript")
		},
		Encode: encodeScript,
	}
}

func ServerLua() rbxmk.Format {
	return rbxmk.Format{
		Name: "server.lua",
		Decode: func(b []byte) (v rbxmk.Value, err error) {
			return decodeScript(b, "Script")
		},
		Encode: encodeScript,
	}
}

func LocalLua() rbxmk.Format {
	return rbxmk.Format{
		Name: "local.lua",
		Decode: func(b []byte) (v rbxmk.Value, err error) {
			return decodeScript(b, "LocalScript")
		},
		Encode: encodeScript,
	}
}
