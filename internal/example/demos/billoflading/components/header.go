package components

import (
	"github.com/looksocial/gofpdf"
)

// RenderHeader renders the date and title section of the Bill of Lading
func RenderHeader(pdf *gofpdf.Fpdf, leftMargin, topMargin, contentWidth float64) float64 {
	currentY := topMargin

	// Date
	pdf.SetFont("Arial", "", 9)
	pdf.SetXY(leftMargin, currentY)
	pdf.Cell(50, 5, "Date: 2/25/2016")

	// Title - centered
	pdf.SetFont("Arial", "B", 18)
	pdf.SetXY(leftMargin, currentY)
	pdf.CellFormat(contentWidth, 8, "BILL OF LADING", "", 0, "C", false, 0, "")

	currentY += 12

	// Set line width for borders
	pdf.SetLineWidth(0.5)

	return currentY
}
