package formats

import (
	"encoding/json"
	"io"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/reflect"
	"github.com/anaminus/rbxmk/rtypes"
	rbxdumpjson "github.com/robloxapi/rbxdump/json"
	"github.com/robloxapi/types"
)

func init() { register(Desc) }
func Desc() rbxmk.Format {
	return rbxmk.Format{
		Name:       "desc.json",
		MediaTypes: []string{"application/json", "text/plain"},
		CanDecode: func(g rtypes.Global, f rbxmk.FormatOptions, typeName string) bool {
			return typeName == rtypes.T_Desc
		},
		Decode: func(g rtypes.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			root, err := rbxdumpjson.Decode(r)
			if err != nil {
				return nil, err
			}
			return &rtypes.Desc{Root: root}, nil
		},
		Encode: func(g rtypes.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			root := v.(*rtypes.Desc).Root
			if err := rbxdumpjson.Encode(w, root); err != nil {
				return err
			}
			return nil
		},
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/desc.json:Summary",
				Description: "Formats/desc.json:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			reflect.Desc,
		},
	}
}

func init() { register(DescPatch) }
func DescPatch() rbxmk.Format {
	return rbxmk.Format{
		Name:        "desc-patch.json",
		EncodeTypes: []string{rtypes.T_DescActions},
		MediaTypes:  []string{"application/json", "text/plain"},
		CanDecode: func(g rtypes.Global, f rbxmk.FormatOptions, typeName string) bool {
			return typeName == rtypes.T_DescActions
		},
		Decode: func(g rtypes.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			var actions rtypes.DescActions
			j := json.NewDecoder(r)
			if err := j.Decode(&actions); err != nil {
				return nil, err
			}
			return rtypes.DescActions(actions), nil
		},
		Encode: func(g rtypes.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			actions := v.(rtypes.DescActions)
			j := json.NewEncoder(w)
			j.SetIndent("", "\t")
			j.SetEscapeHTML(false)
			return j.Encode(actions)
		},
		Dump: func() dump.Format {
			return dump.Format{
				Summary:     "Formats/desc-patch.json:Summary",
				Description: "Formats/desc-patch.json:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			reflect.DescActions,
		},
	}
}
