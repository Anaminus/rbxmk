package main

import (
	"io"
)

type FormatInfo struct {
	Name           string
	Ext            string
	Init           InitFormat
	InputDrills    []InputDrill
	OutputDrills   []OutputDrill
	OutputResolver OutputResolver
}

// InputDrill is used to retrieve data within a Source. inref is used to drill
// to data within the Source, returning the result as another Source.
//
// InputDrill also returns inref after it has been processed. If inref is
// empty, then an EOD (end of drill) error is returned, as well as the
// original Source and reference.
type InputDrill func(opt *Options, insrc *Source, inref []string) (outsrc *Source, outref []string, err error)

// OutputDrill is used to locate data within a Source. inaddr points to some
// data within a Source. After inaddr is resolved, inref is used to drill
// further into the data, returning a SourceAddress that points to the data.
//
// OutputDrill also returns inref after it has been processed. If inref is
// empty, then an EOD (end of drill) error is returned, as well as the
// original Source and reference.
type OutputDrill func(opt *Options, inaddr SourceAddress, inref []string) (outaddr SourceAddress, outref []string, err error)

// OutputResolver is used to apply to contents of a Source to a location
// pointed to by a SourceAddress. A reference is provided for context.
type OutputResolver func(ref []string, addr SourceAddress, src *Source) (err error)

type Format interface {
	Decode(r io.Reader) (src *Source, err error)
	CanEncode(src *Source) bool
	Encode(w io.Writer, src *Source) (err error)
}

type InitFormat func(opt *Options) Format

type Formats map[string]*FormatInfo

var DefaultFormats = Formats{}

func (fs Formats) Register(f FormatInfo) {
	if _, registered := fs[f.Ext]; registered {
		panic("format already registered")
	}
	if f.Init == nil {
		panic("format must have Init function")
	}
	if f.OutputResolver == nil {
		panic("format must have OutputResolver function")
	}

	id := make([]InputDrill, len(f.InputDrills))
	copy(id, f.InputDrills)
	f.InputDrills = id

	od := make([]OutputDrill, len(f.OutputDrills))
	copy(od, f.OutputDrills)
	f.OutputDrills = od

	fs[f.Ext] = &f
}

func (fs Formats) Name(ext string) (name string, registered bool) {
	var f *FormatInfo
	if f, registered = fs[ext]; !registered {
		return "", false
	}
	return f.Name, true
}

func (fs Formats) Init(ext string, opt *Options) (format Format, registered bool) {
	var f *FormatInfo
	if f, registered = fs[ext]; !registered {
		return nil, false
	}
	return f.Init(opt), true
}

func (fs Formats) InputDrills(ext string) (drills []InputDrill, registered bool) {
	var f *FormatInfo
	if f, registered = fs[ext]; !registered {
		return nil, false
	}
	drills = make([]InputDrill, len(f.InputDrills))
	copy(drills, f.InputDrills)
	return drills, true
}

func (fs Formats) OutputDrills(ext string) (drills []OutputDrill, registered bool) {
	var f *FormatInfo
	if f, registered = fs[ext]; !registered {
		return nil, false
	}
	drills = make([]OutputDrill, len(f.OutputDrills))
	copy(drills, f.OutputDrills)
	return f.OutputDrills, true
}

func (fs Formats) OutputResolver(ext string) (resolver OutputResolver, registered bool) {
	var f *FormatInfo
	if f, registered = fs[ext]; !registered {
		return nil, false
	}
	return f.OutputResolver, true
}
