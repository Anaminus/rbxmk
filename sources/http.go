package sources

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	lua "github.com/yuin/gopher-lua"
)

func init() { register(HTTP) }
func HTTP() rbxmk.Source {
	return rbxmk.Source{
		Name: "http",
		Read: func(s rbxmk.State) (b []byte, err error) {
			url := string(s.Pull(1, "string").(types.String))
			return httpGet(url)
		},
		Write: func(s rbxmk.State, b []byte) (err error) {
			url := string(s.Pull(1, "string").(types.String))
			return httpPost(url, b)
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

func httpGet(url string) (b []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}
	return ioutil.ReadAll(resp.Body)
}

func httpPost(url string, b []byte) (err error) {
	resp, err := http.Post(url, "", bytes.NewReader(b))
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

	b, err := httpGet(url)
	if err != nil {
		return s.RaiseError(err.Error())
	}
	v, err := format.Decode(b)
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

	b, err := format.Encode(value)
	if err != nil {
		return s.RaiseError(err.Error())
	}
	if err := httpPost(url, b); err != nil {
		return s.RaiseError(err.Error())
	}
	return 0
}
