package components

import (
	"github.com/looksocial/gofpdf"
)

// RenderFooter renders the footer section with COD Amount, liability notes, and signatures
func RenderFooter(pdf *gofpdf.Fpdf, leftMargin, currentY float64) {
	// ===== FOOTER SECTION =====
	footerStartY := currentY
	
	// Left side - Value declaration
	pdf.SetFont("Arial", "", 6)
	pdf.SetXY(leftMargin, footerStartY)
	pdf.MultiCell(95, 3, "Where the rate is dependent on value, shippers are required to state specifically in writing the agreed or declared value of the property as follows: \"The agreed or declared value of the property is specifically stated by the shipper to be not exceeding _________________ per _______________\"", "1", "L", false)

	// Right side - COD Amount section
	pdf.SetXY(leftMargin+95, footerStartY)
	pdf.CellFormat(47.5, 6, "COD Amount:", "1", 0, "L", false, 0, "")
	pdf.SetXY(leftMargin+142.5, footerStartY)
	pdf.CellFormat(47.5, 6, "5.00", "1", 0, "C", false, 0, "")

	pdf.SetXY(leftMargin+95, footerStartY+6)
	pdf.CellFormat(30, 6, "Fee Terms:", "1", 0, "L", false, 0, "")
	pdf.SetXY(leftMargin+125, footerStartY+6)
	pdf.CellFormat(17.5, 6, "Collect:", "1", 0, "L", false, 0, "")
	pdf.Rect(leftMargin+137, footerStartY+7, 4, 4, "D")
	pdf.SetXY(leftMargin+137.5, footerStartY+7)
	pdf.Cell(3, 4, "X")

	pdf.SetXY(leftMargin+142.5, footerStartY+6)
	pdf.CellFormat(47.5, 6, "", "1", 0, "C", false, 0, "")
	pdf.Rect(leftMargin+185, footerStartY+7, 4, 4, "D")
	pdf.SetXY(leftMargin+185.5, footerStartY+7)
	pdf.Cell(3, 4, "X")

	pdf.SetXY(leftMargin+95, footerStartY+12)
	pdf.CellFormat(95, 6, "Customer Check", "1", 0, "L", false, 0, "")
	pdf.Rect(leftMargin+185, footerStartY+13, 4, 4, "D")
	pdf.SetXY(leftMargin+185.5, footerStartY+13)
	pdf.Cell(3, 4, "X")
	
	currentY = footerStartY + 18

	// NOTE section - calculate height needed for left side
	noteStartY := currentY
	pdf.SetFont("Arial", "B", 7)
	pdf.SetXY(leftMargin, currentY)
	pdf.MultiCell(95, 4.5, "NOTE Liability Limitation for loss or damage in this shipment may be applicable. See 49 U.S.C. 47706(f)(1)(A) and", "", "L", false)
	currentY = pdf.GetY() + 1
	pdf.SetFont("Arial", "", 6)
	pdf.SetXY(leftMargin, currentY)
	pdf.MultiCell(95, 3.5, "RECEIVED, subject to individually determined rates or contracts that have been agreed upon in writing between", "", "L", false)
	currentY = pdf.GetY() + 1
	pdf.SetXY(leftMargin, currentY)
	pdf.MultiCell(95, 3.5, "the carrier and shipper, if applicable, otherwise to the rates, classifications and rules that have been", "", "L", false)
	leftNoteFinalY := pdf.GetY()

	// Shipper Signature section (right side) - height should match left side
	pdf.SetFont("Arial", "", 6)
	pdf.SetXY(leftMargin+95, noteStartY)
	pdf.MultiCell(95, 3.5, "The carrier shall not make delivery of this shipment without payment of freight and all other lawful charges.", "1", "L", false)
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+95, pdf.GetY())
	pdf.CellFormat(95, 7, "Shipper Signature", "1", 0, "C", false, 0, "")
	rightFinalY := pdf.GetY()
	
	// Use the maximum height for currentY
	if rightFinalY > leftNoteFinalY {
		currentY = rightFinalY
	} else {
		currentY = leftNoteFinalY
	}
	currentY += 2

	// ===== SIGNATURE SECTIONS =====
	shipperSectionStart := currentY

	// Row 1: Headers
	pdf.SetFont("Arial", "B", 7)
	pdf.SetXY(leftMargin, currentY)
	pdf.CellFormat(95, 4, "SHIPPER SIGNATURE / DATE", "1", 0, "L", false, 0, "")
	pdf.SetXY(leftMargin+95, currentY)
	pdf.CellFormat(47.5, 4, "Trailer Loaded:", "1", 0, "L", false, 0, "")
	pdf.SetXY(leftMargin+142.5, currentY)
	pdf.CellFormat(47.5, 4, "Freight Counted:", "1", 0, "L", false, 0, "")
	currentY += 4

	// Row 2: Left disclaimer and right checkboxes row 1
	pdf.SetFont("Arial", "", 5)
	pdf.SetXY(leftMargin, currentY)
	pdf.MultiCell(95, 2.5, "This is to certify that the above named materials are properly classified, described, packaged, marked and labeled, and are in proper condition for transportation according to the applicable regulations of the Department of Transportation.", "LR", "L", false)
	leftFinalY := pdf.GetY()
	
	pdf.SetFont("Arial", "", 7)
	rightY := currentY
	pdf.Rect(leftMargin+95, rightY, 4, 4, "D")
	pdf.SetXY(leftMargin+96, rightY)
	pdf.Cell(4, 4, "X")
	pdf.SetXY(leftMargin+100, rightY)
	pdf.CellFormat(42.5, 4, "By Shipper", "R", 0, "L", false, 0, "")

	pdf.Rect(leftMargin+142.5, rightY, 4, 4, "D")
	pdf.SetXY(leftMargin+143.5, rightY)
	pdf.Cell(4, 4, "X")
	pdf.SetXY(leftMargin+147.5, rightY)
	pdf.CellFormat(42.5, 4, "By Shipper", "R", 0, "L", false, 0, "")
	
	// Use the taller of the two sections
	if leftFinalY > currentY+10 {
		currentY = leftFinalY
	} else {
		currentY += 10
	}

	// Row 3: Left empty, right checkboxes row 2
	pdf.SetFont("Arial", "", 5)
	row3StartY := currentY
	pdf.SetXY(leftMargin, currentY)
	pdf.CellFormat(95, 2.5, "", "LRB", 0, "L", false, 0, "")
	leftY := currentY + 2.5
	
	pdf.SetFont("Arial", "", 7)
	// Right side - By Driver checkboxes
	pdf.SetXY(leftMargin+95, row3StartY)
	pdf.CellFormat(47.5, 6, "", "R", 0, "L", false, 0, "")
	pdf.Rect(leftMargin+95, row3StartY+2, 4, 4, "D")
	pdf.SetXY(leftMargin+100, row3StartY+2)
	pdf.CellFormat(42.5, 4, "By Driver", "R", 0, "L", false, 0, "")

	pdf.SetXY(leftMargin+142.5, row3StartY)
	pdf.CellFormat(47.5, 6, "", "R", 0, "L", false, 0, "")
	pdf.Rect(leftMargin+142.5, row3StartY+2, 4, 4, "D")
	pdf.SetXY(leftMargin+143.5, row3StartY+2)
	pdf.Cell(4, 4, "X")
	pdf.SetXY(leftMargin+147.5, row3StartY+2)
	pdf.CellFormat(42.5, 4, "By Driver/pallets", "R", 0, "L", false, 0, "")
	rightRow3Y := row3StartY + 6
	
	// Use the taller of left or right side for Row 3
	if rightRow3Y > leftY {
		currentY = rightRow3Y
	} else {
		currentY = leftY
	}

	// Row 4: Left empty, right "said to contain"
	row4StartY := currentY
	pdf.SetXY(leftMargin, row4StartY)
	pdf.CellFormat(95, 4, "", "LRB", 0, "L", false, 0, "")
	leftY4 := row4StartY + 4
	pdf.SetXY(leftMargin+95, row4StartY)
	pdf.CellFormat(47.5, 4, "", "RB", 0, "L", false, 0, "")
	pdf.SetXY(leftMargin+142.5, row4StartY)
	pdf.CellFormat(47.5, 4, "said to contain", "TRB", 0, "C", false, 0, "")
	currentY = leftY4

	// Row 5: Carrier Signature header
	pdf.SetFont("Arial", "B", 7)
	pdf.SetXY(leftMargin+95, currentY)
	pdf.CellFormat(95, 4, "CARRIER SIGNATURE / PICKUP", "1", 0, "L", false, 0, "")
	currentY += 4
	
	// Carrier disclaimer and final message
	pdf.SetFont("Arial", "", 5)
	pdf.SetXY(leftMargin+95, currentY)
	pdf.MultiCell(95, 2.5, "Carrier acknowledges receipt of packages and required placards. Carrier certifies emergency response information was made available and/or carrier has the U.S. DOT emergency response guidebook or equivalent documentation in the vehicle.", "LR", "L", false)
	currentY = pdf.GetY()
	pdf.SetFont("Arial", "", 6)
	pdf.SetXY(leftMargin+95, currentY)
	pdf.MultiCell(95, 3, "Property described above is received in good order, except as noted.", "LRB", "L", false)

	// Draw shipper section border
	shipperSectionHeight := currentY - shipperSectionStart
	pdf.Rect(leftMargin, shipperSectionStart, 95, shipperSectionHeight, "D")
}
