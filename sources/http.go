package sources

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(HTTP) }
func HTTP() rbxmk.Source {
	return rbxmk.Source{
		Name: "http",
		Read: func(s rbxmk.State) (b []byte, err error) {
			options := s.Pull(1, "HTTPOptions").(rtypes.HTTPOptions)
			options.Method = "GET"
			options.RequestFormat = rtypes.FormatSelector{}
			options.ResponseFormat = rtypes.FormatSelector{Format: "bin"}
			options.Body = nil
			request, err := doHTTPRequest(s, options)
			if err != nil {
				return nil, err
			}
			resp, err := request.Resolve()
			if err != nil {
				return nil, err
			}
			if !resp.Success {
				return nil, fmt.Errorf(resp.StatusMessage)
			}
			body := resp.Body.(types.Stringlike).Stringlike()
			return []byte(body), nil
		},
		Write: func(s rbxmk.State, b []byte) (err error) {
			options := s.Pull(1, "HTTPOptions").(rtypes.HTTPOptions)
			options.Method = "POST"
			options.RequestFormat = rtypes.FormatSelector{Format: "bin"}
			options.ResponseFormat = rtypes.FormatSelector{}
			options.Body = types.BinaryString(b)
			request, err := doHTTPRequest(s, options)
			if err != nil {
				return err
			}
			resp, err := request.Resolve()
			if err != nil {
				return err
			}
			if !resp.Success {
				return fmt.Errorf(resp.StatusMessage)
			}
			return nil
		},
		Library: rbxmk.Library{
			Open: func(s rbxmk.State) *lua.LTable {
				lib := s.L.CreateTable(0, 1)
				lib.RawSetString("request", s.WrapFunc(httpRequest))
				return lib
			},
		},
	}
}

type HTTPRequest struct {
	cancel context.CancelFunc

	respch chan *http.Response
	resp   *rtypes.HTTPResponse

	errch chan error
	err   error

	fmt rbxmk.Format
	sel rtypes.FormatSelector
}

// Type returns a string identifying the type of the value.
func (*HTTPRequest) Type() string {
	return "HTTPRequest"
}

func (r *HTTPRequest) do(client *http.Client, req *http.Request) {
	defer close(r.respch)
	defer close(r.errch)
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		r.errch <- err
		return
	}
	r.respch <- resp
}

func (r *HTTPRequest) Resolve() (*rtypes.HTTPResponse, error) {
	if r.resp != nil || r.err != nil {
		return r.resp, r.err
	}
	select {
	case resp := <-r.respch:
		defer resp.Body.Close()
		r.resp = &rtypes.HTTPResponse{
			Success:       200 <= resp.StatusCode && resp.StatusCode < 300,
			StatusCode:    resp.StatusCode,
			StatusMessage: resp.Status,
			Headers:       rtypes.HTTPHeaders(resp.Header),
		}
		if r.fmt.Name != "" {
			if r.resp.Body, r.err = r.fmt.Decode(r.sel, resp.Body); r.err != nil {
				return nil, r.err
			}
		}
		return r.resp, nil
	case r.err = <-r.errch:
		return nil, r.err
	}
}

func (r *HTTPRequest) Cancel() {
	if r.resp != nil || r.err != nil {
		return
	}
	r.cancel()
	defer close(r.respch)
	defer close(r.errch)
	r.err = <-r.errch
}

func doHTTPRequest(s rbxmk.State, options rtypes.HTTPOptions) (request *HTTPRequest, err error) {
	var buf *bytes.Buffer
	if options.RequestFormat.Format != "" {
		reqfmt := s.Format(options.RequestFormat.Format)
		if reqfmt.Encode == nil {
			return nil, fmt.Errorf("cannot encode with format %s", reqfmt.Name)
		}
		if options.Body != nil {
			buf = new(bytes.Buffer)
			if err := reqfmt.Encode(options.RequestFormat, buf, options.Body); err != nil {
				return nil, fmt.Errorf("encode body: %w", err)
			}
		}
	}
	var respfmt rbxmk.Format
	if options.ResponseFormat.Format != "" {
		respfmt = s.Format(options.ResponseFormat.Format)
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
	req.Header = http.Header(options.Headers)

	// Push request object.
	request = &HTTPRequest{
		cancel: cancel,
		respch: make(chan *http.Response),
		errch:  make(chan error),
		fmt:    respfmt,
		sel:    options.ResponseFormat,
	}
	go request.do(s.Client, req)
	return request, nil
}

func httpRequest(s rbxmk.State) int {
	options := s.Pull(1, "HTTPOptions").(rtypes.HTTPOptions)
	request, err := doHTTPRequest(s, options)
	if err != nil {
		return s.RaiseError("%w", err)
	}
	return s.Push(request)
}
