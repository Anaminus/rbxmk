package library

import (
	"github.com/anaminus/rbxmk"
)

func All() []func(s rbxmk.State) {
	return []func(s rbxmk.State){
		Base,
		Types,
		RBXMK,
		OS,
		File,
	}
}
