package rbxmk

import "github.com/anaminus/rbxmk/rtypes"

// Global contains values that available across an entire World.
type Global struct {
	Desc       *rtypes.RootDesc
	AttrConfig *rtypes.AttrConfig
}

// DescOf returns the root descriptor of an instance. If inst is nil, the global
// descriptor is returned.
func (g Global) DescOf(inst *rtypes.Instance) *rtypes.RootDesc {
	if inst != nil {
		if desc := inst.Desc(); desc != nil {
			return desc
		}
	}
	return g.Desc
}

// AttrConfigOf returns the AttrConfigOf of an instance. If inst is nil, the
// global AttrConfig is returned.
func (g Global) AttrConfigOf(inst *rtypes.Instance) *rtypes.AttrConfig {
	if inst != nil {
		if attrcfg := inst.AttrConfig(); attrcfg != nil {
			return attrcfg
		}
	}
	return g.AttrConfig
}
