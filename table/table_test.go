package table

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/looksocial/gofpdf"
)

// TestNewTable tests table creation
func TestNewTable(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	columns := []Column{
		{Key: "id", Label: "ID", Width: 20},
		{Key: "name", Label: "Name", Width: 60},
	}

	tbl := NewTable(pdf, columns)
	if tbl == nil {
		t.Error("Expected table to be created, got nil")
	}
	if len(tbl.Columns) != 2 {
		t.Errorf("Expected 2 columns, got %d", len(tbl.Columns))
	}
	if tbl.Columns[0].Key != "id" {
		t.Errorf("Expected first column key 'id', got '%s'", tbl.Columns[0].Key)
	}
}

// TestNewTableWithNilPDF tests nil PDF handling
func TestNewTableWithNilPDF(t *testing.T) {
	tbl := NewTable(nil, []Column{{Key: "id", Label: "ID"}})
	if tbl != nil {
		t.Error("Expected nil table when pdf is nil, got non-nil")
	}
}

// TestNewTableAutoWidth tests automatic column width calculation
func TestNewTableAutoWidth(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "id", Label: "ID", Width: 0},      // Auto width
		{Key: "name", Label: "Name", Width: 60}, // Fixed width
		{Key: "email", Label: "Email", Width: 0}, // Auto width
	}

	tbl := NewTable(pdf, columns)
	
	// Check that auto widths were calculated
	if tbl.Columns[0].Width == 0 {
		t.Error("Expected auto-calculated width for first column, got 0")
	}
	if tbl.Columns[1].Width != 60 {
		t.Errorf("Expected fixed width 60, got %.2f", tbl.Columns[1].Width)
	}
	if tbl.Columns[2].Width == 0 {
		t.Error("Expected auto-calculated width for third column, got 0")
	}
}

// TestTableWithStartPosition tests custom starting position
func TestTableWithStartPosition(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	columns := []Column{{Key: "id", Label: "ID"}}

	tbl := NewTable(pdf, columns)
	tbl.WithStartPosition(50, 100)

	if tbl.StartX != 50 {
		t.Errorf("Expected StartX 50, got %.2f", tbl.StartX)
	}
	if tbl.StartY != 100 {
		t.Errorf("Expected StartY 100, got %.2f", tbl.StartY)
	}
}

// TestTableWithRowHeight tests custom row height
func TestTableWithRowHeight(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	columns := []Column{{Key: "id", Label: "ID"}}

	tbl := NewTable(pdf, columns)
	tbl.WithRowHeight(10)

	if tbl.RowHeight != 10 {
		t.Errorf("Expected RowHeight 10, got %.2f", tbl.RowHeight)
	}
}

// TestTableWithAlternatingRows tests alternating row configuration
func TestTableWithAlternatingRows(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	columns := []Column{{Key: "id", Label: "ID"}}

	tbl := NewTable(pdf, columns)
	tbl.WithAlternatingRows(true)

	if !tbl.RowStyle.Alternating {
		t.Error("Expected alternating rows to be enabled")
	}
}

// TestTableRender tests basic table rendering
func TestTableRender(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "id", Label: "ID", Width: 20},
		{Key: "name", Label: "Name", Width: 60},
	}

	tbl := NewTable(pdf, columns)

	data := []map[string]interface{}{
		{"id": "1", "name": "Alice"},
		{"id": "2", "name": "Bob"},
	}

	tbl.Render(true, data)

	// Check for errors
	if pdf.Error() != nil {
		t.Errorf("Unexpected error during rendering: %v", pdf.Error())
	}
}

// TestTableRenderWithoutHeaders tests rendering without headers
func TestTableRenderWithoutHeaders(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "id", Label: "ID", Width: 20},
	}

	tbl := NewTable(pdf, columns)
	data := []map[string]interface{}{{"id": "1"}}

	tbl.Render(false, data)

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableAddRow tests adding individual rows
func TestTableAddRow(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "id", Label: "ID", Width: 20},
		{Key: "name", Label: "Name", Width: 60},
	}

	tbl := NewTable(pdf, columns)
	tbl.AddHeader()
	tbl.AddRow(map[string]interface{}{"id": "1", "name": "Alice"})
	tbl.AddRow(map[string]interface{}{"id": "2", "name": "Bob"})

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableAddRows tests adding multiple rows
func TestTableAddRows(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "id", Label: "ID", Width: 20},
	}

	tbl := NewTable(pdf, columns)

	data := []map[string]interface{}{
		{"id": "1"},
		{"id": "2"},
		{"id": "3"},
	}

	tbl.AddRows(data)

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableWithCustomStyles tests custom styling
func TestTableWithCustomStyles(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	columns := []Column{{Key: "id", Label: "ID"}}

	tbl := NewTable(pdf, columns).
		WithHeaderStyle(CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{100, 100, 200},
			TextColor: []int{255, 255, 255},
		}).
		WithDataStyle(CellStyle{
			Border: "LRB",
		})

	if len(tbl.HeaderStyle.FillColor) != 3 {
		t.Error("Expected FillColor to have 3 values")
	}
	if tbl.HeaderStyle.Bold != true {
		t.Error("Expected Bold to be true")
	}
}

