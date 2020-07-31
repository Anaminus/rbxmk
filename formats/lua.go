package formats

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func decodeScript(b []byte, className string) (v types.Value, err error) {
	script := rtypes.NewInstance(className, nil, nil)
	script.Set("Source", types.ProtectedString(b))
	return script, nil
}

func encodeScript(v types.Value) (b []byte, err error) {
	s := rtypes.Stringlike{Value: v}
	if !s.IsStringlike() {
		return nil, cannotEncode(v)
	}
	return []byte(s.Stringlike()), nil
}

func ModuleScriptLua() rbxmk.Format {
	return rbxmk.Format{
		Name: "modulescript.lua",
		Decode: func(b []byte) (v types.Value, err error) {
			return decodeScript(b, "ModuleScript")
		},
		Encode: encodeScript,
	}
}

func ScriptLua() rbxmk.Format {
	return rbxmk.Format{
		Name: "script.lua",
		Decode: func(b []byte) (v types.Value, err error) {
			return decodeScript(b, "Script")
		},
		Encode: encodeScript,
	}
}

func LocalScriptLua() rbxmk.Format {
	return rbxmk.Format{
		Name: "localscript.lua",
		Decode: func(b []byte) (v types.Value, err error) {
			return decodeScript(b, "LocalScript")
		},
		Encode: encodeScript,
	}
}

func Lua() rbxmk.Format {
	return rbxmk.Format{
		Name: "lua",
		Decode: func(b []byte) (v types.Value, err error) {
			return decodeScript(b, "ModuleScript")
		},
		Encode: encodeScript,
	}
}

func ServerLua() rbxmk.Format {
	return rbxmk.Format{
		Name: "server.lua",
		Decode: func(b []byte) (v types.Value, err error) {
			return decodeScript(b, "Script")
		},
		Encode: encodeScript,
	}
}

func ClientLua() rbxmk.Format {
	return rbxmk.Format{
		Name: "client.lua",
		Decode: func(b []byte) (v types.Value, err error) {
			return decodeScript(b, "LocalScript")
		},
		Encode: encodeScript,
	}
}
