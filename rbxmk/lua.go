package main

import (
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/yuin/gopher-lua"
	"os"
	"reflect"
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

	// Metatables for custom types.
	l.SetFuncs(l.NewTypeMetatable(luaTypeInput), map[string]lua.LGFunction{
		"__type": func(l *lua.LState) int {
			l.Push(lua.LString(luaTypeInput))
			return 1
		},
		"__tostring": func(l *lua.LState) int {
			if typ := reflect.TypeOf(getLuaValue(l.Get(1))); typ == nil {
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

	// Load and filter libraries.
	libs := []struct {
		name   string
		open   lua.LGFunction
		filter map[lua.LValue]bool
	}{
		{lua.BaseLibName, lua.OpenBase, map[lua.LValue]bool{
			lua.LString("_G"):       true,
			lua.LString("_VERSION"): true,
			lua.LString("assert"):   true,
			lua.LString("error"):    true,
			lua.LString("ipairs"):   true,
			lua.LString("next"):     true,
			lua.LString("pairs"):    true,
			lua.LString("pcall"):    true,
			lua.LString("print"):    true,
			lua.LString("select"):   true,
			lua.LString("tonumber"): true,
			lua.LString("tostring"): true,
			lua.LString("type"):     true,
			lua.LString("unpack"):   true,
			lua.LString("xpcall"):   true,
			// lua.LString("collectgarbage"): true,
			// lua.LString("dofile"):         true,
			// lua.LString("getfenv"):        true,
			// lua.LString("getmetatable"):   true,
			// lua.LString("load"):           true,
			// lua.LString("loadfile"):       true,
			// lua.LString("loadstring"):     true,
			// lua.LString("module"):         true,
			// lua.LString("rawequal"):       true,
			// lua.LString("rawget"):         true,
			// lua.LString("rawset"):         true,
			// lua.LString("require"):        true,
			// lua.LString("setfenv"):        true,
			// lua.LString("setmetatable"):   true,
		}},
		// {lua.CoroutineLibName, lua.OpenCoroutine, map[lua.LValue]bool{
		// 	lua.LString("create"):  true,
		// 	lua.LString("resume"):  true,
		// 	lua.LString("running"): true,
		// 	lua.LString("status"):  true,
		// 	lua.LString("wrap"):    true,
		// 	lua.LString("yield"):   true,
		// }},
		// {lua.DebugLibName, lua.OpenDebug, map[lua.LValue]bool{
		// 	lua.LString("debug"):        true,
		// 	lua.LString("getfenv"):      true,
		// 	lua.LString("gethook"):      true,
		// 	lua.LString("getinfo"):      true,
		// 	lua.LString("getlocal"):     true,
		// 	lua.LString("getmetatable"): true,
		// 	lua.LString("getregistry"):  true,
		// 	lua.LString("getupvalue"):   true,
		// 	lua.LString("setfenv"):      true,
		// 	lua.LString("sethook"):      true,
		// 	lua.LString("setlocal"):     true,
		// 	lua.LString("setmetatable"): true,
		// 	lua.LString("setupvalue"):   true,
		// 	lua.LString("traceback"):    true,
		// }},
		// {lua.IoLibName, lua.OpenIo, map[lua.LValue]bool{
		// 	lua.LString("close"):   true,
		// 	lua.LString("flush"):   true,
		// 	lua.LString("input"):   true,
		// 	lua.LString("lines"):   true,
		// 	lua.LString("open"):    true,
		// 	lua.LString("output"):  true,
		// 	lua.LString("popen"):   true,
		// 	lua.LString("read"):    true,
		// 	lua.LString("stderr"):  true,
		// 	lua.LString("stdin"):   true,
		// 	lua.LString("stdout"):  true,
		// 	lua.LString("tmpfile"): true,
		// 	lua.LString("type"):    true,
		// 	lua.LString("write"):   true,
		// }},
		{lua.MathLibName, lua.OpenMath, map[lua.LValue]bool{
			lua.LString("abs"):        true,
			lua.LString("acos"):       true,
			lua.LString("asin"):       true,
			lua.LString("atan"):       true,
			lua.LString("atan2"):      true,
			lua.LString("ceil"):       true,
			lua.LString("cos"):        true,
			lua.LString("cosh"):       true,
			lua.LString("deg"):        true,
			lua.LString("exp"):        true,
			lua.LString("floor"):      true,
			lua.LString("fmod"):       true,
			lua.LString("frexp"):      true,
			lua.LString("huge"):       true,
			lua.LString("ldexp"):      true,
			lua.LString("log"):        true,
			lua.LString("log10"):      true,
			lua.LString("max"):        true,
			lua.LString("min"):        true,
			lua.LString("modf"):       true,
			lua.LString("pi"):         true,
			lua.LString("pow"):        true,
			lua.LString("rad"):        true,
			lua.LString("random"):     true,
			lua.LString("randomseed"): true,
			lua.LString("sin"):        true,
			lua.LString("sinh"):       true,
			lua.LString("sqrt"):       true,
			lua.LString("tan"):        true,
			lua.LString("tanh"):       true,
		}},
		{lua.OsLibName, lua.OpenOs, map[lua.LValue]bool{
			lua.LString("clock"):    true,
			lua.LString("date"):     true,
			lua.LString("difftime"): true,
			lua.LString("time"):     true,
			// lua.LString("execute"):   true,
			// lua.LString("exit"):      true,
			// lua.LString("getenv"):    true,
			// lua.LString("remove"):    true,
			// lua.LString("rename"):    true,
			// lua.LString("setlocale"): true,
			// lua.LString("tmpname"):   true,
		}},
		// {lua.LoadLibName, lua.OpenPackage, map[lua.LValue]bool{
		// 	lua.LString("cpath"):   true,
		// 	lua.LString("loaded"):  true,
		// 	lua.LString("loaders"): true,
		// 	lua.LString("loadlib"): true,
		// 	lua.LString("path"):    true,
		// 	lua.LString("preload"): true,
		// 	lua.LString("seeall"):  true,
		// }},
		{lua.StringLibName, lua.OpenString, map[lua.LValue]bool{
			lua.LString("byte"):    true,
			lua.LString("char"):    true,
			lua.LString("find"):    true,
			lua.LString("format"):  true,
			lua.LString("gmatch"):  true,
			lua.LString("gsub"):    true,
			lua.LString("len"):     true,
			lua.LString("lower"):   true,
			lua.LString("match"):   true,
			lua.LString("rep"):     true,
			lua.LString("reverse"): true,
			lua.LString("sub"):     true,
			lua.LString("upper"):   true,
			// lua.LString("dump"): true,
		}},
		{lua.TabLibName, lua.OpenTable, map[lua.LValue]bool{
			lua.LString("concat"): true,
			lua.LString("insert"): true,
			lua.LString("maxn"):   true,
			lua.LString("remove"): true,
			lua.LString("sort"):   true,
		}},
		// {lua.ChannelLibName, lua.OpenChannel, map[lua.LValue]bool{
		// 	lua.LString("make"):   true,
		// 	lua.LString("select"): true,
		// }},
		{MainLibName, OpenMain, nil},
	}

	for _, lib := range libs {
		l.Push(l.NewFunction(lib.open))
		// LState.OpenLibs passes the library name as an argument for whatever
		// reason.
		l.Push(lua.LString(lib.name))
		// Pass LuaState as an argument.
		ust := l.NewUserData()
		ust.Value = st
		l.Push(ust)

		if lib.filter == nil {
			l.Call(2, 0)
			continue
		}
		l.Call(2, 1)
		table := l.CheckTable(1)
		l.Pop(1)
		for k, _ := table.Next(lua.LNil); k != lua.LNil; k, _ = table.Next(k) {
			if !lib.filter[k] {
				table.RawSet(k, lua.LNil)
			}
		}
	}

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
