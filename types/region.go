package types

import (
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxfile"
)

type RegionRange struct {
	RegA, RegB int
	SelA, SelB int
}

type Region struct {
	Property *Property
	Value    *Stringlike
	Range    []RegionRange
	Append   bool
}

func (r *Region) Get() []byte {
	if r.Value == nil {
		return nil
	}
	if len(r.Range) == 0 {
		return r.Value.Bytes
	}
	return r.Value.Bytes[r.Range[0].SelA:r.Range[0].SelB]
}

func (r *Region) GetStringlike() *Stringlike {
	s := &Stringlike{}
	if r.Value == nil {
		return s
	}
	s.ValueType = r.Value.ValueType
	s.Bytes = r.Get()
	return s
}

func (r *Region) Set(p []byte) {
	if r.Value == nil {
		r.Value = &Stringlike{Bytes: p}
	} else {
		var dst []byte
		{
			n := len(r.Value.Bytes) + len(p)*len(r.Range)
			if !r.Append {
				for _, rng := range r.Range {
					n -= rng.RegB - rng.RegA
				}
			}
			dst = make([]byte, n)
		}
		src := r.Value.Bytes
		dstoff := 0
		srcoff := 0
		if r.Append {
			for _, rng := range r.Range {
				dstoff += copy(dst[dstoff:], src[srcoff:rng.SelB])
				dstoff += copy(dst[dstoff:], p)
				dstoff += copy(dst[dstoff:], src[rng.SelB:rng.RegB])
				srcoff = rng.RegB
			}
		} else {
			for _, rng := range r.Range {
				dstoff += copy(dst[dstoff:], src[srcoff:rng.RegA])
				dstoff += copy(dst[dstoff:], p)
				srcoff = rng.RegB
			}
		}
		copy(dst[dstoff:], src[srcoff:])
		r.Value.Bytes = dst
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

func (indata *Region) Drill(opt *rbxmk.Options, inref []string) (outdata rbxmk.Data, outref []string, err error) {
	return indata, inref, rbxmk.EOD
}

func (indata *Region) Merge(opt *rbxmk.Options, rootdata, drilldata rbxmk.Data) (outdata rbxmk.Data, err error) {
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
