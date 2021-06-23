package main

import (
	"fmt"

	"github.com/anaminus/snek"
)

func init() {
	Program.Register(snek.Def{
		Name: "doc",
		New:  func() snek.Command { return &DocCommand{} },
	})
}

type DocCommand struct {
	List string
}

func (c *DocCommand) SetFlags(flags snek.FlagSet) {
	flags.StringVar(&c.List, "list", "", Doc("Commands/doc:Flags/list"))
}

func (c *DocCommand) Run(opt snek.Options) error {
	if err := opt.ParseFlags(); err != nil {
		return err
	}

	switch ref := opt.Arg(0); {
	case c.List != "":
		if c.List == "." {
			c.List = ""
		}
		topics := ListFragments(c.List)
		fmt.Fprintln(opt.Stderr, "Topics:")
		for _, topic := range topics {
			fmt.Fprint(opt.Stderr, "\t")
			fmt.Fprintln(opt.Stdout, topic)
		}
	case ref != "":
		content := ResolveFragment(ref)
		if content == "" {
			var topics []string
			if ref != "" {
				if ref == "." {
					topics = ListFragments("")
				} else {
					topics = ListFragments(ref)
				}
			}
			if len(topics) == 0 {
				return fmt.Errorf("no content for topic %q", ref)
			}
			fmt.Fprintln(opt.Stderr, "The following sub-topics are available:")
			for _, topic := range topics {
				fmt.Fprintf(opt.Stderr, "\t%s\n", topic)
			}
		} else {
			fmt.Fprintln(opt.Stdout, content)
		}
	default:
		opt.WriteUsageOf(opt.Stderr, opt.Def)
		topics := ListFragments("")
		fmt.Fprintln(opt.Stderr, "The following top-level topics are available:")
		for _, topic := range topics {
			fmt.Fprintf(opt.Stderr, "\t%s\n", topic)
		}
	}

	return nil
}
