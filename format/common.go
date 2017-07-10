package format

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxapi"
	"github.com/robloxapi/rbxfile"
	"strconv"
	"strings"
)

var Formats = rbxmk.NewFormats()

// Stringlike represents string-like Data, which can be one of the following
// types:
//
//     rbxfile.ValueString
//     rbxfile.ValueBinaryString
//     rbxfile.ValueProtectedString
//     string
//     []byte
type Stringlike struct {
	Type  rbxfile.Type
	Bytes []byte
}

// NewStringlike tries to convert a rbxmk.Data to a Stringlike. Returns nil if
// it failed.
func NewStringlike(data rbxmk.Data) *Stringlike {
	s := &Stringlike{}
	if !s.SetFrom(data) {
		return nil
	}
	return s
}

// GetValue returns the content of the Stringlike as a rbxfile.Value. If force
// is true, then an untyped stringlike is returned as a ValueString instead of
// nil.
func (s *Stringlike) GetValue(force bool) rbxfile.Value {
	switch s.Type {
	case rbxfile.TypeString:
		return rbxfile.ValueString(s.Bytes)
	case rbxfile.TypeBinaryString:
		return rbxfile.ValueBinaryString(s.Bytes)
	case rbxfile.TypeProtectedString:
		return rbxfile.ValueProtectedString(s.Bytes)
	}
	if force {
		return rbxfile.ValueString(s.Bytes)
	}
	return nil
}

// GetString returns the content of the Stringlike as a string, if possible.
func (s *Stringlike) GetString() string {
	return string(s.Bytes)
}

// AssignTo assigns the content of the Stringlike to a rbxmk.Data value. The
// data must be a pointer to a string-like value. If force is true, then a
// non-stringlike rbxfile.Value will be overwritten using the Stringlike's
// type.
func (s *Stringlike) AssignTo(data rbxmk.Data, force bool) bool {
	switch data := data.(type) {
	case *rbxfile.Value:
		switch (*data).(type) {
		case rbxfile.ValueString:
			*data = rbxfile.ValueString(s.Bytes)
			return true
		case rbxfile.ValueBinaryString:
			*data = rbxfile.ValueBinaryString(s.Bytes)
			return true
		case rbxfile.ValueProtectedString:
			*data = rbxfile.ValueProtectedString(s.Bytes)
			return true
		default:
			if force {
				*data = s.GetValue(true)
				if *data != nil {
					return true
				}
			}
		}
	case *Stringlike:
		*data = *s
		return true
	case *string:
		*data = string(s.Bytes)
		return true
	case *[]byte:
		*data = s.Bytes
		return true
	}
	return false
}

func (s *Stringlike) AssignToValue(value *rbxfile.Value, force bool) bool {
	switch (*value).(type) {
	case rbxfile.ValueString:
		*value = rbxfile.ValueString(s.Bytes)
		return true
	case rbxfile.ValueBinaryString:
		*value = rbxfile.ValueBinaryString(s.Bytes)
		return true
	case rbxfile.ValueProtectedString:
		*value = rbxfile.ValueProtectedString(s.Bytes)
		return true
	default:
		if *value == nil || force {
			*value = s.GetValue(true)
			return true
		}
	}
	return false
}

// SetFrom sets the content of the Stringlike from a rbxmk.Data value, if
// possible.
func (s *Stringlike) SetFrom(data rbxmk.Data) bool {
	switch data := data.(type) {
	case rbxfile.ValueString:
		s.Type = data.Type()
		s.Bytes = []byte(data)
		return true
	case rbxfile.ValueBinaryString:
		s.Type = data.Type()
		s.Bytes = []byte(data)
		return true
	case rbxfile.ValueProtectedString:
		s.Type = data.Type()
		s.Bytes = []byte(data)
		return true
	case *Stringlike:
		s.Type = data.Type
		s.Bytes = data.Bytes
		return true
	case string:
		s.Type = rbxfile.TypeInvalid
		s.Bytes = []byte(data)
		return true
	case []byte:
		s.Type = rbxfile.TypeInvalid
		s.Bytes = data
		return true
	}
	return false
}

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

type Region struct {
	Property   *Property
	Value      *Stringlike
	RegA, RegB int
	SelA, SelB int
	Append     bool
}

func (r *Region) Get() []byte {
	if r.Value == nil {
		return nil
	}
	return r.Value.Bytes[r.SelA:r.SelB]
}

func (r *Region) GetStringlike() *Stringlike {
	s := &Stringlike{}
	if r.Value == nil {
		return s
	}
	s.Type = r.Value.Type
	s.Bytes = r.Value.Bytes[r.SelA:r.SelB]
	return s
}

