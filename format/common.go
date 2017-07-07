package format

import (
	"errors"
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
	"strconv"
)

var Formats = rbxmk.NewFormats()

// Property is a Data type pointing to a value within a property map.
type Property struct {
	Properties map[string]rbxfile.Value
	Name       string
}

func isAlnum(b byte) bool {
	return ('0' <= b && b <= '9') ||
		('A' <= b && b <= 'Z') ||
		('a' <= b && b <= 'z') ||
		(b == '_')
}

func isDigit(b byte) bool {
	return ('0' <= b && b <= '9')
}

type ParseError struct {
	Index int
	Err   error
}

func (err ParseError) Error() string {
	return fmt.Sprintf("@%d: %s", err.Index, err.Err)
}

func DrillInstance(opt rbxmk.Options, indata rbxmk.Data, inref []string) (outdata rbxmk.Data, outref []string, err error) {
	if len(inref) == 0 {
		err = rbxmk.EOD
		return indata, inref, err
	}

	var instance *rbxfile.Instance
	var instances []*rbxfile.Instance

	switch v := indata.(type) {
	case *[]*rbxfile.Instance:
		instances = *v
	case *rbxfile.Instance:
		if v == nil {
			return indata, inref, fmt.Errorf("*rbxfile.Instance Data cannot be nil")
		}
		instance = v
	default:
		return indata, inref, rbxmk.NewDataTypeError(indata)
	}

	i := 0
	ref := inref[0]
	if ref == "" {
		goto Finish
	}

CheckRef:
	if isDigit(ref[i]) {
		goto ParseIndexedRef
	} else if isAlnum(ref[i]) {
		goto ParseNamedRef
	} else {
		err = fmt.Errorf("unexpected character %q (expected number or word)", ref[i])
		goto Error
	}

ParseIndexedRef:
	// Parse child by index ("0.1.2"; 2nd child of 1st child of 0th child).
	{
		// Parse a number.
		if i >= len(ref) {
			err = errors.New("expected digit")
			goto Error
		}
		j := i
		for ; j < len(ref); j++ {
			if !isDigit(ref[j]) {
				break
			}
		}
		if i == j {
			err = errors.New("expected digit")
			goto Error
		}
		n, e := strconv.Atoi(ref[i:j])
		if e != nil {
			err = fmt.Errorf("failed to parse %q as number", ref[i:j])
			goto Error
		}
		// Number must be positive (negative shouldn't be possible).
		if n < 0 {
			err = errors.New("invalid index")
			goto Error
		}
		if instance == nil {
			// Get the nth child from the root.
			if n >= len(instances) {
				err = errors.New("index exceeds length of parent")
				goto Error
			}
			instance = instances[n]
		} else {
			// Get the nth child from the current parent.
			if n >= len(instance.Children) {
				err = errors.New("index exceeds length of parent")
				goto Error
			}
			instance = instance.Children[n]
		}
		i = j
		goto ParseSep
	}

ParseNamedRef:
	// Parse child by name ("Workspace.Model.Part").
	{
		// Parse a word.
		j := i
		for ; j < len(ref); j++ {
			if !isAlnum(ref[j]) {
				break
			}
		}
		if i == j {
			err = errors.New("expected word")
			goto Error
		}
		name := ref[i:j]
		if instance == nil {
			// Search for child of name in root.
			for _, inst := range instances {
				if inst.Name() == name {
					instance = inst
					break
				}
			}
		} else {
			// Search for child of name in current parent.
			instance = instance.FindFirstChild(name, false)
		}
		i = j
		goto ParseSep
	}

ParseSep:
	// Child must be found.
	if instance == nil {
		err = errors.New("indexed child is nil")
		goto Error
	}
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
	goto CheckRef

Finish:
	if instance == nil {
		return indata, inref, errors.New("no instance selected")
	}
	return instance, inref[1:], nil

Error:
	return indata, inref, ParseError{Index: i, Err: err}
}

func DrillInstanceProperty(opt rbxmk.Options, indata rbxmk.Data, inref []string) (outdata rbxmk.Data, outref []string, err error) {
	if len(inref) == 0 {
		err = rbxmk.EOD
		return indata, inref, err
	}

	var instance *rbxfile.Instance
	switch v := indata.(type) {
	case *[]*rbxfile.Instance:
		if len(*v) == 0 {
			return indata, inref, fmt.Errorf("length of *[]*rbxfile.Instance Data cannot be 0")
		}
		instance = (*v)[0]
	case *rbxfile.Instance:
		if v == nil {
			return indata, inref, fmt.Errorf("*rbxfile.Instance Data cannot be nil")
		}
		instance = v
	default:
		err = rbxmk.NewDataTypeError(indata)
		return indata, inref, err
	}

	ref := inref[0]
	if ref == "" {
		return indata, inref, errors.New("property not specified")
	}
	if ref == "*" {
		// Select all properties.
		return instance.Properties, inref[1:], nil
	}

	// TODO: API?

	if _, exists := instance.Properties[ref]; !exists {
		return indata, inref, fmt.Errorf("property %q not present in instance", ref)
	}
	return Property{Properties: instance.Properties, Name: ref}, inref[1:], nil
}

