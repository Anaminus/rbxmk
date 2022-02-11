package main

import (
	"os"
	"strings"

	"github.com/anaminus/pflag"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
)

func SetCookieFlags(c *rtypes.Cookies, flags *pflag.FlagSet) {
	flags.Var(funcFlag(func(v string) error {
		cookies, err := rbxmk.CookiesFrom(v)
		if err != nil {
			return err
		}
		*c = append(*c, cookies...)
		return nil
	}), "cookies-from", Doc("Flags/cookies:cookies-from"))
	flags.Var(funcFlag(func(v string) error {
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
	}), "cookies-file", Doc("Flags/cookies:cookies-file"))
	flags.Var(funcFlag(func(v string) error {
		content := os.Getenv(v)
		cookies, err := rbxmk.DecodeCookies(strings.NewReader(content))
		if err != nil {
			return err
		}
		*c = append(*c, cookies...)
		return nil
	}), "cookie-var", Doc("Flags/cookies:cookie-var"))
}
