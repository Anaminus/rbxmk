package main

import (
	"errors"
	"fmt"
	"github.com/anaminus/rbxauth"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/config"
	"github.com/anaminus/rbxmk/luautil"
	"github.com/anaminus/rbxmk/scheme"
	"github.com/anaminus/rbxmk/types"
	"github.com/robloxapi/rbxapi"
	"github.com/robloxapi/rbxapi/dump"
	"github.com/yuin/gopher-lua"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
)

const MainLibName = "rbxmk"

const (
	luaTypeInput  = "input"
	luaTypeOutput = "output"
	luaTypeError  = "error"
	luaTypeAPI    = "api"
)

func returnTypedValue(l *lua.LState, value interface{}, valueType string) int {
	ud := l.NewUserData()
	ud.Value = value
	l.SetMetatable(ud, l.GetTypeMetatable(valueType))
	l.Push(ud)
	return 1
}

func getLuaContextUpvalue(l *lua.LState, index int) *luautil.LuaContext {
	return l.CheckUserData(lua.UpvalueIndex(index)).Value.(*luautil.LuaContext)
}

func OpenMain(l *lua.LState) int {
	// Expect a LuaContext as an upvalue.
	ctx := l.CheckUserData(lua.UpvalueIndex(1))

	// Metatables for custom types.
	l.SetFuncs(l.NewTypeMetatable(luaTypeInput), map[string]lua.LGFunction{
		"__type": func(l *lua.LState) int {
			l.Push(lua.LString(luaTypeInput))
			return 1
		},
		"__tostring": func(l *lua.LState) int {
			if typ := reflect.TypeOf(luautil.GetLuaValue(l.Get(1))); typ == nil {
				l.Push(lua.LString("<input(nil)>"))
			} else {
				l.Push(lua.LString("<input(" + typ.String() + ")>"))
			}
			return 1
		},
		"__metatable": func(l *lua.LState) int {
			l.Push(lua.LString("the metatable is locked"))
			return 1
		},
	})
	l.SetFuncs(l.NewTypeMetatable(luaTypeOutput), map[string]lua.LGFunction{
		"__type": func(l *lua.LState) int {
			l.Push(lua.LString(luaTypeOutput))
			return 1
		},
		"__tostring": func(l *lua.LState) int {
			l.Push(lua.LString("<output>"))
			return 1
		},
		"__metatable": func(l *lua.LState) int {
			l.Push(lua.LString("the metatable is locked"))
			return 1
		},
	})
	l.SetFuncs(l.NewTypeMetatable(luaTypeError), map[string]lua.LGFunction{
		"__type": func(l *lua.LState) int {
			l.Push(lua.LString(luaTypeError))
			return 1
		},
		"__tostring": func(l *lua.LState) int {
			lu := l.ToUserData(1)
			if lu != nil {
				if err, ok := lu.Value.(error); ok {
					l.Push(lua.LString(err.Error()))
					return 1
				}
			}
			l.Push(lua.LString("<error>"))
			return 1
		},
		"__metatable": func(l *lua.LState) int {
			l.Push(lua.LString("the metatable is locked"))
			return 1
		},
	})
	l.SetFuncs(l.NewTypeMetatable(luaTypeAPI), map[string]lua.LGFunction{
		"__type": func(l *lua.LState) int {
			l.Push(lua.LString(luaTypeAPI))
			return 1
		},
		"__tostring": func(l *lua.LState) int {
			l.Push(lua.LString("<api>"))
			return 1
		},
		"__metatable": func(l *lua.LState) int {
			l.Push(lua.LString("the metatable is locked"))
			return 1
		},
	})

	// Set LuaContext as upvalue to each function in library.
	l.SetFuncs(l.RegisterModule(MainLibName, nil).(*lua.LTable), mainFuncs, ctx)
	return 1
}

var mainFuncs = map[string]lua.LGFunction{
	"input":     mainInput,
	"output":    mainOutput,
	"filter":    mainFilter,
	"map":       mainMap,
	"delete":    mainDelete,
	"load":      mainLoad,
	"type":      mainType,
	"getenv":    mainGetenv,
	"path":      mainPath,
	"readdir":   mainReaddir,
	"filename":  mainFilename,
	"sprintf":   mainSprintf,
	"printf":    mainPrintf,
	"loadapi":   mainLoadAPI,
	"configure": mainConfigure,
}

const formatIndex = "format"
const apiIndex = "api"

func throwError(l *lua.LState, err error) int {
	l.Error(lua.LString(err.Error()), 1)
	return 0
}

