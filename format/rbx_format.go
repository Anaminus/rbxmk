package format

import (
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxapi"
	"github.com/robloxapi/rbxfile"
	"github.com/robloxapi/rbxfile/bin"
	"github.com/robloxapi/rbxfile/xml"
	"io"
)

func init() {
	register(rbxmk.Format{
		Name: "RBXL",
		Ext:  "rbxl",
		Codec: func(opt rbxmk.Options, ctx interface{}) (codec rbxmk.FormatCodec) {
			return &RBXCodec{model: false, xml: false, api: opt.API}
		},
		InputDrills:  []rbxmk.Drill{DrillInstance, DrillInstanceProperty},
		OutputDrills: []rbxmk.Drill{DrillInstance, DrillInstanceProperty},
		Merger:       MergeInstance,
	})
	register(rbxmk.Format{
		Name: "RBXLX",
		Ext:  "rbxlx",
		Codec: func(opt rbxmk.Options, ctx interface{}) (codec rbxmk.FormatCodec) {
			return &RBXCodec{model: false, xml: true, api: opt.API}
		},
		InputDrills:  []rbxmk.Drill{DrillInstance, DrillInstanceProperty},
		OutputDrills: []rbxmk.Drill{DrillInstance, DrillInstanceProperty},
		Merger:       MergeInstance,
	})
	register(rbxmk.Format{
		Name: "RBXM",
		Ext:  "rbxm",
		Codec: func(opt rbxmk.Options, ctx interface{}) (codec rbxmk.FormatCodec) {
			return &RBXCodec{model: true, xml: false, api: opt.API}
		},
		InputDrills:  []rbxmk.Drill{DrillInstance, DrillInstanceProperty},
		OutputDrills: []rbxmk.Drill{DrillInstance, DrillInstanceProperty},
		Merger:       MergeInstance,
	})
	register(rbxmk.Format{
		Name: "RBXMX",
		Ext:  "rbxmx",
		Codec: func(opt rbxmk.Options, ctx interface{}) (codec rbxmk.FormatCodec) {
			return &RBXCodec{model: true, xml: true, api: opt.API}
		},
		InputDrills:  []rbxmk.Drill{DrillInstance, DrillInstanceProperty},
		OutputDrills: []rbxmk.Drill{DrillInstance, DrillInstanceProperty},
		Merger:       MergeInstance,
	})
}

type RBXCodec struct {
	model bool // Model or Place
	xml   bool // XML or Binary
	api   *rbxapi.API
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
	*data = root.Instances
	return nil
}

func (c *RBXCodec) Encode(w io.Writer, data rbxmk.Data) (err error) {
	instances, ok := data.([]*rbxfile.Instance)
	if !ok {
		return NewDataTypeError(data)
	}
	root := &rbxfile.Root{Instances: instances}
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
