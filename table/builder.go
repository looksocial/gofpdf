package table

// WithStartPosition sets the starting X and Y positions for table rendering.
//
// Parameters:
//   - x: Starting X position in mm. 0 means use current PDF X position (left margin).
//   - y: Starting Y position in mm. 0 means use current PDF Y position.
//
// Returns:
//   - *Table: Returns the Table instance for method chaining.
//
// StartX is maintained for all rows to ensure consistent left alignment.
// StartY is applied only to the first row, then cleared so subsequent rows
// use positions updated by Ln() calls.
func (t *Table) WithStartPosition(x, y float64) *Table {
	t.StartX = x
	t.StartY = y
	return t
}

// WithRowHeight sets the height for all table rows.
//
// Parameters:
//   - height: Row height in mm. Set to 0 for automatic calculation based on font size.
//
// Returns:
//   - *Table: Returns the Table instance for method chaining.
//
// When height is 0, row height is automatically calculated as approximately
// 40% of the current font size in points. Rows with wrapped text or nested
// tables may exceed this height as needed.
func (t *Table) WithRowHeight(height float64) *Table {
	t.RowHeight = height
	return t
}

// WithRowSpacing sets the spacing between table rows.
//
// Parameters:
//   - spacing: Spacing in mm to add after each row's Ln() call.
//
// Returns:
//   - *Table: Returns the Table instance for method chaining.
//
// This spacing is applied after each row, creating vertical gaps between rows.
func (t *Table) WithRowSpacing(spacing float64) *Table {
	t.Spacing = spacing
	return t
}

// WithHeaderStyle sets the styling for header row cells.
//
// Parameters:
//   - style: CellStyle to apply to all header cells.
//
// Returns:
//   - *Table: Returns the Table instance for method chaining.
//
// The style is applied when AddHeader() is called. Header alignment priority:
// Column.HeaderAlign > style.Align > Column.Align.
func (t *Table) WithHeaderStyle(style CellStyle) *Table {
	t.HeaderStyle = style
	return t
}

// WithDataStyle sets the styling for data row cells.
//
// Parameters:
//   - style: CellStyle to apply to all data cells.
//
// Returns:
//   - *Table: Returns the Table instance for method chaining.
//
// The style is applied when AddRow() is called. Data cell alignment priority:
// per-row alignment > style.Align > Column.Align.
func (t *Table) WithDataStyle(style CellStyle) *Table {
	t.DataStyle = style
	return t
}

// WithAlternatingRows enables or disables zebra striping (alternating row colors).
//
// Parameters:
//   - enabled: If true, enables alternating row fill colors using RowStyle.FillColor.
//
// Returns:
//   - *Table: Returns the Table instance for method chaining.
//
// When enabled, even-numbered rows (by position) will have the fill color
// from RowStyle.FillColor applied. Make sure to set RowStyle.FillColor when
// enabling this feature.
func (t *Table) WithAlternatingRows(enabled bool) *Table {
	t.RowStyle.Alternating = enabled
	return t
}

// WithRepeatHeader sets whether to automatically repeat the header row on new pages.
//
// Parameters:
//   - repeat: If true, header is repeated whenever a page break occurs.
//
// Returns:
//   - *Table: Returns the Table instance for method chaining.
//
// When true and a page break occurs (via PageBreakMode), AddHeader() is
// automatically called to render the header at the top of the new page.
// Default: true.
func (t *Table) WithRepeatHeader(repeat bool) *Table {
	t.RepeatHeader = repeat
	return t
}

// WithPageBreakMode enables or disables automatic page breaks for the table.
//
// Parameters:
//   - enabled: If true, automatically inserts page breaks when there's insufficient
//     space for the next row.
//
// Returns:
//   - *Table: Returns the Table instance for method chaining.
//
// When enabled, the table checks available space before rendering each row.
// If the row would exceed PageBreakMargin from the bottom, a new page is added.
// If RepeatHeader is enabled, the header is also rendered on the new page.
// Default: true.
func (t *Table) WithPageBreakMode(enabled bool) *Table {
	t.PageBreakMode = enabled
	return t
}

// WithPageBreakMargin sets the minimum margin from the bottom of the page
// before triggering an automatic page break.
//
// Parameters:
//   - margin: Margin in mm from the bottom of the page. Must be at least
//     the page's bottom margin.
//
// Returns:
//   - *Table: Returns the Table instance for method chaining.
//
// This margin is checked when PageBreakMode is enabled. If a row would
// extend closer than this margin from the bottom, a page break is triggered.
// Default: 20mm.
func (t *Table) WithPageBreakMargin(margin float64) *Table {
	t.PageBreakMargin = margin
	return t
}

