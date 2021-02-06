package library

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/parse"
	"github.com/anaminus/rbxmk"
	reflect "github.com/anaminus/rbxmk/library/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/rbxdump/diff"
	"github.com/robloxapi/types"
	reg "golang.org/x/sys/windows/registry"
)

func init() { register(RBXMK, 0) }

var RBXMK = rbxmk.Library{
	Name: "rbxmk",
	Open: func(s rbxmk.State) *lua.LTable {
		lib := s.L.CreateTable(0, 12)
		lib.RawSetString("loadFile", s.WrapFunc(rbxmkLoadFile))
		lib.RawSetString("loadString", s.WrapFunc(rbxmkLoadString))
		lib.RawSetString("runFile", s.WrapFunc(rbxmkRunFile))
		lib.RawSetString("runString", s.WrapFunc(rbxmkRunString))
		lib.RawSetString("newDesc", s.WrapFunc(rbxmkNewDesc))
		lib.RawSetString("diffDesc", s.WrapFunc(rbxmkDiffDesc))
		lib.RawSetString("patchDesc", s.WrapFunc(rbxmkPatchDesc))
		lib.RawSetString("encodeFormat", s.WrapFunc(rbxmkEncodeFormat))
		lib.RawSetString("decodeFormat", s.WrapFunc(rbxmkDecodeFormat))
		lib.RawSetString("formatCanDecode", s.WrapFunc(rbxmkFormatCanDecode))
		lib.RawSetString("newCookie", s.WrapFunc(rbxmkNewCookie))
		lib.RawSetString("cookiesFrom", s.WrapFunc(rbxmkCookiesFrom))

		for _, f := range reflect.All() {
			r := f()
			s.RegisterReflector(r)
			s.ApplyReflector(r, lib)
		}

		mt := s.L.CreateTable(0, 2)
		mt.RawSetString("__index", s.WrapFunc(func(s rbxmk.State) int {
			switch field := s.Pull(2, "string").(types.String); field {
			case "globalDesc":
				desc := s.Desc(nil)
				if desc == nil {
					return s.Push(rtypes.Nil)
				}
				return s.Push(desc)
			case "globalAttrConfig":
				attrcfg := s.AttrConfig(nil)
				if attrcfg == nil {
					return s.Push(rtypes.Nil)
				}
				return s.Push(attrcfg)
			default:
				return s.RaiseError("unknown field %q", field)
			}
		}))
		mt.RawSetString("__newindex", s.WrapFunc(func(s rbxmk.State) int {
			switch field := s.Pull(2, "string").(types.String); field {
			case "globalDesc":
				desc, _ := s.PullOpt(3, "RootDesc", nil).(*rtypes.RootDesc)
				if desc == nil {
					s.SetDesc(nil)
				}
				s.SetDesc(desc)
				return 0
			case "globalAttrConfig":
				attrcfg, _ := s.PullOpt(3, "AttrConfig", nil).(*rtypes.AttrConfig)
				if attrcfg == nil {
					s.SetAttrConfig(nil)
				}
				s.SetAttrConfig(attrcfg)
				return 0
			default:
				return s.RaiseError("unknown field %q", field)
			}
		}))
		s.L.SetMetatable(lib, mt)

		return lib
	},
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
	fi, err := os.Stat(fileName)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	if err = s.PushFile(rbxmk.FileInfo{Path: fileName, FileInfo: fi}); err != nil {
		return s.RaiseError("%s", err)
	}

	nt := s.L.GetTop()

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
	return s.L.GetTop() - nt
}

func rbxmkRunString(s rbxmk.State) int {
	source := s.CheckString(1)
	nt := s.L.GetTop()

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
	selector := s.Pull(1, "FormatSelector").(rtypes.FormatSelector)
	format := s.Format(selector.Format)
	if format.Name == "" {
		return s.RaiseError("unknown format %q", selector.Format)
	}
	if format.Encode == nil {
		return s.RaiseError("cannot encode with format %s", format.Name)
	}
	var w bytes.Buffer
	if err := format.Encode(selector, &w, s.Pull(2, "Variant")); err != nil {
		return s.RaiseError("%s", err)
	}
	return s.Push(types.BinaryString(w.Bytes()))
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
	v, err := format.Decode(selector, r)
	if err != nil {
		return s.RaiseError("%s", err)
	}
	return s.Push(v)
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
	return s.Push(types.Bool(format.CanDecode(selector, typeName)))
}

func rbxmkNewCookie(s rbxmk.State) int {
	name := string(s.Pull(1, "string").(types.String))
	value := string(s.Pull(2, "string").(types.String))
	cookie := rtypes.Cookie{Cookie: &http.Cookie{Name: name, Value: value}}
	return s.Push(cookie)
}

func rbxmkCookiesFrom(s rbxmk.State) int {
	location := string(s.Pull(1, "string").(types.String))
	switch strings.ToLower(location) {
	case "studio":
		if cookies := cookiesFromStudio(); len(cookies) > 0 {
			fmt.Println("CHECK", cookies[0].Value)
			return s.Push(cookies)
		}
		return s.Push(rtypes.Nil)
	default:
		return s.RaiseError("unknown location %q", location)
	}
}

func cookiesFromStudio() rtypes.Cookies {
	const keyPath = `Software\Roblox\RobloxStudioBrowser\roblox.com`
	key, err := reg.OpenKey(reg.CURRENT_USER, keyPath, reg.QUERY_VALUE)
	if err != nil {
		return nil
	}
	defer key.Close()
	v, _, err := key.GetStringValue(".ROBLOSECURITY")
	if err != nil {
		return nil
	}
	cookie := &http.Cookie{
		Name:   ".ROBLOSECURITY",
		Domain: "roblox.com",
	}
	if !parseRegistryCookie(cookie, v) {
		return nil
	}
	return rtypes.Cookies{rtypes.Cookie{Cookie: cookie}}
}

func parseRegistryCookie(cookie *http.Cookie, v string) bool {
	r := parse.NewTextReader(strings.NewReader(v))
	if r.Is("SEC::<YES>") {
		cookie.Secure = true
		r.Is(",")
	}
	if r.Is("EXP::<") {
		value, ok := r.Until('>')
		if !ok {
			return false
		}
		t, err := time.Parse(time.RFC3339, value)
		if err != nil {
			return false
		}
		cookie.Expires = t
		r.Is(",")
	}
	if r.Is("COOK::<") {
		value, ok := r.UntilEOF()
		if !ok {
			return false
		}
		cookie.Value = value[:len(value)-1]
	}
	return r.IsEOF()
}