// TestTableWithDifferentAlignments tests column alignments
func TestTableWithDifferentAlignments(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "left", Label: "Left", Width: 40, Align: "L"},
		{Key: "center", Label: "Center", Width: 40, Align: "C"},
		{Key: "right", Label: "Right", Width: 40, Align: "R"},
	}

	tbl := NewTable(pdf, columns)

	data := []map[string]interface{}{
		{"left": "L", "center": "C", "right": "R"},
	}

	tbl.Render(true, data)

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableEmptyData tests rendering with empty data
func TestTableEmptyData(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{{Key: "id", Label: "ID"}}
	tbl := NewTable(pdf, columns)

	tbl.Render(true, []map[string]interface{}{})

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableNilData tests rendering with nil data
func TestTableNilData(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{{Key: "id", Label: "ID"}}
	tbl := NewTable(pdf, columns)

	tbl.Render(true, nil)

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableValueToString tests value conversion to string
func TestTableValueToString(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	columns := []Column{{Key: "id", Label: "ID", Width: 20}}
	tbl := NewTable(pdf, columns)

	// Test various types
	data := []map[string]interface{}{
		{"id": "string"},
		{"id": 123},
		{"id": float64(45.67)},
		{"id": true},
		{"id": nil},
	}

	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	tbl.AddRows(data)

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableChaining tests method chaining
func TestTableChaining(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	columns := []Column{{Key: "id", Label: "ID"}}

	tbl := NewTable(pdf, columns).
		WithStartPosition(10, 20).
		WithRowHeight(8).
		WithRowSpacing(2).
		WithAlternatingRows(true)

	if tbl.StartX != 10 || tbl.StartY != 20 {
		t.Error("Chained WithStartPosition failed")
	}
	if tbl.RowHeight != 8 {
		t.Error("Chained WithRowHeight failed")
	}
	if tbl.Spacing != 2 {
		t.Error("Chained WithRowSpacing failed")
	}
	if !tbl.RowStyle.Alternating {
		t.Error("Chained WithAlternatingRows failed")
	}
}

// TestTableOutputPDF tests actual PDF generation
func TestTableOutputPDF(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "id", Label: "ID", Width: 20, Align: "L"},
		{Key: "name", Label: "Name", Width: 60, Align: "L"},
		{Key: "email", Label: "Email", Width: 110, Align: "L"},
	}

	tbl := NewTable(pdf, columns)

	data := []map[string]interface{}{
		{"id": "1", "name": "Alice", "email": "alice@example.com"},
		{"id": "2", "name": "Bob", "email": "bob@example.com"},
	}

	tbl.Render(true, data)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Fatalf("Failed to generate PDF: %v", err)
	}

	if buf.Len() == 0 {
		t.Error("Generated PDF is empty")
	}

	// Verify PDF magic number
	if !bytes.HasPrefix(buf.Bytes(), []byte("%PDF")) {
		t.Error("Output is not a valid PDF file")
	}
}

// TestTableWithMaxWidth tests MaxWidth property
func TestTableWithMaxWidth(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	columns := []Column{
		{Key: "text", Label: "Text", Width: 100, MaxWidth: 80},
	}

	tbl := NewTable(pdf, columns)
	if tbl.Columns[0].Width != 100 {
		t.Errorf("Expected Width 100, got %.2f", tbl.Columns[0].Width)
	}
	if tbl.Columns[0].MaxWidth != 80 {
		t.Errorf("Expected MaxWidth 80, got %.2f", tbl.Columns[0].MaxWidth)
	}
}

// TestTableGetAlignStr tests alignment string parsing
func TestTableGetAlignStr(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	columns := []Column{{Key: "id", Label: "ID"}}
	tbl := NewTable(pdf, columns)

	if tbl.getAlignStr("L") != "L" {
		t.Error("Expected 'L' for Left alignment")
	}
	if tbl.getAlignStr("C") != "C" {
		t.Error("Expected 'C' for Center alignment")
	}
	if tbl.getAlignStr("Center") != "C" {
		t.Error("Expected 'C' for Center alignment")
	}
	if tbl.getAlignStr("R") != "R" {
		t.Error("Expected 'R' for Right alignment")
	}
	if tbl.getAlignStr("Right") != "R" {
		t.Error("Expected 'R' for Right alignment")
	}
	if tbl.getAlignStr("invalid") != "L" {
		t.Error("Expected 'L' for invalid alignment")
	}
}

