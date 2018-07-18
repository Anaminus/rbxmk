package types

import (
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
	"strconv"
)

type Instances []*rbxfile.Instance

func (indata *Instances) Type() string {
	return "Instances"
}

func (indata *Instances) Drill(opt *rbxmk.Options, inref []string) (outdata rbxmk.Data, outref []string, err error) {
	if len(inref) == 0 {
		err = rbxmk.EOD
		return indata, inref, err
	}

	var instance *rbxfile.Instance
	instances := *indata
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
			err = fmt.Errorf("expected digit")
			goto Error
		}
		j := i
		for ; j < len(ref); j++ {
			if !isDigit(ref[j]) {
				break
			}
		}
		if i == j {
			err = fmt.Errorf("expected digit")
			goto Error
		}
		n, e := strconv.Atoi(ref[i:j])
		if e != nil {
			err = fmt.Errorf("failed to parse %q as number", ref[i:j])
			goto Error
		}
		// Number must be positive (negative shouldn't be possible).
		if n < 0 {
			err = fmt.Errorf("invalid index")
			goto Error
		}
		if instance == nil {
			// Get the nth child from the root.
			if n >= len(instances) {
				err = fmt.Errorf("index exceeds length of parent")
				goto Error
			}
			instance = instances[n]
		} else {
			// Get the nth child from the current parent.
			if n >= len(instance.Children) {
				err = fmt.Errorf("index exceeds length of parent")
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
			err = fmt.Errorf("expected word")
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
		err = fmt.Errorf("indexed child is nil")
		goto Error
	}
	// Finish if end of ref was reached.
	if i >= len(ref) {
		goto Finish
	}
	// Expect `.` separator.
	if ref[i] != '.' {
		err = fmt.Errorf("expected '.' separator")
		goto Error
	}
	i++
	goto CheckRef

Finish:
	if instance == nil {
		return indata, inref, fmt.Errorf("no instance selected")
	}
	return Instance{instance}, inref[1:], nil

Error:
	return indata, inref, ParseError{Index: i, Err: err}
}

func (indata *Instances) Merge(opt *rbxmk.Options, rootdata, drilldata rbxmk.Data) (outdata rbxmk.Data, err error) {
	switch drilldata := drilldata.(type) {
	case *Instances:
		*drilldata = append(*drilldata, *indata...)
		return rootdata, nil

	case Instance:
		for _, child := range *indata {
			drilldata.AddChild(child)
		}
		return rootdata, nil

	case nil:
		return indata, nil
	}
	return nil, rbxmk.NewMergeError(indata, drilldata, nil)
}