func string_Format(l *lua.LState) int {
	str := l.CheckString(1)
	args := make([]interface{}, l.GetTop()-1)
	top := l.GetTop()
	for i := 2; i <= top; i++ {
		args[i-2] = l.Get(i)
	}
	npat := strings.Count(str, "%") - strings.Count(str, "%%")
	if len(args) < npat {
		npat = len(args)
	}
	l.Push(lua.LString(fmt.Sprintf(str, args[:npat]...)))
	return 1
}

func mainInput(l *lua.LState) int {
	t := luautil.GetArgs(l, 1)
	ctx := getLuaContextUpvalue(l, 1)

	api, _ := t.FieldTyped(apiIndex, luaTypeAPI, true).(*rbxapi.API)
	config.SetAPI(ctx.Options, api)

	node := &rbxmk.InputNode{}
	node.Format, _ = t.FieldString(formatIndex, true)
	node.Options = ctx.Options

	nt := t.Len()
	if nt == 0 {
		throwError(l, errors.New("at least 1 reference argument is required"))
	}
	i := 1
	if t.TypeOfIndex(i) == "input" {
		node.Data = t.IndexTyped(i, "input", false).(rbxmk.Data)
		i = 2
	}
	for ; i <= nt; i++ {
		node.Reference = append(node.Reference, t.IndexString(i, false))
	}

	data, err := node.ResolveReference()
	if err != nil {
		return throwError(l, err)
	}

	return returnTypedValue(l, data, luaTypeInput)
}

func mainOutput(l *lua.LState) int {
	t := luautil.GetArgs(l, 1)
	ctx := getLuaContextUpvalue(l, 1)

	api, _ := t.FieldTyped(apiIndex, luaTypeAPI, true).(*rbxapi.API)
	config.SetAPI(ctx.Options, api)

	nt := t.Len()
	if nt == 0 {
		throwError(l, errors.New("at least 1 reference argument is required"))
	}

	node := &rbxmk.OutputNode{}
	node.Options = ctx.Options
	i := 1
	if t.TypeOfIndex(i) == "output" {
		originNode := t.IndexTyped(i, "output", false).(*rbxmk.OutputNode)
		node.Format = originNode.Format
		node.Reference = make([]string, len(originNode.Reference), len(originNode.Reference)+nt-1)
		copy(node.Reference, originNode.Reference)
		i = 2
	} else {
		node.Reference = make([]string, 0, nt)
	}
	if format, ok := t.FieldString(formatIndex, true); ok {
		node.Format = format
	}
	for ; i <= nt; i++ {
		node.Reference = append(node.Reference, t.IndexString(i, false))
	}

	return returnTypedValue(l, node, luaTypeOutput)
}

func mainFilter(l *lua.LState) int {
	t := luautil.GetArgs(l, 1)
	ctx := getLuaContextUpvalue(l, 1)

	api, _ := t.FieldTyped(apiIndex, luaTypeAPI, true).(*rbxapi.API)
	config.SetAPI(ctx.Options, api)

	const filterNameIndex = "name"
	var i int = 1
	filterName, ok := t.FieldString(filterNameIndex, true)
	if !ok {
		filterName = t.IndexString(i, false)
		i = 2
	}

	filterFunc := ctx.Options.Filters.Filter(filterName)
	if filterFunc == nil {
		return throwError(l, fmt.Errorf("unknown filter %q", filterName))
	}

	nt := t.Len()
	arguments := make([]interface{}, nt-i+1)
	for o := i; i <= nt; i++ {
		arguments[i-o] = t.IndexValue(i)
	}

	results, err := rbxmk.CallFilter(filterFunc, ctx.Options, arguments...)
	if err != nil {
		return throwError(l, err)
	}

	for _, result := range results {
		var lv lua.LValue
		switch v := result.(type) {
		case bool:
			lv = lua.LBool(v)
		case lua.LGFunction:
			lv = l.NewFunction(v)
		case int:
			lv = lua.LNumber(float64(v))
		case float64:
			lv = lua.LNumber(v)
		case string:
			lv = lua.LString(v)
		case uint:
			lv = lua.LNumber(uint(v))
		case *rbxmk.OutputNode:
			lu := l.NewUserData()
			lu.Value = v
			l.SetMetatable(lu, l.GetTypeMetatable(luaTypeOutput))
			lv = lu
		case rbxmk.Data:
			lu := l.NewUserData()
			lu.Value = v
			l.SetMetatable(lu, l.GetTypeMetatable(luaTypeInput))
			lv = lu
		default:
			lv = lua.LNil
		}
		l.Push(lv)
	}
	return len(results)
}

func mapNodes(l *lua.LState, inputs []rbxmk.Data, outputs []*rbxmk.OutputNode) int {
	for _, input := range inputs {
		for _, output := range outputs {
			if err := output.ResolveReference(input); err != nil {
				return throwError(l, err)
			}
		}
	}
	return 0
}

