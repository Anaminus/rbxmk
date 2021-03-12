package main

import (
	"os"
	"strings"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/anaminus/snek"
)

func SetCookieFlags(c *rtypes.Cookies, flags snek.FlagSet) {
	usage := `Append cookies from a known location. See the documentation of
rbxmk.cookiesFrom for a list of locations. Can be given any number of times.`
	flags.Func("cookies-from", usage, func(v string) error {
		cookies, err := rbxmk.CookiesFrom(v)
		if err != nil {
			return err
		}
		*c = append(*c, cookies...)
		return nil
	})

	usage = `Append cookies from a file. The file is formatted as a number of
Set-Cookie headers. Can be given any number of times.`
	flags.Func("cookies-file", usage, func(v string) error {
		f, err := os.Open(v)
		if err != nil {
			return err
		}
		defer f.Close()
		cookies, err := rbxmk.DecodeCookies(f)
		if err != nil {
			return err
		}
		*c = append(*c, cookies...)
		return nil
	})

	usage = `Append a cookie from an environment variable. The content is
formatted as a number of Set-Cookie headers. Can be given any number of times.`
	flags.Func("cookie-var", usage, func(v string) error {
		content := os.Getenv(v)
		cookies, err := rbxmk.DecodeCookies(strings.NewReader(content))
		if err != nil {
			return err
		}
		*c = append(*c, cookies...)
		return nil
	})
}
