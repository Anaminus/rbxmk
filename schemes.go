package main

import (
	"errors"
	"github.com/robloxapi/rbxfile"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
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

func HandleFileInputScheme(opt *Options, node *InputNode, ref string) (src *Source, err error) {
	// Find extension.
	var ext string
	if node.Format == "" {
		ext = strings.TrimPrefix(filepath.Ext(ref), ".")
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
	file, err := os.Open(ref)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode file with format.
	src, err = newFormat(opt).Decode(file)
	if err != nil {
		return nil, err
	}

	refs := node.Reference[1:]
	if len(refs) == 0 {
		// Return contents of file.
		return src, nil
	}

	// Drill down into file.
	var parent *rbxfile.Instance
	if parent, err = ParseInstanceReference(src, refs[0]); err != nil {
		return nil, err
	}

	refs = refs[1:]
	if len(refs) == 0 {
		// Return instance.
		return &Source{Instances: []*rbxfile.Instance{parent}}, nil
	}
	// Drill down into instance.
	src, err = GetPropertyReference(parent, refs[0])

	refs = refs[1:]
	if len(refs) == 0 {
		return src, err
	}

	// TODO: further processing.
	return src, err
}

func HandleFileOutputScheme(opt *Options, node *OutputNode, ref string, input *Source) (err error) {
	// Find extension.
	var ext string
	if node.Format == "" {
		ext = strings.TrimPrefix(filepath.Ext(ref), ".")
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
	var output *Source
	if len(refs) == 0 {
		// No drilling; content of input overwrites output.
		file, err = os.Create(ref)
		if err != nil {
			return err
		}
		defer file.Close()

		// Append input instances to root of output.
		if len(input.Properties) > 0 || len(input.Values) > 0 {
			return errors.New("cannot map input to file: source must contain only instances")
		}
		output = &Source{
			Instances: make([]*rbxfile.Instance, len(input.Instances)),
		}
		for i, inst := range input.Instances {
			output.Instances[i] = inst.Clone()
		}
	} else {
		// Drilling; open and decode the output file.
		file, err = os.Open(ref)
		if os.IsNotExist(err) {
			return errors.New("cannot drill into file: file does not exist")
		}
		if err != nil {
			return err
		}
		output, err = format.Decode(file)
		file.Close()
		if err != nil {
			return err
		}

		// Drill into instance.
		var parent *rbxfile.Instance
		if parent, err = ParseInstanceReference(output, refs[0]); err != nil {
			return err
		}

		refs = refs[1:]
		if len(refs) == 0 {
			// No drilling; set properties, and append input as children.
			if len(input.Values) > 0 {
				return errors.New("cannot map input to instance: source must not contain values")
			}
			for name, value := range input.Properties {
				// TODO: Use API to make sure properties are correct.
				parent.Properties[name] = value.Copy()
			}
			for _, child := range input.Instances {
				child.Clone().SetParent(parent)
			}
		} else {
			// Drill into property.
			if err = SetPropertyReference(parent, input, refs); err != nil {
				return err
			}
		}

		// Secure file for writing.
		if file, err = os.Create(ref); err != nil {
			return err
		}
		defer file.Close()
	}

	// Write new output to file.
	if !format.CanEncode(output) {
		return errors.New("cannot encode transformed output")
	}
	return format.Encode(file, output)
}

func GetPropertyReference(input *rbxfile.Instance, ref string) (output *Source, err error) {
	if ref == "" {
		return nil, errors.New("property not specified")
	}
	if ref == "*" {
		// Select all properties.
		output = &Source{Properties: input.Properties}
		return
	}

	// TODO: API?

	v, exists := input.Properties[ref]
	if !exists {
		return nil, errors.New("unknown property")
	}
	return &Source{Values: []rbxfile.Value{v}}, nil
}

func SetPropertyReference(parent *rbxfile.Instance, input *Source, refs []string) (err error) {
	ref := refs[0]
	if ref == "" {
		return errors.New("property not specified")
	}

	// TODO: use API to make sure assignment is correct.

	// TODO: drill into properties.
	refs = refs[1:]

	if len(input.Instances) > 0 {
		return errors.New("cannot map input to property: source must not contain instances")
	}
	if len(input.Values) == 1 {
		if len(input.Properties) > 0 {
			return errors.New("cannot map input to property: source must not contain properties while also containing a value")
		}
		// Map value to property.
		parent.Properties[ref] = input.Values[0].Copy()
	} else {
		if len(input.Values) > 0 {
			return errors.New("cannot map input to property: source must not contain values while also containing properties")
		}
		// Map property matching name.
		value, exists := input.Properties[ref]
		if !exists {
			return errors.New("cannot map input to property: cannot find input matching name")
		}
		parent.Properties[ref] = value.Copy()
	}
	return nil
}

func ParseInstanceReference(input *Source, ref string) (output *rbxfile.Instance, err error) {
	// TODO: Drop @ convention; parse mixed references by restricting names
	// from starting with digits; names index by name, numbers index the nth
	// child. ALT: ':' signals a number, '.' signals a name.
	i := 0
	if ref == "" {
		goto Finish
	}
	if ref[i] == '@' {
		i++
		goto ParseChildRef
	}
	if i >= len(ref) {
		goto Finish
	}
ParseNamedRef:
	// Parse child by name ("Workspace.Model.Part").
	{
		// Parse a word.
		j := i
		for ; j < len(ref); j++ {
			if !IsAlnum(string(ref[j])) {
				break
			}
		}
		if i == j {
			err = errors.New("expected word")
			goto Error
		}
		name := ref[i:j]
		if output == nil {
			// Search for child of name in root.
			for _, inst := range input.Instances {
				if inst.Name() == name {
					output = inst
					break
				}
			}
		} else {
			// Search for child of name in current parent.
			output = output.FindFirstChild(name, false)
		}
		// Child of name must be found.
		if output == nil {
			err = errors.New("indexed child is nil")
			goto Error
		}
		i = j
		// Finish if end of ref was reached.
		if i >= len(ref) {
			goto Finish
		}
		// Expect `.` separator.
		if ref[i] != '.' {
			err = errors.New("expected '.' separator")
			goto Error
		}
		i++
		goto ParseNamedRef
	}
ParseChildRef:
	// Parse child of the form "@0.1.2" (2nd child of 1st child of 0th child).
	{
		// Parse a number.
		if i >= len(ref) {
			err = errors.New("expected digit")
			goto Error
		}
		j := i
		for ; j < len(ref); j++ {
			if !IsDigit(string(ref[j])) {
				break
			}
		}
		if i == j {
			err = errors.New("expected digit")
			goto Error
		}
		n, e := strconv.Atoi(ref[i:j])
		if e != nil {
			err = errors.New("failed to parse number")
			goto Error
		}
		// Number must be positive (negative shouldn't be possible).
		if n < 0 {
			err = errors.New("invalid index")
			goto Error
		}
		if output == nil {
			// Get the nth child from the root.
			if n >= len(input.Instances) {
				err = errors.New("index exceeds length of parent")
				goto Error
			}
			output = input.Instances[n]
		} else {
			// Get the nth child from the current parent.
			if n >= len(output.Children) {
				err = errors.New("index exceeds length of parent")
				goto Error
			}
			output = output.Children[n]
		}
		// Child must be found.
		if output == nil {
			err = errors.New("indexed child is nil")
			goto Error
		}
		i = j
		// Finish if end of ref was reached.
		if i >= len(ref) {
			goto Finish
		}
		// Expect `.` separator.
		if ref[i] != '.' {
			err = errors.New("expected '.' separator")
			goto Error
		}
		i++
		goto ParseChildRef
	}
Finish:
	if output == nil {
		err = errors.New("no instance selected")
		goto Error
	}
	return output, nil
Error:
	return nil, err
}
