package main

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

const testdata = "testdata"

var scriptArguments = []string{
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

// TestScripts runs each .lua file in testdata as a Lua script. If the first
// line starts with a comment that contains "fail", then the script is expected
// to throw an error. All scripts receive the arguments from scriptArguments.
func TestScripts(t *testing.T) {
	files, err := ioutil.ReadDir(testdata)
	if err != nil {
		t.Fatalf("missing testdata: %s", err)
	}
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".lua" {
			continue
		}
		script, err := ioutil.ReadFile(filepath.Join(testdata, file.Name()))
		if err != nil {
			t.Errorf("read file: %s", err)
			continue
		}
		r := bytes.NewReader(script)
		first, _ := bufio.NewReader(r).ReadString('\n')
		mustFail := strings.HasPrefix(strings.TrimSpace(first), "--") && strings.Contains(first, "fail")
		r.Reset(script)
		err = Main(scriptArguments, Std{
			in: &dummyFile{r: r, info: &dummyInfo{
				name:  "stdin",
				size:  0,
				isdir: false,
				mode:  69206454,
				time:  time.Now(),
			}},
			out: os.Stdout,
			err: os.Stderr,
		})
		if mustFail && err == nil {
			t.Errorf("script %s: error expected", file.Name())
		} else if !mustFail && err != nil {
			t.Errorf("script %s: %s", file.Name(), err)
		}
	}
}
