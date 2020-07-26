package formats

import (
	"fmt"

	"github.com/anaminus/rbxmk"
)

func All() []func() rbxmk.Format {
	return []func() rbxmk.Format{
		Binary,
		ClientLua,
		LocalScriptLua,
		Lua,
		ModuleScriptLua,
		RBXL,
		RBXLX,
		RBXM,
		RBXMX,
		ScriptLua,
		ServerLua,
		Text,
	}
}

func cannotEncode(v interface{}) error {
	return fmt.Errorf("cannot encode %T", v)
}
