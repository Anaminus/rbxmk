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
	"github.com/anaminus/rbxmk/dump"
)

func init() {
	var c VersionCommand
	var cmd = Register.NewCommand(dump.Command{
		Summary:     "Commands/version:Summary",
		Description: "Commands/version:Description",
	}, &cobra.Command{
		Use:  "version",
		RunE: c.Run,
	})
	c.SetFlags(cmd.Flags())
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
	Prerelease string  `json:",omitempty"`
	Build      string  `json:",omitempty"`
	Go         *GoInfo `json:",omitempty"`
}

type GoInfo struct {
	Version    string
	Compiler   string
	TargetOS   string
	TargetArch string
	Build      *BuildInfo `json:",omitempty"`
}

type BuildInfo struct {
	GoVersion string
	Path      string          `json:",omitempty"`
	Main      debug.Module    `json:",omitempty"`
	Deps      []*debug.Module `json:",omitempty"`
	Settings  []debug.BuildSetting
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
	if v.Go != nil {
		fmt.Fprintf(&s, "go version: %s\n", v.Go.Version)
		fmt.Fprintf(&s, "go compiler: %s\n", v.Go.Compiler)
		fmt.Fprintf(&s, "go target: %s/%s\n", v.Go.TargetOS, v.Go.TargetArch)
		if v.Go.Build != nil {
			fmt.Fprintf(&s, "settings:\n")
			if len(v.Go.Build.Settings) > 0 {
				for _, setting := range v.Go.Build.Settings {
					fmt.Fprintf(&s, "\t%s=%s\n", setting.Key, setting.Value)
				}
			}
			if v.Go.Build.Path != "" {
				fmt.Fprintf(&s, "path: %s\n", v.Go.Build.Path)
			}
			if v.Go.Build.Main != (debug.Module{}) || len(v.Go.Build.Deps) > 0 {
				fmt.Fprintf(&s, "modules:\n")
				writeModuleString(&s, v.Go.Build.Main, "mod")
				for _, dep := range v.Go.Build.Deps {
					writeModuleString(&s, *dep, "dep")
				}
			}
		}
	}
	return s.String()
}

type VersionCommand struct {
	Format  string
	Verbose int
	Error   bool
}

func (c *VersionCommand) SetFlags(flags *pflag.FlagSet) {
	flags.StringVarP(&c.Format, "format", "f", "text", "")
	Register.NewFlag(dump.Flag{Description: "Commands/version:Flags/format"}, flags, "format")

	flags.CountVarP(&c.Verbose, "verbose", "v", "")
	Register.NewFlag(dump.Flag{Description: "Commands/version:Flags/verbose"}, flags, "verbose")
}

func (c *VersionCommand) WriteInfo(w io.Writer) error {
	info := VersionInfo{
		Version:    Version,
		Prerelease: Prerelease,
		Build:      Build,
	}
	if c.Verbose > 0 {
		info.Go = &GoInfo{
			Version:    runtime.Version(),
			Compiler:   runtime.Compiler,
			TargetOS:   runtime.GOOS,
			TargetArch: runtime.GOARCH,
		}
		if c.Verbose > 1 {
			binfo, _ := debug.ReadBuildInfo()
			info.Go.Build = (*BuildInfo)(binfo)
			if c.Verbose < 2 {
				info.Go.Build.Settings = nil
			}
			if c.Verbose < 3 {
				info.Go.Build.Path = ""
				info.Go.Build.Main = debug.Module{}
				info.Go.Build.Deps = nil
			}
		}
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
	var w io.Writer
	if c.Error {
		w = cmd.ErrOrStderr()
	} else {
		w = cmd.OutOrStdout()
	}
	return c.WriteInfo(w)
}
