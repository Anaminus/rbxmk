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
		Name:      "upload-asset",
		Summary:   "Upload an asset.",
		Arguments: "[ FLAGS ] [ -id INT ] PATH",
		Description: `
Uploads an asset to the roblox website.

The -id flag specifies the ID of the asset to upload. If not specified, then a
new asset will be created, and the ID of the asset will be returned.

The first non-flag argument is the path to a file to read from, which is
required. If the path is "-", then the file will be read from standard input.

Each cookie flag appends to the list of cookies that will be sent with the
request. Such flags can be specified any number of times.
`,
		New: func() snek.Command { return &UploadAssetCommand{} },
	})
}

type UploadAssetCommand struct {
	Cookies     rtypes.Cookies
	ID          int64
	AssetFormat string
	FileFormat  string
}

func (c *UploadAssetCommand) SetFlags(flags snek.FlagSet) {
	SetCookieFlags(&c.Cookies, flags)
	flags.Int64Var(&c.ID, "id", 0, "The ID of the asset to download (required).")
	flags.StringVar(&c.AssetFormat, "format", "bin", "The format to encode the asset as.")
	flags.StringVar(&c.FileFormat, "file-format", "", "The format to decode the file as. Defaults to -format.")
}

func (c *UploadAssetCommand) Run(opt snek.Options) error {
	// Parse flags.
	if err := opt.ParseFlags(); err != nil {
		return err
	}
	if c.ID < 0 {
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
	if assetFormat.Encode == nil {
		return fmt.Errorf("cannot encode with format %s", assetFormat.Name)
	}
	var fileFormat rbxmk.Format
	if c.FileFormat == "" {
		fileFormat = assetFormat
	} else {
		fileFormat = world.Format(c.FileFormat)
		if fileFormat.Name == "" {
			return fmt.Errorf("unknown file format %q", c.FileFormat)
		}
		if fileFormat.Decode == nil {
			return fmt.Errorf("cannot decode with format %s", fileFormat.Name)
		}
	}

	// Upload asset.
	var file io.ReadCloser
	switch filename := opt.Arg(0); filename {
	case "":
		return fmt.Errorf("must specify path of file to upload")
	case "-":
		file = opt.Stdin
	default:
		f, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("open file: %w", err)
		}
		file = f
	}
	body, err := fileFormat.Decode(world.Global, rtypes.FormatSelector{Format: fileFormat.Name}, file)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("close file: %w", err)
	}

	options := rtypes.RBXAssetOptions{
		AssetID: c.ID,
		Cookies: c.Cookies,
		Format:  rtypes.FormatSelector{Format: assetFormat.Name},
		Body:    body,
	}
	var id int64 = -1
	if c.ID == 0 {
		id, err = library.RBXAssetIDSource{World: world}.Create(options)
	} else {
		err = library.RBXAssetIDSource{World: world}.Write(options)
	}
	if err != nil {
		return fmt.Errorf("upload asset: %w", err)
	}
	if id >= 0 {
		fmt.Fprintln(opt.Stdout, id)
	}
	return nil
}
