package rtypes

import (
	"net/http"

	"github.com/robloxapi/types"
)

const T_HTTPRequest = "HTTPRequest"

const T_HTTPOptions = "HTTPOptions"

// HTTPOptions specifies options to an HTTP request.
type HTTPOptions struct {
	URL            string
	Method         string
	RequestFormat  FormatSelector
	ResponseFormat FormatSelector
	Headers        HTTPHeaders
	Cookies        Cookies
	Body           types.Value
}

// Type returns a string identifying the type of the value.
func (HTTPOptions) Type() string {
	return T_HTTPOptions
}

const T_HTTPResponse = "HTTPResponse"

// HTTPResponse contains the response to an HTTP request.
type HTTPResponse struct {
	Success       bool
	StatusCode    int
	StatusMessage string
	Headers       HTTPHeaders
	Cookies       Cookies
	Body          types.Value
}

// Type returns a string identifying the type of the value.
func (HTTPResponse) Type() string {
	return T_HTTPResponse
}

const T_HTTPHeaders = "HTTPHeaders"

// HTTPHeaders contains the headers of an HTTP request or response.
type HTTPHeaders http.Header

// AppendCookie formats and adds the given cookies to the Cookie header.
func (h HTTPHeaders) AppendCookies(c Cookies) HTTPHeaders {
	req := http.Request{Header: http.Header(h)}
	for _, cookie := range c {
		req.AddCookie(cookie.Cookie)
	}
	return h
}

// AppendSetCookie formats and adds the given cookies to the Set-Cookie header.
func (h HTTPHeaders) AppendSetCookies(c Cookies) HTTPHeaders {
	for _, cookie := range c {
		if v := cookie.Cookie.String(); v != "" {
			http.Header(h).Add("Set-Cookie", v)
		}
	}
	return h
}

// RetrieveCookies parses the Cookie header.
func (h HTTPHeaders) RetrieveCookies() Cookies {
	return HTTPHeaders{"Set-Cookie": h["Cookie"]}.RetrieveSetCookies()
}

// RetrieveSetCookies parses the Set-Cookie header.
func (h HTTPHeaders) RetrieveSetCookies() Cookies {
	cs := (&http.Response{Header: http.Header(h)}).Cookies()
	cookies := make(Cookies, len(cs))
	for i, c := range cs {
		cookies[i] = Cookie{c}
	}
	return cookies
}

// Type returns a string identifying the type of the value.
func (HTTPHeaders) Type() string {
	return T_HTTPHeaders
}

const T_RBXAssetOptions = "RBXAssetOptions"

// RBXAssetOptions specifies options to a Roblox web request.
type RBXAssetOptions struct {
	AssetID int64
	Cookies Cookies
	Format  FormatSelector
	Body    types.Value
}

// Type returns a string identifying the type of the value.
func (RBXAssetOptions) Type() string {
	return T_RBXAssetOptions
}