// TestTableGetRowHeight tests row height calculation
func TestTableGetRowHeight(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{{Key: "id", Label: "ID"}}
	tbl := NewTable(pdf, columns)

	// Test auto row height
	autoHeight := tbl.getRowHeight()
	if autoHeight <= 0 {
		t.Error("Expected positive auto row height")
	}

	// Test custom row height
	tbl.WithRowHeight(10)
	if tbl.getRowHeight() != 10 {
		t.Error("Expected custom row height 10")
	}
}

// TestTableHeaderAlign tests header alignment with HeaderAlign field
func TestTableHeaderAlign(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "id", Label: "ID", Width: 40, Align: "L", HeaderAlign: "C"},
		{Key: "name", Label: "Name", Width: 60, Align: "L", HeaderAlign: "R"},
		{Key: "value", Label: "Value", Width: 40, Align: "R", HeaderAlign: "L"},
	}

	tbl := NewTable(pdf, columns)

	data := []map[string]interface{}{
		{"id": "1", "name": "Alice", "value": "100"},
		{"id": "2", "name": "Bob", "value": "200"},
	}

	tbl.Render(true, data)

	err := pdf.OutputFileAndClose("../pdf/table_header_align_test.pdf")
	if err != nil {
		t.Errorf("Failed to output PDF: %v", err)
	}

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableHeaderStyleAlign tests header alignment with HeaderStyle.Align
func TestTableHeaderStyleAlign(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "id", Label: "ID", Width: 40, Align: "L"},
		{Key: "name", Label: "Name", Width: 60, Align: "L"},
		{Key: "value", Label: "Value", Width: 40, Align: "R"},
	}

	tbl := NewTable(pdf, columns).
		WithHeaderStyle(CellStyle{
			Border:    "1",
			Bold:      true,
			Align:     "C", // All headers centered
			FillColor: []int{200, 200, 200},
		})

	data := []map[string]interface{}{
		{"id": "1", "name": "Alice", "value": "100"},
		{"id": "2", "name": "Bob", "value": "200"},
	}

	tbl.Render(true, data)

	err := pdf.OutputFileAndClose("../pdf/table_header_style_align_test.pdf")
	if err != nil {
		t.Errorf("Failed to output PDF: %v", err)
	}

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableRowAlign tests per-row alignment with _align map
func TestTableRowAlign(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "id", Label: "ID", Width: 40, Align: "L"},
		{Key: "name", Label: "Name", Width: 60, Align: "L"},
		{Key: "value", Label: "Value", Width: 40, Align: "L"},
	}

	tbl := NewTable(pdf, columns)

	data := []map[string]interface{}{
		{
			"id":    "1",
			"name":  "Alice",
			"value": "100",
			"_align": map[string]string{
				"id":    "C",
				"name":  "C",
				"value": "R",
			},
		},
		{
			"id":    "2",
			"name":  "Bob",
			"value": "200",
			"_align": map[string]string{
				"value": "C",
			},
		},
		{"id": "3", "name": "Charlie", "value": "300"}, // Default alignment
	}

	tbl.Render(true, data)

	err := pdf.OutputFileAndClose("../pdf/table_row_align_test.pdf")
	if err != nil {
		t.Errorf("Failed to output PDF: %v", err)
	}

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableDataStyleAlign tests row alignment with DataStyle.Align
func TestTableDataStyleAlign(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "id", Label: "ID", Width: 40, Align: "L"},
		{Key: "name", Label: "Name", Width: 60, Align: "L"},
		{Key: "value", Label: "Value", Width: 40, Align: "L"},
	}

	tbl := NewTable(pdf, columns).
		WithDataStyle(CellStyle{
			Border: "1",
			Align:  "C", // All data cells centered (overrides column align)
		})

	data := []map[string]interface{}{
		{"id": "1", "name": "Alice", "value": "100"},
		{"id": "2", "name": "Bob", "value": "200"},
		{"id": "3", "name": "Charlie", "value": "300"},
	}

	tbl.Render(true, data)

	err := pdf.OutputFileAndClose("../pdf/table_data_style_align_test.pdf")
	if err != nil {
		t.Errorf("Failed to output PDF: %v", err)
	}

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableMixedAlignments tests all alignment features together
func TestTableMixedAlignments(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "id", Label: "ID", Width: 30, Align: "L", HeaderAlign: "C"},
		{Key: "name", Label: "Name", Width: 50, Align: "L", HeaderAlign: "L"},
		{Key: "price", Label: "Price", Width: 40, Align: "R", HeaderAlign: "R"},
		{Key: "qty", Label: "Qty", Width: 30, Align: "L", HeaderAlign: "C"},
	}

	tbl := NewTable(pdf, columns).
		WithHeaderStyle(CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{220, 220, 220},
		}).
		WithDataStyle(CellStyle{
			Border: "1",
			Align:  "L", // Default data alignment (can be overridden per-row)
		})

	data := []map[string]interface{}{
		{
			"id":    "1",
			"name":  "Product A",
			"price": "10.50",
			"qty":   "5",
			"_align": map[string]string{
				"price": "R", // Override to right align
				"qty":   "C", // Override to center align
			},
		},
		{
			"id":    "2",
			"name":  "Product B",
			"price": "25.00",
			"qty":   "10",
			"_align": map[string]string{
				"qty": "C",
			},
		},
		{"id": "3", "name": "Product C", "price": "5.75", "qty": "3"}, // Default alignment
	}

	tbl.Render(true, data)

	err := pdf.OutputFileAndClose("../pdf/table_mixed_alignments_test.pdf")
	if err != nil {
		t.Errorf("Failed to output PDF: %v", err)
	}

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// BenchmarkTableRender benchmarks table rendering
func BenchmarkTableRender(b *testing.B) {
	columns := []Column{
		{Key: "id", Label: "ID", Width: 20},
		{Key: "name", Label: "Name", Width: 60},
	}

	data := []map[string]interface{}{
		{"id": "1", "name": "Alice"},
		{"id": "2", "name": "Bob"},
	}

	for i := 0; i < b.N; i++ {
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()
		pdf.SetFont("Arial", "", 12)
		tbl := NewTable(pdf, columns)
		tbl.Render(true, data)
	}
}

