package types

import (
	"bytes"
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
	"strings"
)

// Stringlike represents string-like Data, which can be one of the following
// types:
//
//     rbxfile.ValueString
//     rbxfile.ValueBinaryString
//     rbxfile.ValueProtectedString
//     string
//     []byte
type Stringlike struct {
	ValueType rbxfile.Type
	Bytes     []byte
}

// NewStringlike tries to convert a rbxmk.Data to a Stringlike. Returns nil if
// it failed.
func NewStringlike(v interface{}) *Stringlike {
	s := &Stringlike{}
	if !s.SetFrom(v) {
		return nil
	}
	return s
}

// GetValue returns the content of the Stringlike as a rbxfile.Value. If force
// is true, then an untyped stringlike is returned as a ValueString instead of
// nil.
func (s *Stringlike) GetValue(force bool) Value {
	switch s.ValueType {
	case rbxfile.TypeString:
		return Value{rbxfile.ValueString(s.Bytes)}
	case rbxfile.TypeBinaryString:
		return Value{rbxfile.ValueBinaryString(s.Bytes)}
	case rbxfile.TypeProtectedString:
		return Value{rbxfile.ValueProtectedString(s.Bytes)}
	}
	if force {
		return Value{rbxfile.ValueString(s.Bytes)}
	}
	return Value{}
}

// GetString returns the content of the Stringlike as a string, if possible.
func (s *Stringlike) GetString() string {
	return string(s.Bytes)
}

// AssignTo assigns the content of the Stringlike to a rbxmk.Data value. The
// data must be a pointer to a string-like value. If force is true, then a
// non-stringlike rbxfile.Value will be overwritten using the Stringlike's
// type.
func (s *Stringlike) AssignTo(v interface{}, force bool) bool {
	switch v := v.(type) {
	case *Value:
		switch (*v).Value.(type) {
		case rbxfile.ValueString:
			(*v).Value = rbxfile.ValueString(s.Bytes)
			return true
		case rbxfile.ValueBinaryString:
			(*v).Value = rbxfile.ValueBinaryString(s.Bytes)
			return true
		case rbxfile.ValueProtectedString:
			(*v).Value = rbxfile.ValueProtectedString(s.Bytes)
			return true
		default:
			if force {
				*v = s.GetValue(true)
				if (*v).Value != nil {
					return true
				}
			}
		}
	case *rbxfile.Value:
		switch (*v).(type) {
		case rbxfile.ValueString:
			*v = rbxfile.ValueString(s.Bytes)
			return true
		case rbxfile.ValueBinaryString:
			*v = rbxfile.ValueBinaryString(s.Bytes)
			return true
		case rbxfile.ValueProtectedString:
			*v = rbxfile.ValueProtectedString(s.Bytes)
			return true
		default:
			if force {
				*v = s.GetValue(true).Value
				if *v != nil {
					return true
				}
			}
		}
	case *Stringlike:
		*v = *s
		return true
	case *string:
		*v = string(s.Bytes)
		return true
	case *[]byte:
		*v = s.Bytes
		return true
	}
	return false
}

func (s *Stringlike) AssignToValue(value *Value, force bool) bool {
	switch (*value).Value.(type) {
	case rbxfile.ValueString:
		(*value).Value = rbxfile.ValueString(s.Bytes)
		return true
	case rbxfile.ValueBinaryString:
		(*value).Value = rbxfile.ValueBinaryString(s.Bytes)
		return true
	case rbxfile.ValueProtectedString:
		(*value).Value = rbxfile.ValueProtectedString(s.Bytes)
		return true
	default:
		if (*value).Value == nil || force {
			*value = s.GetValue(true)
			return true
		}
	}
	return false
}

// SetFrom sets the content of the Stringlike from a rbxmk.Data value, if
// possible.
func (s *Stringlike) SetFrom(v interface{}) bool {
	switch v := v.(type) {
	case Value:
		switch v := v.Value.(type) {
		case rbxfile.ValueString:
			s.ValueType = v.Type()
			s.Bytes = []byte(v)
			return true
		case rbxfile.ValueBinaryString:
			s.ValueType = v.Type()
			s.Bytes = []byte(v)
			return true
		case rbxfile.ValueProtectedString:
			s.ValueType = v.Type()
			s.Bytes = []byte(v)
			return true

		}
	case rbxfile.ValueString:
		s.ValueType = v.Type()
		s.Bytes = []byte(v)
		return true
	case rbxfile.ValueBinaryString:
		s.ValueType = v.Type()
		s.Bytes = []byte(v)
		return true
	case rbxfile.ValueProtectedString:
		s.ValueType = v.Type()
		s.Bytes = []byte(v)
		return true
	case *Stringlike:
		s.ValueType = v.ValueType
		s.Bytes = v.Bytes
		return true
	case string:
		s.ValueType = rbxfile.TypeInvalid
		s.Bytes = []byte(v)
		return true
	case []byte:
		s.ValueType = rbxfile.TypeInvalid
		s.Bytes = v
		return true
	}
	return false
}

