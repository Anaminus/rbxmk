package main

import (
	"encoding/json"
	"fmt"
	"io"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/anaminus/cobra"
	"github.com/anaminus/pflag"
)

func init() {
	var c VersionCommand
	var cmd = &cobra.Command{
		Use:  "version",
		RunE: c.Run,
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

type VersionInfo struct {
	Version    string
	Prerelease string      `json:",omitempty"`
	Build      string      `json:",omitempty"`
	Config     *ConfigInfo `json:",omitempty"`
	Go         *GoInfo     `json:",omitempty"`
}

type ConfigInfo struct {
	SSLLogVar string `json:",omitempty"`
}

type GoInfo struct {
	Version    string
	Compiler   string
	TargetOS   string
	TargetArch string
	Build      *debug.BuildInfo `json:",omitempty"`
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

func (v VersionInfo) String() string {
	if v.Go == nil {
		return VersionString()
	}
	var s strings.Builder
	fmt.Fprintf(&s, "rbxmk version: %s\n", v.Version)
	if v.Prerelease != "" {
		fmt.Fprintf(&s, "rbxmk prerelease: %s\n", v.Prerelease)
	}
	if v.Build != "" {
		fmt.Fprintf(&s, "rbxmk build: %s\n", v.Build)
	}
	if v.Config != nil {
		if v.Config.SSLLogVar != "" {
			fmt.Fprintf(&s, "ssl log var: %s\n", v.Config.SSLLogVar)
		}
	}
	if v.Go != nil {
		fmt.Fprintf(&s, "go version: %s\n", v.Go.Version)
		fmt.Fprintf(&s, "go compiler: %s\n", v.Go.Compiler)
		fmt.Fprintf(&s, "go target: %s/%s\n", v.Go.TargetOS, v.Go.TargetArch)
		if v.Go.Build != nil {
			fmt.Fprintf(&s, "path: %s\n", v.Go.Build.Path)
			fmt.Fprintf(&s, "modules:\n")
			writeModuleString(&s, v.Go.Build.Main, "mod")
			for _, dep := range v.Go.Build.Deps {
				writeModuleString(&s, *dep, "dep")
			}
		}
	}
	return s.String()
}

type VersionCommand struct {
	Format  string
	Verbose bool
}

func (c *VersionCommand) SetFlags(flags *pflag.FlagSet) {
	flags.StringVarP(&c.Format, "format", "f", "text", DocFlag("Commands/version:Flags/format"))
	flags.BoolVarP(&c.Verbose, "verbose", "v", false, DocFlag("Commands/version:Flags/verbose"))
}

func (c *VersionCommand) WriteInfo(w io.Writer) error {
	info := VersionInfo{
		Version:    Version,
		Prerelease: Prerelease,
		Build:      Build,
	}
	if c.Verbose {
		info.Config = &ConfigInfo{
			SSLLogVar: sslKeyLogFileEnvVar,
		}
		info.Go = &GoInfo{
			Version:    runtime.Version(),
			Compiler:   runtime.Compiler,
			TargetOS:   runtime.GOOS,
			TargetArch: runtime.GOARCH,
		}
		info.Go.Build, _ = debug.ReadBuildInfo()
	}
	switch c.Format {
	case "json":
		je := json.NewEncoder(w)
		je.SetEscapeHTML(false)
		je.SetIndent("", "\t")
		return je.Encode(info)
	case "text":
		_, err := fmt.Fprintln(w, info.String())
		return err
	default:
		return fmt.Errorf("unknown format %q", c.Format)
	}
}

func (c *VersionCommand) Run(cmd *cobra.Command, args []string) error {
	return c.WriteInfo(cmd.OutOrStdout())
}
