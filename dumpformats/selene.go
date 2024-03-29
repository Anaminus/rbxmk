package dumpformats

import (
	"bufio"
	"fmt"
	"io"

	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/dump/dt"
)

func init() { register(Selene) }

var Selene = Format{
	Name: "selene",
	Func: func(w io.Writer, root dump.Root, opts Options) error {
		buf := bufio.NewWriter(w)
		buf.WriteString("[selene]\n")
		buf.WriteString("name = \"rbxmk\"\n")
		structTypes := map[string]struct{}{}
		for _, library := range root.Libraries {
			for t := range library.Types {
				structTypes[t] = struct{}{}
			}
		}
		root.Libraries.ForEach(func(name string, library dump.Library) bool {
			// Library name comment.
			buf.WriteString("\n# ")
			buf.WriteString(name)
			buf.WriteString(" library\n")

			// Structs.
			sortTypeDefs(library.Types, func(defName string, def dump.TypeDef) {
				// Skip if empty.
				if def.Underlying.Kind == nil &&
					len(def.Properties) == 0 &&
					len(def.Methods) == 0 &&
					def.Operators == nil {
					return
				}

				// Library name comment.
				buf.WriteString("# ")
				buf.WriteString(defName)
				buf.WriteString(" struct\n")

				// Struct field.
				buf.WriteString("[selene.structs.")
				buf.WriteString(defName)
				buf.WriteString("]\n")

				// Type field.
				if t := def.Underlying.Kind; t != nil {
					seleneWriteTypeField(buf, structTypes, t, false)
					buf.WriteString("\n")
				}

				// Properties.
				sortProperties(def.Properties, func(propName string, prop dump.Property) {
					if !isName(propName) {
						return
					}
					buf.WriteString("[selene.structs.")
					buf.WriteString(defName)
					buf.WriteString(".")
					buf.WriteString(propName)
					buf.WriteString("]\n")
					seleneWriteProperty(buf, structTypes, prop)
				})

				// Methods.
				sortMethods(def.Methods, func(methodName string, method dump.Function) {
					if !isName(methodName) {
						return
					}
					buf.WriteString("[selene.structs.")
					buf.WriteString(defName)
					buf.WriteString(".")
					buf.WriteString(methodName)
					buf.WriteString("]\n")
					buf.WriteString("\tmethod = true\n")
					seleneWriteParameters(buf, structTypes, method.Parameters)
				})

				// Writability.
				if op := def.Operators; op != nil {
					switch {
					case op.Index != nil && op.Newindex != nil:
						buf.WriteString("[selene.structs.")
						buf.WriteString(defName)
						buf.WriteString(".\"*\"")
						buf.WriteString("]\n")
						buf.WriteString("\tproperty = true\n")
						if len(op.Index.Returns) > 0 {
							seleneWriteTypeField(buf, structTypes, op.Index.Returns[0].Type.Kind, false)
							buf.WriteString("\n")
						}
						buf.WriteString("\twritable = \"overridden\"\n")
					case op.Index != nil:
						buf.WriteString("[selene.structs.")
						buf.WriteString(defName)
						buf.WriteString(".\"*\"")
						buf.WriteString("]\n")
						buf.WriteString("\tproperty = true\n")
						if len(op.Index.Returns) > 0 {
							seleneWriteTypeField(buf, structTypes, op.Index.Returns[0].Type.Kind, false)
							buf.WriteString("\n")
						}
					}
				}

				// Constructors.
				sortConstructors(def.Constructors, func(ctorName string, ctors dump.MultiFunction) {
					buf.WriteString("[")
					if len(library.Import) > 0 {
						buf.WriteString(library.ImportString())
						buf.WriteString(".")
					}
					buf.WriteString(defName)
					buf.WriteString(".")
					buf.WriteString(ctorName)
					buf.WriteString("]\n")
					seleneWriteMultiFunction(buf, structTypes, ctors)
				})

				buf.WriteString("\n")
			})

			// Globals.
			seleneWriteStruct(buf, structTypes, library.ImportString(), library.Struct)
			return true
		})
		return buf.Flush()
	},
}

func seleneWriteStruct(buf *bufio.Writer, structTypes map[string]struct{}, parent string, str dump.Struct) {
	sortFields(str.Fields, func(fieldName string, field dump.Value) {
		buf.WriteString("[")
		if parent != "" {
			fieldName = parent + "." + fieldName
		}
		buf.WriteString(fieldName)
		buf.WriteString("]\n")
		switch v := field.(type) {
		case dump.Property:
			seleneWriteProperty(buf, structTypes, v)
			if p, ok := v.ValueType.Kind.(dt.KindPrim); ok && string(p) == "table" && v.ReadOnly {
				buf.WriteString("\twritable = \"new-fields\"\n")
			}
		case dump.Function:
			seleneWriteParameters(buf, structTypes, v.Parameters)
		case dump.MultiFunction:
			seleneWriteMultiFunction(buf, structTypes, v)
		case dump.Struct:
			seleneWriteStruct(buf, structTypes, fieldName, v)
		}
	})
}

func seleneWriteProperty(buf *bufio.Writer, structTypes map[string]struct{}, prop dump.Property) {
	buf.WriteString("\tproperty = true\n")
	seleneWriteTypeField(buf, structTypes, prop.ValueType.Kind, false)
	buf.WriteString("\n")
	if !prop.ReadOnly {
		buf.WriteString("\twritable = \"overridden\"\n")
	}
}

