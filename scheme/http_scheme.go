package scheme

import (
	"bytes"
	"errors"
	"github.com/anaminus/rbxmk"
	"net/http"
)

func init() {
	input := rbxmk.InputScheme{
		Handler: httpInputSchemeHandler,
	}
	output := rbxmk.OutputScheme{
		Handler:   httpOutputSchemeHandler,
		Finalizer: httpOutputFinalizer,
	}

	registerInput("http", input)
	registerOutput("http", output)

	registerInput("https", input)
	registerOutput("https", output)
}

func httpInputSchemeHandler(opt *rbxmk.Options, node *rbxmk.InputNode, inref []string) (ext string, outref []string, src *rbxmk.Source, err error) {
	ext = node.Format
	format, exists := opt.Formats.Init(ext, opt)
	if !exists {
		return "", nil, nil, errors.New("format is not registered")
	}

	resp, err := http.Get(node.Reference[0])
	if err != nil {
		return "", nil, nil, err
	}
	defer resp.Body.Close()
	if !(200 <= resp.StatusCode && resp.StatusCode < 300) {
		return "", nil, nil, errors.New(resp.Status)
	}

	if src, err = format.Decode(resp.Body); err != nil {
		return "", nil, nil, err
	}
	return ext, inref[1:], src, err
}

func httpOutputSchemeHandler(opt *rbxmk.Options, node *rbxmk.OutputNode, inref []string) (ext string, outref []string, src *rbxmk.Source, err error) {
	return node.Format, inref[1:], &rbxmk.Source{}, nil
}

func httpOutputFinalizer(opt *rbxmk.Options, node *rbxmk.OutputNode, ext string, inref []string, outsrc *rbxmk.Source) (err error) {
	format, exists := opt.Formats.Init(ext, opt)
	if !exists {
		return errors.New("format is not registered")
	}

	if !format.CanEncode(outsrc) {
		return errors.New("cannot encode transformed output")
	}

	var buf bytes.Buffer
	if err := format.Encode(&buf, outsrc); err != nil {
		return err
	}
	resp, err := http.Post(node.Reference[0], "", &buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if !(200 <= resp.StatusCode && resp.StatusCode < 300) {
		return errors.New(resp.Status)
	}
	return nil
}