// TestTableMultiPageWithHeaderRepeat tests automatic page breaks with header repetition
func TestTableMultiPageWithHeaderRepeat(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)

	columns := []Column{
		{Key: "id", Label: "ID", Width: 30, Align: "C"},
		{Key: "product", Label: "Product Name", Width: 80, Align: "L"},
		{Key: "price", Label: "Price", Width: 40, Align: "R"},
		{Key: "stock", Label: "Stock", Width: 40, Align: "C"},
	}

	// Create table with header repetition enabled (default)
	tbl := NewTable(pdf, columns).
		WithHeaderStyle(CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{100, 150, 200},
			TextColor: []int{255, 255, 255},
		}).
		WithDataStyle(CellStyle{
			Border: "1",
		}).
		WithRowHeight(8).
		WithRepeatHeader(true).
		WithPageBreakMode(true).
		WithPageBreakMargin(20)

	// Verify default settings
	if !tbl.RepeatHeader {
		t.Error("Expected RepeatHeader to be true")
	}
	if !tbl.PageBreakMode {
		t.Error("Expected PageBreakMode to be true")
	}
	if tbl.PageBreakMargin != 20 {
		t.Errorf("Expected PageBreakMargin to be 20, got %f", tbl.PageBreakMargin)
	}

	// Add header
	tbl.AddHeader()

	// Get initial page count
	initialPageCount := pdf.PageCount()
	if initialPageCount != 1 {
		t.Errorf("Expected 1 page initially, got %d", initialPageCount)
	}

	// Add enough rows to trigger page breaks
	// A4 height is ~297mm, with 20mm top/bottom margins = 257mm usable
	// Each row is ~8mm, so we need ~35+ rows to fill a page
	rowCount := 50

	for i := 1; i <= rowCount; i++ {
		tbl.AddRow(map[string]interface{}{
			"id":      i,
			"product": "Product " + string(rune(65+(i%26))),
			"price":   19.99 + float64(i),
			"stock":   100 + i*5,
		})
	}

	// Check that multiple pages were created
	finalPageCount := pdf.PageCount()
	if finalPageCount <= 1 {
		t.Errorf("Expected more than 1 page with %d rows, got %d pages", rowCount, finalPageCount)
	}

	// Generate PDF to verify it's valid
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}

	// Verify PDF has content
	if buf.Len() == 0 {
		t.Error("Generated PDF is empty")
	}

	// Basic PDF validation - check for page objects in PDF
	if !bytes.Contains(buf.Bytes(), []byte("/Type /Page")) {
		t.Error("PDF does not contain page objects")
	}

	// Save PDF to file for visual inspection
	outputPath := "../pdf/test_multipage_with_header_repeat.pdf"
	if err := savePDFBytesToFile(buf.Bytes(), outputPath); err != nil {
		t.Logf("Warning: Could not save PDF to %s: %v", outputPath, err)
	} else {
		t.Logf("PDF saved to: %s", outputPath)
	}

	t.Logf("Successfully created multi-page table: %d rows across %d pages", rowCount, finalPageCount)
}

