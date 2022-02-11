package main

import (
	"github.com/anaminus/cobra"
)

func init() {
	var c VersionCommand
	var cmd = &cobra.Command{
		Use: "version",
		Run: c.Run,
	}
	Program.AddCommand(cmd)
}

func VersionString() string {
	s := Version
	if Prerelease != "" {
		s += "-" + Prerelease
	}
	if Build != "" {
		s += "+" + Build
	}
	return s
}

type VersionCommand struct{}

func (VersionCommand) Run(cmd *cobra.Command, args []string) {
	cmd.Println(VersionString())
}
