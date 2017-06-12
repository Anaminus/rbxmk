package format

import (
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxapi"
	"github.com/robloxapi/rbxfile"
	"github.com/robloxapi/rbxfile/xml"
	"io"
)

func init() {
	register(rbxmk.FormatInfo{
		Name:           "XML",
		Ext:            "xml",
		Init:           func(opt *rbxmk.Options) rbxmk.Format { return &XMLFormat{API: opt.API} },
		InputDrills:    nil,
		OutputDrills:   nil,
		OutputResolver: ResolveOutputSource,
	})
}

type XMLFormat struct {
	API *rbxapi.API
}

func (f XMLFormat) Decode(r io.Reader) (src *rbxmk.Source, err error) {
	doc := &xml.Document{}
	if _, err = doc.ReadFrom(r); err != nil {
		return nil, err
	}
	if doc.Root == nil || doc.Root.StartName != "Properties" {
		return &rbxmk.Source{}, nil
	}
	inst := &rbxfile.Instance{Properties: make(map[string]rbxfile.Value, len(doc.Root.Tags))}
	xml.RobloxCodec{API: f.API}.DecodeProperties(doc.Root.Tags, inst, nil)
	return &rbxmk.Source{Properties: inst.Properties}, nil
}

func (f XMLFormat) CanEncode(src *rbxmk.Source) bool {
	return len(src.Instances) == 0 && len(src.Values) == 0
}

func (f XMLFormat) Encode(w io.Writer, src *rbxmk.Source) (err error) {
	doc := &xml.Document{Indent: "\t"}
	root := &xml.Tag{StartName: "Properties"}
	doc.Root = root

	inst := &rbxfile.Instance{Properties: src.Properties}
	root.Tags = xml.RobloxCodec{API: f.API}.EncodeProperties(inst)
	_, err = doc.WriteTo(w)
	return err
}
