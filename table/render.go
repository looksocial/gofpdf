package table

import (
	"fmt"
	"math"
)

// AddHeader renders table headers
func (t *Table) AddHeader() {
	if len(t.Columns) == 0 {
		return
	}

	// Save current position
	startX := t.pdf.GetX()
	startY := t.pdf.GetY()

	if t.StartX > 0 {
		startX = t.StartX
		t.pdf.SetX(startX)
	}
	if t.StartY > 0 {
		startY = t.StartY
		t.pdf.SetY(startY)
	}

	// Apply header style
	t.applyCellStyle(t.HeaderStyle)

	// Render header cells
	xPos := startX
	rowHeight := t.getRowHeight()

	for _, col := range t.Columns {
		// Determine header alignment: HeaderAlign > CellStyle.Align > Column.Align
		align := col.HeaderAlign
		if align == "" {
			if t.HeaderStyle.Align != "" {
				align = t.HeaderStyle.Align
			} else {
				align = col.Align
			}
		}

		if col.ColSpan > 1 {
			// Calculate width for merged cells
			totalWidth := 0.0
			for i := 0; i < col.ColSpan; i++ {
				totalWidth += t.Columns[i].Width
			}
			t.pdf.SetXY(xPos, startY)
			t.pdf.CellFormat(totalWidth, rowHeight, col.Label, t.HeaderStyle.Border, 0,
				t.getAlignStr(align), true, 0, "")
		} else {
			t.pdf.SetXY(xPos, startY)
			t.pdf.CellFormat(col.Width, rowHeight, col.Label, t.HeaderStyle.Border, 0,
				t.getAlignStr(align), true, 0, "")
		}
		xPos += col.Width
	}

	// Move to next line
	t.pdf.Ln(rowHeight)
}

