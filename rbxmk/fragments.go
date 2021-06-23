package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/anaminus/drill"
	"github.com/anaminus/drill/filesys"
	"github.com/anaminus/drill/filesys/markdown"
	"github.com/anaminus/rbxmk/fragments"
	"github.com/anaminus/rbxmk/rbxmk/term"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
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
		{Pattern: "*.md", Func: markdown.NewHandler(
			goldmark.WithRenderer(term.Renderer{Width: termWidth, TabSize: 4}),
			goldmark.WithExtensions(
				extension.Table,
				extension.Footnote,
			),
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
var docCache = map[string]*markdown.Node{}
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

// Doc returns the content of the fragment referred to by fragref. The given
// path is marked to be returned by DocFragments.
//
// Doc should only be used to capture additional fragment references.
// ResolveFragment can be used to resolve a reference without marking it.
func Doc(fragref string) string {
	docMut.Lock()
	defer docMut.Unlock()
	docSeen[fragref] = struct{}{}
	node, _ := resolveFragmentNode(fragref, false)
	if node == nil {
		docFailed[fragref] = struct{}{}
		return "{" + Language + ":" + fragref + "}"
	}
	return strings.TrimSpace(node.Fragment())
}

// ResolveFragment returns the content of the fragment referred to by fragref.
func ResolveFragment(fragref string) string {
	docMut.Lock()
	defer docMut.Unlock()
	node, _ := resolveFragmentNode(fragref, false)
	if node == nil {
		return ""
	}
	return strings.TrimSpace(node.Fragment())
}

const FragSep = ':'
const FragSuffix = ".md"

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
			if node, ok := n.(*markdown.Node); ok {
				docCache[path] = node
			}
		}
	}
	return n, infile
}
