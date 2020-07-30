package rtypes

import (
	"errors"

	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/types"
)

type Instance struct {
	ClassName string
	Reference string
	IsService bool

	properties map[string]types.PropValue
	children   []*Instance
	parent     *Instance
	desc       *rbxdump.Root
	root       bool
}

func NewInstance(className string) *Instance {
	return &Instance{
		ClassName:  className,
		properties: make(map[string]types.PropValue, 0),
	}
}

func NewDataModel() *Instance {
	return &Instance{
		root:       true,
		ClassName:  "DataModel",
		properties: make(map[string]types.PropValue, 0),
	}
}

// IsDataModel returns whether the instance is a root DataModel.
func (inst *Instance) IsDataModel() bool {
	return inst.root
}

type propRef struct {
	Instance *Instance
	Property string
	Value    *Instance
}

// clone returns a deep copy of the instance while managing references.
func (inst *Instance) clone(refs map[*Instance]*Instance, propRefs *[]propRef) *Instance {
	clone := &Instance{
		ClassName:  inst.ClassName,
		Reference:  inst.Reference,
		IsService:  inst.IsService,
		desc:       inst.desc,
		children:   make([]*Instance, len(inst.children)),
		properties: make(map[string]types.PropValue, len(inst.properties)),
	}
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

func (inst *Instance) Children() []*Instance {
	children := make([]*Instance, len(inst.children))
	copy(children, inst.children)
	return children
}

func (inst *Instance) descendants(a *[]*Instance) {
	for _, child := range inst.children {
		*a = append(*a, child)
		child.descendants(a)
	}
}

func (inst *Instance) Descendants() []*Instance {
	var children []*Instance
	inst.descendants(&children)
	return children
}

// Parent returns the parent of the instance. Can return nil if the instance
// has no parent.
func (inst *Instance) Parent() *Instance {
	if inst.root {
		return nil
	}
	return inst.parent
}

// SetParent sets the parent of the instance, removing itself from the
// children of the old parent, and adding itself as a child of the new parent.
// The parent can be set to nil. An error is returned if the parent is a
// descendant of the instance, or if the parent is the instance itself. If the
// new parent is the same as the old parent, then the position of the instance
// in the parent's children is unchanged.
func (inst *Instance) SetParent(parent *Instance) error {
	if inst.root {
		// Error: The Parent property of <Name> is locked, current parent: NULL, new parent <Parent>
		return errors.New("cannot set parent of DataModel")
	}
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

// Get returns the value of a property in the instance. The value will be nil
// if the property is not defined.
func (inst *Instance) Get(property string) (value types.Value) {
	return inst.properties[property]
}

// Set sets the value of a property in the instance. If value is nil, then the
// value will be deleted from the Properties map.
func (inst *Instance) Set(property string, value types.PropValue) {
	if value == nil {
		delete(inst.properties, property)
	} else {
		inst.properties[property] = value
	}
}

func (inst *Instance) Properties() map[string]types.PropValue {
	props := make(map[string]types.PropValue, len(inst.properties))
	for name, value := range inst.properties {
		props[name] = value
	}
	return props
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
	for object != nil && !object.root {
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

// Desc returns the nearest root descriptor for the instance. If the current
// instance does not have a descriptor, then each ancestor is searched. Nil is
// returned if no descriptors are found.
func (inst *Instance) Desc() *rbxdump.Root {
	if inst.desc != nil {
		return inst.desc
	}
	parent := inst.parent
	for parent != nil {
		if parent.desc != nil {
			return parent.desc
		}
		parent = parent.parent
	}
	return nil
}

// SetDesc sets root as the root descriptor for the instance.
func (inst *Instance) SetDesc(root *rbxdump.Root) {
	inst.desc = root
}

// Type returns a string identifying the type.
func (*Instance) Type() string {
	return "Instance"
}

// String implements the fmt.Stringer interface by returning the Name of the
// instance, or the ClassName if Name isn't defined.
func (inst *Instance) String() string {
	if v, ok := (inst.properties["Name"]).(types.Stringlike); ok {
		return v.Stringlike()
	}
	return inst.ClassName
}

// Copy implements types.PropValue.
func (inst *Instance) Copy() types.PropValue {
	return inst.Clone()
}