func (r *Region) Set(p []byte) {
	if r.Value == nil {
		r.Value = &Stringlike{Bytes: p}
	} else {
		var prefix []byte
		var suffix []byte
		if r.Append {
			prefix = r.Value.Bytes[:r.SelB]
			suffix = r.Value.Bytes[r.SelB:]
			r.RegA = r.RegA
			r.SelA = r.SelA
			r.SelB = r.SelB + len(p)
			r.RegB = r.RegB + len(p)
		} else {
			prefix = r.Value.Bytes[:r.RegA]
			suffix = r.Value.Bytes[r.RegB:]
			r.RegA = r.RegA
			r.SelA = r.RegA
			r.SelB = r.SelA + len(p)
			r.RegB = r.SelB
		}
		b := make([]byte, len(prefix)+len(p)+len(suffix))
		copy(b[0:], prefix)
		copy(b[len(prefix):], p)
		copy(b[len(prefix)+len(p):], suffix)

		r.Value.Bytes = b
	}
	if r.Property != nil {
		value := r.Property.Properties[r.Property.Name]
		r.Value.AssignToValue(&value, true)
		r.Property.Properties[r.Property.Name] = value
	}
}

func DrillRegion(opt rbxmk.Options, indata rbxmk.Data, inref []string) (outdata rbxmk.Data, outref []string, err error) {
	if len(inref) == 0 {
		err = rbxmk.EOD
		return indata, inref, err
	}

	region := &Region{}
	switch v := indata.(type) {
	case Property:
		region.Property = &v
		region.Value = &Stringlike{}
		region.Value.SetFrom(v.Properties[v.Name])
	case *Stringlike:
		region.Value = v
	}

	region.RegA = 0
	region.SelA = 0
	region.SelB = len(region.Value.Bytes)
	region.RegB = len(region.Value.Bytes)

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
	root := &tag{}
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
		} else if !end {
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

	current = root
loop:
	for _, sec := range section {
		for _, tag := range current.sub {
			if string(tag.name) == sec {
				current = tag
				continue loop
			}
		}
		current = root
		break
	}
	if current == root || current == nil {
		return indata, inref, fmt.Errorf("failed to find region \"%s\"", strings.Join(section, "."))
	}
	region.RegA = current.regA
	region.RegB = current.regB
	region.SelA = current.selA
	region.SelB = current.selB

	return region, inref[1:], nil
}

func typeOfProperty(api *rbxapi.API, className, propName string) rbxfile.Type {
	if api == nil {
		return rbxfile.TypeInvalid
	}
	class := api.Classes[className]
	if class == nil {
		return rbxfile.TypeInvalid
	}
	prop, ok := class.Members[propName].(*rbxapi.Property)
	if !ok {
		return rbxfile.TypeInvalid
	}
	return rbxfile.TypeFromAPIString(api, prop.ValueType)
}

func propertyIsOfType(api *rbxapi.API, inst *rbxfile.Instance, propName string, typ rbxfile.Type) bool {
	if api == nil {
		v, ok := inst.Properties[propName]
		if !ok {
			// Type cannot be determined, assume given type is correct.
			return true
		}
		return v.Type() == typ
	}
	class := api.Classes[inst.ClassName]
	if class == nil {
		// Unknown class, assume given type is correct.
		return true
	}
	member := class.Members[propName]
	prop, ok := member.(*rbxapi.Property)
	if !ok {
		if member != nil {
			// Incorrect member type.
			return false
		}
		// Unknown property, assume given type is correct.
		return true
	}
	return rbxfile.TypeFromAPIString(api, prop.ValueType) == typ
}

func MergeTable(opt rbxmk.Options, rootdata, drilldata, indata rbxmk.Data) (outdata rbxmk.Data, err error) {
	switch indata := indata.(type) {
	case *[]*rbxfile.Instance:
		return mergeInstances(opt, rootdata, drilldata, indata)
	case *rbxfile.Instance:
		return mergeInstance(opt, rootdata, drilldata, indata)
	case map[string]rbxfile.Value:
		return mergeProperties(opt, rootdata, drilldata, indata)
	case Property:
		return mergeProperty(opt, rootdata, drilldata, indata)
	case *Region:
		return mergeRegion(opt, rootdata, drilldata, indata)
	case *Stringlike:
		return mergeStringlike(opt, rootdata, drilldata, indata)
	case rbxfile.Value:
		return mergeValue(opt, rootdata, drilldata, indata)
	case rbxmk.DeleteData:
		return mergeDelete(opt, rootdata, drilldata, indata)
	case nil:
		return rootdata, nil
	}
	return nil, rbxmk.NewMergeError(indata, drilldata, nil)
}

