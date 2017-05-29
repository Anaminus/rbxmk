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
	u, err := url.Parse(strings.Join(node.Reference, ""))
	if err != nil {
		return
	}

	// Search the end of the URL for a drill separator.
	nextPart := ""
	if u.Fragment != "" {
		i := strings.IndexByte(u.Fragment, ':')
		if i > -1 {
			nextPart = u.Fragment[i+1:]
			u.Fragment = u.Fragment[:i]
		}
	} else if u.RawQuery != "" {
		i := strings.IndexByte(u.RawQuery, ':')
		if i > -1 {
			nextPart = u.RawQuery[i+1:]
			u.RawQuery = u.RawQuery[:i]
		}
	} else if u.Path != "" {
		i := strings.IndexByte(u.Path, ':')
		if i > -1 {
			nextPart = u.Path[i+1:]
			u.Path = u.Path[:i]
		}
	}

	// BUG: URLs cannot contain a ':' without it being detected as a drill.
	// Excludes "host:port" pattern, includes URL-escaped characters.
	//
	// This might be worked around by ignoring URL-escaped characters.
	// However, this requires a custom version of the url package; some parts
	// of the url cannot be acquired in raw format, and the raw parts that can
	// be acquired cannot be unescaped after parsing them.

	// Reconstruct the url without the drill.
	urlPart := u.String()

	_, _ = urlPart, nextPart
	// TODO: get resource; expect a format
	return
}

func parseFilePath(ref, format string) (path, next, ext string) {
	path = ref
	if len(filepath.VolumeName(ref)) == 2 {
		// If path contains drive letter, skip over it.
		i := strings.IndexByte(ref[2:], ':')
		if i > -1 {
			next = ref[i+3:]
			path = ref[:i+2]
		}
	} else {
		i := strings.IndexByte(ref, ':')
		if i > -1 {
			next = ref[i+1:]
			path = ref[:i]
		}
	}

	if format != "" {
		// Format was specified by a flag.
		ext = format
	} else {
		// Guess the format by looking at the file extension.
		ext = strings.TrimPrefix(filepath.Ext(path), ".")
	}
	return
}

func HandleFileInputScheme(opt *Options, node *InputNode, ref string) (src *Source, err error) {
	pathPart, nextPart, ext := parseFilePath(ref, node.Format)
	if ext == "" {
		return nil, errors.New("file must have an extension")
	}
	node.Format = ext

	newFormat, exists := registeredFormats[ext]
	if !exists {
		return nil, errors.New("format is not registered")
	}

	file, err := os.Open(pathPart)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	src, err = newFormat(opt).Decode(file)
	if err != nil {
		return nil, err
	}

	if nextPart == "" {
		return src, nil
	}
	var parent *rbxfile.Instance
	parent, nextPart, err = ParseInstanceReference(src, nextPart)
	if err != nil {
		return nil, err
	}

	if nextPart == "" {
		return &Source{Instances: []*rbxfile.Instance{parent}}, nil
	}
	src, nextPart, err = GetPropertyReference(parent, nextPart)
	return src, err
}

func HandleFileOutputScheme(opt *Options, node *OutputNode, ref string, input *Source) (err error) {
	pathPart, nextPart, ext := parseFilePath(ref, node.Format)
	if ext == "" {
		return errors.New("file must have an extension")
	}
	node.Format = ext

	newFormat, exists := registeredFormats[ext]
	if !exists {
		return errors.New("format is not registered")
	}
	format := newFormat(opt)

	var file *os.File
	var output *Source
	if nextPart == "" {
		file, err = os.Create(pathPart)
		if err != nil {
			return err
		}
		defer file.Close()

		// No drilling; append input instances to root of output.
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
		file, err = os.Open(pathPart)
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
		parent, nextPart, err = ParseInstanceReference(output, nextPart)
		if err != nil {
			return err
		}

		if nextPart == "" {
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
			if err = SetPropertyReference(parent, input, nextPart); err != nil {
				return err
			}
		}

		// Secure file for writing.
		file, err = os.Create(pathPart)
		if err != nil {
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

func GetPropertyReference(input *rbxfile.Instance, ref string) (output *Source, nextPart string, err error) {
	if drill := strings.IndexByte(ref, ':'); drill > -1 {
		ref, nextPart = ref[:drill], ref[drill+1:]
	}
	if ref == "" {
		return nil, "", errors.New("property not specified")
	}
	if ref == "*" {
		// Select all properties.
		output = &Source{Properties: input.Properties}
		return
	}

	// TODO: API?

	v, exists := input.Properties[ref]
	if !exists {
		return nil, "", errors.New("unknown property")
	}
	return &Source{Values: []rbxfile.Value{v}}, nextPart, nil
}

func SetPropertyReference(parent *rbxfile.Instance, input *Source, ref string) (err error) {
	var nextPart string
	if drill := strings.IndexByte(ref, ':'); drill > -1 {
		ref, nextPart = ref[:drill], ref[drill+1:]
	}
	if ref == "" {
		return errors.New("property not specified")
	}

	// TODO: use API to make sure assignment is correct.

	// TODO: drill into properties.
	_ = nextPart

	if len(input.Instances) > 0 {
		return errors.New("cannot map input to property: source must not contain instances")
	}
	if len(input.Values) == 1 {
		if len(input.Properties) > 0 {
			return errors.New("cannot map input to property: source must not contain properties while also containing a value")
		}
		// Map value to property
		parent.Properties[ref] = input.Values[0].Copy()
	} else {
		if len(input.Values) > 0 {
			return errors.New("cannot map input to property: source must not contain values while also containing properties")
		}
		// Map property matching name
		value, exists := input.Properties[ref]
		if !exists {
			return errors.New("cannot map input to property: cannot find input matching name")
		}
		parent.Properties[ref] = value.Copy()
	}
	return nil
}

func ParseInstanceReference(input *Source, ref string) (output *rbxfile.Instance, next string, err error) {
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
	{
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
			for _, inst := range input.Instances {
				if inst.Name() == name {
					output = inst
					break
				}
			}
		} else {
			output = output.FindFirstChild(name, false)
		}
		if output == nil {
			err = errors.New("indexed child is nil")
			goto Error
		}
		i = j
		if i >= len(ref) {
			goto Finish
		}
		switch c := ref[i]; c {
		case '.':
			i++
			goto ParseNamedRef
		case ':':
			i++
			// drill down
			next = ref[i:]
		}
		goto Finish
	}
ParseChildRef:
	{
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
		if n < 0 {
			err = errors.New("invalid index")
			goto Error
		}
		if output == nil {
			if n >= len(input.Instances) {
				err = errors.New("index exceeds length of parent")
				goto Error
			}
			output = input.Instances[n]
		} else {
			if n >= len(output.Children) {
				err = errors.New("index exceeds length of parent")
				goto Error
			}
			output = output.Children[n]
		}
		if output == nil {
			err = errors.New("indexed child is nil")
			goto Error
		}
		i = j
		if i >= len(ref) {
			goto Finish
		}
		switch c := ref[i]; c {
		case '.':
			i++
			goto ParseChildRef
		case ':':
			i++
			// drill down
			next = ref[i:]
		}
		goto Finish
	}
Finish:
	if output == nil {
		err = errors.New("no instance selected")
		goto Error
	}
	return output, next, nil
Error:
	return nil, "", err
}
