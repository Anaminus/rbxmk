package main

import (
	"fmt"
	"strings"

	"github.com/anaminus/cobra"
	"github.com/anaminus/pflag"
)

func init() {
	var c DocCommand
	var cmd = &cobra.Command{
		Use:  "doc",
		Args: cobra.MaximumNArgs(1),
		RunE: c.Run,
	}
	c.SetFlags(cmd.PersistentFlags())
	Program.AddCommand(cmd)
}

type DocCommand struct {
	List bool
}

func (c *DocCommand) SetFlags(flags *pflag.FlagSet) {
	flags.BoolVarP(&c.List, "list", "l", false, DocFlag("Commands/doc:Flags/list"))
}

func (c *DocCommand) Run(cmd *cobra.Command, args []string) error {
	var ref string
	if len(args) > 0 {
		ref = args[0]
	}
	if c.List {
		topics := ListFragments(ref)
		for _, topic := range topics {
			cmd.Println(topic)
		}
		return nil
	} else if ref == "" {
		fmt.Println(ResolveFragment("Messages/doc:Topics"))
		return nil
	}
	content := ResolveFragment(ref)
	if content != "" {
		cmd.Println(content)
		return nil
	}
	topics := ListFragments(ref)
	if len(topics) == 0 {
		cmd.Println(FormatFrag("Messages/doc:NoTopicContent", ref))
		return nil
	}
	cmd.Println(ResolveFragmentWith("Messages/doc:SubTopics", FragOptions{
		TmplFuncs: FuncMap{
			"SubTopics": func() string {
				return "\n\t" + strings.Join(topics, "\n\t")
			},
		},
	}))
	return nil
}
