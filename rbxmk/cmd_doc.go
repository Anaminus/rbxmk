package main

import (
	"fmt"

	"github.com/anaminus/cobra"
)

func init() {
	var c DocCommand
	var cmd = &cobra.Command{
		Use:  "doc",
		RunE: c.Run,
	}
	Program.AddCommand(cmd)
}

type DocCommand struct{}

func (c *DocCommand) Run(cmd *cobra.Command, args []string) error {
	switch len(args) {
	case 0:
		cmd.Usage()
		topics := ListFragments("")
		cmd.PrintErrln("\nThe following top-level topics are available:")
		for _, topic := range topics {
			cmd.PrintErrf("\t%s\n", topic)
		}
	case 1, 2:
		mode := args[0]
		ref := args[1]
		switch mode {
		case "frag":
			content := ResolveFragment(ref)
			if content == "" {
				topics := ListFragments(ref)
				if len(topics) == 0 {
					return fmt.Errorf("no content for topic %q", ref)
				}
				cmd.PrintErrln("The following sub-topics are available:")
				for _, topic := range topics {
					cmd.PrintErrf("\t%s\n", topic)
				}
			} else {
				cmd.Println(content)
			}
		case "list":
			topics := ListFragments(ref)
			cmd.PrintErrln("Topics:")
			for _, topic := range topics {
				cmd.PrintErr("\t")
				cmd.Println(topic)
			}
		default:
			return fmt.Errorf("unknown mode %q (expected frag or list)", mode)
		}
	default:
		return fmt.Errorf("too many arguments")
	}
	return nil
}
