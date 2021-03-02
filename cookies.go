package rbxmk

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/textproto"
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

// DecodeCookies parses cookies from r and returns a list of cookies. Cookies
// are parsed as a number of "Set-Cookie" HTTP headers. Returns an empty list if
// the reader is empty.
func DecodeCookies(r io.Reader) (cookies rtypes.Cookies, err error) {
	// There's no direct way to parse cookies, so we have to cheat a little.
	h, err := textproto.NewReader(bufio.NewReader(r)).ReadMIMEHeader()
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("decode cookies: %w", err)
	}
	resp := http.Response{Header: http.Header(h)}
	cs := resp.Cookies()
	cookies = make(rtypes.Cookies, len(cs))
	for i, c := range cs {
		cookies[i] = rtypes.Cookie{Cookie: c}
	}
	return cookies, nil
}

// EncodeCookies formats a list of cookies as a number of "Set-Cookie" HTTP
// headers and writes them to w.
func EncodeCookies(w io.Writer, cookies rtypes.Cookies) (err error) {
	// More cheating.
	h := http.Header{}
	for _, cookie := range cookies {
		h.Add("Set-Cookie", cookie.Cookie.String())
	}
	if err = h.Write(w); err != nil {
		return fmt.Errorf("encode cookies: %w", err)
	}
	return nil
}
