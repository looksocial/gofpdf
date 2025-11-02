package main_test

import (
	"bytes"
	"testing"

	"github.com/looksocial/gofpdf"
)

func TestInvoiceDemoGeneration(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Colors - Blue theme
	blueColor := []int{33, 150, 243}
	darkGray := []int{51, 51, 51}

	// Header Section
	pdf.SetFont("Arial", "B", 20)
	pdf.SetTextColor(darkGray[0], darkGray[1], darkGray[2])
	pdf.SetXY(10, 10)
	pdf.Cell(80, 8, "Sales Invoice")

	// Invoice Details
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(blueColor[0], blueColor[1], blueColor[2])
	pdf.SetXY(15, 27)
	pdf.Cell(40, 4, "INVOICE NUMBER")
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(darkGray[0], darkGray[1], darkGray[2])
	pdf.SetXY(15, 31)
	pdf.Cell(40, 4, "#9000000001")

	// Customer Information
	pdf.SetXY(10, 70)
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(blueColor[0], blueColor[1], blueColor[2])
	pdf.Cell(40, 4, "CUSTOMER NAME")
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(darkGray[0], darkGray[1], darkGray[2])
	pdf.SetXY(10, 74)
	pdf.Cell(80, 4, "Test Customer")

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}

	if buf.Len() == 0 {
		t.Error("Generated PDF should have content")
	}
}

func TestInvoicePattern1Generation(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.UseEmbeddedFonts()

	// Modern Colors
	primaryBlue := []int{30, 136, 229}
	white := []int{255, 255, 255}
	textDark := []int{33, 33, 33}

	// Header Section
	pdf.SetFillColor(primaryBlue[0], primaryBlue[1], primaryBlue[2])
	pdf.Rect(0, 0, 210, 35, "F")

	// Company name
	pdf.SetTextColor(white[0], white[1], white[2])
	pdf.SetFont("Arial", "B", 16)
	pdf.SetXY(10, 13)
	pdf.Cell(100, 8, "Test Company Ltd.")

	// Invoice Title
	pdf.SetTextColor(textDark[0], textDark[1], textDark[2])
	pdf.SetFont("Arial", "B", 24)
	pdf.SetXY(150, 40)
	pdf.CellFormat(60, 10, "INVOICE", "", 0, "R", false, 0, "")

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestInvoicePattern2Generation(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.UseEmbeddedFonts()

	// Modern Colors
	primaryTeal := []int{0, 150, 136}
	white := []int{255, 255, 255}
	textDark := []int{33, 33, 33}

	// Header Section
	currentY := 10.0
	leftMargin := 10.0

	// Logo placeholder
	pdf.SetFillColor(primaryTeal[0], primaryTeal[1], primaryTeal[2])
	pdf.Circle(leftMargin+15, currentY+15, 15, "F")
	pdf.SetTextColor(white[0], white[1], white[2])
	pdf.SetFont("Arial", "B", 10)
	pdf.SetXY(leftMargin+5, currentY+12)
	pdf.Cell(20, 6, "LOGO")

	// Company info
	pdf.SetTextColor(textDark[0], textDark[1], textDark[2])
	pdf.SetFont("Arial", "B", 14)
	pdf.SetXY(leftMargin, currentY+32)
	pdf.Cell(100, 6, "Test Company Ltd.")

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestQuotationGeneration(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.UseEmbeddedFonts()

	// Colors
	darkBlue := []int{0, 47, 95}
	white := []int{255, 255, 255}

	// Header Section
	pdf.SetFillColor(darkBlue[0], darkBlue[1], darkBlue[2])
	pdf.Rect(0, 0, 210, 40, "F")

	// Logo
	pdf.SetTextColor(white[0], white[1], white[2])
	pdf.SetFont("Arial", "B", 16)
	pdf.SetXY(10, 15)
	pdf.Cell(50, 10, "GOFPDF")

	// Quotation title
	pdf.SetFont("Arial", "B", 28)
	pdf.SetXY(110, 13)
	pdf.CellFormat(100, 10, "Quotation", "", 0, "R", false, 0, "")

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestInvoiceTableGeneration(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 9)

	// Table setup
	tableStartX := 10.0
	tableStartY := 50.0
	rowHeight := 6.0

	// Column widths
	colWidths := []float64{20, 80, 25, 35, 35}
	colHeaders := []string{"No.", "Description", "Qty", "Unit Price", "Total"}
	colAligns := []string{"C", "L", "C", "R", "R"}

	// Table header
	pdf.SetFillColor(242, 242, 242)
	pdf.SetFont("Arial", "B", 10)
	xPos := tableStartX
	for i, header := range colHeaders {
		pdf.SetXY(xPos, tableStartY)
		pdf.CellFormat(colWidths[i], rowHeight, header, "", 0, colAligns[i], true, 0, "")
		xPos += colWidths[i]
	}

	// Table data
	tableData := [][]string{
		{"1", "Test Item 1", "2", "10.00", "20.00"},
		{"2", "Test Item 2", "1", "15.00", "15.00"},
	}

	currentY := tableStartY + rowHeight
	pdf.SetFont("Arial", "", 9)
	for _, row := range tableData {
		xPos = tableStartX
		for i, cell := range row {
			pdf.SetXY(xPos, currentY)
			pdf.CellFormat(colWidths[i], rowHeight, cell, "", 0, colAligns[i], false, 0, "")
			xPos += colWidths[i]
		}
		currentY += rowHeight
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}