// AddRow renders a data row with support for row spans and nested tables
func (t *Table) AddRow(data map[string]interface{}) {
	if len(t.Columns) == 0 {
		return
	}

	startX := t.pdf.GetX()
	currentY := t.pdf.GetY()

	// Check if we should use StartX/StartY (like AddHeader does)
	// StartX should be used for all rows (same X position for each row)
	if t.StartX > 0 {
		startX = t.StartX
		t.pdf.SetX(startX)
	}
	// StartY should only be used for the first row, then cleared
	if t.StartY > 0 {
		currentY = t.StartY
		t.pdf.SetY(currentY)
		// Clear StartY after first use so subsequent rows use updated Y from Ln()
		t.StartY = 0
	}

	// Pre-calculate row height to check if we need a page break
	// Use a rough estimate based on row height
	estimatedRowHeight := t.getRowHeight() * 3 // Conservative estimate for wrapped text

	// Check if we need a page break before rendering this row
	if t.checkPageBreak(estimatedRowHeight) {
		// Page break was added, update positions
		startX = t.pdf.GetX()
		if t.StartX > 0 {
			startX = t.StartX
			t.pdf.SetX(startX)
		}
		currentY = t.pdf.GetY()
	}

	// Apply data style
	t.applyCellStyle(t.DataStyle)

	// Get row height
	rowHeight := t.getRowHeight()
	maxRowSpan := 1 // Track maximum row span in this row

	// Extract per-row alignment map if present
	var rowAlignments map[string]string
	var rowSpans map[string]int
	var colSpans map[string]int
	if data != nil {
		if alignMap, ok := data["_align"].(map[string]interface{}); ok {
			rowAlignments = make(map[string]string)
			for k, v := range alignMap {
				if alignStr, ok := v.(string); ok {
					rowAlignments[k] = alignStr
				}
			}
		} else if alignMap, ok := data["_align"].(map[string]string); ok {
			rowAlignments = alignMap
		}

		// Extract row spans from data
		if spanMap, ok := data["_rowspan"].(map[string]interface{}); ok {
			rowSpans = make(map[string]int)
			for k, v := range spanMap {
				if spanVal, ok := v.(int); ok {
					rowSpans[k] = spanVal
					if spanVal > maxRowSpan {
						maxRowSpan = spanVal
					}
				}
			}
		} else if spanMap, ok := data["_rowspan"].(map[string]int); ok {
			rowSpans = spanMap
			for _, spanVal := range spanMap {
				if spanVal > maxRowSpan {
					maxRowSpan = spanVal
				}
			}
		}

		// Extract column spans from data
		if spanMap, ok := data["_colspan"].(map[string]interface{}); ok {
			colSpans = make(map[string]int)
			for k, v := range spanMap {
				if spanVal, ok := v.(int); ok {
					colSpans[k] = spanVal
				}
			}
		} else if spanMap, ok := data["_colspan"].(map[string]int); ok {
			colSpans = spanMap
		}
	}

	// Track which columns are spanned by previous cells
	colSpanTracker := make(map[int]int) // column index -> remaining spans

	// Pre-calculate nested table heights and text wrapping heights to determine row height
	maxNestedHeight := 0.0
	maxTextHeight := 0.0

	for i, col := range t.Columns {
		if data != nil {
			// Check for column spans from data (need to extract this earlier)
			var cellColSpan int = col.ColSpan
			if cellColSpan == 0 {
				cellColSpan = 1
			}
			if colSpans != nil {
				if span, ok := colSpans[col.Key]; ok && span > 0 {
					cellColSpan = span
				}
			}

			// Check if this cell has a row span
			cellHasRowSpan := false
			if rowSpans != nil {
				if span, ok := rowSpans[col.Key]; ok && span > 1 {
					cellHasRowSpan = true
				}
			}

			// Check for nested tables
			// IMPORTANT: Only adjust row height for nested tables WITHOUT row spans
			// Nested tables WITH row spans will use multiple rows, so they shouldn't
			// affect the height of the current row
			if nested, ok := data["_nested_"+col.Key]; ok {
				if nt, ok := nested.(*Table); ok && !cellHasRowSpan {
					// Calculate cell width for nested table (handle column spans)
					cellWidth := col.Width
					if cellColSpan > 1 {
						cellWidth = 0.0
						for j := 0; j < cellColSpan && (i+j) < len(t.Columns); j++ {
							cellWidth += t.Columns[i+j].Width
						}
					}

					// Calculate nested table height, accounting for text wrapping
					padding := 1.0
					nestedWidth := cellWidth - 2*padding
					requiredHeight := t.calculateNestedTableHeight(nt, nestedWidth) + 2*padding
					if requiredHeight > maxNestedHeight {
						maxNestedHeight = requiredHeight
					}
				}
			}

			// Check if text needs wrapping and calculate required height
			// Skip this calculation if the cell will contain a nested table
			hasNestedTable := false
			if data != nil {
				if _, ok := data["_nested_"+col.Key]; ok {
					hasNestedTable = true
				}
			}

			if !hasNestedTable {
				if val, ok := data[col.Key]; ok {
					value := t.valueToString(val)
					if value != "" {
						cellWidth := col.Width
						cellMargin := 2.0
						usableWidth := cellWidth - cellMargin
						textWidth := t.pdf.GetStringWidth(value)

						// Only calculate wrapping for single-line text that's too long
						// Multi-line text (with \n) is handled differently by MultiCell
						if textWidth > usableWidth || (col.MaxWidth > 0 && textWidth > col.MaxWidth) {
							// More accurate estimation: use SplitLines or estimate with better margin
							// Add 20% margin to account for actual wrapping behavior
							estimatedLines := math.Ceil((textWidth / usableWidth) * 1.2)
							if estimatedLines > 1 {
								// Cap at reasonable maximum (3 lines) to prevent extreme heights
								// Very long text will be clipped rather than making rows too tall
								if estimatedLines > 3 {
									estimatedLines = 3
								}
								wrappedHeight := rowHeight * estimatedLines
								if wrappedHeight > maxTextHeight {
									maxTextHeight = wrappedHeight
								}
							}
						}
					}
				}
			}
		}
	}

	// Adjust base row height if nested tables or wrapped text require more space
	baseRowHeight := rowHeight
	if maxNestedHeight > rowHeight {
		baseRowHeight = maxNestedHeight
		// Update maxRowSpan for proper row tracking
		if maxNestedHeight > rowHeight {
			adjustedRowSpan := int(math.Ceil(maxNestedHeight / rowHeight))
			if adjustedRowSpan > maxRowSpan {
				maxRowSpan = adjustedRowSpan
			}
		}
	}
	if maxTextHeight > baseRowHeight {
		baseRowHeight = maxTextHeight
		// Update maxRowSpan for wrapped text
		if maxTextHeight > rowHeight {
			adjustedRowSpan := int(math.Ceil(maxTextHeight / rowHeight))
			if adjustedRowSpan > maxRowSpan {
				maxRowSpan = adjustedRowSpan
			}
		}
	}

	// Render cells
	xPos := startX
	colIndex := 0
	for i, col := range t.Columns {
		// Skip columns that are part of a column span from previous cells
		if remaining, ok := colSpanTracker[i]; ok && remaining > 0 {
			colSpanTracker[i] = remaining - 1
			if remaining > 1 {
				colSpanTracker[i+1] = remaining - 1
			}
			// Draw border for spanned column (use baseRowHeight for consistency)
			border := "LR"
			if i == 0 {
				border = "L"
			} else if i == len(t.Columns)-1 {
				border = "R"
			}
			t.pdf.SetXY(xPos, currentY)
			t.pdf.CellFormat(col.Width, baseRowHeight, "", border, 0, "", false, 0, "")
			xPos += col.Width
			colIndex++
			continue
		}

		// Check for column span in data
		cellColSpan := col.ColSpan
		if colSpans != nil {
			if span, ok := colSpans[col.Key]; ok && span > 0 {
				cellColSpan = span
			}
		}
		if cellColSpan == 0 {
			cellColSpan = 1
		}

		// Track column spans for subsequent columns
		if cellColSpan > 1 {
			for j := 1; j < cellColSpan; j++ {
				if i+j < len(t.Columns) {
					colSpanTracker[i+j] = cellColSpan - j
				}
			}
		}

		// Also check legacy column span from column definition (for headers)
		if col.ColSpan > 1 && cellColSpan == 1 {
			cellColSpan = col.ColSpan
			for j := 1; j < cellColSpan; j++ {
				if i+j < len(t.Columns) {
					colSpanTracker[i+j] = cellColSpan - j
				}
			}
		}

		key := fmt.Sprintf("%d-%d", i, int(currentY))
		remainingSpan := t.rowSpanTracker[key]

		// Check if this cell is part of an active row span
		if remainingSpan > 0 {
			// This cell is spanned from above, just draw border continuation
			t.rowSpanTracker[key] = remainingSpan - 1
			// Draw vertical borders only (use baseRowHeight for consistency)
			border := "LR"
			if i == 0 {
				border = "L"
			} else if i == len(t.Columns)-1 {
				border = "R"
			}
			t.pdf.SetXY(xPos, currentY)
			t.pdf.CellFormat(col.Width, baseRowHeight, "", border, 0, "", false, 0, "")
			xPos += col.Width
			colIndex++
			continue
		}

		value := ""
		cellRowSpan := col.RowSpan
		if cellRowSpan == 0 {
			cellRowSpan = 1
		}

		// Check if row span is specified in data
		if rowSpans != nil {
			if span, ok := rowSpans[col.Key]; ok && span > 0 {
				cellRowSpan = span
			}
		}

		// Check for nested table
		var nestedTable *Table
		if data != nil {
			if nested, ok := data["_nested_"+col.Key]; ok {
				if nt, ok := nested.(*Table); ok {
					nestedTable = nt
				}
			}
		}

		if data != nil {
			if val, ok := data[col.Key]; ok {
				if nestedTable == nil { // Only get text value if not nested table
					value = t.valueToString(val)
				}
			}
		}

		// Determine cell alignment: per-row > CellStyle.Align > Column.Align
		align := col.Align
		if t.DataStyle.Align != "" {
			align = t.DataStyle.Align
		}
		if rowAlignments != nil {
			if rowAlign, ok := rowAlignments[col.Key]; ok && rowAlign != "" {
				align = rowAlign
			}
		}

		t.pdf.SetXY(xPos, currentY)
		fill := t.shouldFillRow()
		border := t.DataStyle.Border

		// Calculate total height for row span
		// IMPORTANT: When a cell has a row span, we should use the ORIGINAL rowHeight
		// for calculating the span height, NOT the adjusted baseRowHeight.
		// This ensures proper grid alignment across all cells.
		totalHeight := baseRowHeight
		if cellRowSpan > 1 {
			// Use original row height for row span calculations to maintain grid alignment
			originalRowHeight := rowHeight
			totalHeight = originalRowHeight*float64(cellRowSpan) + t.Spacing*float64(cellRowSpan-1)

			// Track this row span - use originalRowHeight for tracking positions
			for j := 1; j < cellRowSpan; j++ {
				nextKey := fmt.Sprintf("%d-%d", i, int(currentY+float64(j)*originalRowHeight))
				t.rowSpanTracker[nextKey] = cellRowSpan - j
			}
		}

		// Handle nested table
		if nestedTable != nil {
			// Calculate cell width (handle column spans)
			cellWidth := col.Width
			if cellColSpan > 1 {
				cellWidth = 0.0
				for j := 0; j < cellColSpan && (i+j) < len(t.Columns); j++ {
					cellWidth += t.Columns[i+j].Width
				}
			}

			// Save current position and state
			savedX := t.pdf.GetX()
			savedY := t.pdf.GetY()

			// Draw cell border
			t.pdf.SetXY(xPos, currentY)
			if border != "" {
				t.pdf.Rect(xPos, currentY, cellWidth, totalHeight, "D")
			}

			// Calculate padding for nested table
			padding := 1.0
			nestedX := xPos + padding
			nestedY := currentY + padding
			nestedWidth := cellWidth - 2*padding
			nestedHeight := totalHeight - 2*padding

			// Scale nested table columns to fit parent cell width
			// Calculate total width of nested table columns
			totalNestedWidth := 0.0
			for _, ncol := range nestedTable.Columns {
				if ncol.Width > 0 {
					totalNestedWidth += ncol.Width
				}
			}

			// Scale factor if nested table is wider than available space
			scaleFactor := 1.0
			if totalNestedWidth > 0 && totalNestedWidth > nestedWidth {
				scaleFactor = nestedWidth / totalNestedWidth
			}

			// Save original column widths before scaling
			originalWidths := make([]float64, len(nestedTable.Columns))
			for j := range nestedTable.Columns {
				originalWidths[j] = nestedTable.Columns[j].Width
			}

			// Apply scaling to nested table columns if needed
			if scaleFactor < 1.0 {
				for j := range nestedTable.Columns {
					if nestedTable.Columns[j].Width > 0 {
						nestedTable.Columns[j].Width *= scaleFactor
					}
				}
			}

			// Save current font size before rendering nested table
			savedFontSize, _ := t.pdf.GetFontSize()

			// Reduce font size for nested tables to fit better in smaller cells
			// Use 70% of parent font size, with a minimum of 6pt
			nestedFontSize := savedFontSize * 0.7
			if nestedFontSize < 6 {
				nestedFontSize = 6
			}
			// Use empty string to keep current font family and style
			t.pdf.SetFont("", "", nestedFontSize)

			// Set clipping area to ensure nested table stays within cell boundaries
			t.pdf.ClipRect(nestedX, nestedY, nestedWidth, nestedHeight, true)

			// Render nested table inside cell with padding
			// Use a special rendering method that doesn't affect parent cursor
			t.renderNestedTable(nestedTable, nestedX, nestedY, nestedWidth, nestedHeight)

			// Clear clipping
			t.pdf.ClipEnd()

			// Restore original font size (keep family and style with empty strings)
			t.pdf.SetFont("", "", savedFontSize)

			// Restore original column widths to avoid side effects on nested table
			for j := range nestedTable.Columns {
				nestedTable.Columns[j].Width = originalWidths[j]
			}

			// Restore position
			t.pdf.SetXY(savedX, savedY)
		} else {
			// Draw simple cell - calculate width for column span
			cellWidth := col.Width
			if cellColSpan > 1 {
				cellWidth = 0.0
				for j := 0; j < cellColSpan && (i+j) < len(t.Columns); j++ {
					cellWidth += t.Columns[i+j].Width
				}
			}
			t.pdf.SetXY(xPos, currentY)

			// Check if text needs wrapping - use MultiCell if text is too long
			// Account for cell margins in width calculation
			cellMargin := 2.0 // Approximate margin
			usableWidth := cellWidth - cellMargin
			textWidth := t.pdf.GetStringWidth(value)
			needsWrapping := textWidth > usableWidth || (col.MaxWidth > 0 && textWidth > col.MaxWidth)

			if needsWrapping && value != "" {
				// For wrapped text, draw border and background manually, then render text
				savedXForCell := t.pdf.GetX()
				savedYForCell := t.pdf.GetY()

				// Draw cell background and border at the calculated baseRowHeight
				if fill {
					t.pdf.Rect(savedXForCell, savedYForCell, cellWidth, baseRowHeight, "F")
				}
				if border != "" && border != "0" {
					t.pdf.Rect(savedXForCell, savedYForCell, cellWidth, baseRowHeight, "D")
				}

				// Add small padding for text inside cell
				textPadding := 0.5
				textX := savedXForCell + textPadding
				textY := savedYForCell + textPadding
				textWidth := cellWidth - 2*textPadding

				t.pdf.SetXY(textX, textY)

				// Use MultiCell for text wrapping WITHOUT borders (we drew them already)
				// Use original row height for line spacing to keep text compact
				lineHeight := rowHeight

				// Render text without borders
				t.pdf.MultiCell(textWidth, lineHeight, value, "", t.getAlignStr(align), false)

				// Restore cursor position to continue with next cell
				// Move X to the right edge of this cell, keep Y at row start
				t.pdf.SetXY(savedXForCell+cellWidth, savedYForCell)
			} else {
				// Use regular CellFormat for single-line text
				// For cells with totalHeight (row spans), use totalHeight
				// For regular cells, use baseRowHeight
				cellHeight := totalHeight
				if totalHeight == baseRowHeight {
					cellHeight = baseRowHeight
				}

				// Render cell directly - clipping can interfere with text rendering
				t.pdf.CellFormat(cellWidth, cellHeight, value, border, 0,
					t.getAlignStr(align), fill, 0, "")
			}
		}

		// Advance xPos by the cell width (or total width if column span)
		if cellColSpan > 1 {
			for j := 0; j < cellColSpan && (i+j) < len(t.Columns); j++ {
				xPos += t.Columns[i+j].Width
			}
		} else {
			xPos += col.Width
		}
		colIndex++
	}

	// Move to next line based on the actual row height
	// Use baseRowHeight for the current row (which accounts for nested tables)
	// Row spans will extend into future rows, but the current row advances by baseRowHeight
	t.pdf.Ln(baseRowHeight + t.Spacing)
}

