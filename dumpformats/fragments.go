package dumpformats

import (
	"bufio"
	"io"
	"path"

	"github.com/anaminus/rbxmk/dump"
)

func init() { register(Fragments) }

var Fragments = Format{
	Name:        "fragments",
	Description: `List of document fragment paths.`,
	Func: func(w io.Writer, root dump.Root) error {
		suffix := ""
		buf := bufio.NewWriter(w)
		for _, library := range root.Libraries {
			prefix := path.Join("libraries", library.Name)
			fragWrite(buf, prefix+suffix)
			fragWriteFields(buf, prefix, suffix, library.Struct.Fields)
			fragWriteTypes(buf, prefix, suffix, library.Types)
		}
		fragWriteTypes(buf, "", suffix, root.Types)
		return buf.Flush()
	},
}

func fragWrite(buf *bufio.Writer, element ...string) {
	buf.WriteString(path.Join(element...))
	buf.WriteString("\n")
}

func fragWriteFields(buf *bufio.Writer, prefix, suffix string, fields dump.Fields) {
	sortFields(fields, func(fieldName string, field dump.Value) {
		fragWrite(buf, prefix, "fields", fieldName+suffix)
		switch value := field.(type) {
		case dump.Struct:
			fragWriteFields(buf, path.Join(prefix, "fields", fieldName), suffix, value.Fields)
		}
	})
}

func fragWriteTypes(buf *bufio.Writer, prefix, suffix string, types dump.TypeDefs) {
	sortTypeDefs(types, func(defName string, def dump.TypeDef) {
		prefix := path.Join(prefix, "types", defName)
		fragWrite(buf, prefix+suffix)
		sortConstructors(def.Constructors, func(ctorName string, ctor dump.MultiFunction) {
			fragWrite(buf, prefix, "constructors", ctorName+suffix)
		})
		sortProperties(def.Properties, func(propName string, prop dump.Property) {
			fragWrite(buf, prefix, "properties", propName+suffix)
		})
		sortMethods(def.Methods, func(methodName string, method dump.Function) {
			fragWrite(buf, prefix, "methods", methodName+suffix)
		})
		if op := def.Operators; op != nil {
			if len(op.Add) > 0 {
				fragWrite(buf, prefix, "operators", "add"+suffix)
			}
			if len(op.Sub) > 0 {
				fragWrite(buf, prefix, "operators", "sub"+suffix)
			}
			if len(op.Mul) > 0 {
				fragWrite(buf, prefix, "operators", "mul"+suffix)
			}
			if len(op.Div) > 0 {
				fragWrite(buf, prefix, "operators", "div"+suffix)
			}
			if len(op.Mod) > 0 {
				fragWrite(buf, prefix, "operators", "div"+suffix)
			}
			if len(op.Pow) > 0 {
				fragWrite(buf, prefix, "operators", "pow"+suffix)
			}
			if len(op.Concat) > 0 {
				fragWrite(buf, prefix, "operators", "concat"+suffix)
			}
			if op.Eq {
				fragWrite(buf, prefix, "operators", "eq"+suffix)
			}
			if op.Le {
				fragWrite(buf, prefix, "operators", "le"+suffix)
			}
			if op.Lt {
				fragWrite(buf, prefix, "operators", "lt"+suffix)
			}
			if op.Len != nil {
				fragWrite(buf, prefix, "operators", "len"+suffix)
			}
			if op.Unm != nil {
				fragWrite(buf, prefix, "operators", "unm"+suffix)
			}
			if op.Call != nil {
				fragWrite(buf, prefix, "operators", "call"+suffix)
			}
			if op.Index != nil {
				fragWrite(buf, prefix, "operators", "index"+suffix)
			}
			if op.Newindex != nil {
				fragWrite(buf, prefix, "operators", "newindex"+suffix)
			}
		}
	})
}
