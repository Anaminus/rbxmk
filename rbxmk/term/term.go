// The term package implements rendering HTML content for the terminal.
package term

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

// Renderer implements a goldmark Renderer by rendering a stripped down format
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

func (r Renderer) Render(w io.Writer, s *goquery.Selection) error {
	return goquery.Render(w, s)
}
