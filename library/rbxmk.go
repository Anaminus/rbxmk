package library

import (
	"os"
	"path/filepath"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

var RBXMK = rbxmk.Library{
	Name: "rbxmk",
	Open: func(s rbxmk.State) *lua.LTable {
		lib := s.L.CreateTable(0, 7)
		lib.RawSetString("load", s.WrapFunc(rbxmkLoad))
		lib.RawSetString("meta", s.WrapFunc(rbxmkMeta))
		lib.RawSetString("newDesc", s.WrapFunc(rbxmkNewDesc))
		lib.RawSetString("encodeFormat", s.WrapFunc(rbxmkEncodeFormat))
		lib.RawSetString("decodeFormat", s.WrapFunc(rbxmkDecodeFormat))
		lib.RawSetString("readSource", s.WrapFunc(rbxmkReadSource))
		lib.RawSetString("writeSource", s.WrapFunc(rbxmkWriteSource))
		return lib
	},
}

func rbxmkLoad(s rbxmk.State) int {
	fileName := filepath.Clean(s.L.CheckString(1))
	fi, err := os.Stat(fileName)
	if err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}
	if err = s.PushFile(rbxmk.FileInfo{Path: fileName, FileInfo: fi}); err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}

	// Load file as function.
	fn, err := s.L.LoadFile(fileName)
	if err != nil {
		s.PopFile()
		s.L.RaiseError(err.Error())
		return 0
	}
	s.L.Push(fn) // +function

	// Push extra arguments as arguments to loaded function.
	nt := s.L.GetTop()
	for i := 2; i <= nt; i++ {
		s.L.Push(s.L.Get(i)) // function, ..., +arg
	}
	// function, +args...

	// Call loaded function.
	err = s.L.PCall(nt-1, lua.MultRet, nil) // -function, -args..., +returns...
	s.PopFile()
	if err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}
	return s.L.GetTop() - 1
}

func metaGet(s rbxmk.State, inst *rtypes.Instance, name string) int {
	switch name {
	case "Reference":
		return s.Push(types.String(inst.Reference))
	case "IsService":
		return s.Push(types.Bool(inst.IsService))
	default:
		s.L.RaiseError("unknown metadata %q", name)
	}
	return 0
}

func metaSet(s rbxmk.State, inst *rtypes.Instance, name string) int {
	switch name {
	case "Reference":
		inst.Reference = string(s.Pull(3, "string").(types.String))
	case "IsService":
		inst.IsService = bool(s.Pull(3, "bool").(types.Bool))
	default:
		s.L.RaiseError("unknown metadata %q", name)
	}
	return 0
}

func rbxmkMeta(s rbxmk.State) int {
	inst := s.Pull(1, "Instance").(*rtypes.Instance)
	name := string(s.Pull(2, "string").(types.String))
	if s.Count() == 3 {
		return metaSet(s, inst, name)
	}
	return metaGet(s, inst, name)
}

func rbxmkNewDesc(s rbxmk.State) int {
	switch name := string(s.Pull(1, "string").(types.String)); name {
	case "Root":
		return s.Push(rtypes.RootDesc{})
	case "Class":
		return s.Push(rtypes.ClassDesc{})
	case "Property":
		return s.Push(rtypes.PropertyDesc{})
	case "Function":
		return s.Push(rtypes.FunctionDesc{})
	case "Event":
		return s.Push(rtypes.EventDesc{})
	case "Callback":
		return s.Push(rtypes.CallbackDesc{})
	case "Parameter":
		return s.Push(rtypes.ParameterDesc{})
	case "Type":
		return s.Push(rtypes.TypeDesc{})
	case "Enum":
		return s.Push(rtypes.EnumDesc{})
	case "EnumItem":
		return s.Push(rtypes.EnumItemDesc{})
	default:
		s.L.RaiseError("unable to create descriptor of type %q", name)
		return 0
	}
}

func rbxmkEncodeFormat(s rbxmk.State) int {
	name := string(s.Pull(1, "string").(types.String))
	format := s.Format(name)
	if format.Name == "" {
		s.L.RaiseError("unknown format %q", name)
		return 0
	}
	if format.Encode == nil {
		s.L.RaiseError("cannot encode with format %s", name)
		return 0
	}
	b, err := format.Encode(s.Pull(2, "Variant"))
	if err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}
	return s.Push(types.BinaryString(b))
}

func rbxmkDecodeFormat(s rbxmk.State) int {
	name := string(s.Pull(1, "string").(types.String))
	format := s.Format(name)
	if format.Name == "" {
		s.L.RaiseError("unknown format %q", name)
		return 0
	}
	if format.Decode == nil {
		s.L.RaiseError("cannot decode with format %s", name)
		return 0
	}
	v, err := format.Decode([]byte(s.Pull(2, "BinaryString").(types.BinaryString)))
	if err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}
	return s.Push(v)
}

func rbxmkReadSource(s rbxmk.State) int {
	name := string(s.Pull(1, "string").(types.String))
	source := s.Source(name)
	if source.Name == "" {
		s.L.RaiseError("unknown source %q", name)
		return 0
	}
	if source.Read == nil {
		s.L.RaiseError("cannot read with format %s", name)
		return 0
	}
	options := make([]interface{}, s.L.GetTop()-1)
	for i := 2; i <= s.L.GetTop(); i++ {
		options[i-2] = s.Pull(i, "Variant")
	}
	b, err := source.Read(options...)
	if err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}
	return s.Push(types.BinaryString(b))
}

func rbxmkWriteSource(s rbxmk.State) int {
	name := string(s.Pull(1, "string").(types.String))
	source := s.Source(name)
	if source.Name == "" {
		s.L.RaiseError("unknown source %q", name)
		return 0
	}
	if source.Write == nil {
		s.L.RaiseError("cannot write with format %s", name)
		return 0
	}
	b := []byte(s.Pull(2, "BinaryString").(types.BinaryString))
	options := make([]interface{}, s.L.GetTop()-2)
	for i := 3; i <= s.L.GetTop(); i++ {
		options[i-3] = s.Pull(i, "Variant")
	}
	if err := source.Write(b, options...); err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}
	return 0
}
