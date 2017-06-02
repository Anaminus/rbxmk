package main

import (
	"errors"
	"fmt"
	"github.com/Shopify/go-lua"
	"os"
	"strings"
)

type LuaState struct {
	options   *Options
	state     *lua.State
	fileStack []os.FileInfo
}

const (
	typeInputNode  = "inputMetatable"
	typeOutputNode = "outputMetatable"
	typeErrorNode  = "errorMetatable"
)

func returnNode(l *lua.State, value interface{}, nodeType string) int {
	l.PushUserData(value)
	lua.SetMetaTableNamed(l, nodeType)
	return 1
}

func returnErrorNode(l *lua.State, err error) int {
	l.PushUserData(err)
	lua.SetMetaTableNamed(l, typeErrorNode)
	return 1
}

func typeOf(l *lua.State, index int) string {
	t := l.TypeOf(index)
	if t == lua.TypeUserData {
		if !l.MetaTable(index) {
			return t.String()
		}
		// +metatable
		l.Field(-1, "__type")
		// metatable, +field
		if l.TypeOf(-1) != lua.TypeFunction {
			l.Pop(2)
			// -metatable, -field
			return t.String()
		}

		l.PushValue(index)
		// metatable, field, +userdata
		l.Call(1, 1)
		// metatable, -field, -userdata, +string
		s, ok := l.ToString(-1)
		l.Pop(2)
		// -metatable, -string
		if ok {
			return s
		}
	}
	return t.String()
}

func mapNodes(l *lua.State, inputs []*Source, outputs []*OutputNode) int {
	// map inputs to outputs, push errors to lua state
	return 0
}

const tableArg = 1

type tArgs struct {
	l *lua.State
}

type exitMarker struct {
	err error
}

func (exitMarker) Error() string {
	return "ExitMarker"
}

func GetArgs(l *lua.State) tArgs {
	t := tArgs{l: l}
	t.Check()
	return t
}

func (t tArgs) Check() {
	if t.l.Top() != 1 {
		lua.Errorf(t.l, "function must have 1 table argument")
	}
	lua.CheckType(t.l, tableArg, lua.TypeTable)
	if t.l.MetaTable(tableArg) {
		t.l.Pop(1)
		lua.Errorf(t.l, "table cannot have metatable")
	}
}

func (t tArgs) Length() int {
	return t.l.RawLength(tableArg)
}

func (t tArgs) FieldError(name string, expected, got string) {
	if got == "" {
		lua.Errorf(t.l, "bad value at field %q: %s expected", name, expected)
	} else {
		lua.Errorf(t.l, "bad value at field %q: %s expected, got %s", name, expected, got)
	}
}

func (t tArgs) IndexError(index int, expected, got string) {
	if got == "" {
		lua.Errorf(t.l, "bad value at index #%d: %s expected", index, expected)
	} else {
		lua.Errorf(t.l, "bad value at index #%d: %s expected, got %s", index, expected, got)
	}
}

func (t tArgs) FieldString(name string, opt bool) (s string, ok bool) {
	t.l.Field(tableArg, name)
	// +field
	s, ok = t.l.ToString(-1)
	if !ok {
		typ := typeOf(t.l, -1)
		if typ != "nil" || !opt {
			t.l.Pop(1)
			// -field
			t.FieldError(name, lua.TypeString.String(), typ)
		}
	}
	t.l.Pop(1)
	// -field
	return s, ok
}

func (t tArgs) IndexString(index int) string {
	t.l.PushInteger(index)
	// +index
	t.l.Table(tableArg)
	// -index, +value
	s, ok := t.l.ToString(-1)
	if !ok {
		typ := typeOf(t.l, -1)
		t.l.Pop(1)
		// -value
		t.IndexError(index, lua.TypeString.String(), typ)
	}
	t.l.Pop(1)
	// -value
	return s
}

func (t tArgs) FieldNode(name string, opt bool) (v interface{}, nodeType string) {
	t.l.Field(tableArg, name)
	// +field
	nodeType = typeOf(t.l, -1)
	switch nodeType {
	case "input", "output", "error":
		v = t.l.ToUserData(-1)
	case "nil":
		if opt {
			nodeType = ""
			goto finish
		}
		fallthrough
	default:
		t.l.Pop(1)
		// -field
		t.FieldError(name, "node", nodeType)
	}
finish:
	t.l.Pop(1)
	// -field
	return v, nodeType
}

func (t tArgs) IndexNode(index int) (v interface{}, nodeType string) {
	t.l.PushInteger(index)
	// +index
	t.l.Table(tableArg)
	// -index, +value
	nodeType = typeOf(t.l, -1)
	switch nodeType {
	case "input", "output", "error":
		v = t.l.ToUserData(-1)
	default:
		t.l.Pop(1)
		// -value
		t.IndexError(index, "node", nodeType)
	}
	t.l.Pop(1)
	// -value
	return v, nodeType
}

