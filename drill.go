package main

import (
	"errors"
	"fmt"
	"github.com/robloxapi/rbxfile"
	"strconv"
)

// SourceAddress points to data within a Source. Data is allowed to be
// resolved as the address is created. For example, an address can point
// directly to an instance within a tree in the Source. Note that this means
// the address may no longer point to the expected location, if the data is
// moved.
type SourceAddress interface {
	// Returns the data being pointed to.
	Get() (v interface{}, err error)
	// Sets the data being pointed to. This may include modifying the current
	// data, rather than replacing it.
	Set(v interface{}) (err error)
}

type addrInstance struct {
	src  *Source
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
	src   *Source
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
	src  *Source
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
	src   *Source
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

// Drill receives a Source and drills into it using inref, returning a
// SourceAddress which points to the resulting data within insrc. It also
// returns the reference after it has been parsed. In case of an error, inref
// is returned. If inref is empty, then an EOD error is returned.
type Drill func(opt *Options, insrc *Source, inref []string) (outaddr SourceAddress, outref []string, err error)

var EOD = errors.New("end of drill")

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

func DrillInstance(opt *Options, insrc *Source, inref []string) (outaddr SourceAddress, outref []string, err error) {
	if len(inref) == 0 {
		err = EOD
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

func DrillProperty(opt *Options, insrc *Source, inref []string) (outaddr SourceAddress, outref []string, err error) {
	if len(inref) == 0 {
		err = EOD
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

func DrillInputInstance(opt *Options, insrc *Source, inref []string) (outsrc *Source, outref []string, err error) {
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
	return &Source{Instances: []*rbxfile.Instance{inst}}, outref, nil
}

func DrillInputProperty(opt *Options, insrc *Source, inref []string) (outsrc *Source, outref []string, err error) {
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
		return &Source{Properties: v}, outref, nil
	case rbxfile.Value:
		return &Source{Values: []rbxfile.Value{v}}, outref, nil
	default:
		panic("unexpected value returned from DrillProperty")
	}
}
