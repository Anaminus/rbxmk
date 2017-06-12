package format

import (
	"errors"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
	"io"
	"io/ioutil"
)

func init() {
	register(rbxmk.FormatInfo{
		Name:           "Lua Script",
		Ext:            "lua",
		Init:           func(_ *rbxmk.Options) rbxmk.Format { return &LuaFormat{Type: LuaScript} },
		InputDrills:    nil,
		OutputDrills:   nil,
		OutputResolver: ResolveOutputSource,
	})
	register(rbxmk.FormatInfo{
		Name:           "Lua LocalScript",
		Ext:            "local.lua",
		Init:           func(_ *rbxmk.Options) rbxmk.Format { return &LuaFormat{Type: LuaLocalScript} },
		InputDrills:    nil,
		OutputDrills:   nil,
		OutputResolver: ResolveOutputSource,
	})
	register(rbxmk.FormatInfo{
		Name:           "Lua ModuleScript",
		Ext:            "module.lua",
		Init:           func(_ *rbxmk.Options) rbxmk.Format { return &LuaFormat{Type: LuaModuleScript} },
		InputDrills:    nil,
		OutputDrills:   nil,
		OutputResolver: ResolveOutputSource,
	})
}

type LuaType uint8

const (
	LuaValue LuaType = iota
	LuaScript
	LuaLocalScript
	LuaModuleScript
)

func (t LuaType) ClassName() string {
	switch t {
	case LuaScript:
		return "Script"
	case LuaLocalScript:
		return "LocalScript"
	case LuaModuleScript:
		return "ModuleScript"
	}
	return ""
}

type LuaFormat struct {
	Type LuaType
}

func (f LuaFormat) Decode(r io.Reader) (src *rbxmk.Source, err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	if f.Type.ClassName() == "" {
		return &rbxmk.Source{Values: []rbxfile.Value{rbxfile.ValueProtectedString(b)}}, nil
	}
	script := rbxfile.NewInstance(f.Type.ClassName(), nil)
	script.Set("Source", rbxfile.ValueProtectedString(b))
	return &rbxmk.Source{Instances: []*rbxfile.Instance{script}}, nil
}

func (f LuaFormat) CanEncode(src *rbxmk.Source) bool {
	if f.Type.ClassName() == "" {
		if len(src.Instances) > 0 || len(src.Properties) > 0 || len(src.Values) != 1 {
			return false
		}
		switch src.Values[0].(type) {
		case rbxfile.ValueString, rbxfile.ValueProtectedString, rbxfile.ValueBinaryString:
			return true
		}
		return false
	}
	if len(src.Instances) != 1 || len(src.Properties) > 0 || len(src.Values) > 0 {
		return false
	}
	return src.Instances[0].ClassName == f.Type.ClassName()
}

func (f LuaFormat) Encode(w io.Writer, src *rbxmk.Source) (err error) {
	var v rbxfile.Value
	switch f.Type {
	case LuaScript, LuaLocalScript, LuaModuleScript:
		if v = src.Instances[0].Get("Source"); v == nil {
			return
		}
	default:
		v = src.Values[0]
	}
	switch v := v.(type) {
	case rbxfile.ValueString:
		_, err = w.Write([]byte(v))
	case rbxfile.ValueProtectedString:
		_, err = w.Write([]byte(v))
	case rbxfile.ValueBinaryString:
		_, err = w.Write([]byte(v))
	default:
		return errors.New("unexpected value type: " + v.Type().String())
	}
	return
}
