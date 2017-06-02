package main

import (
	"github.com/Shopify/go-lua"
	"os"
	"testing"
)

func TestLuaStateInit(t *testing.T) {
	st := NewLuaState(&Options{})
	st.state.PushGoFunction(func(l *lua.State) int {
		typ := l.TypeOf(1)
		l.PushString(typ.String())
		return 1
	})
	st.state.SetGlobal("_TYPE")

	assertFunc := func(name string) {
		err := st.DoString(`return _TYPE(`+name+`) == "function"`, "test func "+name)
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
	st := NewLuaState(&Options{})
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

	st.state.PushGoFunction(func(l *lua.State) int {
		l.PushString("hello world!")
		return 1
	})
	st.state.SetGlobal("_HELLO")
	err = st.DoString(`return _HELLO() == "hello world!"`, "test valid")
	if err != nil {
		t.Errorf("unexpected error: %q", err)
	}
	if b, ok := st.state.ToValue(-1).(bool); !b || !ok {
		t.Errorf("assertion failed")
	}
}

func TestLuaStateDoFile(t *testing.T) {
	st := NewLuaState(&Options{})
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
	st.state.PushGoFunction(func(l *lua.State) int {
		l.PushString("hello world!")
		return 1
	})
	st.state.SetGlobal("_HELLO")
	err = st.DoFile("testdata/test-valid.lua")
	if err != nil {
		t.Errorf("unexpected error: %q", err)
	}
	if b, ok := st.state.ToValue(-1).(bool); !b || !ok {
		t.Errorf("assertion failed")
	}
}

func TestLuaStateDoFileHandle(t *testing.T) {
	st := NewLuaState(&Options{})
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

	st.state.PushGoFunction(func(l *lua.State) int {
		l.PushString("hello world!")
		return 1
	})
	st.state.SetGlobal("_HELLO")

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
