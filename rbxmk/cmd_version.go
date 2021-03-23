package main

import (
	"fmt"

	"github.com/anaminus/snek"
)

func init() {
	Program.Register(snek.Def{
		Name:        "version",
		Summary:     Doc("commands/version.md/Summary"),
		Description: Doc("commands/version.md/Description"),
		New:         func() snek.Command { return VersionCommand{} },
	})
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

func (VersionCommand) Run(opt snek.Options) error {
	if err := opt.ParseFlags(); err != nil {
		return err
	}
	fmt.Fprintln(opt.Stdout, VersionString())
	return nil
}
