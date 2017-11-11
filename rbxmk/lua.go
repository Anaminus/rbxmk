package main

import (
	"errors"
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/types"
	"github.com/robloxapi/rbxapi"
	"github.com/yuin/gopher-lua"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
)

type fileInfo struct {
	path string
	fi   os.FileInfo
}

type LuaState struct {
	options   rbxmk.Options
	state     *lua.LState
	fileStack []*fileInfo
}

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

func throwError(l *lua.LState, err error) int {
	l.Error(lua.LString(err.Error()), 1)
	return 0
}

func typeOf(l *lua.LState, v lua.LValue) string {
	if v.Type() == lua.LTUserData {
		if s, ok := l.CallMeta(v, "__type").(lua.LString); ok {
			return string(s)
		}
	}
	return v.Type().String()
}

const tableArg = 1
const tableMethodArg = 1

type exitMarker struct {
	err error
}

func (exitMarker) Error() string {
	return "ExitMarker"
}

type tArgs struct {
	l *lua.LState
	*lua.LTable
}

func GetArgs(l *lua.LState, index int) tArgs {
	tb := l.Get(index)
	if l.GetTop() != index || tb.Type() != lua.LTTable {
		l.RaiseError("function must have 1 table argument")
	} else if l.GetMetatable(tb) != lua.LNil {
		l.RaiseError("table argument cannot have metatable")
	}
	return tArgs{l: l, LTable: tb.(*lua.LTable)}
}

func (t tArgs) ErrorField(name string, expected, got string) {
	if got == "" {
		t.l.RaiseError("bad value at field %q: %s expected", name, expected)
	} else {
		t.l.RaiseError("bad value at field %q: %s expected, got %s", name, expected, got)
	}
}

func (t tArgs) ErrorIndex(index int, expected, got string) {
	if got == "" {
		t.l.RaiseError("bad value at index #%d: %s expected", index, expected)
	} else {
		t.l.RaiseError("bad value at index #%d: %s expected, got %s", index, expected, got)
	}
}

func (t tArgs) TypeOfField(name string) string {
	return typeOf(t.l, t.RawGetString(name))
}

func (t tArgs) TypeOfIndex(index int) string {
	return typeOf(t.l, t.RawGetInt(index))
}

func (t tArgs) FieldString(name string, opt bool) (s string, ok bool) {
	lv := t.RawGetString(name)
	if typ := typeOf(t.l, lv); typ != "string" {
		if opt && typ == "nil" {
			return "", false
		}
		t.ErrorField(name, "string", typ)
	}
	return string(lv.(lua.LString)), true
}

func (t tArgs) IndexString(index int, opt bool) string {
	lv := t.RawGetInt(index)
	if typ := typeOf(t.l, lv); typ != "string" {
		if opt && typ == "nil" {
			return ""
		}
		t.ErrorIndex(index, "string", typ)
	}
	return string(lv.(lua.LString))
}

func (t tArgs) FieldTyped(name string, valueType string, opt bool) (v interface{}) {
	lv := t.RawGetString(name)
	if typ := typeOf(t.l, lv); typ != valueType {
		if opt && typ == "nil" {
			return nil
		}
		t.ErrorField(name, valueType, typ)
	}
	uv, _ := lv.(*lua.LUserData)
	return uv.Value
}

func (t tArgs) IndexTyped(index int, valueType string, opt bool) (v interface{}) {
	lv := t.RawGetInt(index)
	if typ := typeOf(t.l, lv); typ != valueType {
		if opt && typ == "nil" {
			return nil
		}
		t.ErrorIndex(index, valueType, typ)
	}
	uv, _ := lv.(*lua.LUserData)
	return uv.Value
}

func getLuaValue(lv lua.LValue) interface{} {
	switch v := lv.(type) {
	case lua.LBool:
		return bool(v)
	case lua.LNumber:
		return float64(v)
	case lua.LString:
		return string(v)
	case *lua.LTable:
		return v
	case *lua.LFunction:
		return v
	case *lua.LState:
		return v
	case *lua.LUserData:
		return v.Value
	}
	return nil
}

func (t tArgs) IndexValue(index int) interface{} {
	return getLuaValue(t.RawGetInt(index))
}

func (t tArgs) FieldValue(name string) interface{} {
	return getLuaValue(t.RawGetString(name))
}

// PushAsArgs takes the indices of the table and pushes them to the stack,
// removing the table afterwards.
func (t tArgs) PushAsArgs() {
	t.l.Pop(1)
	nt := t.Len()
	for i := 1; i <= nt; i++ {
		t.l.Push(t.RawGetInt(i))
	}
}

