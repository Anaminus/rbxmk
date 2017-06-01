package main

import (
	"errors"
	"github.com/robloxapi/rbxfile"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	RegisterInputScheme("file", HandleFileInputScheme)
	RegisterInputScheme("http", HandleHTTPInputScheme)
	RegisterInputScheme("https", HandleHTTPInputScheme)

	RegisterOutputScheme("file", HandleFileOutputScheme)
}

func IsAlnum(s string) bool {
	for _, r := range s {
		if (r >= '0' && r <= '9') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= 'a' && r <= 'z') ||
			r == '_' {
			continue
		}
		return false
	}
	return true
}
func IsDigit(s string) bool {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			continue
		}
		return false
	}
	return true
}

func HandleHTTPInputScheme(opt *Options, node *InputNode, _ string) (src *Source, err error) {
	u, err := url.Parse(node.Reference[0])
	if err != nil {
		return
	}

	// Reconstruct the url without the drill.
	urlPart := u.String()
	// nextPart := node.Reference[1]

	_ = urlPart
	// TODO: get resource; expect a format
	return
}

func HandleFileInputScheme(opt *Options, node *InputNode, filename string) (src *Source, err error) {
	// Find extension.
	var ext string
	if node.Format == "" {
		ext = strings.TrimPrefix(filepath.Ext(filename), ".")
		node.Format = ext
	} else {
		ext = node.Format
	}

	// Find format.
	newFormat, exists := registeredFormats[ext]
	if !exists {
		return nil, errors.New("format is not registered")
	}

	// Open file.
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode file with format.
	if src, err = newFormat(opt).Decode(file); err != nil {
		return nil, err
	}

	refs := node.Reference[1:]

	// Try drilling to instance.
	if src, refs, err = DrillInputInstance(opt, src, refs); err != nil && err != EOD {
		return nil, err
	}

	// Try drilling to property.
	if src, refs, err = DrillInputProperty(opt, src, refs); err != nil && err != EOD {
		return nil, err
	}

	return src.Copy(), nil
}

func HandleFileOutputScheme(opt *Options, node *OutputNode, filename string, insrc *Source) (err error) {
	// Find extension.
	var ext string
	if node.Format == "" {
		ext = strings.TrimPrefix(filepath.Ext(filename), ".")
		node.Format = ext
	} else {
		ext = node.Format
	}

	// Find format.
	newFormat, exists := registeredFormats[ext]
	if !exists {
		return errors.New("format is not registered")
	}
	format := newFormat(opt)

	refs := node.Reference[1:]

	var file *os.File
	var outsrc *Source
	if len(refs) == 0 {
		// No drilling; content of input overwrites output.
		file, err = os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		// Append input instances to root of output.
		if len(insrc.Properties) > 0 || len(insrc.Values) > 0 {
			return errors.New("cannot map input to file: source must contain only instances")
		}
		outsrc = insrc
	} else {
		// Drilling; open and decode the output file.
		file, err = os.Open(filename)
		if os.IsNotExist(err) {
			return errors.New("cannot drill into file: file does not exist")
		}
		if err != nil {
			return err
		}
		outsrc, err = format.Decode(file)
		file.Close()
		if err != nil {
			return err
		}

		// Drill to instance.
		var outaddr SourceAddress
		if outaddr, refs, err = DrillInstance(opt, outsrc, refs); err != nil {
			return err
		}

		if len(refs) == 0 {
			// No drilling; set properties, and append input as children.
			if len(insrc.Values) > 0 {
				return errors.New("cannot map input to instance: source must not contain values")
			}
			// TODO: Use API to make sure properties are correct.
			if err := outaddr.Set(insrc.Properties); err != nil {
				return err
			}
			if err := outaddr.Set(insrc.Instances); err != nil {
				return err
			}
		} else {
			// Drill to property.
			inst, err := outaddr.Get()
			if err != nil {
				return err
			}
			src := &Source{Instances: []*rbxfile.Instance{inst.(*rbxfile.Instance)}}
			var name string
			if len(refs) > 0 {
				name = refs[0]
			}
			if outaddr, refs, err = DrillProperty(opt, src, refs); err != nil {
				return err
			}

			if len(insrc.Instances) > 0 {
				return errors.New("cannot map input to property: source must not contain instances")
			}
			if len(insrc.Values) == 1 {
				if len(insrc.Properties) > 0 {
					return errors.New("cannot map input to property: source must not contain properties while also containing a value")
				}
				// Map value to property.
				outaddr.Set(insrc.Values[0].Copy())
			} else {
				if len(insrc.Values) > 0 {
					return errors.New("cannot map input to property: source must not contain values while also containing properties")
				}
				// Map property matching name.
				value, exists := insrc.Properties[name]
				if !exists {
					return errors.New("cannot map input to property: cannot find input matching name")
				}
				outaddr.Set(value.Copy())
			}
		}

		// Secure file for writing.
		if file, err = os.Create(filename); err != nil {
			return err
		}
		defer file.Close()
	}

	// Write new output to file.
	outsrc = outsrc.Copy()
	if !format.CanEncode(outsrc) {
		return errors.New("cannot encode transformed output")
	}
	return format.Encode(file, outsrc)
}
