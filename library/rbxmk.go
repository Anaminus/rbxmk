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
		reflect.Desc,
		reflect.Enums,
		reflect.FormatSelector,
		reflect.Instance,
		reflect.Nil,
		reflect.String,
		reflect.Symbol,
		reflect.Table,
	},
}

func openRBXMK(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 11)
	lib.RawSetString("decodeFormat", s.WrapFunc(rbxmkDecodeFormat))
	lib.RawSetString("encodeFormat", s.WrapFunc(rbxmkEncodeFormat))
	lib.RawSetString("formatCanDecode", s.WrapFunc(rbxmkFormatCanDecode))
	lib.RawSetString("get", s.WrapFunc(rbxmkGet))
	lib.RawSetString("loadFile", s.WrapFunc(rbxmkLoadFile))
	lib.RawSetString("loadString", s.WrapFunc(rbxmkLoadString))
	lib.RawSetString("propType", s.WrapFunc(rbxmkPropType))
	lib.RawSetString("runFile", s.WrapFunc(rbxmkRunFile))
	lib.RawSetString("runString", s.WrapFunc(rbxmkRunString))
	lib.RawSetString("set", s.WrapFunc(rbxmkSet))

	lv, _ := s.MustReflector(rtypes.T_Enums).PushTo(s.Context(), s.Enums())
	lib.RawSetString("Enum", lv)

	mt := s.L.CreateTable(0, 2)
	mt.RawSetString("__index", s.WrapOperator(rbxmkOpIndex))
	mt.RawSetString("__newindex", s.WrapOperator(rbxmkOpNewindex))
	s.L.SetMetatable(lib, mt)

	return lib
}

func rbxmkDecodeFormat(s rbxmk.State) int {
	selector := s.Pull(1, rtypes.T_FormatSelector).(rtypes.FormatSelector)
	format := s.Format(selector.Format)
	if format.Name == "" {
		return s.RaiseError("unknown format %q", selector.Format)
	}
	if format.Decode == nil {
		return s.RaiseError("cannot decode with format %s", format.Name)
	}
	r := strings.NewReader(string(s.Pull(2, rtypes.T_BinaryString).(types.BinaryString)))
	v, err := format.Decode(s.Global, selector, r)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	return s.Push(v)
}

func rbxmkEncodeFormat(s rbxmk.State) int {
	selector := s.Pull(1, rtypes.T_FormatSelector).(rtypes.FormatSelector)
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
	selector := s.Pull(1, rtypes.T_FormatSelector).(rtypes.FormatSelector)
	typeName := string(s.Pull(2, rtypes.T_String).(types.String))
	format := s.Format(selector.Format)
	if format.Name == "" {
		return s.RaiseError("unknown format %q", selector.Format)
	}
	if format.CanDecode == nil {
		return s.RaiseError("undefined decode type for %s", format.Name)
	}
	return s.Push(types.Bool(format.CanDecode(s.Global, selector, typeName)))
}

func fallbacksFromArg(s rbxmk.State, n int, instance *rtypes.Instance) (desc *rtypes.Desc, rfl rbxmk.Reflector) {
	var fallbackDesc *rtypes.Desc
	switch v := s.PullAnyOf(n, rtypes.T_String, rtypes.T_Desc, rtypes.T_Nil).(type) {
	case types.String:
		if rfl = s.Reflector(string(v)); rfl.Name == "" {
			s.RaiseError("unknown type %q", string(v))
			return desc, rfl
		}
	case *rtypes.Desc:
		fallbackDesc = v
	case rtypes.NilType:
	default:
		s.ReflectorError(n)
		return desc, rfl
	}
	if desc = s.Desc.Of(instance); desc == nil {
		desc = fallbackDesc
	}
	return desc, rfl
}

