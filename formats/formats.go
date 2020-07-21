package formats

import (
	"fmt"

	"github.com/anaminus/rbxmk"
)

func All() []func() rbxmk.Format {
	return []func() rbxmk.Format{
		RBXL,
		RBXLX,
		RBXM,
		RBXMX,

func cannotEncode(v interface{}, s bool) error {
	if s {
		return fmt.Errorf("cannot encode %s", v)
	}
	return fmt.Errorf("cannot encode %T", v)
}
