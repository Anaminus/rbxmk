package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	RegisterInputScheme("file", InputScheme{
		Handler: fileInputSchemeHandler,
	})
	RegisterOutputScheme("file", OutputScheme{
		Handler:   fileOutputSchemeHandler,
		Finalizer: fileOutputFinalizer,
	})
}

func fileInputSchemeHandler(opt *Options, node *InputNode, filename string) (ext string, src *Source, err error) {
	// Find extension.
	if node.Format == "" {
		ext = strings.TrimPrefix(filepath.Ext(filename), ".")
	} else {
		ext = node.Format
	}

	// Find format.
	format, exists := DefaultFormats.Init(ext, opt)
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

func fileOutputSchemeHandler(opt *Options, node *OutputNode, filename string) (ext string, src *Source, err error) {
	// Find extension.
	if node.Format == "" {
		ext = strings.TrimPrefix(filepath.Ext(filename), ".")
	} else {
		ext = node.Format
	}

	// Find format.
	format, exists := DefaultFormats.Init(ext, opt)
	if !exists {
		return "", nil, errors.New("format is not registered")
	}

	// Open file.
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return ext, &Source{}, nil
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

func fileOutputFinalizer(opt *Options, node *OutputNode, filename, ext string, outsrc *Source) (err error) {
	format, exists := DefaultFormats.Init(ext, opt)
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
