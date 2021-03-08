package library

import (
	"encoding/json"
	"fmt"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	reflect "github.com/anaminus/rbxmk/library/rbxassetid"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

const rbxassetidReadURL = "https://assetdelivery.roblox.com/v1/assetId/%d"
const rbxassetidWriteURL = "https://data.roblox.com/Data/Upload.ashx?assetid=%d"

func init() { register(RBXAssetIDSource, 10) }

var RBXAssetIDSource = rbxmk.Library{
	Name: "rbxassetid",
	Open: func(s rbxmk.State) *lua.LTable {
		lib := s.L.CreateTable(0, 2)
		lib.RawSetString("read", s.WrapFunc(func(s rbxmk.State) int {
			options := s.Pull(1, "RBXAssetOptions").(rtypes.RBXAssetOptions)
			body, err := RBXAssetID{World: s.World}.Read(options)
			if err != nil {
				return s.RaiseError("%s", err)
			}
			return s.Push(body)
		}))
		lib.RawSetString("write", s.WrapFunc(func(s rbxmk.State) int {
			options := s.Pull(1, "RBXAssetOptions").(rtypes.RBXAssetOptions)
			err := RBXAssetID{World: s.World}.Write(options)
			if err != nil {
				return s.RaiseError("%s", err)
			}
			return 0
		}))

		for _, f := range reflect.All() {
			r := f()
			s.RegisterReflector(r)
			s.ApplyReflector(r, lib)
		}

		return lib
	},
	Dump: func(s rbxmk.State) dump.Library {
		return dump.Library{
			Struct: dump.Struct{
				Fields: dump.Fields{
					"read": dump.Function{
						Parameters: dump.Parameters{
							{Name: "options", Type: dt.Prim("RBXAssetOptions")},
						},
						Returns: dump.Parameters{
							{Name: "value", Type: dt.Prim("any")},
						},
					},
					"write": dump.Function{
						Parameters: dump.Parameters{
							{Name: "options", Type: dt.Prim("RBXAssetOptions")},
						},
					},
				},
			},
		}
	},
}

// RBXAssetID provides access to assets on the Roblox website.
type RBXAssetID struct {
	*rbxmk.World
}

// Read downloads an asset according to the given options.
func (s RBXAssetID) Read(options rtypes.RBXAssetOptions) (body types.Value, err error) {
	if options.Format.Format == "" {
		return nil, fmt.Errorf("must specify Format for decoding")
	}
	resp, err := rbxmk.DoHTTPRequest(s.World, rtypes.HTTPOptions{
		URL:            fmt.Sprintf(rbxassetidReadURL, options.AssetID),
		Method:         "GET",
		ResponseFormat: rtypes.FormatSelector{Format: "bin"},
		Headers:        rtypes.HTTPHeaders{}.AppendCookies(options.Cookies),
	})
	if err != nil {
		return nil, err
	}
	var assetResponse struct {
		Location string `json:"location"`
	}
	if err := json.Unmarshal(resp.Body.(types.BinaryString), &assetResponse); err != nil {
		return nil, fmt.Errorf("decode asset response: %s", err)
	}
	resp, err = rbxmk.DoHTTPRequest(s.World, rtypes.HTTPOptions{
		URL:            assetResponse.Location,
		Method:         "GET",
		ResponseFormat: options.Format,
		Headers:        rtypes.HTTPHeaders{}.AppendCookies(options.Cookies),
	})
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

// Write uploads an asset according to the given options.
func (s RBXAssetID) Write(options rtypes.RBXAssetOptions) error {
	if options.Format.Format == "" {
		return fmt.Errorf("must specify Format for encoding")
	}
	_, err := rbxmk.DoHTTPRequest(s.World, rtypes.HTTPOptions{
		URL:           fmt.Sprintf(rbxassetidWriteURL, options.AssetID),
		Method:        "POST",
		RequestFormat: options.Format,
		Headers:       rtypes.HTTPHeaders{}.AppendCookies(options.Cookies),
		Body:          options.Body,
	})
	return err
}
