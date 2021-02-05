package rtypes

import "net/http"

type Cookie struct {
	*http.Cookie
}

// Type returns a string identifying the type of the value.
func (Cookie) Type() string {
	return "Cookie"
}

// String returns a string representation of the value.
func (c Cookie) String() string {
	return "Cookie: " + c.Name
}

type Cookies []Cookie

// Type returns a string identifying the type of the value.
func (Cookies) Type() string {
	return "Cookies"
}