func seleneWriteMultiFunction(buf *bufio.Writer, structTypes map[string]struct{}, funcs dump.MultiFunction) {
	var fn dump.Function
	switch len(funcs) {
	case 0:
		return
	case 1:
		fn = funcs[0]
	default:
		// Generate single function with merged arguments.
		min, max := -1, 0
		for _, fn := range funcs {
			n := len(fn.Parameters)
			if n > max {
				max = n
			}
			if n < min || min < 0 {
				min = n
			}
		}
		if max > 0 {
			fn.Parameters = make(dump.Parameters, max)
			for _, fn := range funcs {
				// If the type of the nth parameter from two
				// functions do not match, convert to any.
				for i, param := range fn.Parameters {
					if fn.Parameters[i].Type.Kind == nil {
						fn.Parameters[i].Type = param.Type
					} else if fn.Parameters[i].Type != param.Type {
						fn.Parameters[i].Type = dt.Prim("any")
					}
				}
			}
			// Parameters after the minimum number of arguments
			// are treated as optional.
			for i := min; i < max; i++ {
				if _, ok := fn.Parameters[i].Type.Kind.(dt.KindOptional); !ok {
					fn.Parameters[i].Type = dt.Optional(fn.Parameters[i].Type)
				}
			}
		}
	}
	seleneWriteParameters(buf, structTypes, fn.Parameters)
}

func seleneWriteParameters(buf *bufio.Writer, structTypes map[string]struct{}, params dump.Parameters) {
	if len(params) == 0 {
		buf.WriteString("\targs = []\n")
		return
	}
	buf.WriteString("\targs = [")
	for i, param := range params {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString("\n\t\t{")

		// Try variadic.
		if param.Name == "..." {
			buf.WriteString("type = \"...\"")
			goto continueArg
		}

		// Try enums.
		if len(param.Enums) > 0 {
			for _, enum := range param.Enums {
				// Must be a list of strings.
				if len(enum) < 2 || enum[0] != '"' || enum[len(enum)-1] != '"' {
					goto skipEnum
				}
			}
			buf.WriteString("type = [")
			for i, enum := range param.Enums {
				if i > 0 {
					buf.WriteString(", ")
				}
				buf.WriteString(enum)
			}
			buf.WriteString("]")
			goto continueArg
		}
	skipEnum:

		seleneWriteTypeField(buf, structTypes, param.Type.Kind, true)

	continueArg:
		// Check optional.
		if _, ok := param.Type.Kind.(dt.KindOptional); ok {
			buf.WriteString(", required = false")
		}

		// Finish argument.
		buf.WriteString("}")
	}
	buf.WriteString("\n\t]\n")
}

func seleneWriteTypeField(buf *bufio.Writer, structTypes map[string]struct{}, t dt.Kind, arg bool) {
	if !arg {
		buf.WriteString("\t")
	}

	switch t.(type) {
	case dt.KindArray:
		buf.WriteString("type = \"table\"")
		return
	case dt.KindStruct:
		buf.WriteString("type = \"table\"")
		return
	case dt.KindDictionary:
		buf.WriteString("type = \"table\"")
		return
	case dt.KindMap:
		buf.WriteString("type = \"table\"")
		return
	case dt.KindTable:
		buf.WriteString("type = \"table\"")
		return
	case dt.KindFunction:
		buf.WriteString("type = \"function\"")
		return
	}

	// Try known primitive.
	if p := getPrim(t); p != nil {
		switch p.String() {
		case "boolean":
			buf.WriteString("type = \"bool\"")
			return
		case "int", "int64", "float", "double":
			buf.WriteString("type = \"number\"")
			return
		case "any", "bool", "function", "nil", "number", "string", "table":
			buf.WriteString("type = ")
			seleneEscapeString(buf, p.String())
			return
		case "Variant":
			buf.WriteString("type = \"any\"")
			return
		}
	}

	// Try struct.
	if p := getPrim(t); p != nil {
		if _, ok := structTypes[p.String()]; ok {
			if arg {
				buf.WriteString("type = {display = ")
				seleneEscapeString(buf, p.String())
				buf.WriteString("}")
			} else {
				buf.WriteString("struct = ")
				seleneEscapeString(buf, p.String())
			}
			return
		}
	}

	// Fallback to any.
	buf.WriteString("type = \"any\"")
}

// Get the underlying primitive type.
func getPrim(t dt.Kind) dt.Kind {
	switch t := t.(type) {
	case dt.KindPrim:
		return t
	case dt.KindOptional:
		return getPrim(t.Type.Kind)
	case dt.KindGroup:
		return getPrim(t.Type.Kind)
	default:
		return nil
	}
}

// Write an escaped TOML string.
func seleneEscapeString(w *bufio.Writer, s string) {
	w.WriteByte('"')
	for _, r := range s {
		switch {
		case r == '\b':
			w.WriteString(`\b`)
		case r == '\t':
			w.WriteString(`\t`)
		case r == '\n':
			w.WriteString(`\n`)
		case r == '\r':
			w.WriteString(`\r`)
		case r == '"':
			w.WriteString(`\"`)
		case r == '\\':
			w.WriteString(`\\`)
		case r < 0x20:
			fmt.Fprintf(w, `\u%04X`, r)
		default:
			w.WriteRune(r)
		}
	}
	w.WriteByte('"')
}
