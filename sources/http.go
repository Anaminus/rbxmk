package sources

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
)

func init() { register(HTTP) }
func HTTP() rbxmk.Source {
	return rbxmk.Source{
		Name: "http",
		Read: func(s rbxmk.State) (b []byte, err error) {
			url := string(s.Pull(1, "string").(types.String))
			r, err := httpGet(url)
			if err != nil {
				return nil, err
			}
			defer r.Close()
			return ioutil.ReadAll(r)
		},
		Write: func(s rbxmk.State, b []byte) (err error) {
			url := string(s.Pull(1, "string").(types.String))
			return httpPost(url, bytes.NewReader(b))
		},
		Library: rbxmk.Library{
			Open: func(s rbxmk.State) *lua.LTable {
				lib := s.L.CreateTable(0, 2)
				lib.RawSetString("read", s.WrapFunc(httpRead))
				lib.RawSetString("write", s.WrapFunc(httpWrite))
				return lib
			},
		},
	}
}

func httpGet(url string) (r io.ReadCloser, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}
	return resp.Body, nil
}

func httpPost(url string, r io.Reader) (err error) {
	resp, err := http.Post(url, "", r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("bad status: %s", resp.Status)
	}
	return nil
}

func httpRead(s rbxmk.State) int {
	url := string(s.Pull(1, "string").(types.String))
	formatName := string(s.PullOpt(2, "string", types.String("")).(types.String))
	format := s.Format(formatName)
	if format.Name == "" {
		return s.RaiseError("unknown format %q", formatName)
	}
	if format.Decode == nil {
		return s.RaiseError("cannot decode with format %s", format.Name)
	}

	r, err := httpGet(url)
	if err != nil {
		return s.RaiseError(err.Error())
	}
	defer r.Close()
	v, err := format.Decode(rbxmk.FormatOptions{}, r)
	if err != nil {
		return s.RaiseError(err.Error())
	}
	return s.Push(v)
}

func httpWrite(s rbxmk.State) int {
	url := string(s.Pull(1, "string").(types.String))
	formatName := string(s.Pull(2, "string").(types.String))
	value := s.Pull(3, "Variant")
	format := s.Format(formatName)
	if format.Name == "" {
		return s.RaiseError("unknown format %q", formatName)
	}
	if format.Encode == nil {
		return s.RaiseError("cannot encode with format %s", format.Name)
	}

	var w bytes.Buffer
	if err := format.Encode(rbxmk.FormatOptions{}, &w, value); err != nil {
		return s.RaiseError(err.Error())
	}
	if err := httpPost(url, &w); err != nil {
		return s.RaiseError(err.Error())
	}
	return 0
}
