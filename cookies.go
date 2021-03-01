package rbxmk

import (
	"fmt"
	"strings"

	"github.com/anaminus/rbxmk/rtypes"
)

// CookiesFrom retrieves cookies from a known location. location is
// case-insensitive. The following locations are implemented:
//
//     - studio: Returns the cookies used for authentication when logging into
//       Roblox Studio.
func CookiesFrom(location string) (cookies rtypes.Cookies, err error) {
	switch strings.ToLower(location) {
	case "studio":
		cookies = cookiesFromStudio()
		return cookies, nil
	default:
		return nil, fmt.Errorf("unknown location %q", location)
	}
}
