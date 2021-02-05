package sources

import (
	"encoding/json"
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

const rbxassetidReadURL = "https://assetdelivery.roblox.com/v1/assetId/%d"
const rbxassetidWriteURL = "https://data.roblox.com/Data/Upload.ashx?assetid=%d"

func init() { register(RBXAssetID) }
func RBXAssetID() rbxmk.Source {
	return rbxmk.Source{
		Name: "rbxassetid",
		Library: rbxmk.Library{
			Open: func(s rbxmk.State) *lua.LTable {
				lib := s.L.CreateTable(0, 2)
				lib.RawSetString("read", s.WrapFunc(rbxassetidRead))
				lib.RawSetString("write", s.WrapFunc(rbxassetidWrite))
				return lib
			},
		},
	}
}

func doGeneralRequest(s rbxmk.State, options rtypes.HTTPOptions) (resp *rtypes.HTTPResponse, err error) {
	request, err := doHTTPRequest(s, options)
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

func rbxassetidRead(s rbxmk.State) int {
	options := s.Pull(1, "RBXAssetOptions").(rtypes.RBXAssetOptions)
	if options.Format.Format == "" {
		return s.RaiseError("must specify Format for decoding")
	}
	resp, err := doGeneralRequest(s, rtypes.HTTPOptions{
		URL:            fmt.Sprintf(rbxassetidReadURL, options.AssetID),
		Method:         "GET",
		ResponseFormat: rtypes.FormatSelector{Format: "bin"},
		Headers:        rtypes.HTTPHeaders{}.AppendCookies(options.Cookies),
	})
	if err != nil {
		return s.RaiseError("%w", err)
	}
	var assetResponse struct {
		Location string `json:"location"`
	}
	if err := json.Unmarshal(resp.Body.(types.BinaryString), &assetResponse); err != nil {
		return s.RaiseError("decode asset response: %w", err)
	}
	resp, err = doGeneralRequest(s, rtypes.HTTPOptions{
		URL:            assetResponse.Location,
		Method:         "GET",
		ResponseFormat: options.Format,
		Headers:        rtypes.HTTPHeaders{}.AppendCookies(options.Cookies),
	})
	if err != nil {
		return s.RaiseError("%w", err)
	}
	return s.Push(resp.Body)
}

func rbxassetidWrite(s rbxmk.State) int {
	options := s.Pull(1, "RBXAssetOptions").(rtypes.RBXAssetOptions)
	if options.Format.Format == "" {
		return s.RaiseError("must specify Format for encoding")
	}
	_, err := doGeneralRequest(s, rtypes.HTTPOptions{
		URL:           fmt.Sprintf(rbxassetidWriteURL, options.AssetID),
		Method:        "POST",
		RequestFormat: options.Format,
		Headers:       rtypes.HTTPHeaders{}.AppendCookies(options.Cookies),
		Body:          options.Body,
	})
	if err != nil {
		return s.RaiseError("%w", err)
	}
	return 0
}
