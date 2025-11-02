//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/boombuler/barcode/qr"
	"github.com/looksocial/gofpdf"
	"github.com/looksocial/gofpdf/contrib/barcode"
)

func main() {
	// Create new PDF
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Use embedded fonts (Arial works with or without this, but kept for compatibility)
	pdf.UseEmbeddedFonts()

	// Add a page
	pdf.AddPage()

	// Generate the invoice
	generateInvoicePattern2(pdf)

	// Ensure pdf directory exists
	pdfDir := filepath.Join("..", "..", "..", "..", "pdf")
	err := os.MkdirAll(pdfDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Save to file in pdf directory
	outputPath := filepath.Join(pdfDir, "invoice_pattern2.pdf")
	err = pdf.OutputFileAndClose(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	// Get absolute path for display
	absPath, _ := filepath.Abs(outputPath)
	fmt.Printf("✓ Invoice Pattern 2 created successfully: %s\n", absPath)
}

func generateInvoicePattern2(pdf *gofpdf.Fpdf) {
	// Modern Colors
	primaryTeal := []int{0, 150, 136} // Modern teal (#009688)
	lightTeal := []int{224, 242, 241} // Light teal accent
	white := []int{255, 255, 255}     // White
	textDark := []int{33, 33, 33}     // Dark text

	// Page width and margins
	pageWidth, _ := pdf.GetPageSize()
	leftMargin, topMargin, rightMargin, _ := pdf.GetMargins()

	currentY := topMargin

	// ===== HEADER SECTION =====
	// Logo placeholder (circle)
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
	pdf.Cell(100, 6, "gofpdf Solutions Ltd.")

	pdf.SetFont("Arial", "", 9)
	pdf.SetXY(leftMargin, currentY+38)
	pdf.Cell(100, 4, "456 Tech Street, Silicon Valley, CA 94025")

	pdf.SetXY(leftMargin, currentY+42)
	pdf.Cell(100, 4, "www.gofpdf.com | info@gofpdf.com")

	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(leftMargin, currentY+46)
	pdf.Cell(100, 4, "+1 (555) 123-4567")

	// Right side - Page info
	pdf.SetFont("Arial", "", 9)
	infoX := pageWidth - rightMargin - 60
	pdf.SetXY(infoX, currentY)
	pdf.Cell(30, 5, "Page")
	pdf.SetXY(infoX+35, currentY)
	pdf.Cell(25, 5, "1 of 1")

	pdf.SetXY(infoX, currentY+5)
	pdf.Cell(30, 5, "Date")
	pdf.SetXY(infoX+35, currentY+5)
	pdf.Cell(25, 5, "01/15/2024")

	pdf.SetXY(infoX, currentY+10)
	pdf.Cell(30, 5, "Date of Expiry")
	pdf.SetXY(infoX+35, currentY+10)
	pdf.Cell(25, 5, "02/14/2024")

	pdf.SetXY(infoX, currentY+15)
	pdf.Cell(30, 5, "Estimate No.")
	pdf.SetXY(infoX+35, currentY+15)
	pdf.Cell(25, 5, "EST-2024-023")

	pdf.SetXY(infoX, currentY+20)
	pdf.Cell(30, 5, "Customer ID")
	pdf.SetXY(infoX+35, currentY+20)
	pdf.Cell(25, 5, "DI-2024-015")

	currentY += 55

	// ===== BILL TO / SHIP TO SECTION =====
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(leftMargin, currentY)
	pdf.Cell(30, 5, "BILL TO")

	pdf.SetXY(pageWidth/2, currentY)
	pdf.Cell(30, 5, "SHIP TO")

	currentY += 5
	pdf.SetFont("Arial", "", 8)

	// Bill To details
	billToLines := []string{"Michael Chen", "TechStart Industries", "1520 Innovation Drive", "Austin, TX 78701", "mchen@techstart.io"}
	billToStartY := currentY
	for _, line := range billToLines {
		pdf.SetXY(leftMargin, billToStartY)
		pdf.Cell(80, 4, line)
		billToStartY += 4
	}

	// Ship To details
	shipToStartY := currentY
	shipToLines := []string{"Michael Chen", "TechStart Industries", "1520 Innovation Drive", "Austin, TX 78701", "+1 (512) 555-9876"}
	for _, line := range shipToLines {
		pdf.SetXY(pageWidth/2, shipToStartY)
		pdf.Cell(80, 4, line)
		shipToStartY += 4
	}

	currentY = topMargin + 80

	// ===== SHIPMENT INFORMATION =====
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(leftMargin, currentY)
	pdf.Cell(100, 5, "SHIPMENT INFORMATION")

	currentY += 5
	pdf.SetFont("Arial", "", 8)

	shipmentItems := [][]string{
		{"P.O. #", "PO-2024-TS-0789", "Mode of Transportation", "Standard Ground"},
		{"P.O. Date", "12/20/2023", "Transportation Terms", "FOB Origin"},
		{"Letter of Credit #", "LC-INTL-0422", "Number of Packages", "3 crates"},
		{"Currency", "USD", "Est. Gross Weight", "125 kg"},
		{"Payment Terms", "Net 30", "Est. Net Weight", "95 kg"},
		{"Est. Ship Date", "01/22/2024", "Carrier", "FedEx Ground"},
	}

	for i := 0; i < len(shipmentItems); i++ {
		item := shipmentItems[i]
		pdf.SetXY(leftMargin, currentY)
		pdf.Cell(40, 4, item[0])
		pdf.SetX(leftMargin + 40)
		pdf.Cell(50, 4, item[1])
		pdf.SetXY(pageWidth/2, currentY)
		pdf.Cell(40, 4, item[2])
		pdf.SetX(pageWidth/2 + 40)
		pdf.Cell(40, 4, item[3])
		currentY += 4
	}

	currentY += 3

	// ===== ITEMS TABLE =====
	tableStartY := currentY
	tableStartX := leftMargin
	rowHeight := 5.0

	// Column widths (optimized to fit A4 without overflow)
	colWidths := []float64{15, 50, 15, 12, 28, 22, 28} // Total: 170mm
	colHeaders := []string{"ITEM", "DESCRIPTION", "UNIT", "QTY", "UNIT PRICE", "TAX", "TOTAL"}
	colAligns := []string{"L", "L", "C", "C", "R", "R", "R"}

	totalTableWidth := 0.0
	for _, w := range colWidths {
		totalTableWidth += w
	}

	// Header
	pdf.SetFillColor(primaryTeal[0], primaryTeal[1], primaryTeal[2])
	pdf.SetTextColor(white[0], white[1], white[2])
	pdf.SetFont("Arial", "B", 8)

	xPos := tableStartX
	for i, header := range colHeaders {
		pdf.SetXY(xPos, tableStartY)
		pdf.CellFormat(colWidths[i], rowHeight, header, "1", 0, colAligns[i], true, 0, "")
		xPos += colWidths[i]
	}

	pdf.SetTextColor(textDark[0], textDark[1], textDark[2])
	currentY = tableStartY + rowHeight

	// Data rows with actual content (reduced to fit one page)
	tableData := [][]string{
		{"1", "PDF Library Enterprise", "Lic.", "5", "2,499", "250", "12,745"},
		{"2", "Table Module w/Nesting", "Lic.", "5", "449", "90", "2,295"},
		{"3", "Font Integration Kit", "Lic.", "3", "599", "60", "1,857"},
		{"4", "Priority Support (1yr)", "Year", "1", "999", "200", "1,199"},
		{"5", "API Documentation Pkg", "Set", "1", "249", "50", "299"},
	}

	// Alternating row backgrounds
	useLightBg := false
	pdf.SetFont("Arial", "", 7) // Slightly smaller font
	for _, row := range tableData {
		// Alternate row background
		if useLightBg {
			pdf.SetFillColor(lightTeal[0], lightTeal[1], lightTeal[2])
			pdf.Rect(tableStartX, currentY-0.3, totalTableWidth, rowHeight+0.6, "F")
		}
		xPos = tableStartX
		for j, cell := range row {
			pdf.SetXY(xPos, currentY)
			if j >= 4 && cell != "" {
				// Format numbers with $ sign
				pdf.CellFormat(colWidths[j], rowHeight, "$"+cell, "1", 0, colAligns[j], false, 0, "")
			} else {
				pdf.CellFormat(colWidths[j], rowHeight, cell, "1", 0, colAligns[j], false, 0, "")
			}
			xPos += colWidths[j]
		}
		currentY += rowHeight
		useLightBg = !useLightBg
	}

	// Empty rows for remaining space (reduced count)
	for i := 0; i < 3; i++ {
		if useLightBg {
			pdf.SetFillColor(lightTeal[0], lightTeal[1], lightTeal[2])
			pdf.Rect(tableStartX, currentY-0.3, totalTableWidth, rowHeight+0.6, "F")
		}
		xPos = tableStartX
		for j := 0; j < len(colWidths); j++ {
			pdf.SetXY(xPos, currentY)
			pdf.CellFormat(colWidths[j], rowHeight, "", "1", 0, colAligns[j], false, 0, "")
			xPos += colWidths[j]
		}
		currentY += rowHeight
		useLightBg = !useLightBg
	}

	currentY += 2

	// ===== SUMMARY SECTION =====
	summaryX := pageWidth - rightMargin - 80

	summaryItems := [][]string{
		{"SUBTOTAL", "18,395.00"},
		{"SUBTOTAL LESS DISCOUNT", "17,555.00"},
		{"SUBJECT TO SALES TAX", "17,555.00"},
		{"TAX RATE", "8.25%"},
		{"TOTAL TAX", "1,448.29"},
		{"SHIPPING/HANDLING", "0.00"},
		{"INSURANCE", "250.00"},
		{"INSTALLATION FEE", "500.00"},
	}

	pdf.SetFont("Arial", "", 7)
	for _, item := range summaryItems {
		pdf.SetXY(summaryX, currentY)
		pdf.Cell(50, 4, item[0])
		pdf.SetX(summaryX + 50)
		if item[1] != "" {
			pdf.CellFormat(30, 4, "$"+item[1], "", 0, "R", false, 0, "")
		}
		currentY += 4
	}

	// Quote Total
	pdf.SetLineWidth(0.3)
	pdf.Line(summaryX, currentY, summaryX+80, currentY)
	currentY += 1
	pdf.Line(summaryX, currentY, summaryX+80, currentY)
	currentY += 3

	pdf.SetFont("Arial", "B", 10)
	pdf.SetFillColor(primaryTeal[0], primaryTeal[1], primaryTeal[2])
	pdf.SetTextColor(white[0], white[1], white[2])
	quoteTotalY := currentY - 1
	pdf.SetXY(summaryX, quoteTotalY)
	pdf.CellFormat(50, 7, "QUOTE TOTAL", "1", 0, "L", true, 0, "")
	pdf.SetXY(summaryX+50, quoteTotalY)
	pdf.CellFormat(30, 7, "$19,753", "1", 0, "R", true, 0, "")
	pdf.SetTextColor(textDark[0], textDark[1], textDark[2])

	// Special notes section
	currentY = tableStartY + rowHeight*9 + 4
	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(leftMargin, currentY)
	pdf.Cell(80, 4, "TERMS & CONDITIONS")

	// QR Code - Quote details (placed under TERMS & CONDITIONS, aligned with QUOTE TOTAL)
	qrData := "Quote: EST-2024-023\nDate: 2024-01-15\nTotal: $19,753"
	qrKey := barcode.RegisterQR(pdf, qrData, qr.M, qr.Auto)
	// Position QR code to align with QUOTE TOTAL row
	barcode.Barcode(pdf, qrKey, leftMargin, quoteTotalY, 20, 20, false)

	// Declaration and signature
	currentY = 260
	pdf.SetFont("Arial", "I", 8)
	pdf.SetXY(leftMargin, currentY)
	pdf.Cell(100, 4, "I declare that the above information is true and correct to the best of my knowledge.")

	currentY += 6
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin, currentY)
	pdf.Cell(30, 4, "Signature")
	pdf.SetXY(leftMargin+60, currentY)
	pdf.Cell(30, 4, "Date")

	// Signature lines
	pdf.Line(leftMargin, currentY+5, leftMargin+50, currentY+5)
	pdf.Line(leftMargin+60, currentY+5, leftMargin+110, currentY+5)
}
