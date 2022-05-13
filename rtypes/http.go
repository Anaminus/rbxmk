package rtypes

import (
	"net/http"

	"github.com/robloxapi/types"
)

const T_HttpRequest = "HttpRequest"

const T_HttpOptions = "HttpOptions"

// HttpOptions specifies options to an HTTP request.
type HttpOptions struct {
	URL            string
	Method         string
	RequestFormat  FormatSelector
	ResponseFormat FormatSelector
	Headers        HttpHeaders
	Cookies        Cookies
	Body           types.Value
}

// Type returns a string identifying the type of the value.
func (HttpOptions) Type() string {
	return T_HttpOptions
}

const T_HttpResponse = "HttpResponse"

// HttpResponse contains the response to an HTTP request.
type HttpResponse struct {
	Success       bool
	StatusCode    int
	StatusMessage string
	Headers       HttpHeaders
	Cookies       Cookies
	Body          types.Value
}

// Type returns a string identifying the type of the value.
func (HttpResponse) Type() string {
	return T_HttpResponse
}

const T_HttpHeaders = "HttpHeaders"

// HttpHeaders contains the headers of an HTTP request or response.
type HttpHeaders http.Header

// AppendCookie formats and adds the given cookies to the Cookie header.
func (h HttpHeaders) AppendCookies(c Cookies) HttpHeaders {
	req := http.Request{Header: http.Header(h)}
	for _, cookie := range c {
		req.AddCookie(cookie.Cookie)
	}
	return h
}

// AppendSetCookie formats and adds the given cookies to the Set-Cookie header.
func (h HttpHeaders) AppendSetCookies(c Cookies) HttpHeaders {
	for _, cookie := range c {
		if v := cookie.Cookie.String(); v != "" {
			http.Header(h).Add("Set-Cookie", v)
		}
	}
	return h
}

// RetrieveCookies parses the Cookie header.
func (h HttpHeaders) RetrieveCookies() Cookies {
	return HttpHeaders{"Set-Cookie": h["Cookie"]}.RetrieveSetCookies()
}

// RetrieveSetCookies parses the Set-Cookie header.
func (h HttpHeaders) RetrieveSetCookies() Cookies {
	cs := (&http.Response{Header: http.Header(h)}).Cookies()
	cookies := make(Cookies, len(cs))
	for i, c := range cs {
		cookies[i] = Cookie{c}
	}
	return cookies
}

// Type returns a string identifying the type of the value.
func (HttpHeaders) Type() string {
	return T_HttpHeaders
}

const T_RbxAssetOptions = "RbxAssetOptions"

// RbxAssetOptions specifies options to a Roblox web request.
type RbxAssetOptions struct {
	AssetId int64
	Cookies Cookies
	Format  FormatSelector
	Body    types.Value
}

// Type returns a string identifying the type of the value.
func (RbxAssetOptions) Type() string {
	return T_RbxAssetOptions
}
