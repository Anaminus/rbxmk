package library

import (
	"bytes"
	"net/http"
	"path/filepath"
	"strings"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
	"github.com/anaminus/rbxmk/reflect"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/rbxdump/diff"
	"github.com/robloxapi/types"
)

func init() { register(RBXMK, 0) }

var RBXMK = rbxmk.Library{
	Name:       "rbxmk",
	ImportedAs: "rbxmk",
	Open:       openRBXMK,
	Dump:       dumpRBXMK,
	Types: []func() rbxmk.Reflector{
		reflect.AttrConfig,
		reflect.CallbackDesc,
		reflect.ClassDesc,
		reflect.Cookie,
		reflect.Cookies,
		reflect.DescAction,
		reflect.DescActions,
		reflect.EnumDesc,
		reflect.EnumItemDesc,
		reflect.EventDesc,
		reflect.FormatSelector,
		reflect.FunctionDesc,
		reflect.Nil,
		reflect.ParameterDesc,
		reflect.PropertyDesc,
		reflect.RootDesc,
		reflect.String,
		reflect.Symbol,
		reflect.Table,
		reflect.TypeDesc,
	},
}

func openRBXMK(s rbxmk.State) *lua.LTable {
	lib := s.L.CreateTable(0, 13)
	lib.RawSetString("cookiesFrom", s.WrapFunc(rbxmkCookiesFrom))
	lib.RawSetString("decodeFormat", s.WrapFunc(rbxmkDecodeFormat))
	lib.RawSetString("diffDesc", s.WrapFunc(rbxmkDiffDesc))
	lib.RawSetString("encodeFormat", s.WrapFunc(rbxmkEncodeFormat))
	lib.RawSetString("formatCanDecode", s.WrapFunc(rbxmkFormatCanDecode))
	lib.RawSetString("loadFile", s.WrapFunc(rbxmkLoadFile))
	lib.RawSetString("loadString", s.WrapFunc(rbxmkLoadString))
	lib.RawSetString("newAttrConfig", s.WrapFunc(rbxmkNewAttrConfig))
	lib.RawSetString("newCookie", s.WrapFunc(rbxmkNewCookie))
	lib.RawSetString("newDesc", s.WrapFunc(rbxmkNewDesc))
	lib.RawSetString("patchDesc", s.WrapFunc(rbxmkPatchDesc))
	lib.RawSetString("runFile", s.WrapFunc(rbxmkRunFile))
	lib.RawSetString("runString", s.WrapFunc(rbxmkRunString))

	mt := s.L.CreateTable(0, 2)
	mt.RawSetString("__index", s.WrapOperator(rbxmkOpIndex))
	mt.RawSetString("__newindex", s.WrapOperator(rbxmkOpNewindex))
	s.L.SetMetatable(lib, mt)

	return lib
}

