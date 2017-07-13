package format

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/types"
	"github.com/robloxapi/rbxfile"
	"io"
	"io/ioutil"
)

func init() {
	Formats.Register(rbxmk.Format{
		Name: "Text",
		Ext:  "txt",
		Codec: func(rbxmk.Options, interface{}) rbxmk.FormatCodec {
			return &TextCodec{Binary: false}
		},
	})
	Formats.Register(rbxmk.Format{
		Name: "Binary",
		Ext:  "bin",
		Codec: func(rbxmk.Options, interface{}) rbxmk.FormatCodec {
			return &TextCodec{Binary: true}
		},
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
		*data = &types.Stringlike{ValueType: rbxfile.TypeBinaryString, Bytes: b}
	} else {
		*data = &types.Stringlike{ValueType: rbxfile.TypeString, Bytes: b}
	}
	return nil
}

func (c *TextCodec) Encode(w io.Writer, data rbxmk.Data) (err error) {
	if s := types.NewStringlike(data); s != nil {
		data = s
	}
	switch v := data.(type) {
	case *types.Stringlike:
		_, err = w.Write(v.Bytes)
	case nil:
		// Write nothing.
	default:
		err = rbxmk.NewDataTypeError(data)
	}
	return err
}
