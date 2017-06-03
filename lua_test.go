package main

import (
	"fmt"
	"github.com/Shopify/go-lua"
	"os"
	"strings"
	"testing"
)

func init() {
	RegisterInputScheme("test", func(opt *Options, node *InputNode, ref string) (src *Source, err error) {
		switch ref {
		case "":
			return &Source{}, nil
		}
		return nil, fmt.Errorf("unknown ref string %q", ref)
	})
}

func InitTestLib(opt *Options) *LuaState {
	st := NewLuaState(opt)
	st.state.PushGlobalTable()
	lua.SetFunctions(st.state, []lua.RegistryFunction{
		{"_type", func(l *lua.State) int {
			typ := l.TypeOf(1)
			l.PushString(typ.String())
			return 1
		}},
		{"_hello", func(l *lua.State) int {
			l.PushString("hello world!")
			return 1
		}},
		{"_assertInput", func(l *lua.State) int {
			if typeOf(l, 1) != "input" {
				l.PushString("_assertInput: value is not input type")
				l.Error()
			}
			if _, ok := l.ToUserData(1).(*Source); !ok {
				l.PushString("_assertInput: value is not *Source")
				l.Error()
			}
			return 0
		}},
		{"_assertOutput", func(l *lua.State) int {
			if typeOf(l, 1) != "output" {
				l.PushString("_assertOutput: value is not output type")
				l.Error()
			}
			if _, ok := l.ToUserData(1).(*OutputNode); !ok {
				l.PushString("_assertOutput: value is not *OutputNode")
				l.Error()
			}
			return 0
		}},
	}, 0)
	st.state.Pop(1)
	return st
}

func TestLuaStateInit(t *testing.T) {
	st := InitTestLib(nil)

	assertFunc := func(name string) {
		err := st.DoString(`return _type(`+name+`) == "function"`, "test func "+name)
		if err != nil {
			t.Errorf("error testing function %s: %s", name, err)
		}
		if !st.state.ToBoolean(1) {
			t.Errorf("missing function %q: %s", name, err)
		}
		st.state.Pop(1)
	}

	assertFunc("input")
	assertFunc("output")
	assertFunc("map")
	assertFunc("filter")
	assertFunc("load")
	assertFunc("error")
	assertFunc("exit")
	assertFunc("type")
	assertFunc("pcall")
}

func TestLuaStateDoString(t *testing.T) {
	st := InitTestLib(nil)
	var err error

	err = st.DoString(`func`, "test syntax")
	if _, ok := err.(LuaSyntaxError); !ok {
		t.Errorf("expected LuaSyntaxError (got %q)", err)
	}
	st.state.Pop(1)

	err = st.DoString(`_DOES_NOT_EXIST()`, "test runtime")
	if _, ok := err.(lua.RuntimeError); !ok {
		t.Errorf("expected lua.RuntimeError (got %q)", err)
	}
	st.state.Pop(1)

	err = st.DoString(`return _hello() == "hello world!"`, "test valid")
	if err != nil {
		t.Errorf("unexpected error: %q", err)
	}
	if b, ok := st.state.ToValue(-1).(bool); !b || !ok {
		t.Errorf("assertion failed")
	}
}

func TestLuaStateDoFile(t *testing.T) {
	st := InitTestLib(nil)
	var err error

	// Test syntax error.
	err = st.DoFile("testdata/test-syntax.lua")
	if _, ok := err.(LuaSyntaxError); !ok {
		t.Errorf("expected LuaSyntaxError (got %q)", err)
	}
	st.state.Pop(1)

	// Test runtime error.
	err = st.DoFile("testdata/test-runtime.lua")
	if _, ok := err.(lua.RuntimeError); !ok {
		t.Errorf("expected lua.RuntimeError (got %q)", err)
	}
	st.state.Pop(1)

	// Test no error.
	err = st.DoFile("testdata/test-valid.lua")
	if err != nil {
		t.Errorf("unexpected error: %q", err)
	}
	if b, ok := st.state.ToValue(-1).(bool); !b || !ok {
		t.Errorf("assertion failed")
	}
}

