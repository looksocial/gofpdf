package table

import (
	"github.com/looksocial/gofpdf"
)

// NewTable creates a new table instance
func NewTable(pdf *gofpdf.Fpdf, columns []Column) *Table {
	if pdf == nil {
		// Can't log error without pdf, so return nil and let caller handle
		return nil
	}

	t := &Table{
		pdf:             pdf,
		Columns:         columns,
		StartX:          0,
		StartY:          0,
		AutoWidth:       false,
		RowHeight:       0,
		Spacing:         0,
		RepeatHeader:    true,  // Default: repeat headers on new pages
		PageBreakMode:   true,  // Default: enable automatic page breaks
		PageBreakMargin: 20.0,  // Default: 20mm margin from bottom
		rowSpanTracker:  make(map[string]int),
		storedRows:      nil,
		HeaderStyle: CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{200, 200, 200},
		},
		DataStyle: CellStyle{
			Border: "1",
		},
		RowStyle: RowStyle{
			Border: "1",
		},
	}

	// Set default column widths if not specified
	t.calculateColumnWidths()

	return t
}

// calculateColumnWidths sets default widths for columns without specified width
func (t *Table) calculateColumnWidths() {
	totalWidth := 0.0
	colsWithoutWidth := 0

	// Count columns without specified width and sum existing widths
	for i := range t.Columns {
		if t.Columns[i].Width == 0 {
			colsWithoutWidth++
		} else {
			totalWidth += t.Columns[i].Width
		}
		// Set ColSpan from MergeCell for backward compatibility
		if t.Columns[i].MergeCell && t.Columns[i].ColSpan == 0 {
			t.Columns[i].ColSpan = 1
		}
		if t.Columns[i].ColSpan == 0 {
			t.Columns[i].ColSpan = 1
		}
	}

	// Auto-calculate remaining columns if needed
	if colsWithoutWidth > 0 {
		left, _, right, _ := t.pdf.GetMargins()
		pageWidth, _ := t.pdf.GetPageSize()
		// Calculate usable width
		usableWidth := pageWidth - left - right - totalWidth
		avgWidth := usableWidth / float64(colsWithoutWidth)

		for i := range t.Columns {
			if t.Columns[i].Width == 0 {
				t.Columns[i].Width = avgWidth
			}
		}
	}
}
