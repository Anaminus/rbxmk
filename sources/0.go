package sources

import (
	"github.com/anaminus/rbxmk"
)

func All() []func() rbxmk.Source {
	return []func() rbxmk.Source{
		File,
	}
}
