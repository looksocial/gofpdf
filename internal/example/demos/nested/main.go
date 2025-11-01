package main

import (
	"os"
	"path/filepath"

	"github.com/looksocial/gofpdf"
	"github.com/looksocial/gofpdf/table"
)

func main() {
	// Create new PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)

	// Example 1: Simple nested table
	nestedTableExample(pdf)

	pdf.AddPage()

	// Example 2: Row span example
	rowSpanExample(pdf)

	pdf.AddPage()

	// Example 3: Summary/Total rows
	summaryTotalExample(pdf)

	pdf.AddPage()

	// Example 4: Complex nested table with row spans
	complexNestedExample(pdf)

	pdf.AddPage()

	// Example 5: Extreme text wrapping test
	extremeTextWrappingExample(pdf)

	// Save PDF
	// Get the output path relative to the project root
	outputPath := filepath.Join("pdf", "nested_table_demo.pdf")

	// Ensure the pdf directory exists
	pdfDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(pdfDir, 0755); err != nil {
		panic(err)
	}

	err := pdf.OutputFileAndClose(outputPath)
	if err != nil {
		panic(err)
	}
}

// Example 1: Simple nested table
func nestedTableExample(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "Example 1: Simple Nested Table")
	pdf.Ln(15)

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

	// Row 1 - with nested table in column 2
	nestedData := []map[string]interface{}{
		{"ncol1": "Inner Row 1 Col 1", "ncol2": "Inner Row 1 Col 2"},
		{"ncol1": "Inner Row 2 Col 1", "ncol2": "Inner Row 2 Col 2"},
		{"ncol1": "Inner Row 3 Col 1", "ncol2": "Inner Row 3 Col 2"},
	}
	nestedTbl.AddRows(nestedData)

	row1 := map[string]interface{}{
		"col1":         "Main Row 1 Col 1",
		"col2":         "", // Nested table will go here
		"col3":         "Main Row 1 Col 3",
		"_nested_col2": nestedTbl,
	}

	mainTbl.AddRow(row1)

	// Row 2 - regular row
	mainTbl.AddRow(map[string]interface{}{
		"col1": "Main Row 2 Col 1",
		"col2": "Main Row 2 Col 2",
		"col3": "Main Row 2 Col 3",
	})
}

// Example 2: Row span example
func rowSpanExample(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "Example 2: Row Span (Cells spanning multiple rows)")
	pdf.Ln(15)

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

	// Row 1 - item spans 3 rows
	tbl.AddRow(map[string]interface{}{
		"item":    "Item A (spans 3 rows)",
		"details": "Detail 1",
		"value1":  "100",
		"value2":  "200",
		"_rowspan": map[string]int{
			"item": 3,
		},
	})

	// Row 2 - continuation of Item A
	tbl.AddRow(map[string]interface{}{
		"details": "Detail 2",
		"value1":  "150",
		"value2":  "250",
	})

	// Row 3 - continuation of Item A
	tbl.AddRow(map[string]interface{}{
		"details": "Detail 3",
		"value1":  "200",
		"value2":  "300",
	})

	// Row 4 - new item
	tbl.AddRow(map[string]interface{}{
		"item":    "Item B",
		"details": "Detail 1",
		"value1":  "50",
		"value2":  "75",
	})
}

// Example 3: Summary/Total rows
func summaryTotalExample(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "Example 3: Summary and Total Rows")
	pdf.Ln(15)

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
		{"product": "Product C", "qty": "2", "price": "$15.00", "total": "$30.00"},
	}

	tbl.AddRows(data)

	// Summary row
	tbl.AddSummaryRow("SUB-TOTAL", 1, map[string]interface{}{
		"qty":   "10",
		"price": "",
		"total": "$140.00",
	}, table.CellStyle{
		Border:    "1",
		Bold:      true,
		FillColor: []int{240, 240, 240},
	})

	// Grand total row
	tbl.AddTotalRow("GRAND TOTAL", map[string]interface{}{
		"qty":   "",
		"price": "",
		"total": "$140.00",
	}, table.CellStyle{
		Border:    "1",
		Bold:      true,
		FillColor: []int{200, 200, 200},
		TextColor: []int{0, 0, 0},
	})
}

