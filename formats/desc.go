package formats

import (
	"bytes"
	"encoding/json"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	rbxdumpjson "github.com/robloxapi/rbxdump/json"
	"github.com/robloxapi/types"
)

func init() { register(Desc) }
func Desc() rbxmk.Format {
	return rbxmk.Format{
		Name: "desc.json",
		Decode: func(b []byte) (v types.Value, err error) {
			root, err := rbxdumpjson.Decode(bytes.NewReader(b))
			if err != nil {
				return nil, err
			}
			return &rtypes.RootDesc{Root: root}, nil
		},
		Encode: func(v types.Value) (b []byte, err error) {
			root := v.(*rtypes.RootDesc).Root
			var buf bytes.Buffer
			if err := rbxdumpjson.Encode(&buf, root); err != nil {
				return nil, err
			}
			return buf.Bytes(), nil
		},
	}
}

func init() { register(DescPatch) }
func DescPatch() rbxmk.Format {
	return rbxmk.Format{
		Name: "desc-patch.json",
		Decode: func(b []byte) (v types.Value, err error) {
			var actions rtypes.DescActions
			if err := json.Unmarshal(b, &actions); err != nil {
				return nil, err
			}
			return rtypes.DescActions(actions), nil
		},
		Encode: func(v types.Value) (b []byte, err error) {
			actions := v.(rtypes.DescActions)
			return json.MarshalIndent(actions, "", "\t")
		},
	}
}
