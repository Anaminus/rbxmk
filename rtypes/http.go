package rtypes

import (
	"net/http"

	"github.com/robloxapi/types"
)

// HTTPOptions specifies options to an HTTP request.
type HTTPOptions struct {
	URL            string
	Method         string
	RequestFormat  FormatSelector
	ResponseFormat FormatSelector
	Headers        HTTPHeaders
	Body           types.Value
}

// Type returns a string identifying the type of the value.
func (HTTPOptions) Type() string {
	return "HTTPOptions"
}

// HTTPResponse contains the response to an HTTP request.
type HTTPResponse struct {
	Success       bool
	StatusCode    int
	StatusMessage string
	Headers       HTTPHeaders
	Body          types.Value
}

// Type returns a string identifying the type of the value.
func (HTTPResponse) Type() string {
	return "HTTPResponse"
}

// HTTPHeaders contains the headers of an HTTP request or response.
type HTTPHeaders http.Header

// Type returns a string identifying the type of the value.
func (HTTPHeaders) Type() string {
	return "HTTPHeaders"
}

// RBXAssetOptions specifies options to a Roblox web request.
type RBXAssetOptions struct {
	AssetID int64
	Cookies []string
	Format  FormatSelector
	Body    types.Value
}

// Type returns a string identifying the type of the value.
func (RBXAssetOptions) Type() string {
	return "RBXAssetOptions"
}