// Set the __index metamethod to a table of functions.
func SetIndexFunctions(l *lua.LState, tb *lua.LTable, functions map[string]lua.LGFunction, upValues ...lua.LValue) {
	idx := l.CreateTable(0, len(functions))
	l.SetFuncs(idx, functions, upValues...)
	tb.RawSetString("__index", idx)
}

func NewLuaState(opt rbxmk.Options) *LuaState {
	st := &LuaState{}
	l := lua.NewState(lua.Options{SkipOpenLibs: true})
	st.options = opt
	st.state = l
	st.fileStack = make([]*fileInfo, 0, 1)

	string_Format := func(l *lua.LState) int {
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

	{
		mt := l.NewTypeMetatable(luaTypeInput)
		l.SetFuncs(mt, map[string]lua.LGFunction{
			"__type": func(l *lua.LState) int {
				l.Push(lua.LString(luaTypeInput))
				return 1
			},
			"__tostring": func(l *lua.LState) int {
				l.Push(lua.LString("<input>"))
				return 1
			},
			"__metatable": func(l *lua.LState) int {
				l.Push(lua.LString("the metatable is locked"))
				return 1
			},
		})
	}
	{
		mt := l.NewTypeMetatable(luaTypeOutput)
		l.SetFuncs(mt, map[string]lua.LGFunction{
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
	}

	{
		mt := l.NewTypeMetatable(luaTypeError)
		l.SetFuncs(mt, map[string]lua.LGFunction{
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
	}

	{
		mt := l.NewTypeMetatable(luaTypeAPI)
		l.SetFuncs(mt, map[string]lua.LGFunction{
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
	}

	globalmt := l.NewTable()
	const formatIndex = "format"
	const apiIndex = "api"
	SetIndexFunctions(l, globalmt, map[string]lua.LGFunction{
		"input": func(l *lua.LState) int {
			t := GetArgs(l, 1)

			opt := st.options
			opt.Config.API, _ = t.FieldTyped(apiIndex, luaTypeAPI, true).(*rbxapi.API)

			node := &rbxmk.InputNode{}
			node.Format, _ = t.FieldString(formatIndex, true)
			node.Options = opt

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
		},
		"output": func(l *lua.LState) int {
			t := GetArgs(l, 1)

			opt := st.options
			opt.Config.API, _ = t.FieldTyped(apiIndex, luaTypeAPI, true).(*rbxapi.API)

			node := &rbxmk.OutputNode{}
			node.Format, _ = t.FieldString(formatIndex, true)
			node.Options = opt

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
		},
		"filter": func(l *lua.LState) int {
			t := GetArgs(l, 1)

			opt := st.options
			opt.Config.API, _ = t.FieldTyped(apiIndex, luaTypeAPI, true).(*rbxapi.API)

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
		},
		"map": func(l *lua.LState) int {
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

			return st.mapNodes(inputs, outputs)
		},
		"delete": func(l *lua.LState) int {
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

			return st.mapNodes([]rbxmk.Data{types.Delete{}}, outputs)
		},
		"load": func(l *lua.LState) int {
			t := GetArgs(l, 1)

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
		},
		"error": func(l *lua.LState) int {
			return throwError(l, errors.New(GetArgs(l, 1).IndexString(1, false)))
		},
		"exit": func(l *lua.LState) int {
			t := GetArgs(l, 1)
			v := t.IndexTyped(1, luaTypeError, false)
			err, _ := v.(error)
			panic(exitMarker{err: err})
		},
		"type": func(l *lua.LState) int {
			t := GetArgs(l, 1)
			l.Push(lua.LString(typeOf(l, t.RawGetInt(1))))
			return 1
		},
		"pcall": func(l *lua.LState) int {
			t := GetArgs(l, 1)

			lv := t.RawGetInt(1)
			fn, ok := lv.(*lua.LFunction)
			if !ok {
				t.ErrorIndex(1, "function", lv.Type().String())
			}
			l.Push(fn)

			nt := t.Len()
			for i := 2; i < nt; i++ {
				l.Push(t.RawGetInt(i))
			}
			if err := l.PCall(nt-1, lua.MultRet, nil); err != nil {
				l.Push(lua.LFalse)
				if aerr, ok := err.(*lua.ApiError); ok {
					l.Push(aerr.Object)
				} else {
					l.Push(lua.LString(err.Error()))
				}
				return 2
			}
			l.Insert(lua.LTrue, 1)
			return l.GetTop()
		},
		"getenv": func(l *lua.LState) int {
			t := GetArgs(l, 1)
			value, ok := os.LookupEnv(t.IndexString(1, false))
			if ok {
				l.Push(lua.LString(value))
			} else {
				l.Push(lua.LNil)
			}
			return 1
		},
		"path": func(l *lua.LState) int {
			t := GetArgs(l, 1)
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
		},
		"readdir": func(l *lua.LState) int {
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
		},
		"print": func(l *lua.LState) int {
			t := GetArgs(l, 1)
			nt := t.Len()
			s := make([]interface{}, nt)
			for i := 1; i <= nt; i++ {
				typ := t.TypeOfIndex(i)
				switch typ {
				case "input":
					if typ := reflect.TypeOf(t.IndexValue(i)); typ == nil {
						s[i-1] = "<input(nil)>"
					} else {
						s[i-1] = "<input(" + typ.String() + ")>"
					}
				case "output":
					s[i-1] = "<output>"
				default:
					s[i-1] = t.IndexValue(i)
				}
			}
			fmt.Println(s...)
			return 0
		},
		"sprintf": func(l *lua.LState) int {
			t := GetArgs(l, 1)
			t.PushAsArgs()
			string_Format(l)
			return 1
		},
		"printf": func(l *lua.LState) int {
			t := GetArgs(l, 1)
			t.PushAsArgs()
			string_Format(l)
			s := l.ToString(-1)
			l.Pop(1)
			fmt.Print(s)
			return 0
		},
		"loadapi": func(l *lua.LState) int {
			t := GetArgs(l, 1)
			path := t.IndexString(1, false)
			api, err := rbxmk.LoadAPI(path)
			if err != nil {
				return throwError(l, err)
			}
			return returnTypedValue(l, api, luaTypeAPI)
		},
		"configure": func(l *lua.LState) int {
			t := GetArgs(l, 1)
			if v := t.FieldValue("api"); v != nil {
				st.options.Config.API, _ = v.(*rbxapi.API)
			}
			return 0
		},
	})
	l.SetFuncs(globalmt, map[string]lua.LGFunction{
		"__metatable": func(l *lua.LState) int {
			l.Push(lua.LString("the metatable is locked"))
			return 1
		},
	})
	l.SetMetatable(l.Get(lua.GlobalsIndex), globalmt)
	return st
}

func (st *LuaState) pushFile(fi *fileInfo) error {
	for _, f := range st.fileStack {
		if os.SameFile(fi.fi, f.fi) {
			return fmt.Errorf("\"%s\" is already running", fi.path)
		}
	}
	st.fileStack = append(st.fileStack, fi)
	return nil
}

func (st *LuaState) popFile() {
	if len(st.fileStack) > 0 {
		st.fileStack[len(st.fileStack)-1] = nil
		st.fileStack = st.fileStack[:len(st.fileStack)-1]
	}
}

type LuaSyntaxError string

func (err LuaSyntaxError) Error() string {
	return "syntax error: " + string(err)
}

type LuaError struct {
	Where string
	Err   error
}

func (err LuaError) Error() string {
	if err.Where == "" {
		return err.Err.Error()
	}
	return err.Where + " " + err.Err.Error()
}

func (st *LuaState) DoString(s, name string, args int) (err error) {
	fn, err := st.state.Load(strings.NewReader(s), name)
	if err != nil {
		return err
	}
	st.state.Insert(fn, -args-1)
	return st.state.PCall(args, lua.MultRet, nil)
}

func (st *LuaState) DoFile(fileName string, args int) error {
	fi, err := os.Stat(fileName)
	if err != nil {
		return err
	}
	if err = st.pushFile(&fileInfo{fileName, fi}); err != nil {
		return err
	}

	fn, err := st.state.LoadFile(fileName)
	if err != nil {
		st.popFile()
		return err
	}
	st.state.Insert(fn, -args-1)
	err = st.state.PCall(args, lua.MultRet, nil)
	st.popFile()
	return err
}

func (st *LuaState) DoFileHandle(f *os.File, args int) error {
	fi, err := f.Stat()
	if err != nil {
		return err
	}
	if err = st.pushFile(&fileInfo{f.Name(), fi}); err != nil {
		return err
	}

	fn, err := st.state.Load(f, fi.Name())
	if err != nil {
		st.popFile()
		return err
	}
	st.state.Insert(fn, -args-1)
	err = st.state.PCall(args, lua.MultRet, nil)
	st.popFile()
	return err
}

func (st *LuaState) mapNodes(inputs []rbxmk.Data, outputs []*rbxmk.OutputNode) int {
	for _, input := range inputs {
		for _, output := range outputs {
			if err := output.ResolveReference(input); err != nil {
				return throwError(st.state, err)
			}
		}
	}
	return 0
}
