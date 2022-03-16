package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/anaminus/cobra"
	"github.com/anaminus/pflag"
)

func init() {
	var c VersionCommand
	var cmd = &cobra.Command{
		Use: "version",
		Run: c.Run,
	}
	c.SetFlags(cmd.PersistentFlags())
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

func writeModuleString(s *strings.Builder, mod debug.Module, prefix string) {
	s.WriteString("\t")
	s.WriteString(prefix)
	s.WriteString("\t")
	s.WriteString(mod.Path)
	if mod.Version != "" {
		s.WriteString("\t")
		s.WriteString(mod.Version)
	}
	if mod.Sum != "" {
		s.WriteString("\t")
		s.WriteString(mod.Sum)
	}
	s.WriteString("\n")
	if mod.Replace != nil {
		writeModuleString(s, *mod.Replace, "=>")
	}
}

func DetailedInfoString() string {
	var s strings.Builder
	fmt.Fprintf(&s, "rbxmk version: %s\n", Version)
	if Prerelease != "" {
		fmt.Fprintf(&s, "rbxmk prerelease: %s\n", Prerelease)
	}
	if Build != "" {
		fmt.Fprintf(&s, "rbxmk build: %s\n", Build)
	}
	if sslKeyLogFileEnvVar != "" {
		fmt.Fprintf(&s, "ssl log var: %s\n", sslKeyLogFileEnvVar)
	}
	fmt.Fprintf(&s, "go version: %s\n", runtime.Version())
	fmt.Fprintf(&s, "go compiler: %s\n", runtime.Compiler)
	fmt.Fprintf(&s, "go target: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	if info, ok := debug.ReadBuildInfo(); ok {
		fmt.Fprintf(&s, "path: %s\n", info.Path)
		fmt.Fprintf(&s, "modules:\n")
		writeModuleString(&s, info.Main, "mod")
		for _, dep := range info.Deps {
			writeModuleString(&s, *dep, "dep")
		}
	}
	return s.String()
}

type VersionCommand struct {
	Verbose bool
}

func (c *VersionCommand) SetFlags(flags *pflag.FlagSet) {
	flags.BoolVarP(&c.Verbose, "verbose", "v", false, DocFlag("Commands/version:Flags/verbose"))
}

func (c *VersionCommand) Run(cmd *cobra.Command, args []string) {
	if c.Verbose {
		cmd.Println(DetailedInfoString())
	} else {
		cmd.Println(VersionString())
	}
}
