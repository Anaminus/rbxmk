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
		suffix := ".md"
		buf := bufio.NewWriter(w)
		fragWriteDir(buf, "libraries")
		for _, library := range root.Libraries {
			prefix := path.Join("libraries", library.Name)
			fragWriteFile(buf, prefix+suffix)
			if len(library.Struct.Fields) > 0 || len(library.Types) > 0 {
				fragWriteDir(buf, prefix)
			}
			fragWriteFields(buf, prefix, suffix, library.Struct.Fields)
			fragWriteTypes(buf, prefix, suffix, library.Types)
		}
		fragWriteTypes(buf, "", suffix, root.Types)
		return buf.Flush()
	},
}

func fragWriteFile(buf *bufio.Writer, element ...string) {
	buf.WriteString(path.Join(element...))
	buf.WriteString("\n")
}

func fragWriteDir(buf *bufio.Writer, element ...string) {
	buf.WriteString(path.Join(element...))
	buf.WriteString("/\n")
}

func fragWriteFields(buf *bufio.Writer, prefix, suffix string, fields dump.Fields) {
	if len(fields) == 0 {
		return
	}
	fragWriteDir(buf, prefix, "fields")
	sortFields(fields, func(fieldName string, field dump.Value) {
		fragWriteFile(buf, prefix, "fields", fieldName+suffix)
		switch value := field.(type) {
		case dump.Struct:
			fragWriteDir(buf, prefix, "fields", fieldName)
			fragWriteFields(buf, path.Join(prefix, "fields", fieldName), suffix, value.Fields)
		}
	})
}

func fragWriteTypes(buf *bufio.Writer, prefix, suffix string, types dump.TypeDefs) {
	if len(types) == 0 {
		return
	}
	fragWriteDir(buf, prefix, "types")
	sortTypeDefs(types, func(defName string, def dump.TypeDef) {
		prefix := path.Join(prefix, "types", defName)
		fragWriteFile(buf, prefix+suffix)
		if len(def.Constructors) > 0 ||
			len(def.Properties) > 0 ||
			len(def.Methods) > 0 ||
			def.Operators != nil {
			fragWriteDir(buf, prefix)
		}
		if len(def.Constructors) > 0 {
			fragWriteDir(buf, prefix, "constructors")
			sortConstructors(def.Constructors, func(ctorName string, ctor dump.MultiFunction) {
				fragWriteFile(buf, prefix, "constructors", ctorName+suffix)
			})
		}
		if len(def.Properties) > 0 {
			fragWriteDir(buf, prefix, "properties")
			sortProperties(def.Properties, func(propName string, prop dump.Property) {
				fragWriteFile(buf, prefix, "properties", propName+suffix)
			})
		}
		if len(def.Methods) > 0 {
			fragWriteDir(buf, prefix, "methods")
			sortMethods(def.Methods, func(methodName string, method dump.Function) {
				fragWriteFile(buf, prefix, "methods", methodName+suffix)
			})
		}
		if op := def.Operators; op != nil {
			fragWriteDir(buf, prefix, "operators")
			if len(op.Add) > 0 {
				fragWriteFile(buf, prefix, "operators", "add"+suffix)
			}
			if len(op.Sub) > 0 {
				fragWriteFile(buf, prefix, "operators", "sub"+suffix)
			}
			if len(op.Mul) > 0 {
				fragWriteFile(buf, prefix, "operators", "mul"+suffix)
			}
			if len(op.Div) > 0 {
				fragWriteFile(buf, prefix, "operators", "div"+suffix)
			}
			if len(op.Mod) > 0 {
				fragWriteFile(buf, prefix, "operators", "div"+suffix)
			}
			if len(op.Pow) > 0 {
				fragWriteFile(buf, prefix, "operators", "pow"+suffix)
			}
			if len(op.Concat) > 0 {
				fragWriteFile(buf, prefix, "operators", "concat"+suffix)
			}
			if op.Eq {
				fragWriteFile(buf, prefix, "operators", "eq"+suffix)
			}
			if op.Le {
				fragWriteFile(buf, prefix, "operators", "le"+suffix)
			}
			if op.Lt {
				fragWriteFile(buf, prefix, "operators", "lt"+suffix)
			}
			if op.Len != nil {
				fragWriteFile(buf, prefix, "operators", "len"+suffix)
			}
			if op.Unm != nil {
				fragWriteFile(buf, prefix, "operators", "unm"+suffix)
			}
			if op.Call != nil {
				fragWriteFile(buf, prefix, "operators", "call"+suffix)
			}
			if op.Index != nil {
				fragWriteFile(buf, prefix, "operators", "index"+suffix)
			}
			if op.Newindex != nil {
				fragWriteFile(buf, prefix, "operators", "newindex"+suffix)
			}
		}
	})
}
