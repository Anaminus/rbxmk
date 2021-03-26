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
	Name:        "fragments",
	Description: `List of document fragment paths.`,
	Func: func(w io.Writer, root dump.Root) error {
		buf := bufio.NewWriter(w)
		p := "libraries"
		for _, library := range root.Libraries {
			p := path.Join(p, library.Name)
			fragWrite(buf, path.Join(p, "Summary"), library.Struct.Summary)
			fragWrite(buf, path.Join(p, "Description"), library.Struct.Description)
			fragWriteFields(buf, path.Join(p, "Fields"), library.Struct.Fields)
			fragWriteTypes(buf, path.Join(p, "Types"), library.Types)
		}
		fragWriteTypes(buf, p, root.Types)
		for _, frag := range root.Fragments {
			fragWrite(buf, "fragments", frag)
		}
		return buf.Flush()
	},
}

func fragWrite(buf *bufio.Writer, p string, path string) {
	switch path {
	case "":
		buf.WriteString(">>>> ")
		buf.WriteString(p)
	case "$TODO":
		buf.WriteString("---- ")
		buf.WriteString(p)
	default:
		buf.WriteString(path)
	}
	buf.WriteString("\n")
}

func fragWriteFields(buf *bufio.Writer, p string, fields dump.Fields) {
	sortFields(fields, func(fieldName string, field dump.Value) {
		p := path.Join(p, fieldName)
		switch value := field.(type) {
		case dump.Property:
			fragWrite(buf, path.Join(p, "Summary"), value.Summary)
			fragWrite(buf, path.Join(p, "Description"), value.Description)
		case dump.Struct:
			fragWrite(buf, path.Join(p, "Summary"), value.Summary)
			fragWrite(buf, path.Join(p, "Description"), value.Description)
			fragWriteFields(buf, p, value.Fields)
		case dump.Function:
			fragWrite(buf, path.Join(p, "Summary"), value.Summary)
			fragWrite(buf, path.Join(p, "Description"), value.Description)
		case dump.MultiFunction:
			for i, value := range value {
				fragWrite(buf, path.Join(p, strconv.Itoa(i), "Summary"), value.Summary)
				fragWrite(buf, path.Join(p, strconv.Itoa(i), "Description"), value.Description)
			}
		}
	})
}

func fragWriteTypes(buf *bufio.Writer, p string, types dump.TypeDefs) {
	sortTypeDefs(types, func(defName string, def dump.TypeDef) {
		p := path.Join(p, defName)
		fragWrite(buf, path.Join(p, "Summary"), def.Summary)
		fragWrite(buf, path.Join(p, "Description"), def.Description)
		sortConstructors(def.Constructors, func(ctorName string, ctor dump.MultiFunction) {
			p := path.Join(p, "Constructors", ctorName)
			for i, fn := range ctor {
				p := path.Join(p, strconv.Itoa(i))
				fragWrite(buf, path.Join(p, "Summary"), fn.Summary)
				fragWrite(buf, path.Join(p, "Description"), fn.Description)
			}
		})
		sortProperties(def.Properties, func(propName string, prop dump.Property) {
			p := path.Join(p, "Properties", propName)
			fragWrite(buf, path.Join(p, "Summary"), prop.Summary)
			fragWrite(buf, path.Join(p, "Description"), prop.Description)
		})
		sortProperties(def.Symbols, func(symName string, prop dump.Property) {
			p := path.Join(p, "Symbols", symName)
			fragWrite(buf, path.Join(p, "Summary"), prop.Summary)
			fragWrite(buf, path.Join(p, "Description"), prop.Description)
		})
		sortMethods(def.Methods, func(methodName string, method dump.Function) {
			p := path.Join(p, "Methods", methodName)
			fragWrite(buf, path.Join(p, "Summary"), method.Summary)
			fragWrite(buf, path.Join(p, "Description"), method.Description)
		})
		if op := def.Operators; op != nil {
			p := path.Join(p, "Operators")
			for i, fn := range op.Add {
				fragWrite(buf, path.Join(p, "Add", strconv.Itoa(i), "Summary"), fn.Summary)
				fragWrite(buf, path.Join(p, "Add", strconv.Itoa(i), "Description"), fn.Description)
			}
			for i, fn := range op.Sub {
				fragWrite(buf, path.Join(p, "Sub", strconv.Itoa(i), "Summary"), fn.Summary)
				fragWrite(buf, path.Join(p, "Sub", strconv.Itoa(i), "Description"), fn.Description)
			}
			for i, fn := range op.Mul {
				fragWrite(buf, path.Join(p, "Mul", strconv.Itoa(i), "Summary"), fn.Summary)
				fragWrite(buf, path.Join(p, "Mul", strconv.Itoa(i), "Description"), fn.Description)
			}
			for i, fn := range op.Div {
				fragWrite(buf, path.Join(p, "Div", strconv.Itoa(i), "Summary"), fn.Summary)
				fragWrite(buf, path.Join(p, "Div", strconv.Itoa(i), "Description"), fn.Description)
			}
			for i, fn := range op.Mod {
				fragWrite(buf, path.Join(p, "Mod", strconv.Itoa(i), "Summary"), fn.Summary)
				fragWrite(buf, path.Join(p, "Mod", strconv.Itoa(i), "Description"), fn.Description)
			}
			for i, fn := range op.Pow {
				fragWrite(buf, path.Join(p, "Pow", strconv.Itoa(i), "Summary"), fn.Summary)
				fragWrite(buf, path.Join(p, "Pow", strconv.Itoa(i), "Description"), fn.Description)
			}
			for i, fn := range op.Concat {
				fragWrite(buf, path.Join(p, "Concat", strconv.Itoa(i), "Summary"), fn.Summary)
				fragWrite(buf, path.Join(p, "Concat", strconv.Itoa(i), "Description"), fn.Description)
			}
			if op.Eq != nil {
				fragWrite(buf, path.Join(p, "Eq", "Summary"), op.Eq.Summary)
				fragWrite(buf, path.Join(p, "Eq", "Description"), op.Eq.Description)
			}
			if op.Lt != nil {
				fragWrite(buf, path.Join(p, "Lt", "Summary"), op.Lt.Summary)
				fragWrite(buf, path.Join(p, "Lt", "Description"), op.Lt.Description)
			}
			if op.Le != nil {
				fragWrite(buf, path.Join(p, "Le", "Summary"), op.Le.Summary)
				fragWrite(buf, path.Join(p, "Le", "Description"), op.Le.Description)
			}
			if op.Len != nil {
				fragWrite(buf, path.Join(p, "Len", "Summary"), op.Len.Summary)
				fragWrite(buf, path.Join(p, "Len", "Description"), op.Len.Description)
			}
			if op.Unm != nil {
				fragWrite(buf, path.Join(p, "Unm", "Summary"), op.Unm.Summary)
				fragWrite(buf, path.Join(p, "Unm", "Description"), op.Unm.Description)
			}
			if op.Call != nil {
				fragWrite(buf, path.Join(p, "Call", "Summary"), op.Call.Summary)
				fragWrite(buf, path.Join(p, "Call", "Description"), op.Call.Description)
			}
			if op.Index != nil {
				fragWrite(buf, path.Join(p, "Index", "Summary"), op.Index.Summary)
				fragWrite(buf, path.Join(p, "Index", "Description"), op.Index.Description)
			}
			if op.Newindex != nil {
				fragWrite(buf, path.Join(p, "Newindex", "Summary"), op.Newindex.Summary)
				fragWrite(buf, path.Join(p, "Newindex", "Description"), op.Newindex.Description)
			}
		}
	})
}
