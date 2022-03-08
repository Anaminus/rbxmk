package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"sync"
	"text/template"

	"github.com/PuerkitoBio/goquery"
	"github.com/anaminus/drill"
	"github.com/anaminus/drill/filesys"
	"github.com/anaminus/rbxmk/fragments"
	"github.com/anaminus/rbxmk/rbxmk/htmldrill"
	"github.com/anaminus/rbxmk/rbxmk/term"
	"golang.org/x/net/html"
	terminal "golang.org/x/term"
)

// Fragments are formatted as HTML templates. They are parsed and rendered as
// regular HTML by the drill, but this is lossless enough that that template
// directives survive the process.
//
// More technically, a particular fragment file is a concatenation of a number
// of independently evaluated template fragments. That is, the data in one
// section does not necessarily correspond to the data in another section.
//
// The resulting render is a HTML template, ready for evaluation. Data produced
// by the execution are expected to be formatted in HTML.
//
// The final result of the template is regular HTML. This result is passed to a
// configured renderer that converts it to a finalized format.

// Language determines the language of documentation text.
var Language = "en-us"

// panicLanguage writes the languages that have been embedded, then panics.
func panicLanguage() {
	langs := make([]string, 0, len(fragments.Languages))
	for lang := range fragments.Languages {
		langs = append(langs, lang)
	}
	sort.Strings(langs)
	if len(langs) == 0 {
		fmt.Fprintln(os.Stderr, "no languages are embedded")
	} else {
		fmt.Fprintln(os.Stderr, "the following languages are embedded:")
		for _, lang := range langs {
			fmt.Fprintf(os.Stderr, "\t%s\n", lang)
		}
	}
	panic(fmt.Sprintf("unsupported language %q", Language))
}

func initFrags() drill.Node {
	lang, ok := fragments.Languages[Language]
	if !ok {
		panicLanguage()
	}
	f, err := filesys.NewFS(lang, filesys.Handlers{
		{Pattern: "*.html", Func: htmldrill.NewHandler()},
	})
	if err != nil {
		panic(err)
	}
	node := f.UnorderedChild(Language)
	if node == nil {
		panicLanguage()
	}
	return node
}

var fragfs = initFrags()
var fragMut sync.RWMutex
var fragCache = map[string]*htmldrill.Node{}

// parseFragRef receives a fragment reference and converts it to a file path.
//
// A fragment reference is like a regular file path, except that filesep is the
// separator that descends into a file. After the first occurrence of filesep,
// the preceding element is appended with the given suffix, and then descending
// continues.
//
// For example, with ".md" and ":":
//
//     libraries/roblox/types/Region3:Properties/CFrame/Description
//
// is split into the following elements:
//
//     libraries, roblox, types, Region3.md, Properties, CFrame, Description
//
// The file portion of the fragment reference is converted to lowercase.
//
// If no separator was found in the reference, then the final element will have
// suffix appended unless dir is true.
//
// infile returns whether the reference drilled into a file.
func parseFragRef(s, suffix string, filesep rune, dir bool) (items []string, infile bool) {
	if s == "" {
		return []string{}, false
	}
	i := strings.IndexRune(s, filesep)
	if i < 0 {
		if dir {
			return strings.Split(strings.ToLower(s), "/"), false
		} else {
			return strings.Split(strings.ToLower(s)+suffix, "/"), false
		}
	}
	items = make([]string, 0, strings.Count(s, "/")+1)
	items = append(items, strings.Split(strings.ToLower(s[:i])+suffix, "/")...)
	items = append(items, strings.Split(s[i+1:], "/")...)
	return items, true
}

type Renderer = func(w io.Writer, s *goquery.Selection) error

type FuncMap = template.FuncMap

var fragTmplFuncs = FuncMap{
	// List of top-level fragment topics.
	"Topics": func() string {
		return "\n\t" + strings.Join(ListFragments(""), "\n\t")
	},
}

type FragOptions struct {
	// Data included with the executed template.
	TmplData interface{}
	// Functions included with the executed template.
	TmplFuncs FuncMap
	// Renderer used if a node is htmldrill.Node.
	Renderer Renderer
}

// ExecuteFragTmpl renders converts the result of node, in template format, to a
// final rendering.
func ExecuteFragTmpl(fragref string, node drill.Node, opt FragOptions) string {
	// Parse template.
	tmplText := strings.TrimSpace(node.Fragment())
	t := template.New("root")
	t.Funcs(fragTmplFuncs)
	t.Funcs(opt.TmplFuncs)
	t, err := t.Parse(tmplText)
	if err != nil {
		panic(fmt.Errorf("parse template %q: %w", fragref, err))
	}

	// Execute template.
	var buf bytes.Buffer
	err = t.Execute(&buf, opt.TmplData)
	if err != nil {
		panic(fmt.Errorf("execute template %q: %w", fragref, err))
	}

	// If no renderer, return directly as HTML.
	if opt.Renderer == nil {
		return strings.TrimSpace(buf.String())
	}

	// Parse HTML.
	root, err := html.Parse(&buf)
	if err != nil {
		panic(fmt.Errorf("parse HTML %q: %w", fragref, err))
	}
	doc := goquery.NewDocumentFromNode(root)

	// Render HTML.
	buf.Reset()
	if err := opt.Renderer(&buf, doc.Selection); err != nil {
		panic(fmt.Errorf("render HTML %q: %w", fragref, err))
	}
	return strings.TrimSpace(buf.String())
}