func TestLuaStateDoFileHandle(t *testing.T) {
	st := InitTestLib(nil)
	var f *os.File
	var err error

	f, err = os.Open("testdata/test-syntax.lua")
	if err != nil {
		t.Fatal(err)
	}

	// Test syntax error.
	err = st.DoFileHandle(f)
	f.Close()
	if _, ok := err.(LuaSyntaxError); !ok {
		t.Errorf("expected LuaSyntaxError (got %q)", err)
	}
	st.state.Pop(1)

	f, err = os.Open("testdata/test-runtime.lua")
	if err != nil {
		t.Fatal(err)
	}

	// Test runtime error.
	err = st.DoFileHandle(f)
	f.Close()
	if _, ok := err.(lua.RuntimeError); !ok {
		t.Errorf("expected lua.RuntimeError (got %q)", err)
	}
	st.state.Pop(1)

	// Test no error.
	f, err = os.Open("testdata/test-valid.lua")
	if err != nil {
		t.Fatal(err)
	}

	err = st.DoFileHandle(f)
	f.Close()
	if err != nil {
		t.Errorf("unexpected error: %q", err)
	}
	if b, ok := st.state.ToValue(-1).(bool); !b || !ok {
		t.Errorf("assertion failed")
	}
}

func mustNotError(t *testing.T, st *LuaState, name, script string) {
	if err := st.DoString(script, ""); err != nil {
		t.Errorf("%s: unexpected error: %s", name, err)
	}
}
func mustError(t *testing.T, st *LuaState, name, msg, script string) {
	if err := st.DoString(script, ""); err == nil {
		t.Errorf("%s: expected error", name)
	} else if msg != "" {
		if !strings.Contains(err.Error(), msg) {
			t.Errorf("%s: unexpected error: %s", name, err)
		}
	}
}

func testTArgs(t *testing.T, st *LuaState, name string) {
	mustError(t, st, name+": no args",
		`function must have 1 table argument`,
		name+`()`,
	)
	mustError(t, st, name+": non-table arg",
		`function must have 1 table argument`,
		name+`(true)`,
	)
	mustError(t, st, name+": 2 non-table arg",
		`function must have 1 table argument`,
		name+`(true,true)`,
	)
	mustError(t, st, name+": multi-arg; table, bool",
		`function must have 1 table argument`,
		name+`({},true)`,
	)
	mustError(t, st, name+": multi-arg; bool, table",
		`function must have 1 table argument`,
		name+`(true,{})`,
	)
	mustError(t, st, name+": multi-arg: table, table",
		`function must have 1 table argument`,
		name+`({},{})`,
	)
}

func TestLuaFuncInput(t *testing.T) {
	st := InitTestLib(nil)
	funcName := "input"
	testTArgs(t, st, funcName)
	mustError(t, st, funcName+": empty targs",
		`at least 1 reference argument is required`,
		funcName+`{}`,
	)
	mustError(t, st, funcName+": non-string index targ 1",
		`bad value at index #1: string expected, got boolean`,
		funcName+`{true}`,
	)
	mustError(t, st, funcName+": non-string index targ 2",
		`bad value at index #2: string expected, got boolean`,
		funcName+`{"",true}`,
	)
	mustError(t, st, funcName+": non-string field targ",
		`bad value at field "format": string expected, got boolean`,
		funcName+`{format=true,""}`,
	)
	mustError(t, st, funcName+": field, no index",
		`at least 1 reference argument is required`,
		funcName+`{format=""}`,
	)

	mustNotError(t, st, funcName+": expect node",
		`_assertInput(input{"test://"})`,
	)
}
func TestLuaFuncOutput(t *testing.T) {
	st := InitTestLib(nil)
	funcName := "output"
	testTArgs(t, st, funcName)
	mustError(t, st, funcName+": empty targs",
		`at least 1 reference argument is required`,
		funcName+`{}`,
	)
	mustError(t, st, funcName+": non-string index targ 1",
		`bad value at index #1: string expected, got boolean`,
		funcName+`{true}`,
	)
	mustError(t, st, funcName+": non-string index targ 2",
		`bad value at index #2: string expected, got boolean`,
		funcName+`{"",true}`,
	)
	mustError(t, st, funcName+": non-string field targ",
		`bad value at field "format": string expected, got boolean`,
		funcName+`{format=true,""}`,
	)
	mustError(t, st, funcName+": field, no index",
		`at least 1 reference argument is required`,
		funcName+`{format=""}`,
	)

	mustNotError(t, st, funcName+": expect node",
		`_assertOutput(output{"test://"})`,
	)
}

func TestLuaFuncFilter(t *testing.T) {}
func TestLuaFuncMap(t *testing.T)    {}

func TestLuaFuncLoad(t *testing.T) {
	st := InitTestLib(nil)

	if err := st.DoFile("testdata/test-loop.lua"); err == nil {
		t.Fatalf("expected error")
	}
}

func TestLuaFuncError(t *testing.T) {}
func TestLuaFuncExit(t *testing.T)  {}
func TestLuaFuncType(t *testing.T)  {}
func TestLuaFuncPCall(t *testing.T) {}
