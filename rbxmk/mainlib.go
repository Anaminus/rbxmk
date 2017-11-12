package main

import (
	"errors"
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/scheme"
	"github.com/anaminus/rbxmk/types"
	"github.com/robloxapi/rbxapi"
	"github.com/yuin/gopher-lua"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const MainLibName = "rbxmk"

func OpenMain(l *lua.LState) int {
	// Expect a LuaState as the second argument.
	st := l.CheckUserData(2)
	_ = st.Value.(*LuaState)
	// Set LuaState as upvalue to each function in library.
	l.SetFuncs(l.RegisterModule(MainLibName, nil).(*lua.LTable), mainFuncs, st)
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

func getLuaStateUpvalue(l *lua.LState, index int) *LuaState {
	return l.CheckUserData(lua.UpvalueIndex(index)).Value.(*LuaState)
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
	t := GetArgs(l, 1)
	st := getLuaStateUpvalue(l, 1)

	st.options.Config.API, _ = t.FieldTyped(apiIndex, luaTypeAPI, true).(*rbxapi.API)

	node := &rbxmk.InputNode{}
	node.Format, _ = t.FieldString(formatIndex, true)
	node.Options = st.options

	nt := t.Len()
	if nt == 0 {
		throwError(l, errors.New("at least 1 reference argument is required"))
	}
	i := 1
	if data, ok := t.IndexValue(i).(rbxmk.Data); ok {
		node.Data = data
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
	st := getLuaStateUpvalue(l, 1)

	st.options.Config.API, _ = t.FieldTyped(apiIndex, luaTypeAPI, true).(*rbxapi.API)

	node := &rbxmk.OutputNode{}
	node.Format, _ = t.FieldString(formatIndex, true)
	node.Options = st.options

	nt := t.Len()
	if nt == 0 {
		throwError(l, errors.New("at least 1 reference argument is required"))
	}
	i := 1
	if data, ok := t.IndexValue(i).(rbxmk.Data); ok {
		node.Data = data
		i = 2
	}
	for ; i <= nt; i++ {
		node.Reference = append(node.Reference, t.IndexString(i, false))
	}

	return returnTypedValue(l, node, luaTypeOutput)
}

func mainFilter(l *lua.LState) int {
	t := GetArgs(l, 1)
	st := getLuaStateUpvalue(l, 1)

	st.options.Config.API, _ = t.FieldTyped(apiIndex, luaTypeAPI, true).(*rbxapi.API)

	const filterNameIndex = "name"
	var i int = 1
	filterName, ok := t.FieldString(filterNameIndex, true)
	if !ok {
		filterName = t.IndexString(i, false)
		i = 2
	}

	filterFunc := st.options.Filters.Filter(filterName)
	if filterFunc == nil {
		return throwError(l, fmt.Errorf("unknown filter %q", filterName))
	}

	nt := t.Len()
	arguments := make([]interface{}, nt-i+1)
	for o := i; i <= nt; i++ {
		arguments[i-o] = t.IndexValue(i)
	}

	results, err := rbxmk.CallFilter(filterFunc, st.options, arguments...)
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

func mainMap(l *lua.LState) int {
	t := GetArgs(l, 1)
	st := getLuaStateUpvalue(l, 1)

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

	return st.mapNodes(inputs, outputs)
}

func mainDelete(l *lua.LState) int {
	t := GetArgs(l, 1)
	st := getLuaStateUpvalue(l, 1)

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

	return st.mapNodes([]rbxmk.Data{types.Delete{}}, outputs)
}

func mainLoad(l *lua.LState) int {
	t := GetArgs(l, 1)
	st := getLuaStateUpvalue(l, 1)

	fileName := t.IndexString(1, false)
	fileName = shortenPath(filepath.Clean(fileName))
	fi, err := os.Stat(fileName)
	if err != nil {
		return throwError(l, err)
	}
	if err = st.pushFile(&fileInfo{fileName, fi}); err != nil {
		return throwError(l, err)
	}

	// Load file as function.
	fn, err := l.LoadFile(fileName)
	if err != nil {
		st.popFile()
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
	st.popFile()
	if err != nil {
		return throwError(l, err)
	}
	return l.GetTop() - 1
}

func mainType(l *lua.LState) int {
	t := GetArgs(l, 1)
	l.Push(lua.LString(typeOf(l, t.RawGetInt(1))))
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
	st := getLuaStateUpvalue(l, 1)

	s := make([]string, t.Len())
	for i := 1; i <= t.Len(); i++ {
		s[i-1] = os.Expand(t.IndexString(i, false), func(v string) string {
			switch v {
			case "script_name", "sn":
				n := len(st.fileStack)
				if n == 0 {
					l.Push(lua.LString(""))
					break
				}
				path, _ := filepath.Abs(st.fileStack[n-1].path)
				return filepath.Base(path)
			case "script_directory", "script_dir", "sd":
				n := len(st.fileStack)
				if n == 0 {
					l.Push(lua.LString(""))
					break
				}
				path, _ := filepath.Abs(st.fileStack[n-1].path)
				return filepath.Dir(path)
			case "working_directory", "working_dir", "wd":
				wd, _ := os.Getwd()
				return wd
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
	t := GetArgs(l, 1)
	st := getLuaStateUpvalue(l, 1)

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
		ext := scheme.GuessFileExtension(st.options, "", path)
		if ext != "" && ext != "." {
			ext = "." + ext
		}
		result = filepath.Base(path)
		result = result[:len(result)-len(ext)]
	case "fext":
		result = scheme.GuessFileExtension(st.options, "", path)
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
	path := t.IndexString(1, false)
	api, err := rbxmk.LoadAPI(path)
	if err != nil {
		return throwError(l, err)
	}
	return returnTypedValue(l, api, luaTypeAPI)
}

func mainConfigure(l *lua.LState) int {
	t := GetArgs(l, 1)
	st := getLuaStateUpvalue(l, 1)

	if v := t.FieldValue("api"); v != nil {
		st.options.Config.API, _ = v.(*rbxapi.API)
	}
	return 0
}
