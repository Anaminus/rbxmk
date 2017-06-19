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

func fileInputSchemeHandler(opt *rbxmk.Options, node *rbxmk.InputNode, inref []string) (ext string, outref []string, data rbxmk.Data, err error) {
	if ext = guessFileExtension(opt, node.Format, inref[0]); ext == "" {
		return "", nil, nil, errors.New("failed to guess format")
	}

	// Find format.
	if !opt.Formats.Registered(ext) {
		return "", nil, nil, errors.New("format is not registered")
	}

	// Open file.
	file, err := os.Open(inref[0])
	if err != nil {
		return "", nil, nil, err
	}
	defer file.Close()

	// Decode file with format.
	format, err := opt.Formats.Decoder(ext, opt, file)
	if err != nil {
		return "", nil, nil, err
	}
	if err = format.Decode(&data); err != nil {
		return "", nil, nil, err
	}

	return ext, inref[1:], data, nil
}

func fileOutputSchemeHandler(opt *rbxmk.Options, node *rbxmk.OutputNode, inref []string) (ext string, outref []string, data rbxmk.Data, err error) {
	if ext = guessFileExtension(opt, node.Format, inref[0]); ext == "" {
		return "", nil, nil, errors.New("failed to guess format")
	}

	// Find format.
	if !opt.Formats.Registered(ext) {
		return "", nil, nil, errors.New("format is not registered")
	}

	// Open file.
	file, err := os.Open(inref[0])
	if os.IsNotExist(err) {
		return ext, inref[1:], nil, nil
	}
	if err != nil {
		return "", nil, nil, err
	}
	defer file.Close()

	// Decode file with format.
	format, err := opt.Formats.Decoder(ext, opt, file)
	if err != nil {
		return "", nil, nil, err
	}
	if err = format.Decode(&data); err != nil {
		return "", nil, nil, err
	}
	return ext, inref[1:], data, nil
}

func fileOutputFinalizer(opt *rbxmk.Options, node *rbxmk.OutputNode, inref []string, ext string, outdata rbxmk.Data) (err error) {
	if !opt.Formats.Registered(ext) {
		return errors.New("format is not registered")
	}
	format, err := opt.Formats.Encoder(ext, opt, outdata)
	if err != nil {
		return err
	}
	file, err := os.Create(inref[0])
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = format.WriteTo(file)
	return err
}
