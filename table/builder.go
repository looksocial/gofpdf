package table

// WithStartPosition sets the starting position of the table
func (t *Table) WithStartPosition(x, y float64) *Table {
	t.StartX = x
	t.StartY = y
	return t
}

// WithRowHeight sets the row height
func (t *Table) WithRowHeight(height float64) *Table {
	t.RowHeight = height
	return t
}

// WithRowSpacing sets spacing between rows
func (t *Table) WithRowSpacing(spacing float64) *Table {
	t.Spacing = spacing
	return t
}

// WithHeaderStyle sets the header styling
func (t *Table) WithHeaderStyle(style CellStyle) *Table {
	t.HeaderStyle = style
	return t
}

// WithDataStyle sets the data row styling
func (t *Table) WithDataStyle(style CellStyle) *Table {
	t.DataStyle = style
	return t
}

// WithAlternatingRows enables/disables zebra striping
func (t *Table) WithAlternatingRows(enabled bool) *Table {
	t.RowStyle.Alternating = enabled
	return t
}

// WithRepeatHeader sets whether to repeat header on new pages
func (t *Table) WithRepeatHeader(repeat bool) *Table {
	t.RepeatHeader = repeat
	return t
}

// WithPageBreakMode enables/disables automatic page breaks
func (t *Table) WithPageBreakMode(enabled bool) *Table {
	t.PageBreakMode = enabled
	return t
}

// WithPageBreakMargin sets the margin from bottom before page break
func (t *Table) WithPageBreakMargin(margin float64) *Table {
	t.PageBreakMargin = margin
	return t
}

