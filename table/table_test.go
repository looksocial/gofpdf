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