func DrillProperty(opt rbxmk.Options, indata rbxmk.Data, inref []string) (outdata rbxmk.Data, outref []string, err error) {
	if len(inref) == 0 {
		err = rbxmk.EOD
		return indata, inref, err
	}

	props, ok := indata.(map[string]rbxfile.Value)
	if !ok {
		return indata, inref, rbxmk.NewDataTypeError(indata)
	}
	if _, exists := props[inref[0]]; !exists {
		return indata, inref, fmt.Errorf("property %q not present in instance", inref[0])
	}
	return Property{Properties: props, Name: inref[0]}, inref[1:], nil
}

// MergeOverwrite is a rbxmk.Merger that overrides the output data with
// the input data.
func MergeOverwrite(opt rbxmk.Options, rootdata, drilldata, indata rbxmk.Data) (outdata rbxmk.Data, err error) {
	switch indata.(type) {
	case rbxmk.DeleteData:
		return nil, nil
	}
	return indata, nil
}

func MergeInstance(opt rbxmk.Options, rootdata, drilldata, indata rbxmk.Data) (outdata rbxmk.Data, err error) {
	switch drilldata := drilldata.(type) {
	case nil:
		switch indata := indata.(type) {
		case *[]*rbxfile.Instance:
			return indata, nil
		case *rbxfile.Instance:
			v := []*rbxfile.Instance{indata}
			return &v, nil
		case rbxmk.DeleteData:
			return nil, nil
		}
	case *[]*rbxfile.Instance:
		switch indata := indata.(type) {
		case *[]*rbxfile.Instance:
			*drilldata = append(*drilldata, *indata...)
			return rootdata, nil
		case *rbxfile.Instance:
			*drilldata = append(*drilldata, indata)
			return rootdata, nil
		case rbxmk.DeleteData:
			*drilldata = (*drilldata)[:0]
			return rootdata, nil
		}
	case *rbxfile.Instance:
		switch indata := indata.(type) {
		case *[]*rbxfile.Instance:
			for _, child := range *indata {
				drilldata.AddChild(child)
			}
			return rootdata, nil
		case *rbxfile.Instance:
			drilldata.AddChild(indata)
			return rootdata, nil
		case map[string]rbxfile.Value:
			for name, value := range indata {
				drilldata.Properties[name] = value
			}
			return rootdata, nil
		case Property:
			drilldata.Properties[indata.Name] = indata.Properties[indata.Name]
			return rootdata, nil
		case rbxmk.DeleteData:
			drilldata.SetParent(nil)
			return rootdata, nil
		}
	}
	return MergeProperties(opt, rootdata, drilldata, indata)
}

func MergeProperties(opt rbxmk.Options, rootdata, drilldata, indata rbxmk.Data) (outdata rbxmk.Data, err error) {
	switch drilldata := drilldata.(type) {
	case nil:
		switch indata := indata.(type) {
		case map[string]rbxfile.Value:
			return indata, nil
		case rbxmk.DeleteData:
			return nil, nil
		}
	case map[string]rbxfile.Value:
		switch indata := indata.(type) {
		case map[string]rbxfile.Value:
			for name, value := range indata {
				drilldata[name] = value
			}
			return rootdata, nil
		case Property:
			drilldata[indata.Name] = indata.Properties[indata.Name]
			return rootdata, nil
		case rbxmk.DeleteData:
			for k := range drilldata {
				delete(drilldata, k)
			}
			return rootdata, nil
		}
	case Property:
		switch indata := indata.(type) {
		case Property:
			drilldata.Properties[drilldata.Name] = indata.Properties[indata.Name]
			return rootdata, nil
		case rbxfile.Value:
			drilldata.Properties[drilldata.Name] = indata
			return rootdata, nil
		case rbxmk.DeleteData:
			delete(drilldata.Properties, drilldata.Name)
			return rootdata, nil
		}
	}
	return nil, rbxmk.NewMergeError(drilldata, indata)
}