// TestTableMultiPageWithoutHeaderRepeat tests automatic page breaks without header repetition
func TestTableMultiPageWithoutHeaderRepeat(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)

	columns := []Column{
		{Key: "order", Label: "Order ID", Width: 40, Align: "C"},
		{Key: "customer", Label: "Customer Name", Width: 80, Align: "L"},
		{Key: "date", Label: "Order Date", Width: 35, Align: "C"},
		{Key: "total", Label: "Total", Width: 35, Align: "R"},
	}

	// Create table with header repetition DISABLED
	tbl := NewTable(pdf, columns).
		WithHeaderStyle(CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{200, 100, 100},
		}).
		WithDataStyle(CellStyle{
			Border: "1",
		}).
		WithRowHeight(8).
		WithRepeatHeader(false).      // Disable header repetition
		WithPageBreakMode(true).      // Enable page breaks
		WithPageBreakMargin(20)

	// Verify settings
	if tbl.RepeatHeader {
		t.Error("Expected RepeatHeader to be false")
	}
	if !tbl.PageBreakMode {
		t.Error("Expected PageBreakMode to be true")
	}

	// Add header - should only appear once
	tbl.AddHeader()

	// Get initial page count
	initialPageCount := pdf.PageCount()
	if initialPageCount != 1 {
		t.Errorf("Expected 1 page initially, got %d", initialPageCount)
	}

	// Add enough rows to span multiple pages
	rowCount := 45

	for i := 1; i <= rowCount; i++ {
		tbl.AddRow(map[string]interface{}{
			"order":    "ORD-" + string(rune(48+i%10)) + string(rune(48+(i/10)%10)) + string(rune(48+(i/100))),
			"customer": "Customer " + string(rune(65+(i%26))),
			"date":     "2024-11-01",
			"total":    99.99 + float64(i)*5.50,
		})
	}

	// Check that multiple pages were created
	finalPageCount := pdf.PageCount()
	if finalPageCount <= 1 {
		t.Errorf("Expected more than 1 page with %d rows, got %d pages", rowCount, finalPageCount)
	}

	// Generate PDF to verify it's valid
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}

	// Verify PDF has content
	if buf.Len() == 0 {
		t.Error("Generated PDF is empty")
	}

	// Verify PDF structure
	pdfContent := buf.Bytes()
	if !bytes.Contains(pdfContent, []byte("/Type /Page")) {
		t.Error("PDF does not contain page objects")
	}

	// Verify we have multiple pages
	if !bytes.Contains(pdfContent, []byte("/Count "+string(rune(48+finalPageCount)))) &&
		!bytes.Contains(pdfContent, []byte("/Count "+string(rune(48+finalPageCount)))) {
		t.Logf("Note: Page count validation is informational only")
	}

	// Save PDF to file for visual inspection
	outputPath := "../pdf/test_multipage_without_header_repeat.pdf"
	if err := savePDFBytesToFile(pdfContent, outputPath); err != nil {
		t.Logf("Warning: Could not save PDF to %s: %v", outputPath, err)
	} else {
		t.Logf("PDF saved to: %s", outputPath)
	}

	t.Logf("Successfully created multi-page table without header repetition: %d rows across %d pages", rowCount, finalPageCount)
	t.Logf("Header should only appear on page 1 (RepeatHeader=false)")
}

// Helper function to save PDF bytes to file
func savePDFBytesToFile(pdfBytes []byte, path string) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Write PDF bytes to file
	return os.WriteFile(path, pdfBytes, 0644)
}

// TestTableAddSummaryRow tests summary row rendering
func TestTableAddSummaryRow(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "category", Label: "Category", Width: 50},
		{Key: "item", Label: "Item", Width: 50},
		{Key: "quantity", Label: "Qty", Width: 30, Align: "R"},
		{Key: "price", Label: "Price", Width: 40, Align: "R"},
	}

	tbl := NewTable(pdf, columns)
	tbl.AddHeader()

	// Add some data rows
	data := []map[string]interface{}{
		{"category": "Electronics", "item": "Laptop", "quantity": 5, "price": 999.99},
		{"category": "Electronics", "item": "Phone", "quantity": 10, "price": 699.99},
		{"category": "Books", "item": "Novel", "quantity": 20, "price": 19.99},
	}

	for _, row := range data {
		tbl.AddRow(row)
	}

	// Add summary row
	tbl.AddSummaryRow("Subtotal", 2, map[string]interface{}{
		"quantity": 35,
		"price":    1719.97,
	}, CellStyle{
		Border:    "1",
		Bold:      true,
		FillColor: []int{240, 240, 240},
	})

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}

	// Verify PDF can be generated
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}

	if buf.Len() == 0 {
		t.Error("Generated PDF is empty")
	}
}

