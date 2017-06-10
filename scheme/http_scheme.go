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

func httpInputSchemeHandler(opt *rbxmk.Options, node *rbxmk.InputNode, _ string) (ext string, src *rbxmk.Source, err error) {
	ext = node.Format
	format, exists := opt.Formats.Init(ext, opt)
	if !exists {
		return "", nil, errors.New("format is not registered")
	}

	resp, err := http.Get(node.Reference[0])
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()
	if !(200 <= resp.StatusCode && resp.StatusCode < 300) {
		return "", nil, errors.New(resp.Status)
	}

	if src, err = format.Decode(resp.Body); err != nil {
		return "", nil, err
	}
	return ext, src, err
}

func httpOutputSchemeHandler(opt *rbxmk.Options, node *rbxmk.OutputNode, _ string) (ext string, src *rbxmk.Source, err error) {
	return node.Format, &rbxmk.Source{}, nil
}

func httpOutputFinalizer(opt *rbxmk.Options, node *rbxmk.OutputNode, _, ext string, outsrc *rbxmk.Source) (err error) {
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
