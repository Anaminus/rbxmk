package htmldrill

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func cut(s, sep string) (before, after string, found bool) {
	if i := strings.Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}

// FixTemplateDirectives splits TextNodes by template directives. Directives are
// inserted as RawNodes. Directives are assumed not to span across multiple
// TextNodes.
func FixTemplateDirectives(s *goquery.Selection) {
	nodes := make([]*html.Node, len(s.Nodes))
	copy(nodes, s.Nodes)
	var textNodes []*html.Node
	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			switch c.Type {
			case html.TextNode:
				textNodes = append(textNodes, c)
			case html.RawNode:
			default:
				nodes = append(nodes, c)
			}
		}
	}
	for _, text := range textNodes {
		for {
			prev, middle, ok := cut(text.Data, "{{")
			if !ok {
				break
			}
			middle, next, ok := cut(middle, "}}")
			if !ok {
				break
			}

			// Reuse current node to contain previous content.
			text.Data = prev

			// Insert RawNode containing directive.
			directive := &html.Node{Type: html.RawNode, Data: "{{" + middle + "}}"}
			text.Parent.InsertBefore(directive, text.NextSibling)

			// Insert next content.
			after := &html.Node{Type: html.TextNode, Data: next}
			directive.Parent.InsertBefore(after, directive.NextSibling)

			// Look for next directive.
			text = after
		}
	}
}
