package scheme

import (
	"errors"
	"fmt"
	"github.com/anaminus/rbxmk"
	"github.com/robloxapi/rbxapi"
	"github.com/robloxapi/rbxfile"
	"github.com/robloxapi/rbxfile/declare"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func init() {
	registerInput("generate", rbxmk.InputScheme{
		Handler: generateInputSchemeHandler,
	})
}

func generateInputSchemeHandler(opt rbxmk.Options, node *rbxmk.InputNode, inref []string) (ext string, outref []string, data rbxmk.Data, err error) {
	ext = node.Format
	switch inref[0] {
	case "Instance":
		if len(inref) < 2 {
			err = errors.New("expected Instance reference")
			break
		}
		data, err = generateInstance(opt, inref[1])

	case "Property":
		if len(inref) < 2 {
			err = errors.New("expected Property reference")
			break
		}
		data, err = generateProperty(opt, inref[1])

	case "Value":
		if len(inref) < 2 {
			err = errors.New("expected Value reference")
			break
		}
		data, err = generateValue(opt, inref[1])
	}
	return ext, inref[2:], data, err
}

func generateInstance(opt rbxmk.Options, s string) (data rbxmk.Data, err error) {
	p := &genParser{s: s, refs: make(rbxfile.References), api: opt.API}
	if !p.parseClasses() {
		return nil, p.errmsg
	}
	if !p.eof() {
		if p.errmsg == nil {
			p.err(fmt.Errorf("expected EOF"))
		}
		return nil, p.errmsg
	}
	p.processRefs()
	return p.instances, nil
}

func generateProperty(opt rbxmk.Options, s string) (data rbxmk.Data, err error) {
	p := &genParser{s: s}
	if !p.parseProperties() {
		return nil, p.errmsg
	}
	if !p.eof() {
		if p.errmsg == nil {
			p.err(fmt.Errorf("expected EOF"))
		}
		return nil, p.errmsg
	}
	return p.instances[0].Properties, nil
}

func generateValue(opt rbxmk.Options, s string) (data rbxmk.Data, err error) {
	p := &genParser{s: s}
	var propType string
	propType, ok := p.parseName("type")
	if !ok {
		return nil, p.errmsg
	}
	typ := declare.TypeFromString(propType)
	if typ == 0 {
		p.err(fmt.Errorf("unknown type %q", propType))
		return nil, p.errmsg
	}
	var comp []interface{}
	if p.try("=") {
		if comp, ok = p.parseValue(); !ok {
			return nil, p.errmsg
		}
	}
	if !p.eof() {
		if p.errmsg == nil {
			p.err(fmt.Errorf("expected EOF"))
		}
		return nil, p.errmsg
	}
	return declare.Property("", typ, comp...).Declare(), nil
}

// GenError represents an error that occurred with the generate scheme parser.
type GenError struct {
	Offset int   // Location of the error.
	Err    error // Underlying error.
}

func (err GenError) Error() string {
	return fmt.Sprintf("parse error at %d: %s", err.Offset, err.Err.Error())
}

type genParser struct {
	s         string
	i         int
	errmsg    error
	api       *rbxapi.API
	instances []*rbxfile.Instance
	refs      rbxfile.References
	prefs     []rbxfile.PropRef
	ref       string
}

func (p *genParser) processRefs() {
	for _, pref := range p.prefs {
		p.refs.Resolve(pref)
	}
}

// Look at the next rune.
func (p *genParser) peek() (rune, int) {
	return utf8.DecodeRuneInString(p.s[p.i:])
}

// Increment location.
func (p *genParser) inc(n int) {
	if p.i+n > len(p.s) {
		p.i = len(p.s)
	} else if p.i+n < 0 {
		p.i = 0
	} else {
		p.i += n
	}
}

// Set or update error at current location.
func (p *genParser) err(e error) {
	if v, ok := e.(GenError); ok {
		e = v.Err
	}
	p.errmsg = GenError{Offset: p.i, Err: e}
}

func (p *genParser) parseDigit() bool {
	r, w := p.peek()
	if '0' <= r && r <= '9' {
		p.inc(w)
		return true
	}
	return false
}

func (p *genParser) parseHexdigit() bool {
	if p.parseDigit() {
		return true
	}
	r, w := p.peek()
	if ('a' <= r && r <= 'f') || ('A' <= r && r <= 'F') {
		p.inc(w)
		return true
	}
	return false
}

// Skip whitespace.
func (p *genParser) white() {
	for {
		r, w := p.peek()
		if !unicode.IsSpace(r) {
			return
		}
		p.inc(w)
	}
}

// Check if at EOF.
func (p *genParser) eof() bool {
	j := p.i
	defer func() {
		p.i = j
	}()
	p.white()
	return p.i == len(p.s)
}

// Consume if current location matches given string.
func (p *genParser) try(s string) bool {
	j := p.i
	p.white()
	if len(p.s[p.i:]) >= len(s) && p.s[p.i:p.i+len(s)] == s {
		p.inc(len(s))
		return true
	}
	p.i = j
	return false
}

// Match current location with given string.
func (p *genParser) look(s string) bool {
	j := p.i
	defer func() {
		p.i = j
	}()
	p.white()
	return len(p.s[p.i:]) >= len(s) && p.s[p.i:p.i+len(s)] == s
}

// Parse a boolean.
func (p *genParser) parseBool() (b, ok bool) {
	switch {
	case p.try("false"):
		return false, true
	case p.try("true"):
		return true, true
	}
	return false, false
}

// Parse a number (int, float, hex, oct).
func (p *genParser) parseNumber() (n float64, ok bool) {
	p.white()
	j := p.i
	defer func() {
		if !ok {
			p.i = j
		}
	}()

	n = 0
	neg := p.try("-")
	if p.try("0x") {
		// Hexidecimal
		if !p.parseHexdigit() {
			p.err(fmt.Errorf("expected hexidecimal digit"))
			return 0, false
		}
		for p.parseHexdigit() {
		}
		k := j + 2 // 0x
		if neg {
			k++
		}
		v, err := strconv.ParseInt(p.s[k:p.i], 16, 64)
		if err != nil {
			p.err(fmt.Errorf("error parsing hexidecimal: %s", err))
			return 0, false
		}
		n = float64(v)
		if neg {
			n = -n
		}
		return n, true
	}
	{
		r, _ := p.peek()
		hasDigit := false
		notOct := 0
		for p.parseDigit() {
			hasDigit = true
		}
		if p.try(".") {
			notOct++
			if !p.parseDigit() && !hasDigit {
				p.err(fmt.Errorf("expected digit"))
				return 0, false
			}
			for p.parseDigit() {
			}
		} else if !hasDigit {
			p.err(fmt.Errorf("expected digit"))
			return 0, false
		}
		if p.try("e") || p.try("E") {
			notOct++
			switch {
			case p.try("-"), p.try("+"):
			}
			if !p.parseDigit() {
				p.err(fmt.Errorf("expected digit"))
				return 0, false
			}
			for p.parseDigit() {
			}
		}
		if r == '0' && notOct == 0 {
			// Octal
			v, err := strconv.ParseInt(p.s[j:p.i], 8, 64)
			if err != nil {
				p.err(fmt.Errorf("error parsing octal: %s", err))
				return 0, false
			}
			n = float64(v)
		}
		// Float
		var err error
		if n, err = strconv.ParseFloat(p.s[j:p.i], 64); err != nil {
			p.err(fmt.Errorf("error parsing float: %s", err))
			return 0, false
		}
	}
	return n, true
}

// Parse a quoted string (" or ').
func (p *genParser) parseString() (str string, ok bool) {
	p.white()
	j := p.i
	defer func() {
		if !ok {
			p.i = j
		}
	}()

	var quot rune
	switch {
	case p.try(`"`):
		quot = '"'
	case p.try(`'`):
		quot = '\''
	default:
		p.err(fmt.Errorf("expected string"))
		return "", false
	}
	s := make([]rune, 0)
loop:
	for {
		r, w := p.peek()
		if p.eof() {
			p.err(fmt.Errorf("unexpected EOF"))
			return "", false
		}
		switch r {
		case '\\':
			p.inc(w)
			r, w = p.peek()
			switch {
			case p.parseDigit():
				p.parseDigit()
				p.parseDigit()
				n, err := strconv.ParseUint(p.s[j:p.i], 10, 8)
				if err != nil {
					p.err(fmt.Errorf("error in escape sequence: %s", err.(*strconv.NumError).Err))
					return "", false
				}
				s = append(s, rune(n))
			case r == 'x':
				p.inc(w)
				j := p.i
				for i := 0; i < 2; i++ {
					if !p.parseHexdigit() {
						if p.eof() {
							p.err(fmt.Errorf("non-hex character in escape sequence"))
						} else {
							p.err(fmt.Errorf("non-hex character in escape sequence: %s", string(p.s[p.i])))
						}
						return "", false
					}
				}
				n, err := strconv.ParseUint(p.s[j:p.i], 16, 8)
				if err != nil {
					p.err(fmt.Errorf("error in escape sequence: %s", err.(*strconv.NumError).Err))
					return "", false
				}
				s = append(s, rune(n))
			case r == 'u':
				p.inc(w)
				j := p.i
				for i := 0; i < 4; i++ {
					if !p.parseHexdigit() {
						if p.eof() {
							p.err(fmt.Errorf("non-hex character in escape sequence"))
						} else {
							p.err(fmt.Errorf("non-hex character in escape sequence: %s", string(p.s[p.i])))
						}
						return "", false
					}
				}
				n, err := strconv.ParseUint(p.s[j:p.i], 16, 16)
				if err != nil {
					p.err(fmt.Errorf("error in escape sequence: %s", err.(*strconv.NumError).Err))
					return "", false
				}
				s = append(s, rune(n))
			case r == 'U':
				p.inc(w)
				j := p.i
				for i := 0; i < 8; i++ {
					if !p.parseHexdigit() {
						if p.eof() {
							p.err(fmt.Errorf("non-hex character in escape sequence"))
						} else {
							p.err(fmt.Errorf("non-hex character in escape sequence: %s", string(p.s[p.i])))
						}
						return "", false
					}
				}
				n, err := strconv.ParseUint(p.s[j:p.i], 16, 32)
				if err != nil {
					p.err(fmt.Errorf("error in escape sequence: %s", err.(*strconv.NumError).Err))
					return "", false
				}
				s = append(s, rune(n))
			case r == 'a':
				s = append(s, '\a')
				p.inc(1)
			case r == 'b':
				s = append(s, '\b')
				p.inc(1)
			case r == 'f':
				s = append(s, '\f')
				p.inc(1)
			case r == 'n':
				s = append(s, '\n')
				p.inc(1)
			case r == 'r':
				s = append(s, '\r')
				p.inc(1)
			case r == 't':
				s = append(s, '\t')
				p.inc(1)
			case r == 'v':
				s = append(s, '\v')
				p.inc(1)
			case r == '\\':
				s = append(s, '\\')
				p.inc(1)
			case r == '\'':
				s = append(s, '\'')
				p.inc(1)
			case r == '"':
				s = append(s, '"')
				p.inc(1)
			default:
				p.err(fmt.Errorf("unknown escape sequence"))
				return "", false
			}
			r, w = p.peek()
		case quot:
			p.inc(w)
			break loop
		}
		s = append(s, r)
		p.inc(w)
	}
	return string(s), true
}

// Parse an identifier ([A-Za-z_][0-9A-Za-z_]*).
func (p *genParser) parseIdent() (str string, ok bool) {
	p.white()
	j := p.i

loop:
	for first := true; ; {
		r, w := p.peek()
		switch {
		case !first && '0' <= r && r <= '9':
		case 'a' <= r && r <= 'z':
		case 'A' <= r && r <= 'Z':
		case r == '_':
		default:
			break loop
		}
		p.inc(w)
		first = false
	}
	return p.s[j:p.i], true
}

// Parse a list of one or more classes (<class> { ';' <class> } [ ';' ]).
func (p *genParser) parseClasses() bool {
	name, ok := p.parseName("name")
	if !ok {
		return false
	}
	inst, ok := p.parseClass(name)
	if !ok {
		return false
	}
	p.instances = append(p.instances, inst)
	for p.try(";") {
		if p.eof() {
			break
		}
		name, ok := p.parseName("name")
		if !ok {
			return false
		}
		inst, ok := p.parseClass(name)
		if !ok {
			return false
		}
		p.instances = append(p.instances, inst)
	}
	return true
}

// Parse a single class (<name> '{' [ <item> { ';' <item> } [ ';' ] ] '}').
func (p *genParser) parseClass(name string) (inst *rbxfile.Instance, ok bool) {
	inst = rbxfile.NewInstance(name, nil)
	if !p.try("{") {
		p.err(fmt.Errorf("expected '{'"))
		return nil, false
	}
	if p.api != nil {
		// Fill in properties from API.
		if classAPI := p.api.Classes[name]; classAPI != nil {
			for _, member := range classAPI.MemberList() {
				if prop, _ := member.(*rbxapi.Property); prop != nil {
					if typ := rbxfile.TypeFromAPIString(p.api, prop.ValueType); typ != rbxfile.TypeInvalid {
						inst.Properties[prop.MemberName] = rbxfile.NewValue(typ)
					}
				}
			}
		}
	}
	if !p.try("}") {
		if !p.parseClassItem(inst) {
			return nil, false
		}
		for p.try(";") {
			if p.look("}") || p.eof() {
				break
			}
			if !p.parseClassItem(inst) {
				return nil, false
			}
		}
		if !p.try("}") {
			p.err(fmt.Errorf("expected '}'"))
			return nil, false
		}
	}
	return inst, true
}

// Parse one item within a class (<class> | <property> | <metaprop>).
func (p *genParser) parseClassItem(parent *rbxfile.Instance) bool {
	name, ok := p.parseName("name")
	if !ok {
		return false
	}
	switch {
	case p.look("{"):
		inst, ok := p.parseClass(name)
		if !ok {
			return false
		}
		inst.SetParent(parent)
	case p.look(":"), p.look("="):
		value, ok := p.parseProperty(name, parent)
		if !ok {
			return false
		}
		if _, ok := value.(rbxfile.ValueReference); ok && p.ref != "" {
			p.prefs = append(p.prefs, rbxfile.PropRef{
				Instance:  parent,
				Property:  name,
				Reference: p.ref,
			})
		} else {
			parent.Properties[name] = value
		}
	case p.look("("):
		if !p.parseMetaProp(parent, name) {
			return false
		}
	default:
		return false
	}
	return true
}

// Parse a meta property (<name> '(' <component> ')').
func (p *genParser) parseMetaProp(parent *rbxfile.Instance, name string) bool {
	if !p.try("(") {
		p.err(fmt.Errorf("expected '('"))
		return false
	}
	switch name {
	case "IsService":
		b, ok := p.parseBool()
		if !ok {
			p.err(fmt.Errorf("expected bool"))
			return false
		}
		parent.IsService = b
	case "Reference":
		ref, ok := p.parseName("name")
		if !ok {
			return false
		}
		parent.Reference = ref
		p.refs[ref] = parent
	default:
		p.err(fmt.Errorf("unknown meta property %q", name))
		return false
	}
	if !p.try(")") {
		p.err(fmt.Errorf("expected ')'"))
		return false
	}
	return true
}

// Parse a list of one or more properties (<property> { ';' <property> } [ ';' ]).
func (p *genParser) parseProperties() bool {
	p.instances = []*rbxfile.Instance{rbxfile.NewInstance("", nil)}
	name, ok := p.parseName("name")
	if !ok {
		return false
	}
	value, ok := p.parseProperty(name, nil)
	if !ok {
		return false
	}
	p.instances[0].Properties[name] = value
	for p.try(";") {
		if p.eof() {
			break
		}
		name, ok := p.parseName("name")
		if !ok {
			return false
		}
		value, ok := p.parseProperty(name, nil)
		if !ok {
			return false
		}
		p.instances[0].Properties[name] = value
	}
	return true
}

// Parse a single property (<name> [ ':' <type> ] '=' <value>).
func (p *genParser) parseProperty(name string, parent *rbxfile.Instance) (value rbxfile.Value, ok bool) {
	var propType string
	if p.api != nil && parent != nil {
		var propDesc *rbxapi.Property
		if classDesc := p.api.Classes[parent.ClassName]; classDesc != nil {
			propDesc, _ = classDesc.Members[name].(*rbxapi.Property)
		}
		if p.try(":") {
			if propType, ok = p.parseName("type"); !ok {
				return nil, false
			}
			propType = rbxfile.TypeFromAPIString(p.api, propType).String()
			if propDesc != nil && propType != rbxfile.TypeFromAPIString(p.api, propDesc.ValueType).String() {
				p.err(fmt.Errorf("expected type %s for property %s.%s (got %s)", propDesc.ValueType, parent.ClassName, name, propType))
				return nil, false
			}
		} else if propDesc != nil {
			propType = rbxfile.TypeFromAPIString(p.api, propDesc.ValueType).String()
		} else {
			p.err(fmt.Errorf("expected ':'"))
			return nil, false
		}
	} else {
		if !p.try(":") {
			p.err(fmt.Errorf("expected ':'"))
			return nil, false
		}
		if propType, ok = p.parseName("type"); !ok {
			return nil, false
		}
	}
	typ := declare.TypeFromString(propType)
	if typ == 0 {
		p.err(fmt.Errorf("unknown type %q", propType))
		return nil, false
	}
	var comp []interface{}
	if p.try("=") {
		if comp, ok = p.parseValue(); !ok {
			return nil, false
		}
	}
	if typ == declare.Reference {
		p.ref = ""
		if len(comp) > 0 {
			if v, ok := comp[0].(string); ok {
				p.ref = v
			}
		}
	}
	return declare.Property(name, typ, comp...).Declare(), true
}

// Parse a name (<string> | <ident>).
func (p *genParser) parseName(t string) (s string, ok bool) {
	if s, ok = p.parseString(); ok {
		return s, true
	}
	if p.look("\"") || p.look("'") {
		return "", false
	}
	p.errmsg = nil
	if s, ok = p.parseIdent(); ok && s != "" {
		return s, true
	}
	p.err(fmt.Errorf("expected %s", t))
	return "", false
}

// Parse a value (<component> { ',' <component> } [ ',' ]).
func (p *genParser) parseValue() (comp []interface{}, ok bool) {
	if v, ok := p.parseComponent(); !ok {
		return nil, false
	} else {
		comp = append(comp, v)
	}
	for p.try(",") {
		if p.look(";") || p.look("}") || p.eof() {
			break
		}
		if v, ok := p.parseComponent(); !ok {
			return nil, false
		} else {
			comp = append(comp, v)
		}
	}
	return comp, true
}

// Parse a component (<bool> | <number> | <string>).
func (p *genParser) parseComponent() (v interface{}, ok bool) {
	if v, ok = p.parseBool(); ok {
		return v, true
	}
	if v, ok = p.parseNumber(); ok {
		return v, true
	}
	if r, _ := p.peek(); '0' <= r && r <= '9' {
		return nil, false
	}
	if v, ok = p.parseString(); ok {
		return v, true
	}
	if r, _ := p.peek(); r != '"' && r != '\'' {
		p.err(fmt.Errorf("expected value"))
	}
	return nil, false
}
