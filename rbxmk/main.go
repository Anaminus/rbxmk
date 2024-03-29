package main

import (
	"context"
	"os"
	"os/signal"
	"text/template"

	"github.com/anaminus/cobra"
	"github.com/anaminus/pflag"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/rbxmk/render/term"
	terminal "golang.org/x/term"
)

var Frag = NewFragments(initFragRoot(), template.FuncMap{
	"expand": os.Expand,
})

var docState = NewDocState(Frag)

func Doc(fragref string) string     { return docState.Doc(fragref) }
func DocFlag(fragref string) string { return docState.DocFlag(fragref) }
func DocFragments() []string        { return docState.DocFragments() }
func UnresolvedFragments()          { docState.UnresolvedFragments() }

var Program = Register.NewCommand(dump.Command{
	Summary:     "Commands:Summary",
	Description: "Commands:Summary",
}, &cobra.Command{
	Use:           "rbxmk",
	SilenceUsage:  true,
	SilenceErrors: true,
})

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

// Returns all local and inherited flags on a command.
func allFlags(c *cobra.Command) *pflag.FlagSet {
	out := pflag.NewFlagSet(c.Name(), pflag.ContinueOnError)
	c.LocalFlags().VisitAll(func(f *pflag.Flag) { out.AddFlag(f) })
	c.InheritedFlags().VisitAll(func(f *pflag.Flag) { out.AddFlag(f) })
	return out
}

func init() {
	cobra.AddTemplateFunc("frag", func(fragref string) string {
		return Frag.ResolveWith(fragref, FragOptions{
			Renderer: term.NewRenderer(-1).Render,
			Inner:    true,
		})
	})
	cobra.AddTemplateFunc("expand", expand)
	cobra.AddTemplateFunc("width", func() int {
		width, _, _ := terminal.GetSize(int(os.Stdout.Fd()))
		return width
	})
	Program.SetUsageTemplate(usageTemplate)

	Program.PersistentFlags().BoolP("help", "h", false, "")
	Register.NewFlag(dump.Flag{
		Persistent:  true,
		Description: "Commands/help:Flags/help",
	}, Program.PersistentFlags(), "help")

	Program.PersistentPreRun = func(_ *cobra.Command, _ []string) {
		for _, cmd := range Program.Commands() {
			switch cmd.Name() {
			case "completion":
				Register.NewCommand(dump.Command{
					Summary:     "Commands/completion:Summary",
					Description: "Commands/completion:Description",
				}, cmd)
				flags := allFlags(cmd)
				flags.VisitAll(func(f *pflag.Flag) {
					if Register.Flag[f] == nil {
						Register.NewFlag(dump.Flag{
							Description: "Commands/completion:Flags/" + f.Name,
						}, flags, f.Name)
					}
				})
				for _, sub := range cmd.Commands() {
					Register.NewCommand(dump.Command{
						Summary:     "Commands/completion/" + sub.Name() + ":Summary",
						Description: "Commands/completion/" + sub.Name() + ":Description",
					}, sub)
					flags := allFlags(sub)
					flags.VisitAll(func(f *pflag.Flag) {
						if Register.Flag[f] == nil {
							Register.NewFlag(dump.Flag{
								Description: "Commands/completion:Flags/" + f.Name,
							}, flags, f.Name)
						}
					})
				}
			case "help":
				Register.NewCommand(dump.Command{
					Arguments:   "Commands/help:Arguments",
					Summary:     "Commands/help:Summary",
					Description: "Commands/help:Description",
				}, cmd)
				flags := allFlags(cmd)
				flags.VisitAll(func(f *pflag.Flag) {
					if Register.Flag[f] == nil {
						Register.NewFlag(dump.Flag{
							Description: "Commands/help:Flags/" + f.Name,
						}, flags, f.Name)
					}
				})
			}
		}
	}
}

var ProgramContext, ProgramExit = context.WithCancel(context.Background())

// Include build information when the program panics.
func PanicWithBuildInfo() {
	if v := recover(); v != nil {
		(&VersionCommand{
			Format:  "text",
			Verbose: 2,
			Error:   true,
		}).WriteInfo(os.Stderr)
		panic(v)
	}
}

func main() {
	defer PanicWithBuildInfo()
	ctx, stop := signal.NotifyContext(ProgramContext, os.Kill)
	defer stop()

	Program.SetIn(os.Stdin)
	Program.SetOut(os.Stdout)
	Program.SetErr(os.Stderr)

	UnresolvedFragments()

	if err := Starter()(ctx); err != nil {
		Program.PrintErrln(err)
	}
}
