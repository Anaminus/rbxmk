package rtypes

import (
	"errors"
	"sort"

	"github.com/robloxapi/types"
)

// Instance contains data describing an instance of a particular class. It
// corresponds to the Instance type in Roblox.
type Instance struct {
	ClassName string
	Reference string
	IsService bool

	properties     map[string]types.PropValue
	children       []*Instance
	parent         *Instance
	desc           *Desc
	descBlocked    bool
	attrcfg        *AttrConfig
	attrcfgBlocked bool

	// Contains model metadata. Non-nil also signals that Instance is a
	// DataModel.
	metadata map[string]string
}

// NewInstance returns a new Instance of the given class name. Optionally,
// parent sets the Parent of the instance.
func NewInstance(className string, parent *Instance) *Instance {
	inst := &Instance{
		ClassName:  className,
		properties: map[string]types.PropValue{},
	}
	if parent != nil {
		inst.SetParent(parent)
	}
	return inst
}

// NewDataModel returns a special Instance of the DataModel class, to be used as
// the root of a tree of instances.
func NewDataModel() *Instance {
	return &Instance{
		ClassName:  "DataModel",
		properties: map[string]types.PropValue{},
		metadata:   map[string]string{},
	}
}

// Type returns a string identifying the type of the value.
func (*Instance) Type() string {
	return "Instance"
}

// String returns a string representation of the instance by returning the Name,
// or the ClassName if Name isn't defined.
func (inst *Instance) String() string {
	if v, ok := (inst.properties["Name"]).(types.Stringlike); ok {
		return v.Stringlike()
	}
	return inst.ClassName
}

// Copy is an alias for Clone that allows Instance to implement types.PropValue.
func (inst *Instance) Copy() types.PropValue {
	return inst.Clone()
}

// IsDataModel returns whether the instance is a root DataModel.
func (inst *Instance) IsDataModel() bool {
	return inst.metadata != nil
}

// Metadata returns the metadata of a DataModel. Returns nil if the Instance is
// not a DataModel.
func (inst *Instance) Metadata() map[string]string {
	return inst.metadata
}

// propRef holds the value of an instance property to be resolved later.
type propRef struct {
	Instance *Instance
	Property string
	Value    *Instance
}

// clone returns a deep copy of the instance while managing references.
func (inst *Instance) clone(refs map[*Instance]*Instance, propRefs *[]propRef) *Instance {
	c := *inst
	clone := &c
	clone.children = make([]*Instance, len(inst.children))
	clone.properties = make(map[string]types.PropValue, len(inst.properties))
	clone.parent = nil
	refs[inst] = clone
	for name, v := range inst.properties {
		switch v := v.(type) {
		case *Instance:
			*propRefs = append(*propRefs, propRef{
				Instance: clone,
				Property: name,
				Value:    v,
			})
			continue
		case types.PropValue:
			clone.properties[name] = v.Copy()
		}
		clone.properties[name] = v
	}
	for i, child := range inst.children {
		c := child.clone(refs, propRefs)
		clone.children[i] = c
		c.parent = clone
	}
	return clone
}

// Clone returns a copy of the instance. Each property and all descendants are
// copied as well. Unlike Roblox's implementation, the Archivable property is
// ignored.
//
// A copied reference within the tree is resolved so that it points to the
// corresponding copy of the original referent. Copied references that point
// to an instance which isn't being copied will still point to the same
// instance.
func (inst *Instance) Clone() *Instance {
	refs := make(map[*Instance]*Instance)
	propRefs := make([]propRef, 0, 8)
	clone := inst.clone(refs, &propRefs)
	for _, propRef := range propRefs {
		if c, ok := refs[propRef.Value]; ok {
			propRef.Instance.Set(propRef.Property, c)
		}
	}
	return clone
}

// IsAncestorOf returns whether the instance is the ancestor of another
// instance.
func (inst *Instance) IsAncestorOf(descendant *Instance) bool {
	if descendant != nil {
		return descendant.IsDescendantOf(inst)
	}
	return false
}

// IsDescendantOf returns whether the instance is the descendant of another
// instance.
func (inst *Instance) IsDescendantOf(ancestor *Instance) bool {
	parent := inst.parent
	for parent != nil {
		if parent == ancestor {
			return true
		}
		parent = parent.parent
	}
	return false
}