// AddRows stores rows for later rendering (used for nested tables)
// If you want to render immediately, use Render() instead
func (t *Table) AddRows(data []map[string]interface{}) {
	if t.storedRows == nil {
		t.storedRows = make([]map[string]interface{}, 0)
	}
	t.storedRows = append(t.storedRows, data...)
}

// AddSummaryRow adds a summary row with a label spanning multiple columns and totals
func (t *Table) AddSummaryRow(label string, labelSpan int, totals map[string]interface{}, style CellStyle) {
	if len(t.Columns) == 0 {
		return
	}

	startX := t.pdf.GetX()
	currentY := t.pdf.GetY()

	// Apply style
	t.applyCellStyle(style)

	rowHeight := t.getRowHeight()
	xPos := startX
	colIndex := 0

	// Calculate total width for label columns
	labelWidth := 0.0
	for i := 0; i < labelSpan && i < len(t.Columns); i++ {
		labelWidth += t.Columns[i].Width
	}

	// Render label cell spanning multiple columns
	t.pdf.SetXY(xPos, currentY)
	border := style.Border
	if border == "" {
		border = t.DataStyle.Border
	}
	t.pdf.CellFormat(labelWidth, rowHeight, label, border, 0, "L", false, 0, "")
	xPos += labelWidth
	colIndex = labelSpan

	// Render total columns
	for i := labelSpan; i < len(t.Columns); i++ {
		value := ""
		if totals != nil {
			if val, ok := totals[t.Columns[i].Key]; ok {
				value = t.valueToString(val)
			}
		}

		align := t.Columns[i].Align
		if style.Align != "" {
			align = style.Align
		}

		t.pdf.SetXY(xPos, currentY)
		t.pdf.CellFormat(t.Columns[i].Width, rowHeight, value, border, 0,
			t.getAlignStr(align), false, 0, "")

		xPos += t.Columns[i].Width
		colIndex++
	}

	t.pdf.Ln(rowHeight + t.Spacing)
}

