package rtypes

import (
	"github.com/robloxapi/rbxfile"
	"github.com/robloxapi/types"
)

const T_UniqueId = "UniqueId"

// UniqueId represents a unique identifier.
type UniqueId struct {
	Random types.Int64
	Time   uint32
	Index  uint32
}

// Type returns a string identifying the type of the value.
func (UniqueId) Type() string {
	return T_UniqueId
}

// String returns a string representation of the value.
func (u UniqueId) String() string {
	return rbxfile.ValueUniqueId(rbxfile.ValueUniqueId{
		Random: int64(u.Random),
		Time:   u.Time,
		Index:  u.Index,
	}).String()
}

// Copy returns a copy of the value.
func (u UniqueId) Copy() types.PropValue {
	return u
}
