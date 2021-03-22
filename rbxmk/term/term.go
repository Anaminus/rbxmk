// The term package implements a goldmark renderer for rendering Markdown in the
// terminal.
package term

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"strings"
	"unicode"

	"github.com/anaminus/drill/filesys/markdown"
	"github.com/bbrks/wrap/v2"
	"github.com/olekukonko/tablewriter"
	"github.com/yuin/goldmark/ast"
	east "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// textOf returns the text of a node without escaped punctuation.
func textOf(n ast.Node, source []byte) string {
	return string(util.UnescapePunctuations(n.Text(source)))
}

// isPrefix returns whether r should be treated as indentation.
func isPrefix(r rune) bool {
	return unicode.IsSpace(r) || unicode.IsDigit(r) || r == '-' || r == '.'
}

// writer writes formatted content.
type writer struct {
	w    *bufio.Writer
	wrap wrap.Wrapper

	maxWidth int // Max line width.
	tabSize  int // Size of tabs.

	line []rune // Accumulated line content.

	nolead bool   // True if no writing indentation.
	indent []rune // Accumulated indentation, normalized to spacing.

	newlines int // Number of consecutive newlines written.
}

// Raw is called after writing raw content to the underlying buffer. This resets
// all line information.
func (w *writer) Raw() {
	w.line = w.line[:0]
	w.indent = w.indent[:0]
	w.nolead = false
	w.newlines = 0
}

func (w *writer) Write(p []byte) (n int, err error) {
	for _, r := range string(p) {
		switch r {
		case '\n':
			// Wrap content to width.
			w.wrap.OutputLinePrefix = string(w.indent)
			s := w.wrap.Wrap(string(w.line), w.maxWidth)
			s = strings.TrimPrefix(s, w.wrap.OutputLinePrefix)
			if w.newlines < 2 {
				// Only allow up to two consecutive newlines.
				w.w.WriteString(s)
			}
			w.line = w.line[:0]
			w.indent = w.indent[:0]
			w.nolead = false
			w.newlines++
		case '\t':
			w.newlines = 0
			// Replace tab with spaces snapped to the nearest tab stop.
			n := (len(w.line)/w.tabSize+1)*w.tabSize - len(w.line)
			for i := 0; i < n; i++ {
				w.line = append(w.line, ' ')
				if !w.nolead {
					w.indent = append(w.indent, ' ')
				}
			}
		default:
			w.newlines = 0
			if !w.nolead && isPrefix(r) {
				if unicode.IsSpace(r) {
					// Accumulate directly.
					w.indent = append(w.indent, r)
				} else {
					// Accumulate as a space.
					w.indent = append(w.indent, ' ')
				}
			} else {
				// End of indentation.
				w.nolead = true
			}
			w.line = append(w.line, r)
		}
	}
	return len(p), nil
}

func (w *writer) WriteString(s string) (n int, err error) {
	return w.Write([]byte(s))
}

// Renderer implements a goldmark Renderer by rendering a stripped down format
// suitable for the terminal.
type Renderer struct {
	// Width sets the assumed width of the terminal, which aids in wrapping
	// text. If <= 0, then the width is treated as unbounded.
	Width int
	// TabSize indicates the number of characters that the terminal uses to
	// render a tab. If <= 0, then tabs are assumed to have a size of 8.
	TabSize int
}

