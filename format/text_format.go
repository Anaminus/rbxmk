package format

import (
	"github.com/anaminus/rbxmk"
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
		InputDrills:  nil,
		OutputDrills: nil,
		Merger:       MergeTable,
	})
	Formats.Register(rbxmk.Format{
		Name: "Binary",
		Ext:  "bin",
		Codec: func(rbxmk.Options, interface{}) rbxmk.FormatCodec {
			return &TextCodec{Binary: true}
		},
		InputDrills:  nil,
		OutputDrills: nil,
		Merger:       MergeTable,
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
		*data = &Stringlike{Type: rbxfile.TypeBinaryString, Bytes: b}
	} else {
		*data = &Stringlike{Type: rbxfile.TypeString, Bytes: b}
	}
	return nil
}

func (c *TextCodec) Encode(w io.Writer, data rbxmk.Data) (err error) {
	if s := NewStringlike(data); s != nil {
		data = s
	}
	switch v := data.(type) {
	case *Stringlike:
		_, err = w.Write(v.Bytes)
	case nil:
		// Write nothing.
	default:
		err = rbxmk.NewDataTypeError(data)
	}
	return err
}
