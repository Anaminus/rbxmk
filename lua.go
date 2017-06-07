package main

import (
	"errors"
	"fmt"
	"github.com/Shopify/go-lua"
	"os"
	"strings"
)

/*
Stack Annotations:
    separate values : a, b
	insert          : +a
	remove          : -a
	insert from top : >a
	replace         : a>b
	group of values : a...

	stack   : a, b, c, d
	push    : a, b, c, d, +e
	pop     : a, b, c, d, -e
	insert  : >d, a, b, c
	replace : d>c, a, b
	copy    : c, a>c, b
	remove  : c, -c, b
*/

type LuaState struct {
	options   *Options
	state     *lua.State
	fileStack []os.FileInfo
}

const (
	luaTypeInput  = "input"
	luaTypeOutput = "output"
	luaTypeError  = "error"
)

func returnNode(l *lua.State, value interface{}, nodeType string) int {
	l.PushUserData(value)
	lua.SetMetaTableNamed(l, nodeType)
	return 1
}

func throwError(l *lua.State, err error) int {
	l.PushString(err.Error())
	l.Error()
	return 0
}

func typeOf(l *lua.State, index int) string {
	t := l.TypeOf(index)
	if t == lua.TypeUserData && lua.CallMeta(l, index, "__type") {
		s, ok := l.ToString(-1)
		l.Pop(1)
		if ok {
			return s
		}
	}
	return t.String()
}

const tableArg = 1
const tableMethodArg = 1

type tArgs struct {
	l   *lua.State
	off int
}

type exitMarker struct {
	err error
}

func (exitMarker) Error() string {
	return "ExitMarker"
}

func GetArgs(l *lua.State) tArgs {
	t := tArgs{l: l, off: tableArg}
	t.Check()
	return t
}

func GetMethodArgs(l *lua.State) tArgs {
	t := tArgs{l: l, off: tableMethodArg}
	t.Check()
	return t
}

func (t tArgs) Check() {
	if t.l.Top() != 1 || typeOf(t.l, t.off) != "table" {
		lua.Errorf(t.l, "function must have 1 table argument")
	}
	if t.l.MetaTable(t.off) {
		t.l.Pop(1)
		lua.Errorf(t.l, "table cannot have metatable")
	}
}

func (t tArgs) Length() int {
	return t.l.RawLength(t.off)
}

func luaErrorf(l *lua.State, format string, a ...interface{}) {
	lua.Where(l, 1)
	l.PushString(fmt.Sprintf(format, a...))
	l.Concat(2)
	l.Error()
}

func (t tArgs) FieldError(name string, expected, got string) {
	if got == "" {
		luaErrorf(t.l, "bad value at field %q: %s expected", name, expected)
	} else {
		luaErrorf(t.l, "bad value at field %q: %s expected, got %s", name, expected, got)
	}
}

func (t tArgs) IndexError(index int, expected, got string) {
	if got == "" {
		luaErrorf(t.l, "bad value at index #%d: %s expected", index, expected)
	} else {
		luaErrorf(t.l, "bad value at index #%d: %s expected, got %s", index, expected, got)
	}
}

func (t tArgs) FieldString(name string, opt bool) (s string, ok bool) {
	t.l.Field(t.off, name) // +field
	s, ok = t.l.ToString(-1)
	if !ok {
		typ := typeOf(t.l, -1)
		if typ != "nil" || !opt {
			t.l.Pop(1) // -field
			t.FieldError(name, lua.TypeString.String(), typ)
		}
	}
	t.l.Pop(1) // -field
	return s, ok
}

func (t tArgs) IndexString(index int) string {
	t.l.PushInteger(index) // +index
	t.l.Table(t.off)       // -index, +value
	s, ok := t.l.ToString(-1)
	if !ok {
		typ := typeOf(t.l, -1)
		t.l.Pop(1) // -value
		t.IndexError(index, lua.TypeString.String(), typ)
	}
	t.l.Pop(1) // -value
	return s
}

func (t tArgs) FieldNode(name string, opt bool) (v interface{}, nodeType string) {
	t.l.Field(t.off, name) // +field
	nodeType = typeOf(t.l, -1)
	switch nodeType {
	case luaTypeInput, luaTypeOutput:
		v = t.l.ToUserData(-1)
	case "nil":
		if opt {
			nodeType = ""
			goto finish
		}
		fallthrough
	default:
		t.l.Pop(1) // -field
		t.FieldError(name, "node", nodeType)
	}
finish:
	t.l.Pop(1) // -field
	return v, nodeType
}

