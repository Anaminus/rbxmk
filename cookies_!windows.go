//go:build !windows

package rbxmk

import (
	"github.com/anaminus/rbxmk/rtypes"
)

func cookiesFromStudio() rtypes.Cookies {
	return nil
}
