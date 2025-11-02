package main_test

import (
	"bytes"
	"testing"

	"github.com/looksocial/gofpdf"
	"github.com/looksocial/gofpdf/table"
)

func TestSimpleTable(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []table.Column{
		{Key: "id", Label: "ID", Width: 20, Align: "L"},
		{Key: "name", Label: "Name", Width: 60, Align: "L"},
		{Key: "email", Label: "Email", Width: 110, Align: "L"},
	}

	tbl := table.NewTable(pdf, columns)

	data := []map[string]interface{}{
		{"id": "1", "name": "Alice Johnson", "email": "alice@example.com"},
		{"id": "2", "name": "Bob Smith", "email": "bob@example.com"},
		{"id": "3", "name": "Charlie Brown", "email": "charlie@example.com"},
	}

	tbl.Render(true, data)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestStyledTable(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []table.Column{
		{Key: "country", Label: "Country", Width: 40, Align: "L"},
		{Key: "capital", Label: "Capital", Width: 40, Align: "L"},
		{Key: "area", Label: "Area (sq km)", Width: 45, Align: "R"},
		{Key: "population", Label: "Population", Width: 50, Align: "R"},
	}

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

	data := []map[string]interface{}{
		{"country": "France", "capital": "Paris", "area": "551,695", "population": "67,000,000"},
		{"country": "Germany", "capital": "Berlin", "area": "357,022", "population": "83,000,000"},
	}

	tbl.Render(true, data)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestAlternatingTable(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []table.Column{
		{Key: "rank", Label: "Rank", Width: 25, Align: "C"},
		{Key: "team", Label: "Team", Width: 70, Align: "L"},
		{Key: "points", Label: "Points", Width: 35, Align: "R"},
		{Key: "wins", Label: "Wins", Width: 30, Align: "R"},
		{Key: "losses", Label: "Losses", Width: 30, Align: "R"},
	}

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

	data := []map[string]interface{}{
		{"rank": "1", "team": "Team A", "points": "89", "wins": "28", "losses": "3"},
		{"rank": "2", "team": "Team B", "points": "84", "wins": "26", "losses": "5"},
		{"rank": "3", "team": "Team C", "points": "75", "wins": "23", "losses": "7"},
	}

	tbl.Render(true, data)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestHeaderAlignmentTable(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []table.Column{
		{Key: "id", Label: "ID", Width: 30, Align: "L", HeaderAlign: "C"},
		{Key: "product", Label: "Product", Width: 60, Align: "L", HeaderAlign: "L"},
		{Key: "price", Label: "Price", Width: 40, Align: "R", HeaderAlign: "R"},
		{Key: "stock", Label: "Stock", Width: 35, Align: "L", HeaderAlign: "C"},
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
		WithRowHeight(8)

	data := []map[string]interface{}{
		{"id": "001", "product": "Widget A", "price": "$19.99", "stock": "50"},
		{"id": "002", "product": "Widget B", "price": "$29.99", "stock": "25"},
	}

	tbl.Render(true, data)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestStyleAlignmentTable(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []table.Column{
		{Key: "name", Label: "Name", Width: 50, Align: "L"},
		{Key: "category", Label: "Category", Width: 50, Align: "L"},
		{Key: "value", Label: "Value", Width: 50, Align: "L"},
	}

	tbl := table.NewTable(pdf, columns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			Align:     "C",
			FillColor: []int{220, 220, 220},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
			Align:  "C",
		}).
		WithRowHeight(8)

	data := []map[string]interface{}{
		{"name": "Item A", "category": "Type 1", "value": "100"},
		{"name": "Item B", "category": "Type 2", "value": "200"},
	}

	tbl.Render(true, data)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestPerRowAlignmentTable(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []table.Column{
		{Key: "id", Label: "ID", Width: 30, Align: "L"},
		{Key: "item", Label: "Item", Width: 60, Align: "L"},
		{Key: "price", Label: "Price", Width: 40, Align: "L"},
		{Key: "qty", Label: "Quantity", Width: 40, Align: "L"},
	}

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

	data := []map[string]interface{}{
		{
			"id":    "1",
			"item":  "Product A",
			"price": "$10.50",
			"qty":   "5",
			"_align": map[string]string{
				"id":    "C",
				"price": "R",
				"qty":   "C",
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
	}

	tbl.Render(true, data)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestMixedAlignmentTable(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []table.Column{
		{Key: "id", Label: "ID", Width: 25, Align: "L", HeaderAlign: "C"},
		{Key: "product", Label: "Product", Width: 55, Align: "L", HeaderAlign: "L"},
		{Key: "price", Label: "Price", Width: 40, Align: "R", HeaderAlign: "R"},
		{Key: "qty", Label: "Qty", Width: 30, Align: "L", HeaderAlign: "C"},
		{Key: "total", Label: "Total", Width: 40, Align: "R", HeaderAlign: "R"},
	}

	tbl := table.NewTable(pdf, columns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{200, 200, 200},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
			Align:  "L",
		}).
		WithRowHeight(8)

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
	}

	tbl.Render(true, data)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestCustomPositionTable(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []table.Column{
		{Key: "item", Label: "Item", Width: 50, Align: "L"},
		{Key: "status", Label: "Status", Width: 50, Align: "C"},
	}

	tbl := table.NewTable(pdf, columns).
		WithStartPosition(50, 50).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{250, 200, 200},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(8)

	data := []map[string]interface{}{
		{"item": "Task A", "status": "Done"},
		{"item": "Task B", "status": "In Progress"},
		{"item": "Task C", "status": "Pending"},
	}

	tbl.Render(true, data)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestAutoWidthTable(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	columns := []table.Column{
		{Key: "name", Label: "Name", Width: 0},
		{Key: "email", Label: "Email", Width: 100},
		{Key: "phone", Label: "Phone", Width: 0},
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
		WithRowHeight(8)

	data := []map[string]interface{}{
		{"name": "Alice", "email": "alice@example.com", "phone": "555-0101"},
		{"name": "Bob", "email": "bob@example.com", "phone": "555-0102"},
	}

	tbl.Render(true, data)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}