func mergeInstances(opt rbxmk.Options, rootdata, drilldata rbxmk.Data, indata *[]*rbxfile.Instance) (outdata rbxmk.Data, err error) {
	switch drilldata := drilldata.(type) {
	case *[]*rbxfile.Instance:
		*drilldata = append(*drilldata, *indata...)
		return rootdata, nil

	case *rbxfile.Instance:
		for _, child := range *indata {
			drilldata.AddChild(child)
		}
		return rootdata, nil

	case map[string]rbxfile.Value:
		// Incompatible

	case Property:
		// Incompatible

	case *Stringlike:
		// Incompatible

	case rbxfile.Value:
		// Incompatible

	case *Region:
		// Incompatible

	case nil:
		return indata, nil
	}
	return nil, rbxmk.NewMergeError(indata, drilldata, nil)
}

func mergeInstance(opt rbxmk.Options, rootdata, drilldata rbxmk.Data, indata *rbxfile.Instance) (outdata rbxmk.Data, err error) {
	switch drilldata := drilldata.(type) {
	case *[]*rbxfile.Instance:
		*drilldata = append(*drilldata, indata)
		return rootdata, nil

	case *rbxfile.Instance:
		drilldata.AddChild(indata)
		return rootdata, nil

	case map[string]rbxfile.Value:
		// Incompatible

	case Property:
		if typeOfProperty(opt.Config.API, indata.ClassName, drilldata.Name) == rbxfile.TypeReference ||
			drilldata.Properties[drilldata.Name].Type() != rbxfile.TypeReference {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("property must be a Reference"))
		}
		drilldata.Properties[drilldata.Name] = rbxfile.ValueReference{Instance: indata}
		return rootdata, nil

	case *Region:
		// Incompatible

	case *Stringlike:
		// Incompatible

	case rbxfile.Value:
		// Incompatible

	case nil:
		return indata, nil
	}
	return nil, rbxmk.NewMergeError(indata, drilldata, nil)
}

func mergeProperties(opt rbxmk.Options, rootdata, drilldata rbxmk.Data, indata map[string]rbxfile.Value) (outdata rbxmk.Data, err error) {
	switch drilldata := drilldata.(type) {
	case *[]*rbxfile.Instance:
		for _, inst := range *drilldata {
			for name, value := range indata {
				if propertyIsOfType(opt.Config.API, inst, name, value.Type()) {
					inst.Properties[name] = value
				}
			}
		}
		return rootdata, nil

	case *rbxfile.Instance:
		for name, value := range indata {
			if propertyIsOfType(opt.Config.API, drilldata, name, value.Type()) {
				drilldata.Properties[name] = value
			}
		}
		return rootdata, nil

	case map[string]rbxfile.Value:
		for name, value := range indata {
			if v, _ := drilldata[name]; v == nil || v.Type() == value.Type() {
				drilldata[name] = value
			}
		}
		return rootdata, nil

	case Property:
		value, ok := indata[drilldata.Name]
		if !ok {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("\"%s\" not found in properties", drilldata.Name))
		}
		drilldata.Properties[drilldata.Name] = value
		return rootdata, nil

	case *Region:
		if drilldata.Property == nil {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("Region must have Property"))
		}
		value, ok := indata[drilldata.Property.Name]
		if !ok {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("\"%s\" not found in properties", drilldata.Property.Name))
		}
		s := NewStringlike(value)
		if s == nil {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("value of \"%s\" must be stringlike", drilldata.Property.Name))
		}
		drilldata.Set(s.Bytes)
		return rootdata, nil

	case *Stringlike:
		// Incompatible

	case rbxfile.Value:
		// Incompatible

	case nil:
		return indata, nil
	}
	return nil, rbxmk.NewMergeError(indata, drilldata, nil)
}

