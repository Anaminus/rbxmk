package library

import (
	"os"
	"path/filepath"

	"github.com/anaminus/rbxmk"
	reflect "github.com/anaminus/rbxmk/library/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/rbxdump/diff"
	"github.com/robloxapi/types"
	lua "github.com/yuin/gopher-lua"
)

func init() { register(RBXMK, 0) }

var RBXMK = rbxmk.Library{
	Name: "rbxmk",
	Open: func(s rbxmk.State) *lua.LTable {
		lib := s.L.CreateTable(0, 11)
		lib.RawSetString("loadFile", s.WrapFunc(rbxmkLoadFile))
		lib.RawSetString("loadString", s.WrapFunc(rbxmkLoadString))
		lib.RawSetString("runFile", s.WrapFunc(rbxmkRunFile))
		lib.RawSetString("runString", s.WrapFunc(rbxmkRunString))
		lib.RawSetString("newDesc", s.WrapFunc(rbxmkNewDesc))
		lib.RawSetString("diffDesc", s.WrapFunc(rbxmkDiffDesc))
		lib.RawSetString("patchDesc", s.WrapFunc(rbxmkPatchDesc))
		lib.RawSetString("encodeFormat", s.WrapFunc(rbxmkEncodeFormat))
		lib.RawSetString("decodeFormat", s.WrapFunc(rbxmkDecodeFormat))
		lib.RawSetString("readSource", s.WrapFunc(rbxmkReadSource))
		lib.RawSetString("writeSource", s.WrapFunc(rbxmkWriteSource))

		for _, f := range reflect.All() {
			r := f()
			s.RegisterReflector(r)
			s.ApplyReflector(r, lib)
		}

		mt := s.L.CreateTable(0, 2)
		mt.RawSetString("__index", s.WrapFunc(func(s rbxmk.State) int {
			if field := s.Pull(2, "string").(types.String); field != "globalDesc" {
				return s.RaiseError("unknown field %q", field)
			}
			desc := s.Desc(nil)
			if desc == nil {
				return s.Push(rtypes.Nil)
			}
			return s.Push(desc)
		}))
		mt.RawSetString("__newindex", s.WrapFunc(func(s rbxmk.State) int {
			if field := s.Pull(2, "string").(types.String); field != "globalDesc" {
				return s.RaiseError("unknown field %q", field)
			}
			desc, _ := s.PullOpt(3, "RootDesc", nil).(*rtypes.RootDesc)
			if desc == nil {
				s.SetDesc(nil)
			}
			s.SetDesc(desc)
			return 0
		}))
		s.L.SetMetatable(lib, mt)

		return lib
	},
}

func rbxmkLoadFile(s rbxmk.State) int {
	fileName := filepath.Clean(s.CheckString(1))
	fn, err := s.L.LoadFile(fileName)
	if err != nil {
		return s.RaiseError(err.Error())
	}
	s.L.Push(fn)
	return 1
}

func rbxmkLoadString(s rbxmk.State) int {
	source := s.CheckString(1)
	fn, err := s.L.LoadString(source)
	if err != nil {
		return s.RaiseError(err.Error())
	}
	s.L.Push(fn)
	return 1
}

func rbxmkRunFile(s rbxmk.State) int {
	fileName := filepath.Clean(s.CheckString(1))
	fi, err := os.Stat(fileName)
	if err != nil {
		return s.RaiseError(err.Error())
	}
	if err = s.PushFile(rbxmk.FileInfo{Path: fileName, FileInfo: fi}); err != nil {
		return s.RaiseError(err.Error())
	}

	nt := s.L.GetTop()

	// Load file as function.
	fn, err := s.L.LoadFile(fileName)
	if err != nil {
		s.PopFile()
		return s.RaiseError(err.Error())
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
		return s.RaiseError(err.Error())
	}
	return s.L.GetTop() - nt
}

func rbxmkRunString(s rbxmk.State) int {
	source := s.CheckString(1)
	nt := s.L.GetTop()

	// Load file as function.
	fn, err := s.L.LoadString(source)
	if err != nil {
		return s.RaiseError(err.Error())
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
		return s.RaiseError(err.Error())
	}
	return s.L.GetTop() - nt
}