func (t tArgs) IndexNode(index int) (v interface{}, nodeType string) {
	t.l.PushInteger(index) // +index
	t.l.Table(t.off)       // -index, +value
	nodeType = typeOf(t.l, -1)
	switch nodeType {
	case luaTypeInput, luaTypeOutput:
		v = t.l.ToUserData(-1)
	default:
		t.l.Pop(1) // -value
		t.IndexError(index, "node", nodeType)
	}
	t.l.Pop(1) // -value
	return v, nodeType
}

func (t tArgs) IndexValue(index int) interface{} {
	t.l.PushInteger(index)  // +index
	t.l.Table(t.off)        // -index, +value
	v := t.l.ToUserData(-1) // value
	t.l.Pop(1)              // -value
	return v
}

func (t tArgs) FieldValue(name string) interface{} {
	t.l.Field(t.off, name)  // +field
	v := t.l.ToUserData(-1) // field
	t.l.Pop(1)              // -field
	return v
}

// PushAsArgs takes the indices of the table and pushes them to the stack,
// removing the table afterwards.
func (t tArgs) PushAsArgs() {
	nt := t.Length()
	for i := 1; i <= nt; i++ {
		t.l.PushInteger(i)
		t.l.Table(t.off)
	}
	// table, args...
	t.l.Remove(t.off) // -table, args...
}

// Set the __index metamethod to a table of functions.
func SetIndexFunctions(l *lua.State, functions []lua.RegistryFunction, upValueCount uint8) {
	uvCount := int(upValueCount)
	lua.CheckStackWithMessage(l, uvCount, "too many upvalues")
	l.CreateTable(0, len(functions)) // metatable, up..., +table
	l.Insert(-(uvCount + 1))         // metatable, >table, up...
	for _, r := range functions {
		for i := 0; i < uvCount; i++ {
			l.PushValue(-uvCount)
		} // metatable, table, up..., +up...
		l.PushGoClosure(r.Function, upValueCount) // metatable, table, up..., +func, -up...
		l.SetField(-(uvCount + 2), r.Name)        // metatable, table, up..., -func
	} // metatable, table, up...
	l.Pop(uvCount)          // metatable, table, -up...
	l.PushString("__index") // metatable, table, +index
	l.Insert(-2)            // metatable, >index, table
	l.SetTable(-3)          // metatable, -index, -table
}

