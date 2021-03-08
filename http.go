package rbxmk

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/anaminus/rbxmk/rtypes"
)

// HTTPRequest performs and HTTP request with a promise-like API.
type HTTPRequest struct {
	global Global

	cancel context.CancelFunc

	respch chan *http.Response
	resp   *rtypes.HTTPResponse

	errch chan error
	err   error

	fmt Format
	sel rtypes.FormatSelector
}

// Type returns a string identifying the type of the value.
func (*HTTPRequest) Type() string {
	return "HTTPRequest"
}

// do concurrently begins the request.
func (r *HTTPRequest) do(client *Client, req *http.Request) {
	defer close(r.respch)
	defer close(r.errch)
	resp, err := client.Do(req)
	if err != nil {
		r.errch <- err
		return
	}
	r.respch <- resp
}

// Resolve blocks until the request resolves.
func (r *HTTPRequest) Resolve() (*rtypes.HTTPResponse, error) {
	if r.resp != nil || r.err != nil {
		return r.resp, r.err
	}
	select {
	case resp := <-r.respch:
		defer resp.Body.Close()
		headers := rtypes.HTTPHeaders(resp.Header)
		r.resp = &rtypes.HTTPResponse{
			Success:       200 <= resp.StatusCode && resp.StatusCode < 300,
			StatusCode:    resp.StatusCode,
			StatusMessage: resp.Status,
			Headers:       headers,
			Cookies:       headers.RetrieveSetCookies(),
		}
		if r.fmt.Name != "" {
			if r.resp.Body, r.err = r.fmt.Decode(r.global, r.sel, resp.Body); r.err != nil {
				return nil, r.err
			}
		}
		return r.resp, nil
	case r.err = <-r.errch:
		return nil, r.err
	}
}

// Cancel cancels the request.
func (r *HTTPRequest) Cancel() {
	if r.resp != nil || r.err != nil {
		return
	}
	r.cancel()
	defer close(r.respch)
	defer close(r.errch)
	r.err = <-r.errch
}

// BeginHTTPRequest begins an HTTP request according to the given options, in
// the context of the given world.
//
// The request starts immediately, and can either be resolved or canceled.
func BeginHTTPRequest(w *World, options rtypes.HTTPOptions) (request *HTTPRequest, err error) {
	var buf *bytes.Buffer
	if options.RequestFormat.Format != "" {
		reqfmt := w.Format(options.RequestFormat.Format)
		if reqfmt.Encode == nil {
			return nil, fmt.Errorf("cannot encode with format %s", reqfmt.Name)
		}
		if options.Body != nil {
			buf = new(bytes.Buffer)
			if err := reqfmt.Encode(w.Global, options.RequestFormat, buf, options.Body); err != nil {
				return nil, fmt.Errorf("encode body: %w", err)
			}
		}
	}
	var respfmt Format
	if options.ResponseFormat.Format != "" {
		respfmt = w.Format(options.ResponseFormat.Format)
		if respfmt.Decode == nil {
			return nil, fmt.Errorf("cannot decode with format %s", respfmt.Name)
		}
	}

	// Create request.
	ctx, cancel := context.WithCancel(context.TODO())
	var req *http.Request
	if buf != nil {
		// Use of *bytes.Buffer guarantees that req.GetBody will be set.
		req, err = http.NewRequestWithContext(ctx, options.Method, options.URL, buf)
	} else {
		req, err = http.NewRequestWithContext(ctx, options.Method, options.URL, nil)
	}
	if err != nil {
		cancel()
		return nil, err
	}
	if options.Headers == nil {
		options.Headers = rtypes.HTTPHeaders{}
	}
	req.Header = http.Header(options.Headers.AppendCookies(options.Cookies))

	// Push request object.
	request = &HTTPRequest{
		global: w.Global,
		cancel: cancel,
		respch: make(chan *http.Response),
		errch:  make(chan error),
		fmt:    respfmt,
		sel:    options.ResponseFormat,
	}
	go request.do(w.Client, req)
	return request, nil
}

// DoHTTPRequest begins and resolves an HTTPRequest. Returns an error if the
// reponse did not return a successful status.
func DoHTTPRequest(w *World, options rtypes.HTTPOptions) (resp *rtypes.HTTPResponse, err error) {
	request, err := BeginHTTPRequest(w, options)
	if err != nil {
		return nil, err
	}
	if resp, err = request.Resolve(); err != nil {
		return nil, err
	}
	if !resp.Success {
		return nil, fmt.Errorf("%s", resp.StatusMessage)
	}
	return resp, nil
}
