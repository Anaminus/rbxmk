package main

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/PuerkitoBio/goquery"
	"github.com/anaminus/cobra"
	"github.com/anaminus/pflag"
	"github.com/anaminus/rbxmk/rbxmk/term"
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
	List   bool
	Format string
	Width  int
}

func (c *DocCommand) SetFlags(flags *pflag.FlagSet) {
	flags.BoolVarP(&c.List, "list", "l", false, DocFlag("Commands/doc:Flags/list"))
	flags.StringVarP(&c.Format, "format", "f", "terminal", DocFlag("Commands/doc:Flags/format"))
	flags.IntVarP(&c.Width, "width", "w", -1, DocFlag("Commands/doc:Flags/width"))
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
	var content string
	switch c.Format {
	case "", "html":
		content = Frag.ResolveWith(ref, FragOptions{
			Renderer: goquery.Render,
		})
	case "terminal":
		content = Frag.ResolveWith(ref, FragOptions{
			Renderer: term.Renderer{Width: c.Width, TabSize: 4}.Render,
		})
	}
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
				return "<ul>\n\t<li>" + strings.Join(topics, "</li>\n\t<li>") + "</li>\n</ul>"
			},
		},
	}))
	return nil
}
