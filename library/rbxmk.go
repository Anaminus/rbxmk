package library

import (
	"bytes"
	"path/filepath"
	"strings"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/reflect"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

func init() { register(RBXMK) }

var RBXMK = rbxmk.Library{
	Name:       "rbxmk",
	ImportedAs: "rbxmk",
	Priority:   1,
	Open:       openRBXMK,
	Dump:       dumpRBXMK,
	Types: []func() rbxmk.Reflector{
		reflect.AttrConfig,
		reflect.Enums,
		reflect.FormatSelector,
		reflect.Nil,
		reflect.RootDesc,
		reflect.String,
		reflect.Symbol,
		reflect.Table,
	},
}

func openRBXMK(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 8)
	lib.RawSetString("decodeFormat", s.WrapFunc(rbxmkDecodeFormat))
	lib.RawSetString("encodeFormat", s.WrapFunc(rbxmkEncodeFormat))
	lib.RawSetString("formatCanDecode", s.WrapFunc(rbxmkFormatCanDecode))
	lib.RawSetString("loadFile", s.WrapFunc(rbxmkLoadFile))
	lib.RawSetString("loadString", s.WrapFunc(rbxmkLoadString))
	lib.RawSetString("runFile", s.WrapFunc(rbxmkRunFile))
	lib.RawSetString("runString", s.WrapFunc(rbxmkRunString))

	lv, _ := s.MustReflector("Enums").PushTo(s.Context(), s.Enums())
	lib.RawSetString("Enum", lv)

	mt := s.L.CreateTable(0, 2)
	mt.RawSetString("__index", s.WrapOperator(rbxmkOpIndex))
	mt.RawSetString("__newindex", s.WrapOperator(rbxmkOpNewindex))
	s.L.SetMetatable(lib, mt)

	return lib
}

func rbxmkDecodeFormat(s rbxmk.State) int {
	selector := s.Pull(1, "FormatSelector").(rtypes.FormatSelector)
	format := s.Format(selector.Format)
	if format.Name == "" {
		return s.RaiseError("unknown format %q", selector.Format)
	}
	if format.Decode == nil {
		return s.RaiseError("cannot decode with format %s", format.Name)
	}
	r := strings.NewReader(string(s.Pull(2, "BinaryString").(types.BinaryString)))
	v, err := format.Decode(s.Global, selector, r)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	return s.Push(v)
}

func rbxmkEncodeFormat(s rbxmk.State) int {
	selector := s.Pull(1, "FormatSelector").(rtypes.FormatSelector)
	format := s.Format(selector.Format)
	if format.Name == "" {
		return s.RaiseError("unknown format %q", selector.Format)
	}
	if format.Encode == nil {
		return s.RaiseError("cannot encode with format %s", format.Name)
	}
	value := s.PullEncodedFormat(2, format)
	var w bytes.Buffer
	if err := format.Encode(s.Global, selector, &w, value); err != nil {
		return s.RaiseError("%s", err)
	}
	return s.Push(types.BinaryString(w.Bytes()))
}

func rbxmkFormatCanDecode(s rbxmk.State) int {
	selector := s.Pull(1, "FormatSelector").(rtypes.FormatSelector)
	typeName := string(s.Pull(2, "string").(types.String))
	format := s.Format(selector.Format)
	if format.Name == "" {
		return s.RaiseError("unknown format %q", selector.Format)
	}
	if format.CanDecode == nil {
		return s.RaiseError("undefined decode type for %s", format.Name)
	}
	return s.Push(types.Bool(format.CanDecode(s.Global, selector, typeName)))
}

func rbxmkLoadFile(s rbxmk.State) int {
	fileName := filepath.Clean(s.CheckString(1))
	fn, err := s.L.LoadFile(fileName)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	s.L.Push(fn)
	return 1
}

func rbxmkLoadString(s rbxmk.State) int {
	source := s.CheckString(1)
	fn, err := s.L.LoadString(source)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	s.L.Push(fn)
	return 1
}

func rbxmkRunFile(s rbxmk.State) int {
	fileName := filepath.Clean(s.CheckString(1))
	fi, err := s.FS.Stat(fileName)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	if err = s.PushFile(rbxmk.FileEntry{Path: fileName, FileInfo: fi}); err != nil {
		return s.RaiseError("%s", err)
	}

	nt := s.Count()

	// Load file as function.
	fn, err := s.L.LoadFile(fileName)
	if err != nil {
		s.PopFile()
		return s.RaiseError("%s", err)
	}
	s.L.Push(fn) // +function

	// Push extra arguments as arguments to loaded function.
	for i := 2; i <= nt; i++ {
		s.L.Push(s.L.Get(i)) // function, ..., +arg
	}
	// function, +args...

	// Call loaded function.
	err = s.L.PCall(nt-1, lua.MultRet, nil) // -function, -args..., +returns...
	s.PopFile()
	if err != nil {
		return s.RaiseError("%s", err)
	}
	return s.Count() - nt
}

func rbxmkRunString(s rbxmk.State) int {
	source := s.CheckString(1)
	nt := s.Count()

	// Load file as function.
	fn, err := s.L.LoadString(source)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	s.L.Push(fn) // +function

	// Push extra arguments as arguments to loaded function.
	for i := 2; i <= nt; i++ {
		s.L.Push(s.L.Get(i)) // function, ..., +arg
	}
	// function, +args...

	// Call loaded function.
	err = s.L.PCall(nt-1, lua.MultRet, nil) // -function, -args..., +returns...
	s.PopFile()
	if err != nil {
		return s.RaiseError("%s", err)
	}
	return s.Count() - nt
}

