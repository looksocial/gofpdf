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

	// Example 1: Simple table
	simpleTable(pdf)

	// Example 2: Table with styling
	styledTable(pdf)

	// Example 3: Table with alternating rows
	alternatingTable(pdf)

	// Example 4: Header alignment (HeaderAlign)
	headerAlignmentTable(pdf)

	// Example 5: Style-level alignment
	styleAlignmentTable(pdf)

	// Example 6: Per-row alignment
	perRowAlignmentTable(pdf)

	// Example 7: Mixed alignments
	mixedAlignmentTable(pdf)

	// Example 8: Column spanning
	columnSpanTable(pdf)

	// Example 9: Auto-width columns
	autoWidthTable(pdf)

	// Example 10: Custom positioning
	customPositionTable(pdf)

	// Save PDF
	// Get the output path relative to the project root
	outputPath := filepath.Join("pdf", "table_component_demo.pdf")
	
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

// Example 1: Simple table
func simpleTable(pdf *gofpdf.Fpdf) {
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, "Example 1: Simple Table")
	pdf.Ln(15)

	// Define columns
	columns := []table.Column{
		{Key: "id", Label: "ID", Width: 20, Align: "L"},
		{Key: "name", Label: "Name", Width: 60, Align: "L"},
		{Key: "email", Label: "Email", Width: 110, Align: "L"},
	}

	// Create table
	tbl := table.NewTable(pdf, columns)

	// Data
	data := []map[string]interface{}{
		{"id": "1", "name": "Alice Johnson", "email": "alice@example.com"},
		{"id": "2", "name": "Bob Smith", "email": "bob@example.com"},
		{"id": "3", "name": "Charlie Brown", "email": "charlie@example.com"},
	}

	// Render table
	tbl.Render(true, data)
	pdf.Ln(20)
}

// Example 2: Styled table
func styledTable(pdf *gofpdf.Fpdf) {
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, "Example 2: Styled Table (Colors, Fonts, Borders)")
	pdf.Ln(15)

	// Define columns
	columns := []table.Column{
		{Key: "country", Label: "Country", Width: 40, Align: "L"},
		{Key: "capital", Label: "Capital", Width: 40, Align: "L"},
		{Key: "area", Label: "Area (sq km)", Width: 45, Align: "R"},
		{Key: "population", Label: "Population", Width: 50, Align: "R"},
	}

	// Create table with custom styling
	tbl := table.NewTable(pdf, columns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{100, 100, 200},
			TextColor: []int{255, 255, 255},
			FontSize:  13,
		}).
		WithDataStyle(table.CellStyle{
			Border:    "LR",
			TextColor: []int{0, 0, 0},
		}).
		WithRowHeight(8)

	// Data
	data := []map[string]interface{}{
		{"country": "France", "capital": "Paris", "area": "551,695", "population": "67,000,000"},
		{"country": "Germany", "capital": "Berlin", "area": "357,022", "population": "83,000,000"},
		{"country": "Italy", "capital": "Rome", "area": "301,340", "population": "60,000,000"},
	}

	// Render table
	tbl.Render(true, data)
	pdf.Ln(20)
}

// Example 3: Table with alternating rows (zebra stripes)
func alternatingTable(pdf *gofpdf.Fpdf) {
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, "Example 3: Table with Alternating Rows (Zebra Stripes)")
	pdf.Ln(15)

	// Define columns
	columns := []table.Column{
		{Key: "rank", Label: "Rank", Width: 25, Align: "C"},
		{Key: "team", Label: "Team", Width: 70, Align: "L"},
		{Key: "points", Label: "Points", Width: 35, Align: "R"},
		{Key: "wins", Label: "Wins", Width: 30, Align: "R"},
		{Key: "losses", Label: "Losses", Width: 30, Align: "R"},
	}

	// Create table with alternating rows
	tbl := table.NewTable(pdf, columns).
		WithAlternatingRows(true).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{70, 70, 70},
			TextColor: []int{255, 255, 255},
		}).
		WithDataStyle(table.CellStyle{
			Border: "LR",
		}).
		WithRowHeight(7)

	// Data
	data := []map[string]interface{}{
		{"rank": "1", "team": "Manchester City", "points": "89", "wins": "28", "losses": "3"},
		{"rank": "2", "team": "Arsenal", "points": "84", "wins": "26", "losses": "5"},
		{"rank": "3", "team": "Manchester United", "points": "75", "wins": "23", "losses": "7"},
		{"rank": "4", "team": "Newcastle", "points": "71", "wins": "19", "losses": "10"},
		{"rank": "5", "team": "Liverpool", "points": "67", "wins": "19", "losses": "12"},
	}

	// Render table
	tbl.Render(true, data)
	pdf.Ln(20)
}

