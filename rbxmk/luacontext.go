package main

import (
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/yuin/gopher-lua"
	"os"
	"strings"
)

type LuaContext struct {
	Options   *rbxmk.Options
	state     *lua.LState
	fileStack []FileInfo
}

type FileInfo struct {
	Path string
	os.FileInfo
}

func NewLuaContext(opt *rbxmk.Options) *LuaContext {
	ctx := &LuaContext{}
	l := lua.NewState(lua.Options{SkipOpenLibs: true})

	ctx.Options = opt
	ctx.state = l
	ctx.fileStack = make([]FileInfo, 0, 1)

	return ctx
}

func (ctx *LuaContext) State() *lua.LState {
	return ctx.state
}

func (ctx *LuaContext) PushFile(fi FileInfo) error {
	for _, f := range ctx.fileStack {
		if os.SameFile(fi.FileInfo, f.FileInfo) {
			return fmt.Errorf("\"%s\" is already running", fi.Path)
		}
	}
	ctx.fileStack = append(ctx.fileStack, fi)
	return nil
}

func (ctx *LuaContext) PopFile() {
	if len(ctx.fileStack) > 0 {
		ctx.fileStack[len(ctx.fileStack)-1] = FileInfo{}
		ctx.fileStack = ctx.fileStack[:len(ctx.fileStack)-1]
	}
}

func (ctx *LuaContext) PeekFile() (fi FileInfo, ok bool) {
	if len(ctx.fileStack) == 0 {
		return
	}
	fi = ctx.fileStack[len(ctx.fileStack)-1]
	ok = true
	return
}

func (ctx *LuaContext) DoString(s, name string, args int) (err error) {
	fn, err := ctx.state.Load(strings.NewReader(s), name)
	if err != nil {
		return err
	}
	ctx.state.Insert(fn, -args-1)
	return ctx.state.PCall(args, lua.MultRet, nil)
}

func (ctx *LuaContext) DoFile(fileName string, args int) error {
	fi, err := os.Stat(fileName)
	if err != nil {
		return err
	}
	if err = ctx.PushFile(FileInfo{fileName, fi}); err != nil {
		return err
	}

	fn, err := ctx.state.LoadFile(fileName)
	if err != nil {
		ctx.PopFile()
		return err
	}
	ctx.state.Insert(fn, -args-1)
	err = ctx.state.PCall(args, lua.MultRet, nil)
	ctx.PopFile()
	return err
}

func (ctx *LuaContext) DoFileHandle(f *os.File, args int) error {
	fi, err := f.Stat()
	if err != nil {
		return err
	}
	if err = ctx.PushFile(FileInfo{f.Name(), fi}); err != nil {
		return err
	}

	fn, err := ctx.state.Load(f, fi.Name())
	if err != nil {
		ctx.PopFile()
		return err
	}
	ctx.state.Insert(fn, -args-1)
	err = ctx.state.PCall(args, lua.MultRet, nil)
	ctx.PopFile()
	return err
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