func rbxmkNewDesc(s rbxmk.State) int {
	switch name := string(s.Pull(1, "string").(types.String)); name {
	case "RootDesc":
		return s.Push(&rtypes.RootDesc{Root: &rbxdump.Root{
			Classes: make(map[string]*rbxdump.Class),
			Enums:   make(map[string]*rbxdump.Enum),
		}})
	case "ClassDesc":
		return s.Push(rtypes.ClassDesc{Class: &rbxdump.Class{
			Members: make(map[string]rbxdump.Member),
		}})
	case "PropertyDesc":
		return s.Push(rtypes.PropertyDesc{Property: &rbxdump.Property{
			ReadSecurity:  "None",
			WriteSecurity: "None",
		}})
	case "FunctionDesc":
		return s.Push(rtypes.FunctionDesc{Function: &rbxdump.Function{
			Security: "None",
		}})
	case "EventDesc":
		return s.Push(rtypes.EventDesc{Event: &rbxdump.Event{
			Security: "None",
		}})
	case "CallbackDesc":
		return s.Push(rtypes.CallbackDesc{Callback: &rbxdump.Callback{
			Security: "None",
		}})
	case "ParameterDesc":
		var param rbxdump.Parameter
		param.Type = s.PullOpt(2, "TypeDesc", rtypes.TypeDesc{}).(rtypes.TypeDesc).Embedded
		param.Name = string(s.PullOpt(3, "string", types.String("")).(types.String))
		switch def := s.PullOpt(4, "string", rtypes.Nil).(type) {
		case rtypes.NilType:
			param.Optional = false
		case types.String:
			param.Optional = true
			param.Default = string(def)
		}
		return s.Push(rtypes.ParameterDesc{Parameter: param})
	case "TypeDesc":
		category := string(s.PullOpt(2, "string", types.String("")).(types.String))
		name := string(s.PullOpt(3, "string", types.String("")).(types.String))
		return s.Push(rtypes.TypeDesc{Embedded: rbxdump.Type{
			Category: category,
			Name:     name,
		}})
	case "EnumDesc":
		return s.Push(rtypes.EnumDesc{Enum: &rbxdump.Enum{
			Items: make(map[string]*rbxdump.EnumItem),
		}})
	case "EnumItemDesc":
		return s.Push(rtypes.EnumItemDesc{EnumItem: &rbxdump.EnumItem{}})
	default:
		return s.RaiseError("unable to create descriptor of type %q", name)
	}
}

func rbxmkDiffDesc(s rbxmk.State) int {
	var prev *rbxdump.Root
	var next *rbxdump.Root
	switch v := s.PullAnyOf(1, "RootDesc", "nil").(type) {
	case rtypes.NilType:
	case *rtypes.RootDesc:
		prev = v.Root
	}
	switch v := s.PullAnyOf(2, "RootDesc", "nil").(type) {
	case rtypes.NilType:
	case *rtypes.RootDesc:
		next = v.Root
	}
	actions := diff.Diff{Prev: prev, Next: next}.Diff()
	descActions := make(rtypes.DescActions, len(actions))
	for i, action := range actions {
		descActions[i] = &rtypes.DescAction{Action: action}
	}
	return s.Push(descActions)
}

func rbxmkPatchDesc(s rbxmk.State) int {
	desc := s.Pull(1, "RootDesc").(*rtypes.RootDesc).Root
	descActions := s.Pull(2, "DescActions").(rtypes.DescActions)
	actions := make([]diff.Action, len(descActions))
	for i, action := range descActions {
		actions[i] = action.Action
	}
	diff.Patch{Root: desc}.Patch(actions)
	return 0
}

func rbxmkEncodeFormat(s rbxmk.State) int {
	name := string(s.Pull(1, "string").(types.String))
	format := s.Format(name)
	if format.Name == "" {
		return s.RaiseError("unknown format %q", name)
	}
	if format.Encode == nil {
		return s.RaiseError("cannot encode with format %s", name)
	}
	b, err := format.Encode(s.Pull(2, "Variant"))
	if err != nil {
		return s.RaiseError(err.Error())
	}
	return s.Push(types.BinaryString(b))
}

func rbxmkDecodeFormat(s rbxmk.State) int {
	name := string(s.Pull(1, "string").(types.String))
	format := s.Format(name)
	if format.Name == "" {
		return s.RaiseError("unknown format %q", name)
	}
	if format.Decode == nil {
		return s.RaiseError("cannot decode with format %s", name)
	}
	v, err := format.Decode([]byte(s.Pull(2, "BinaryString").(types.BinaryString)))
	if err != nil {
		return s.RaiseError(err.Error())
	}
	return s.Push(v)
}

func rbxmkReadSource(s rbxmk.State) int {
	name := string(s.Pull(1, "string").(types.String))
	source := s.Source(name)
	if source.Name == "" {
		return s.RaiseError("unknown source %q", name)
	}
	if source.Read == nil {
		return s.RaiseError("cannot read with format %s", name)
	}
	options := make([]interface{}, s.L.GetTop()-1)
	for i := 2; i <= s.L.GetTop(); i++ {
		options[i-2] = s.Pull(i, "Variant")
	}
	b, err := source.Read(options...)
	if err != nil {
		return s.RaiseError(err.Error())
	}
	return s.Push(types.BinaryString(b))
}

func rbxmkWriteSource(s rbxmk.State) int {
	name := string(s.Pull(1, "string").(types.String))
	source := s.Source(name)
	if source.Name == "" {
		return s.RaiseError("unknown source %q", name)
	}
	if source.Write == nil {
		return s.RaiseError("cannot write with format %s", name)
	}
	b := []byte(s.Pull(2, "BinaryString").(types.BinaryString))
	options := make([]interface{}, s.L.GetTop()-2)
	for i := 3; i <= s.L.GetTop(); i++ {
		options[i-3] = s.Pull(i, "Variant")
	}
	if err := source.Write(b, options...); err != nil {
		return s.RaiseError(err.Error())
	}
	return 0
}