func rbxmkGet(s rbxmk.State) int {
	instance := s.Pull(1, rtypes.T_Instance).(*rtypes.Instance)
	property := string(s.Pull(2, rtypes.T_String).(types.String))
	desc, rfl := fallbacksFromArg(s, 3, instance)
	lv, err := reflect.GetProperty(s, instance, property, desc, rfl)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	s.L.Push(lv)
	return 1
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

func rbxmkPropType(s rbxmk.State) int {
	instance := s.Pull(1, rtypes.T_Instance).(*rtypes.Instance)
	property := string(s.Pull(2, rtypes.T_String).(types.String))
	if value := instance.Get(property); value != nil {
		return s.Push(types.String(value.Type()))
	}
	return s.Push(rtypes.Nil)
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

func rbxmkSet(s rbxmk.State) int {
	instance := s.Pull(1, rtypes.T_Instance).(*rtypes.Instance)
	property := string(s.Pull(2, rtypes.T_String).(types.String))
	value := s.CheckAny(3)
	desc, rfl := fallbacksFromArg(s, 4, instance)
	err := reflect.SetProperty(s, instance, property, value, desc, rfl)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	return 0
}

func rbxmkOpIndex(s rbxmk.State) int {
	switch field := s.Pull(2, rtypes.T_String).(types.String); field {
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
	switch field := s.Pull(2, rtypes.T_String).(types.String); field {
	case "globalDesc":
		desc, _ := s.PullOpt(3, nil, rtypes.T_Desc).(*rtypes.Desc)
		if desc == nil {
			s.Desc = nil
		}
		s.Desc = desc
		return 0
	case "globalAttrConfig":
		attrcfg, _ := s.PullOpt(3, nil, rtypes.T_AttrConfig).(*rtypes.AttrConfig)
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
					ValueType:   dt.Prim(rtypes.T_Enums),
					Summary:     "Libraries/rbxmk:Fields/Enum/Summary",
					Description: "Libraries/rbxmk:Fields/Enum/Description",
				},
				"decodeFormat": dump.Function{
					Parameters: dump.Parameters{
						{Name: "format", Type: dt.Prim(rtypes.T_String)},
						{Name: "bytes", Type: dt.Prim(rtypes.T_BinaryString)},
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
						{Name: "format", Type: dt.Prim(rtypes.T_String)},
						{Name: "value", Type: dt.Prim("any")},
					},
					Returns: dump.Parameters{
						{Name: "bytes", Type: dt.Prim(rtypes.T_BinaryString)},
					},
					CanError:    true,
					Summary:     "Libraries/rbxmk:Fields/encodeFormat/Summary",
					Description: "Libraries/rbxmk:Fields/encodeFormat/Description",
				},
				"formatCanDecode": dump.Function{
					Parameters: dump.Parameters{
						{Name: "format", Type: dt.Prim(rtypes.T_String)},
						{Name: "type", Type: dt.Prim(rtypes.T_String)},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim(rtypes.T_Bool)},
					},
					CanError:    true,
					Summary:     "Libraries/rbxmk:Fields/formatCanDecode/Summary",
					Description: "Libraries/rbxmk:Fields/formatCanDecode/Description",
				},
				"get": dump.Function{
					Parameters: dump.Parameters{
						{Name: "instance", Type: dt.Prim(rtypes.T_Instance)},
						{Name: "property", Type: dt.Prim(rtypes.T_String)},
						{Name: "fallback", Type: dt.Or{
							dt.Prim(rtypes.T_String),
							dt.Prim(rtypes.T_Desc),
							dt.Prim(rtypes.T_Nil),
						}},
					},
					Returns: dump.Parameters{
						{Name: "value", Type: dt.Prim("any")},
					},
					CanError:    true,
					Summary:     "Libraries/rbxmk:Fields/get/Summary",
					Description: "Libraries/rbxmk:Fields/get/Description",
				},
				"globalAttrConfig": dump.Property{
					ValueType:   dt.Optional{T: dt.Prim(rtypes.T_AttrConfig)},
					Summary:     "Libraries/rbxmk:Fields/globalAttrConfig/Summary",
					Description: "Libraries/rbxmk:Fields/globalAttrConfig/Description",
				},
				"globalDesc": dump.Property{
					ValueType:   dt.Optional{T: dt.Prim(rtypes.T_Desc)},
					Summary:     "Libraries/rbxmk:Fields/globalDesc/Summary",
					Description: "Libraries/rbxmk:Fields/globalDesc/Description",
				},
				"loadFile": dump.Function{
					Parameters: dump.Parameters{
						{Name: "path", Type: dt.Prim(rtypes.T_String)},
					},
					Returns: dump.Parameters{
						{Name: "func", Type: dt.Prim(rtypes.T_LuaFunction)},
					},
					CanError:    true,
					Summary:     "Libraries/rbxmk:Fields/loadFile/Summary",
					Description: "Libraries/rbxmk:Fields/loadFile/Description",
				},
				"loadString": dump.Function{
					Parameters: dump.Parameters{
						{Name: "source", Type: dt.Prim(rtypes.T_String)},
					},
					Returns: dump.Parameters{
						{Name: "func", Type: dt.Prim(rtypes.T_LuaFunction)},
					},
					CanError:    true,
					Summary:     "Libraries/rbxmk:Fields/loadString/Summary",
					Description: "Libraries/rbxmk:Fields/loadString/Description",
				},
				"propType": dump.Function{
					Parameters: dump.Parameters{
						{Name: "instance", Type: dt.Prim(rtypes.T_Instance)},
						{Name: "property", Type: dt.Prim(rtypes.T_String)},
					},
					Returns: dump.Parameters{
						{Name: "type", Type: dt.Optional{T: dt.Prim(rtypes.T_String)}},
					},
					CanError:    true,
					Summary:     "Libraries/rbxmk:Fields/propType/Summary",
					Description: "Libraries/rbxmk:Fields/propType/Description",
				},
				"runFile": dump.Function{
					Parameters: dump.Parameters{
						{Name: "path", Type: dt.Prim(rtypes.T_String)},
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
						{Name: "source", Type: dt.Prim(rtypes.T_String)},
						{Name: "...", Type: dt.Prim("any")},
					},
					Returns: dump.Parameters{
						{Name: "...", Type: dt.Prim("any")},
					},
					CanError:    true,
					Summary:     "Libraries/rbxmk:Fields/runString/Summary",
					Description: "Libraries/rbxmk:Fields/runString/Description",
				},
				"set": dump.Function{
					Parameters: dump.Parameters{
						{Name: "instance", Type: dt.Prim(rtypes.T_Instance)},
						{Name: "property", Type: dt.Prim(rtypes.T_String)},
						{Name: "value", Type: dt.Prim("any")},
						{Name: "fallback", Type: dt.Or{
							dt.Prim(rtypes.T_String),
							dt.Prim(rtypes.T_Desc),
							dt.Prim(rtypes.T_Nil),
						}},
					},
					CanError:    true,
					Summary:     "Libraries/rbxmk:Fields/set/Summary",
					Description: "Libraries/rbxmk:Fields/set/Description",
				},
			},
			Summary:     "Libraries/rbxmk:Summary",
			Description: "Libraries/rbxmk:Description",
		},
	}
	return lib
}
