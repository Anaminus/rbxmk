package dumpformats

import (
	"bufio"
	"io"
	"path"
	"strconv"

	"github.com/anaminus/rbxmk/dump"
)

func init() { register(Fragments) }

var Fragments = Format{
	Name: "fragments",
	Func: func(w io.Writer, root dump.Root, opts Options) error {
		buf := bufio.NewWriter(w)
		ListFragments(root, func(path, ref string) {
			switch ref {
			case "":
				buf.WriteString(">>>> ")
				buf.WriteString(path)
			case "$TODO":
				buf.WriteString("---- ")
				buf.WriteString(path)
			default:
				buf.WriteString(ref)
			}
			buf.WriteString("\n")
		})
		return buf.Flush()
	},
}

// ListFragments traverses root, calling write for each fragment reference that
// was found. ref is the fragment reference, and path describes the location of
// the reference within root.
func ListFragments(root dump.Root, write func(path, ref string)) {
	write("libraries", "Libraries")
	root.Libraries.ForEach(func(name string, library dump.Library) bool {
		p := path.Join("libraries", name)
		write(path.Join(p, "Summary"), library.Struct.Summary)
		write(path.Join(p, "Description"), library.Struct.Description)
		fragWriteFields(write, path.Join(p, "Fields"), library.Struct.Fields)
		fragWriteTypes(write, path.Join(p, "Types"), library.Types)
		return true
	})
	write("types", "Types")
	fragWriteTypes(write, "types", root.Types)
	write("formats", "Formats")
	fragWriteFormats(write, "formats", root.Formats)
}

func fragWriteFormats(write func(path, ref string), p string, formats dump.Formats) {
	sortFormats(formats, func(formatName string, format dump.Format) {
		p := path.Join(p, formatName)
		write(path.Join(p, "Summary"), format.Summary)
		write(path.Join(p, "Description"), format.Description)
	})
}

func fragWriteFields(write func(path, ref string), p string, fields dump.Fields) {
	sortFields(fields, func(fieldName string, field dump.Value) {
		p := path.Join(p, fieldName)
		switch value := field.(type) {
		case dump.Property:
			write(path.Join(p, "Summary"), value.Summary)
			write(path.Join(p, "Description"), value.Description)
		case dump.Struct:
			write(path.Join(p, "Summary"), value.Summary)
			write(path.Join(p, "Description"), value.Description)
			fragWriteFields(write, p, value.Fields)
		case dump.Function:
			write(path.Join(p, "Summary"), value.Summary)
			write(path.Join(p, "Description"), value.Description)
		case dump.MultiFunction:
			for i, value := range value {
				write(path.Join(p, strconv.Itoa(i), "Summary"), value.Summary)
				write(path.Join(p, strconv.Itoa(i), "Description"), value.Description)
			}
		}
	})
}

func fragWriteTypes(write func(path, ref string), p string, types dump.TypeDefs) {
	sortTypeDefs(types, func(defName string, def dump.TypeDef) {
		p := path.Join(p, defName)
		write(path.Join(p, "Summary"), def.Summary)
		write(path.Join(p, "Description"), def.Description)
		sortConstructors(def.Constructors, func(ctorName string, ctor dump.MultiFunction) {
			p := path.Join(p, "Constructors", ctorName)
			for i, fn := range ctor {
				p := path.Join(p, strconv.Itoa(i))
				write(path.Join(p, "Summary"), fn.Summary)
				write(path.Join(p, "Description"), fn.Description)
			}
		})
		sortProperties(def.Properties, func(propName string, prop dump.Property) {
			p := path.Join(p, "Properties", propName)
			write(path.Join(p, "Summary"), prop.Summary)
			write(path.Join(p, "Description"), prop.Description)
		})
		sortProperties(dump.Properties(def.Symbols), func(symName string, prop dump.Property) {
			p := path.Join(p, "Symbols", symName)
			write(path.Join(p, "Summary"), prop.Summary)
			write(path.Join(p, "Description"), prop.Description)
		})
		sortMethods(def.Methods, func(methodName string, method dump.Function) {
			p := path.Join(p, "Methods", methodName)
			write(path.Join(p, "Summary"), method.Summary)
			write(path.Join(p, "Description"), method.Description)
		})
		if op := def.Operators; op != nil {
			p := path.Join(p, "Operators")
			for i, fn := range op.Add {
				write(path.Join(p, "Add", strconv.Itoa(i), "Summary"), fn.Summary)
				write(path.Join(p, "Add", strconv.Itoa(i), "Description"), fn.Description)
			}
			for i, fn := range op.Sub {
				write(path.Join(p, "Sub", strconv.Itoa(i), "Summary"), fn.Summary)
				write(path.Join(p, "Sub", strconv.Itoa(i), "Description"), fn.Description)
			}
			for i, fn := range op.Mul {
				write(path.Join(p, "Mul", strconv.Itoa(i), "Summary"), fn.Summary)
				write(path.Join(p, "Mul", strconv.Itoa(i), "Description"), fn.Description)
			}
			for i, fn := range op.Div {
				write(path.Join(p, "Div", strconv.Itoa(i), "Summary"), fn.Summary)
				write(path.Join(p, "Div", strconv.Itoa(i), "Description"), fn.Description)
			}
			for i, fn := range op.Mod {
				write(path.Join(p, "Mod", strconv.Itoa(i), "Summary"), fn.Summary)
				write(path.Join(p, "Mod", strconv.Itoa(i), "Description"), fn.Description)
			}
			for i, fn := range op.Pow {
				write(path.Join(p, "Pow", strconv.Itoa(i), "Summary"), fn.Summary)
				write(path.Join(p, "Pow", strconv.Itoa(i), "Description"), fn.Description)
			}
			for i, fn := range op.Concat {
				write(path.Join(p, "Concat", strconv.Itoa(i), "Summary"), fn.Summary)
				write(path.Join(p, "Concat", strconv.Itoa(i), "Description"), fn.Description)
			}
			if op.Eq != nil {
				write(path.Join(p, "Eq", "Summary"), op.Eq.Summary)
				write(path.Join(p, "Eq", "Description"), op.Eq.Description)
			}
			if op.Lt != nil {
				write(path.Join(p, "Lt", "Summary"), op.Lt.Summary)
				write(path.Join(p, "Lt", "Description"), op.Lt.Description)
			}
			if op.Le != nil {
				write(path.Join(p, "Le", "Summary"), op.Le.Summary)
				write(path.Join(p, "Le", "Description"), op.Le.Description)
			}
			if op.Len != nil {
				write(path.Join(p, "Len", "Summary"), op.Len.Summary)
				write(path.Join(p, "Len", "Description"), op.Len.Description)
			}
			if op.Unm != nil {
				write(path.Join(p, "Unm", "Summary"), op.Unm.Summary)
				write(path.Join(p, "Unm", "Description"), op.Unm.Description)
			}
			if op.Call != nil {
				write(path.Join(p, "Call", "Summary"), op.Call.Summary)
				write(path.Join(p, "Call", "Description"), op.Call.Description)
			}
			if op.Index != nil {
				write(path.Join(p, "Index", "Summary"), op.Index.Summary)
				write(path.Join(p, "Index", "Description"), op.Index.Description)
			}
			if op.Newindex != nil {
				write(path.Join(p, "Newindex", "Summary"), op.Newindex.Summary)
				write(path.Join(p, "Newindex", "Description"), op.Newindex.Description)
			}
		}
	})
}