// Example 4: Header alignment (HeaderAlign field)
func headerAlignmentTable(pdf *gofpdf.Fpdf) {
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, "Example 4: Header Alignment (HeaderAlign)")
	pdf.Ln(15)

	// Define columns with different header and data alignments
	columns := []table.Column{
		{Key: "id", Label: "ID", Width: 30, Align: "L", HeaderAlign: "C"},
		{Key: "product", Label: "Product", Width: 60, Align: "L", HeaderAlign: "L"},
		{Key: "price", Label: "Price", Width: 40, Align: "R", HeaderAlign: "R"},
		{Key: "stock", Label: "Stock", Width: 35, Align: "L", HeaderAlign: "C"},
	}

	// Create table
	tbl := table.NewTable(pdf, columns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{240, 240, 240},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(8)

	// Data
	data := []map[string]interface{}{
		{"id": "001", "product": "Widget A", "price": "$19.99", "stock": "50"},
		{"id": "002", "product": "Widget B", "price": "$29.99", "stock": "25"},
		{"id": "003", "product": "Widget C", "price": "$39.99", "stock": "75"},
	}

	// Render table
	tbl.Render(true, data)
	pdf.Ln(20)
}

// Example 5: Style-level alignment
func styleAlignmentTable(pdf *gofpdf.Fpdf) {
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, "Example 5: Style-Level Alignment (CellStyle.Align)")
	pdf.Ln(15)

	// Define columns
	columns := []table.Column{
		{Key: "name", Label: "Name", Width: 50, Align: "L"},
		{Key: "category", Label: "Category", Width: 50, Align: "L"},
		{Key: "value", Label: "Value", Width: 50, Align: "L"},
	}

	// Create table with style-level alignment override
	tbl := table.NewTable(pdf, columns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			Align:     "C", // All headers centered
			FillColor: []int{220, 220, 220},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
			Align:  "C", // All data cells centered
		}).
		WithRowHeight(8)

	// Data
	data := []map[string]interface{}{
		{"name": "Item A", "category": "Type 1", "value": "100"},
		{"name": "Item B", "category": "Type 2", "value": "200"},
		{"name": "Item C", "category": "Type 1", "value": "150"},
	}

	// Render table
	tbl.Render(true, data)
	pdf.Ln(20)
}

// Example 6: Per-row alignment
func perRowAlignmentTable(pdf *gofpdf.Fpdf) {
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, "Example 6: Per-Row Alignment (_align map)")
	pdf.Ln(15)

	// Define columns
	columns := []table.Column{
		{Key: "id", Label: "ID", Width: 30, Align: "L"},
		{Key: "item", Label: "Item", Width: 60, Align: "L"},
		{Key: "price", Label: "Price", Width: 40, Align: "L"},
		{Key: "qty", Label: "Quantity", Width: 40, Align: "L"},
	}

	// Create table
	tbl := table.NewTable(pdf, columns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{230, 230, 230},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(8)

	// Data with per-row alignment overrides
	data := []map[string]interface{}{
		{
			"id":    "1",
			"item":  "Product A",
			"price": "$10.50",
			"qty":   "5",
			"_align": map[string]string{
				"id":    "C",  // Center align ID
				"price": "R",  // Right align price
				"qty":   "C",  // Center align quantity
			},
		},
		{
			"id":    "2",
			"item":  "Product B",
			"price": "$25.00",
			"qty":   "10",
			"_align": map[string]string{
				"price": "R",
				"qty":   "C",
			},
		},
		{"id": "3", "item": "Product C", "price": "$5.75", "qty": "3"}, // Default alignment
	}

	// Render table
	tbl.Render(true, data)
	pdf.Ln(20)
}

