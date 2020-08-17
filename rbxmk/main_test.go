package main

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
	lua "github.com/yuin/gopher-lua"
)

const testdata = "testdata"

var scriptArguments = [...]string{
	"rbxmk_test",
	"-",
	"true",
	"false",
	"nil",
	"42",
	"3.141592653589793",
	"-1e-8",
	"hello, world!",
	"hello\000world!",
}

type dummyFile struct {
	r    io.Reader
	info *dummyInfo
}

func (d *dummyFile) Name() string               { return "test" }
func (d *dummyFile) Stat() (os.FileInfo, error) { return d.info, nil }
func (d *dummyFile) Read(b []byte) (int, error) { return d.r.Read(b) }
func (d *dummyFile) Write([]byte) (int, error)  { return 0, nil }

type dummyInfo struct {
	name  string
	size  int64
	mode  os.FileMode
	time  time.Time
	isdir bool
}

func (d *dummyInfo) Name() string       { return d.name }
func (d *dummyInfo) Size() int64        { return d.size }
func (d *dummyInfo) Mode() os.FileMode  { return d.mode }
func (d *dummyInfo) ModTime() time.Time { return d.time }
func (d *dummyInfo) IsDir() bool        { return d.isdir }
func (d *dummyInfo) Sys() interface{}   { return d }

func initMain(s rbxmk.State, t *testing.T) {
	T := s.L.CreateTable(0, 2)

	// Pass makes a positive assertion. The first argument is a string that
	// describes the assertion, which is included with an emitted error. If the
	// second argument is a non-function, then an error is emitted if the value
	// is falsy. Otherwise, the function is called. If the call errors or
	// returns a falsy value, then an error is emitted. Returning no value
	// counts as truthy. The error is emitted to testing.T, but does not cause a
	// Lua error to be thrown.
	T.RawSetString("Pass", s.L.NewFunction(func(l *lua.LState) int {
		msg := l.CheckString(1)
		v := l.CheckAny(2)
		switch v := v.(type) {
		case *lua.LFunction:
			n := s.L.GetTop()
			s.L.Push(v)
			if err := s.L.PCall(0, lua.MultRet, nil); err != nil {
				t.Errorf("%s: %s", msg, err.Error())
				return 0
			}
			if s.L.GetTop() > n {
				if !s.L.ToBool(n + 1) {
					t.Errorf(msg)
					return 0
				}
			}
		default:
			if !lua.LVAsBool(v) {
				t.Errorf(msg)
				return 0
			}
		}
		return 0
	}))

	// Fail makes a negative assertion. The first argument is a string that
	// describes the assertion, which is included with an emitted error. If the
	// second argument is a non-function, then an error is emitted if the value
	// is truthy. Otherwise, the function is called. If the call does not error
	// or returns a truthy value, then an error is emitted. Returning no value
	// counts as falsy. The error is emitted to testing.T, but does not cause a
	// Lua error to be thrown.
	T.RawSetString("Fail", s.L.NewFunction(func(l *lua.LState) int {
		msg := l.CheckString(1)
		v := l.CheckAny(2)
		switch v := v.(type) {
		case *lua.LFunction:
			n := s.L.GetTop()
			s.L.Push(v)
			if err := s.L.PCall(0, lua.MultRet, nil); err == nil {
				if s.L.GetTop() > n {
					if s.L.ToBool(n + 1) {
						t.Errorf(msg)
					}
					return 0
				}
				t.Errorf(msg)
				return 0
			} else if lua.LVAsBool(s.L.GetGlobal("SHOW_ERRORS")) {
				t.Logf("ERROR: %s\n", err)
			}
		default:
			if lua.LVAsBool(v) {
				t.Errorf(msg)
				return 0
			}
		}
		return 0
	}))

	// GC runs the garbage collector.
	T.RawSetString("GC", s.L.NewFunction(func(l *lua.LState) int {
		runtime.GC()
		return 0
	}))

	// UserDataCacheLen returns the number of cached userdata values.
	T.RawSetString("UserDataCacheLen", s.WrapFunc(func(s rbxmk.State) int {
		return s.Push(types.Int(s.UserDataCacheLen()))
	}))

	T.RawSetString("DummySymbol", s.UserDataOf(rtypes.Symbol{Name: "DummySymbol"}, "Symbol"))

	s.L.SetGlobal("T", T)
}

// TestScripts runs each .lua file in testdata as a Lua script. If the first
// line starts with a comment that contains "fail", then the script is expected
// to throw an error. All scripts receive the arguments from scriptArguments.
func TestScripts(t *testing.T) {
	var files []string
	wd, _ := os.Getwd()
	for _, arg := range os.Args[2:] {
		if strings.HasPrefix(arg, "-test.") {
			continue
		}
		rel, err := filepath.Rel(wd, arg)
		if err != nil {
			rel = arg
		}
		files = append(files, rel)
	}
	if len(files) == 0 {
		err := filepath.Walk(testdata, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() &&
				filepath.Ext(info.Name()) == ".lua" &&
				!strings.HasPrefix(filepath.Base(info.Name()), "_") {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("error walking testdata: %s", err)
		}
	}
	for _, file := range files {
		t.Run(filepath.ToSlash(file), func(t *testing.T) {
			args := scriptArguments
			args[1] = file
			err := Main(args[:], Std{
				in:  os.Stdin,
				out: os.Stdout,
				err: os.Stderr,
			}, func(s rbxmk.State) { initMain(s, t) })
			if err != nil {
				t.Errorf("script %s: %s", file, err)
			}
		})
	}
}