// assertLoop returns an error if an instance being the child of a parent would
// create a circular reference.
func assertLoop(child, parent *Instance) error {
	if parent == nil {
		return nil
	}
	if child.IsDataModel() {
		// Error: The Parent property of <Name> is locked, current parent: NULL, new parent <Parent>
		return errors.New("cannot set parent of DataModel")
	}
	if parent == child {
		return errors.New("attempt to set instance as its own parent")
	}
	if parent.IsDescendantOf(child) {
		return errors.New("attempt to set parent would result in circular reference")
	}
	return nil
}

// addChild appends a child to the instance, and sets its parent. If the child
// is already the child of another instance, it is first removed.
func (inst *Instance) addChild(child *Instance) {
	if child.parent != nil {
		child.parent.RemoveChild(child)
	}
	inst.children = append(inst.children, child)
	child.parent = inst
}

// AddChild appends a child instance to the instance's list of children. If
// the child has a parent, it is first removed. The parent of the child is set
// to the instance. An error is returned if the instance is a descendant of
// the child, or if the child is the instance itself.
func (inst *Instance) AddChild(child *Instance) error {
	if err := assertLoop(child, inst); err != nil {
		return err
	}
	inst.addChild(child)
	return nil
}

// AddChildAt inserts a child instance into the instance's list of children at
// a given position. If the child has a parent, it is first removed. The
// parent of the child is set to the instance. If the index is outside the
// bounds of the list, then it is constrained. An error is returned if the
// instance is a descendant of the child, or if the child is the instance
// itself.
func (inst *Instance) AddChildAt(index int, child *Instance) error {
	if err := assertLoop(child, inst); err != nil {
		return err
	}
	if index < 0 {
		index = 0
	} else if index >= len(inst.children) {
		inst.addChild(child)
		return nil
	}
	if child.parent != nil {
		child.parent.RemoveChild(child)
	}
	inst.children = append(inst.children, nil)
	copy(inst.children[index+1:], inst.children[index:])
	inst.children[index] = child
	child.parent = inst
	return nil
}

// removeChildAt removes the child at the given index, which is assumed to be
// within bounds.
func (inst *Instance) removeChildAt(i int) (child *Instance) {
	child = inst.children[i]
	child.parent = nil
	copy(inst.children[i:], inst.children[i+1:])
	inst.children[len(inst.children)-1] = nil
	inst.children = inst.children[:len(inst.children)-1]
	return child
}

// removeChildAtFast removes the child at the given index, which is assumed to
// be within bounds. The order of children is not preserved.
func (inst *Instance) removeChildAtFast(i int) (child *Instance) {
	child = inst.children[i]
	child.parent = nil
	if n := len(inst.children); n == 1 {
		inst.children[i] = nil
		inst.children = inst.children[:0]
	} else {
		inst.children[i] = inst.children[n-1]
		inst.children = inst.children[:n-1]
	}
	return child
}

// RemoveChild removes a child instance from the instance's list of children.
// The parent of the child is set to nil. Returns the removed child.
func (inst *Instance) RemoveChild(child *Instance) *Instance {
	for index, c := range inst.children {
		if c == child {
			return inst.removeChildAt(index)
		}
	}
	return nil
}

// RemoveChildFast is like RemoveChild, but does not preserve the order of
// children.
func (inst *Instance) RemoveChildFast(child *Instance) *Instance {
	for index, c := range inst.children {
		if c == child {
			return inst.removeChildAtFast(index)
		}
	}
	return nil
}

// RemoveChildAt removes the child at a given position from the instance's
// list of children. The parent of the child is set to nil. If the index is
// outside the bounds of the list, then no children are removed. Returns the
// removed child.
func (inst *Instance) RemoveChildAt(index int) *Instance {
	if index < 0 || index >= len(inst.children) {
		return nil
	}
	return inst.removeChildAt(index)
}

// RemoveChildAtFast is like RemoveChildAt, but does not preserve the order of
// children.
func (inst *Instance) RemoveChildAtFast(index int) *Instance {
	if index < 0 || index >= len(inst.children) {
		return nil
	}
	return inst.removeChildAtFast(index)
}

// RemoveAll remove every child from the instance. The parent of each child is
// set to nil.
func (inst *Instance) RemoveAll() {
	for i, child := range inst.children {
		child.parent = nil
		inst.children[i] = nil
	}
	inst.children = inst.children[:0]
}

// Children returns the children of the instance in a list.
func (inst *Instance) Children() []*Instance {
	children := make([]*Instance, len(inst.children))
	copy(children, inst.children)
	return children
}

