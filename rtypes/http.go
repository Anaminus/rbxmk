package rtypes

import (
	"net/http"

	"github.com/robloxapi/types"
)

type HTTPOptions struct {
	URL            string
	Method         string
	RequestFormat  FormatSelector
	ResponseFormat FormatSelector
	Headers        http.Header
	Body           types.Value
}

// Type returns a string identifying the type of the value.
func (HTTPOptions) Type() string {
	return "HTTPOptions"
}

type HTTPResponse struct {
	Success       bool
	StatusCode    int
	StatusMessage string
	Headers       http.Header
	Body          types.Value
}

// Type returns a string identifying the type of the value.
func (HTTPResponse) Type() string {
	return "HTTPResponse"
}