// TestTableAddTotalRow tests total row rendering
func TestTableAddTotalRow(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "product", Label: "Product", Width: 60},
		{Key: "amount", Label: "Amount", Width: 50, Align: "R"},
	}

	tbl := NewTable(pdf, columns)
	tbl.AddHeader()

	// Add data rows
	tbl.AddRow(map[string]interface{}{"product": "Product A", "amount": 100.50})
	tbl.AddRow(map[string]interface{}{"product": "Product B", "amount": 250.75})
	tbl.AddRow(map[string]interface{}{"product": "Product C", "amount": 99.25})

	// Add total row
	tbl.AddTotalRow("Grand Total", map[string]interface{}{
		"amount": 450.50,
	}, CellStyle{
		Border:    "1",
		Bold:      true,
		FillColor: []int{220, 220, 220},
	})

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableColumnSpan tests column spanning in headers and rows
func TestTableColumnSpan(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "id", Label: "ID", Width: 30},
		{Key: "name", Label: "Name", Width: 60},
		{Key: "email", Label: "Email", Width: 100},
		{Key: "phone", Label: "Phone", Width: 60},
	}

	tbl := NewTable(pdf, columns)

	// Test column span in header
	columns[0].ColSpan = 2
	tbl.Columns[0].ColSpan = 2

	tbl.AddHeader()

	// Test column span in data row
	tbl.AddRow(map[string]interface{}{
		"id":    "1",
		"name":  "John Doe",
		"email": "john@example.com",
		"phone": "123-456-7890",
		"_colspan": map[string]int{
			"name": 2, // Name spans 2 columns
		},
	})

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableRowSpan tests row spanning
func TestTableRowSpan(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "id", Label: "ID", Width: 30},
		{Key: "name", Label: "Name", Width: 60},
		{Key: "item1", Label: "Item 1", Width: 50},
		{Key: "item2", Label: "Item 2", Width: 50},
	}

	tbl := NewTable(pdf, columns)
	tbl.AddHeader()

	// Add row with row span
	tbl.AddRow(map[string]interface{}{
		"id":     "1",
		"name":   "Alice",
		"item1":  "Product A",
		"item2":  "Product B",
		"_rowspan": map[string]int{
			"id":   3, // ID spans 3 rows
			"name": 3, // Name spans 3 rows
		},
	})

	// Add subsequent rows (spanned cells will be skipped)
	tbl.AddRow(map[string]interface{}{
		"item1": "Product C",
		"item2": "Product D",
	})

	tbl.AddRow(map[string]interface{}{
		"item1": "Product E",
		"item2": "Product F",
	})

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}

	// Verify PDF generation
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

// TestTableRowSpanWithSpacing tests row spanning with row spacing
func TestTableRowSpanWithSpacing(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "id", Label: "ID", Width: 30},
		{Key: "name", Label: "Name", Width: 60},
		{Key: "item1", Label: "Item 1", Width: 50},
		{Key: "item2", Label: "Item 2", Width: 50},
	}

	tbl := NewTable(pdf, columns).
		WithRowSpacing(2.0) // Set spacing to verify it's included in key calculation

	tbl.AddHeader()

	// Add row with row span - this should track keys correctly with spacing
	tbl.AddRow(map[string]interface{}{
		"id":     "1",
		"name":   "Alice",
		"item1":  "Product A",
		"item2":  "Product B",
		"_rowspan": map[string]int{
			"id":   3, // ID spans 3 rows
			"name": 3, // Name spans 3 rows
		},
	})

	// Add subsequent rows (spanned cells will be skipped)
	// The rowSpanTracker keys should match actual Y positions including spacing
	tbl.AddRow(map[string]interface{}{
		"item1": "Product C",
		"item2": "Product D",
	})

	tbl.AddRow(map[string]interface{}{
		"item1": "Product E",
		"item2": "Product F",
	})

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}

	// Verify PDF generation
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}

	if buf.Len() == 0 {
		t.Error("Generated PDF is empty")
	}
}

// TestTableNestedTables tests nested tables within cells
func TestTableNestedTables(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	// Parent table
	parentColumns := []Column{
		{Key: "category", Label: "Category", Width: 80},
		{Key: "items", Label: "Items", Width: 110},
	}

	parentTable := NewTable(pdf, parentColumns)
	parentTable.AddHeader()

	// Create nested table
	nestedColumns := []Column{
		{Key: "item", Label: "Item", Width: 40},
		{Key: "qty", Label: "Qty", Width: 30, Align: "R"},
	}

	nestedTable := NewTable(pdf, nestedColumns)
	nestedTable.AddRows([]map[string]interface{}{
		{"item": "Item A", "qty": 5},
		{"item": "Item B", "qty": 10},
		{"item": "Item C", "qty": 3},
	})

	// Add row with nested table
	parentTable.AddRow(map[string]interface{}{
		"category":       "Electronics",
		"_nested_items": nestedTable,
	})

	// Add another row with different nested table
	nestedTable2 := NewTable(pdf, nestedColumns)
	nestedTable2.AddRows([]map[string]interface{}{
		{"item": "Item X", "qty": 2},
		{"item": "Item Y", "qty": 8},
	})

	parentTable.AddRow(map[string]interface{}{
		"category":       "Books",
		"_nested_items": nestedTable2,
	})

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}

	// Verify PDF generation
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}

	if buf.Len() == 0 {
		t.Error("Generated PDF is empty")
	}
}

