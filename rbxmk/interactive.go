package main

import (
	"context"
	"errors"
	"strings"

	"github.com/anaminus/cobra"
	"github.com/anaminus/pflag"
	"github.com/anaminus/rbxmk/rbxmk/render/term"
	"github.com/kballard/go-shellquote"
	"github.com/peterh/liner"
)

var exitCommand = &cobra.Command{
	Use:   "exit " + Doc("Commands/exit:Arguments"),
	Short: Doc("Commands/exit:Summary"),
	Long:  Doc("Commands/exit:Description"),
	Args:  cobra.NoArgs,
	Run:   ExitCommand{}.Run,
}

type ExitCommand struct{}

func (ExitCommand) Run(cmd *cobra.Command, args []string) {
	ProgramExit()
}

func resetFlags(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		f.Changed = false
	})
	for _, sub := range cmd.Commands() {
		resetFlags(sub)
	}
}

// InteractiveMode allows the user to run program commands interactively.
func InteractiveMode(ctx context.Context) (err error) {
	Program.AddCommand(exitCommand)

	line := liner.NewLiner()
	line.SetCtrlCAborts(true)

	resetFlags(Program)

	format := Frag.ResolveWith("interactive", FragOptions{
		Renderer: term.NewRenderer(0).Render,
	})
	Program.PrintErrln(strings.TrimSpace(format) + "\n")

repl:
	for {
		var command string
		if command, err = line.Prompt("rbxmk "); err != nil {
			if errors.Is(err, liner.ErrPromptAborted) {
				err = nil
			}
			break repl
		}
		if command == "" {
			continue
		}
		line.AppendHistory(command)
		args, err := shellquote.Split(command)
		if err != nil {
			Program.PrintErrln(err)
			continue
		}
		Program.SetArgs(args)
		if err := Program.ExecuteContext(ctx); err != nil {
			Program.PrintErrln(err)
		}
		select {
		case <-ctx.Done():
			break repl
		default:
		}
	}

	if e := line.Close(); err == nil {
		return e
	}
	return err
}
