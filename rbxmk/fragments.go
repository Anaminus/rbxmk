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

// docString extracts and formats the fragment from the given node. Panics if
// the node was not found.
func docString(fragref string, node drill.Node) string {
	if node == nil {
		docFailed[fragref] = struct{}{}
		return "{" + Language + ":" + fragref + "}"
	}
	return strings.TrimSpace(node.Fragment())
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
func parseFragRef(s, suffix string, filesep rune) []string {
	if s == "" {
		return []string{}
	}
	i := strings.IndexRune(s, filesep)
	if i < 0 {
		return strings.Split(strings.ToLower(s)+suffix, "/")
	}
	items := make([]string, 0, strings.Count(s, "/")+1)
	items = append(items, strings.Split(strings.ToLower(s[:i])+suffix, "/")...)
	items = append(items, strings.Split(s[i+1:], "/")...)
	return items
}

// ResolveFragment returns the content of the fragment referred to by fragref.
func ResolveFragment(fragref string) string {
	docMut.Lock()
	defer docMut.Unlock()
	return resolveFragment(fragref)
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
	return resolveFragment(fragref)
}

const filesep = ':'
const suffix = ".md"

func resolveFragment(fragref string) string {
	var n drill.Node = docfs
	var path string
	names := parseFragRef(fragref, suffix, filesep)
	for _, name := range names {
		if name == "" {
			return docString(fragref, nil)
		}
		path += "/" + name
		if node, ok := docCache[path]; ok {
			n = node
		} else {
			switch v := n.(type) {
			case drill.UnorderedBranch:
				n = v.UnorderedChild(name)
			default:
				return docString(fragref, nil)
			}
			if node, ok := n.(*markdown.Node); ok {
				docCache[path] = node
			}
		}
	}
	return docString(fragref, n)
}
