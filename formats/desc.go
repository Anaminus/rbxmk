package formats

import (
	"bytes"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump/json"
	"github.com/robloxapi/types"
)

func Desc() rbxmk.Format {
	return rbxmk.Format{
		Name: "desc.json",
		Decode: func(b []byte) (v types.Value, err error) {
			root, err := json.Decode(bytes.NewReader(b))
			if err != nil {
				return nil, err
			}
			return rtypes.RootDesc{Root: root}, nil
		},
		Encode: func(v types.Value) (b []byte, err error) {
			root := v.(rtypes.RootDesc).Root
			var buf bytes.Buffer
			if err := json.Encode(&buf, root); err != nil {
				return nil, err
			}
			return buf.Bytes(), nil
		},
	}
}
