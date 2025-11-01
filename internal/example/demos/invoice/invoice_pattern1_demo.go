//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/looksocial/gofpdf"
	"github.com/looksocial/gofpdf/contrib/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	// Create new PDF
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Use embedded fonts for Thai language support
	pdf.UseEmbeddedFonts()

	// Add a page
	pdf.AddPage()

	// Generate the invoice
	generateInvoicePattern1(pdf)

	// Ensure pdf directory exists
	pdfDir := filepath.Join("..", "..", "..", "..", "pdf")
	err := os.MkdirAll(pdfDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Save to file in pdf directory
	outputPath := filepath.Join(pdfDir, "invoice_pattern1.pdf")
	err = pdf.OutputFileAndClose(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	// Get absolute path for display
	absPath, _ := filepath.Abs(outputPath)
	fmt.Printf("✓ Invoice Pattern 1 created successfully: %s\n", absPath)
}

func generateInvoicePattern1(pdf *gofpdf.Fpdf) {
	// Modern Colors
	primaryBlue := []int{30, 136, 229}  // Modern blue (#1E88E5)
	lightBlue := []int{227, 242, 253}   // Light blue accent
	white := []int{255, 255, 255}       // White
	textDark := []int{33, 33, 33}       // Dark text

	// Page width and margins
	pageWidth, _ := pdf.GetPageSize()
	leftMargin, topMargin, rightMargin, _ := pdf.GetMargins()

	// ===== HEADER SECTION =====
	// Modern blue header background
	pdf.SetFillColor(primaryBlue[0], primaryBlue[1], primaryBlue[2])
	pdf.Rect(0, 0, pageWidth, 35, "F")

	// Company name
	pdf.SetTextColor(white[0], white[1], white[2])
	pdf.SetFont("Sarabun", "B", 16)
	pdf.SetXY(leftMargin, topMargin+3)
	pdf.Cell(100, 8, "gofpdf Solutions Ltd.")

	// Company address
	pdf.SetFont("Sarabun", "", 9)
	pdf.SetXY(leftMargin, topMargin+11)
	pdf.Cell(100, 5, "456 Tech Street, Silicon Valley, CA 94025")

	// Phone and Email
	pdf.SetXY(leftMargin, topMargin+16)
	pdf.Cell(100, 5, "Phone: +1 (555) 123-4567 | Email: info@gofpdf.com")

	// Reset text color
	pdf.SetTextColor(textDark[0], textDark[1], textDark[2])

	// ===== INVOICE TITLE AND INFO =====
	currentY := 40.0
	pdf.SetFont("Sarabun", "B", 24)
	pdf.SetXY(pageWidth-rightMargin-60, currentY)
	pdf.CellFormat(60, 10, "INVOICE", "", 0, "R", false, 0, "")

	currentY += 12
	pdf.SetFont("Sarabun", "", 9)
	pdf.SetXY(pageWidth-rightMargin-60, currentY)
	pdf.Cell(30, 5, "DATE:")
	pdf.SetXY(pageWidth-rightMargin-30, currentY)
	pdf.Cell(30, 5, "Jan 15, 2024")

	currentY += 5
	pdf.SetXY(pageWidth-rightMargin-60, currentY)
	pdf.Cell(30, 5, "INVOICE NO:")
	pdf.SetXY(pageWidth-rightMargin-30, currentY)
	pdf.Cell(30, 5, "INV-2024-0042")

	currentY += 5
	pdf.SetFont("Sarabun", "I", 8)
	pdf.SetXY(pageWidth-rightMargin-100, currentY)
	pdf.MultiCell(100, 4, "Payment due within 30 days of invoice date", "", "R", false)

	currentY = 70.0

	// ===== BILL TO / SHIP TO SECTION =====
	// Bill To
	pdf.SetFont("Sarabun", "B", 10)
	pdf.SetXY(leftMargin, currentY)
	pdf.Cell(40, 6, "BILL TO")

	// Ship To
	pdf.SetXY(pageWidth/2, currentY)
	pdf.Cell(40, 6, "SHIP TO")

	currentY += 7
	pdf.SetFont("Sarabun", "", 9)

	// Bill To details
	billToLines := []string{"Sarah Johnson", "", "Digital Innovations Corp.", "789 Business Blvd, Suite 200", "New York, NY 10001"}
	for _, line := range billToLines {
		pdf.SetXY(leftMargin, currentY)
		pdf.Cell(80, 4, line)
		currentY += 4
	}

	// Ship To details
	currentY = 77.0
	shipToLines := []string{"Sarah Johnson - Accounting Dept", "Digital Innovations Corp.", "789 Business Blvd, Suite 200", "New York, NY 10001"}
	for _, line := range shipToLines {
		pdf.SetXY(pageWidth/2, currentY)
		pdf.Cell(80, 4, line)
		currentY += 4
	}

	currentY = 105.0

	// ===== ITEMS TABLE =====
	tableStartY := currentY
	tableStartX := leftMargin
	rowHeight := 6.0

	// Column widths
	colWidths := []float64{80, 25, 35, 35}
	colHeaders := []string{"DESCRIPTION", "QTY", "UNIT PRICE", "TOTAL"}
	colAligns := []string{"L", "C", "R", "R"}

	totalTableWidth := 0.0
	for _, w := range colWidths {
		totalTableWidth += w
	}

	// Header
	pdf.SetFillColor(primaryBlue[0], primaryBlue[1], primaryBlue[2])
	pdf.SetTextColor(white[0], white[1], white[2])
	pdf.SetFont("Sarabun", "B", 10)

	xPos := tableStartX
	for i, header := range colHeaders {
		pdf.SetXY(xPos, tableStartY)
		pdf.CellFormat(colWidths[i], rowHeight, header, "", 0, colAligns[i], true, 0, "")
		xPos += colWidths[i]
	}

	pdf.SetTextColor(textDark[0], textDark[1], textDark[2])
	currentY = tableStartY + rowHeight

	// Data rows with alternating colors
	useLightBg := false
	tableData := [][]string{
		{"Professional PDF Library License", "1", "299.00", "299.00"},
		{"Advanced Table Module", "2", "149.00", "298.00"},
		{"Custom Font Integration Kit", "1", "199.00", "199.00"},
		{"Technical Support (Annual)", "1", "499.00", "499.00"},
		{"API Documentation Bundle", "1", "149.00", "149.00"},
	}

	pdf.SetFont("Sarabun", "", 9)
	for _, row := range tableData {
		// Alternate row background
		if useLightBg {
			pdf.SetFillColor(lightBlue[0], lightBlue[1], lightBlue[2])
			pdf.Rect(tableStartX, currentY-0.5, totalTableWidth, rowHeight+1, "F")
		}
		xPos = tableStartX
		for i, cell := range row {
			pdf.SetXY(xPos, currentY)
			if i == 2 && cell != "" { // Unit price column
				pdf.CellFormat(colWidths[i], rowHeight, "$ "+cell, "", 0, colAligns[i], false, 0, "")
			} else if i == 3 && cell != "" { // Total column
				pdf.CellFormat(colWidths[i], rowHeight, "$ "+cell, "", 0, colAligns[i], false, 0, "")
			} else {
				pdf.CellFormat(colWidths[i], rowHeight, cell, "", 0, colAligns[i], false, 0, "")
			}
			xPos += colWidths[i]
		}
		currentY += rowHeight
		useLightBg = !useLightBg
	}

	// Draw table borders
	pdf.SetLineWidth(0.1)
	pdf.Rect(tableStartX, tableStartY, totalTableWidth, rowHeight*float64(len(tableData)+1), "D")

	// Vertical lines
	xPos = tableStartX
	for i := 0; i < len(colWidths)-1; i++ {
		xPos += colWidths[i]
		pdf.Line(xPos, tableStartY, xPos, tableStartY+rowHeight*float64(len(tableData)+1))
	}

	// Horizontal line after header
	pdf.Line(tableStartX, tableStartY+rowHeight, tableStartX+totalTableWidth, tableStartY+rowHeight)

	// ===== SUMMARY SECTION =====
	currentY = tableStartY + rowHeight*float64(len(tableData)+1) + 2

	// Remarks section
	pdf.SetFont("Sarabun", "", 9)
	pdf.SetXY(leftMargin, currentY)
	pdf.Cell(80, 5, "Remarks / Payment Instructions:")
	pdf.SetXY(leftMargin, currentY+5)
	pdf.Cell(80, 4, "Payment accepted via ACH, Wire Transfer, or Check.")
	pdf.SetXY(leftMargin, currentY+9)
	pdf.Cell(80, 4, "Please include invoice number with payment.")

	// Summary calculations
	summaryX := tableStartX + totalTableWidth - 70
	summaryStartY := currentY
	summaryItems := [][]string{
		{"SUBTOTAL", "$ 1,444.00"},
		{"DISCOUNT (10%)", "$ 144.40"},
		{"TAX RATE", "7.50%"},
		{"TOTAL TAX", "$ 97.47"},
		{"SHIPPING/HANDLING", "$ 0.00"},
	}

	for _, item := range summaryItems {
		pdf.SetXY(summaryX, summaryStartY)
		pdf.Cell(35, 5, item[0])
		pdf.SetX(summaryX + 35)
		pdf.CellFormat(35, 5, item[1], "", 0, "R", false, 0, "")
		summaryStartY += 5
	}
	currentY = summaryStartY

	// Balance Due
	pdf.SetFillColor(primaryBlue[0], primaryBlue[1], primaryBlue[2])
	pdf.SetTextColor(white[0], white[1], white[2])
	pdf.SetFont("Sarabun", "B", 12)
	pdf.SetXY(summaryX, currentY)
	pdf.CellFormat(35, 8, "Balance Due", "", 0, "L", true, 0, "")
	pdf.SetX(summaryX + 35)
	pdf.CellFormat(35, 8, "$ 1,397.07", "", 0, "R", true, 0, "")

	// Reset colors
	pdf.SetTextColor(textDark[0], textDark[1], textDark[2])

	// QR Code - Invoice details (placed at bottom left)
	qrData := "Invoice: INV-2024-0042\nDate: 2024-01-15\nAmount: $1,397.07"
	qrKey := barcode.RegisterQR(pdf, qrData, qr.M, qr.Auto)
	barcode.Barcode(pdf, qrKey, leftMargin, 250, 20, 20, false)

	// Footer line
	pdf.SetFillColor(primaryBlue[0], primaryBlue[1], primaryBlue[2])
	pdf.Rect(0, 280, pageWidth, 5, "F")
}