// descendants recursively accumulates descendants of an instance in a.
func (inst *Instance) descendants(a *[]*Instance) {
	for _, child := range inst.children {
		*a = append(*a, child)
		child.descendants(a)
	}
}

// Descendants returns the descendants of the instance in a list.
func (inst *Instance) Descendants() []*Instance {
	var children []*Instance
	inst.descendants(&children)
	return children
}

// Parent returns the parent of the instance. Returns nil if the instance has no
// parent.
func (inst *Instance) Parent() *Instance {
	return inst.parent
}

// SetParent sets the parent of the instance, removing itself from the
// children of the old parent, and adding itself as a child of the new parent.
// The parent can be set to nil. An error is returned if the parent is a
// descendant of the instance, or if the parent is the instance itself. If the
// new parent is the same as the old parent, then the position of the instance
// in the parent's children is unchanged.
func (inst *Instance) SetParent(parent *Instance) error {
	if inst.parent == parent {
		return nil
	}
	if err := assertLoop(inst, parent); err != nil {
		return err
	}
	if inst.parent != nil {
		inst.parent.RemoveChild(inst)
	}
	if parent != nil {
		parent.addChild(inst)
	}
	return nil
}

// FindFirstAncestor returns the nearest ancestor of the instance that matches
// the given name, or nil if no such instance was found.
func (inst *Instance) FindFirstAncestor(name string) *Instance {
	parent := inst.parent
	for parent != nil {
		if parent.Name() == name {
			return parent
		}
		parent = parent.parent
	}
	return nil
}

// FindFirstAncestorOfClass returns the nearest ancestor of the instance that
// matches the given class name, or nil if no such instance was found.
func (inst *Instance) FindFirstAncestorOfClass(class string) *Instance {
	parent := inst.parent
	for parent != nil {
		if parent.ClassName == class {
			return parent
		}
		parent = parent.parent
	}
	return nil
}

// FindFirstChild returns the first child instance of the given name. If recurse
// is true, then descendants will also be searched top-down.
func (inst *Instance) FindFirstChild(name string, recurse bool) *Instance {
	for _, child := range inst.children {
		if child.Name() == name {
			return child
		}
		if recurse {
			if descendant := child.FindFirstChild(name, true); descendant != nil {
				return descendant
			}
		}
	}
	return nil
}

// FindFirstChildOfClass returns the first child instance of the given class
// name. If recurse is true, then descendants will also be searched top-down.
func (inst *Instance) FindFirstChildOfClass(class string, recurse bool) *Instance {
	for _, child := range inst.children {
		if child.ClassName == class {
			return child
		}
		if recurse {
			if descendant := child.FindFirstChildOfClass(class, true); descendant != nil {
				return descendant
			}
		}
	}
	return nil
}

// Descend returns a descendant of the instance by recursively searching for
// each name in succession according to FindFirstChild. Returns the instance
// itself if no arguments are given. Returns nil if a child could not be found.
func (inst *Instance) Descend(names ...string) *Instance {
	child := inst
	for _, name := range names {
		if child = child.FindFirstChild(name, false); child == nil {
			return nil
		}
	}
	return child
}

// Get returns the value of a property in the instance. The value will be nil
// if the property is not defined.
//
// If property is "ClassName", then the ClassName field is returned as a
// types.String.
//
// If property is "Parent", then the result of the Parent method is returned.
func (inst *Instance) Get(property string) (value types.PropValue) {
	switch property {
	case "ClassName":
		return types.String(inst.ClassName)
	case "Parent":
		return inst.Parent()
	}
	return inst.properties[property]
}

// Set sets the value of a property in the instance. If value is nil, then the
// value will be deleted from the Properties map.
//
// If property is "ClassName", then the ClassName field is set. Set panics if
// value does not implement types.Stringlike.
//
// If property is "Parent", then the SetParent method is called. Set panics if
// value is not an *Instance or nil, or if SetParent returns an error.
func (inst *Instance) Set(property string, value types.PropValue) {
	switch property {
	case "ClassName":
		if s, ok := value.(types.Stringlike); ok {
			inst.ClassName = s.Stringlike()
			return
		}
		panic("value of ClassName must be Stringlike")
	case "Parent":
		switch parent := value.(type) {
		case nil:
			if err := inst.SetParent(nil); err != nil {
				panic(err)
			}
		case *Instance:
			if err := inst.SetParent(parent); err != nil {
				panic(err)
			}
		}
		panic("value of Parent must be *Instance or nil")
	}
	if value == nil {
		delete(inst.properties, property)
	} else {
		inst.properties[property] = value
	}
}

