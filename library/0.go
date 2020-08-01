package library

import (
	"github.com/anaminus/rbxmk"
)

func All() []rbxmk.Library {
	return []rbxmk.Library{
		Base,
		Types,
		Sources,
		RBXMK,
		OS,
	}
}
