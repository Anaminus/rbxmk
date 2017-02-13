package main

import (
	"io"
)

type Format interface {
	Name() string
	Ext() string
	Decode(r io.Reader) (src *Source, err error)
	CanEncode(src *Source) bool
	Encode(w io.Writer, src *Source) (err error)
}

type NewFormat func(opt *Options) Format

var registeredFormats = map[string]NewFormat{}

func RegisterFormat(f NewFormat) {
	if f == nil {
		panic("cannot register nil format")
	}
	format := f(&Options{})
	if format == nil {
		panic("format function must return non-nil format")
	}
	ext := format.Ext()
	if _, registered := registeredFormats[ext]; registered {
		panic("format already registered")
	}
	registeredFormats[ext] = f
}
