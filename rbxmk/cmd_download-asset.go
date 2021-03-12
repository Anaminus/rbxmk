package main

import (
	"fmt"
	"io"
	"os"

	lua "github.com/anaminus/gopher-lua"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/formats"
	"github.com/anaminus/rbxmk/library"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/anaminus/snek"
)

func init() {
	Program.Register(snek.Def{
		Name:      "download-asset",
		Summary:   "Download an asset.",
		Arguments: "[ FLAGS ] -id INT [ PATH ]",
		Description: `
Downloads an asset from the roblox website.

The -id flag, which is required, specifies the ID of the asset to download.

The first non-flag argument is the path to a file to write to. If not specified,
then the file will be written to standard output.

Each cookie flag appends to the list of cookies that will be sent with the
request. Such flags can be specified any number of times.
`,
		New: func() snek.Command { return &DownloadAssetCommand{} },
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
	flags.Int64Var(&c.ID, "id", 0, "The ID of the asset to download (required).")
	flags.StringVar(&c.AssetFormat, "format", "bin", "The format to decode the asset as.")
	flags.StringVar(&c.FileFormat, "file-format", "", "The format to encode the file as. Defaults to -format.")
}

func (c DownloadAssetCommand) Run(opt snek.Options) error {
	if err := opt.ParseFlags(); err != nil {
		return err
	}
	if c.ID <= 0 {
		return fmt.Errorf("must specify valid asset ID with -id flag")
	}

	world := rbxmk.NewWorld(lua.NewState(lua.Options{
		SkipOpenLibs: true,
	}))
	for _, f := range formats.All() {
		world.RegisterFormat(f())
	}
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
