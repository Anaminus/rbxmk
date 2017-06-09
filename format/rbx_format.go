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
	register(rbxmk.FormatInfo{
		Name:           "RBXL",
		Ext:            "rbxl",
		Init:           func(opt *rbxmk.Options) rbxmk.Format { return &RBXFormat{Model: false, XML: false, API: opt.API} },
		InputDrills:    []rbxmk.InputDrill{DrillInputInstance, DrillInputProperty},
		OutputDrills:   []rbxmk.OutputDrill{DrillOutputInstance, DrillOutputProperty},
		OutputResolver: ResolveOutputInstance,
	})
	register(rbxmk.FormatInfo{
		Name:           "RBXLX",
		Ext:            "rbxlx",
		Init:           func(opt *rbxmk.Options) rbxmk.Format { return &RBXFormat{Model: false, XML: true, API: opt.API} },
		InputDrills:    []rbxmk.InputDrill{DrillInputInstance, DrillInputProperty},
		OutputDrills:   []rbxmk.OutputDrill{DrillOutputInstance, DrillOutputProperty},
		OutputResolver: ResolveOutputInstance,
	})
	register(rbxmk.FormatInfo{
		Name:           "RBXM",
		Ext:            "rbxm",
		Init:           func(opt *rbxmk.Options) rbxmk.Format { return &RBXFormat{Model: true, XML: false, API: opt.API} },
		InputDrills:    []rbxmk.InputDrill{DrillInputInstance, DrillInputProperty},
		OutputDrills:   []rbxmk.OutputDrill{DrillOutputInstance, DrillOutputProperty},
		OutputResolver: ResolveOutputInstance,
	})
	register(rbxmk.FormatInfo{
		Name:           "RBXMX",
		Ext:            "rbxmx",
		Init:           func(opt *rbxmk.Options) rbxmk.Format { return &RBXFormat{Model: true, XML: true, API: opt.API} },
		InputDrills:    []rbxmk.InputDrill{DrillInputInstance, DrillInputProperty},
		OutputDrills:   []rbxmk.OutputDrill{DrillOutputInstance, DrillOutputProperty},
		OutputResolver: ResolveOutputInstance,
	})
}

////////////////////////////////

type RBXFormat struct {
	Model bool // Model or Place
	XML   bool // XML or Binary
	API   *rbxapi.API
}

func (f *RBXFormat) Decode(r io.Reader) (src *rbxmk.Source, err error) {
	var root *rbxfile.Root
	if f.XML {
		root, err = xml.Deserialize(r, f.API)
	} else {
		if f.Model {
			root, err = bin.DeserializeModel(r, f.API)
		} else {
			root, err = bin.DeserializePlace(r, f.API)
		}
	}

	if err != nil {
		return nil, err
	}
	return &rbxmk.Source{Instances: root.Instances}, nil
}

func (f *RBXFormat) CanEncode(src *rbxmk.Source) bool {
	return len(src.Properties) == 0 && len(src.Values) == 0
}

func (f *RBXFormat) Encode(w io.Writer, src *rbxmk.Source) (err error) {
	root := &rbxfile.Root{Instances: src.Instances}
	if f.XML {
		err = xml.Serialize(w, f.API, root)
	} else {
		if f.Model {
			err = bin.SerializeModel(w, f.API, root)
		} else {
			err = bin.SerializePlace(w, f.API, root)
		}
	}
	return err
}
