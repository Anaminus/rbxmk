package main

import (
	"github.com/robloxapi/rbxfile"
	"io"
	"io/ioutil"
)

func init() {
	RegisterFormat(func(_ *Options) Format { return &LuaFormat{} })
}

type LuaFormat struct{}

func (LuaFormat) Name() string {
	return "Lua"
}

func (LuaFormat) Ext() string {
	return "lua"
}

func (LuaFormat) Decode(r io.Reader) (src *Source, err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return &Source{Values: []rbxfile.Value{rbxfile.ValueProtectedString(b)}}, nil
}
func (LuaFormat) CanEncode(src *Source) bool {
	if len(src.Instances) > 0 || len(src.Properties) > 0 || len(src.Values) != 1 {
		return false
	}
	_, ok := src.Values[0].(rbxfile.ValueProtectedString)
	return ok
}

func (LuaFormat) Encode(w io.Writer, src *Source) (err error) {
	_, err = w.Write([]byte(src.Values[0].(rbxfile.ValueProtectedString)))
	return
}