// Example 7: Mixed alignments (all alignment features combined)
func mixedAlignmentTable(pdf *gofpdf.Fpdf) {
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, "Example 7: Mixed Alignments (All Features Combined)")
	pdf.Ln(15)

	// Define columns with HeaderAlign
	columns := []table.Column{
		{Key: "id", Label: "ID", Width: 25, Align: "L", HeaderAlign: "C"},
		{Key: "product", Label: "Product", Width: 55, Align: "L", HeaderAlign: "L"},
		{Key: "price", Label: "Price", Width: 40, Align: "R", HeaderAlign: "R"},
		{Key: "qty", Label: "Qty", Width: 30, Align: "L", HeaderAlign: "C"},
		{Key: "total", Label: "Total", Width: 40, Align: "R", HeaderAlign: "R"},
	}

	// Create table with style-level alignment
	tbl := table.NewTable(pdf, columns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{200, 200, 200},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
			Align:  "L", // Default data alignment (can be overridden per-row)
		}).
		WithRowHeight(8)

	// Data with per-row alignment overrides
	data := []map[string]interface{}{
		{
			"id":      "1",
			"product": "Widget A",
			"price":   "$19.99",
			"qty":     "5",
			"total":   "$99.95",
			"_align": map[string]string{
				"id":    "C",
				"price": "R",
				"qty":   "C",
				"total": "R",
			},
		},
		{
			"id":      "2",
			"product": "Widget B",
			"price":   "$29.99",
			"qty":     "3",
			"total":   "$89.97",
			"_align": map[string]string{
				"qty": "C",
			},
		},
		{"id": "3", "product": "Widget C", "price": "$39.99", "qty": "2", "total": "$79.98"}, // Default alignment
	}

	// Render table
	tbl.Render(true, data)
	pdf.Ln(20)
}

// Example 8: Column spanning
func columnSpanTable(pdf *gofpdf.Fpdf) {
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, "Example 8: Column Spanning")
	pdf.Ln(15)

	// Define columns with column spanning
	columns := []table.Column{
		{Key: "quarter1", Label: "Q1 2024", Width: 40, Align: "C", ColSpan: 2},
		{Key: "q1sales", Label: "", Width: 0}, // Skipped (part of Q1 span)
		{Key: "q1profit", Label: "Sales", Width: 0}, // Skipped (part of Q1 span)
		{Key: "quarter2", Label: "Q2 2024", Width: 40, Align: "C", ColSpan: 2},
		{Key: "q2sales", Label: "", Width: 0}, // Skipped (part of Q2 span)
		{Key: "q2profit", Label: "Sales", Width: 0}, // Skipped (part of Q2 span)
	}

	// Create table
	tbl := table.NewTable(pdf, columns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{180, 180, 200},
			TextColor: []int{255, 255, 255},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
			Align:  "C",
		}).
		WithRowHeight(8)

	// Render header only (column spanning is mainly for headers)
	tbl.AddHeader()

	// Note about column spanning
	pdf.Ln(15)
	pdf.SetFont("Arial", "I", 10)
	pdf.Cell(0, 5, "Note: Column spanning allows headers to span multiple columns. Data cells still map to individual column keys.")
	pdf.Ln(10)
}

// Example 9: Auto-width columns
func autoWidthTable(pdf *gofpdf.Fpdf) {
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, "Example 9: Auto-Width Columns (Width: 0)")
	pdf.Ln(15)

	// Define columns with auto-width (Width: 0)
	columns := []table.Column{
		{Key: "name", Label: "Name", Width: 0},      // Auto-calculated
		{Key: "email", Label: "Email", Width: 100}, // Fixed width
		{Key: "phone", Label: "Phone", Width: 0},   // Auto-calculated
	}

	// Create table
	tbl := table.NewTable(pdf, columns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{240, 240, 240},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(8)

	// Data
	data := []map[string]interface{}{
		{"name": "Alice", "email": "alice@example.com", "phone": "555-0101"},
		{"name": "Bob", "email": "bob@example.com", "phone": "555-0102"},
		{"name": "Charlie", "email": "charlie@example.com", "phone": "555-0103"},
	}

	// Render table
	tbl.Render(true, data)
	pdf.Ln(20)
}

// Example 10: Custom positioning
func customPositionTable(pdf *gofpdf.Fpdf) {
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, "Example 10: Custom Positioning")
	pdf.Ln(30)

	// Add some text before the table
	pdf.SetFont("Arial", "I", 10)
	pdf.Cell(0, 8, "Table positioned at custom coordinates (X: 50mm, Y: 50mm)")
	pdf.Ln(5)

	// Define columns
	columns := []table.Column{
		{Key: "item", Label: "Item", Width: 50, Align: "L"},
		{Key: "status", Label: "Status", Width: 50, Align: "C"},
	}

	// Create table with custom starting position
	tbl := table.NewTable(pdf, columns).
		WithStartPosition(50, 50). // X: 50mm, Y: 50mm from top
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{250, 200, 200},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(8)

	// Data
	data := []map[string]interface{}{
		{"item": "Task A", "status": "Done"},
		{"item": "Task B", "status": "In Progress"},
		{"item": "Task C", "status": "Pending"},
	}

	// Render table
	tbl.Render(true, data)
	pdf.Ln(20)
}