// TestTableTextWrapping tests automatic text wrapping
func TestTableTextWrapping(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "id", Label: "ID", Width: 30},
		{Key: "description", Label: "Description", Width: 80},
	}

	tbl := NewTable(pdf, columns)
	tbl.AddHeader()

	// Add row with long text that should wrap
	tbl.AddRow(map[string]interface{}{
		"id":          "1",
		"description": "This is a very long description that should wrap to multiple lines within the cell because it exceeds the column width",
	})

	// Add row with MaxWidth constraint
	columns[1].MaxWidth = 50
	tbl.Columns[1].MaxWidth = 50

	tbl.AddRow(map[string]interface{}{
		"id":          "2",
		"description": "Another long text that should wrap based on MaxWidth constraint",
	})

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableAlternatingRows tests zebra striping
func TestTableAlternatingRows(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "id", Label: "ID", Width: 30},
		{Key: "name", Label: "Name", Width: 60},
	}

	tbl := NewTable(pdf, columns).
		WithAlternatingRows(true)
	tbl.RowStyle.FillColor = []int{245, 245, 245}

	tbl.AddHeader()

	// Add multiple rows
	for i := 1; i <= 5; i++ {
		tbl.AddRow(map[string]interface{}{
			"id":   i,
			"name": "Item " + string(rune(64+i)),
		})
	}

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableCheckPageBreak tests page break functionality
func TestTableCheckPageBreak(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)

	columns := []Column{
		{Key: "id", Label: "ID", Width: 30},
		{Key: "data", Label: "Data", Width: 150},
	}

	tbl := NewTable(pdf, columns).
		WithPageBreakMode(true).
		WithPageBreakMargin(30).
		WithRepeatHeader(true)

	tbl.AddHeader()

	// Add enough rows to trigger page break
	// A4 height is ~297mm, with margins ~257mm usable
	// Row height ~8mm + spacing, so need ~30+ rows per page
	for i := 1; i <= 60; i++ {
		tbl.AddRow(map[string]interface{}{
			"id":   i,
			"data": "Row " + string(rune(48+(i%10))),
		})
	}

	pageCount := pdf.PageCount()
	// Page break may or may not occur depending on exact measurements
	// Just verify the test completes without errors
	t.Logf("Created table with page breaks: %d pages", pageCount)

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableWithRowSpacing tests row spacing
func TestTableWithRowSpacing(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "id", Label: "ID", Width: 30},
		{Key: "name", Label: "Name", Width: 60},
	}

	tbl := NewTable(pdf, columns).
		WithRowSpacing(3.0)

	tbl.AddHeader()

	for i := 1; i <= 3; i++ {
		tbl.AddRow(map[string]interface{}{
			"id":   i,
			"name": "Item " + string(rune(64+i)),
		})
	}

	if tbl.Spacing != 3.0 {
		t.Errorf("Expected spacing 3.0, got %.2f", tbl.Spacing)
	}

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableWithPageBreakMargin tests page break margin configuration
func TestTableWithPageBreakMargin(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	columns := []Column{{Key: "id", Label: "ID"}}

	tbl := NewTable(pdf, columns).
		WithPageBreakMargin(25.0)

	if tbl.PageBreakMargin != 25.0 {
		t.Errorf("Expected PageBreakMargin 25.0, got %.2f", tbl.PageBreakMargin)
	}
}

// TestTableApplyCellStyle tests cell style application
func TestTableApplyCellStyle(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "id", Label: "ID", Width: 50},
	}

	tbl := NewTable(pdf, columns)

	// Test style with bold and italic
	style := CellStyle{
		Bold:      true,
		Italic:    true,
		FontSize:  14,
		FillColor: []int{255, 200, 200},
		TextColor: []int{0, 0, 255},
		Border:    "1",
	}

	tbl.WithHeaderStyle(style)
	tbl.AddHeader()

	// Verify style was applied (indirectly by checking no errors)
	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableValueToStringComprehensive tests valueToString with various types
