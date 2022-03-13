package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/anaminus/cobra"
	"github.com/anaminus/rbxmk/rbxmk/term"
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

// Template function that expands environment variables. Each argument is an
// alternating key and value. The last value is the string to expand.
func expand(p ...string) string {
	if len(p) == 0 {
		return ""
	}
	s, p := p[len(p)-1], p[:len(p)-1]
	m := make(map[string]string, len(p)/2)
	for i := 1; i < len(p); i += 2 {
		m[p[i-1]] = p[i]
	}
	return os.Expand(s, func(s string) string { return m[s] })
}

func init() {
	fragTmplFuncs["frag"] = FragContent
	fragTmplFuncs["fraglist"] = Frag.List
	fragTmplFuncs["expand"] = os.Expand

	cobra.AddTemplateFunc("frag", func(fragref string) string {
		return Frag.ResolveWith(fragref, FragOptions{
			Renderer: term.Renderer{Width: -1}.Render,
			Inner:    true,
		})
	})
	cobra.AddTemplateFunc("expand", expand)
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
