package types

import (
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
)

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
	s.ValueType = r.Value.ValueType
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
		value := Value{r.Property.Properties[r.Property.Name]}
		r.Value.AssignToValue(&value, true)
		r.Property.Properties[r.Property.Name] = value.Value
	}
}

func (indata *Region) Type() string {
	return "Region"
}

func (indata *Region) Drill(opt rbxmk.Options, inref []string) (outdata rbxmk.Data, outref []string, err error) {
	return indata, inref, rbxmk.EOD
}

func (indata *Region) Merge(opt rbxmk.Options, rootdata, drilldata rbxmk.Data) (outdata rbxmk.Data, err error) {
	if indata.Property == nil {
		return indata.GetStringlike().Merge(opt, rootdata, drilldata)
	}
	return Property{
		Name: indata.Property.Name,
		Properties: map[string]rbxfile.Value{
			indata.Property.Name: indata.GetStringlike().GetValue(true).Value,
		},
	}.Merge(opt, rootdata, drilldata)
}