// Properties returns the properties of the instance as names mapped to values.
//
// ClassName and Parent are not included.
func (inst *Instance) Properties() map[string]types.PropValue {
	props := make(map[string]types.PropValue, len(inst.properties))
	for name, value := range inst.properties {
		props[name] = value
	}
	return props
}

// PropertyNames returns a list of names of properties set on the instance.
func (inst *Instance) PropertyNames() []string {
	props := make([]string, 0, len(inst.properties))
	for name := range inst.properties {
		props = append(props, name)
	}
	sort.Strings(props)
	return props
}

// SetProperties replaces the properties of the instance with props. Each entry
// in props is a property name mapped to a value. ClassName and Parent are
// ignored.
//
// If replace is true, then properties in the instance that are nil in props are
// removed. If false, such properties are retained. If props is nil or empty,
// and replace is true, then all properties are removed from the instance.
func (inst *Instance) SetProperties(props map[string]types.PropValue, replace bool) {
	if replace {
		for name := range inst.properties {
			if props[name] == nil {
				delete(inst.properties, name)
			}
		}
	}
	for name, value := range props {
		if value == nil {
			delete(inst.properties, name)
		} else {
			inst.properties[name] = value
		}
	}
}

// Name returns the Name property of the instance, or an empty string if it is
// invalid or not defined.
func (inst *Instance) Name() string {
	if v, ok := inst.properties["Name"].(types.Stringlike); ok {
		return v.Stringlike()
	}
	return ""
}

// SetName sets the Name property of the instance.
func (inst *Instance) SetName(name string) {
	inst.properties["Name"] = types.String(name)
}

// GetFullName returns the "full" name of the instance, which is the combined
// names of the instance and every ancestor, separated by a `.` character.
func (inst *Instance) GetFullName() string {
	// Note: Roblox's GetFullName stops at the first ancestor that is a
	// ServiceProvider. Since recreating this behavior would require
	// information about the class hierarchy, this implementation simply
	// includes every ancestor.
	names := make([]string, 0, 8)
	object := inst
	for object != nil && object.metadata == nil {
		names = append(names, object.Name())
		object = object.Parent()
	}
	full := make([]byte, 0, 64)
	for i := len(names) - 1; i > 0; i-- {
		full = append(full, []byte(names[i])...)
		full = append(full, '.')
	}
	full = append(full, []byte(names[0])...)
	return string(full)
}

// Desc returns the nearest root descriptor for the instance. If the descriptor
// of current instance is nil, then the parent is searched, and so on, until a
// non-nil or blocked descriptor is found. Nil is returned if no descriptors are
// found.
func (inst *Instance) Desc() *Desc {
	parent := inst
	for parent != nil {
		if parent.descBlocked {
			return nil
		}
		if parent.desc != nil {
			return parent.desc
		}
		parent = parent.parent
	}
	return nil
}

// RawDesc returns the root descriptor for the instance, and whether it is
// blocked.
func (inst *Instance) RawDesc() (desc *Desc, blocked bool) {
	return inst.desc, inst.descBlocked
}

// SetDesc sets root as the root descriptor for the instance. If blocked is
// true, then the root descriptor is set to nil, and Desc will return nil if it
// reaches the instance.
func (inst *Instance) SetDesc(root *Desc, blocked bool) {
	if blocked {
		inst.desc = nil
		inst.descBlocked = true
		return
	}
	inst.desc = root
	inst.descBlocked = false
}

// IsA returns whether the instance's class inherits from className. If the
// instance has no descriptor, then only its ClassName is compared.
func (inst *Instance) IsA(className string) bool {
	desc := inst.Desc()
	if desc == nil {
		return inst.ClassName == className
	}
	class := desc.Classes[inst.ClassName]
	for class != nil {
		if class.Name == className {
			return true
		}
		class = desc.Classes[class.Superclass]
	}
	return false
}

// FindFirstChildWhichIsA returns the first child instance that inherits from
// the given class name. If recurse is true, then descendants will also be
// searched top-down. If an instance has no descriptor, then only its ClassName
// is compared.
func (inst *Instance) FindFirstChildWhichIsA(class string, recurse bool) *Instance {
	for _, child := range inst.children {
		// TODO: improve performance by handling descriptors directly.
		if child.IsA(class) {
			return child
		}
		if recurse {
			if descendant := child.FindFirstChildWhichIsA(class, true); descendant != nil {
				return descendant
			}
		}
	}
	return nil
}

