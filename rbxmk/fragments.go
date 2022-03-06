package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"text/template"

	"github.com/anaminus/drill"
	"github.com/anaminus/drill/filesys"
	"github.com/anaminus/rbxmk/fragments"
	"github.com/anaminus/rbxmk/rbxmk/htmldrill"
	"github.com/anaminus/rbxmk/rbxmk/term"
	terminal "golang.org/x/term"
)

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

func initDocs() drill.Node {
	termWidth, _, _ := terminal.GetSize(int(os.Stdout.Fd()))
	lang, ok := fragments.Languages[Language]
	if !ok {
		panicLanguage()
	}
	f, err := filesys.NewFS(lang, filesys.Handlers{
		{Pattern: "*.html", Func: htmldrill.NewHandler(
			htmldrill.WithRenderer(term.Renderer{Width: termWidth, TabSize: 4}.Render),
		)},
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

var docfs = initDocs()
var docMut sync.RWMutex
var docCache = map[string]*htmldrill.Node{}
var docFailed = map[string]struct{}{}
var docSeen = map[string]struct{}{}

// DocFragments returns a list of requested fragments.
func DocFragments() []string {
	docMut.RLock()
	defer docMut.RUnlock()
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

type FuncMap = template.FuncMap

var docTmplFuncs = FuncMap{
	// List of top-level fragment topics.
	"Topics": func() string {
		return "\n\t" + strings.Join(ListFragments(""), "\n\t")
	},
}

func executeDocTmpl(fragref, tmplText string, data interface{}, funcs FuncMap) string {
	t := template.New("root")
	t.Funcs(docTmplFuncs)
	t.Funcs(funcs)
	t, err := t.Parse(tmplText)
	if err != nil {
		panic(fmt.Errorf("parse %q: %w", fragref, err))
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		panic(fmt.Errorf("execute %q: %w", fragref, err))
	}
	return strings.TrimSpace(buf.String())
}

type FragOptions struct {
	// Data included with the executed template.
	TmplData interface{}
	// Functions included with the executed template.
	TmplFuncs FuncMap
	// Renderer used if a node is htmldrill.Node.
	Renderer htmldrill.Renderer
}

// Doc returns the content of the fragment referred to by fragref. The given
// path is marked to be returned by DocFragments. If no content was found, then
// a string indicating an unresolved reference is returned.
//
// Doc should only be used to capture additional fragment references.
// ResolveFragment can be used to resolve a reference without marking it.
//
// The content of the fragment executed as a template with docTmplFuncs included
// as functions.
func Doc(fragref string) string {
	return DocWith(fragref, FragOptions{})
}

// DocWith is like Doc, but with configurable options.
func DocWith(fragref string, opt FragOptions) string {
	docMut.Lock()
	defer docMut.Unlock()
	docSeen[fragref] = struct{}{}
	node, _ := resolveFragmentNode(fragref, false)
	if node == nil {
		docFailed[fragref] = struct{}{}
		return "{" + Language + ":" + fragref + "}"
	}
	if opt.Renderer != nil {
		if n, ok := node.(*htmldrill.Node); ok {
			node = n.WithRenderer(opt.Renderer)
		}
	}
	tmplText := strings.TrimSpace(node.Fragment())
	return executeDocTmpl(fragref, tmplText, opt.TmplData, opt.TmplFuncs)
}

// ResolveFragment returns the content of the fragment referred to by fragref.
// Returns an empty string if no content was found.
//
// The content of the fragment executed as a template with docTmplFuncs included
// as functions.
func ResolveFragment(fragref string) string {
	return ResolveFragmentWith(fragref, FragOptions{})
}

// ResolveFragmentWith is like ResolveFragment, but with configurable options.
func ResolveFragmentWith(fragref string, opt FragOptions) string {
	docMut.Lock()
	defer docMut.Unlock()
	node, _ := resolveFragmentNode(fragref, false)
	if node == nil {
		return ""
	}
	if opt.Renderer != nil {
		if n, ok := node.(*htmldrill.Node); ok {
			node = n.WithRenderer(opt.Renderer)
		}
	}
	tmplText := strings.TrimSpace(node.Fragment())
	return executeDocTmpl(fragref, tmplText, opt.TmplData, opt.TmplFuncs)
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
	n = docfs
	var path string
	names, infile := parseFragRef(fragref, FragSuffix, FragSep, dir)
	for _, name := range names {
		if name == "" {
			return nil, false
		}
		path += "/" + name
		if node, ok := docCache[path]; ok {
			n = node
		} else {
			switch v := n.(type) {
			case drill.UnorderedBranch:
				n = v.UnorderedChild(name)
			default:
				return nil, false
			}
			if node, ok := n.(*htmldrill.Node); ok {
				docCache[path] = node
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