// AddTotalRow adds a grand total row spanning all columns (or specified columns)
func (t *Table) AddTotalRow(label string, totals map[string]interface{}, style CellStyle) {
	t.AddSummaryRow(label, len(t.Columns), totals, style)
}

// Render renders the complete table
func (t *Table) Render(headers bool, data []map[string]interface{}) {
	if headers {
		t.AddHeader()
	}

	// Use provided data if available, otherwise use stored rows
	rowsToRender := data
	if len(rowsToRender) == 0 && len(t.storedRows) > 0 {
		rowsToRender = t.storedRows
	}

	if len(rowsToRender) > 0 {
		// For nested tables, we need to ensure proper Y positioning
		// Set initial position if StartX/StartY are specified
		if t.StartY > 0 {
			t.pdf.SetY(t.StartY)
		}
		if t.StartX > 0 {
			t.pdf.SetX(t.StartX)
		}

		// Render stored rows (but don't call AddRows to avoid recursion)
		for i, row := range rowsToRender {
			// For first row, StartY will be used and cleared by AddRow
			// For subsequent rows, ensure StartY is clear so AddRow uses GetY()
			// which should be updated by the previous row's Ln() call
			if i > 0 {
				t.StartY = 0
				// Ensure X position is maintained for all rows
				if t.StartX > 0 {
					t.pdf.SetX(t.StartX)
				}
				// Explicitly get current Y to ensure we're using the position
				// updated by the previous row's Ln() call
				currentY := t.pdf.GetY()
				t.pdf.SetY(currentY)
			}

			t.AddRow(row)
		}

		// Clear StartY after rendering all rows
		t.StartY = 0
	}
}

