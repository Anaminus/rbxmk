package library

import (
	"encoding/json"
	"fmt"
	"strings"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/reflect"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

const rbxassetidReadURL = "https://assetdelivery.roblox.com/v1/assetId/%d"
const rbxassetidWriteURL = "https://data.roblox.com/Data/Upload.ashx?assetid=%d"

func init() { register(RBXAssetID) }

var RBXAssetID = rbxmk.Library{
	Name:       "rbxassetid",
	ImportedAs: "rbxassetid",
	Priority:   10,
	Open:       openRBXAssetID,
	Dump:       dumpRBXAssetID,
	Types: []func() rbxmk.Reflector{
		reflect.RBXAssetOptions,
	},
}

func openRBXAssetID(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 2)
	lib.RawSetString("read", s.WrapFunc(rbxassetidRead))
	lib.RawSetString("write", s.WrapFunc(rbxassetidWrite))
	return lib
}

func rbxassetidRead(s rbxmk.State) int {
	options := s.Pull(1, reflect.T_RBXAssetOptions).(rtypes.RBXAssetOptions)
	body, err := RBXAssetIDSource{World: s.World}.Read(options)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	return s.Push(body)
}

func rbxassetidWrite(s rbxmk.State) int {
	options := s.Pull(1, reflect.T_RBXAssetOptions).(rtypes.RBXAssetOptions)
	err := RBXAssetIDSource{World: s.World}.Write(options)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	return 0
}

func dumpRBXAssetID(s rbxmk.State) dump.Library {
	lib := dump.Library{
		Struct: dump.Struct{
			Fields: dump.Fields{
				"read": dump.Function{
					Parameters: dump.Parameters{
						{Name: "options", Type: dt.Prim(reflect.T_RBXAssetOptions)},
					},
					Returns: dump.Parameters{
						{Name: "value", Type: dt.Prim("any")},
					},
					CanError:    true,
					Summary:     "Libraries/rbxassetid:Fields/read/Summary",
					Description: "Libraries/rbxassetid:Fields/read/Description",
				},
				"write": dump.Function{
					Parameters: dump.Parameters{
						{Name: "options", Type: dt.Prim(reflect.T_RBXAssetOptions)},
					},
					CanError:    true,
					Summary:     "Libraries/rbxassetid:Fields/write/Summary",
					Description: "Libraries/rbxassetid:Fields/write/Description",
				},
			},
			Summary:     "Libraries/rbxassetid:Summary",
			Description: "Libraries/rbxassetid:Description",
		},
		Types: dump.TypeDefs{},
	}
	return lib
}

// RBXAssetIDSource provides access to assets on the Roblox website.
type RBXAssetIDSource struct {
	*rbxmk.World
}

type assetError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err assetError) Error() string {
	return fmt.Sprintf("%d: %s", err.Code, err.Message)
}

// Read downloads an asset according to the given options.
func (s RBXAssetIDSource) Read(options rtypes.RBXAssetOptions) (body types.Value, err error) {
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
		return nil, fmt.Errorf("get asset location: %w", err)
	}
	var assetResponse struct {
		Location string       `json:"location"`
		Errors   []assetError `json:"errors"`
	}
	if err := json.Unmarshal([]byte(resp.Body.(types.BinaryString)), &assetResponse); err != nil {
		return nil, fmt.Errorf("decode asset response: %s", err)
	}
	switch n := len(assetResponse.Errors); n {
	case 0:
	case 1:
		return nil, assetResponse.Errors[0]
	default:
		s := make([]string, n)
		for i, err := range assetResponse.Errors {
			s[i] = "\t" + err.Error()
		}
		return nil, fmt.Errorf("response errors:\n%s", strings.Join(s, "\n"))
	}
	resp, err = rbxmk.DoHTTPRequest(s.World, rtypes.HTTPOptions{
		URL:            assetResponse.Location,
		Method:         "GET",
		ResponseFormat: options.Format,
		Headers:        rtypes.HTTPHeaders{}.AppendCookies(options.Cookies),
	})
	if err != nil {
		return nil, fmt.Errorf("get asset content: %w", err)
	}
	return resp.Body, nil
}

// Write uploads an asset according to the given options.
func (s RBXAssetIDSource) Write(options rtypes.RBXAssetOptions) error {
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

// Create creates and uploads a new asset according to the given options.
// Returns the ID of the created asset.
func (s RBXAssetIDSource) Create(options rtypes.RBXAssetOptions) (assetID int64, err error) {
	if options.Format.Format == "" {
		return -1, fmt.Errorf("must specify Format for encoding")
	}
	//TODO: Implement.
	return -1, fmt.Errorf("creating new assets not implemented")
}
