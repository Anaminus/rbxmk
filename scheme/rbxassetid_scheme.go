package scheme

import (
	"bytes"
	"errors"
	"github.com/anaminus/rbxauth"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/config"
	"net/http"
	"net/url"
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
const rbxassetidDownloadPath = "/asset"
const rbxassetidUploadPath = "/ide/publish/uploadexistingasset"

func setCookies(req *http.Request, opt rbxmk.Options, cred rbxauth.Cred) (err error) {
	users := config.RobloxAuth(opt)
	cookies := users[cred]
	if len(cookies) == 0 {
		auth := &rbxauth.Config{Host: config.Host(opt)}
		if cred, cookies, err = auth.PromptCred(cred); err != nil {
			return err
		}
		if len(cookies) > 0 {
			users[cred] = cookies
		}
	}
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	return nil
}

func rbxassetidInputSchemeHandler(opt rbxmk.Options, node *rbxmk.InputNode, inref []string) (outref []string, data rbxmk.Data, err error) {
	ext := node.Format
	if !opt.Formats.Registered(ext) {
		return nil, nil, errors.New("format is not registered")
	}

	assetURL := url.URL{
		Scheme:   "https",
		Host:     wwwSubdomain + "." + config.Host(opt),
		Path:     rbxassetidDownloadPath,
		RawQuery: url.Values{"id": []string{node.Reference[0]}}.Encode(),
	}
	req, _ := http.NewRequest("GET", assetURL.String(), nil)
	if err := setCookies(req, opt, node.User); err != nil {
		return nil, nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	if !(200 <= resp.StatusCode && resp.StatusCode < 300) {
		return nil, nil, errors.New(resp.Status)
	}

	if err := opt.Formats.Decode(ext, opt, nil, resp.Body, &data); err != nil {
		return nil, nil, err
	}
	return inref[1:], data, err
}

func rbxassetidOutputSchemeHandler(opt rbxmk.Options, node *rbxmk.OutputNode, inref []string) (ext string, outref []string, data rbxmk.Data, err error) {
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
		Host:     wwwSubdomain + "." + config.Host(opt),
		Path:     rbxassetidDownloadPath,
		RawQuery: url.Values{"id": []string{node.Reference[0]}}.Encode(),
	}
	req, _ := http.NewRequest("GET", assetURL.String(), nil)
	if err := setCookies(req, opt, node.User); err != nil {
		return "", nil, nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", nil, nil, err
	}
	defer resp.Body.Close()
	if !(200 <= resp.StatusCode && resp.StatusCode < 300) {
		return "", nil, nil, errors.New(resp.Status)
	}

	if err := opt.Formats.Decode(ext, opt, nil, resp.Body, &data); err != nil {
		return "", nil, nil, err
	}
	return node.Format, inref[1:], data, err
}

func rbxassetidOutputFinalizer(opt rbxmk.Options, node *rbxmk.OutputNode, inref []string, ext string, outdata rbxmk.Data) (err error) {
	if !opt.Formats.Registered(ext) {
		return errors.New("format is not registered")
	}
	var buf bytes.Buffer
	if err = opt.Formats.Encode(ext, opt, nil, &buf, outdata); err != nil {
		return err
	}

	uploadURL := url.URL{
		Scheme:   "https",
		Host:     wwwSubdomain + "." + config.Host(opt),
		Path:     rbxassetidUploadPath,
		RawQuery: url.Values{"assetID": []string{node.Reference[0]}}.Encode(),
	}
	req, _ := http.NewRequest("GET", uploadURL.String(), &buf)
	if err := setCookies(req, opt, node.User); err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if !(200 <= resp.StatusCode && resp.StatusCode < 300) {
		return errors.New(resp.Status)
	}
	return nil
}
