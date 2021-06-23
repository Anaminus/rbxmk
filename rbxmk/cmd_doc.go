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

type DocCommand struct{}

func (c *DocCommand) Run(opt snek.Options) error {
	if err := opt.ParseFlags(); err != nil {
		return err
	}

	switch args := opt.Args(); len(args) {
	case 0:
		opt.WriteUsageOf(opt.Stderr, opt.Def)
		topics := ListFragments("")
		fmt.Fprintln(opt.Stderr, "\nThe following top-level topics are available:")
		for _, topic := range topics {
			fmt.Fprintf(opt.Stderr, "\t%s\n", topic)
		}
	case 1:
		ref := args[0]
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
	case 2:
		mode := args[0]
		ref := args[1]
		switch mode {
		case "list":
			if ref == "." {
				ref = ""
			}
			topics := ListFragments(ref)
			fmt.Fprintln(opt.Stderr, "Topics:")
			for _, topic := range topics {
				fmt.Fprint(opt.Stderr, "\t")
				fmt.Fprintln(opt.Stdout, topic)
			}
		default:
			return fmt.Errorf("unknown mode %q (expected list)", mode)
		}
	default:
		return fmt.Errorf("too many arguments")
	}
	return nil
}
