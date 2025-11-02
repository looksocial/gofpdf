//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/looksocial/gofpdf"
)

func main() {
	// Create new PDF
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Use embedded fonts (Arial works with or without this, but kept for compatibility)
	pdf.UseEmbeddedFonts()

	// Add a page
	pdf.AddPage()

	// Generate the quotation
	generateQuotation(pdf)

	// Ensure pdf directory exists
	pdfDir := filepath.Join("..", "..", "..", "..", "pdf")
	err := os.MkdirAll(pdfDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Save to file in pdf directory
	outputPath := filepath.Join(pdfDir, "quotation.pdf")
	err = pdf.OutputFileAndClose(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	// Get absolute path for display
	absPath, _ := filepath.Abs(outputPath)
	fmt.Printf("✓ Quotation created successfully: %s\n", absPath)
}

func generateQuotation(pdf *gofpdf.Fpdf) {
	// Colors
	darkBlue := []int{0, 47, 95}      // #002F5F
	white := []int{255, 255, 255}     // #FFFFFF
	lightGray := []int{242, 242, 242} // #F2F2F2
	beige := []int{245, 239, 231}     // #F5EFE7

	// Page width and margins
	pageWidth, _ := pdf.GetPageSize()
	leftMargin, topMargin, rightMargin, _ := pdf.GetMargins()

	// ===== HEADER SECTION =====
	// Blue header background
	pdf.SetFillColor(darkBlue[0], darkBlue[1], darkBlue[2])
	pdf.Rect(0, 0, pageWidth, 40, "F")

	// Logo placeholder (text-based)
	pdf.SetTextColor(white[0], white[1], white[2])
	pdf.SetFont("Arial", "B", 16)
	pdf.SetXY(leftMargin, topMargin+5)
	pdf.Cell(50, 10, "GOFPDF")

	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin, topMargin+14)
	pdf.Cell(50, 5, "REAL ESTATE AGENT")

	// Quotation title on the right
	pdf.SetFont("Arial", "B", 28)
	pdf.SetXY(pageWidth-rightMargin-100, topMargin+3)
	pdf.CellFormat(100, 10, "Quotation", "", 0, "R", false, 0, "")

	// Reset text color to black
	pdf.SetTextColor(0, 0, 0)
	pdf.SetY(45) // Move below header

	// ===== BILLING INFORMATION SECTION =====
	currentY := 48.0

	// Left side - Billed to
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(leftMargin, currentY)
	pdf.Cell(40, 5, "Customer Name")
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(leftMargin + 40)
	pdf.Cell(70, 5, "Salford & Co.")

	// Right side - Quotation details
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(pageWidth-rightMargin-90, currentY)
	pdf.Cell(30, 5, "Quotation No.")
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(pageWidth - rightMargin - 60)
	pdf.Cell(60, 5, "01234")

	currentY += 5

	// Address
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(leftMargin, currentY)
	pdf.Cell(40, 5, "Address")
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(leftMargin + 40)
	pdf.Cell(70, 5, "123 Anywhere St., Any City, ST 12345")

	// Date
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(pageWidth-rightMargin-90, currentY)
	pdf.Cell(30, 5, "Date")
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(pageWidth - rightMargin - 60)
	pdf.Cell(60, 5, "January 1, 2026")

	currentY += 5

	// Tax ID
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(leftMargin, currentY)
	pdf.Cell(40, 5, "Tax ID")
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(leftMargin + 40)
	pdf.Cell(70, 5, "1234567890")

	// Due date
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(pageWidth-rightMargin-90, currentY)
	pdf.Cell(30, 5, "Due Date")
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(pageWidth - rightMargin - 60)
	pdf.Cell(60, 5, "February 1, 2026")

	currentY += 5

	// Contact
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(leftMargin, currentY)
	pdf.Cell(40, 5, "Contact")
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(leftMargin + 40)
	pdf.Cell(70, 5, "Benjamin Hart")

	// Reference
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(pageWidth-rightMargin-90, currentY)
	pdf.Cell(30, 5, "Reference")

	currentY += 8

	// ===== BILLING FROM SECTION =====
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(leftMargin, currentY)
	pdf.Cell(40, 5, "Issued By")
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(leftMargin + 40)
	pdf.Cell(70, 5, "Wardiere Inc.")

	// Tax ID
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(pageWidth-rightMargin-90, currentY)
	pdf.Cell(30, 5, "Tax ID")
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(pageWidth - rightMargin - 60)
	pdf.Cell(60, 5, "1234567890")

	currentY += 5

	// Address
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(leftMargin, currentY)
	pdf.Cell(40, 5, "Address")
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(leftMargin + 40)
	pdf.Cell(70, 5, "123 Anywhere St., Any City, ST 12345")

	// Phone
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(pageWidth-rightMargin-90, currentY)
	pdf.Cell(30, 5, "Phone")
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(pageWidth - rightMargin - 60)
	pdf.Cell(60, 5, "123-456-7890")

	currentY += 5

	// Email
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(pageWidth-rightMargin-90, currentY)
	pdf.Cell(30, 5, "Email")
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(pageWidth - rightMargin - 60)
	pdf.Cell(60, 5, "hello@reallygreatsite.com")

	currentY += 8

	// ===== ITEMS TABLE =====
	pdf.SetY(currentY)
	pdf.SetFont("Arial", "", 9)

	// Table setup
	tableStartX := leftMargin
	tableStartY := currentY
	rowHeight := 6.0

	// Column widths
	colWidths := []float64{20, 80, 25, 35, 35}
	colHeaders := []string{"No.", "Description", "Qty", "Unit Price", "Total"}
	colAligns := []string{"C", "L", "C", "R", "R"}

	// Table data
	tableData := [][]string{
		{"1", "Building A", "1", "1,200,000.00", "1,200,000.00"},
		{"2", "Interior Design", "1", "300,000.00", "300,000.00"},
		{"3", "Processing Fee", "1", "150,000.00", "150,000.00"},
	}

	// Calculate total table width
	totalTableWidth := 0.0
	for _, w := range colWidths {
		totalTableWidth += w
	}

	// Draw table header
	pdf.SetFillColor(lightGray[0], lightGray[1], lightGray[2])
	pdf.SetFont("Arial", "B", 10)
	xPos := tableStartX
	for i, header := range colHeaders {
		pdf.SetXY(xPos, tableStartY)
		pdf.CellFormat(colWidths[i], rowHeight, header, "", 0, colAligns[i], true, 0, "")
		xPos += colWidths[i]
	}

	// Draw header bottom border
	pdf.SetLineWidth(0.2)
	pdf.Line(tableStartX, tableStartY+rowHeight, tableStartX+totalTableWidth, tableStartY+rowHeight)

	currentY = tableStartY + rowHeight

	// Draw data rows
	pdf.SetFillColor(white[0], white[1], white[2])
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

	// Draw table borders
	pdf.SetLineWidth(0.2)

	// Outer border
	pdf.Rect(tableStartX, tableStartY, totalTableWidth, rowHeight*float64(len(tableData)+1), "D")

	// Vertical lines
	xPos = tableStartX
	for i := 0; i < len(colWidths); i++ {
		xPos += colWidths[i]
		if i < len(colWidths)-1 { // Don't draw after last column
			pdf.Line(xPos, tableStartY, xPos, tableStartY+rowHeight*float64(len(tableData)+1))
		}
	}

	// Horizontal lines (between data rows)
	yPos := tableStartY + rowHeight
	for i := 0; i < len(tableData); i++ {
		yPos += rowHeight
		if i < len(tableData)-1 { // Don't draw after last row (outer border handles it)
			pdf.Line(tableStartX, yPos, tableStartX+totalTableWidth, yPos)
		}
	}

	currentY = tableStartY + rowHeight*float64(len(tableData)+1) + 3

	// ===== SUMMARY SECTION =====
	summaryX := pageWidth - rightMargin - 100

	// Notes section on the left
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(leftMargin, currentY)
	pdf.Cell(40, 5, "Notes")

	// Summary on the right
	pdf.SetFont("Arial", "", 9)
	pdf.SetXY(summaryX, currentY)
	pdf.Cell(60, 5, "Subtotal")
	pdf.SetX(summaryX + 60)
	pdf.CellFormat(40, 5, "1,650,000.00", "", 0, "R", false, 0, "")

	currentY += 5

	pdf.SetXY(summaryX, currentY)
	pdf.Cell(60, 5, "VAT (7%)")
	pdf.SetX(summaryX + 60)
	pdf.CellFormat(40, 5, "115,500.00", "", 0, "R", false, 0, "")

	currentY += 5

	pdf.SetXY(summaryX, currentY)
	pdf.Cell(60, 5, "Discount")
	pdf.SetX(summaryX + 60)
	pdf.CellFormat(40, 5, "0.00", "", 0, "R", false, 0, "")

	currentY += 6

	// Line separator
	pdf.SetLineWidth(0.5)
	pdf.Line(summaryX, currentY, summaryX+100, currentY)

	currentY += 5

	// Total - with beige background
	pdf.SetFillColor(beige[0], beige[1], beige[2])
	pdf.SetFont("Arial", "B", 10)
	pdf.SetXY(summaryX, currentY)
	pdf.CellFormat(60, 6, "Total Amount", "", 0, "L", true, 0, "")
	pdf.SetX(summaryX + 60)
	pdf.CellFormat(40, 6, "1,765,500.00", "", 0, "R", true, 0, "")

	currentY += 10

	// ===== TERMS AND CONDITIONS =====
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(leftMargin, currentY)
	pdf.Cell(40, 5, "Terms & Conditions")

	// Approval section on the right
	pdf.SetXY(summaryX, currentY)
	pdf.CellFormat(100, 5, "Quotation Issuer", "", 0, "C", false, 0, "")

	currentY += 6

	pdf.SetFont("Arial", "", 7)
	pdf.SetXY(leftMargin, currentY)
	termsText := `• Lorem ipsum dolor sit amet, consectetur adipiscing elit.
  Phasellus at egestas odio. Vestibulum ante ipsum primis in
  faucibus orci luctus et ultrices posuere cubilia curae;
• Phasellus congue metus quis vehicula ultrices. Fusce at
  tristique lacus. Nullam sit amet lobortis sem, ut luctus odio.
  Duis semper odio et odio bibendum aliquam.`

	pdf.MultiCell(90, 3, termsText, "", "L", false)

	// Signature lines for issuer
	pdf.SetXY(summaryX, currentY+18)
	pdf.CellFormat(100, 4, "(Signature)", "", 0, "C", false, 0, "")

	pdf.SetXY(summaryX, currentY+22)
	pdf.CellFormat(100, 4, "Sales Representative", "", 0, "C", false, 0, "")

	pdf.SetXY(summaryX, currentY+26)
	pdf.Cell(20, 4, "Date")
	pdf.SetX(summaryX + 25)
	pdf.Cell(60, 4, ".........................")

	currentY += 32

	// ===== RECIPIENT SECTION =====
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(summaryX, currentY)
	pdf.CellFormat(100, 5, "Recipient", "", 0, "C", false, 0, "")

	currentY += 6

	// Signature lines for recipient
	pdf.SetXY(summaryX, currentY+18)
	pdf.CellFormat(100, 4, "(Signature)", "", 0, "C", false, 0, "")

	pdf.SetXY(summaryX, currentY+22)
	pdf.CellFormat(100, 4, "Manager", "", 0, "C", false, 0, "")

	pdf.SetXY(summaryX, currentY+26)
	pdf.Cell(20, 4, "Date")
	pdf.SetX(summaryX + 25)
	pdf.Cell(60, 4, ".........................")
}