func NewLuaState(opt *Options) *LuaState {
	st := &LuaState{}
	l := lua.NewState()
	st.options = opt
	st.state = l
	st.fileStack = make([]os.FileInfo, 0, 1)

	// switch t := l.TypeOf(1); t {
	// case lua.TypeNone:
	// 	// error: invalid type
	// default:
	// 	return
	// case lua.TypeUserData, lua.TypeLightUserData:
	// }

	lua.NewMetaTable(l, "inputMetaTable")
	lua.SetFunctions(l, []lua.RegistryFunction{
		{"__type", func(l *lua.State) int {
			l.PushString("input")
			return 1
		}},
		{"__metatable", func(l *lua.State) int {
			l.PushString("the metatable is locked")
			return 1
		}},
	}, 0)
	l.Pop(1)

	lua.NewMetaTable(l, "outputMetaTable")
	lua.SetFunctions(l, []lua.RegistryFunction{
		{"__type", func(l *lua.State) int {
			l.PushString("output")
			return 1
		}},
		{"__metatable", func(l *lua.State) int {
			l.PushString("the metatable is locked")
			return 1
		}},
	}, 0)
	l.Pop(1)

	lua.NewMetaTable(l, "errorMetaTable")
	lua.SetFunctions(l, []lua.RegistryFunction{
		{"__type", func(l *lua.State) int {
			l.PushString("error")
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

	l.PushGlobalTable()

	const formatIndex = "format"
	lua.SetFunctions(l, []lua.RegistryFunction{
		{"input", func(l *lua.State) int {
			t := GetArgs(l)
			node := &InputNode{}

			node.Format, _ = t.FieldString(formatIndex, false)

			nt := t.Length()
			for i := 1; i <= nt; i++ {
				node.Reference = append(node.Reference, t.IndexString(i))
			}

			src, err := node.ResolveReference(st.options)
			if err != nil {
				return returnNode(l, err, typeErrorNode)
			}

			return returnNode(l, src, typeInputNode)
		}},
		{"output", func(l *lua.State) int {
			t := GetArgs(l)
			node := &OutputNode{}

			node.Format, _ = t.FieldString(formatIndex, false)

			nt := t.Length()
			for i := 1; i <= nt; i++ {
				node.Reference = append(node.Reference, t.IndexString(i))
			}

			return returnNode(l, node, typeOutputNode)
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

			// lookup filter
			_ = filterName

			return lua.MultipleReturns
		}},
		{"map", func(l *lua.State) int {
			t := GetArgs(l)

			inputs := make([]*Source, 1)
			outputs := make([]*OutputNode, 1)

			nt := t.Length()
			for i := 1; i <= nt; i++ {
				switch v, typ := t.IndexNode(i); typ {
				case "input":
					inputs = append(inputs, v.(*Source))
				case "output":
					outputs = append(outputs, v.(*OutputNode))
				default:
					// error node
				}
			}

			return mapNodes(l, inputs, outputs)
		}},
		{"load", func(l *lua.State) int {
			t := GetArgs(l)

			fileName := t.IndexString(1)
			fi, err := os.Stat(fileName)
			if err != nil {
				return returnErrorNode(l, err)
			}
			if err = st.pushFile(fi); err != nil {
				return returnErrorNode(l, err)
			}

			// Load file as function.
			err = lua.LoadFile(l, fileName, "")
			if err != nil {
				st.popFile()
				return returnErrorNode(l, err)
			}
			// +function

			// Push extra arguments as arguments to loaded function.
			nt := t.Length()
			for i := 2; i <= nt; i++ {
				l.PushInteger(i)
				// function, ..., +int
				l.Table(tableArg)
				// function, ..., -int, +arg
			}
			// function, +args...

			// Call loaded function.
			err = l.ProtectedCall(nt-1, lua.MultipleReturns, 0)
			// -function, -args..., +returns...
			st.popFile()
			if err != nil {
				return returnErrorNode(l, err)
			}
			return lua.MultipleReturns
		}},
		{"error", func(l *lua.State) int {
			return returnErrorNode(l, errors.New(lua.CheckString(l, 1)))
		}},
		{"exit", func(l *lua.State) int {
			v := l.ToUserData(1)
			err, _ := v.(error)
			panic(exitMarker{err: err})
		}},
		{"type", func(l *lua.State) int {
			l.PushString(typeOf(l, 1))
			return 1
		}},
	}, 0)
	l.Pop(1)
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

func (st *LuaState) DoFileHandle(f *os.File) error {
	fi, err := f.Stat()
	if err != nil {
		return err
	}
	if err = st.pushFile(fi); err != nil {
		return err
	}
	if err := st.state.Load(f, fi.Name(), ""); err != nil {
		st.popFile()
		return err
	}
	err = st.state.ProtectedCall(0, lua.MultipleReturns, 0)
	st.popFile()
	return err
}

func (st *LuaState) DoFile(fileName string) error {
	fi, err := os.Stat(fileName)
	if err != nil {
		return err
	}
	if err = st.pushFile(fi); err != nil {
		return err
	}
	err = lua.DoFile(st.state, fileName)
	st.popFile()
	if err != nil {
		return err
	}
	return nil
}

func (st *LuaState) DoString(s, name string) error {
	return st.state.Load(strings.NewReader(s), name, "")
}