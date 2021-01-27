package formats

import (
	"encoding/json"
	"io"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	rbxdumpjson "github.com/robloxapi/rbxdump/json"
	"github.com/robloxapi/types"
)

func init() { register(Desc) }
func Desc() rbxmk.Format {
	return rbxmk.Format{
		Name:       "desc.json",
		MediaTypes: []string{"application/json", "text/plain"},
		CanDecode: func(typeName string) bool {
			return typeName == "RootDesc"
		},
		Decode: func(f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			root, err := rbxdumpjson.Decode(r)
			if err != nil {
				return nil, err
			}
			return &rtypes.RootDesc{Root: root}, nil
		},
		Encode: func(f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			root := v.(*rtypes.RootDesc).Root
			if err := rbxdumpjson.Encode(w, root); err != nil {
				return err
			}
			return nil
		},
	}
}

func init() { register(DescPatch) }
func DescPatch() rbxmk.Format {
	return rbxmk.Format{
		Name:       "desc-patch.json",
		MediaTypes: []string{"application/json", "text/plain"},
		CanDecode: func(typeName string) bool {
			return typeName == "DescActions"
		},
		Decode: func(f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			var actions rtypes.DescActions
			j := json.NewDecoder(r)
			if err := j.Decode(&actions); err != nil {
				return nil, err
			}
			return rtypes.DescActions(actions), nil
		},
		Encode: func(f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			actions := v.(rtypes.DescActions)
			j := json.NewEncoder(w)
			j.SetIndent("", "\t")
			j.SetEscapeHTML(false)
			return j.Encode(actions)
		},
	}
}
