package main

import (
	"errors"
	"fmt"
	"github.com/anaminus/rbxauth"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/scheme"
	"github.com/anaminus/rbxmk/types"
	"github.com/robloxapi/rbxapi"
	"github.com/robloxapi/rbxapi/rbxapidump"
	"github.com/yuin/gopher-lua"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
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

func getLuaContextUpvalue(l *lua.LState, index int) *LuaContext {
	return l.CheckUserData(lua.UpvalueIndex(index)).Value.(*LuaContext)
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
			if typ := reflect.TypeOf(GetLuaValue(l.Get(1))); typ == nil {
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
const userIndex = "user"
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

func getUserCred(v interface{}) (cred rbxauth.Cred, err error) {
	switch user := v.(type) {
	case *lua.LTable:
		cred = rbxauth.Cred{
			Type:  user.RawGetString("type").String(),
			Ident: user.RawGetString("ident").String(),
		}
	case lua.LString:
		// Assume Username.
		cred = rbxauth.Cred{Type: "Username", Ident: string(user)}
	case lua.LNumber:
		// Assume UserID.
		cred = rbxauth.Cred{Type: "UserID", Ident: user.String()}
	case nil:
	default:
		err = fmt.Errorf("'%s' field must be a table, string, or number", userIndex)
	}
	return cred, err
}

func mainInput(l *lua.LState) int {
	t := GetArgs(l, 1)
	ctx := getLuaContextUpvalue(l, 1)

	node := &rbxmk.InputNode{}
	node.Options = ctx.Options.Copy()

	api, _ := t.FieldTyped(apiIndex, luaTypeAPI, true).(rbxapi.Root)
	node.Options.Config["API"] = api

	node.Format, _ = t.FieldString(formatIndex, true)
	if user, err := getUserCred(t.FieldValue(userIndex)); err != nil {
		return throwError(l, err)
	} else {
		node.User = user
	}

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
	t := GetArgs(l, 1)
	ctx := getLuaContextUpvalue(l, 1)

	nt := t.Len()
	if nt == 0 {
		throwError(l, errors.New("at least 1 reference argument is required"))
	}

	node := &rbxmk.OutputNode{}
	node.Options = ctx.Options.Copy()

	api, _ := t.FieldTyped(apiIndex, luaTypeAPI, true).(rbxapi.Root)
	node.Options.Config["API"] = api

	i := 1
	if t.TypeOfIndex(i) == "output" {
		originNode := t.IndexTyped(i, "output", false).(*rbxmk.OutputNode)
		node.Data = originNode.Data
		node.Format = originNode.Format
		node.User = originNode.User
		node.Reference = make([]string, len(originNode.Reference), len(originNode.Reference)+nt-1)
		copy(node.Reference, originNode.Reference)
		i = 2
	} else if t.TypeOfIndex(i) == "input" {
		node.Data = t.IndexTyped(i, "input", false).(rbxmk.Data)
		i = 2
	} else {
		node.Reference = make([]string, 0, nt)
	}
	if format, ok := t.FieldString(formatIndex, true); ok {
		node.Format = format
	}
	if user, err := getUserCred(t.FieldValue(userIndex)); err != nil {
		return throwError(l, err)
	} else {
		node.User = user
	}

	for ; i <= nt; i++ {
		node.Reference = append(node.Reference, t.IndexString(i, false))
	}

	return returnTypedValue(l, node, luaTypeOutput)
}

func mainFilter(l *lua.LState) int {
	t := GetArgs(l, 1)
	ctx := getLuaContextUpvalue(l, 1)

	opt := ctx.Options.Copy()

	api, _ := t.FieldTyped(apiIndex, luaTypeAPI, true).(rbxapi.Root)
	opt.Config["API"] = api

	const filterNameIndex = "name"
	var i int = 1
	filterName, ok := t.FieldString(filterNameIndex, true)
	if !ok {
		filterName = t.IndexString(i, false)
		i = 2
	}

	filterFunc := opt.Filters.Filter(filterName)
	if filterFunc == nil {
		return throwError(l, fmt.Errorf("unknown filter %q", filterName))
	}

	nt := t.Len()
	arguments := make([]interface{}, nt-i+1)
	for o := i; i <= nt; i++ {
		arguments[i-o] = t.IndexValue(i)
	}

	results, err := rbxmk.CallFilter(filterFunc, opt, arguments...)
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
	t := GetArgs(l, 1)

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
	t := GetArgs(l, 1)

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
	t := GetArgs(l, 1)
	ctx := getLuaContextUpvalue(l, 1)

	fileName := t.IndexString(1, false)
	fileName = shortenPath(filepath.Clean(fileName))
	fi, err := os.Stat(fileName)
	if err != nil {
		return throwError(l, err)
	}
	if err = ctx.PushFile(FileInfo{fileName, fi}); err != nil {
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
	t := GetArgs(l, 1)
	l.Push(lua.LString(t.TypeOfIndex(1)))
	return 1
}

func mainGetenv(l *lua.LState) int {
	t := GetArgs(l, 1)
	value, ok := os.LookupEnv(t.IndexString(1, false))
	if ok {
		l.Push(lua.LString(value))
	} else {
		l.Push(lua.LNil)
	}
	return 1
}

func mainPath(l *lua.LState) int {
	t := GetArgs(l, 1)
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
	t := GetArgs(l, 1)
	dirname := t.IndexString(1, false)
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return throwError(l, err)
	}
	tfiles := l.CreateTable(len(files), 0)
	for _, info := range files {
		tinfo := l.CreateTable(0, 4)
		tinfo.RawSetString("name", lua.LString(info.Name()))
		tinfo.RawSetString("isdir", lua.LBool(info.IsDir()))
		tinfo.RawSetString("size", lua.LNumber(info.Size()))
		tinfo.RawSetString("modtime", lua.LNumber(info.ModTime().Unix()))
		tfiles.Append(tinfo)
	}
	l.Push(tfiles)
	return 1
}

func mainFilename(l *lua.LState) int {
	t := GetArgs(l, 1)
	ctx := getLuaContextUpvalue(l, 1)

	path := t.IndexString(1, false)
	for i := 2; i <= t.Len(); i++ {
		var result string
		switch typ := t.IndexString(i, false); typ {
		case "dir":
			result = filepath.Dir(path)
		case "base":
			result = filepath.Base(path)
		case "ext":
			result = filepath.Ext(path)
		case "stem":
			result = filepath.Base(path)
			result = result[:len(result)-len(filepath.Ext(path))]
		case "fext":
			result = scheme.GuessFileExtension(ctx.Options, "", path)
			if result != "" && result != "." {
				result = "." + result
			}
		case "fstem":
			ext := scheme.GuessFileExtension(ctx.Options, "", path)
			if ext != "" && ext != "." {
				ext = "." + ext
			}
			result = filepath.Base(path)
			result = result[:len(result)-len(ext)]
		default:
			return throwError(l, fmt.Errorf("unknown argument %q", typ))
		}
		l.Push(lua.LString(result))
	}
	return t.Len() - 1
}

func mainSprintf(l *lua.LState) int {
	t := GetArgs(l, 1)
	t.PushAsArgs()
	string_Format(l)
	return 1
}

func mainPrintf(l *lua.LState) int {
	t := GetArgs(l, 1)
	t.PushAsArgs()
	string_Format(l)
	s := l.ToString(-1)
	l.Pop(1)
	fmt.Print(s)
	return 0
}

func mainLoadAPI(l *lua.LState) int {
	t := GetArgs(l, 1)
	file, err := os.Open(t.IndexString(1, false))
	if err != nil {
		return throwError(l, fmt.Errorf("failed to open API file: %s", err))
	}
	defer file.Close()
	api, err := rbxapidump.Decode(file)
	if err != nil {
		return throwError(l, fmt.Errorf("failed to decode API file: %s", err))
	}
	return returnTypedValue(l, api, luaTypeAPI)
}

func mainConfigure(l *lua.LState) int {
	t := GetArgs(l, 1)
	ctx := getLuaContextUpvalue(l, 1)

	if v := t.FieldValue("api"); v != nil {
		api, _ := v.(rbxapi.Root)
		ctx.Options.Config["API"] = api
	}
	if v := t.FieldValue("undef"); v != nil {
		env := ctx.Options.Config["PPEnv"].([]*lua.LTable)[PPEnvScript]
		undefs := v.(*lua.LTable)
		for i := 1; i <= undefs.Len(); i++ {
			k := undefs.RawGetInt(i)
			if k.Type() != lua.LTString {
				continue
			}
			key := string(k.(lua.LString))
			if !CheckStringVar(key) {
				continue
			}
			env.RawSetString(key, lua.LNil)
		}
	}
	if v := t.FieldValue("define"); v != nil {
		env := ctx.Options.Config["PPEnv"].([]*lua.LTable)[PPEnvScript]
		defs := v.(*lua.LTable)
		defs.ForEach(func(k, v lua.LValue) {
			if k.Type() != lua.LTString {
				return
			}
			key := string(k.(lua.LString))
			if !CheckStringVar(key) {
				return
			}
			switch v.Type() {
			case lua.LTBool, lua.LTNumber, lua.LTString:
				env.RawSetString(key, v)
			default:
				if typeOf(l, v) == "input" {
					str, err := types.ToString(v.(*lua.LUserData).Value)
					if err != nil {
						return
					}
					env.RawSetString(key, lua.LString(str.GetString()))
				}
			}
		})
	}
	if v := t.FieldValue("rbxauth"); v != nil {
		users, _ := ctx.Options.Config["RobloxAuth"].(map[rbxauth.Cred][]*http.Cookie)
		defs := v.(*lua.LTable)
		for i := 1; i <= defs.Len(); i++ {
			entry, ok := defs.RawGetInt(i).(*lua.LTable)
			if !ok {
				continue
			}

			cred := rbxauth.Cred{}
			if v, ok := entry.RawGetString("type").(lua.LString); ok {
				cred.Type = string(v)
			}
			if v, ok := entry.RawGetString("ident").(lua.LString); ok {
				cred.Type = string(v)
			}

			if logout, ok := entry.RawGetString("logout").(lua.LBool); ok && logout == lua.LTrue {
				cookies := users[cred]
				if len(cookies) == 0 {
					return throwError(l, errors.New("cannot logout: unknown credentials"))
				}
				stream := rbxauth.StandardStream()
				if err := stream.Logout(cookies); err != nil {
					return throwError(l, fmt.Errorf("Error logging out: %s", err))
				}
			}

			var cookies []*http.Cookie
			if path := entry.RawGetString("file"); path.Type() == lua.LTString {
				f, err := os.Open(path.String())
				if err != nil {
					return throwError(l, err)
				}
				cookies, err = rbxauth.ReadCookies(f)
				f.Close()
				if err != nil {
					return throwError(l, err)
				}
			} else if prompt := entry.RawGetString("prompt"); prompt.Type() == lua.LTBool && prompt.(lua.LBool) == lua.LTrue {
				stream := rbxauth.StandardStream()
				var err error
				if cred, cookies, err = stream.PromptCred(cred); err != nil {
					return throwError(l, err)
				}
			}
			if len(cookies) > 0 {
				if users == nil {
					users = make(map[rbxauth.Cred][]*http.Cookie)
					ctx.Options.Config["RobloxAuth"] = users
				}
				users[cred] = cookies
			}
		}
	}
	if v := t.FieldValue("host"); v != nil {
		ctx.Options.Config["Host"] = string(v.(lua.LString))
	}
	return 0
}
