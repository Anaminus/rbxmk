package format

import (
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
	"io"
	"io/ioutil"
)

func init() {
	rbxmk.DefaultFormats.Register(rbxmk.FormatInfo{
		Name:           "Lua",
		Ext:            "lua",
		Init:           func(_ *rbxmk.Options) rbxmk.Format { return &LuaFormat{} },
		InputDrills:    nil,
		OutputDrills:   nil,
		OutputResolver: ResolveOutputSource,
	})
}

type LuaFormat struct{}

func (LuaFormat) Decode(r io.Reader) (src *rbxmk.Source, err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return &rbxmk.Source{Values: []rbxfile.Value{rbxfile.ValueProtectedString(b)}}, nil
}
func (LuaFormat) CanEncode(src *rbxmk.Source) bool {
	if len(src.Instances) > 0 || len(src.Properties) > 0 || len(src.Values) != 1 {
		return false
	}
	_, ok := src.Values[0].(rbxfile.ValueProtectedString)
	return ok
}

func (LuaFormat) Encode(w io.Writer, src *rbxmk.Source) (err error) {
	_, err = w.Write([]byte(src.Values[0].(rbxfile.ValueProtectedString)))
	return
}