// ResolveFragmentWith is like ResolveFragment, but with configurable options.
func ResolveFragmentWith(fragref string, opt FragOptions) string {
	fragMut.Lock()
	defer fragMut.Unlock()
	node, _ := resolveFragmentNode(fragref, false)
	if node == nil {
		return ""
	}
	return ExecuteFragTmpl(fragref, node, opt)
}

// ResolveFragment returns the content of the fragment referred to by fragref.
// Returns an empty string if no content was found.
func ResolveFragment(fragref string) string {
	return ResolveFragmentWith(fragref, FragOptions{})
}

const FragSep = ':'
const FragSuffix = ".html"

func ListFragments(fragref string) []string {
	frags := map[string]struct{}{}
	if fragref == "" {
		node, _ := resolveFragmentNode("", false)
		switch node := node.(type) {
		case drill.UnorderedBranch:
			children := node.UnorderedChildren()
			for name := range children {
				if name != "" {
					name = strings.TrimSuffix(name, FragSuffix)
					frags[name] = struct{}{}
				}
			}
		}
	} else {
		node, infile := resolveFragmentNode(fragref, false)
		switch node := node.(type) {
		case drill.UnorderedBranch:
			children := node.UnorderedChildren()
			for name := range children {
				if name != "" {
					if infile {
						name = strings.TrimSuffix(name, FragSuffix)
						frags["/"+name] = struct{}{}
					} else {
						frags[string(FragSep)+name] = struct{}{}
					}
				}
			}
		}
		if node == nil || !infile {
			node, infile := resolveFragmentNode(fragref, true)
			switch node := node.(type) {
			case drill.UnorderedBranch:
				children := node.UnorderedChildren()
				for name := range children {
					if name != "" {
						if !infile {
							name = strings.TrimSuffix(name, FragSuffix)
						}
						frags["/"+name] = struct{}{}
					}
				}
			}
		}
	}
	list := make([]string, 0, len(frags))
	for frag := range frags {
		list = append(list, frag)
	}
	sort.Strings(list)
	return list
}

// resolveFragmentNode resolves a fragment reference into a drill Node by
// walking through the components of the reference.
func resolveFragmentNode(fragref string, dir bool) (n drill.Node, infile bool) {
	n = fragfs
	var path string
	names, infile := parseFragRef(fragref, FragSuffix, FragSep, dir)
	for _, name := range names {
		if name == "" {
			return nil, false
		}
		path += "/" + name
		if node, ok := fragCache[path]; ok {
			n = node
		} else {
			switch v := n.(type) {
			case drill.UnorderedBranch:
				n = v.UnorderedChild(name)
			default:
				return nil, false
			}
			if node, ok := n.(*htmldrill.Node); ok {
				fragCache[path] = node
			}
		}
	}
	return n, infile
}

// ErrorFrag returns an error according to the fragment section of the given
// name. The result is passed to fmt.Errorf with args.
func ErrorFrag(name string, args ...interface{}) error {
	format := ResolveFragment("Errors:" + name)
	return fmt.Errorf(strings.TrimSpace(format), args...)
}

// FormatFrag returns a formatted string according to the fragment of the given
// reference. The result is passed to fmt.Sprintf with args.
func FormatFrag(fragref string, args ...interface{}) string {
	format := ResolveFragment(fragref)
	return fmt.Sprintf(strings.TrimSpace(format), args...)
}

var docFailed = map[string]struct{}{}
var docSeen = map[string]struct{}{}

// DocFragments returns a list of requested fragments.
func DocFragments() []string {
	fragMut.RLock()
	defer fragMut.RUnlock()
	frags := make([]string, 0, len(docSeen))
	for frag := range docSeen {
		frags = append(frags, frag)
	}
	sort.Strings(frags)
	return frags
}

// UnresolvedFragments writes to stderr a list of fragment references that
// failed to resolve. Panics if any references failed.
func UnresolvedFragments() {
	if len(docFailed) == 0 {
		return
	}
	refs := make([]string, 0, len(docFailed))
	for ref := range docFailed {
		refs = append(refs, ref)
	}
	sort.Strings(refs)
	var s strings.Builder
	fmt.Fprintf(&s, "unresolved fragments for %q:\n", Language)
	for _, ref := range refs {
		fmt.Fprintf(&s, "\t%s\n", ref)
	}
	fmt.Fprintf(&s, "\nversion: %s", VersionString())
	panic(s.String())
}

// DocWith is like Doc, but with configurable options.
func DocWith(fragref string, opt FragOptions) string {
	fragMut.Lock()
	defer fragMut.Unlock()
	docSeen[fragref] = struct{}{}
	node, _ := resolveFragmentNode(fragref, false)
	if node == nil {
		docFailed[fragref] = struct{}{}
		return "{" + Language + ":" + fragref + "}"
	}
	return ExecuteFragTmpl(fragref, node, opt)
}

// Doc returns the content of the fragment referred to by fragref. The given
// path is marked to be returned by DocFragments. If no content was found, then
// a string indicating an unresolved reference is returned.
//
// Doc should be used only to process descriptions for command-line elements.
// Descriptions are rendered in a format suitable for the terminal.
// ResolveFragment can be used to resolve a reference without marking it.
func Doc(fragref string) string {
	termWidth, _, _ := terminal.GetSize(int(os.Stdout.Fd()))
	return DocWith(fragref, FragOptions{
		Renderer: term.Renderer{Width: termWidth, TabSize: 4}.Render,
	})
}

// DocFlag returns the content for a flag by configuring the renderer with 0
// width, so that it can be properly formatted by the usage template.
func DocFlag(fragref string) string {
	return DocWith(fragref, FragOptions{
		Renderer: term.Renderer{Width: 0, TabSize: 4}.Render,
	})
}