func mainMap(l *lua.LState) int {
	t := luautil.GetArgs(l, 1)

	inputs := make([]rbxmk.Data, 0, 1)
	outputs := make([]*rbxmk.OutputNode, 0, 1)

	nt := t.Len()
	for i := 1; i <= nt; i++ {
		switch t.TypeOfIndex(i) {
		case "input":
			inputs = append(inputs, t.IndexValue(i).(rbxmk.Data))
		case "output":
			outputs = append(outputs, t.IndexValue(i).(*rbxmk.OutputNode))
		}
	}
	if len(inputs) == 0 {
		return throwError(l, errors.New("at least 1 input is expected"))
	}
	if len(outputs) == 0 {
		return throwError(l, errors.New("at least 1 output is expected"))
	}

	return mapNodes(l, inputs, outputs)
}

func mainDelete(l *lua.LState) int {
	t := luautil.GetArgs(l, 1)

	outputs := make([]*rbxmk.OutputNode, 0, 1)
	nt := t.Len()
	for i := 1; i <= nt; i++ {
		switch t.TypeOfIndex(i) {
		case "output":
			outputs = append(outputs, t.IndexValue(i).(*rbxmk.OutputNode))
		}
	}
	if len(outputs) == 0 {
		return throwError(l, errors.New("at least 1 output is expected"))
	}

	return mapNodes(l, []rbxmk.Data{types.Delete{}}, outputs)
}

func mainLoad(l *lua.LState) int {
	t := luautil.GetArgs(l, 1)
	ctx := getLuaContextUpvalue(l, 1)

	fileName := t.IndexString(1, false)
	fileName = shortenPath(filepath.Clean(fileName))
	fi, err := os.Stat(fileName)
	if err != nil {
		return throwError(l, err)
	}
	if err = ctx.PushFile(luautil.FileInfo{fileName, fi}); err != nil {
		return throwError(l, err)
	}

	// Load file as function.
	fn, err := l.LoadFile(fileName)
	if err != nil {
		ctx.PopFile()
		return throwError(l, err)
	}
	l.Push(fn) // +function

	// Push extra arguments as arguments to loaded function.
	nt := t.Len()
	for i := 2; i <= nt; i++ {
		l.Push(t.RawGetInt(i)) // function, ..., +arg
	}
	// function, +args...

	// Call loaded function.
	err = l.PCall(nt-1, lua.MultRet, nil) // -function, -args..., +returns...
	ctx.PopFile()
	if err != nil {
		return throwError(l, err)
	}
	return l.GetTop() - 1
}

func mainType(l *lua.LState) int {
	t := luautil.GetArgs(l, 1)
	l.Push(lua.LString(t.TypeOfIndex(1)))
	return 1
}

func mainGetenv(l *lua.LState) int {
	t := luautil.GetArgs(l, 1)
	value, ok := os.LookupEnv(t.IndexString(1, false))
	if ok {
		l.Push(lua.LString(value))
	} else {
		l.Push(lua.LNil)
	}
	return 1
}

func mainPath(l *lua.LState) int {
	t := luautil.GetArgs(l, 1)
	ctx := getLuaContextUpvalue(l, 1)

	s := make([]string, t.Len())
	for i := 1; i <= t.Len(); i++ {
		s[i-1] = os.Expand(t.IndexString(i, false), func(v string) string {
			switch v {
			case "script_name", "sn":
				fi, ok := ctx.PeekFile()
				if !ok {
					l.Push(lua.LString(""))
					break
				}
				path, _ := filepath.Abs(fi.Path)
				return filepath.Base(path)
			case "script_directory", "script_dir", "sd":
				fi, ok := ctx.PeekFile()
				if !ok {
					l.Push(lua.LString(""))
					break
				}
				path, _ := filepath.Abs(fi.Path)
				return filepath.Dir(path)
			case "working_directory", "working_dir", "wd":
				wd, _ := os.Getwd()
				return wd
			case "temp_directory", "temp_dir", "tmp":
				return os.TempDir()
			}
			return ""
		})
	}
	filename := shortenPath(filepath.Join(s...))
	l.Push(lua.LString(filename))
	return 1
}

func mainReaddir(l *lua.LState) int {
	t := luautil.GetArgs(l, 1)
	dirname := t.IndexString(1, false)
	f, err := os.Open(dirname)
	if err != nil {
		return throwError(l, err)
	}
	defer f.Close()
	names, err := f.Readdirnames(-1)
	if err != nil {
		return throwError(l, err)
	}
	sort.Strings(names)
	tnames := l.CreateTable(len(names), 0)
	for _, name := range names {
		tnames.Append(lua.LString(name))
	}
	l.Push(tnames)
	return 1
}