func rbxmkCookiesFrom(s rbxmk.State) int {
	location := string(s.Pull(1, "string").(types.String))
	cookies, err := rbxmk.CookiesFrom(location)
	if err != nil {
		return s.RaiseError("unknown location %q", location)
	}
	if len(cookies) == 0 {
		s.Push(rtypes.Nil)
	}
	return s.Push(cookies)
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

func rbxmkDiffDesc(s rbxmk.State) int {
	var prev *rbxdump.Root
	var next *rbxdump.Root
	switch v := s.PullAnyOf(1, "RootDesc", "nil").(type) {
	case rtypes.NilType:
	case *rtypes.RootDesc:
		prev = v.Root
	default:
		return s.ReflectorError(1)
	}
	switch v := s.PullAnyOf(2, "RootDesc", "nil").(type) {
	case rtypes.NilType:
	case *rtypes.RootDesc:
		next = v.Root
	default:
		return s.ReflectorError(2)
	}
	actions := diff.Diff{Prev: prev, Next: next}.Diff()
	descActions := make(rtypes.DescActions, len(actions))
	for i, action := range actions {
		descActions[i] = &rtypes.DescAction{Action: action}
	}
	return s.Push(descActions)
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

func rbxmkNewAttrConfig(s rbxmk.State) int {
	var v rtypes.AttrConfig
	v.Property = string(s.PullOpt(1, "string", types.String("")).(types.String))
	return s.Push(&v)
}

func rbxmkNewCookie(s rbxmk.State) int {
	name := string(s.Pull(1, "string").(types.String))
	value := string(s.Pull(2, "string").(types.String))
	cookie := rtypes.Cookie{Cookie: &http.Cookie{Name: name, Value: value}}
	return s.Push(cookie)
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
		desc, _ := s.PullOpt(3, "RootDesc", nil).(*rtypes.RootDesc)
		if desc == nil {
			s.Desc = nil
		}
		s.Desc = desc
		return 0
	case "globalAttrConfig":
		attrcfg, _ := s.PullOpt(3, "AttrConfig", nil).(*rtypes.AttrConfig)
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
				"cookiesFrom": dump.Function{
					Parameters: dump.Parameters{
						{Name: "location", Type: dt.Prim("string"),
							Enums: dt.Enums{
								`"studio"`,
							},
						},
					},
					Returns: dump.Parameters{
						{Name: "cookies", Type: dt.Prim("Cookies")},
					},
					CanError:    true,
					Summary:     "Libraries/rbxmk:Fields/cookiesFrom/Summary",
					Description: "Libraries/rbxmk:Fields/cookiesFrom/Description",
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
				"diffDesc": dump.Function{
					Parameters: dump.Parameters{
						{Name: "prev", Type: dt.Optional{T: dt.Prim("RootDesc")}},
						{Name: "next", Type: dt.Optional{T: dt.Prim("RootDesc")}},
					},
					Returns: dump.Parameters{
						{Name: "diff", Type: dt.Array{T: dt.Prim("DescAction")}},
					},
					Summary:     "Libraries/rbxmk:Fields/diffDesc/Summary",
					Description: "Libraries/rbxmk:Fields/diffDesc/Description",
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
				"newAttrConfig": dump.Function{
					Parameters: dump.Parameters{
						{Name: "property", Type: dt.Prim("string")},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim("AttrConfig")},
					},
					Summary:     "Libraries/rbxmk:Fields/newAttrConfig/Summary",
					Description: "Libraries/rbxmk:Fields/newAttrConfig/Description",
				},
				"newCookie": dump.Function{
					Parameters: dump.Parameters{
						{Name: "name", Type: dt.Prim("string")},
						{Name: "value", Type: dt.Prim("string")},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim("Cookie")},
					},
					Summary:     "Libraries/rbxmk:Fields/newCookie/Summary",
					Description: "Libraries/rbxmk:Fields/newCookie/Description",
				},
				"newDesc": dump.Function{
					Parameters: dump.Parameters{
						{Name: "name", Type: dt.Prim("string"),
							Enums: dt.Enums{
								`"RootDesc"`,
								`"ClassDesc"`,
								`"PropertyDesc"`,
								`"FunctionDesc"`,
								`"EventDesc"`,
								`"CallbackDesc"`,
								`"ParameterDesc"`,
								`"TypeDesc"`,
								`"EnumDesc"`,
								`"EnumItemDesc"`,
							},
						},
					},
					Returns: dump.Parameters{
						{Type: dt.Prim("Descriptor")},
					},
					CanError:    true,
					Summary:     "Libraries/rbxmk:Fields/newDesc/Summary",
					Description: "Libraries/rbxmk:Fields/newDesc/Description",
				},
				"patchDesc": dump.Function{
					Parameters: dump.Parameters{
						{Name: "desc", Type: dt.Prim("RootDesc")},
						{Name: "actions", Type: dt.Array{T: dt.Prim("DescAction")}},
					},
					Summary:     "Libraries/rbxmk:Fields/patchDesc/Summary",
					Description: "Libraries/rbxmk:Fields/patchDesc/Description",
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
		Types: dump.TypeDefs{
			"Descriptor": dump.TypeDef{
				Underlying: dt.Or{
					dt.Prim("RootDesc"),
					dt.Prim("ClassDesc"),
					dt.Prim("PropertyDesc"),
					dt.Prim("FunctionDesc"),
					dt.Prim("EventDesc"),
					dt.Prim("CallbackDesc"),
					dt.Prim("ParameterDesc"),
					dt.Prim("TypeDesc"),
					dt.Prim("EnumDesc"),
					dt.Prim("EnumItemDesc"),
				},
				Summary:     "Libraries/rbxmk/Types/Descriptor:Summary",
				Description: "Libraries/rbxmk/Types/Descriptor:Description",
			},
		},
	}
	return lib
}
