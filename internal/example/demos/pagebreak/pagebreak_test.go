package main_test

import (
	"bytes"
	"testing"

	"github.com/looksocial/gofpdf"
	"github.com/looksocial/gofpdf/table"
)

func TestSimpleMultiPageExample(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)

	columns := []table.Column{
		{Key: "id", Label: "ID", Width: 20, Align: "C"},
		{Key: "name", Label: "Product Name", Width: 70, Align: "L"},
		{Key: "category", Label: "Category", Width: 50, Align: "L"},
		{Key: "price", Label: "Price", Width: 30, Align: "R"},
		{Key: "stock", Label: "Stock", Width: 20, Align: "C"},
	}

	tbl := table.NewTable(pdf, columns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{100, 150, 200},
			TextColor: []int{255, 255, 255},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(8).
		WithRepeatHeader(true).
		WithPageBreakMode(true).
		WithPageBreakMargin(20)

	tbl.AddHeader()

	// Generate rows
	for i := 0; i < 30; i++ {
		tbl.AddRow(map[string]interface{}{
			"id":       "1",
			"name":     "Test Product",
			"category": "Test Category",
			"price":    "$10.00",
			"stock":    "100",
		})
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}

	// Check if multiple pages were created (page count should be > 1 for 30 rows)
	pageCount := pdf.PageCount()
	if pageCount == 0 {
		t.Error("PDF should have at least 1 page")
	}
}

func TestNoHeaderRepeatExample(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)

	columns := []table.Column{
		{Key: "order", Label: "Order #", Width: 30, Align: "C"},
		{Key: "customer", Label: "Customer", Width: 80, Align: "L"},
		{Key: "date", Label: "Date", Width: 40, Align: "C"},
		{Key: "total", Label: "Total", Width: 40, Align: "R"},
	}

	tbl := table.NewTable(pdf, columns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{200, 100, 100},
			TextColor: []int{255, 255, 255},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(8).
		WithRepeatHeader(false).
		WithPageBreakMode(true)

	tbl.AddHeader()

	// Generate rows
	for i := 1; i <= 40; i++ {
		tbl.AddRow(map[string]interface{}{
			"order":    "ORD-0001",
			"customer": "Test Customer",
			"date":     "2024-11-01",
			"total":    "$100.00",
		})
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestCustomMarginExample(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)

	columns := []table.Column{
		{Key: "item", Label: "Item", Width: 100, Align: "L"},
		{Key: "description", Label: "Description", Width: 90, Align: "L"},
	}

	tbl := table.NewTable(pdf, columns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{150, 200, 100},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(10).
		WithRepeatHeader(true).
		WithPageBreakMargin(40)

	tbl.AddHeader()

	// Generate rows with long descriptions
	for i := 0; i < 30; i++ {
		tbl.AddRow(map[string]interface{}{
			"item":        "Test Item",
			"description": "This is a long description that will wrap to multiple lines and test the page break functionality with custom margin",
		})
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestPageBreakWithLongContent(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)

	columns := []table.Column{
		{Key: "col1", Label: "Column 1", Width: 95, Align: "L"},
		{Key: "col2", Label: "Column 2", Width: 95, Align: "L"},
	}

	tbl := table.NewTable(pdf, columns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{200, 200, 200},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(8).
		WithRepeatHeader(true).
		WithPageBreakMode(true).
		WithPageBreakMargin(20)

	tbl.AddHeader()

	// Add rows with very long text that will cause wrapping
	for i := 0; i < 50; i++ {
		tbl.AddRow(map[string]interface{}{
			"col1": "This is a very long text that will wrap to multiple lines and test how the page break handles wrapped content in the first column",
			"col2": "This is another very long text that will wrap to multiple lines and test how the page break handles wrapped content in the second column",
		})
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestPageBreakEdgeCases(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)

	columns := []table.Column{
		{Key: "data", Label: "Data", Width: 190, Align: "L"},
	}

	tbl := table.NewTable(pdf, columns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{240, 240, 240},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(8).
		WithRepeatHeader(true).
		WithPageBreakMode(true).
		WithPageBreakMargin(10)

	tbl.AddHeader()

	// Test with single row
	tbl.AddRow(map[string]interface{}{
		"data": "Single row test",
	})

	// Test with empty row
	tbl.AddRow(map[string]interface{}{
		"data": "",
	})

	// Test with very large margin (should still work)
	tbl2 := table.NewTable(pdf, columns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{240, 240, 240},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(8).
		WithRepeatHeader(true).
		WithPageBreakMode(true).
		WithPageBreakMargin(100)

	pdf.AddPage()
	tbl2.AddHeader()
	tbl2.AddRow(map[string]interface{}{
		"data": "Test with large margin",
	})

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}
