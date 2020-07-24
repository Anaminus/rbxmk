package rbxmk

import "github.com/robloxapi/types"

// Stringlike is any Value that is string-like. Note that this is distinct from
// a string-representation of the value. Rather, Stringlike indicates that the
// value has string-like properties.
type Stringlike = types.Stringlike

// Numberlike is any value that can be converted to a floating-point number.
type Numberlike = types.Numberlike

// Intlike is any value that can be converted to an integer.
type Intlike = types.Intlike
