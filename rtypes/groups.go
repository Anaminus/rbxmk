package rtypes

import (
	"github.com/anaminus/rbxmk"
)

type Array []rbxmk.Value

func (Array) Type() string { return "Array" }

type Dictionary map[string]rbxmk.Value

func (Dictionary) Type() string { return "Dictionary" }

type Tuple []rbxmk.Value

func (Tuple) Type() string { return "Tuple" }

type Objects []*Instance

func (Objects) Type() string { return "Objects" }