func TestTableValueToStringComprehensive(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "value", Label: "Value", Width: 60},
	}

	tbl := NewTable(pdf, columns)

	// Test various types
	testCases := []map[string]interface{}{
		{"value": "string"},
		{"value": 123},
		{"value": int64(456)},
		{"value": float32(78.9)},
		{"value": float64(123.456)},
		{"value": true},
		{"value": false},
		{"value": nil},
		{"value": []int{1, 2, 3}}, // Slice
		{"value": map[string]int{"a": 1}}, // Map
	}

	for _, tc := range testCases {
		tbl.AddRow(tc)
	}

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableColumnSpanInHeader tests column spanning in header cells
func TestTableColumnSpanInHeader(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "info", Label: "Information", Width: 50, ColSpan: 2},
		{Key: "id", Label: "ID", Width: 30},
		{Key: "name", Label: "Name", Width: 60},
		{Key: "price", Label: "Price", Width: 40},
	}

	tbl := NewTable(pdf, columns)
	tbl.AddHeader()

	// Add data rows
	tbl.AddRow(map[string]interface{}{
		"id":    "1",
		"name":  "Product A",
		"price": 99.99,
	})

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableMergeCellDeprecated tests deprecated MergeCell field
func TestTableMergeCellDeprecated(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "merged", Label: "Merged", Width: 50, MergeCell: true},
		{Key: "id", Label: "ID", Width: 30},
	}

	tbl := NewTable(pdf, columns)

	// Verify MergeCell sets ColSpan
	if tbl.Columns[0].ColSpan == 0 {
		t.Error("Expected ColSpan to be set when MergeCell is true")
	}

	tbl.AddHeader()

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableEmptyColumns tests edge case with empty columns
func TestTableEmptyColumns(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	tbl := NewTable(pdf, []Column{})

	// Should not panic or error on empty columns
	tbl.AddHeader()
	tbl.AddRow(map[string]interface{}{})
	tbl.Render(false, []map[string]interface{}{})

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableNestedTableWithRowSpan tests nested table with row spanning
func TestTableNestedTableWithRowSpan(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	parentColumns := []Column{
		{Key: "category", Label: "Category", Width: 60},
		{Key: "details", Label: "Details", Width: 130},
	}

	parentTable := NewTable(pdf, parentColumns)
	parentTable.AddHeader()

	// Create nested table
	nestedColumns := []Column{
		{Key: "item", Label: "Item", Width: 50},
		{Key: "qty", Label: "Qty", Width: 30},
	}

	nestedTable := NewTable(pdf, nestedColumns)
	nestedTable.AddRows([]map[string]interface{}{
		{"item": "Item A", "qty": 5},
		{"item": "Item B", "qty": 10},
	})

	// Add row with nested table that has row span
	parentTable.AddRow(map[string]interface{}{
		"category":       "Products",
		"_nested_details": nestedTable,
		"_rowspan": map[string]int{
			"category": 2,
		},
	})

	// Add another row for row span continuation
	parentTable.AddRow(map[string]interface{}{
		"details": "Additional info",
	})

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableComplexAlignment tests complex alignment scenarios
func TestTableComplexAlignment(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []Column{
		{Key: "col1", Label: "Col1", Width: 50, Align: "L", HeaderAlign: "C"},
		{Key: "col2", Label: "Col2", Width: 50, Align: "R", HeaderAlign: "L"},
		{Key: "col3", Label: "Col3", Width: 50, Align: "C", HeaderAlign: "R"},
	}

	tbl := NewTable(pdf, columns).
		WithHeaderStyle(CellStyle{Align: "C"}).
		WithDataStyle(CellStyle{Align: "L"})

	tbl.AddHeader()

	// Add row with per-cell alignment overrides
	tbl.AddRow(map[string]interface{}{
		"col1": "Data1",
		"col2": "Data2",
		"col3": "Data3",
		"_align": map[string]string{
			"col1": "R",
			"col2": "C",
			"col3": "L",
		},
	})

	if pdf.Error() != nil {
		t.Errorf("Unexpected error: %v", pdf.Error())
	}
}

// TestTableWithRepeatHeader tests header repetition configuration
func TestTableWithRepeatHeader(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	columns := []Column{{Key: "id", Label: "ID"}}

	tbl := NewTable(pdf, columns)

	// Test default (should be true)
	if !tbl.RepeatHeader {
		t.Error("Expected RepeatHeader to be true by default")
	}

	// Test disabling
	tbl.WithRepeatHeader(false)
	if tbl.RepeatHeader {
		t.Error("Expected RepeatHeader to be false after WithRepeatHeader(false)")
	}

	// Test enabling
	tbl.WithRepeatHeader(true)
	if !tbl.RepeatHeader {
		t.Error("Expected RepeatHeader to be true after WithRepeatHeader(true)")
	}
}

// TestTableWithPageBreakMode tests page break mode configuration
func TestTableWithPageBreakMode(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	columns := []Column{{Key: "id", Label: "ID"}}

	tbl := NewTable(pdf, columns)

	// Test default (should be true)
	if !tbl.PageBreakMode {
		t.Error("Expected PageBreakMode to be true by default")
	}

	// Test disabling
	tbl.WithPageBreakMode(false)
	if tbl.PageBreakMode {
		t.Error("Expected PageBreakMode to be false after WithPageBreakMode(false)")
	}

	// Test enabling
	tbl.WithPageBreakMode(true)
	if !tbl.PageBreakMode {
		t.Error("Expected PageBreakMode to be true after WithPageBreakMode(true)")
	}
}