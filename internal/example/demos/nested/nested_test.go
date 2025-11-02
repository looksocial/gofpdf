package main_test

import (
	"bytes"
	"testing"

	"github.com/looksocial/gofpdf"
	"github.com/looksocial/gofpdf/table"
)

func TestNestedTableExample(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)

	// Main table columns
	mainColumns := []table.Column{
		{Key: "col1", Label: "Main Table Column 1", Width: 60, Align: "L"},
		{Key: "col2", Label: "Main Table Column 2", Width: 80, Align: "L"},
		{Key: "col3", Label: "Main Table Column 3", Width: 50, Align: "L"},
	}

	mainTbl := table.NewTable(pdf, mainColumns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{200, 220, 200},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(8)

	// Create nested table
	nestedColumns := []table.Column{
		{Key: "ncol1", Label: "Inner Col 1", Width: 30, Align: "L"},
		{Key: "ncol2", Label: "Inner Col 2", Width: 35, Align: "L"},
	}

	nestedTbl := table.NewTable(pdf, nestedColumns).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(6)

	// Render main table header
	mainTbl.AddHeader()

	// Row with nested table
	nestedData := []map[string]interface{}{
		{"ncol1": "Inner Row 1 Col 1", "ncol2": "Inner Row 1 Col 2"},
		{"ncol1": "Inner Row 2 Col 1", "ncol2": "Inner Row 2 Col 2"},
	}
	nestedTbl.AddRows(nestedData)

	row1 := map[string]interface{}{
		"col1":         "Main Row 1 Col 1",
		"col2":         "",
		"col3":         "Main Row 1 Col 3",
		"_nested_col2": nestedTbl,
	}

	mainTbl.AddRow(row1)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestRowSpanExample(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)

	columns := []table.Column{
		{Key: "item", Label: "Item", Width: 50, Align: "L"},
		{Key: "details", Label: "Details", Width: 60, Align: "L"},
		{Key: "value1", Label: "Value 1", Width: 40, Align: "R"},
		{Key: "value2", Label: "Value 2", Width: 40, Align: "R"},
	}

	tbl := table.NewTable(pdf, columns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{220, 220, 220},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(8)

	tbl.AddHeader()

	// Row with row span
	tbl.AddRow(map[string]interface{}{
		"item":    "Item A (spans 3 rows)",
		"details": "Detail 1",
		"value1":  "100",
		"value2":  "200",
		"_rowspan": map[string]int{
			"item": 3,
		},
	})

	// Continuation rows
	tbl.AddRow(map[string]interface{}{
		"details": "Detail 2",
		"value1":  "150",
		"value2":  "250",
	})

	tbl.AddRow(map[string]interface{}{
		"details": "Detail 3",
		"value1":  "200",
		"value2":  "300",
	})

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestSummaryTotalExample(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)

	columns := []table.Column{
		{Key: "product", Label: "Product", Width: 60, Align: "L"},
		{Key: "qty", Label: "Qty", Width: 30, Align: "C"},
		{Key: "price", Label: "Price", Width: 40, Align: "R"},
		{Key: "total", Label: "Total", Width: 40, Align: "R"},
	}

	tbl := table.NewTable(pdf, columns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{180, 180, 200},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(8)

	tbl.AddHeader()

	// Data rows
	data := []map[string]interface{}{
		{"product": "Product A", "qty": "5", "price": "$10.00", "total": "$50.00"},
		{"product": "Product B", "qty": "3", "price": "$20.00", "total": "$60.00"},
	}

	tbl.AddRows(data)

	// Summary row
	tbl.AddSummaryRow("SUB-TOTAL", 1, map[string]interface{}{
		"qty":   "8",
		"price": "",
		"total": "$110.00",
	}, table.CellStyle{
		Border:    "1",
		Bold:      true,
		FillColor: []int{240, 240, 240},
	})

	// Grand total row
	tbl.AddTotalRow("GRAND TOTAL", map[string]interface{}{
		"qty":   "",
		"price": "",
		"total": "$110.00",
	}, table.CellStyle{
		Border:    "1",
		Bold:      true,
		FillColor: []int{200, 200, 200},
	})

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestComplexNestedExample(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)

	// Main table
	mainColumns := []table.Column{
		{Key: "header1", Label: "Header Column 1", Width: 40, Align: "L"},
		{Key: "header2", Label: "Header Column 2", Width: 50, Align: "L"},
		{Key: "header3", Label: "Header Column 3", Width: 50, Align: "L"},
		{Key: "header4", Label: "Header Column 4", Width: 50, Align: "L"},
	}

	mainTbl := table.NewTable(pdf, mainColumns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{220, 220, 220},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(8)

	mainTbl.AddHeader()

	// Nested table
	nestedCols1 := []table.Column{
		{Key: "n1", Label: "Row Header", Width: 25, Align: "L"},
		{Key: "n2", Label: "item", Width: 20, Align: "L"},
	}

	nestedTbl1 := table.NewTable(pdf, nestedCols1).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(6)

	nestedData1 := []map[string]interface{}{
		{"n1": "Row 1 Header", "n2": "item"},
		{"n1": "Row 2 Header", "n2": "item"},
	}
	nestedTbl1.AddRows(nestedData1)

	// Row with nested table and row span
	mainTbl.AddRow(map[string]interface{}{
		"header1":         "Row 2 - Item 1",
		"header2":         "Row 2 - Item 2",
		"header3":         "",
		"header4":         "Row 2 - Item 4",
		"_nested_header3": nestedTbl1,
	})

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestExtremeTextWrappingExample(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)

	// Main table
	mainColumns := []table.Column{
		{Key: "col1", Label: "Column 1", Width: 50, Align: "L"},
		{Key: "col2", Label: "Column 2 (Nested)", Width: 80, Align: "L"},
		{Key: "col3", Label: "Column 3", Width: 60, Align: "L"},
	}

	mainTbl := table.NewTable(pdf, mainColumns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{220, 220, 220},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(8)

	mainTbl.AddHeader()

	// Create nested table with long text
	nestedCols := []table.Column{
		{Key: "n1", Label: "Description", Width: 40, Align: "L"},
		{Key: "n2", Label: "Value", Width: 35, Align: "L"},
	}

	nestedTbl := table.NewTable(pdf, nestedCols).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(6)

	nestedData := []map[string]interface{}{
		{
			"n1": "Long description text that wraps",
			"n2": "Value 1",
		},
		{
			"n1": "Short desc",
			"n2": "Longer value text here",
		},
	}
	nestedTbl.AddRows(nestedData)

	// Row with nested table
	mainTbl.AddRow(map[string]interface{}{
		"col1":         "Regular text in column 1",
		"col2":         "",
		"col3":         "Long text in column 3 that wraps properly",
		"_nested_col2": nestedTbl,
	})

	// Row with all wrapped text
	mainTbl.AddRow(map[string]interface{}{
		"col1": "Column 1 with text that wraps",
		"col2": "Column 2 with wrapped text content",
		"col3": "Column 3 has wrapping text",
	})

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}