// FindFirstAncestorWhichIsA returns the nearest ancestor of the instance that
// inherits from the given class name, or nil if no such instance was found. If
// an instance has no descriptor, then only its ClassName is compared.
func (inst *Instance) FindFirstAncestorWhichIsA(class string) *Instance {
	parent := inst.parent
	for parent != nil {
		if parent.IsA(class) {
			return parent
		}
		parent = parent.parent
	}
	return nil
}

// WithDescIsA is like IsA, but includes a global descriptor.
func (inst *Instance) WithDescIsA(globalDesc *Desc, className string) bool {
	desc := globalDesc.Of(inst)
	if desc == nil {
		return inst.ClassName == className
	}
	class := desc.Classes[inst.ClassName]
	for class != nil {
		if class.Name == className {
			return true
		}
		class = desc.Classes[class.Superclass]
	}
	return false
}

// WithDescFindFirstChildWhichIsA is like FindFirstChildWhichIsA, but includes a
// global descriptor.
func (inst *Instance) WithDescFindFirstChildWhichIsA(globalDesc *Desc, class string, recurse bool) *Instance {
	for _, child := range inst.children {
		// TODO: improve performance by handling descriptors directly.
		if child.WithDescIsA(globalDesc, class) {
			return child
		}
		if recurse {
			if descendant := child.WithDescFindFirstChildWhichIsA(globalDesc, class, true); descendant != nil {
				return descendant
			}
		}
	}
	return nil
}

// WithDescFindFirstAncestorWhichIsA is like FindFirstAncestorWhichIsA, but
// includes a global descriptor.
func (inst *Instance) WithDescFindFirstAncestorWhichIsA(globalDesc *Desc, class string) *Instance {
	parent := inst.parent
	for parent != nil {
		if parent.WithDescIsA(globalDesc, class) {
			return parent
		}
		parent = parent.parent
	}
	return nil
}

// AttrConfig returns the nearest AttrConfig for the instance. If the AttrConfig
// of current instance is nil, then the parent is searched, and so on, until a
// non-nil or blocked AttrConfig is found. Nil is returned if no Attrs are
// found.
func (inst *Instance) AttrConfig() *AttrConfig {
	parent := inst
	for parent != nil {
		if parent.attrcfgBlocked {
			return nil
		}
		if parent.attrcfg != nil {
			return parent.attrcfg
		}
		parent = parent.parent
	}
	return nil
}

// RawAttrConfig returns the AttrConfig for the instance, and whether it is
// blocked.
func (inst *Instance) RawAttrConfig() (attrcfg *AttrConfig, blocked bool) {
	return inst.attrcfg, inst.attrcfgBlocked
}

// SetAttrConfig sets attrcfg as the AttrConfig for the instance. If blocked is
// true, then the AttrConfig is set to nil, and AttrConfig will return nil if it
// reaches the instance.
func (inst *Instance) SetAttrConfig(attrcfg *AttrConfig, blocked bool) {
	if blocked {
		inst.attrcfg = nil
		inst.attrcfgBlocked = true
		return
	}
	inst.attrcfg = attrcfg
	inst.attrcfgBlocked = false
}

// ForEachChild iterates over each child of the instance.
//
// If cb returns an error, iteration stops, and the error is returned.
func (inst *Instance) ForEachChild(cb func(child *Instance) error) error {
	for _, child := range inst.children {
		if err := cb(child); err != nil {
			return err
		}

	}
	return nil
}

// SkipChildren causes Instance.ForEachDescendant to skip iterating an
// instance's children.
var SkipChildren = errors.New("skip children")

// ForEachDescendant iterates over each descendant of the instance, depth-first.
//
// If cb returns SkipChildren, then the method moves to the next sibling of desc
// without visiting the children of desc.
//
// Otherwise, if cb returns an error, iteration stops, and the error is
// returned.
func (inst *Instance) ForEachDescendant(cb func(desc *Instance) error) error {
	for _, child := range inst.children {
		if err := cb(child); err != nil {
			if err == SkipChildren {
				continue
			}
			return err
		}
		if err := child.ForEachDescendant(cb); err != nil {
			return err
		}
	}
	return nil
}

// ForEachProperty iterates over each property of the instance.
//
// The order in which properties are visited is undefined.
//
// If cb returns an error, iteration stops, and the error is returned.
func (inst *Instance) ForEachProperty(cb func(name string, value types.PropValue) error) error {
	for name, value := range inst.properties {
		if err := cb(name, value); err != nil {
			return err
		}
	}
	return nil
}
