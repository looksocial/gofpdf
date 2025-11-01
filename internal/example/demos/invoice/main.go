package main

import (
	"github.com/looksocial/gofpdf"
	"github.com/looksocial/gofpdf/contrib/barcode"
	"github.com/looksocial/gofpdf/internal/example"
	"github.com/looksocial/gofpdf/table"
)

func main() {
	// Create new PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Set font
	pdf.SetFont("Arial", "", 12)

	// Colors - Blue theme
	blueColor := []int{33, 150, 243} // Blue for accents (Material Design Blue 500)
	darkGray := []int{51, 51, 51}    // Dark gray for text

	// Header Section - Compact
	pdf.SetFont("Arial", "B", 20)
	pdf.SetTextColor(darkGray[0], darkGray[1], darkGray[2])
	pdf.SetXY(10, 10)
	pdf.Cell(80, 8, "Sales Invoice")

	// Logo (simplified - just text representation)
	pdf.SetFont("Arial", "B", 14)
	pdf.SetFillColor(blueColor[0], blueColor[1], blueColor[2])
	pdf.SetXY(170, 10)
	pdf.Rect(170, 10, 5, 12, "F") // Blue stripe
	pdf.SetXY(177, 12)
	pdf.SetTextColor(0, 0, 0)
	pdf.Cell(20, 8, "gofpdf")

	// Invoice Details Section - Compact
	pdf.SetXY(10, 25)
	pdf.SetFillColor(blueColor[0], blueColor[1], blueColor[2])
	pdf.Rect(10, 25, 2, 38, "F") // Blue vertical line

	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(blueColor[0], blueColor[1], blueColor[2])
	pdf.SetXY(15, 27)
	pdf.Cell(40, 4, "INVOICE NUMBER")
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(darkGray[0], darkGray[1], darkGray[2])
	pdf.SetXY(15, 31)
	pdf.Cell(40, 4, "#9000000001")

	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(blueColor[0], blueColor[1], blueColor[2])
	pdf.SetXY(15, 36)
	pdf.Cell(40, 4, "ORDER")
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(darkGray[0], darkGray[1], darkGray[2])
	pdf.SetXY(15, 40)
	pdf.Cell(40, 4, "#9000000001")

	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(blueColor[0], blueColor[1], blueColor[2])
	pdf.SetXY(15, 45)
	pdf.Cell(40, 4, "ORDER DATE")
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(darkGray[0], darkGray[1], darkGray[2])
	pdf.SetXY(15, 49)
	pdf.Cell(70, 4, "Dec 11, 2020, 10:56:14 AM")

	// Barcode using contrib/barcode
	invoiceNumber := "9000000001"
	// Register Code128 barcode
	barcodeKey := barcode.RegisterCode128(pdf, invoiceNumber)
	// Display barcode (smaller size to fit)
	barcode.Barcode(pdf, barcodeKey, 170, 25, 25, 12, false)

	// Customer Information - Compact
	pdf.SetXY(10, 70)
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(blueColor[0], blueColor[1], blueColor[2])
	pdf.Cell(40, 4, "CUSTOMER NAME")
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(darkGray[0], darkGray[1], darkGray[2])
	pdf.SetXY(10, 74)
	pdf.Cell(80, 4, "Veronica Costello")

	pdf.SetXY(10, 80)
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(blueColor[0], blueColor[1], blueColor[2])
	pdf.Cell(40, 4, "ADDRESS")
	pdf.SetFont("Arial", "", 8)
	pdf.SetTextColor(darkGray[0], darkGray[1], darkGray[2])
	pdf.SetXY(10, 84)
	pdf.Cell(75, 3, "Veronica Costello")
	pdf.SetXY(10, 87)
	pdf.Cell(75, 3, "6146 Honey Bluff Parkway")
	pdf.SetXY(10, 90)
	pdf.Cell(75, 3, "Calder, Michigan, 49628-7978")
	pdf.SetXY(10, 93)
	pdf.Cell(75, 3, "United States")
	pdf.SetXY(10, 96)
	pdf.Cell(75, 3, "T: (555) 229-3326")

	pdf.SetXY(10, 102)
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(blueColor[0], blueColor[1], blueColor[2])
	pdf.Cell(50, 4, "SHIPPING METHOD")
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(darkGray[0], darkGray[1], darkGray[2])
	pdf.SetXY(10, 106)
	pdf.Cell(80, 4, "Flat Rate - Fixed")

	pdf.SetXY(10, 112)
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(blueColor[0], blueColor[1], blueColor[2])
	pdf.Cell(50, 4, "PAYMENT METHOD")
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(darkGray[0], darkGray[1], darkGray[2])
	pdf.SetXY(10, 116)
	pdf.Cell(80, 4, "Check / Money order")

	// Items Table - Compact positioning
	// A4 page width: 210mm, so max usable width is ~190mm (with 10mm margins)
	tableStartX := 10.0
	tableStartY := 125.0  // Moved up significantly
	itemColWidth := 100.0 // Adjusted to fit A4 better
	qtyColWidth := 30.0
	subtotalColWidth := 50.0
	rowHeight := 10.0 // Reduced row height

	// Define table columns for header only
	columns := []table.Column{
		{Key: "item", Label: "ITEMS", Width: itemColWidth, Align: "L", HeaderAlign: "L"},
		{Key: "qty", Label: "QTY", Width: qtyColWidth, Align: "C", HeaderAlign: "C"},
		{Key: "subtotal", Label: "SUBTOTAL", Width: subtotalColWidth, Align: "R", HeaderAlign: "R"},
	}

	// Create table for header
	tbl := table.NewTable(pdf, columns).
		WithStartPosition(tableStartX, tableStartY).
		WithHeaderStyle(table.CellStyle{
			Border:    "0",
			Bold:      true,
			FillColor: blueColor,
			TextColor: []int{255, 255, 255},
			FontSize:  11,
		}).
		WithRowHeight(rowHeight)

	// Render header
	tbl.AddHeader()

	// Item data with separate name and SKU
	type InvoiceItem struct {
		Name     string
		SKU      string
		Quantity string
		Subtotal string
	}

	items := []InvoiceItem{
		{"Endurance Watch", "24-MG01", "1", "$49.00"},
		{"Fusion Backpack", "24-MB02", "1", "$50.00"},
		{"Affirm Water Bottle", "24-UG06", "1", "$7.00"},
		{"Pursuit Lumaflex™ Tone Band", "24-UG02", "1", "$16.00"},
		{"Go-Get'r Pushup Grips", "24-UG05", "1", "$19.00"},
	}

	// Render item rows manually with MultiCell for item name/SKU
	currentY := tableStartY + rowHeight
	pdf.SetFont("Arial", "", 10)
	pdf.SetTextColor(darkGray[0], darkGray[1], darkGray[2])

	for _, item := range items {
		xPos := tableStartX
		pdf.SetXY(xPos, currentY)

		// Item column - use MultiCell for name and SKU
		pdf.SetFillColor(255, 255, 255)
		pdf.Rect(xPos, currentY, itemColWidth, rowHeight, "D") // Draw border

		// Item name
		pdf.SetXY(xPos+1, currentY+1.5)
		pdf.SetFont("Arial", "", 9)
		pdf.Cell(itemColWidth-2, 4, item.Name)

		// SKU in smaller font
		pdf.SetFont("Arial", "", 7)
		pdf.SetXY(xPos+1, currentY+6)
		pdf.SetTextColor(100, 100, 100)
		pdf.Cell(itemColWidth-2, 3, item.SKU)
		pdf.SetTextColor(darkGray[0], darkGray[1], darkGray[2])
		pdf.SetFont("Arial", "", 9)

		// Quantity column
		xPos += itemColWidth
		pdf.SetXY(xPos, currentY)
		pdf.Rect(xPos, currentY, qtyColWidth, rowHeight, "D")
		pdf.CellFormat(qtyColWidth, rowHeight, item.Quantity, "0", 0, "C", false, 0, "")

		// Subtotal column
		xPos += qtyColWidth
		pdf.SetXY(xPos, currentY)
		pdf.Rect(xPos, currentY, subtotalColWidth, rowHeight, "D")
		pdf.CellFormat(subtotalColWidth, rowHeight, item.Subtotal, "0", 0, "R", false, 0, "")

		currentY += rowHeight
	}

	// Summary Section (right-aligned) - Compact
	// Position summary to fit within A4 width (210mm)
	pageWidth, _ := pdf.GetPageSize()
	leftMargin, _, rightMargin, _ := pdf.GetMargins()
	usableWidth := pageWidth - leftMargin - rightMargin
	summaryWidth := 60.0                               // Width for summary values
	labelWidth := 40.0                                 // Width for labels
	summaryX := pageWidth - rightMargin - summaryWidth // Right-aligned
	labelX := summaryX - labelWidth                    // Labels to the left of values
	summaryY := currentY + 5                           // Reduced spacing
	lineHeight := 5.0                                  // Reduced line height

	// Ensure summary doesn't exceed A4 page (297mm height)
	if summaryY > 230 {
		summaryY = 230 // Keep margin from bottom
	}

	// Ensure labels don't overflow to the left
	if labelX < leftMargin {
		// Adjust if labels would overflow
		labelX = leftMargin
		summaryX = labelX + labelWidth
		// Recalculate if summary would overflow
		if summaryX+summaryWidth > pageWidth-rightMargin {
			summaryWidth = pageWidth - rightMargin - summaryX
		}
	}

	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(darkGray[0], darkGray[1], darkGray[2])

	// Subtotal
	pdf.SetXY(labelX, summaryY)
	pdf.Cell(labelWidth, lineHeight, "SUBTOTAL")
	pdf.SetXY(summaryX, summaryY)
	pdf.CellFormat(summaryWidth, lineHeight, "$141.00", "0", 0, "R", false, 0, "")

	// Discount
	summaryY += lineHeight + 1
	pdf.SetXY(labelX, summaryY)
	pdf.Cell(labelWidth, lineHeight, "DISCOUNT")
	summaryY += lineHeight + 1
	pdf.SetXY(labelX, summaryY)
	pdf.SetFont("Arial", "", 7)
	discountText1 := "(EYHPAOMMT9O9FXDH, FREE SHIPPING"
	pdf.Cell(labelWidth, lineHeight-1, discountText1)
	summaryY += lineHeight
	pdf.SetXY(labelX, summaryY)
	discountText2 := "ON ANY PURCHASE OVER $50)"
	pdf.Cell(labelWidth, lineHeight-1, discountText2)
	summaryY += lineHeight - 1
	pdf.SetXY(summaryX, summaryY)
	pdf.SetFont("Arial", "", 9)
	pdf.CellFormat(summaryWidth, lineHeight, "-$14.10", "0", 0, "R", false, 0, "")

	// Tax
	summaryY += lineHeight + 1
	pdf.SetXY(labelX, summaryY)
	pdf.Cell(labelWidth, lineHeight, "TAX")
	pdf.SetXY(summaryX, summaryY)
	pdf.CellFormat(summaryWidth, lineHeight, "$10.47", "0", 0, "R", false, 0, "")

	// Shipping & Handling
	summaryY += lineHeight + 1
	pdf.SetXY(labelX, summaryY)
	pdf.Cell(labelWidth, lineHeight, "SHIPPING & HANDLING")
	pdf.SetXY(summaryX, summaryY)
	pdf.SetDrawColor(blueColor[0], blueColor[1], blueColor[2])
	pdf.SetLineWidth(0.5)
	pdf.CellFormat(summaryWidth, lineHeight, "$25.00", "B", 0, "R", false, 0, "")

	// Grand Total
	summaryY += lineHeight + 2
	pdf.SetXY(labelX, summaryY)
	pdf.SetFont("Arial", "B", 9)
	pdf.Cell(labelWidth, lineHeight, "GRAND TOTAL")
	pdf.SetXY(summaryX, summaryY)
	pdf.SetDrawColor(blueColor[0], blueColor[1], blueColor[2])
	pdf.CellFormat(summaryWidth, lineHeight, "$162.37", "B", 0, "R", false, 0, "")

	// Footer Section - Compact
	// Position footer within A4 page bounds (297mm height)
	footerY := summaryY + 15 // Position after summary
	if footerY > 270 {
		footerY = 270 // Keep margin from bottom of A4 page
	}
	pdf.SetXY(leftMargin, footerY)
	pdf.SetFont("Arial", "B", 11)
	pdf.SetTextColor(darkGray[0], darkGray[1], darkGray[2])
	pdf.CellFormat(usableWidth, 6, "Thank you for your order!", "0", 0, "C", false, 0, "")

	footerY += 6
	pdf.SetXY(leftMargin, footerY)
	pdf.SetFont("Arial", "", 8)
	pdf.SetTextColor(darkGray[0], darkGray[1], darkGray[2])
	// Remove email from prompt text since it's shown in company info below
	pdf.CellFormat(usableWidth, 4, "If you have questions about your order, you can contact us.", "0", 0, "C", false, 0, "")

	// Company Address - Compact
	footerY += 10
	pdf.SetXY(leftMargin, footerY)
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(darkGray[0], darkGray[1], darkGray[2])
	pdf.CellFormat(usableWidth, 4, "gofpdf", "0", 0, "C", false, 0, "")
	footerY += 4
	pdf.SetXY(leftMargin, footerY)
	pdf.SetFont("Arial", "", 8)

	// Check each line to ensure it fits
	companyLines := []string{
		"Go PDF Library for Go Programming",
		"https://github.com/looksocial/gofpdf",
	}

	for _, line := range companyLines {
		lineWidth := pdf.GetStringWidth(line)
		if lineWidth > usableWidth {
			// Text too long, use MultiCell or wrap
			pdf.SetXY(leftMargin, footerY)
			pdf.MultiCell(usableWidth, 3, line, "0", "C", false)
			footerY = pdf.GetY()
		} else {
			pdf.SetXY(leftMargin, footerY)
			pdf.CellFormat(usableWidth, 3, line, "0", 0, "C", false, 0, "")
			footerY += 3
		}
	}

	// Save PDF
	outputPath := example.PdfFile("invoice_demo.pdf")
	err := pdf.OutputFileAndClose(outputPath)
	if err != nil {
		panic(err)
	}
}
