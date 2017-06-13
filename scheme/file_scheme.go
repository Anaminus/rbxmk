package scheme

import (
	"errors"
	"github.com/anaminus/rbxmk"
	"os"
)

func init() {
	registerInput("file", rbxmk.InputScheme{
		Handler: fileInputSchemeHandler,
	})
	registerOutput("file", rbxmk.OutputScheme{
		Handler:   fileOutputSchemeHandler,
		Finalizer: fileOutputFinalizer,
	})
}

func guessFileExtension(opt *rbxmk.Options, format, filename string) (ext string) {
	ext = format
	if ext == "" {
		// Try to guess the format.
		if fi, err := os.Stat(filename); err == nil && fi.IsDir() {
			ext = "directory"
		} else {
			{
				i := len(filename) - 1
				for ; i >= 0 && !os.IsPathSeparator(filename[i]); i-- {
				}
				filename = filename[i+1:]
			}
			for {
				for i := 0; i < len(filename); i++ {
					if filename[i] == '.' {
						filename = filename[i+1:]
						goto check
					}
				}
				return ""
			check:
				if opt.Formats.Registered(filename) {
					return filename
				}
			}
		}
	}
	return ext
}

func fileInputSchemeHandler(opt *rbxmk.Options, node *rbxmk.InputNode, filename string) (ext string, src *rbxmk.Source, err error) {
	if ext = guessFileExtension(opt, node.Format, filename); ext == "" {
		return "", nil, errors.New("failed to guess format")
	}

	// Find format.
	format, exists := opt.Formats.Init(ext, opt)
	if !exists {
		return "", nil, errors.New("format is not registered")
	}

	// Open file.
	file, err := os.Open(filename)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	// Decode file with format.
	if src, err = format.Decode(file); err != nil {
		return "", nil, err
	}

	return ext, src.Copy(), nil
}

func fileOutputSchemeHandler(opt *rbxmk.Options, node *rbxmk.OutputNode, filename string) (ext string, src *rbxmk.Source, err error) {
	if ext = guessFileExtension(opt, node.Format, filename); ext == "" {
		return "", nil, errors.New("failed to guess format")
	}

	// Find format.
	format, exists := opt.Formats.Init(ext, opt)
	if !exists {
		return "", nil, errors.New("format is not registered")
	}

	// Open file.
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return ext, &rbxmk.Source{}, nil
	}
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	// Decode file with format.
	if src, err = format.Decode(file); err != nil {
		return "", nil, err
	}
	return ext, src, nil
}

func fileOutputFinalizer(opt *rbxmk.Options, node *rbxmk.OutputNode, filename, ext string, outsrc *rbxmk.Source) (err error) {
	format, exists := opt.Formats.Init(ext, opt)
	if !exists {
		return errors.New("format is not registered")
	}
	if !format.CanEncode(outsrc) {
		return errors.New("cannot encode transformed output")
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return format.Encode(file, outsrc.Copy())
}
