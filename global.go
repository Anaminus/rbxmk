package rbxmk

import "github.com/anaminus/rbxmk/rtypes"

// Global contains values that available across an entire World.
type Global struct {
	Desc       *rtypes.RootDesc
	AttrConfig *rtypes.AttrConfig
}
