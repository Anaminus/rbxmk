package library

import (
	"encoding/json"
	"fmt"
	"strings"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/formats"
	"github.com/anaminus/rbxmk/reflect"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

const rbxassetidReadURL = "https://assetdelivery.roblox.com/v1/assetId/%d"
const rbxassetidWriteURL = "https://data.roblox.com/Data/Upload.ashx?assetid=%d"

func init() { register(RbxAssetId) }

var RbxAssetId = rbxmk.Library{
	Name:     "rbxassetid",
	Import:   []string{"rbxassetid"},
	Priority: 10,
	Open:     openRbxAssetId,
	Dump:     dumpRbxAssetId,
	Types: []func() rbxmk.Reflector{
		reflect.RbxAssetOptions,
	},
}

func openRbxAssetId(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 2)
	lib.RawSetString("read", s.WrapFunc(rbxassetidRead))
	lib.RawSetString("write", s.WrapFunc(rbxassetidWrite))
	return lib
}

func rbxassetidRead(s rbxmk.State) int {
	options := s.Pull(1, rtypes.T_RbxAssetOptions).(rtypes.RbxAssetOptions)
	body, err := RbxAssetIdSource{World: s.World}.Read(options)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	return s.Push(body)
}

func rbxassetidWrite(s rbxmk.State) int {
	options := s.Pull(1, rtypes.T_RbxAssetOptions).(rtypes.RbxAssetOptions)
	err := RbxAssetIdSource{World: s.World}.Write(options)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	return 0
}

func dumpRbxAssetId(s rbxmk.State) dump.Library {
	lib := dump.Library{
		Struct: dump.Struct{
			Fields: dump.Fields{
				"read": dump.Function{
					Parameters: dump.Parameters{
						{Name: "options", Type: dt.Prim(rtypes.T_RbxAssetOptions)},
					},
					Returns: dump.Parameters{
						{Name: "value", Type: dt.Prim(rtypes.T_Any)},
					},
					CanError:    true,
					Summary:     "Libraries/rbxassetid:Fields/read/Summary",
					Description: "Libraries/rbxassetid:Fields/read/Description",
				},
				"write": dump.Function{
					Parameters: dump.Parameters{
						{Name: "options", Type: dt.Prim(rtypes.T_RbxAssetOptions)},
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

// RbxAssetIdSource provides access to assets on the Roblox website.
type RbxAssetIdSource struct {
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
func (s RbxAssetIdSource) Read(options rtypes.RbxAssetOptions) (body types.Value, err error) {
	if options.Format.Format == "" {
		return nil, fmt.Errorf("must specify Format for decoding")
	}
	resp, err := rbxmk.DoHttpRequest(s.World, rtypes.HttpOptions{
		URL:            fmt.Sprintf(rbxassetidReadURL, options.AssetId),
		Method:         "GET",
		ResponseFormat: rtypes.FormatSelector{Format: formats.F_Binary},
		Headers:        rtypes.HttpHeaders{}.AppendCookies(options.Cookies),
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
	resp, err = rbxmk.DoHttpRequest(s.World, rtypes.HttpOptions{
		URL:            assetResponse.Location,
		Method:         "GET",
		ResponseFormat: options.Format,
		Headers:        rtypes.HttpHeaders{}.AppendCookies(options.Cookies),
	})
	if err != nil {
		return nil, fmt.Errorf("get asset content: %w", err)
	}
	return resp.Body, nil
}

// Write uploads an asset according to the given options.
func (s RbxAssetIdSource) Write(options rtypes.RbxAssetOptions) error {
	if options.Format.Format == "" {
		return fmt.Errorf("must specify Format for encoding")
	}
	_, err := rbxmk.DoHttpRequest(s.World, rtypes.HttpOptions{
		URL:           fmt.Sprintf(rbxassetidWriteURL, options.AssetId),
		Method:        "POST",
		RequestFormat: options.Format,
		Headers:       rtypes.HttpHeaders{}.AppendCookies(options.Cookies),
		Body:          options.Body,
	})
	return err
}

// Create creates and uploads a new asset according to the given options.
// Returns the ID of the created asset.
func (s RbxAssetIdSource) Create(options rtypes.RbxAssetOptions) (assetID int64, err error) {
	if options.Format.Format == "" {
		return -1, fmt.Errorf("must specify Format for encoding")
	}
	//TODO: Implement.
	return -1, fmt.Errorf("creating new assets not implemented")
}