func (r Renderer) Render(w io.Writer, source []byte, n ast.Node) error {
	buf := writer{
		w: bufio.NewWriter(w),
		wrap: wrap.Wrapper{
			Breakpoints:               " -",
			Newline:                   "\n",
			LimitIncludesPrefixSuffix: true,
			CutLongWords:              true,
		},
		tabSize:  r.TabSize,
		maxWidth: r.Width,
	}
	if buf.maxWidth <= 0 {
		buf.maxWidth = -1
	}
	if buf.tabSize <= 0 {
		buf.tabSize = 8
	}
	indent := 0
	ast.Walk(n, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			switch n := n.(type) {
			case *ast.Document:
			case *markdown.Section:
			case *ast.Heading:
				buf.WriteString(strings.Repeat("#", n.Level))
				buf.WriteString(" ")
			case *ast.Text:
				buf.WriteString(textOf(n, source))
			case *ast.Paragraph:
			case *ast.List:
			case *ast.ListItem:
				if !n.HasBlankPreviousLines() {
					buf.WriteString("\n")
				}
				buf.WriteString(strings.Repeat(" ", buf.tabSize*indent))
				list := n.Parent().(*ast.List)
				if list.IsOrdered() {
					i := 1
					for n := n.PreviousSibling(); n != nil; n = n.PreviousSibling() {
						i++
					}
					buf.WriteString(strconv.Itoa(i))
					buf.WriteString(". ")
				} else {
					buf.WriteString("- ")
				}
				indent++
			case *ast.Emphasis:
				if n.Level == 1 {
					buf.WriteString("*")
				}
			case *ast.CodeBlock:
				buf.WriteString("\n")
				buf.WriteString("\t")
				lines := n.Lines()
				prev := false
				for i := 0; i < lines.Len(); i++ {
					line := lines.At(i)
					s := bytes.TrimSpace(line.Value(source))
					if prev {
						buf.WriteString(" ")
					}
					prev = len(s) > 0
					buf.Write(s)
				}
			case *ast.FencedCodeBlock:
				buf.WriteString("\n")
				lines := n.Lines()
				for i := 0; i < lines.Len(); i++ {
					line := lines.At(i)
					s := line.Value(source)
					if len(bytes.TrimSpace(s)) > 0 {
						buf.WriteString("\t")
					}
					buf.Write(s)
				}
			case *east.Table:
				// Prepare for tablewriter.
				table := make([][]string, 0, n.ChildCount())
				var headers []string
				for row := n.FirstChild(); row != nil; row = row.NextSibling() {
					cols := make([]string, 0, row.ChildCount())
					for col := row.FirstChild(); col != nil; col = col.NextSibling() {
						cell := strings.TrimSpace(textOf(col, source))
						cols = append(cols, cell)
					}
					if _, ok := row.(*east.TableHeader); ok {
						headers = cols
					} else {
						table = append(table, cols)
					}
				}

				tw := tablewriter.NewWriter(buf.w)
				cols := make([]int, len(n.Alignments))
				for i, a := range n.Alignments {
					switch a {
					case east.AlignLeft:
						cols[i] = tablewriter.ALIGN_LEFT
					case east.AlignRight:
						cols[i] = tablewriter.ALIGN_RIGHT
					case east.AlignCenter:
						cols[i] = tablewriter.ALIGN_CENTER
					case east.AlignNone:
						cols[i] = tablewriter.ALIGN_DEFAULT
					}
				}
				tw.SetColumnAlignment(cols)
				tw.SetHeader(headers)
				tw.SetHeaderLine(true)
				tw.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
				tw.SetAutoFormatHeaders(false)
				tw.SetBorders(tablewriter.Border{Left: true, Right: true})
				tw.SetColumnSeparator("")
				tw.SetCenterSeparator("-")
				tw.SetColWidth((buf.maxWidth - 3 - 2*len(cols)) / len(cols))
				tw.AppendBulk(table)
				tw.Render()
				buf.Raw()
				buf.WriteString("\n")
				return ast.WalkSkipChildren, nil
			case *ast.CodeSpan:
				buf.WriteString("`")
			default:
				_ = n
			}
		} else {
			switch n := n.(type) {
			case *ast.Document:
			case *markdown.Section:
			case *ast.Heading:
				buf.WriteString("\n")
			case *ast.Text:
				if n.SoftLineBreak() {
					buf.WriteString(" ")
				} else if n.HardLineBreak() {
					buf.WriteString("\n")
				}
			case *ast.Paragraph:
				buf.WriteString("\n\n")
			case *ast.List:
				if n.Parent().Kind() != ast.KindListItem {
					buf.WriteString("\n\n")
				}
			case *ast.ListItem:
				indent--
			case *ast.Emphasis:
				if n.Level == 1 {
					buf.WriteString("*")
				}
			case *ast.CodeBlock:
				buf.WriteString("\n\n")
			case *ast.FencedCodeBlock:
				buf.WriteString("\n\n")
			case *east.Table:
			case *ast.CodeSpan:
				buf.WriteString("`")
			default:
				_ = n
			}
		}
		return ast.WalkContinue, nil
	})
	return buf.w.Flush()
}

func (Renderer) AddOptions(...renderer.Option) {}
