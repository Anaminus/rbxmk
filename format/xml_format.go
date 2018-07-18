package format

import (
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/config"
	"github.com/anaminus/rbxmk/types"
	"github.com/robloxapi/rbxapi"
	"github.com/robloxapi/rbxfile"
	"github.com/robloxapi/rbxfile/xml"
	"io"
)

func init() {
	Formats.Register(rbxmk.Format{
		Name: "XML Properties",
		Ext:  "properties.xml",
		Codec: func(opt *rbxmk.Options, ctx interface{}) rbxmk.FormatCodec {
			return &XMLCodec{API: config.API(opt)}
		},
	})
}

type XMLCodec struct {
	API rbxapi.Root
}

func (c *XMLCodec) Decode(r io.Reader, data *rbxmk.Data) (err error) {
	doc := &xml.Document{}
	if _, err = doc.ReadFrom(r); err != nil {
		return err
	}
	if doc.Root == nil || doc.Root.StartName != "Properties" {
		return fmt.Errorf("expected Properties tag")
	}
	inst := &rbxfile.Instance{Properties: make(map[string]rbxfile.Value, len(doc.Root.Tags))}
	xml.RobloxCodec{API: c.API}.DecodeProperties(doc.Root.Tags, inst, nil)
	*data = types.Properties(inst.Properties)
	return nil
}

func (c *XMLCodec) Encode(w io.Writer, data rbxmk.Data) (err error) {
	switch v := data.(type) {
	case *types.Instances:
		if len(*v) > 0 {
			data = types.Properties((*v)[0].Properties)
		}
	case *types.Instance:
		data = types.Properties(v.Properties)
	case types.Property:
		data = types.Properties{v.Name: v.Properties[v.Name]}
	case nil:
		data = types.Properties{}
	}

	props, ok := data.(types.Properties)
	if !ok {
		return rbxmk.NewDataTypeError(data)
	}

	doc := &xml.Document{Indent: "\t"}
	root := &xml.Tag{StartName: "Properties"}
	doc.Root = root

	inst := &rbxfile.Instance{Properties: map[string]rbxfile.Value(props)}
	root.Tags = xml.RobloxCodec{API: c.API}.EncodeProperties(inst)
	_, err = doc.WriteTo(w)
	return err
}
