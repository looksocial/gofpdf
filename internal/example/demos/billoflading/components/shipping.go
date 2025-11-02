package components

import (
	"github.com/looksocial/gofpdf"
	"github.com/looksocial/gofpdf/contrib/barcode"
)

// RenderShipFrom renders the Ship From section with barcode
func RenderShipFrom(pdf *gofpdf.Fpdf, leftMargin, currentY, contentWidth float64, black, white []int) float64 {
	sectionHeight := 28.0

	// Header bar for "SHIP FROM"
	pdf.SetFillColor(black[0], black[1], black[2])
	pdf.Rect(leftMargin, currentY, contentWidth/2, 6, "F")

	// "SHIP FROM" text
	pdf.SetTextColor(white[0], white[1], white[2])
	pdf.SetFont("Arial", "B", 10)
	pdf.SetXY(leftMargin+2, currentY+1)
	pdf.Cell(contentWidth/2-4, 5, "SHIP FROM")

	// Bill of Lading Number section header
	pdf.Rect(leftMargin+contentWidth/2, currentY, contentWidth/2, 6, "F")
	pdf.SetXY(leftMargin+contentWidth/2+2, currentY+1)
	pdf.CellFormat(contentWidth/2-4, 5, "Bill of Lading Number: GOFPDF2024001", "", 0, "L", false, 0, "")

	pdf.SetTextColor(black[0], black[1], black[2])
	currentY += 6

	// Ship from details - left side
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+3, currentY+2)
	pdf.Cell(contentWidth/2-6, 5, "Name: gofpdf Shipping Co.")
	pdf.SetXY(leftMargin+3, currentY+7)
	pdf.Cell(contentWidth/2-6, 5, "Address: 1234 Demo Street")
	pdf.SetXY(leftMargin+3, currentY+12)
	pdf.Cell(contentWidth/2-6, 5, "City/State/Zip: Sample City, ST 12345")
	pdf.SetXY(leftMargin+3, currentY+17)
	pdf.Cell(contentWidth/2-6, 5, "SID#: GOFPDF-001")
	pdf.SetXY(leftMargin+3, currentY+22)
	pdf.Cell(20, 5, "FOB:")
	pdf.Rect(leftMargin+24, currentY+22, 4, 4, "D")

	// Barcode - right side
	barcodeKey := barcode.RegisterCode128(pdf, "GOFPDF2024001")
	barcode.Barcode(pdf, barcodeKey, leftMargin+contentWidth/2+15, currentY+2, 50, 12, false)

	// Barcode number below barcode
	pdf.SetFont("Arial", "", 7)
	pdf.SetXY(leftMargin+contentWidth/2+2, currentY+17)
	pdf.CellFormat(contentWidth/2-4, 4, "GOFPDF2024001", "", 0, "C", false, 0, "")

	// Draw border
	pdf.Rect(leftMargin, currentY-6, contentWidth, sectionHeight, "D")
	pdf.Line(leftMargin+contentWidth/2, currentY-6, leftMargin+contentWidth/2, currentY+sectionHeight-6)

	return currentY + sectionHeight - 6
}

// RenderShipTo renders the Ship To and Carrier section
func RenderShipTo(pdf *gofpdf.Fpdf, leftMargin, currentY, contentWidth float64, black, white []int) float64 {
	sectionHeight := 28.0

	// Header bar
	pdf.SetFillColor(black[0], black[1], black[2])
	pdf.Rect(leftMargin, currentY, contentWidth/2, 6, "F")
	pdf.SetTextColor(white[0], white[1], white[2])
	pdf.SetFont("Arial", "B", 10)
	pdf.SetXY(leftMargin+2, currentY+1)
	pdf.Cell(contentWidth/2-4, 5, "SHIP TO")

	// Carrier section header
	pdf.Rect(leftMargin+contentWidth/2, currentY, contentWidth/2, 6, "F")
	pdf.SetXY(leftMargin+contentWidth/2+2, currentY+1)
	pdf.Cell(contentWidth/2-4, 5, "CARRIER")

	pdf.SetTextColor(black[0], black[1], black[2])
	currentY += 6

	// Ship to details - left side
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+3, currentY+2)
	pdf.Cell(65, 5, "Name: gofpdf Receiving LLC")
	pdf.SetXY(leftMargin+70, currentY+2)
	pdf.Cell(25, 5, "Location #:")

	pdf.SetXY(leftMargin+3, currentY+7)
	pdf.Cell(contentWidth/2-6, 5, "Address: 5678 Delivery Avenue")

	pdf.SetXY(leftMargin+3, currentY+12)
	pdf.Cell(contentWidth/2-6, 5, "City/State/Zip: Example Town, EX 54321")

	pdf.SetXY(leftMargin+3, currentY+17)
	pdf.Cell(40, 5, "CID#: GOFPDF-REC-001")
	pdf.SetXY(leftMargin+45, currentY+17)
	pdf.Cell(20, 5, "FOB:")
	pdf.Rect(leftMargin+66, currentY+17, 4, 4, "D")

	// Carrier details - right side
	pdf.SetXY(leftMargin+contentWidth/2+3, currentY+2)
	pdf.Cell(contentWidth/2-6, 5, "GOFPDF FREIGHT SERVICES")
	pdf.SetXY(leftMargin+contentWidth/2+3, currentY+7)
	pdf.Cell(60, 5, "Trailer number:")
	pdf.SetXY(leftMargin+contentWidth/2+3, currentY+12)
	pdf.Cell(60, 5, "Seal number(s):")
	pdf.SetXY(leftMargin+contentWidth/2+3, currentY+17)
	pdf.Cell(60, 5, "SCAC: GOFP")
	pdf.SetXY(leftMargin+contentWidth/2+3, currentY+22)
	pdf.Cell(60, 5, "Pro number:")

	pdf.Rect(leftMargin, currentY-6, contentWidth, sectionHeight, "D")
	pdf.Line(leftMargin+contentWidth/2, currentY-6, leftMargin+contentWidth/2, currentY+sectionHeight-6)

	return currentY + sectionHeight - 6
}

