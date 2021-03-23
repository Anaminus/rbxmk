package main

import (
	"os"
	"strings"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/anaminus/snek"
)

func SetCookieFlags(c *rtypes.Cookies, flags snek.FlagSet) {
	flags.Func("cookies-from", Doc("commands/cookie_flags.md/cookies-from"), func(v string) error {
		cookies, err := rbxmk.CookiesFrom(v)
		if err != nil {
			return err
		}
		*c = append(*c, cookies...)
		return nil
	})
	flags.Func("cookies-file", Doc("commands/cookie_flags.md/cookies-file"), func(v string) error {
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
	flags.Func("cookie-var", Doc("commands/cookie_flags.md/cookie-var"), func(v string) error {
		content := os.Getenv(v)
		cookies, err := rbxmk.DecodeCookies(strings.NewReader(content))
		if err != nil {
			return err
		}
		*c = append(*c, cookies...)
		return nil
	})
}
