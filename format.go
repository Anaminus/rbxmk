package rbxmk

import (
	"errors"
	"io"
)

type Data interface{}

// Drill receives a Data and drills into it using inref, returning a Data to
// represent the result. It also returns the reference after it has been
// parsed. In case of an error, the original Data and inref is returned. If
// inref is empty, then an EOD error is returned.
type Drill func(opt Options, indata Data, inref []string) (outdata Data, outref []string, err error)

// Merger is used to merge input Data into output Data. Three kinds of Data
// are received. rootdata is the top-level data returned by an output scheme
// handler. drilldata is a value within the rootdata, which was selected by
// one or more Drills. If no drills were used, then drilldata will be equal to
// rootdata. Depending on the scheme, both rootdata and drilldata may be nil.
// indata is the data to be merged. Merger returns the resulting Data after it
// has been merged.
type Merger func(opt Options, rootdata, drilldata, indata Data) (outdata Data, err error)

var EOD = errors.New("end of drill")

type Format struct {
	Name         string
	Ext          string
	Codec        InitFormatCodec
	InputDrills  []Drill
	OutputDrills []Drill
	Merger       Merger
	// CanEncode    func(data Data) bool
}

type FormatDecoder interface {
	Decode(r io.Reader, data *Data) (err error)
}

type FormatEncoder interface {
	Encode(w io.Writer, data Data) (err error)
}

type FormatCodec interface {
	FormatDecoder
	FormatEncoder
}

type InitFormatCodec func(opt Options, ctx interface{}) (codec FormatCodec)

type Formats struct {
	f map[string]*Format
}

func NewFormats() *Formats {
	return &Formats{f: map[string]*Format{}}
}

func (fs *Formats) Register(f Format) {
	if _, registered := fs.f[f.Ext]; registered {
		panic("format already registered")
	}
	if f.Codec == nil {
		panic("format must have Codec function")
	}

	id := make([]Drill, len(f.InputDrills))
	copy(id, f.InputDrills)
	f.InputDrills = id

	od := make([]Drill, len(f.OutputDrills))
	copy(od, f.OutputDrills)
	f.OutputDrills = od

	fs.f[f.Ext] = &f
}

func (fs *Formats) Registered(ext string) (registered bool) {
	_, registered = fs.f[ext]
	return registered
}

func (fs *Formats) Name(ext string) (name string) {
	f, registered := fs.f[ext]
	if !registered {
		return ""
	}
	return f.Name
}

func (fs *Formats) Decoder(ext string, opt Options, ctx interface{}) (dec FormatDecoder) {
	f, registered := fs.f[ext]
	if !registered {
		return nil
	}
	return f.Codec(opt, ctx)
}

func (fs *Formats) Decode(ext string, opt Options, ctx interface{}, r io.Reader, data *Data) (err error) {
	f, registered := fs.f[ext]
	if !registered {
		return nil
	}
	return f.Codec(opt, ctx).Decode(r, data)
}

func (fs *Formats) Encoder(ext string, opt Options, ctx interface{}) (enc FormatEncoder) {
	f, registered := fs.f[ext]
	if !registered {
		return nil
	}
	return f.Codec(opt, ctx)
}

func (fs *Formats) Encode(ext string, opt Options, ctx interface{}, w io.Writer, data Data) (err error) {
	f, registered := fs.f[ext]
	if !registered {
		return nil
	}
	return f.Codec(opt, ctx).Encode(w, data)
}

func (fs *Formats) InputDrills(ext string) (drills []Drill) {
	f, registered := fs.f[ext]
	if !registered {
		return nil
	}
	drills = make([]Drill, len(f.InputDrills))
	copy(drills, f.InputDrills)
	return drills
}

func (fs *Formats) OutputDrills(ext string) (drills []Drill) {
	f, registered := fs.f[ext]
	if !registered {
		return nil
	}
	drills = make([]Drill, len(f.OutputDrills))
	copy(drills, f.OutputDrills)
	return f.OutputDrills
}

func (fs *Formats) Merger(ext string) (merger Merger) {
	f, registered := fs.f[ext]
	if !registered {
		return nil
	}
	return f.Merger
}