func rbxmkOpIndex(s rbxmk.State) int {
	switch field := s.Pull(2, "string").(types.String); field {
	case "globalDesc":
		desc := s.Desc.Of(nil)
		if desc == nil {
			return s.Push(rtypes.Nil)
		}
		return s.Push(desc)
	case "globalAttrConfig":
		attrcfg := s.AttrConfig.Of(nil)
		if attrcfg == nil {
			return s.Push(rtypes.Nil)
		}
		return s.Push(attrcfg)
	default:
		return s.RaiseError("unknown field %q", field)
	}
}

func rbxmkOpNewindex(s rbxmk.State) int {
	switch field := s.Pull(2, "string").(types.String); field {
	case "globalDesc":
		desc, _ := s.PullOpt(3, nil, "RootDesc").(*rtypes.RootDesc)
		if desc == nil {
			s.Desc = nil
		}
		s.Desc = desc
		return 0
	case "globalAttrConfig":
		attrcfg, _ := s.PullOpt(3, nil, "AttrConfig").(*rtypes.AttrConfig)
		if attrcfg == nil {
			s.AttrConfig = nil
		}
		s.AttrConfig = attrcfg
		return 0
	default:
		return s.RaiseError("unknown field %q", field)
	}
}

func dumpRBXMK(s rbxmk.State) dump.Library {
	lib := dump.Library{
		Struct: dump.Struct{
			Fields: dump.Fields{
				"Enum": dump.Property{
					ValueType:   dt.Prim("Enums"),
					Summary:     "Libraries/rbxmk:Fields/Enum/Summary",
					Description: "Libraries/rbxmk:Fields/Enum/Description",
				},
				"decodeFormat": dump.Function{
					Parameters: dump.Parameters{
						{Name: "format", Type: dt.Prim("string")},
						{Name: "bytes", Type: dt.Prim("BinaryString")},
					},
					Returns: dump.Parameters{
						{Name: "value", Type: dt.Prim("any")},
					},
					CanError:    true,
					Summary:     "Libraries/rbxmk:Fields/decodeFormat/Summary",
					Description: "Libraries/rbxmk:Fields/decodeFormat/Description",
				},
				"encodeFormat": dump.Function{
					Parameters: dump.Parameters{
						{Name: "format", Type: dt.Prim("string")},
						{Name: "value", Type: dt.Prim("any")},
					},
					Returns: dump.Parameters{
						{Name: "bytes", Type: dt.Prim("BinaryString")},
					},
					CanError:    true,
					Summary:     "Libraries/rbxmk:Fields/encodeFormat/Summary",
					Description: "Libraries/rbxmk:Fields/encodeFormat/Description",
				},
				"formatCanDecode": dump.Function{
					Parameters: dump.Parameters{
						{Name: "format", Type: dt.Prim("string")},
						{Name: "type", Type: dt.Prim("string")},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim("bool")},
					},
					CanError:    true,
					Summary:     "Libraries/rbxmk:Fields/formatCanDecode/Summary",
					Description: "Libraries/rbxmk:Fields/formatCanDecode/Description",
				},
				"globalAttrConfig": dump.Property{
					ValueType:   dt.Optional{T: dt.Prim("AttrConfig")},
					Summary:     "Libraries/rbxmk:Fields/globalAttrConfig/Summary",
					Description: "Libraries/rbxmk:Fields/globalAttrConfig/Description",
				},
				"globalDesc": dump.Property{
					ValueType:   dt.Optional{T: dt.Prim("RootDesc")},
					Summary:     "Libraries/rbxmk:Fields/globalDesc/Summary",
					Description: "Libraries/rbxmk:Fields/globalDesc/Description",
				},
				"loadFile": dump.Function{
					Parameters: dump.Parameters{
						{Name: "path", Type: dt.Prim("string")},
					},
					Returns: dump.Parameters{
						{Name: "func", Type: dt.Prim("function")},
					},
					CanError:    true,
					Summary:     "Libraries/rbxmk:Fields/loadFile/Summary",
					Description: "Libraries/rbxmk:Fields/loadFile/Description",
				},
				"loadString": dump.Function{
					Parameters: dump.Parameters{
						{Name: "source", Type: dt.Prim("string")},
					},
					Returns: dump.Parameters{
						{Name: "func", Type: dt.Prim("function")},
					},
					CanError:    true,
					Summary:     "Libraries/rbxmk:Fields/loadString/Summary",
					Description: "Libraries/rbxmk:Fields/loadString/Description",
				},
				"runFile": dump.Function{
					Parameters: dump.Parameters{
						{Name: "path", Type: dt.Prim("string")},
						{Name: "...", Type: dt.Prim("any")},
					},
					Returns: dump.Parameters{
						{Name: "...", Type: dt.Prim("any")},
					},
					CanError:    true,
					Summary:     "Libraries/rbxmk:Fields/runFile/Summary",
					Description: "Libraries/rbxmk:Fields/runFile/Description",
				},
				"runString": dump.Function{
					Parameters: dump.Parameters{
						{Name: "source", Type: dt.Prim("string")},
						{Name: "...", Type: dt.Prim("any")},
					},
					Returns: dump.Parameters{
						{Name: "...", Type: dt.Prim("any")},
					},
					CanError:    true,
					Summary:     "Libraries/rbxmk:Fields/runString/Summary",
					Description: "Libraries/rbxmk:Fields/runString/Description",
				},
			},
			Summary:     "Libraries/rbxmk:Summary",
			Description: "Libraries/rbxmk:Description",
		},
	}
	return lib
}
