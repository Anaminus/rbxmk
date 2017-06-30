package format

import (
	"errors"
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
	"reflect"
	"strconv"
)

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

type DataTypeError struct {
	dataName string
}

func (err DataTypeError) Error() string {
	return fmt.Sprintf("unexpected Data type: %s", err.dataName)
}

func NewDataTypeError(data rbxmk.Data) error {
	return DataTypeError{dataName: reflect.TypeOf(data).String()}
}

func DrillInstance(opt rbxmk.Options, indata rbxmk.Data, inref []string) (outdata rbxmk.Data, outref []string, err error) {
	if len(inref) == 0 {
		err = rbxmk.EOD
		return indata, inref, err
	}

	var instance *rbxfile.Instance
	var instances []*rbxfile.Instance

	switch v := indata.(type) {
	case []*rbxfile.Instance:
		instances = v
	case *rbxfile.Instance:
		if v == nil {
			return indata, inref, fmt.Errorf("*rbxfile.Instance Data cannot be nil")
		}
		instance = v
	default:
		return indata, inref, NewDataTypeError(indata)
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
	case []*rbxfile.Instance:
		if len(v) == 0 {
			return indata, inref, fmt.Errorf("length of []*rbxfile.Instance Data cannot be 0")
		}
		instance = v[0]
	case *rbxfile.Instance:
		if v == nil {
			return indata, inref, fmt.Errorf("*rbxfile.Instance Data cannot be nil")
		}
		instance = v
	default:
		err = NewDataTypeError(indata)
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
		return indata, inref, NewDataTypeError(indata)
	}
	if _, exists := props[inref[0]]; !exists {
		return indata, inref, fmt.Errorf("property %q not present in instance", inref[0])
	}
	return Property{Properties: props, Name: inref[0]}, inref[1:], nil
}

// ResolveOverwrite is a rbxmk.Resolver that overrides the output data with
// the input data.
func ResolveOverwrite(opt rbxmk.Options, indata, data rbxmk.Data) (outdata rbxmk.Data, err error) {
	return data, nil
}

func ResolveInstance(opt rbxmk.Options, indata, data rbxmk.Data) (outdata rbxmk.Data, err error) {
	switch indata := indata.(type) {
	case nil:
		switch data := data.(type) {
		case []*rbxfile.Instance:
			return data, nil
		case *rbxfile.Instance:
			return []*rbxfile.Instance{data}, nil
		}
	case []*rbxfile.Instance:
		switch data := data.(type) {
		case []*rbxfile.Instance:
			return append(indata, data...), nil
		case *rbxfile.Instance:
			return append(indata, data), nil
		}
	case *rbxfile.Instance:
		switch data := data.(type) {
		case []*rbxfile.Instance:
			for _, child := range data {
				indata.AddChild(child)
			}
			return []*rbxfile.Instance{indata}, nil
		case *rbxfile.Instance:
			indata.AddChild(data)
			return []*rbxfile.Instance{indata}, nil
		case map[string]rbxfile.Value:
			for name, value := range data {
				indata.Properties[name] = value
			}
			return []*rbxfile.Instance{indata}, nil
		case Property:
			indata.Properties[data.Name] = data.Properties[data.Name]
			return []*rbxfile.Instance{indata}, nil
		}
	}
	return ResolveProperties(opt, indata, data)
}

func ResolveProperties(opt rbxmk.Options, indata, data rbxmk.Data) (outdata rbxmk.Data, err error) {
	switch indata := indata.(type) {
	case nil:
		switch data := data.(type) {
		case map[string]rbxfile.Value:
			return data, nil
		}
	case map[string]rbxfile.Value:
		switch data := data.(type) {
		case map[string]rbxfile.Value:
			for name, value := range data {
				indata[name] = value
			}
			return indata, nil
		case Property:
			indata[data.Name] = data.Properties[data.Name]
			return indata, nil
		}
	case Property:
		switch data := data.(type) {
		case Property:
			indata.Properties[indata.Name] = data.Properties[data.Name]
			return indata, nil
		case rbxfile.Value:
			indata.Properties[indata.Name] = data
			return indata, nil
		}
	}
	return nil, NewDataTypeError(indata)
}