func (indata *Stringlike) Type() string {
	return "Stringlike"
}

func (indata *Stringlike) Drill(opt rbxmk.Options, inref []string) (outdata rbxmk.Data, outref []string, err error) {
	if len(inref) == 0 {
		err = rbxmk.EOD
		return indata, inref, err
	}

	region := &Region{Value: indata}

	ref := inref[0]
	section := []string{}
	for i, j := 0, 0; i < len(ref); i++ {
		switch c := ref[i]; {
		case isAlnum(c):
			if i == len(ref)-1 {
				section = append(section, string(ref[j:]))
				j = i + 1
			}
		case c == '.':
			section = append(section, string(ref[j:i]))
			j = i + 1
		case c == '+' && i == len(ref)-1:
			section = append(section, string(ref[j:i]))
			j = i + 1
			region.Append = true
		default:
			return indata, inref, ParseError{Index: i, Err: fmt.Errorf("expected '.' or alphanumeric character")}
		}
	}

	if len(section) == 0 {
		return region, inref[1:], nil
	}

	type tag struct {
		name       []byte
		regA, regB int
		selA, selB int
		sub        []*tag
		parent     *tag
	}
	prefix := bytes.HasPrefix
	root := &tag{
		regA: 0,
		selA: 0,
		selB: len(indata.Bytes),
		regB: len(indata.Bytes),
	}
	current := root
	v := region.Value.Bytes

	i := 0
scanTag:
	name := []byte{}
	end := false
	inline := false
	eof := false
	tagA := 0
	tagB := 0
scanTagLoop:
	for ; i < len(v); i++ {
		j := i
		eq := 0
		switch {
		case prefix(v[i:], []byte("--#")):
			j += 3
			inline = false
		case prefix(v[i:], []byte("--[")):
			j += 3
			for ; prefix(v[j+eq:], []byte("=")); eq++ {
			}
			j += eq
			if !prefix(v[j:], []byte("[#")) {
				i++
				goto scanTag
			}
			j += 2
			inline = true
		default:
			continue scanTagLoop
		}
		if prefix(v[j:], []byte("/")) {
			end = true
			j++
		} else {
			end = false
		}
		k := 0
		for ; j+k < len(v) && isAlnum(v[j+k]); k++ {
		}
		name = v[j : j+k]
		if len(name) == 0 {
			i++
			goto scanTag
		}
		j += k
		if inline {
			if !prefix(v[j:], []byte("]")) {
				i++
				goto scanTag
			}
			j++
			if eq > 0 {
				b := bytes.Repeat([]byte("="), eq)
				if !prefix(v[j:], b) {
					i++
					goto scanTag
				}
				j += len(b)
			}
			if !prefix(v[j:], []byte("]")) {
				i++
				goto scanTag
			}
			j++
		} else {
			if bytes.HasPrefix(v[j:], []byte("\r")) {
				j++
			}
			if bytes.HasPrefix(v[j:], []byte("\n")) {
				j++
			}
		}
		tagA = i
		tagB = j
		i = j
		goto finishScanTag
	}
	eof = true
finishScanTag:
	if !eof {
		if end {
			last := current
			for !bytes.Equal(name, current.name) {
				current = current.parent
				if current == root || current == nil {
					// Unmatched end tag; create one from scratch.
					current = &tag{
						name:   name,
						regA:   tagA,
						selA:   tagA,
						parent: last,
					}
					last.sub = append(last.sub, current)
					break
				}
			}
			current.selB = tagA
			current.regB = tagB
			current = current.parent
			goto scanTag
		} else {
			tag := &tag{
				name: name,
				regA: tagA,
				selA: tagB,
				selB: len(v),
				regB: len(v),
			}
			tag.parent = current
			current.sub = append(current.sub, tag)
			current = tag
			goto scanTag
		}
	}

	var recurseTags func(current *tag, sec []string)
	recurseTags = func(current *tag, sec []string) {
		if len(sec) == 0 {
			region.Range = append(region.Range, RegionRange{
				RegA: current.regA,
				RegB: current.regB,
				SelA: current.selA,
				SelB: current.selB,
			})
			return
		}
		for _, sub := range current.sub {
			if string(sub.name) == sec[0] {
				recurseTags(sub, sec[1:])
			}
		}
	}
	recurseTags(root, section)

	if len(region.Range) == 0 {
		return indata, inref, RegionError(strings.Join(section, "."))
	}
	return region, inref[1:], nil
}

