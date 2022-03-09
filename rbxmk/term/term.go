// The term package implements rendering HTML content for the terminal.
package term

import (
	"bytes"
	"errors"
	"io"
	"runtime"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

// Renderer implements a rendering HTML content in a stripped down format
// suitable for the terminal.
type Renderer struct {
	// Width sets the assumed width of the terminal, which aids in wrapping
	// text. If == 0, then the width is treated as 80. If < 0, then the width is
	// treated as unbounded.
	Width int
	// TabSize indicates the number of characters that the terminal uses to
	// render a tab. If <= 0, then tabs are assumed to have a size of 8.
	TabSize int
}

var (
	stop         = errors.New("stop")          // Stop walking without erroring.
	skipChildren = errors.New("skip children") // Skip all child nodes.
	skipText     = errors.New("skip text")     // Skip child text nodes.
)

type walker func(node *html.Node, entering bool) error

func walk(node *html.Node, cb walker) error {
	err := cb(node, true)
	if err != nil && err != skipChildren && err != skipText {
		return err
	}
	if err != skipChildren {
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode && err == skipText {
				continue
			}
			if err := walk(c, cb); err != nil {
				return err
			}
		}
	}
	if err = cb(node, false); err != nil && err != skipChildren && err != skipText {
		return err
	}
	return err
}

// Facilitates wrapped text.
type writer struct {
	bytes.Buffer
	w io.Writer

	// If wrapping, wrap to this width.
	width int

	// Size of tab character.
	tabSize int
}

func (w *writer) Flush() error {
	b := w.Buffer.Bytes()
	defer w.Buffer.Reset()

	b = bytes.TrimSpace(b)
	var paragraphs [][]rune
	for i := 0; i < len(b); {
		if b[i] == '\n' {
			// Bounds check unneeded; if b[i], being a \n, was at the end, it
			// would have been trimmed.
			if b[i+1] == '\n' {
				// New paragraph.
				paragraphs = append(paragraphs, []rune(string(b[:i])))
				i += 2
				for b[i] == '\n' {
					// Collapse extra newlines.
					i++
				}
				b = b[i:]
				i = 0
				continue
			} else {
				// Unwrap.
				b[i] = ' '
			}
		}
		i++
	}
	paragraphs = append(paragraphs, []rune(string(b)))
	if w.width > 0 {
		for j, p := range paragraphs {
			for i := 0; i < len(p); {
				n := i + w.width
				if n+1 >= len(p) {
					break
				}
				for n > i && !unicode.IsSpace(p[n]) {
					n--
				}
				if n <= i {
					// Long word.
					n = i + w.width
					p = append(p, 0)
					copy(p[n+1:], p[n:])
					p[n] = '\n'
					i = n + 1
				} else {
					p[n] = '\n'
					i = n + 1
				}
			}
			paragraphs[j] = p
		}
		if runtime.GOOS == "windows" {
			// Remove newlines that are inserted just after the edge of the
			// width. This prevents the terminal from incorrectly producing an
			// extra gap when wrapping. This removes the newline entirely, so it
			// is assumed that the wrapping will separate the previous word from
			// the next word.
			for i, p := range paragraphs {
				for h, i := 0, 0; i < len(p); i++ {
					if p[i] != '\n' {
						continue
					}
					if i-h >= w.width {
						copy(p[i:], p[i+1:])
						p = p[:len(p)-1]
						h = i
						i--
					} else {
						h = i + 1
					}
				}
				paragraphs[i] = p
			}
		}
	}
	for i, p := range paragraphs {
		if i > 0 {
			w.w.Write([]byte{'\n', '\n'})
		}
		if _, err := w.w.Write([]byte(string(p))); err != nil {
			return err
		}
	}
	return nil
}

var sectionCounter = cascadia.MustCompile("body > section")

func (r Renderer) Render(w io.Writer, s *goquery.Selection) error {
	buf := &writer{
		w:       w,
		width:   r.Width,
		tabSize: r.TabSize,
	}
	var state walkState
	for _, node := range s.Nodes {
		if node.Type == html.TextNode {
			continue
		}
		// If the body contains more than one section, increase the depth to
		// force the section names to be rendered.
		isRoot := s.FindMatcher(sectionCounter).Length() > 1
		if isRoot {
			state.depth++
		}
		err := walk(node, func(node *html.Node, entering bool) error {
			switch node.Type {
			case html.ErrorNode:
			case html.TextNode:
				if entering {
					buf.WriteString(node.Data)
				}
			case html.DocumentNode:
			case html.ElementNode:
				h := handlers[elementMatcher{node.Data, entering}]
				if h != nil {
					return h(buf, node, &state)
				}
			case html.CommentNode:
			case html.DoctypeNode:
			case html.RawNode:
			}
			return nil
		})
		if err != nil && err != stop {
			return err
		}
		if isRoot {
			state.depth--
		}
	}
	return buf.Flush()
}

type elementMatcher struct {
	data     string
	entering bool
}

type walkState struct {
	depth int
}

type nodeHandlers map[elementMatcher]func(w *writer, node *html.Node, s *walkState) error

func isElement(node *html.Node, tag string) bool {
	return node.Type == html.ElementNode && node.Data == tag
}

func sectionName(node *html.Node) string {
	for _, attr := range node.Attr {
		if attr.Key == "data-name" {
			return attr.Val
		}
	}
	return ""
}

var handlers = nodeHandlers{
	{"code", true}: func(w *writer, node *html.Node, s *walkState) error {
		if isElement(node.Parent, "pre") { // May have syntax: `class="language-*"`
			// Block
			w.WriteString(node.FirstChild.Data)
			return skipChildren
		} else {
			// Inline
			return w.WriteByte('`')
		}
	},
	{"code", false}: func(w *writer, node *html.Node, s *walkState) error {
		if isElement(node.Parent, "pre") {
			// Block
		} else {
			// Inline
			return w.WriteByte('`')
		}
		return nil
	},
	{"section", true}: func(w *writer, node *html.Node, s *walkState) error {
		if s.depth > 0 {
			if name := sectionName(node); name != "" {
				w.WriteString(strings.Repeat("#", s.depth))
				w.WriteString(" ")
				w.WriteString(name)
				w.WriteString("\n\n")
			}
		}
		s.depth++
		return skipText
	},
	{"section", false}: func(w *writer, node *html.Node, s *walkState) error {
		s.depth--
		return nil
	},
	{"p", true}: func(w *writer, node *html.Node, s *walkState) error {
		_, err := w.WriteString("\n\n")
		return err
	},
	{"p", false}: func(w *writer, node *html.Node, s *walkState) error {
		_, err := w.WriteString("\n\n")
		return err
	},
}