func mergeProperty(opt rbxmk.Options, rootdata, drilldata rbxmk.Data, indata Property) (outdata rbxmk.Data, err error) {
	value := indata.Properties[indata.Name]
	if value == nil {
		return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("input property \"%s\" cannot be nil", indata.Name))
	}
	switch drilldata := drilldata.(type) {
	case *[]*rbxfile.Instance:
		for _, inst := range *drilldata {
			if propertyIsOfType(opt.Config.API, inst, indata.Name, value.Type()) {
				inst.Properties[indata.Name] = value
			}
		}
		return rootdata, nil

	case *rbxfile.Instance:
		if propertyIsOfType(opt.Config.API, drilldata, indata.Name, value.Type()) {
			drilldata.Properties[indata.Name] = value
		}
		return rootdata, nil

	case map[string]rbxfile.Value:
		if v, _ := drilldata[indata.Name]; v == nil || v.Type() == value.Type() {
			drilldata[indata.Name] = value
		}
		return rootdata, nil

	case Property:
		if v, _ := drilldata.Properties[drilldata.Name]; v == nil || v.Type() == value.Type() {
			drilldata.Properties[drilldata.Name] = value
		} else {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf(
				"input property \"%s\" cannot be assigned to output property \"%s\": expected %s, got %s",
				indata.Name, drilldata.Name, v.Type(), value.Type(),
			))
		}
		return rootdata, nil

	case *Region:
		s := NewStringlike(value)
		if s == nil {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("input property \"%s\" must be stringlike", indata.Name))
		}
		drilldata.Set(s.Bytes)
		if drilldata.Property == nil {
			return drilldata.Value, nil
		}
		return rootdata, nil

	case *Stringlike:
		if !drilldata.SetFrom(value) {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("input property \"%s\" must be stringlike", indata.Name))
		}
		return drilldata, nil

	case rbxfile.Value:
		if drilldata.Type() != value.Type() {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("expected input type %s, got %s", drilldata.Type(), value.Type()))
		}
		return value, nil

	case nil:
		return indata.Properties[indata.Name], nil
	}
	return nil, rbxmk.NewMergeError(indata, drilldata, nil)
}

func mergeStringlike(opt rbxmk.Options, rootdata, drilldata rbxmk.Data, indata *Stringlike) (outdata rbxmk.Data, err error) {
	switch drilldata := drilldata.(type) {
	case *[]*rbxfile.Instance:
		// Incompatible

	case *rbxfile.Instance:
		// Incompatible

	case map[string]rbxfile.Value:
		// Incompatible

	case Property:
		v := drilldata.Properties[drilldata.Name]
		if !indata.AssignToValue(&v, false) {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("output property \"%s\" must be stringlike", drilldata.Name))
		}
		drilldata.Properties[drilldata.Name] = v
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

	case rbxfile.Value:
		if !indata.AssignToValue(&drilldata, false) {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("output value must be stringlike"))
		}
		return drilldata, nil

	case nil:
		return indata, nil
	}
	return nil, rbxmk.NewMergeError(indata, drilldata, nil)
}

func mergeValue(opt rbxmk.Options, rootdata, drilldata rbxmk.Data, indata rbxfile.Value) (outdata rbxmk.Data, err error) {
	switch drilldata := drilldata.(type) {
	case Property:
		if v := drilldata.Properties[drilldata.Name]; v != nil && indata.Type() != v.Type() {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("expected input type %s, got %s", v.Type(), indata.Type()))
		}
		drilldata.Properties[drilldata.Name] = indata
		return rootdata, nil

	case rbxfile.Value:
		if drilldata != nil && indata.Type() != drilldata.Type() {
			return nil, rbxmk.NewMergeError(indata, drilldata, fmt.Errorf("expected input type %s, got %s", drilldata.Type(), indata.Type()))
		}
		return indata, nil

	case nil:
		return indata, nil
	}
	s := NewStringlike(indata)
	if s == nil {
		return nil, rbxmk.NewMergeError(indata, drilldata, nil)
	}
	return mergeStringlike(opt, rootdata, drilldata, s)
}

func mergeRegion(opt rbxmk.Options, rootdata, drilldata rbxmk.Data, indata *Region) (outdata rbxmk.Data, err error) {
	if indata.Property == nil {
		return mergeStringlike(opt, rootdata, drilldata, indata.GetStringlike())
	}
	return mergeProperty(opt, rootdata, drilldata, Property{
		Name: indata.Property.Name,
		Properties: map[string]rbxfile.Value{
			indata.Property.Name: indata.GetStringlike().GetValue(true),
		},
	})
}

func mergeDelete(opt rbxmk.Options, rootdata, drilldata rbxmk.Data, indata rbxmk.DeleteData) (outdata rbxmk.Data, err error) {
	switch drilldata := drilldata.(type) {
	case *[]*rbxfile.Instance:
		*drilldata = (*drilldata)[:0]
		return rootdata, nil

	case *rbxfile.Instance:
		drilldata.SetParent(nil)
		return rootdata, nil

	case map[string]rbxfile.Value:
		for k := range drilldata {
			delete(drilldata, k)
		}
		return rootdata, nil

	case Property:
		delete(drilldata.Properties, drilldata.Name)
		return rootdata, nil

	case *Region:
		drilldata.Set(nil)
		if drilldata.Property == nil {
			return drilldata.Value, nil
		}
		return rootdata, nil

	case *Stringlike:
		drilldata.Type = rbxfile.TypeInvalid
		drilldata.Bytes = nil
		return drilldata, nil

	case rbxfile.Value:
		return nil, nil

	case nil:
		return nil, nil
	}
	return nil, rbxmk.NewMergeError(indata, drilldata, nil)
}
