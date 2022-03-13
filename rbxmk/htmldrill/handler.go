// The htmldrill package implements a drill/filesys.Handler for the HTML format.
// Drill nodes are delimited by "section" elements. Unordered nodes are
// implemented as section elements with the "data-name" attribute.
package htmldrill

import (
	"bytes"
	"io"
	"io/fs"

	"github.com/PuerkitoBio/goquery"
	"github.com/anaminus/drill"
	"github.com/anaminus/drill/filesys"
	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

// Renderer renders a selection, writing the result to w.
type Renderer func(w io.Writer, s *goquery.Selection) error

// config holds options
type config struct {
	parseOptions []html.ParseOption
	renderer     Renderer
}

// Option configures the package's filesys handler.
type Option func(*config)

// WithParseOptions configures how a handler parses HTML.
func WithParseOptions(opts ...html.ParseOption) Option {
	return func(cfg *config) {
		cfg.parseOptions = append(cfg.parseOptions, opts...)
	}
}

// WithRenderer configures a handler to return nodes that use the given
// renderer.
func WithRenderer(renderer Renderer) Option {
	return func(cfg *config) {
		cfg.renderer = renderer
	}
}

// NewHandler returns a filesys.HandlerFunc that parses a file as an HTML
// document.
func NewHandler(opts ...Option) filesys.HandlerFunc {
	var cfg config
	for _, opt := range opts {
		opt(&cfg)
	}
	return func(fsys fs.FS, name string) drill.Node {
		b, err := fs.ReadFile(fsys, name)
		if err != nil {
			return nil
		}
		root, err := html.ParseWithOptions(bytes.NewBuffer(b), cfg.parseOptions...)
		if err != nil {
			return nil
		}
		doc := goquery.NewDocumentFromNode(root)
		StripComments(doc.Selection)
		FixTemplateDirectives(doc.Selection)
		return NewNode(doc, cfg.renderer)
	}
}

const (
	sectionName = "section"   // Section element name.
	sectionAttr = "data-name" // Section element attribute name.
)

var bodyMatcher = goquery.SingleMatcher(cascadia.MustCompile("body"))
var sectionMatcher = cascadia.MustCompile(sectionName)
var usectionMatcher = cascadia.MustCompile(sectionName + "[" + sectionAttr + "]")

// Node implements drill.Node.
type Node struct {
	document  *goquery.Document
	selection *goquery.Selection
	renderer  Renderer
}

// NewNode returns a Node that wraps the given goquery.Document and optional
// renderer.
func NewNode(doc *goquery.Document, renderer Renderer) *Node {
	return &Node{
		document:  doc,
		selection: doc.FindMatcher(bodyMatcher).First(),
		renderer:  renderer,
	}
}

// derive returns a Node that wraps selection instead of the selection from n.
func (n *Node) derive(selection *goquery.Selection) *Node {
	d := *n
	d.selection = selection
	return &d
}

// Document returns the root goquery.Document.
func (n *Node) Document() *goquery.Document {
	return n.document
}

// Selection returns the wrapped goquery.Selection.
func (n *Node) Selection() *goquery.Selection {
	return n.selection
}

func (n *Node) ochildren() *goquery.Selection {
	return n.selection.ChildrenMatcher(sectionMatcher)
}

func (n *Node) uchildren() *goquery.Selection {
	return n.selection.ChildrenMatcher(usectionMatcher)
}

// Fragment renders the wrapped node as a string using the node's renderer, or
// goquery.Render if the node has no renderer. Returns an empty string if an
// error occurs.
func (n *Node) Fragment() string {
	renderer := n.renderer
	if renderer == nil {
		renderer = goquery.Render
	}
	var buf bytes.Buffer
	if err := renderer(&buf, n.selection); err != nil {
		return ""
	}
	return buf.String()
}

// render renders n to w.
func render(n *Node, w *io.PipeWriter) {
	renderer := n.renderer
	if renderer == nil {
		renderer = goquery.Render
	}
	if err := renderer(w, n.selection); err != nil {
		w.CloseWithError(err)
	}
	w.Close()
}

// FragmentReader returns a ReadCloser that renders the wrapped node according
// the node's renderer, or goquery.Render if the node has no renderer. Errors
// are passed to the reader.
func (n *Node) FragmentReader() (r io.ReadCloser, err error) {
	r, w := io.Pipe()
	go render(n, w)
	return r, nil
}

// WithRenderer returns a copy of the node that uses the given renderer. r may
// be nil.
func (n *Node) WithRenderer(r Renderer) *Node {
	d := *n
	d.renderer = r
	return &d
}

// Len returns the number of child section elements.
func (n *Node) Len() int {
	return n.ochildren().Length()
}

func (n *Node) orderedChild(i int) *Node {
	children := n.ochildren()
	if i = drill.Index(i, children.Length()); i < 0 {
		return nil
	}
	return n.derive(children.Eq(i))
}

// OrderedChild returns a Node that wraps the ordered child section element at
// index i. Returns nil if the index is out of bounds.
func (n *Node) OrderedChild(i int) drill.Node {
	node := n.orderedChild(i)
	if node == nil {
		return nil
	}
	return node
}

// OrderedChildren returns a list of Nodes that wrap each ordered child Section.
func (n *Node) OrderedChildren() []drill.Node {
	var sections []drill.Node
	n.ochildren().Each(func(i int, s *goquery.Selection) {
		sections = append(sections, n.derive(s))
	})
	return sections
}

func (n *Node) unorderedChild(name string) *Node {
	var selection *goquery.Selection
	n.uchildren().EachWithBreak(func(i int, s *goquery.Selection) bool {
		if value, ok := s.Attr(sectionAttr); !ok || value != name {
			return true
		}
		selection = s
		return false
	})
	if selection == nil {
		return nil
	}
	return n.derive(selection)
}

// UnorderedChild returns a Node that wraps the unordered child section element
// whose data-name attribute is equal to name.
func (n *Node) UnorderedChild(name string) drill.Node {
	node := n.unorderedChild(name)
	if node == nil {
		return nil
	}
	return node
}

// UnorderedChildren returns a map of names to Nodes that wrap each unordered
// child section element.
func (n *Node) UnorderedChildren() map[string]drill.Node {
	children := n.uchildren()
	sections := make(map[string]drill.Node, children.Length())
	children.Each(func(i int, s *goquery.Selection) {
		if value, ok := s.Attr(sectionAttr); ok {
			sections[value] = n.derive(s)
		}
	})
	return sections
}

// Descend recursively descends into the unordered child section elements
// matching each given name. Returns nil if a child could not be found at any
// point.
func (n *Node) Descend(names ...string) drill.Node {
	for _, name := range names {
		node := n.unorderedChild(name)
		if node == nil {
			return nil
		}
		n = node
	}
	return n
}

// Query recursively descends into the child nodes that match the given queries.
// A query is either a string or an int. If an int, then the next node is
// acquired using the OrderedChild method of the current node. If a string, then
// the next node is acquired using the UnorderedChild method of the current
// node. Returns nil if a child could not be found at any point.
func (n *Node) Query(queries ...interface{}) drill.Node {
	for _, query := range queries {
		switch q := query.(type) {
		case string:
			node := n.unorderedChild(q)
			if node == nil {
				return nil
			}
			n = node
		case int:
			node := n.orderedChild(q)
			if node == nil {
				return nil
			}
			n = node
		default:
			return nil
		}
	}
	return n
}
