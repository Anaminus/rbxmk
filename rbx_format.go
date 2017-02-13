package main

import (
	"github.com/robloxapi/rbxapi"
	"github.com/robloxapi/rbxfile"
	"github.com/robloxapi/rbxfile/bin"
	"github.com/robloxapi/rbxfile/xml"
	"io"
)

func init() {
	RegisterFormat(func(opt *Options) Format { return &RBXFormat{false, false, opt.API} }) // rbxl
	RegisterFormat(func(opt *Options) Format { return &RBXFormat{false, true, opt.API} })  // rbxlx
	RegisterFormat(func(opt *Options) Format { return &RBXFormat{true, false, opt.API} })  // rbxm
	RegisterFormat(func(opt *Options) Format { return &RBXFormat{true, true, opt.API} })   // rbxmx
}

////////////////////////////////

type RBXFormat struct {
	Model bool // Model or Place
	XML   bool // XML or Binary
	API   *rbxapi.API
}

func (f RBXFormat) Name() string {
	if f.Model {
		if f.XML {
			return "RBXMX"
		} else {
			return "RBXM"
		}
	} else {
		if f.XML {
			return "RBXLX"
		} else {
			return "RBXL"
		}
	}
}

func (f RBXFormat) Ext() string {
	if f.Model {
		if f.XML {
			return "rbxmx"
		} else {
			return "rbxm"
		}
	} else {
		if f.XML {
			return "rbxlx"
		} else {
			return "rbxl"
		}
	}
}

func (f *RBXFormat) Decode(r io.Reader) (src *Source, err error) {
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
	return &Source{Instances: root.Instances}, nil
}

func (f *RBXFormat) CanEncode(src *Source) bool {
	return len(src.Properties) == 0 && len(src.Values) == 0
}

func (f *RBXFormat) Encode(w io.Writer, src *Source) (err error) {
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
