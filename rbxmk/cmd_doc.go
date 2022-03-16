package main

import (
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
	List      bool
	Format    string
	Width     int
	ForExport bool
}

func (c *DocCommand) SetFlags(flags *pflag.FlagSet) {
	flags.BoolVarP(&c.List, "list", "l", false, DocFlag("Commands/doc:Flags/list"))
	flags.StringVarP(&c.Format, "format", "f", "terminal", DocFlag("Commands/doc:Flags/format"))
	flags.IntVarP(&c.Width, "width", "w", 0, DocFlag("Commands/doc:Flags/width"))
	flags.BoolVarP(&c.ForExport, "export", "", false, DocFlag("Commands/doc:Flags/export"))
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
		ref = "Messages/doc:Topics"
	}
	var renderer Renderer = goquery.Render
	switch c.Format {
	case "", "html":
		renderer = goquery.Render
	case "terminal":
		renderer = term.Renderer{
			Width:     c.Width,
			ForOutput: c.ForExport,
		}.Render
	}
	content := Frag.ResolveWith(ref, FragOptions{
		Renderer:         renderer,
		TrailingNewlines: 1,
	})
	count := Frag.Count(ref)
	if content == "" && count == 0 {
		cmd.Println(Frag.Format("Messages/doc:NoTopicContent", ref))
		return nil
	}
	if content != "" {
		cmd.Println(content)
	}
	if count > 0 {
		cmd.Println(Frag.ResolveWith("Messages/doc:SubTopics", FragOptions{
			Renderer:         term.Renderer{}.Render,
			TmplData:         ref,
			TrailingNewlines: 1,
		}))
	}
	return nil
}
