// Package table provides a flexible and feature-rich table component for gofpdf.
//
// The table package enables creation of professional PDF tables with extensive
// customization options including:
//
//   - Column and row spanning for merged cells
//   - Nested tables within table cells
//   - Automatic text wrapping
//   - Customizable styling (colors, fonts, borders, alignment)
//   - Zebra striping (alternating row colors)
//   - Automatic page breaks with header repetition
//   - Per-cell, per-row, and per-column alignment overrides
//
// Quick Start:
//
//	columns := []table.Column{
//	    {Key: "id", Label: "ID", Width: 20},
//	    {Key: "name", Label: "Name", Width: 60},
//	}
//	tbl := table.NewTable(pdf, columns)
//	tbl.Render(true, data)
//
// For more examples and detailed documentation, see the README in this package.
package table

import (
	"github.com/looksocial/gofpdf"
)

// NewTable creates a new table instance with the specified columns.
//
// Parameters:
//   - pdf: The gofpdf Fpdf instance to render the table on. Must not be nil.
//   - columns: A slice of Column definitions specifying the table structure.
//
// Returns:
//   - *Table: A new Table instance with default styling and settings, or nil if pdf is nil.
//
// The table is initialized with:
//   - Default header style: bold text with gray fill (RGB: 200, 200, 200)
//   - Default data style: simple borders
//   - Automatic column width calculation for columns without specified width
//   - Header repetition enabled on new pages
//   - Automatic page breaks enabled with 20mm margin
//
// Example:
//
//	columns := []table.Column{
//	    {Key: "id", Label: "ID", Width: 20},
//	    {Key: "name", Label: "Name", Width: 60},
//	}
//	tbl := table.NewTable(pdf, columns)
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
		RepeatHeader:    true, // Default: repeat headers on new pages
		PageBreakMode:   true, // Default: enable automatic page breaks
		PageBreakMargin: 20.0, // Default: 20mm margin from bottom
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

// calculateColumnWidths automatically calculates and sets column widths for columns
// that don't have an explicit width specified. Columns with Width = 0 will be
// evenly distributed across the remaining available page width after accounting
// for margins and columns with fixed widths.
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
