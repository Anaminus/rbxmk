package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/anaminus/cobra"
	terminal "golang.org/x/term"
)

var Frag = NewFragments(initFragRoot())

// FragContent renders in HTML without the outer section or body.
func FragContent(fragref string) string {
	return Frag.ResolveWith(fragref, FragOptions{Inner: true})
}

var docState = NewDocState(Frag)

func Doc(fragref string) string     { return docState.Doc(fragref) }
func DocFlag(fragref string) string { return docState.DocFlag(fragref) }
func DocFragments() []string        { return docState.DocFragments() }
func UnresolvedFragments()          { docState.UnresolvedFragments() }

var Program = &cobra.Command{
	Use:           "rbxmk",
	Short:         Doc("Commands:Summary"),
	Long:          Doc("Commands:Summary"),
	SilenceUsage:  false,
	SilenceErrors: true,
}

func init() {
	fragTmplFuncs["frag"] = FragContent
	fragTmplFuncs["fraglist"] = Frag.List

	cobra.AddTemplateFunc("frag", FragContent)
	cobra.AddTemplateFunc("width", func() int {
		width, _, _ := terminal.GetSize(int(os.Stdout.Fd()))
		return width
	})
	Program.SetUsageTemplate(usageTemplate)
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Kill)
	defer stop()

	Program.SetIn(os.Stdin)
	Program.SetOut(os.Stdout)
	Program.SetErr(os.Stderr)

	DocumentCommands()
	UnresolvedFragments()
	if err := Program.ExecuteContext(ctx); err != nil {
		Program.PrintErrln(err)
	}
}
