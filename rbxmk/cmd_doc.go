package main

import (
	"fmt"
	"strings"
	"text/template"

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
		topics := Frag.List(ref)
		for _, topic := range topics {
			cmd.Println(topic)
		}
		return nil
	} else if ref == "" {
		fmt.Println(Frag.Resolve("Messages/doc:Topics"))
		return nil
	}
	content := Frag.Resolve(ref)
	if content != "" {
		cmd.Println(content)
		return nil
	}
	topics := Frag.List(ref)
	if len(topics) == 0 {
		cmd.Println(Frag.Format("Messages/doc:NoTopicContent", ref))
		return nil
	}
	cmd.Println(Frag.ResolveWith("Messages/doc:SubTopics", FragOptions{
		TmplFuncs: template.FuncMap{
			"SubTopics": func() string {
				return "\n\t" + strings.Join(topics, "\n\t")
			},
		},
	}))
	return nil
}
