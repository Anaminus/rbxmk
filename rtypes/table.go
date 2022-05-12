package rtypes

import (
	lua "github.com/anaminus/gopher-lua"
)

const T_Table = "table"

// Table wraps a Lua table to implement types.Value.
type Table struct {
	*lua.LTable
}

// Type returns a string identifying the type of the value.
func (Table) Type() string {
	return T_Table
}

// String returns a string representation of the value.
func (t Table) String() string {
	return t.LTable.String()
}