func (indata *Stringlike) Merge(opt rbxmk.Options, rootdata, drilldata rbxmk.Data) (outdata rbxmk.Data, err error) {
	switch drilldata := drilldata.(type) {
	case Property:
		v := Value{drilldata.Properties[drilldata.Name]}
		if !indata.AssignToValue(&v, false) {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("output property \"%s\" must be stringlike", drilldata.Name))
		}
		drilldata.Properties[drilldata.Name] = v.Value
		return rootdata, nil

	case *Region:
		drilldata.Set(indata.Bytes)
		if drilldata.Property == nil {
			return drilldata.Value, nil
		}
		return rootdata, nil

	case *Stringlike:
		drilldata.SetFrom(indata)
		return drilldata, nil

	case Value:
		if !indata.AssignToValue(&drilldata, false) {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("output value must be stringlike"))
		}
		return drilldata, nil

	case nil:
		return indata, nil
	}
	return nil, rbxmk.NewMergeError(indata, drilldata, nil)
}

// ProcessStringlikeCallback receives and modifies a Stringlike value.
type ProcessStringlikeCallback func(s *Stringlike) error

// ProcessStringlikeInterface receives an arbitrary value, and converts it
// into a types.Stringlike, if possible. The Stringlike is processed by a
// callback, and the result is applied back to the original location of the
// value. Returns the result as a rbxmk.Data.
//
// The following types are handled:
//
//     - types.Instance (Script, LocalScript, or ModuleScript; modifies the Source property)
//     - *rbxfile.Instances (modifies each instance)
//     - types.Property (any string-like property)
//     - types.Value (any string-like value)
//     - *types.Stringlike
//     - string (returns as a Stringlike)
//     - []byte (returns as a Stringlike)
func ProcessStringlikeInterface(cb ProcessStringlikeCallback, v interface{}) (out rbxmk.Data, err error) {
	switch v := v.(type) {
	case rbxmk.Data:
		switch v := v.(type) {
		case *Instances:
			for _, inst := range *v {
				if err := processStringlikeInstance(cb, inst, false); err != nil {
					return nil, err
				}
			}
			return v, nil
		case Instance:
			if err := processStringlikeInstance(cb, v.Instance, true); err != nil {
				return nil, err
			}
			return v, nil
		case Property:
			value, err := processStringlikeValue(cb, Value{v.Properties[v.Name]})
			if err != nil {
				return nil, err
			}
			v.Properties[v.Name] = value.Value
			return v, nil
		case Value:
			return processStringlikeValue(cb, v)
		case *Stringlike:
			if err := cb(v); err != nil {
				return nil, err
			}
			return v, nil
		default:
			return nil, rbxmk.NewDataTypeError(v)
		}
	case string, []byte:
		s := NewStringlike(v)
		if err := cb(s); err != nil {
			return nil, err
		}
		return s, nil
	case nil:
		return nil, nil
	}
	return nil, fmt.Errorf("unexpected type")
}

func processStringlikeInstance(cb ProcessStringlikeCallback, inst *rbxfile.Instance, fail bool) (err error) {
	switch inst.ClassName {
	case "Script", "LocalScript", "ModuleScript":
		if source, ok := inst.Properties["Source"]; ok {
			value, _ := processStringlikeValue(cb, Value{source})
			inst.Properties["Source"] = value.Value
		}
		return nil
	}
	if fail {
		return fmt.Errorf("instance must be script-like")
	}
	return nil
}

func processStringlikeValue(cb ProcessStringlikeCallback, value Value) (out Value, err error) {
	if s := NewStringlike(value); s != nil {
		if err := cb(s); err != nil {
			return out, err
		}
		return s.GetValue(true), nil
	}
	return out, fmt.Errorf("value must be string-like")
}
