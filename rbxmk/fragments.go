package main

import (
	"fmt"
	"os"
	"path"
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

func initDocs() drill.Node {
	termWidth, _, _ := terminal.GetSize(int(os.Stdout.Fd()))
	lang, ok := fragments.Languages[Language]
	if !ok {
		panic(fmt.Sprintf("unsupported language %q", Language))
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
		panic(fmt.Sprintf("unsupported language %q", Language))
	}
	return node
}

var docfs = initDocs()
var docMut sync.Mutex
var docCache = map[string]*markdown.Node{}

// docString extracts and formats the fragment from the given node. Panics if
// the node was not found.
func docString(fragpath string, node drill.Node) string {
	if node == nil {
		panic(fmt.Sprintf("unknown fragment {%s:%s}", Language, fragpath))
	}
	return strings.TrimSpace(node.Fragment())
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
