package scheme

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/anaminus/rbxauth"
	"github.com/anaminus/rbxmk"
)

func init() {
	Schemes.Register(rbxmk.Scheme{
		Name: "rbxassetid",
		Input: &rbxmk.InputScheme{
			Handler: rbxassetidInputSchemeHandler,
		},
		Output: &rbxmk.OutputScheme{
			Handler:   rbxassetidOutputSchemeHandler,
			Finalizer: rbxassetidOutputFinalizer,
		},
	})
}

const wwwSubdomain = "www"
const defaultHost = "roblox.com"
const rbxassetidDownloadPath = "/asset"
const rbxassetidUploadPath = "/ide/publish/uploadexistingasset"

func getHost(opt *rbxmk.Options) (host string) {
	host, _ = opt.Config["Host"].(string)
	if host == "" {
		host = defaultHost
	}
	return
}

func setCookies(opt *rbxmk.Options, req *http.Request, cred rbxauth.Cred, force bool) (err error) {
	users, _ := opt.Config["RobloxAuth"].(map[rbxauth.Cred][]*http.Cookie)
	cookies := users[cred]
	if len(cookies) == 0 && force {
		stream := rbxauth.StandardStream()
		if cred, cookies, err = stream.PromptCred(cred); err != nil {
			return err
		}
		if len(cookies) > 0 {
			if users == nil {
				users = make(map[rbxauth.Cred][]*http.Cookie)
				opt.Config["RobloxAuth"] = users
			}
			users[cred] = cookies
		}
	}
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	return nil
}

func requestWithAuth(opt *rbxmk.Options, req *http.Request, cred rbxauth.Cred) (io.ReadCloser, error) {
	client := &http.Client{}
	token := false
	for force := false; ; {
		req.Header.Del("Cookie")
		if err := setCookies(opt, req, cred, force); err != nil {
			return nil, err
		}

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		switch resp.StatusCode {
		case 409: // Sent by Roblox when user is not authorized to download asset.
			// Retry, forcing the user to login.
			resp.Body.Close()
			if req.GetBody != nil {
				req.Body, _ = req.GetBody()
			}
			if force {
				// Failed with supplied creds, retry with full prompt.
				cred = rbxauth.Cred{}
			}
			force = true
			continue
		case 403: // Sent when token is required.
			if token {
				break
			}
			token = true
			resp.Body.Close()
			req.Header.Add("X-CSRF-TOKEN", resp.Header.Get("X-CSRF-TOKEN"))
			if req.GetBody != nil {
				req.Body, _ = req.GetBody()
			}
			continue
		}

		if !(200 <= resp.StatusCode && resp.StatusCode < 300) {
			resp.Body.Close()
			return nil, errors.New(resp.Status)
		}

		return resp.Body, nil
	}
}

func rbxassetidInputSchemeHandler(opt *rbxmk.Options, node *rbxmk.InputNode, inref []string) (outref []string, data rbxmk.Data, err error) {
	ext := node.Format
	if !opt.Formats.Registered(ext) {
		return nil, nil, errors.New("format is not registered")
	}

	assetURL := url.URL{
		Scheme:   "https",
		Host:     wwwSubdomain + "." + getHost(opt),
		Path:     rbxassetidDownloadPath,
		RawQuery: url.Values{"id": []string{inref[0]}}.Encode(),
	}
	req, _ := http.NewRequest("GET", assetURL.String(), nil)

	resp, err := requestWithAuth(opt, req, node.User)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Close()

	if err := opt.Formats.Decode(ext, opt, nil, resp, &data); err != nil {
		return nil, nil, err
	}
	return inref[1:], data, err
}

func rbxassetidOutputSchemeHandler(opt *rbxmk.Options, node *rbxmk.OutputNode, inref []string) (ext string, outref []string, data rbxmk.Data, err error) {
	ext = node.Format
	if !opt.Formats.Registered(ext) {
		return "", nil, nil, errors.New("format is not registered")
	}

	if len(inref[1:]) == 0 {
		// Avoid downloading if the output is not drilled into.
		return ext, inref[1:], nil, nil
	}

	assetURL := url.URL{
		Scheme:   "https",
		Host:     wwwSubdomain + "." + getHost(opt),
		Path:     rbxassetidDownloadPath,
		RawQuery: url.Values{"id": []string{inref[0]}}.Encode(),
	}
	req, _ := http.NewRequest("GET", assetURL.String(), nil)

	resp, err := requestWithAuth(opt, req, node.User)
	if err != nil {
		return "", nil, nil, err
	}
	defer resp.Close()

	if err := opt.Formats.Decode(ext, opt, nil, resp, &data); err != nil {
		return "", nil, nil, err
	}
	return node.Format, inref[1:], data, err
}

func rbxassetidOutputFinalizer(opt *rbxmk.Options, node *rbxmk.OutputNode, inref []string, ext string, outdata rbxmk.Data) (err error) {
	if !opt.Formats.Registered(ext) {
		return errors.New("format is not registered")
	}
	var buf strings.Builder
	if err = opt.Formats.Encode(ext, opt, nil, &buf, outdata); err != nil {
		return err
	}

	uploadURL := url.URL{
		Scheme:   "https",
		Host:     wwwSubdomain + "." + getHost(opt),
		Path:     rbxassetidUploadPath,
		RawQuery: url.Values{"assetID": []string{inref[0]}}.Encode(),
	}
	req, _ := http.NewRequest("POST", uploadURL.String(), strings.NewReader(buf.String()))
	req.GetBody = func() (io.ReadCloser, error) {
		return ioutil.NopCloser(strings.NewReader(buf.String())), nil
	}

	resp, err := requestWithAuth(opt, req, node.User)
	if err != nil {
		return err
	}
	defer resp.Close()

	return nil
}
