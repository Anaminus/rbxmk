// The term package implements rendering HTML content for the terminal.
package term

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/andybalholm/cascadia"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/net/html"
)

// Renderer implements a rendering HTML content in a stripped down format
// suitable for the terminal.
type Renderer struct {
	// Width sets the assumed width of the terminal, which aids in wrapping
	// text. If < 0, then the width is unbounded.
	//
	// If == 0, then the width is determined by the current terminal, or, if
	// ForOutput is true, then the width is 80.
	Width int
	// TabSize indicates the number of characters that the terminal uses to
	// render a tab. If <= 0, then tabs are assumed to have a size of 8.
	TabSize int
	// ForOutput indicates whether content is to be renderered for output,
	// rather than for display in the current terminal.
	ForOutput bool
	// UnorderedListFormat is the text used for unordered list markers.
	UnorderedListFormat string
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

// Change to debug inserted spaces.
const z = ' '

// Get base-10 digit-count of integer.
func intlen(n int) (c int) {
	for n != 0 {
		n /= 10
		c++
	}
	return c
}

// Contains the state of a list marker.
type listState struct {
	// Slice of numeric portion of marker.
	i, j int
	// Current numeric value of marker.
	num int
	// Current marker content.
	marker []byte
}

// Facilitates wrapped text.
type writer struct {
	bytes.Buffer
	w io.Writer

	// If wrapping, wrap to this width.
	width int

	// Width of tab character.
	tabWidth int

	// Level of indentation.
	indent int

	// String used for single level of indentation.
	indentString []byte

	// Contains list marker states.
	listStack []listState

	// Current leading space for each line. Includes extra spacing equal to list
	// marker width.
	leading []byte

	// Displayed width of leading space.
	leadingWidth int

	// Remaining number of newlines allowed within a margin. Resets by entering
	// a block.
	margin int

	// Whether a block was just entered. Prevents writing another margin if a
	// block was entered multiple times (e.g. nested blocks).
	block bool

	// Whether a marker was just inserted. Prevents leading space from be added
	// to the start of a paragraph.
	marked bool
}

// Calculate leading space inserted before lines.
func (w *writer) updateLeadingSpace() {
	ind := w.indent
	if ind <= 0 {
		ind = 0
	} else {
		// Indentation is used primarily for lists, so reduce indent by one for
		// markers.
		ind--
	}
	indent := len(w.indentString) * ind
	marker := 0
	if len(w.listStack) > 0 {
		marker = len(w.listStack[len(w.listStack)-1].marker)
	}
	if cap(w.leading) < indent+marker {
		w.leading = make([]byte, 0, indent+marker)
	} else {
		w.leading = w.leading[:0]
	}
	for i := 0; i < ind; i++ {
		w.leading = append(w.leading, w.indentString...)
	}
	for i := 0; i < marker; i++ {
		//TODO: Should match grapheme width.
		w.leading = append(w.leading, z)
	}
	w.leadingWidth = 0
	for _, c := range w.leading {
		if c == '\t' {
			w.leadingWidth += w.tabWidth
		} else {
			w.leadingWidth += 1
		}
	}
}

// Start inserting list markers. prefix is the content of the marker appearing
// before the index. max is the maximum index that will be displayed, while
// start is the starting index. suffix is the content appearing after the index.
// If max is <= 0, start is < 0, or max < start, then no index is displayed.
func (w *writer) List(prefix string, start, max int, suffix string) {
	var state listState
	numlen := intlen(max)
	n := len(prefix) + numlen + len(suffix)
	state.marker = make([]byte, n)
	b := state.marker[:0]
	b = append(b, prefix...)
	if max > 0 && start >= 0 && max >= start {
		state.num = start
		for i := 0; i < numlen; i++ {
			b = append(b, z)
		}
		state.i = len(prefix)
		state.j = state.i + numlen
	}
	b = append(b, suffix...)
	w.listStack = append(w.listStack, state)
	w.Indent(1)
	w.updateLeadingSpace()
}

// Insert a new marker. The marker is inserted immediately, ahead of the buffer.
// Every line except the first after a marker will include additional leading
// spacing equal to the length of the marker.
func (w *writer) ListMarker() {
	if len(w.listStack) == 0 {
		return
	}
	last := len(w.listStack) - 1
	state := w.listStack[last]
	if state.j > 0 {
		b := state.marker[state.i:state.j]
		b = b[len(b)-intlen(state.num):]
		b = b[:0]
		strconv.AppendInt(b, int64(state.num), 10)
		state.num++
		w.listStack[last] = state
	}
	w.w.Write(w.leading[:len(w.leading)-len(state.marker)])
	w.w.Write(state.marker)
	w.marked = true
}

// Stop rendering list markers and leading marker space.
func (w *writer) ClearList() {
	if len(w.listStack) == 0 {
		return
	}
	w.listStack = w.listStack[:len(w.listStack)-1]
	w.Indent(-1)
	w.updateLeadingSpace()
	if w.indent <= 0 {
		// Compensate for shortened li margin.
		w.margin++
	}
}

// Increase or decrease the level of indentation.
func (w *writer) Indent(delta int) {
	w.indent += delta
	w.updateLeadingSpace()
}

// Format accumulated paragraph text by wrapping.
func (w *writer) Flush() error {
	b := w.Buffer.Bytes()
	defer w.Buffer.Reset()

	b = bytes.TrimSpace(b)
	var paragraphs [][]byte
	for i := 0; i < len(b); {
		if b[i] == '\n' {
			// Bounds check unneeded; if b[i], being a \n, was at the end, it
			// would have been trimmed.
			if b[i+1] == '\n' {
				// New paragraph.
				if len(bytes.TrimSpace(b[:i])) > 0 {
					paragraphs = append(paragraphs, b[:i])
				}
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
	if len(bytes.TrimSpace(b)) > 0 {
		paragraphs = append(paragraphs, b)
	}

	if width := w.width; width > 0 {
		for j, p := range paragraphs {
			i := w.leadingWidth
			if i >= width {
				// Lost-cause. Keep as-is.
				//
				//TODO: Fallback to unwrapped (<=0 width).
				continue
			}
			s := bufio.NewScanner(bytes.NewBuffer(p))
			s.Split(bufio.ScanWords)
			var buf bytes.Buffer
			if !w.marked {
				// Insert leading space. i is already set.
				buf.Write(w.leading)
			} else {
				w.marked = false
			}
			var prevWord bool
			for s.Scan() {
				word := s.Bytes()
				// 1 added to represent preceding space.
				if i+len(word)+1 > width {
					// Append newline with leading space.
					buf.WriteByte('\n')
					buf.Write(w.leading)
					i = w.leadingWidth
					prevWord = false
				} else {
					// Append entire word.
					if prevWord {
						buf.WriteByte(z)
						i++
					}
					prevWord = true
					buf.Write(word)
					i += len(word)
					continue
				}
				// Append word in chunks of the available space.
				if prevWord {
					buf.WriteByte(z)
					i++
				}
				prevWord = true
				for len(word) > 0 {
					n := width - i
					if n == 0 {
						// Ensure at least one character is written.
						n = 1
					} else if n >= len(word) {
						// Leftover word fits.
						buf.Write(word)
						i += len(word)
						break
					}
					buf.Write(word[:n])
					buf.WriteByte('\n')
					buf.Write(w.leading)
					i = w.leadingWidth
					word = word[n:]
				}
			}
			paragraphs[j] = buf.Bytes()
		}
	} else {
		for j, p := range paragraphs {
			s := bufio.NewScanner(bytes.NewBuffer(p))
			s.Split(bufio.ScanWords)
			var buf bytes.Buffer
			for s.Scan() {
				if buf.Len() > 0 {
					buf.WriteByte(z)
				}
				buf.Write(s.Bytes())
			}
			paragraphs[j] = buf.Bytes()
		}
	}
	for i, p := range paragraphs {
		if i > 0 {
			w.Block(2)
		}
		if _, err := w.w.Write([]byte(string(p))); err != nil {
			return err
		}
		w.Margin()
	}
	return nil
}

// Margin inserts up to two newlines between blocks.
func (w *writer) Margin() {
	if w.margin > 0 {
		// Because margin is initialized to 0, it wont be written at the start
		// of the content.
		w.w.Write([]byte{'\n'})
		w.margin--
	}
	w.block = false
}

// Block adds a margin, then starts a new block, which resets the available
// margins.
func (w *writer) Block(margin int) {
	if !w.block {
		w.Margin()
	}
	w.margin = margin
	w.block = true
}

var sectionCounter = cascadia.MustCompile("body > section")

// Elements that have margins inserted before and after. Value indicates the
// maximum allowed margin size after the block.
var block = map[string]int{
	"p":       2,
	"ol":      2,
	"ul":      2,
	"pre":     2,
	"table":   2,
	"section": 2,
	"li":      1,
}

func (r Renderer) Render(w io.Writer, s *goquery.Selection) error {
	buf := &writer{
		w:            w,
		width:        r.Width,
		tabWidth:     r.TabSize,
		indentString: []byte{'\t'},
	}
	if buf.tabWidth <= 0 {
		//TODO: Get from terminal somehow. 8 is the default most of the time,
		//but this can be adjusted by the user, so this value wont always be
		//correct. However, it at least wont overflow if the actual tab width is
		//less than the value.
		buf.tabWidth = 8
	}
	if buf.width == 0 {
		if r.ForOutput {
			buf.width = 80
		} else {
			buf.width, _, _ = terminal.GetSize(int(os.Stdout.Fd()))
			if runtime.GOOS == "windows" && buf.width > 1 {
				// If a newline occurs on the edge, cmd.exe will incorrectly
				// wrap it around, producing a gap. Reducing the width by one
				// ensures this gap wont appear.
				buf.width--
			}
		}
	}
	isRoot := s.Is("body")
	state := walkState{renderer: r}
	for _, node := range s.Nodes {
		if isRoot {
			// If an entire document is received, increase the depth to force
			// the top-level section names to be rendered.
			state.depth++
		}
		err := walk(node, func(node *html.Node, entering bool) error {
			switch node.Type {
			case html.ErrorNode:
			case html.TextNode:
				if entering {
					if strings.TrimSpace(node.Data) != "" {
						buf.WriteString(node.Data)
					}
				}
			case html.DocumentNode:
			case html.ElementNode:
				if m := block[node.Data]; m != 0 && entering {
					buf.Flush()
					buf.Block(m)
				}
				h := handlers[elementMatcher{node.Data, entering}]
				var err error
				if h != nil {
					err = h(buf, node, &state)
				}
				if block[node.Data] != 0 && !entering {
					buf.Flush()
					buf.Margin()
				}
				return err
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
	renderer Renderer
	depth    int
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

			var buf bytes.Buffer
			walk(node, func(node *html.Node, entering bool) error {
				if entering && node.Type == html.TextNode {
					buf.WriteString(node.Data)
				}
				return nil
			})
			b := bytes.TrimSpace(buf.Bytes())

			ind := 1
			if w.indent <= 0 {
				// Compensate for reduced list indentation.
				ind = 2
			}
			w.Indent(ind)
			c := make([]byte, len(w.leading)+1)
			c[0] = '\n'
			copy(c[1:], w.leading)
			b = bytes.ReplaceAll(b, c[:1], c)
			w.w.Write(w.leading)
			w.w.Write(b)
			w.Indent(-ind)

			return skipChildren
		} else {
			// Inline

			// TODO: Content must be treated as one word.
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
				w.Block(2)
				w.WriteString(strings.Repeat("#", s.depth))
				w.WriteString(" ")
				w.WriteString(name)
				w.Flush()
			}
		}
		s.depth++
		return nil
	},
	{"section", false}: func(w *writer, node *html.Node, s *walkState) error {
		s.depth--
		return nil
	},
	{"ul", true}: func(w *writer, node *html.Node, s *walkState) error {
		prefix := s.renderer.UnorderedListFormat
		if prefix == "" {
			prefix = " - "
		}
		w.List(" - ", 0, 0, "")
		return nil
	},
	{"ul", false}: func(w *writer, node *html.Node, s *walkState) error {
		w.ClearList()
		return nil
	},
	{"ol", true}: func(w *writer, node *html.Node, s *walkState) error {
		n := 0
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode && c.Data == "li" {
				n++
			}
		}
		w.List(" ", 1, n, ". ")
		return nil
	},
	{"ol", false}: func(w *writer, node *html.Node, s *walkState) error {
		w.ClearList()
		return nil
	},
	{"li", true}: func(w *writer, node *html.Node, s *walkState) error {
		w.ListMarker()
		return nil
	},
	{"li", false}: func(w *writer, node *html.Node, s *walkState) error {
		w.Flush()
		return nil
	},
	{"table", true}: func(w *writer, node *html.Node, s *walkState) error {
		w.w.Write([]byte("<TODO:table>"))
		return skipChildren
	},
	{"table", false}: func(w *writer, node *html.Node, s *walkState) error {
		w.w.Write([]byte("</TODO:table>"))
		return nil
	},
}
