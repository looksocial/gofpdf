package table

import (
	"fmt"
)

// shouldFillRow returns true if current row should be filled (for zebra striping)
func (t *Table) shouldFillRow() bool {
	if t.RowStyle.Alternating {
		// Alternating based on logical row index, which is independent of page layout
		// Odd-indexed rows (1, 3, 5, ...) are filled
		return t.currentRow%2 == 1
	}
	return false
}

// getAlignStr returns alignment string
func (t *Table) getAlignStr(align string) string {
	switch align {
	case "C", "Center":
		return "C"
	case "R", "Right":
		return "R"
	default:
		return "L"
	}
}

// applyCellStyle applies cell styling
func (t *Table) applyCellStyle(style CellStyle) {
	// Get current font info using GetFontSize
	ptSize, _ := t.pdf.GetFontSize()
	fontSize := ptSize

	// Build font style
	fontStyle := ""
	if style.Bold {
		fontStyle = "B"
	}
	if style.Italic {
		if fontStyle == "" {
			fontStyle = "I"
		} else {
			fontStyle += "I"
		}
	}

	if style.FontSize > 0 {
		fontSize = style.FontSize
	}

	// Set font - use empty string to keep current font family
	t.pdf.SetFont("", fontStyle, fontSize)

	// Set colors
	if len(style.FillColor) >= 3 {
		t.pdf.SetFillColor(style.FillColor[0], style.FillColor[1], style.FillColor[2])
	}
	// Always set text color - if not specified in style, use black (0,0,0)
	if len(style.TextColor) >= 3 {
		t.pdf.SetTextColor(style.TextColor[0], style.TextColor[1], style.TextColor[2])
	} else {
		// Default to black text if not specified
		t.pdf.SetTextColor(0, 0, 0)
	}
}

// getRowHeight returns the row height
func (t *Table) getRowHeight() float64 {
	if t.RowHeight > 0 {
		return t.RowHeight
	}
	// Get default line height based on font
	ptSize, _ := t.pdf.GetFontSize()
	return ptSize * 0.4 // Default multiplier
}

// valueToString converts value to string
func (t *Table) valueToString(val interface{}) string {
	if val == nil {
		return ""
	}

	switch v := val.(type) {
	case string:
		return v
	case fmt.Stringer:
		return v.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}

// containsRune checks if string contains rune
func containsRune(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}

// checkPageBreak checks if we need to add a page break before rendering the next row
// Returns true if a page break was added
func (t *Table) checkPageBreak(requiredHeight float64) bool {
	if !t.PageBreakMode {
		return false
	}

	// Get current position and page dimensions
	currentY := t.pdf.GetY()
	_, pageHeight := t.pdf.GetPageSize()
	_, _, _, bottomMargin := t.pdf.GetMargins()

	// Calculate available space on current page
	availableHeight := pageHeight - bottomMargin - t.PageBreakMargin - currentY

	// Check if we need a page break
	if requiredHeight > availableHeight {
		// Add new page
		t.pdf.AddPage()

		// Repeat header if enabled - check for custom header callback first
		if t.CustomRepeatHeader != nil {
			// Call custom header function which returns new Y position
			newY := t.CustomRepeatHeader()
			t.pdf.SetY(newY)
		} else if t.RepeatHeader {
			t.AddHeader()
		}

		return true
	}

	return false
}

