package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/anaminus/snek"
	"github.com/robloxapi/types"
)

const testdata = "testdata"

// Replace scriptArguments[x] with test script.
const replaceIndex = 3

var scriptArguments = [...]string{
	"rbxmk_test",
	"run",
	"--debug",
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

func deepeq(s rbxmk.State, t *testing.T, msg string, a, b lua.LValue) bool {
	if b, ok := b.(*lua.LTable); ok {
		if a, ok := a.(*lua.LTable); ok {
			visited := map[lua.LValue]struct{}{}
			b.ForEach(func(k, v lua.LValue) error {
				switch k := k.(type) {
				case lua.LNumber:
					deepeq(s, t, msg+"["+k.String()+"]", a.RawGetInt(int(k)), v)
				case lua.LString:
					deepeq(s, t, msg+"."+k.String(), a.RawGetString(string(k)), v)
				default:
					deepeq(s, t, msg+"["+k.String()+"]", a.RawGet(k), v)
				}
				visited[k] = struct{}{}
				return nil
			})
			a.ForEach(func(k, v lua.LValue) error {
				if _, ok := visited[k]; ok {
					return nil
				}
				switch k := k.(type) {
				case lua.LNumber:
					deepeq(s, t, msg+"["+k.String()+"]", v, b.RawGetInt(int(k)))
				case lua.LString:
					deepeq(s, t, msg+"."+k.String(), v, b.RawGetString(string(k)))
				default:
					deepeq(s, t, msg+"["+k.String()+"]", v, b.RawGet(k))
				}
				return nil
			})
			return true
		}
	}
	if !s.L.Equal(a, b) {
		if a.Type() != b.Type() {
			lineError(s, t, "expected type %s, got %s", msg, b.Type(), a.Type())
			return false
		}
		if a.Type() == lua.LTString {
			lineError(s, t, "expected %q, got %q", msg, b.String(), a.String())
			return false
		}
		lineError(s, t, "expected %s, got %s", msg, b.String(), a.String())
		return false
	}
	return true
}

func lineError(s rbxmk.State, t *testing.T, f, msg string, v ...interface{}) {
	d, ok := s.L.GetStack(-1)
	if ok {
		s.L.GetInfo("l", d, lua.LNil)
		if d.CurrentLine > 0 {
			v = append([]interface{}{d.CurrentLine, msg}, v...)
			t.Errorf("line %d: %s: "+f, v...)
			return
		}
	}
	v = append([]interface{}{msg}, v...)
	t.Errorf("%s: "+f, v...)
}

func initMain(s rbxmk.State, t *testing.T) {
	T := s.L.CreateTable(0, 2)

	// Pass makes a positive assertion. If the first argument is a non-function,
	// then an error is emitted if the value is falsy. Otherwise, the function
	// is called. If the call errors or returns a falsy value, then an error is
	// emitted. Returning no value counts as truthy. The error is emitted to
	// testing.T, but does not cause a Lua error to be thrown. The second
	// optional argument is a string that describes the assertion, which is
	// included with an emitted error.
	T.RawSetString("Pass", s.WrapFunc(func(s rbxmk.State) int {
		v := s.CheckAny(1)
		msg := s.OptString(2, "expected pass")
		switch v := v.(type) {
		case *lua.LFunction:
			n := s.Count()
			s.L.Push(v)
			if err := s.L.PCall(0, lua.MultRet, nil); err != nil {
				lineError(s, t, "%s", msg, err.Error())
				return 0
			}
			if s.Count() > n {
				if !s.L.ToBool(n + 1) {
					if m := s.L.ToString(n + 2); m != "" {
						lineError(s, t, "%s", msg, m)
					} else {
						lineError(s, t, "", msg)
					}
					return 0
				}
			}
		default:
			if !lua.LVAsBool(v) {
				lineError(s, t, "", msg)
				return 0
			}
		}
		return 0
	}))

	// Fail makes a negative assertion. If the first argument is a non-function,
	// then an error is emitted if the value is truthy. Otherwise, the function
	// is called. If the call does not error or returns a truthy value, then an
	// error is emitted. Returning no value counts as falsy. The error is
	// emitted to testing.T, but does not cause a Lua error to be thrown. The
	// second optional argument is a string that describes the assertion, which
	// is included with an emitted error.
	T.RawSetString("Fail", s.WrapFunc(func(s rbxmk.State) int {
		v := s.CheckAny(1)
		msg := s.OptString(2, "expected fail")
		switch v := v.(type) {
		case *lua.LFunction:
			n := s.Count()
			s.L.Push(v)
			if err := s.L.PCall(0, lua.MultRet, nil); err == nil {
				if s.Count() > n {
					if s.L.ToBool(n + 1) {
						if m := s.L.ToString(n + 2); m != "" {
							lineError(s, t, "%s", msg, m)
						} else {
							lineError(s, t, "", msg)
						}
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
				lineError(s, t, "", msg)
				return 0
			}
		}
		return 0
	}))

	// Equal asserts whether two values are deeply equal.
	T.RawSetString("Equal", s.L.NewFunction(func(l *lua.LState) int {
		msg := s.CheckString(1)
		a := s.CheckAny(2)
		b := s.CheckAny(3)
		deepeq(s, t, msg, a, b)
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

	// UserDataID returns a string that identifies a userdata value.
	T.RawSetString("UserDataID", s.WrapFunc(func(s rbxmk.State) int {
		return s.Push(types.String(fmt.Sprintf("%p", s.CheckUserData(1))))
	}))

	T.RawSetString("DummySymbol", s.UserDataOf(rtypes.Symbol{Name: "DummySymbol"}, "Symbol"))

	s.L.SetGlobal("T", T)
}

// TestScripts runs each .lua file in testdata as a Lua script. If the first
// line starts with a comment that contains "fail", then the script is expected
// to throw an error. All scripts receive the arguments from scriptArguments.
func TestScripts(t *testing.T) {
	program := snek.NewProgram("", scriptArguments[:])
	program.Register(snek.Def{
		Name: "run",
		New: func() snek.Command {
			return &RunCommand{Init: func(s rbxmk.State) { initMain(s, t) }}
		},
	})

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
			args := make([]string, len(scriptArguments))
			copy(args, scriptArguments[:])
			args[replaceIndex] = file
			err := program.RunWithInput("run", snek.Input{
				Program:   args[0],
				Arguments: args[2:],
			})
			if err != nil {
				t.Errorf("script %s: %s", file, err)
			}
		})
	}
}