func NewLuaState(opt *Options) *LuaState {
	st := &LuaState{}
	l := lua.NewState()
	st.options = opt
	st.state = l
	st.fileStack = make([]os.FileInfo, 0, 1)

	lua.NewMetaTable(l, luaTypeInput)
	SetIndexFunctions(l, []lua.RegistryFunction{
		{"CheckInstance", func(l *lua.State) int {
			src := l.ToUserData(1).(*Source)
			t := GetMethodArgs(l)

			nt := t.Length()
			ref := make([]string, nt)
			for i := 1; i <= nt; i++ {
				ref[i-1] = t.IndexString(i)
			}

			var err error
			if src, ref, err = DrillInputInstance(st.options, src, ref); err != nil && err != EOD {
				l.PushBoolean(false)
				return 1
			}
			if src, ref, err = DrillInputProperty(st.options, src, ref); err != nil && err != EOD {
				l.PushBoolean(false)
				return 1
			}
			l.PushBoolean(true)
			return 1
		}},
		{"CheckProperty", func(l *lua.State) int {
			src := l.ToUserData(1).(*Source)
			t := GetMethodArgs(l)
			name := t.IndexString(1)
			_, exists := src.Properties[name]
			l.PushBoolean(exists)
			return 1
		}},
	}, 0)
	lua.SetFunctions(l, []lua.RegistryFunction{
		{"__type", func(l *lua.State) int {
			l.PushString(luaTypeInput)
			return 1
		}},
		{"__metatable", func(l *lua.State) int {
			l.PushString("the metatable is locked")
			return 1
		}},
	}, 0)
	l.Pop(1)

	lua.NewMetaTable(l, luaTypeOutput)
	lua.SetFunctions(l, []lua.RegistryFunction{
		{"__type", func(l *lua.State) int {
			l.PushString(luaTypeOutput)
			return 1
		}},
		{"__metatable", func(l *lua.State) int {
			l.PushString("the metatable is locked")
			return 1
		}},
	}, 0)
	l.Pop(1)

	lua.NewMetaTable(l, luaTypeError)
	lua.SetFunctions(l, []lua.RegistryFunction{
		{"__type", func(l *lua.State) int {
			l.PushString(luaTypeError)
			return 1
		}},
		{"__tostring", func(l *lua.State) int {
			err, ok := l.ToUserData(1).(error)
			if ok {
				l.PushString(err.Error())
			} else {
				l.PushString("<error>")
			}
			return 1
		}},
		{"__metatable", func(l *lua.State) int {
			l.PushString("the metatable is locked")
			return 1
		}},
	}, 0)
	l.Pop(1)

	l.PushGlobalTable()                    // +global
	lua.NewMetaTable(l, "globalMetatable") // global, +metatable

	const formatIndex = "format"
	SetIndexFunctions(l, []lua.RegistryFunction{
		{"input", func(l *lua.State) int {
			t := GetArgs(l)
			node := &InputNode{}

			node.Format, _ = t.FieldString(formatIndex, true)

			nt := t.Length()
			if nt == 0 {
				throwError(l, errors.New("at least 1 reference argument is required"))
			}
			for i := 1; i <= nt; i++ {
				node.Reference = append(node.Reference, t.IndexString(i))
			}

			src, err := node.ResolveReference(st.options)
			if err != nil {
				return throwError(l, err)
			}

			return returnNode(l, src, luaTypeInput)
		}},
		{"output", func(l *lua.State) int {
			t := GetArgs(l)
			node := &OutputNode{}

			node.Format, _ = t.FieldString(formatIndex, true)

			nt := t.Length()
			if nt == 0 {
				throwError(l, errors.New("at least 1 reference argument is required"))
			}
			for i := 1; i <= nt; i++ {
				node.Reference = append(node.Reference, t.IndexString(i))
			}

			return returnNode(l, node, luaTypeOutput)
		}},
		{"filter", func(l *lua.State) int {
			t := GetArgs(l)

			const filterNameIndex = "name"
			var i int = 1
			filterName, ok := t.FieldString(filterNameIndex, true)
			if !ok {
				filterName = t.IndexString(i)
				i = 2
			}

			filterFunc, exists := registeredFilters[filterName]
			if !exists {
				return throwError(l, fmt.Errorf("unknown filter %q", filterName))
			}

			nt := t.Length()
			arguments := make([]interface{}, nt-i+1)
			for o := i; i <= nt; i++ {
				arguments[i-o] = t.IndexValue(i)
			}

			results, err := CallFilter(filterFunc, arguments...)
			if err != nil {
				return throwError(l, err)
			}

			for _, result := range results {
				switch v := result.(type) {
				case bool:
					l.PushBoolean(v)
				case lua.Function:
					l.PushGoFunction(v)
				case int:
					l.PushInteger(v)
				case float64:
					l.PushNumber(v)
				case string:
					l.PushString(v)
				case uint:
					l.PushUnsigned(v)
				case *Source:
					l.PushUserData(v)
					lua.SetMetaTableNamed(l, luaTypeInput)
				case *OutputNode:
					l.PushUserData(v)
					lua.SetMetaTableNamed(l, luaTypeOutput)
				default:
					l.PushNil()
				}
			}
			return len(results)
		}},
		{"map", func(l *lua.State) int {
			t := GetArgs(l)

			inputs := make([]*Source, 1)
			outputs := make([]*OutputNode, 1)

			nt := t.Length()
			for i := 1; i <= nt; i++ {
				switch v, typ := t.IndexNode(i); typ {
				case luaTypeInput:
					inputs = append(inputs, v.(*Source))
				case luaTypeOutput:
					outputs = append(outputs, v.(*OutputNode))
				}
			}

			return st.mapNodes(inputs, outputs)
		}},
		{"load", func(l *lua.State) int {
			t := GetArgs(l)

			fileName := t.IndexString(1)
			fi, err := os.Stat(fileName)
			if err != nil {
				return throwError(l, err)
			}
			if err = st.pushFile(fi); err != nil {
				return throwError(l, err)
			}

			// Load file as function.
			if err = lua.LoadFile(l, fileName, ""); err != nil {
				st.popFile()
				return throwError(l, err)
			}
			// +function

			// Push extra arguments as arguments to loaded function.
			nt := t.Length()
			for i := 2; i <= nt; i++ {
				l.PushInteger(i)  // function, ..., +int
				l.Table(tableArg) // function, ..., -int, +arg
			}
			// function, +args...

			// Call loaded function.
			err = l.ProtectedCall(nt-1, lua.MultipleReturns, 0) // -function, -args..., +returns...
			st.popFile()
			if err != nil {
				return throwError(l, err)
			}
			return lua.MultipleReturns
		}},
		{"error", func(l *lua.State) int {
			return throwError(l, errors.New(lua.CheckString(l, 1)))
		}},
		{"exit", func(l *lua.State) int {
			t := GetArgs(l)
			var err error
			if v, typ := t.IndexNode(1); typ == "error" {
				err, _ = v.(error)
			}
			panic(exitMarker{err: err})
		}},
		{"type", func(l *lua.State) int {
			GetArgs(l)
			l.PushInteger(1)
			l.Table(tableArg)
			typ := typeOf(l, -1)
			l.Pop(1)
			l.PushString(typ)
			return 1
		}},
		{"pcall", func(l *lua.State) int {
			finishPCall := func(l *lua.State, status bool) int {
				// nil, results...
				if !l.CheckStack(1) {
					l.SetTop(0)                    // -nil, -results...
					l.PushBoolean(false)           // +false
					l.PushString("stack overflow") // false, +msg
					return 2
				}
				l.PushBoolean(status) // nil, results..., +status
				l.Replace(1)          // nil>status, results...
				return l.Top()
			}

			t := GetArgs(l)    // table
			t.PushAsArgs()     // -table, +func, +args...
			lua.CheckAny(l, 1) // func, args...
			l.PushNil()        // func, args..., +nil
			l.Insert(1)        // >nil, func, args...
			status := nil == l.ProtectedCallWithContinuation(l.Top()-2, lua.MultipleReturns, 0, 0, func(l *lua.State) int {
				_, shouldYield, _ := l.Context()
				return finishPCall(l, shouldYield)
			})
			// nil, -func, -args..., +results...
			return finishPCall(l, status) // status, results...
		}},
		{"getenv", func(l *lua.State) int {
			t := GetArgs(l)
			value, ok := os.LookupEnv(t.IndexString(1))
			if ok {
				l.PushString(value)
			} else {
				l.PushNil()
			}
			return 1
		}},
	}, 0)
	lua.SetFunctions(l, []lua.RegistryFunction{
		{"__metatable", func(l *lua.State) int {
			l.PushString("the metatable is locked")
			return 1
		}},
	}, 0)
	l.SetMetaTable(-2) // global, -metatable
	l.Pop(1)           // -global
	return st
}

