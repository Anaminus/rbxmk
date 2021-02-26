package main

import (
	"github.com/anaminus/but"
	"github.com/anaminus/snek"
)

func init() {
	Program.Register(snek.Def{
		Name:        "version",
		Summary:     "Display the version.",
		Description: "Displays the current version of rbxmk.",
		New:         func() snek.Command { return VersionCommand{} },
	})
}

type VersionCommand struct{}

func (VersionCommand) Run(opt snek.Options) error {
	if err := opt.ParseFlags(); err != nil {
		return err
	}
	but.Log(Version)
	return nil
}