// calculateNestedTableHeight calculates the total height needed for a nested table
// including text wrapping in nested cells
func (t *Table) calculateNestedTableHeight(nestedTable *Table, availableWidth float64) float64 {
	if nestedTable == nil {
		return 0
	}
	if len(nestedTable.storedRows) == 0 {
		return nestedTable.getRowHeight()
	}

	nestedRowHeight := nestedTable.getRowHeight()
	nestedSpacing := nestedTable.Spacing
	totalHeight := 0.0

	// Scale nested table columns to fit available width
	totalNestedWidth := 0.0
	for _, ncol := range nestedTable.Columns {
		if ncol.Width > 0 {
			totalNestedWidth += ncol.Width
		}
	}

	scaleFactor := 1.0
	if totalNestedWidth > 0 && totalNestedWidth > availableWidth {
		scaleFactor = availableWidth / totalNestedWidth
	}

	// Save current font size
	savedFontSize, _ := t.pdf.GetFontSize()

	// Apply nested table font size (70% of parent, minimum 6pt)
	nestedFontSize := savedFontSize * 0.7
	if nestedFontSize < 6 {
		nestedFontSize = 6
	}
	// Use empty string to keep current font family and style
	t.pdf.SetFont("", "", nestedFontSize)

	// Calculate height for each row in the nested table
	for _, row := range nestedTable.storedRows {
		maxRowHeight := nestedRowHeight

		// Check each column in the nested row for text wrapping
		for _, ncol := range nestedTable.Columns {
			if row != nil {
				if val, ok := row[ncol.Key]; ok {
					value := t.valueToString(val)
					if value != "" {
						// Calculate scaled column width
						scaledColWidth := ncol.Width
						if scaleFactor < 1.0 {
							scaledColWidth = ncol.Width * scaleFactor
						}

						// Check if text needs wrapping - use GetStringWidth with scaled font
						cellMargin := 2.0
						usableWidth := scaledColWidth - cellMargin
						textWidth := t.pdf.GetStringWidth(value)

						if textWidth > usableWidth || (ncol.MaxWidth > 0 && textWidth > ncol.MaxWidth) {
							// Estimate lines needed for wrapped text
							estimatedLines := math.Ceil(textWidth / usableWidth)
							if estimatedLines < 1 {
								estimatedLines = 1
							}
							wrappedHeight := nestedRowHeight * estimatedLines
							if wrappedHeight > maxRowHeight {
								maxRowHeight = wrappedHeight
							}
						}
					}
				}
			}
		}

		totalHeight += maxRowHeight
		if totalHeight > 0 {
			totalHeight += nestedSpacing
		}
	}

	// Restore original font size (keep family and style with empty strings)
	t.pdf.SetFont("", "", savedFontSize)

	// Remove last spacing
	if totalHeight > nestedSpacing {
		totalHeight -= nestedSpacing
	}

	// Ensure minimum height
	if totalHeight < nestedRowHeight {
		totalHeight = nestedRowHeight
	}

	return totalHeight
}

