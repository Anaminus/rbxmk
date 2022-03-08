package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/anaminus/cobra"
	terminal "golang.org/x/term"
)

var Frag = NewFragments(initFragRoot())
var docState = NewDocState(Frag)

func Doc(fragref string) string {
	return docState.Doc(fragref)
}

func DocFlag(fragref string) string {
	return docState.DocFlag(fragref)
}

const usageTemplate = `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{width | .LocalFlags.FlagUsagesWrapped | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{width | .InheritedFlags.FlagUsagesWrapped | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

var Program = &cobra.Command{
	Use:           "rbxmk",
	Short:         Doc("Commands:Summary"),
	Long:          Doc("Commands:Summary"),
	SilenceUsage:  false,
	SilenceErrors: true,
}

func init() {
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
	docState.UnresolvedFragments()
	if err := Program.ExecuteContext(ctx); err != nil {
		Program.PrintErrln(err)
	}
}
