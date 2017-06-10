package format

import (
	"errors"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
	"io"
	"io/ioutil"
)

func init() {
	register(rbxmk.FormatInfo{
		Name:           "Text",
		Ext:            "txt",
		Init:           func(_ *rbxmk.Options) rbxmk.Format { return &TextFormat{Binary: false} },
		InputDrills:    nil,
		OutputDrills:   nil,
		OutputResolver: ResolveOutputSource,
	})
	register(rbxmk.FormatInfo{
		Name:           "Binary",
		Ext:            "bin",
		Init:           func(_ *rbxmk.Options) rbxmk.Format { return &TextFormat{Binary: true} },
		InputDrills:    nil,
		OutputDrills:   nil,
		OutputResolver: ResolveOutputSource,
	})
}

type TextFormat struct {
	Binary bool
}

func (f TextFormat) Decode(r io.Reader) (src *rbxmk.Source, err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var v rbxfile.Value
	if f.Binary {
		v = rbxfile.ValueBinaryString(b)
	} else {
		v = rbxfile.ValueString(b)
	}
	return &rbxmk.Source{Values: []rbxfile.Value{v}}, nil
}

func (TextFormat) CanEncode(src *rbxmk.Source) bool {
	if len(src.Instances) > 0 || len(src.Properties) > 0 || len(src.Values) != 1 {
		return false
	}
	switch src.Values[0].(type) {
	case rbxfile.ValueString, rbxfile.ValueProtectedString, rbxfile.ValueBinaryString:
		return true
	}
	return false
}

func (TextFormat) Encode(w io.Writer, src *rbxmk.Source) (err error) {
	switch v := src.Values[0].(type) {
	case rbxfile.ValueString:
		_, err = w.Write([]byte(v))
	case rbxfile.ValueProtectedString:
		_, err = w.Write([]byte(v))
	case rbxfile.ValueBinaryString:
		_, err = w.Write([]byte(v))
	default:
		return errors.New("unexpected value type: " + v.Type().String())
	}
	return nil
}
