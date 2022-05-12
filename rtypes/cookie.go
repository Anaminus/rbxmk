package rtypes

import "net/http"

const T_Cookie = "Cookie"

type Cookie struct {
	*http.Cookie
}

// Type returns a string identifying the type of the value.
func (Cookie) Type() string {
	return T_Cookie
}

// String returns a string representation of the value.
func (c Cookie) String() string {
	return "Cookie: " + c.Name
}

const T_Cookies = "Cookies"

type Cookies []Cookie

// Type returns a string identifying the type of the value.
func (Cookies) Type() string {
	return T_Cookies
}
