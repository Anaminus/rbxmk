package rbxmk

import (
	"github.com/robloxapi/rbxfile"
)

type Source struct {
	FileName   string
	Instances  []*rbxfile.Instance
	Properties map[string]rbxfile.Value
	Values     []rbxfile.Value
	Sources    []*Source
}

func (src *Source) Copy() *Source {
	dst := &Source{
		FileName:   src.FileName,
		Instances:  make([]*rbxfile.Instance, len(src.Instances)),
		Properties: make(map[string]rbxfile.Value, len(src.Properties)),
		Values:     make([]rbxfile.Value, len(src.Values)),
		Sources:    make([]*Source, len(src.Sources)),
	}
	for i, inst := range src.Instances {
		dst.Instances[i] = inst.Clone()
	}
	for name, value := range src.Properties {
		dst.Properties[name] = value.Copy()
	}
	for i, value := range src.Values {
		dst.Values[i] = value.Copy()
	}
	for i, s := range src.Sources {
		dst.Sources[i] = s.Copy()
	}
	return dst
}
