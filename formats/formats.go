package formats

import (
	"github.com/anaminus/rbxmk"
)

func All() []func() rbxmk.Format {
	return []func() rbxmk.Format{
		RBXL,
		RBXLX,
	}
}