func mainFilename(l *lua.LState) int {
	t := luautil.GetArgs(l, 1)
	ctx := getLuaContextUpvalue(l, 1)

	typ := t.IndexString(1, false)
	path := t.IndexString(2, false)
	var result string
	switch typ {
	case "dir":
		result = filepath.Dir(path)
	case "name":
		result = filepath.Base(path)
	case "base":
		result = filepath.Base(path)
		result = result[:len(result)-len(filepath.Ext(path))]
	case "ext":
		result = filepath.Ext(path)
	case "fbase":
		ext := scheme.GuessFileExtension(ctx.Options, "", path)
		if ext != "" && ext != "." {
			ext = "." + ext
		}
		result = filepath.Base(path)
		result = result[:len(result)-len(ext)]
	case "fext":
		result = scheme.GuessFileExtension(ctx.Options, "", path)
		if result != "" && result != "." {
			result = "." + result
		}
	default:
		return throwError(l, fmt.Errorf("unknown argument %q", typ))
	}
	l.Push(lua.LString(result))
	return 1
}

func mainSprintf(l *lua.LState) int {
	t := luautil.GetArgs(l, 1)
	t.PushAsArgs()
	string_Format(l)
	return 1
}

func mainPrintf(l *lua.LState) int {
	t := luautil.GetArgs(l, 1)
	t.PushAsArgs()
	string_Format(l)
	s := l.ToString(-1)
	l.Pop(1)
	fmt.Print(s)
	return 0
}

func mainLoadAPI(l *lua.LState) int {
	t := luautil.GetArgs(l, 1)
	file, err := os.Open(t.IndexString(1, false))
	if err != nil {
		return throwError(l, fmt.Errorf("failed to open API file: %s", err))
	}
	defer file.Close()
	api, err := dump.Decode(file)
	if err != nil {
		return throwError(l, fmt.Errorf("failed to decode API file: %s", err))
	}
	return returnTypedValue(l, api, luaTypeAPI)
}

func mainConfigure(l *lua.LState) int {
	t := luautil.GetArgs(l, 1)
	ctx := getLuaContextUpvalue(l, 1)

	if v := t.FieldValue("api"); v != nil {
		api, _ := v.(*rbxapi.API)
		config.SetAPI(ctx.Options, api)
	}
	if v := t.FieldValue("undef"); v != nil {
		env := config.PPEnvs(ctx.Options)[config.PPEnvScript]
		undefs := v.(*lua.LTable)
		for i := 1; i <= undefs.Len(); i++ {
			k := undefs.RawGetInt(i)
			if k.Type() != lua.LTString {
				continue
			}
			key := string(k.(lua.LString))
			if !luautil.CheckStringVar(key) {
				continue
			}
			env.RawSetString(key, lua.LNil)
		}
	}
	if v := t.FieldValue("define"); v != nil {
		env := config.PPEnvs(ctx.Options)[config.PPEnvScript]
		defs := v.(*lua.LTable)
		defs.ForEach(func(k, v lua.LValue) {
			if k.Type() != lua.LTString {
				return
			}
			key := string(k.(lua.LString))
			if !luautil.CheckStringVar(key) {
				return
			}
			switch v.Type() {
			case lua.LTBool, lua.LTNumber, lua.LTString:
				env.RawSetString(key, v)
			}
		})
	}
	if v := t.FieldValue("rbxauth"); v != nil {
		users := config.RobloxAuth(ctx.Options)
		defs := v.(*lua.LTable)
		var terr error
		defs.ForEach(func(k, v lua.LValue) {
			if terr != nil {
				return
			}
			if k.Type() != lua.LTString && k.Type() != lua.LTNumber {
				return
			}
			ident := k.String()
			identOpt, ok := v.(*lua.LTable)
			if !ok {
				return
			}

			cred := rbxauth.Cred{
				Type:  identOpt.RawGetString("type").String(),
				Ident: ident,
			}

			var cookies []*http.Cookie
			if path := identOpt.RawGetString("file"); path.Type() == lua.LTString {
				f, err := os.Open(path.String())
				if err != nil {
					terr = err
				}
				cookies, err = rbxauth.ReadCookies(f)
				f.Close()
				if err != nil {
					terr = err
					return
				}
			} else if prompt := identOpt.RawGetString("prompt"); prompt.Type() == lua.LTBool && prompt.(lua.LBool) == lua.LTrue {
				auth := &rbxauth.Config{Host: config.Host(ctx.Options)}
				var err error
				if cred, cookies, err = auth.PromptCred(cred); err != nil {
					terr = err
					return
				}
			}
			if len(cookies) > 0 {
				users[cred] = cookies
			}
		})
		if terr != nil {
			return throwError(l, terr)
		}
	}
	if v := t.FieldValue("host"); v != nil {
		config.SetHost(ctx.Options, string(v.(lua.LString)))
	}
	return 0
}