// Example 4: Complex nested table with row spans
func complexNestedExample(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "Example 4: Complex Nested Table with Row Spans")
	pdf.Ln(15)

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

	// Nested table 1 - with text that demonstrates wrapping
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
		{"n1": "Row 1 Header with text", "n2": "item"},
		{"n1": "Row 2 Header", "n2": "item overflow"},
		{"n1": "Row 3 Header", "n2": "item"},
	}
	nestedTbl1.AddRows(nestedData1)

	// Row 2 of main table
	mainTbl.AddRow(map[string]interface{}{
		"header1":         "Row 2 - Item 1",
		"header2":         "Row 2 - Item 2 with text that wraps",
		"header3":         "", // Nested table will go here
		"header4":         "Row 2 - Item 4\nA second line\nThird line",
		"_nested_header3": nestedTbl1,
	})

	// Nested table 2 - with longer text
	nestedCols2 := []table.Column{
		{Key: "n1", Label: "Row Header", Width: 25, Align: "L"},
		{Key: "n2", Label: "item", Width: 20, Align: "L"},
	}

	nestedTbl2 := table.NewTable(pdf, nestedCols2).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(6)

	nestedData2 := []map[string]interface{}{
		{"n1": "Row 1 Header with long text", "n2": "item data here"},
		{"n1": "Row 2 Header", "n2": "item with more data"},
	}
	nestedTbl2.AddRows(nestedData2)

	// Row 3 with row span and nested table
	mainTbl.AddRow(map[string]interface{}{
		"header1":         "", // Nested table will go here (with row span)
		"header2":         "Row 3 - Item 2 with wrap text",
		"header3":         "",
		"header4":         "Row 3 - Item 3 spans 2 rows",
		"_nested_header1": nestedTbl2,
		"_rowspan": map[string]int{
			"header1": 2, // Nested table spans 2 rows
			"header4": 2, // Item 3 spans 2 rows
		},
	})

	// Row 4
	mainTbl.AddRow(map[string]interface{}{
		"header2": "Row 4 - Item 2",
		"header3": "Row 4 - Item 3",
	})

	// Row 5 - spanning all columns
	mainTbl.AddTotalRow("Row 5 - Last row of outer table", map[string]interface{}{},
		table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{240, 240, 240},
		})
}

// Example 5: Extreme text wrapping test
func extremeTextWrappingExample(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "Example 5: Extreme Text Wrapping Test")
	pdf.Ln(15)

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

	// Create nested table with very long text
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
		{
			"n1": "Description 3",
			"n2": "Value 3",
		},
	}
	nestedTbl.AddRows(nestedData)

	// Row 1 - with nested table
	mainTbl.AddRow(map[string]interface{}{
		"col1":         "Regular text in column 1",
		"col2":         "", // Nested table
		"col3":         "Long text in column 3 that wraps properly",
		"_nested_col2": nestedTbl,
	})

	// Row 2 - all wrapped text
	mainTbl.AddRow(map[string]interface{}{
		"col1": "Column 1 with text that wraps",
		"col2": "Column 2 with wrapped text content",
		"col3": "Column 3 has wrapping text",
	})

	// Row 3 - nested table with row span
	nestedTbl2 := table.NewTable(pdf, nestedCols).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(6)

	nestedData2 := []map[string]interface{}{
		{
			"n1": "Long text A that wraps",
			"n2": "Value A",
		},
		{
			"n1": "Long text B",
			"n2": "Value B with more data",
		},
	}
	nestedTbl2.AddRows(nestedData2)

	mainTbl.AddRow(map[string]interface{}{
		"col1":         "", // Nested table with row span
		"col2":         "Regular text",
		"col3":         "Text spans 2 rows with wrap",
		"_nested_col1": nestedTbl2,
		"_rowspan": map[string]int{
			"col1": 2,
			"col3": 2,
		},
	})

	// Row 4
	mainTbl.AddRow(map[string]interface{}{
		"col2": "Row 4 column 2",
	})
}
