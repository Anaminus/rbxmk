// +build windows

package rbxmk

import (
	"net/http"
	"strings"
	"time"

	"github.com/anaminus/parse"
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

func parseRegistryCookie(cookie *http.Cookie, v string) bool {
	r := parse.NewTextReader(strings.NewReader(v))
	if r.Is("SEC::<YES>") {
		cookie.Secure = true
		r.Is(",")
	}
	if r.Is("EXP::<") {
		value, ok := r.Until('>')
		if !ok {
			return false
		}
		t, err := time.Parse(time.RFC3339, value)
		if err != nil {
			return false
		}
		cookie.Expires = t
		r.Is(",")
	}
	if r.Is("COOK::<") {
		value, ok := r.UntilEOF()
		if !ok {
			return false
		}
		cookie.Value = value[:len(value)-1]
	}
	return r.IsEOF()
}
