package table

import (
	"fmt"
	"math"
)

// shouldFillRow returns true if current row should be filled (for zebra striping)
func (t *Table) shouldFillRow() bool {
	if t.RowStyle.Alternating {
		// Simple alternating based on Y position
		currentY := t.pdf.GetY()
		cellHeight := t.getRowHeight()
		rows := int(math.Round(currentY / (cellHeight + t.Spacing)))
		return rows%2 == 1
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
	if len(style.TextColor) >= 3 {
		t.pdf.SetTextColor(style.TextColor[0], style.TextColor[1], style.TextColor[2])
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

