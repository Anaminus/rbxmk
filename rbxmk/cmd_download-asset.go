package main

import (
	"fmt"
	"io"
	"os"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/library"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/anaminus/snek"
)

func init() {
	Program.Register(snek.Def{
		Name:        "download-asset",
		Summary:     Doc("commands/download-asset.md/Summary"),
		Arguments:   Doc("commands/download-asset.md/Arguments"),
		Description: Doc("commands/download-asset.md/Description"),
		New:         func() snek.Command { return &DownloadAssetCommand{} },
	})
}

type DownloadAssetCommand struct {
	Cookies     rtypes.Cookies
	ID          int64
	AssetFormat string
	FileFormat  string
}

func (c *DownloadAssetCommand) SetFlags(flags snek.FlagSet) {
	SetCookieFlags(&c.Cookies, flags)
	flags.Int64Var(&c.ID, "id", 0, Doc("commands/download-asset.md/Flags/id"))
	flags.StringVar(&c.AssetFormat, "format", "bin", Doc("commands/download-asset.md/Flags/format"))
	flags.StringVar(&c.FileFormat, "file-format", "", Doc("commands/download-asset.md/Flags/file-format"))
}

func (c DownloadAssetCommand) Run(opt snek.Options) error {
	// Parse flags.
	if err := opt.ParseFlags(); err != nil {
		return err
	}
	if c.ID <= 0 {
		return fmt.Errorf("must specify valid asset ID with -id flag")
	}

	// Initialize world.
	world, err := InitWorld(WorldOpt{
		WorldFlags:       WorldFlags{Debug: false},
		ExcludeRoots:     true,
		ExcludeLibraries: true,
		ExcludeVersion:   true,
	})
	if err != nil {
		return err
	}

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
	var file io.WriteCloser
	if filename := opt.Arg(0); filename == "" {
		file = opt.Stdout
	} else {
		f, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("create file: %w", err)
		}
		file = f
	}
	body, err := library.RBXAssetIDSource{World: world}.Read(rtypes.RBXAssetOptions{
		AssetID: c.ID,
		Cookies: c.Cookies,
		Format:  rtypes.FormatSelector{Format: assetFormat.Name},
	})
	if err != nil {
		return fmt.Errorf("download asset: %w", err)
	}
	err = fileFormat.Encode(world.Global, rtypes.FormatSelector{Format: fileFormat.Name}, file, body)
	if err != nil {
		return fmt.Errorf("encode file: %w", err)
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("close file: %w", err)
	}
	return nil
}
