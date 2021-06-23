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

	switch len(opt.Args()) {
	case 0:
		opt.WriteUsageOf(opt.Stderr, opt.Def)
		topics := ListFragments("")
		fmt.Fprintln(opt.Stderr, "\nThe following top-level topics are available:")
		for _, topic := range topics {
			fmt.Fprintf(opt.Stderr, "\t%s\n", topic)
		}
	case 1, 2:
		mode := opt.Arg(0)
		ref := opt.Arg(1)
		switch mode {
		case "frag":
			content := ResolveFragment(ref)
			if content == "" {
				topics := ListFragments(ref)
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
		case "list":
			topics := ListFragments(ref)
			fmt.Fprintln(opt.Stderr, "Topics:")
			for _, topic := range topics {
				fmt.Fprint(opt.Stderr, "\t")
				fmt.Fprintln(opt.Stdout, topic)
			}
		default:
			return fmt.Errorf("unknown mode %q (expected frag or list)", mode)
		}
	default:
		return fmt.Errorf("too many arguments")
	}
	return nil
}
