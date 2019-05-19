package rbxmk

import (
	"errors"
	"fmt"
	"io"
	"sort"
)

// EOD indicates that a Data cannot be drilled into.
var EOD = errors.New("end of drill")

type Data interface {
	// Type returns a string representation of the Data's type.
	Type() string

	// Drill drills into the Data using inref, returning another Data that
	// represents the result. It also returns the reference after it has been
	// processed. In case of an error, the original Data and inref is
	// returned. If inref is empty, or the Data cannot be drilled into, then
	// an EOD error is returned.
	Drill(opt *Options, inref []string) (outdata Data, outref []string, err error)

	// Merge is used to merge input Data into output Data. Three kinds of Data
	// are received. rootdata is the top-level data returned by an output
	// scheme handler. drilldata is a value within the rootdata, which was
	// selected by one or more Drills. If no drills were used, then drilldata
	// will be equal to rootdata. Depending on the scheme, both rootdata and
	// drilldata may be nil. The receiver is the data to be merged. Merge
	// returns the resulting Data after it has been merged.
	Merge(opt *Options, rootdata, drilldata Data) (outdata Data, err error)
}

// DataTypeError is returned when a Data of an unexpected type is received.
type DataTypeError struct {
	dataName string
}

func (err DataTypeError) Error() string {
	return fmt.Sprintf("unexpected Data type: %s", err.dataName)
}

// NewDataTypeError returns a DataTypeError with the given Data as the
// unexpected type.
func NewDataTypeError(data Data) error {
	if data == nil {
		return DataTypeError{dataName: "nil"}
	}
	return DataTypeError{dataName: data.Type()}
}

// MergeError is returned when two Data could not be merged.
type MergeError struct {
	indata, drilldata string
	msg               error
}

func (err *MergeError) Error() string {
	if err.msg == nil {
		return fmt.Sprintf("cannot merge %s into %s", err.indata, err.drilldata)
	}
	return fmt.Sprintf("cannot merge %s into %s: %s", err.indata, err.drilldata, err.msg.Error())
}

// NewMergeError returns a MergeError with the two given Data values, and an
// optional message indicating why the merge failed.
func NewMergeError(indata, drilldata Data, msg error) error {
	err := &MergeError{"nil", "nil", msg}
	if indata != nil {
		err.indata = indata.Type()
	}
	if drilldata != nil {
		err.drilldata = drilldata.Type()
	}
	return err
}

// Format represents a rbxmk format.
type Format struct {
	Name  string
	Ext   string
	Codec InitFormatCodec
	// CanEncode    func(data Data) bool
}

// FormatDecoder is the interface that wraps the format Decode method.
type FormatDecoder interface {
	// Decode receives a Reader, decodes it, and sets the results to `data`.
	Decode(r io.Reader, data *Data) (err error)
}

// FormatEncoder is the interface that wrap the format Encode method.
type FormatEncoder interface {
	// Encode receives `data`, encodes it, and write the results to a Writer.
	Encode(w io.Writer, data Data) (err error)
}

// FormatCodec is the interface that groups the Decode and Encode methods,
// representing an object that can both decode and encode a format.
type FormatCodec interface {
	FormatDecoder
	FormatEncoder
}

// InitFormatCodec is a function that initializes a FormatCodec. The ctx
// argument is a value passed in from a scheme, which may be used by the
// format to provide context.
type InitFormatCodec func(opt *Options, ctx interface{}) (codec FormatCodec)

// Formats is a container of rbxmk formats.
type Formats struct {
	f map[string]*Format
}

// NewFormats creates and initializes a new Formats container.
func NewFormats() *Formats {
	return &Formats{f: map[string]*Format{}}
}

// Copy returns a copy of Formats. Changes to the copy will not affect the
// original.
func (fs *Formats) Copy() *Formats {
	c := Formats{
		f: make(map[string]*Format, len(fs.f)),
	}
	for k, v := range fs.f {
		c.f[k] = v
	}
	return &c
}

// Remove leading dot character.
func normalizeExt(ext string) string {
	if len(ext) > 0 && ext[0] == '.' {
		ext = ext[1:]
	}
	return ext
}

// Register registers a number of rbxmk formats with the container. An error
// is returned if a format of the same extension is already registered.
func (fs *Formats) Register(formats ...Format) error {
	for i := range formats {
		formats[i].Ext = normalizeExt(formats[i].Ext)
	}
	for i, f := range formats {
		if f.Ext == "" {
			return fmt.Errorf("format #%d must have non-empty Ext", i)
		}
		if _, registered := fs.f[f.Ext]; registered {
			return fmt.Errorf("format \"%s\" is already registered", f.Ext)
		}
		if f.Codec == nil {
			return fmt.Errorf("format \"%s\" must have Codec function", f.Ext)
		}
	}
	for _, f := range formats {
		format := f
		fs.f[format.Ext] = &format
	}
	return nil
}

// List returns a list of rbxmk formats registered with the container. The
// list is sorted by extension.
func (fs *Formats) List() []Format {
	l := make([]Format, len(fs.f))
	i := 0
	for _, f := range fs.f {
		l[i] = *f
		i++
	}
	sort.Slice(l, func(i, j int) bool {
		return l[i].Ext < l[j].Ext
	})
	return l
}

// Registered returns whether a given extension is registered.
func (fs *Formats) Registered(ext string) (registered bool) {
	_, registered = fs.f[normalizeExt(ext)]
	return registered
}

// Name returns the name of a format from a given extension. Returns an empty
// string if the extension is not registered.
func (fs *Formats) Name(ext string) (name string) {
	f, registered := fs.f[normalizeExt(ext)]
	if !registered {
		return ""
	}
	return f.Name
}

// Decoder returns a FormatDecoder from a given extension. Returns nil if the
// extension is not registered.
func (fs *Formats) Decoder(ext string, opt *Options, ctx interface{}) (dec FormatDecoder) {
	f, registered := fs.f[normalizeExt(ext)]
	if !registered {
		return nil
	}
	return f.Codec(opt, ctx)
}

// Decode directly calls a format codec's Decode method.
func (fs *Formats) Decode(ext string, opt *Options, ctx interface{}, r io.Reader, data *Data) (err error) {
	f, registered := fs.f[normalizeExt(ext)]
	if !registered {
		return nil
	}
	return f.Codec(opt, ctx).Decode(r, data)
}

// Encoder returns a FormatEncoder from a given extension. Returns nil if the
// extension is not registered.
func (fs *Formats) Encoder(ext string, opt *Options, ctx interface{}) (enc FormatEncoder) {
	f, registered := fs.f[normalizeExt(ext)]
	if !registered {
		return nil
	}
	return f.Codec(opt, ctx)
}

// Encode directly calls a format codec's Encode method.
func (fs *Formats) Encode(ext string, opt *Options, ctx interface{}, w io.Writer, data Data) (err error) {
	f, registered := fs.f[normalizeExt(ext)]
	if !registered {
		return nil
	}
	return f.Codec(opt, ctx).Encode(w, data)
}
