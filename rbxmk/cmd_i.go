package main

import (
	"errors"
	"fmt"
	"os"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/gopher-lua/parse"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/formats"
	"github.com/anaminus/rbxmk/library"
	"github.com/anaminus/snek"
	"github.com/peterh/liner"
)

func init() {
	Program.Register(snek.Def{
		Name:    "i",
		Summary: "Enter interactive mode.",
		Description: `
Enters interactive mode. Each prompt executes a chunk of Lua code.

If a prompt begins with '=', then the comma-separated list of expressions that
follow are evaluated and printed to standard output.

The environment contains the os.exit function. When called, interactive mode is
terminated, and the program exits.

Within supported terminals, the following shortcuts are available:

	Ctrl-A, Home             Move cursor to beginning of line.
	Ctrl-E, End              Move cursor to end of line
	Ctrl-B, Left             Move cursor one character left.
	Ctrl-F, Right            Move cursor one character right.
	Ctrl-Left, Alt-B         Move cursor to previous word.
	Ctrl-Right, Alt-F        Move cursor to next word
	Ctrl-D, Del              If line is not empty, delete character under cursor.
	Ctrl-D                   If line is empty, end of file.
	Ctrl-C                   Reset input (create new empty prompt).
	Ctrl-L                   Clear screen (line is unmodified).
	Ctrl-T                   Transpose previous character with current character.
	Ctrl-H, BackSpace        Delete character before cursor.
	Ctrl-W, Alt-BackSpace    Delete word leading up to cursor.
	Alt-D                    Delete word following cursor.
	Ctrl-K                   Delete from cursor to end of line.
	Ctrl-U                   Delete from start of line to cursor.
	Ctrl-P, Up               Previous match from history.
	Ctrl-N, Down             Next match from history.
	Ctrl-R                   Reverse Search history (Ctrl-S forward, Ctrl-G cancel).
	Ctrl-Y                   Paste from Yank buffer (Alt-Y to paste next yank instead).
	Tab                      Next completion.
	Shift-Tab                (after Tab) Previous completion.
`,
		New: func() snek.Command { return &InteractiveCommand{} },
	})
}

type InteractiveCommand struct {
	IncludedRoots []string
	InsecurePaths bool
	Debug         bool
	Init          func(rbxmk.State)
}

func (c *InteractiveCommand) SetFlags(flags snek.FlagSet) {
	flags.Var((*repeatedString)(&c.IncludedRoots), "include-root", "Mark a path as an accessible root directory. May be specified any number of times.")
	flags.BoolVar(&c.InsecurePaths, "allow-insecure-paths", false, "Disable path restrictions, allowing scripts to access any path in the file system.")
	flags.BoolVar(&c.Debug, "debug", false, "Display stack traces when an error occurs.")
}

func (c *InteractiveCommand) Run(opt snek.Options) (err error) {
	// Parse flags.
	if err := opt.ParseFlags(); err != nil {
		return err
	}

	// Initialize world.
	world := rbxmk.NewWorld(lua.NewState(lua.Options{
		SkipOpenLibs:        true,
		IncludeGoStackTrace: c.Debug,
	}))
	if c.InsecurePaths {
		world.FS.SetSecured(false)
	}
	if wd, err := os.Getwd(); err == nil {
		// Working directory is an accessible root.
		world.FS.AddRoot(wd)
	}
	for _, root := range c.IncludedRoots {
		world.FS.AddRoot(root)
	}
	for _, f := range formats.All() {
		world.RegisterFormat(f())
	}
	for _, lib := range library.All() {
		if err := world.Open(lib); err != nil {
			return err
		}
	}

	state := world.State()
	state.SetGlobal("_RBXMK_VERSION", lua.LString(VersionString()))
	exit := make(chan struct{})
	state.GetGlobal("os").(*lua.LTable).RawSetString("exit", world.WrapFunc(func(s rbxmk.State) int {
		close(exit)
		return 0
	}))

	if c.Init != nil {
		c.Init(rbxmk.State{World: world, L: state})
	}

	// Initialize terminal prompt.
	line := liner.NewLiner()
	line.SetCtrlCAborts(true)
	line.SetMultiLineMode(true)

	// Begin read-eval-print loop.
repl:
	for {
		var chunk string
		if chunk, err = loadLine(state, line); err != nil {
			if errors.Is(err, liner.ErrPromptAborted) {
				err = nil
				break repl
			}
			if !errors.Is(err, expr) {
				break repl
			}
		}
		if chunk == "" {
			continue
		}
		if err := world.DoString(chunk, "stdin", 0); err != nil {
			fmt.Fprintln(opt.Stderr, err)
			continue
		}
		if err == expr {
			// Print values returned by chunk.
			err = nil
			n := state.GetTop()
			s := make([]interface{}, n)
			for i := 1; i <= n; i++ {
				s[i-1] = state.ToStringMeta(state.Get(i))
			}
			state.Pop(n)
			fmt.Println(s...)
		}
		// Check if os.exit was called.
		select {
		case <-exit:
			break repl
		default:
		}
	}

	if e := line.Close(); err == nil {
		return e
	}
	return err
}

// expr indicates that a chunk is an expression.
var expr = errors.New("expression")

// loadLine prompts for a Lua chunk. If the chunk begins with '=', it is
// interpreted as a return statement, and returns the expr error.
func loadLine(l *lua.LState, line *liner.State) (string, error) {
	chunk, err := line.Prompt("> ")
	if err != nil {
		return "", err
	}
	if chunk == "" {
		return "", nil
	}
	if chunk[0] == '=' {
		chunk = chunk[1:]
		if _, err := l.LoadString("return " + chunk); err == nil {
			line.AppendHistory(chunk)
			return "return " + chunk, expr
		}
	}
	if chunk, err = loadMultiline(chunk, l, line); err != nil {
		return "", err
	}
	line.AppendHistory(chunk)
	return chunk, nil
}

// loadMultiline continually prompts until a parsed Lua chunk is complete.
func loadMultiline(chunk string, l *lua.LState, line *liner.State) (string, error) {
	for {
		if _, err := l.LoadString(chunk); !isIncomplete(err) {
			return chunk, nil
		}
		line, err := line.Prompt(">> ")
		if err != nil {
			return "", err
		}
		chunk = chunk + "\n" + line
	}
}

// isIncomplete returns whether the error indicates that a parsed Lua chunk is
// incomplete.
func isIncomplete(err error) bool {
	if err == nil {
		return false
	}
	if lerr, ok := err.(*lua.ApiError); ok {
		if perr, ok := lerr.Cause.(*parse.Error); ok {
			return perr.Pos.Line == parse.EOF
		}
	}
	return false
}
