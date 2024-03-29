package main

import (
	"github.com/anaminus/cobra"
	"github.com/anaminus/pflag"
	"github.com/anaminus/rbxmk/dump"
)

type CommandRegistry struct {
	Command map[*cobra.Command]*dump.Command
	Flag    map[*pflag.Flag]*dump.Flag
}

// NewCommand registers cmd to be associated with def. If def has Arguments,
// Summary, or Description fields, they will be resolved using Doc, and applied
// to the corresponding field on cmd.
//
//     - def.Arguments sets cmd.Use
//     - def.Summary sets cmd.Short
//     - def.Description sets cmd.Long
//     - def.Deprecated sets cmd.Deprecated
//
// Returns cmd.
func (c *CommandRegistry) NewCommand(def dump.Command, cmd *cobra.Command) *cobra.Command {
	if c.Command == nil {
		c.Command = map[*cobra.Command]*dump.Command{}
	}

	def.Aliases = make([]string, len(cmd.Aliases))
	copy(def.Aliases, cmd.Aliases)
	def.Hidden = cmd.Hidden

	c.Command[cmd] = &def
	if def.Arguments != "" {
		name := cmd.Name()
		if cmd.Use == name {
			cmd.Use += " " + Doc(def.Arguments)
		}
	}
	if def.Summary != "" {
		cmd.Short = Doc(def.Summary)
	}
	if def.Description != "" {
		cmd.Long = Doc(def.Description)
	}
	if def.Deprecated != "" {
		cmd.Deprecated = Doc(def.Deprecated)
	}
	return cmd
}

// NewFlag registers flag to be associated with def. If def has the Description
// field, it will be resolved using Doc, and applied to the flag.Usage field.
// Returns flag.
func (c *CommandRegistry) NewFlag(def dump.Flag, flags *pflag.FlagSet, name string) *pflag.Flag {
	if c.Flag == nil {
		c.Flag = map[*pflag.Flag]*dump.Flag{}
	}
	flag := flags.Lookup(name)
	if flag == nil {
		panic("flag does not exist")
	}

	def.Shorthand = flag.Shorthand
	def.Default = flag.DefValue
	def.Hidden = flag.Hidden

	c.Flag[flag] = &def
	if def.Type == "" {
		def.Type = flag.Value.Type()
	}
	if def.Description != "" {
		flag.Usage = DocFlag(def.Description)
	}
	if def.Deprecated != "" {
		flag.Deprecated = DocFlag(def.Deprecated)
	}
	if def.ShorthandDeprecated != "" {
		flag.ShorthandDeprecated = DocFlag(def.ShorthandDeprecated)
	}
	return flag
}

var Register CommandRegistry
