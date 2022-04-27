// The html package implements rendering content as HTML fragments.
package html

import (
	"fmt"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type Renderer struct {
	// Receives a heading name and returns an identifier to be used for the
	// heading.
	GenerateHeadingID func(name string) string

	ResolveLink func(link string) string
}

func NewRenderer() Renderer {
	return Renderer{}
}

func (r Renderer) Render(w io.Writer, s *goquery.Selection) error {
	// Ensure content is wrapped in body.
	if !s.Is("body") {
		s = s.WrapAllHtml("<body>")
	}

	// Remove gap text between sections.
	s.Find("section + section").Each(func(i int, s *goquery.Selection) {
		if n := s.Nodes[0]; n.PrevSibling.Type == html.TextNode {
			if strings.TrimSpace(n.PrevSibling.Data) == "" {
				n.Parent.RemoveChild(n.PrevSibling)
			}
		}
	})

	// Remove gap text after last section.
	s.Find("section > section:last-child").Each(func(i int, s *goquery.Selection) {
		if n := s.Nodes[0]; n.NextSibling.Type == html.TextNode {
			if strings.TrimSpace(n.NextSibling.Data) == "" {
				n.Parent.RemoveChild(n.NextSibling)
			}
		}
	})

	// Insert headings at start of each selection.
	for i := 6; i >= 1; i-- {
		s.Find("body" + strings.Repeat(" > section", i)).Each(func(_ int, s *goquery.Selection) {
			if name, ok := s.Attr("data-name"); ok {
				s.PrependHtml(fmt.Sprintf("<h%[1]d>%[2]s</h%[1]d>", i, name))
			}
		})
	}

	// Dissolve section elements.
	s.Find("section").Children().Unwrap()

	// Resolve heading IDs.
	if r.GenerateHeadingID != nil {
		s.Find("h1,h2,h3,h4,h5,h6").Each(func(i int, s *goquery.Selection) {
			s.SetAttr("id", r.GenerateHeadingID(s.Text()))
		})
	}

	// Resolve links.
	if r.ResolveLink != nil {
		s.Find("a[href]").Each(func(i int, s *goquery.Selection) {
			s.SetAttr("href", r.ResolveLink(s.AttrOr("href", "")))
		})
	}

	// Render contents of body.
	var err error
	s.Contents().EachWithBreak(func(i int, s *goquery.Selection) bool {
		if err = goquery.Render(w, s); err != nil {
			return false
		}
		return true
	})
	return err
}
