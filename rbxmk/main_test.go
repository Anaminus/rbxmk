package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/anaminus/rbxmk"
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

func initMain(s rbxmk.State) {
	// PASS makes a positive assertion. If the first argument is a non-function,
	// then an error is thrown if the value is falsy. Otherwise, the function is
	// called. If the call errors or returns a falsy value, then an error is
	// thrown. Returning no value counts as truthy. The second argument is an
	// optional message to display with the error.
	s.L.SetGlobal("PASS", s.L.NewFunction(func(l *lua.LState) int {
		v := l.CheckAny(1)
		msg := l.OptString(2, "expected pass")
		switch v := v.(type) {
		case *lua.LFunction:
			n := s.L.GetTop()
			s.L.Push(v)
			if err := s.L.PCall(0, lua.MultRet, nil); err != nil {
				return s.RaiseError("%s: %s", msg, err.Error())
			}
			if s.L.GetTop() > n {
				if !s.L.ToBool(n + 1) {
					return s.RaiseError(msg)
				}
			}
		default:
			if !lua.LVAsBool(v) {
				return s.RaiseError(msg)
			}
		}
		return 0
	}))
	// FAIL makes a negative assertion. If the first argument is a non-function,
	// then an error is thrown if the value is truthy. Otherwise, the function
	// is called. If the call does not error or returns a truthy value, then an
	// error is thrown. Returning no value counts as falsy. The second argument
	// is an optional message to display with the error.
	s.L.SetGlobal("FAIL", s.L.NewFunction(func(l *lua.LState) int {
		v := l.CheckAny(1)
		msg := l.OptString(2, "expected fail")
		switch v := v.(type) {
		case *lua.LFunction:
			n := s.L.GetTop()
			s.L.Push(v)
			if err := s.L.PCall(0, lua.MultRet, nil); err == nil {
				if s.L.GetTop() > n {
					if s.L.ToBool(n + 1) {
						return s.RaiseError(msg)
					}
					return 0
				}
				return s.RaiseError(msg)
			} else if lua.LVAsBool(s.L.GetGlobal("SHOW_ERRORS")) {
				fmt.Printf("ERROR: %s\n", err)
			}
		default:
			if lua.LVAsBool(v) {
				return s.RaiseError(msg)
			}
		}
		return 0
	}))
}

type Directive int

const (
	Fail Directive = 1 << iota
	Skip
	None = 0
)

func checkDirective(path string) (d Directive) {
	f, err := os.Open(path)
	if err != nil {
		return d
	}
	defer f.Close()
	first, _ := bufio.NewReader(f).ReadString('\n')
	if !strings.HasPrefix(strings.TrimSpace(first), "--") {
		return d
	}
	if strings.Contains(first, "skip") {
		d |= Skip
	}
	if strings.Contains(first, "fail") {
		d |= Fail
	}
	return d
}

// TestScripts runs each .lua file in testdata as a Lua script. If the first
// line starts with a comment that contains "fail", then the script is expected
// to throw an error. All scripts receive the arguments from scriptArguments.
func TestScripts(t *testing.T) {
	var files []string
	err := filepath.Walk(testdata, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".lua" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("error walking testdata: %s", err)
	}
	for _, file := range files {
		d := checkDirective(file)
		if d&Skip != 0 {
			continue
		}
		t.Run(filepath.ToSlash(file), func(t *testing.T) {
			args := scriptArguments
			args[1] = file
			err := Main(args[:], Std{
				in:  os.Stdin,
				out: os.Stdout,
				err: os.Stderr,
			}, initMain)
			if d&Fail != 0 && err == nil {
				t.Errorf("script %s: error expected", file)
			} else if d&Fail == 0 && err != nil {
				t.Errorf("script %s: %s", file, err)
			}
		})
	}
}
