// The dump package describes Lua APIs.
package dump

import (
	"github.com/anaminus/rbxmk/dump/dt"
)

// Value is a value that has a Type.
type Value interface {
	// Kind returns a name describing the kind of type.
	Kind() string
	// Type returns a type definition.
	Type() dt.Type

	// Index returns the element of the value indexed by name. Returns nil if
	// the value does not have elements, or no such element exists. Appends to
	// path the path elements required to reach the indexed element.
	Index(path []string, name string) ([]string, Value)

	// Indices returns the indices of the elements of the value. Returns an
	// empty slice if the value has zero elements, and nil if the value cannot
	// have elements.
	Indices() []string

	v()
}

// Node is a node within a dump tree, with a Root as the top-level node.
type Node interface {
	// Resolve recursively resolves path, returning the referred value. Returns
	// nil if no subnode could be found. Returns the node if path is empty.
	Resolve(path ...string) any
}

// resolveValue ensures that a path does not return a value if the path is
// longer than it should be.
func resolveValue(path []string, v any) any {
	if len(path) > 0 {
		return nil
	}
	return v
}
