//go:build windows

package rbxmk

import (
	"bytes"
	"net/http"

	"github.com/anaminus/rbxmk/rtypes"
	"github.com/danieljoos/wincred"
)

func cookiesFromStudio() rtypes.Cookies {
	const domain = `roblox.com`
	const credBaseName = `https://www.` + domain + `:RobloxStudioAuth`
	const cookieListName = credBaseName + `Cookies`

	cookieList, err := wincred.GetGenericCredential(cookieListName)
	if err != nil {
		return nil
	}
	acceptedCookies := map[string]string{}
	for _, name := range bytes.Split(cookieList.CredentialBlob, []byte{';'}) {
		if len(name) == 0 {
			continue
		}
		acceptedCookies[credBaseName+string(name)] = string(name)
	}

	creds, err := wincred.List()
	if err != nil {
		return nil
	}
	var cookies rtypes.Cookies
	for _, cred := range creds {
		if name, ok := acceptedCookies[cred.TargetName]; ok {
			cookies = append(cookies, rtypes.Cookie{Cookie: &http.Cookie{
				Name:   name,
				Domain: domain,
				Secure: true,
				Value:  string(cred.CredentialBlob),
			}})
		}
	}
	return cookies
}
