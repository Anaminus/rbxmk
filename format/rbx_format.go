package format

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/config"
	"github.com/anaminus/rbxmk/types"
	"github.com/robloxapi/rbxapi"
	"github.com/robloxapi/rbxfile"
	"github.com/robloxapi/rbxfile/bin"
	"github.com/robloxapi/rbxfile/xml"
	"io"
)

func init() {
	Formats.Register(rbxmk.Format{
		Name: "RBXL",
		Ext:  "rbxl",
		Codec: func(opt rbxmk.Options, ctx interface{}) (codec rbxmk.FormatCodec) {
			return &RBXCodec{model: false, xml: false, api: config.API(opt)}
		},
	})
	Formats.Register(rbxmk.Format{
		Name: "RBXLX",
		Ext:  "rbxlx",
		Codec: func(opt rbxmk.Options, ctx interface{}) (codec rbxmk.FormatCodec) {
			return &RBXCodec{model: false, xml: true, api: config.API(opt)}
		},
	})
	Formats.Register(rbxmk.Format{
		Name: "RBXM",
		Ext:  "rbxm",
		Codec: func(opt rbxmk.Options, ctx interface{}) (codec rbxmk.FormatCodec) {
			return &RBXCodec{model: true, xml: false, api: config.API(opt)}
		},
	})
	Formats.Register(rbxmk.Format{
		Name: "RBXMX",
		Ext:  "rbxmx",
		Codec: func(opt rbxmk.Options, ctx interface{}) (codec rbxmk.FormatCodec) {
			return &RBXCodec{model: true, xml: true, api: config.API(opt)}
		},
	})
}

type RBXCodec struct {
	model bool // Model or Place
	xml   bool // XML or Binary
	api   rbxapi.Root
}

func (c *RBXCodec) Decode(r io.Reader, data *rbxmk.Data) (err error) {
	var root *rbxfile.Root
	if c.xml {
		root, err = xml.Deserialize(r, c.api)
	} else {
		if c.model {
			root, err = bin.DeserializeModel(r, c.api)
		} else {
			root, err = bin.DeserializePlace(r, c.api)
		}
	}
	if err != nil {
		return err
	}
	instances := types.Instances(root.Instances)
	*data = &instances
	return nil
}

func (c *RBXCodec) Encode(w io.Writer, data rbxmk.Data) (err error) {
	var root *rbxfile.Root
	switch v := data.(type) {
	case *types.Instances:
		root = &rbxfile.Root{Instances: []*rbxfile.Instance(*v)}
	case types.Instance:
		root = &rbxfile.Root{Instances: []*rbxfile.Instance{v.Instance}}
	case nil:
		root = &rbxfile.Root{}
	default:
		return rbxmk.NewDataTypeError(data)
	}
	if c.xml {
		err = xml.Serialize(w, c.api, root)
	} else {
		if c.model {
			err = bin.SerializeModel(w, c.api, root)
		} else {
			err = bin.SerializePlace(w, c.api, root)
		}
	}
	return err
}