// RenderThirdParty renders the Third Party Freight Charges Bill To section
func RenderThirdParty(pdf *gofpdf.Fpdf, leftMargin, currentY, contentWidth float64, black, white []int) float64 {
	sectionHeight := 18.0

	pdf.SetFillColor(black[0], black[1], black[2])
	pdf.Rect(leftMargin, currentY, contentWidth, 6, "F")
	pdf.SetTextColor(white[0], white[1], white[2])
	pdf.SetFont("Arial", "B", 10)
	pdf.SetXY(leftMargin+2, currentY+1)
	pdf.Cell(contentWidth-4, 5, "THIRD PARTY FREIGHT CHARGES BILL TO")

	pdf.SetTextColor(black[0], black[1], black[2])
	currentY += 6

	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+2, currentY+1)
	pdf.Cell(contentWidth/2-4, 4, "Name: gofpdf Billing Inc.")
	pdf.SetXY(leftMargin+2, currentY+5)
	pdf.Cell(contentWidth/2-4, 4, "Address: 9999 Billing Blvd.")
	pdf.SetXY(leftMargin+2, currentY+9)
	pdf.Cell(contentWidth/2-4, 4, "City/State/Zip: Sample City, ST 99999")

	// Right side - Freight Charge Terms
	pdf.SetXY(leftMargin+contentWidth/2+2, currentY+1)
	pdf.Cell(contentWidth/2-4, 4, "Freight Charge Terms: (freight charges are prepaid")
	pdf.SetXY(leftMargin+contentWidth/2+2, currentY+5)
	pdf.Cell(contentWidth/2-4, 4, "unless marked otherwise)")

	pdf.Rect(leftMargin, currentY-6, contentWidth, sectionHeight, "D")
	pdf.Line(leftMargin+contentWidth/2, currentY-6, leftMargin+contentWidth/2, currentY+sectionHeight-6)

	return currentY + sectionHeight - 6
}

// RenderSpecialInstructions renders the Special Instructions section
func RenderSpecialInstructions(pdf *gofpdf.Fpdf, leftMargin, currentY, contentWidth float64) float64 {
	sectionHeight := 25.0

	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+2, currentY+1)
	pdf.Cell(contentWidth/2-4, 4, "SPECIAL INSTRUCTIONS:")

	// Right side - Prepaid/Collect checkboxes
	pdf.SetXY(leftMargin+contentWidth/2+2, currentY+1)
	pdf.Cell(20, 4, "Prepaid: ____")
	pdf.SetXY(leftMargin+contentWidth/2+25, currentY+1)
	pdf.Cell(20, 4, "Collect: ____")
	pdf.SetXY(leftMargin+contentWidth/2+48, currentY+1)
	pdf.Cell(20, 4, "3rd Party:")

	pdf.SetXY(leftMargin+contentWidth/2+2, currentY+7)
	pdf.Rect(leftMargin+contentWidth/2+2, currentY+7, 4, 4, "D")
	pdf.SetXY(leftMargin+contentWidth/2+6.5, currentY+7)
	pdf.Cell(4, 4, "X")
	pdf.SetXY(leftMargin+contentWidth/2+12, currentY+7)
	pdf.Cell(contentWidth/2-20, 4, "Master Bill of Lading: with attached")
	pdf.SetXY(leftMargin+contentWidth/2+2, currentY+11)
	pdf.Cell(contentWidth/2-4, 4, "(check box)  underlying Bills of Lading")

	pdf.Rect(leftMargin, currentY, contentWidth, sectionHeight, "D")
	pdf.Line(leftMargin+contentWidth/2, currentY, leftMargin+contentWidth/2, currentY+sectionHeight)

	return currentY + sectionHeight
}
