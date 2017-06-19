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
type Drill func(opt *Options, indata Data, inref []string) (outdata Data, outref []string, err error)

type OutputResolver func(ref []string, indata, data Data) (outdata Data, err error)

var EOD = errors.New("end of drill")

type FormatInfo struct {
	Name           string
	Ext            string
	Decoder        InitFormatDecoder
	Encoder        InitFormatEncoder
	InputDrills    []Drill
	OutputDrills   []Drill
	OutputResolver OutputResolver
}

type FormatDecoder interface {
	Decode(data *Data) (err error)
}

type FormatEncoder interface {
	io.Reader
	io.WriterTo
}

type InitFormatDecoder func(opt *Options, r io.Reader) (format FormatDecoder, err error)
type InitFormatEncoder func(opt *Options, data Data) (format FormatEncoder, err error)

type Formats struct {
	f map[string]*FormatInfo
}

func NewFormats() *Formats {
	return &Formats{f: map[string]*FormatInfo{}}
}

func (fs *Formats) Register(f FormatInfo) {
	if _, registered := fs.f[f.Ext]; registered {
		panic("format already registered")
	}
	if f.Decoder == nil {
		panic("format must have Decoder function")
	}
	if f.Encoder == nil {
		panic("format must have Encoder function")
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

func (fs *Formats) Decoder(ext string, opt *Options, r io.Reader) (format FormatDecoder, err error) {
	f, registered := fs.f[ext]
	if !registered {
		return nil, nil
	}
	return f.Decoder(opt, r)
}

func (fs *Formats) Encoder(ext string, opt *Options, data Data) (format FormatEncoder, err error) {
	f, registered := fs.f[ext]
	if !registered {
		return nil, nil
	}
	return f.Encoder(opt, data)
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

func (fs *Formats) OutputResolver(ext string) (resolver OutputResolver) {
	f, registered := fs.f[ext]
	if !registered {
		return nil
	}
	return f.OutputResolver
}
