package format

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/scheme"
	"github.com/robloxapi/rbxfile"
	"io"
	"io/ioutil"
	"path/filepath"
)

func getFileNameCtx(opt rbxmk.Options, ctx interface{}) string {
	path, _ := ctx.(string)
	if path == "" {
		return path
	}
	base := filepath.Base(path)
	ext := scheme.GuessFileExtension(opt, "", base)
	if ext == "" {
		return base[:len(base)]
	}
	return base[:len(base)-len(ext)-1]
}

func init() {
	register(rbxmk.Format{
		Name: "Lua",
		Ext:  "lua",
		Codec: func(opt rbxmk.Options, ctx interface{}) rbxmk.FormatCodec {
			return &LuaCodec{Type: LuaValue, Name: getFileNameCtx(opt, ctx)}
		},
		InputDrills:  nil,
		OutputDrills: nil,
		Merger:       MergeOverwrite,
	})
	register(rbxmk.Format{
		Name: "Lua Script",
		Ext:  "script.lua",
		Codec: func(opt rbxmk.Options, ctx interface{}) rbxmk.FormatCodec {
			return &LuaCodec{Type: LuaScript, Name: getFileNameCtx(opt, ctx)}
		},
		InputDrills:  nil,
		OutputDrills: nil,
		Merger:       MergeOverwrite,
	})
	register(rbxmk.Format{
		Name: "Lua LocalScript",
		Ext:  "localscript.lua",
		Codec: func(opt rbxmk.Options, ctx interface{}) rbxmk.FormatCodec {
			return &LuaCodec{Type: LuaLocalScript, Name: getFileNameCtx(opt, ctx)}
		},
		InputDrills:  nil,
		OutputDrills: nil,
		Merger:       MergeOverwrite,
	})
	register(rbxmk.Format{
		Name: "Lua ModuleScript",
		Ext:  "modulescript.lua",
		Codec: func(opt rbxmk.Options, ctx interface{}) rbxmk.FormatCodec {
			return &LuaCodec{Type: LuaModuleScript, Name: getFileNameCtx(opt, ctx)}
		},
		InputDrills:  nil,
		OutputDrills: nil,
		Merger:       MergeOverwrite,
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

type LuaCodec struct {
	Type LuaType
	Name string
}

func (c LuaCodec) Decode(r io.Reader, data *rbxmk.Data) (err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if c.Type.ClassName() == "" {
		*data = rbxfile.ValueProtectedString(b)
		return nil
	}
	script := rbxfile.NewInstance(c.Type.ClassName(), nil)
	if c.Name != "" {
		script.SetName(c.Name)
	}
	script.Set("Source", rbxfile.ValueProtectedString(b))
	*data = script
	return nil
}

func (c LuaCodec) Encode(w io.Writer, data rbxmk.Data) (err error) {
	var script *rbxfile.Instance
	switch v := data.(type) {
	case []*rbxfile.Instance:
		if len(v) > 0 {
			script = v[0]
		}
	case *rbxfile.Instance:
		script = v
	}
	if script != nil {
		switch script.ClassName {
		case "Script", "LocalScript", "ModuleScript":
			if source, ok := script.Properties["Source"]; ok {
				data = source
			}
		}
	}
	switch v := data.(type) {
	case rbxfile.ValueProtectedString:
		_, err = w.Write([]byte(v))
	case rbxfile.ValueBinaryString:
		_, err = w.Write([]byte(v))
	case rbxfile.ValueString:
		_, err = w.Write([]byte(v))
	default:
		err = NewDataTypeError(data)
	}
	return err
}
