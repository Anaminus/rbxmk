package htmldrill

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// StripComments removes comment nodes from the selection.
func StripComments(s *goquery.Selection) {
	s.FilterFunction(func(i int, s *goquery.Selection) bool {
		return s.Nodes[0].Type == html.CommentNode
	}).Remove()
}