func (st *LuaState) pushFile(fi os.FileInfo) error {
	for _, f := range st.fileStack {
		if os.SameFile(fi, f) {
			return fmt.Errorf("cannot load file %q: file is already running", fi.Name())
		}
	}
	st.fileStack = append(st.fileStack, fi)
	return nil
}

func (st *LuaState) popFile() {
	st.fileStack = st.fileStack[:len(st.fileStack)-1]
}

type LuaSyntaxError string

func (err LuaSyntaxError) Error() string {
	return "syntax error: " + string(err)
}

func (st *LuaState) DoString(s, name string, args int) (err error) {
	if err = st.state.Load(strings.NewReader(s), name, ""); err != nil {
		if err == lua.SyntaxError {
			return LuaSyntaxError(fmt.Sprintf("%s", st.state.ToValue(-1)))
		}
		return err
	} // args..., +func
	st.state.Insert(-args - 1) // >func, args...
	if err = st.state.ProtectedCall(args, lua.MultipleReturns, 0); err != nil {
		return err
	} // +results..., -func, -args...
	return nil
}

func (st *LuaState) DoFile(fileName string, args int) error {
	fi, err := os.Stat(fileName)
	if err != nil {
		return err
	}
	if err = st.pushFile(fi); err != nil {
		return err
	}
	if err := lua.LoadFile(st.state, fileName, ""); err != nil {
		st.popFile()
		if err == lua.SyntaxError {
			return LuaSyntaxError(fmt.Sprintf("%s", st.state.ToValue(-1)))
		}
		return err
	} // args..., +func
	st.state.Insert(-args - 1)                                 // >func, args...
	err = st.state.ProtectedCall(args, lua.MultipleReturns, 0) // +results..., -func, -args...
	st.popFile()
	return err
}

func (st *LuaState) DoFileHandle(f *os.File, args int) error {
	fi, err := f.Stat()
	if err != nil {
		return err
	}
	if err = st.pushFile(fi); err != nil {
		return err
	}
	if err = st.state.Load(f, fi.Name(), ""); err != nil {
		st.popFile()
		if err == lua.SyntaxError {
			return LuaSyntaxError(fmt.Sprintf("%s", st.state.ToValue(-1)))
		}
		return err
	} // args..., +func
	st.state.Insert(-args - 1)                                 // >func, args...
	err = st.state.ProtectedCall(args, lua.MultipleReturns, 0) // +results..., -func, -args...
	st.popFile()
	return err
}

func (st *LuaState) mapNodes(inputs []*Source, outputs []*OutputNode) int {
	for _, input := range inputs {
		for _, output := range outputs {
			if err := output.ResolveReference(st.options, input); err != nil {
				return throwError(st.state, err)
			}
		}
	}
	return 0
}
