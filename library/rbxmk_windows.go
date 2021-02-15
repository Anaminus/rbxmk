// +build windows

package library

import (
	"net/http"

	"github.com/anaminus/rbxmk/rtypes"
	reg "golang.org/x/sys/windows/registry"
)

func cookiesFromStudio() rtypes.Cookies {
	const keyPath = `Software\Roblox\RobloxStudioBrowser\roblox.com`
	key, err := reg.OpenKey(reg.CURRENT_USER, keyPath, reg.QUERY_VALUE)
	if err != nil {
		return nil
	}
	defer key.Close()
	v, _, err := key.GetStringValue(".ROBLOSECURITY")
	if err != nil {
		return nil
	}
	cookie := &http.Cookie{
		Name:   ".ROBLOSECURITY",
		Domain: "roblox.com",
	}
	if !parseRegistryCookie(cookie, v) {
		return nil
	}
	return rtypes.Cookies{rtypes.Cookie{Cookie: cookie}}
}
