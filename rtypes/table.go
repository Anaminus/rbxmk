package rtypes

import (
	"github.com/yuin/gopher-lua"
)

type Table struct {
	*lua.LTable
}

func (Table) Type() string     { return "table" }
func (t Table) String() string { return t.LTable.String() }
