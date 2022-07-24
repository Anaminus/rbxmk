package main

import (
	"errors"

	"github.com/anaminus/cobra"
	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/gopher-lua/parse"
	"github.com/anaminus/pflag"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/library"
	"github.com/peterh/liner"
)

func init() {
	var c InteractiveCommand
	var cmd = Register.NewCommand(dump.Command{
		Summary:     "Commands/interactive:Summary",
		Description: "Commands/interactive:Description",
	}, &cobra.Command{
		Use:     "interactive",
		Aliases: []string{"i"},
		RunE:    c.Run,
	})
	c.SetFlags(cmd.PersistentFlags())
	Program.AddCommand(cmd)
}

type InteractiveCommand struct {
	WorldFlags
	DescFlags
	Init func(rbxmk.State)
}

func (c *InteractiveCommand) SetFlags(flags *pflag.FlagSet) {
	c.WorldFlags.SetFlags(flags)
	c.DescFlags.SetFlags(flags)
}

func (c *InteractiveCommand) Run(cmd *cobra.Command, args []string) (err error) {
	// Initialize world.
	world, err := InitWorld(WorldOpt{
		WorldFlags:       c.WorldFlags,
		IncludeLibraries: library.All(),
	})
	if err != nil {
		return err
	}
	injectSSLKeyLogFile(world, cmd.ErrOrStderr())
	state := world.LuaState()
	exit := make(chan struct{})
	state.GetGlobal("os").(*lua.LTable).RawSetString("exit", world.WrapFunc(func(s rbxmk.State) int {
		close(exit)
		return 0
	}))
	if c.Init != nil {
		c.Init(world.State())
	}

	// Initialize global descriptor.
	world.Desc, err = c.DescFlags.Resolve(world.Client)
	if err != nil {
		return err
	}
	world.SetEnumGlobal()

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
			cmd.PrintErrln(err)
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
			cmd.Println(s...)
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
		if _, err := l.LoadString("return " + chunk[1:]); err == nil {
			line.AppendHistory(chunk)
			return "return " + chunk[1:], expr
		}
		chunk = chunk[1:]
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
