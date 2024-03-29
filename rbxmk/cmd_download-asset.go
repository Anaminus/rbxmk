package main

import (
	"fmt"
	"io"
	"os"

	"github.com/anaminus/cobra"
	"github.com/anaminus/pflag"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/formats"
	"github.com/anaminus/rbxmk/library"
	"github.com/anaminus/rbxmk/rtypes"
)

func init() {
	var c DownloadAssetCommand
	var cmd = Register.NewCommand(dump.Command{
		Arguments:   "Commands/download-asset:Arguments",
		Summary:     "Commands/download-asset:Summary",
		Description: "Commands/download-asset:Description",
	}, &cobra.Command{
		Use:  "download-asset",
		RunE: c.Run,
	})
	c.SetFlags(cmd.Flags())
	Program.AddCommand(cmd)
}

type DownloadAssetCommand struct {
	Cookies     rtypes.Cookies
	ID          int64
	AssetFormat string
	FileFormat  string
}

func (c *DownloadAssetCommand) SetFlags(flags *pflag.FlagSet) {
	SetCookieFlags(&c.Cookies, flags)

	flags.Int64Var(&c.ID, "id", 0, "")
	Register.NewFlag(dump.Flag{Description: "Commands/download-asset:Flags/id"}, flags, "id")

	flags.StringVar(&c.AssetFormat, "format", formats.F_Binary, "")
	Register.NewFlag(dump.Flag{Description: "Commands/download-asset:Flags/format"}, flags, "format")

	flags.StringVar(&c.FileFormat, "file-format", "", "")
	Register.NewFlag(dump.Flag{Description: "Commands/download-asset:Flags/file-format"}, flags, "file-format")
}

func (c *DownloadAssetCommand) Run(cmd *cobra.Command, args []string) error {
	if c.ID <= 0 {
		return fmt.Errorf("must specify valid asset ID with -id flag")
	}

	// Initialize world.
	world, err := InitWorld(WorldOpt{
		WorldFlags:     WorldFlags{Debug: false},
		ExcludeRoots:   true,
		ExcludeEnums:   true,
		ExcludeProgram: true,
	})
	if err != nil {
		return err
	}
	injectSSLKeyLogFile(world, cmd.ErrOrStderr())

	// Check formats.
	assetFormat := world.Format(c.AssetFormat)
	if assetFormat.Name == "" {
		return fmt.Errorf("unknown asset format %q", c.AssetFormat)
	}
	if assetFormat.Decode == nil {
		return fmt.Errorf("cannot decode with format %s", assetFormat.Name)
	}
	var fileFormat rbxmk.Format
	if c.FileFormat == "" {
		fileFormat = assetFormat
	} else {
		fileFormat = world.Format(c.FileFormat)
		if fileFormat.Name == "" {
			return fmt.Errorf("unknown file format %q", c.FileFormat)
		}
		if fileFormat.Encode == nil {
			return fmt.Errorf("cannot encode with format %s", fileFormat.Name)
		}
	}

	// Download asset.
	body, err := library.RbxAssetIdSource{World: world}.Read(rtypes.RbxAssetOptions{
		AssetId: c.ID,
		Cookies: c.Cookies,
		Format:  rtypes.FormatSelector{Format: assetFormat.Name},
	})
	if err != nil {
		return fmt.Errorf("download asset: %w", err)
	}
	var file io.Writer
	if filename := args[0]; filename == "" {
		file = cmd.OutOrStdout()
	} else {
		f, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("create file: %w", err)
		}
		defer func() {
			if e := f.Close(); e != nil {
				err = fmt.Errorf("close file: %w", e)
			}
		}()
		file = f
	}
	err = fileFormat.Encode(world.Global, rtypes.FormatSelector{Format: fileFormat.Name}, file, body)
	if err != nil {
		return fmt.Errorf("encode file: %w", err)
	}
	return nil
}