// renderNestedTable renders a nested table at absolute positions without affecting parent cursor
// maxWidth and maxHeight define the clipping boundaries for the nested table
func (t *Table) renderNestedTable(nestedTable *Table, startX, startY, maxWidth, maxHeight float64) {
	// Get rows to render
	rowsToRender := nestedTable.storedRows
	if len(rowsToRender) == 0 {
		return
	}

	// Save parent PDF state (use nested table's PDF reference for consistency)
	savedX := nestedTable.pdf.GetX()
	savedY := nestedTable.pdf.GetY()
	// Disable automatic page breaks while rendering nested content to prevent
	// creation of empty pages when inner rows call Ln(). We'll restore this
	// immediately after rendering the nested table.
	autoPB, autoPBMargin := nestedTable.pdf.GetAutoPageBreak()
	nestedTable.pdf.SetAutoPageBreak(false, 0)

	// Calculate nested table row height
	nestedRowHeight := nestedTable.getRowHeight()
	nestedSpacing := nestedTable.Spacing

	// Set initial position
	currentY := startY

	// Store original nested table settings
	originalStartX := nestedTable.StartX
	originalStartY := nestedTable.StartY
	originalSpacing := nestedTable.Spacing

	// Temporarily set spacing to 0 for nested tables to prevent excessive gaps
	// Nested tables within a cell should be tightly packed
	nestedTable.Spacing = 0

	// Clear row span tracker to ensure clean state for nested table rendering
	// Each nested table rendering should start fresh
	if nestedTable.rowSpanTracker == nil {
		nestedTable.rowSpanTracker = make(map[string]int)
	} else {
		// Clear existing tracker
		for k := range nestedTable.rowSpanTracker {
			delete(nestedTable.rowSpanTracker, k)
		}
	}

	// Set nested table to use absolute positioning
	nestedTable.StartX = startX

	// Render each row at absolute positions
	for i, row := range rowsToRender {
		// Quick check to avoid obvious overflow (exact height will be computed after AddRow)
		if currentY+nestedRowHeight > startY+maxHeight {
			break
		}

		// Position the nested table absolutely for this row
		nestedTable.StartX = startX
		if i == 0 {
			nestedTable.StartY = currentY
		} else {
			nestedTable.StartY = currentY
		}
		nestedTable.pdf.SetXY(startX, currentY)

		// Let AddRow perform its full layout (including wrapping and spans)
		rowStartY := nestedTable.pdf.GetY()
		nestedTable.AddRow(row)
		rowEndY := nestedTable.pdf.GetY()

		// Compute the actual vertical advance that AddRow applied
		deltaY := rowEndY - rowStartY
		if deltaY < 0 {
			// Safety: if something went wrong, fall back to nominal height
			deltaY = nestedRowHeight + nestedSpacing
		}

		// Restore position so parent table state is unchanged
		nestedTable.pdf.SetXY(savedX, savedY)

		// Advance to next row based on the actual delta applied by AddRow
		currentY += deltaY
		if currentY > startY+maxHeight {
			break
		}
	}

	// Restore nested table's original StartX/StartY and spacing after all rows are rendered
	nestedTable.StartX = originalStartX
	nestedTable.StartY = originalStartY
	nestedTable.Spacing = originalSpacing

	// Restore parent PDF state (use nested table's PDF reference for consistency)
	nestedTable.pdf.SetXY(savedX, savedY)
	// Restore original automatic page break settings
	nestedTable.pdf.SetAutoPageBreak(autoPB, autoPBMargin)
}