package format

import (
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
	"io"
	"io/ioutil"
)

func init() {
	register(rbxmk.Format{
		Name: "Text",
		Ext:  "txt",
		Codec: func(rbxmk.Options, interface{}) rbxmk.FormatCodec {
			return &TextCodec{Binary: false}
		},
		InputDrills:  nil,
		OutputDrills: nil,
		Resolver:     ResolveOverwrite,
	})
	register(rbxmk.Format{
		Name: "Binary",
		Ext:  "bin",
		Codec: func(rbxmk.Options, interface{}) rbxmk.FormatCodec {
			return &TextCodec{Binary: true}
		},
		InputDrills:  nil,
		OutputDrills: nil,
		Resolver:     ResolveOverwrite,
	})
}

type TextCodec struct {
	Binary bool
}

func (c *TextCodec) Decode(r io.Reader, data *rbxmk.Data) (err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if c.Binary {
		*data = rbxfile.ValueBinaryString(b)
	} else {
		*data = rbxfile.ValueString(b)
	}
	return nil
}

func (c *TextCodec) Encode(w io.Writer, data rbxmk.Data) (err error) {
	switch v := data.(type) {
	case rbxfile.ValueProtectedString:
		_, err = w.Write([]byte(v))
	case rbxfile.ValueBinaryString:
		_, err = w.Write([]byte(v))
	case rbxfile.ValueString:
		_, err = w.Write([]byte(v))
	default:
		err = NewDataTypeError(data)
	}
	return err
}
