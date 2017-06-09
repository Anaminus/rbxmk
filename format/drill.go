package format

import (
	"errors"
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
	"strconv"
)

type addrInstance struct {
	src  *rbxmk.Source
	inst *rbxfile.Instance
}

func (a addrInstance) Get() (v interface{}, err error) {
	return a.inst, nil
}

func (a addrInstance) Set(v interface{}) (err error) {
	switch v := v.(type) {
	case map[string]rbxfile.Value:
		for name, value := range v {
			a.inst.Properties[name] = value
		}
	case *rbxfile.Instance:
		if err := a.inst.AddChild(v); err != nil {
			return err
		}
	case []*rbxfile.Instance:
		for _, child := range v {
			if err := a.inst.AddChild(child); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("received unexpected type")
	}
	return nil
}

type addrProperty struct {
	src   *rbxmk.Source
	props map[string]rbxfile.Value
	name  string
}

func (a addrProperty) Get() (v interface{}, err error) {
	return a.props[a.name], nil
}

func (a addrProperty) Set(v interface{}) (err error) {
	value, ok := v.(rbxfile.Value)
	if !ok {
		return fmt.Errorf("expected rbxfile.Value")
	}
	a.props[a.name] = value
	return nil
}

type addrProperties struct {
	src  *rbxmk.Source
	inst *rbxfile.Instance
}

func (a addrProperties) Get() (v interface{}, err error) {
	return a.inst.Properties, nil
}

func (a addrProperties) Set(v interface{}) (err error) {
	props, ok := v.(map[string]rbxfile.Value)
	if !ok {
		return fmt.Errorf("expected map[string]rbxfile.Value")
	}
	a.inst.Properties = props
	return nil
}

type addrValue struct {
	src   *rbxmk.Source
	index int
}

func (a addrValue) Get() (v interface{}, err error) {
	if a.index < 0 || a.index >= len(a.src.Values) {
		return nil, fmt.Errorf("cannot get value at address")
	}
	return a.src.Values[a.index], nil
}

func (a addrValue) Set(v interface{}) (err error) {
	value, ok := v.(rbxfile.Value)
	if !ok {
		return fmt.Errorf("expected rbxfile.Value")
	}
	if a.index < 0 {
		return fmt.Errorf("cannot set value at address")
	}
	if a.index == len(a.src.Values) {
		a.src.Values = append(a.src.Values, value)
		return nil
	}
	a.src.Values[a.index] = value
	return nil
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

func DrillInstance(opt *rbxmk.Options, insrc *rbxmk.Source, inref []string) (outaddr rbxmk.SourceAddress, outref []string, err error) {
	if len(inref) == 0 {
		err = rbxmk.EOD
		return nil, inref, err
	}
	if len(insrc.Instances) == 0 {
		err = fmt.Errorf("source must contain at least one instance")
		return nil, inref, err
	}
	if len(insrc.Properties) > 0 || len(insrc.Values) > 0 {
		err = fmt.Errorf("source must contain only instances")
		return nil, inref, err
	}

	var instance *rbxfile.Instance
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
			if n >= len(insrc.Instances) {
				err = errors.New("index exceeds length of parent")
				goto Error
			}
			instance = insrc.Instances[n]
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
			for _, inst := range insrc.Instances {
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
		return nil, inref, errors.New("no instance selected")
	}
	return addrInstance{src: insrc, inst: instance}, inref[1:], nil

Error:
	return nil, inref, ParseError{Index: i, Err: err}
}

func DrillProperty(opt *rbxmk.Options, insrc *rbxmk.Source, inref []string) (outaddr rbxmk.SourceAddress, outref []string, err error) {
	if len(inref) == 0 {
		err = rbxmk.EOD
		return nil, inref, err
	}
	if len(insrc.Instances) != 1 {
		err = fmt.Errorf("source must contain exactly one instance")
		return nil, inref, err
	}
	if len(insrc.Properties) > 0 || len(insrc.Values) > 0 {
		err = fmt.Errorf("source must contain only instances")
		return nil, inref, err
	}

	ref := inref[0]
	if ref == "" {
		return nil, inref, errors.New("property not specified")
	}
	if ref == "*" {
		// Select all properties.
		return addrProperties{src: insrc, inst: insrc.Instances[0]}, inref[1:], nil
	}

	// TODO: API?

	if _, exists := insrc.Instances[0].Properties[ref]; !exists {
		return nil, inref, fmt.Errorf("property %q not present in instance", ref)
	}

	return addrProperty{src: insrc, props: insrc.Instances[0].Properties, name: ref}, inref[1:], nil
}

func DrillInputInstance(opt *rbxmk.Options, insrc *rbxmk.Source, inref []string) (outsrc *rbxmk.Source, outref []string, err error) {
	outaddr, outref, err := DrillInstance(opt, insrc, inref)
	if err != nil {
		return insrc, inref, err
	}
	v, err := outaddr.Get()
	if err != nil {
		return insrc, inref, err
	}
	inst, ok := v.(*rbxfile.Instance)
	if !ok {
		panic("unexpected value returned from DrillInstance")
	}
	return &rbxmk.Source{Instances: []*rbxfile.Instance{inst}}, outref, nil
}

func DrillInputProperty(opt *rbxmk.Options, insrc *rbxmk.Source, inref []string) (outsrc *rbxmk.Source, outref []string, err error) {
	outaddr, outref, err := DrillProperty(opt, insrc, inref)
	if err != nil {
		return insrc, inref, err
	}
	v, err := outaddr.Get()
	if err != nil {
		return insrc, inref, err
	}
	switch v := v.(type) {
	case map[string]rbxfile.Value:
		return &rbxmk.Source{Properties: v}, outref, nil
	case rbxfile.Value:
		return &rbxmk.Source{Values: []rbxfile.Value{v}}, outref, nil
	default:
		panic("unexpected value returned from DrillProperty")
	}
}

func DrillOutputInstance(opt *rbxmk.Options, inaddr rbxmk.SourceAddress, inref []string) (outaddr rbxmk.SourceAddress, outref []string, err error) {
	v, err := inaddr.Get()
	inst := v.([]*rbxfile.Instance)
	insrc := &rbxmk.Source{Instances: inst}
	return DrillInstance(opt, insrc, inref)
}

func DrillOutputProperty(opt *rbxmk.Options, inaddr rbxmk.SourceAddress, inref []string) (outaddr rbxmk.SourceAddress, outref []string, err error) {
	v, err := inaddr.Get()
	inst := v.(*rbxfile.Instance)
	insrc := &rbxmk.Source{Instances: []*rbxfile.Instance{inst}}
	return DrillProperty(opt, insrc, inref)
}

func ResolveOutputSource(ref []string, addr rbxmk.SourceAddress, src *rbxmk.Source) (err error) {
	// addrSource
	return addr.Set(src)
}

func ResolveOutputInstance(ref []string, addr rbxmk.SourceAddress, src *rbxmk.Source) (err error) {
	switch len(ref) {
	case 1: // addrSource
		// Content of input overwrites output.
		if len(src.Properties) > 0 || len(src.Values) > 0 {
			return errors.New("cannot map input to file: source must contain only instances")
		}
		return addr.Set(src)
	case 2: // addrInstance
		// No drilling; set properties, and append input as children.
		if len(src.Values) > 0 {
			return errors.New("cannot map input to instance: source must not contain values")
		}
		// TODO: Use API to make sure properties are correct.
		if err := addr.Set(src.Properties); err != nil {
			return err
		}
		if err := addr.Set(src.Instances); err != nil {
			return err
		}
		return nil
	case 3: // addrProperty / addrProperties
		if len(src.Instances) > 0 {
			return errors.New("cannot map input to property: source must not contain instances")
		}
		if len(src.Values) == 1 {
			if len(src.Properties) > 0 {
				return errors.New("cannot map input to property: source must not contain properties while also containing a value")
			}
			// Map value to property.
			if err := addr.Set(src.Values[0].Copy()); err != nil {
				return err
			}
		} else {
			if len(src.Values) > 0 {
				return errors.New("cannot map input to property: rbxmk.source must not contain values while also containing properties")
			}
			// Map property matching name.
			value, exists := src.Properties[ref[2]]
			if !exists {
				return errors.New("cannot map input to property: cannot find input matching name")
			}
			if err := addr.Set(value.Copy()); err != nil {
				return err
			}
		}
	}
	return nil
}
