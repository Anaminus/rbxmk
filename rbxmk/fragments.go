package main

import (
	"fmt"
	"os"
	"path"
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
var docMut sync.Mutex
var docCache = map[string]*markdown.Node{}
var docFails = map[string]struct{}{}

// docString extracts and formats the fragment from the given node. Panics if
// the node was not found.
func docString(fragpath string, node drill.Node) string {
	if node == nil {
		docFails[fragpath] = struct{}{}
		return fmt.Sprintf("{%s:%s}", Language, fragpath)
	}
	return strings.TrimSpace(node.Fragment())
}

// UnresolvedFragments writes to stderr a list of fragment references that
// failed to resolve. Panics if any references failed.
func UnresolvedFragments() {
	if len(docFails) == 0 {
		return
	}
	refs := make([]string, 0, len(docFails))
	for ref := range docFails {
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

func Doc(fragpath string) string {
	docMut.Lock()
	defer docMut.Unlock()
	var n drill.Node = docfs
	var p string
	names := strings.Split(fragpath, "/")
	for _, name := range names {
		if name == "" {
			return docString(fragpath, nil)
		}
		p = path.Join(p, name)
		if node, ok := docCache[p]; ok {
			n = node
		} else {
			switch v := n.(type) {
			case drill.UnorderedBranch:
				n = v.UnorderedChild(name)
			default:
				return docString(fragpath, nil)
			}
			if node, ok := n.(*markdown.Node); ok {
				docCache[p] = node
			}
		}
	}
	return docString(fragpath, n)
}
