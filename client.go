package rbxmk

import (
	"fmt"
	"net/http"
	"sync"
)

// UserAgent is the User-Agent header string sent with HTTP requests made by
// rbxmk. It includes components that ensure the client will operate with Roblox
// website APIs.
const UserAgent = "RobloxStudio/WinInet rbxmk/0.0"

// Client wraps an http.Client to handle various additional behavior.
type Client struct {
	*http.Client

	mtx sync.Mutex
	// Maps host name to token value.
	csrfTokens map[string]string
}

// NewClient returns an initialized Client. If *client* is nil, then
// http.DefaultClient is used.
func NewClient(client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{
		Client:     client,
		csrfTokens: make(map[string]string, 1),
	}
}

// Do sends a request, with the following additional behaviors:
//
//     - Includes a configured user agent header with the request, if the header
//       is unset.
//     - Handles CSRF token validation.
func (c *Client) Do(req *http.Request) (resp *http.Response, err error) {
	for i := 2; i > 0; i-- {
		// Merge headers.
		c.mtx.Lock()
		if req.Header.Get("User-Agent") == "" {
			req.Header.Set("User-Agent", UserAgent)
		}
		if token, ok := c.csrfTokens[req.URL.Host]; ok {
			req.Header.Set("X-Csrf-Token", token)
		}
		c.mtx.Unlock()

		// Do request.
		if resp, err = c.Client.Do(req); err != nil {
			return resp, err
		}
		switch resp.StatusCode {
		// Check for failed CSRF token.
		case 403:
			token := resp.Header.Get("X-Csrf-Token")
			if token == "" {
				// No token; regular 403.
				return resp, err
			}
			c.mtx.Lock()
			c.csrfTokens[req.URL.Host] = token
			c.mtx.Unlock()
			// Reset body if needed.
			if req.Body != nil {
				if req.GetBody == nil {
					return nil, fmt.Errorf("retry failed: cannot reset body")
				}
				if req.Body, err = req.GetBody(); err != nil {
					return nil, fmt.Errorf("retry failed: reset body: %w", err)
				}
			}
			// Retry with new token.
		default:
			return resp, err
		}
	}
	return resp, err
}
