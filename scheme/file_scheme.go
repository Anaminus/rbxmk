package scheme

import (
	"errors"
	"github.com/anaminus/rbxmk"
	"os"
)

func init() {
	Schemes.Register(rbxmk.Scheme{
		Name: "file",
		Input: &rbxmk.InputScheme{
			Handler: fileInputSchemeHandler,
		},
		Output: &rbxmk.OutputScheme{
			Handler:   fileOutputSchemeHandler,
			Finalizer: fileOutputFinalizer,
		},
	})
}

func GuessFileExtension(opt rbxmk.Options, format, filename string) (ext string) {
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

func fileInputSchemeHandler(opt rbxmk.Options, node *rbxmk.InputNode, inref []string) (outref []string, data rbxmk.Data, err error) {
	var ext string
	if ext = GuessFileExtension(opt, node.Format, inref[0]); ext == "" {
		return nil, nil, errors.New("failed to guess format")
	}

	// Find format.
	if !opt.Formats.Registered(ext) {
		return nil, nil, errors.New("format is not registered")
	}

	// Open file.
	file, err := os.Open(inref[0])
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	// Decode file with format.
	if err = opt.Formats.Decode(ext, opt, file.Name(), file, &data); err != nil {
		return nil, nil, err
	}

	return inref[1:], data, nil
}

func fileOutputSchemeHandler(opt rbxmk.Options, node *rbxmk.OutputNode, inref []string) (ext string, outref []string, data rbxmk.Data, err error) {
	if ext = GuessFileExtension(opt, node.Format, inref[0]); ext == "" {
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
	if n, _ := file.Seek(0, os.SEEK_END); n == 0 {
		// Empty file, do not decode.
		return ext, inref[1:], nil, nil
	}
	file.Seek(0, os.SEEK_SET)

	// Decode file with format.
	if err = opt.Formats.Decode(ext, opt, file.Name(), file, &data); err != nil {
		return "", nil, nil, err
	}
	return ext, inref[1:], data, nil
}

func fileOutputFinalizer(opt rbxmk.Options, node *rbxmk.OutputNode, inref []string, ext string, outdata rbxmk.Data) (err error) {
	if !opt.Formats.Registered(ext) {
		return errors.New("format is not registered")
	}
	file, err := os.Create(inref[0])
	if err != nil {
		return err
	}
	defer file.Close()
	if err = opt.Formats.Encode(ext, opt, file.Name(), file, outdata); err != nil {
		return err
	}
	return file.Sync()
}
