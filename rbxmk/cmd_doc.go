package main

import (
	"sort"
	"strings"

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
	Recursive bool
}

func (c *DocCommand) SetFlags(flags *pflag.FlagSet) {
	flags.BoolVarP(&c.List, "list", "l", false, DocFlag("Commands/doc:Flags/list"))
	flags.StringVarP(&c.Format, "format", "f", "terminal", DocFlag("Commands/doc:Flags/format"))
	flags.IntVarP(&c.Width, "width", "w", 0, DocFlag("Commands/doc:Flags/width"))
	flags.BoolVarP(&c.ForExport, "export", "", false, DocFlag("Commands/doc:Flags/export"))
	flags.BoolVarP(&c.Recursive, "recursive", "r", false, DocFlag("Commands/doc:Flags/recursive"))
}

func (c *DocCommand) Run(cmd *cobra.Command, args []string) error {
	var ref string
	if len(args) > 0 {
		ref = args[0]
	}
	if c.List || c.Recursive {
		var topics []string
		if c.Recursive {
			list := Frag.List(ref)
			for i, sub := range list {
				list[i] = ref + sub
			}
			for r := ""; len(list) > 0; {
				r, list = list[len(list)-1], list[:len(list)-1]
				topics = append(topics, r)
				for _, sub := range Frag.List(r) {
					list = append(list, r+sub)
				}
			}
			sort.Strings(topics)
		} else {
			topics = Frag.List(ref)
		}
		for _, topic := range topics {
			cmd.Println(strings.TrimPrefix(topic, ref))
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
		r := term.NewRenderer(c.Width)
		r.ForOutput = c.ForExport
		renderer = r.Render
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
		//TODO: Skip unless --format is terminal and --export is false?
		cmd.PrintErrln(Frag.ResolveWith("Messages/doc:SubTopics", FragOptions{
			Renderer:         term.NewRenderer(0).Render,
			TmplData:         ref,
			TrailingNewlines: 1,
		}))
	}
	return nil
}
