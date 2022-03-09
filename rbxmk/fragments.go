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

// initFragRoot initializes the top-level drill node, starting from the root of
// the embedded fragment file system.
func initFragRoot() drill.Node {
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

// ParseFragRef receives a fragment reference and converts it to a file path.
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
func ParseFragRef(s, suffix string, filesep rune, dir bool) (items []string, infile bool) {
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

// Renderer renders a finalized HTML fragment.
type Renderer = func(w io.Writer, s *goquery.Selection) error

// Global template functions.
var fragTmplFuncs = template.FuncMap{
	// List of top-level fragment topics.
	"Topics": func() string {
		return "\n\t" + strings.Join(Frag.List(""), "\n\t")
	},
}

// FragOptions configure ExecuteFragTmpl.
type FragOptions struct {
	// Data included with the executed template.
	TmplData interface{}
	// Functions included with the executed template.
	TmplFuncs template.FuncMap
	// Renderer used if a node is htmldrill.Node.
	Renderer Renderer
}

// ExecuteFragTmpl converts the result of node, in template format, to a final
// rendering.
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

// Fragments provides methods for resolving fragments.
type Fragments struct {
	mut   sync.RWMutex
	root  drill.Node
	cache map[string]*htmldrill.Node

	// Sep is the separator used for fragment references to separate path
	// content from section content. Defaults to ':'.
	Sep rune

	// Suffix is the Suffix applied to the base file portion of fragment
	// references. Defaults to ".html".
	Suffix string
}

// NewFragments returns a new Fragments initialized with root.
func NewFragments(root drill.Node) *Fragments {
	return &Fragments{
		root:   root,
		cache:  map[string]*htmldrill.Node{},
		Sep:    ':',
		Suffix: ".html",
	}
}

// Resolve returns the content of the fragment referred to by fragref.
// Returns an empty string if no content was found.
func (f *Fragments) Resolve(fragref string) string {
	return f.ResolveWith(fragref, FragOptions{})
}

// ResolveWith is like Resolve, but with configurable options.
func (f *Fragments) ResolveWith(fragref string, opt FragOptions) string {
	f.mut.Lock()
	defer f.mut.Unlock()
	node, _ := f.resolveNode(fragref, false)
	if node == nil {
		return ""
	}
	return ExecuteFragTmpl(fragref, node, opt)
}

// List returns a list of child fragment references available under the fragment
// referred to by fragref.
func (f *Fragments) List(fragref string) []string {
	frags := map[string]struct{}{}
	if fragref == "" {
		node, _ := f.resolveNode("", false)
		switch node := node.(type) {
		case drill.UnorderedBranch:
			children := node.UnorderedChildren()
			for name := range children {
				if name != "" {
					name = strings.TrimSuffix(name, f.Suffix)
					frags[name] = struct{}{}
				}
			}
		}
	} else {
		node, infile := f.resolveNode(fragref, false)
		switch node := node.(type) {
		case drill.UnorderedBranch:
			children := node.UnorderedChildren()
			for name := range children {
				if name != "" {
					if infile {
						name = strings.TrimSuffix(name, f.Suffix)
						frags["/"+name] = struct{}{}
					} else {
						frags[string(f.Sep)+name] = struct{}{}
					}
				}
			}
		}
		if node == nil || !infile {
			node, infile := f.resolveNode(fragref, true)
			switch node := node.(type) {
			case drill.UnorderedBranch:
				children := node.UnorderedChildren()
				for name := range children {
					if name != "" {
						if !infile {
							name = strings.TrimSuffix(name, f.Suffix)
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

// resolveNode resolves a fragment reference into a drill Node by walking
// through the components of the reference. If found, returns the node, and
// whether the node has drilled into a file.
//
// If fragref does not contain a separator, then the last element is assumed to
// be a directory if dir is true, and a file otherwise.
func (f *Fragments) resolveNode(fragref string, dir bool) (n drill.Node, infile bool) {
	n = f.root
	var path string
	names, infile := ParseFragRef(fragref, f.Suffix, f.Sep, dir)
	for _, name := range names {
		if name == "" {
			return nil, false
		}
		path += "/" + name
		if node, ok := f.cache[path]; ok {
			n = node
		} else {
			switch v := n.(type) {
			case drill.UnorderedBranch:
				n = v.UnorderedChild(name)
			default:
				return nil, false
			}
			if node, ok := n.(*htmldrill.Node); ok {
				f.cache[path] = node
			}
		}
	}
	return n, infile
}

// Error returns an error according to the fragment section of the given
// name. The result is passed to fmt.Errorf with args.
func (f *Fragments) Error(name string, args ...interface{}) error {
	format := f.ResolveWith("Errors:"+name, FragOptions{
		Renderer: term.Renderer{Width: 0, TabSize: 4}.Render,
	})
	return fmt.Errorf(strings.TrimSpace(format), args...)
}

// Format returns a formatted string according to the fragment of the given
// reference. The result is passed to fmt.Sprintf with args.
func (f *Fragments) Format(fragref string, args ...interface{}) string {
	format := f.ResolveWith(fragref, FragOptions{
		Renderer: term.Renderer{Width: 0, TabSize: 4}.Render,
	})
	return fmt.Sprintf(strings.TrimSpace(format), args...)
}

// DocState contains the state of fragments used to resolve documentation for
// commands.
type DocState struct {
	fragState *Fragments
	failed    map[string]struct{}
	seen      map[string]struct{}
}

// NewDocState returns a DocState initialized with a FragState.
func NewDocState(f *Fragments) *DocState {
	return &DocState{
		fragState: f,
		failed:    map[string]struct{}{},
		seen:      map[string]struct{}{},
	}
}

// DocFragments returns a list of requested fragments.
func (d *DocState) DocFragments() []string {
	d.fragState.mut.RLock()
	defer d.fragState.mut.RUnlock()
	frags := make([]string, 0, len(d.seen))
	for frag := range d.seen {
		frags = append(frags, frag)
	}
	sort.Strings(frags)
	return frags
}

// UnresolvedFragments writes to stderr a list of fragment references that
// failed to resolve. Panics if any references failed.
func (d *DocState) UnresolvedFragments() {
	if len(d.failed) == 0 {
		return
	}
	refs := make([]string, 0, len(d.failed))
	for ref := range d.failed {
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
func (d *DocState) DocWith(fragref string, opt FragOptions) string {
	d.fragState.mut.Lock()
	defer d.fragState.mut.Unlock()
	d.seen[fragref] = struct{}{}
	node, _ := d.fragState.resolveNode(fragref, false)
	if node == nil {
		d.failed[fragref] = struct{}{}
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
func (d *DocState) Doc(fragref string) string {
	termWidth, _, _ := terminal.GetSize(int(os.Stdout.Fd()))
	return d.DocWith(fragref, FragOptions{
		Renderer: term.Renderer{Width: termWidth, TabSize: 4}.Render,
	})
}

// DocFlag returns the content for a flag by configuring the renderer with 0
// width, so that it can be properly formatted by the usage template.
func (d *DocState) DocFlag(fragref string) string {
	return d.DocWith(fragref, FragOptions{
		Renderer: term.Renderer{Width: 0, TabSize: 4}.Render,
	})
}
