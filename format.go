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

// Resolver is used to merge one Data into another. A reference is provided
// for context.
type Resolver func(ref []string, indata, data Data) (outdata Data, err error)

var EOD = errors.New("end of drill")

type Format struct {
	Name         string
	Ext          string
	Codec        InitFormatCodec
	InputDrills  []Drill
	OutputDrills []Drill
	Resolver     Resolver
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

type InitFormatCodec func(opt *Options) (codec FormatCodec)

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

func (fs *Formats) Decoder(ext string, opt *Options) (dec FormatDecoder) {
	f, registered := fs.f[ext]
	if !registered {
		return nil
	}
	return f.Codec(opt)
}

func (fs *Formats) Decode(ext string, opt *Options, r io.Reader, data *Data) (err error) {
	f, registered := fs.f[ext]
	if !registered {
		return nil
	}
	return f.Codec(opt).Decode(r, data)
}

func (fs *Formats) Encoder(ext string, opt *Options) (enc FormatEncoder) {
	f, registered := fs.f[ext]
	if !registered {
		return nil
	}
	return f.Codec(opt)
}

func (fs *Formats) Encode(ext string, opt *Options, w io.Writer, data Data) (err error) {
	f, registered := fs.f[ext]
	if !registered {
		return nil
	}
	return f.Codec(opt).Encode(w, data)
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

func (fs *Formats) Resolver(ext string) (resolver Resolver) {
	f, registered := fs.f[ext]
	if !registered {
		return nil
	}
	return f.Resolver
}
