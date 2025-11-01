package main

import (
	"github.com/looksocial/gofpdf"
	"github.com/looksocial/gofpdf/contrib/barcode"
	"github.com/looksocial/gofpdf/internal/example"
	"github.com/boombuler/barcode/qr"
)

func main() {
	// Create new PDF - A4 portrait
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetAutoPageBreak(false, 0)

	// Colors
	darkBlue := []int{0, 51, 102}       // Dark blue for headers
	lightGray := []int{240, 240, 240}   // Light gray for section headers
	mediumGray := []int{128, 128, 128}  // Medium gray for labels
	black := []int{0, 0, 0}             // Black for text

	// Page setup
	leftMargin := 10.0
	rightMargin := 200.0
	currentY := 10.0

	// ==================== HEADER SECTION ====================
	// Logo/Brand area
	pdf.SetFillColor(darkBlue[0], darkBlue[1], darkBlue[2])
	pdf.Rect(leftMargin, currentY, 50, 20, "F")

	pdf.SetFont("Arial", "B", 24)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetXY(leftMargin+5, currentY+6)
	pdf.Cell(40, 8, "gofpdf")

	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+5, currentY+14)
	pdf.Cell(40, 4, "Go PDF Library")

	// Booking Acknowledgement Title
	pdf.SetFont("Arial", "B", 16)
	pdf.SetTextColor(black[0], black[1], black[2])
	pdf.SetXY(70, currentY+8)
	pdf.Cell(60, 8, "Booking Acknowledgement")

	// Booking Number, Date, Version
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(150, currentY)
	pdf.Cell(30, 4, "BOOKING NUMBER:")
	pdf.SetFont("Arial", "", 9)
	pdf.SetXY(180, currentY)
	pdf.Cell(20, 4, "2715061641")

	// QR Code
	qrKey := barcode.RegisterQR(pdf, "https://github.com/looksocial/gofpdf/2715061641", qr.M, qr.Auto)
	barcode.Barcode(pdf, qrKey, 150, currentY+5, 15, 15, false)

	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(170, currentY+5)
	pdf.Cell(20, 4, "DATE:")
	pdf.SetFont("Arial", "", 9)
	pdf.SetXY(180, currentY+5)
	pdf.Cell(20, 4, "03 Mar 2025 17:38")

	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(170, currentY+10)
	pdf.Cell(20, 4, "VERSION:")
	pdf.SetFont("Arial", "", 9)
	pdf.SetXY(188, currentY+10)
	pdf.Cell(10, 4, "3")

	currentY += 22

	// Contact Information
	pdf.SetFillColor(lightGray[0], lightGray[1], lightGray[2])
	pdf.Rect(leftMargin, currentY, rightMargin-leftMargin, 6, "F")
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(black[0], black[1], black[2])
	pdf.SetXY(leftMargin+2, currentY+1)
	pdf.Cell(60, 4, "FROM: gofpdf Library")
	pdf.SetXY(leftMargin+70, currentY+1)
	pdf.Cell(60, 4, "CONTACT: +1 234 567 8900")
	pdf.SetXY(leftMargin+140, currentY+1)
	pdf.Cell(60, 4, "EMAIL: info@gofpdf.com")

	currentY += 8

	// Info message
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin, currentY)
	pdf.Cell(180, 4, "Track Your Shipments 24/7 via gofpdf Portal or Mobile App! Available through your app store.")
	currentY += 6

	// ==================== BOOKING ACKNOWLEDGEMENT REMARK ====================
	renderSectionHeader(pdf, "BOOKING ACKNOWLEDGEMENT REMARK", leftMargin, currentY, rightMargin-leftMargin, lightGray)
	currentY += 6
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+2, currentY)
	pdf.MultiCell(180, 4, "Thank you for choosing gofpdf for your PDF generation needs. This booking has been confirmed and is ready for processing.", "", "L", false)
	currentY += 10

	// ==================== BOOKING DETAILS ====================
	renderInfoRow(pdf, "BOOKING NUMBER:", "2715061641", leftMargin, currentY)
	currentY += 5
	renderInfoRow(pdf, "BOOKING STATUS:", "Confirmed", leftMargin, currentY)
	currentY += 5
	renderInfoRow(pdf, "RATE AGREEMENT NUMBER:", "00103751", leftMargin, currentY)
	currentY += 7

	// ==================== EXTERNAL REFERENCE INFORMATION ====================
	renderSectionHeader(pdf, "EXTERNAL REFERENCE INFORMATION", leftMargin, currentY, rightMargin-leftMargin, lightGray)
	currentY += 6
	renderInfoRow(pdf, "CS Reference Number:", "CS6476839896", leftMargin, currentY)
	currentY += 7

	// ==================== PARTIES INFORMATION ====================
	renderSectionHeader(pdf, "PARTIES INFORMATION", leftMargin, currentY, rightMargin-leftMargin, lightGray)
	currentY += 6
	renderInfoRow(pdf, "BOOKING PARTY:", "MERIDIAN GLOBAL SHIPPING AND LOGISTICS PVT LTD", leftMargin, currentY)
	currentY += 5
	renderInfoRow(pdf, "FORWARDER:", "MERIDIAN GLOBAL SHIPPING AND LOGISTICS PVT LTD", leftMargin, currentY)
	currentY += 5
	renderInfoRow(pdf, "SHIPPER:", "MERIDIAN GLOBAL SHIPPING AND LOGISTICS PVT LTD", leftMargin, currentY)
	currentY += 7

	// ==================== ROUTE INFORMATION ====================
	renderSectionHeader(pdf, "ROUTE INFORMATION", leftMargin, currentY, rightMargin-leftMargin, lightGray)
	currentY += 6
	renderInfoRow(pdf, "TOTAL BOOKING CONTAINER QTY SIZE/TYPE:", "1 x 40' Hi-Cube Container", leftMargin, currentY)
	currentY += 5

	// Place of Receipt
	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(leftMargin+2, currentY)
	pdf.Cell(60, 4, "PLACE OF RECEIPT:")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+55, currentY)
	pdf.Cell(80, 4, "Nhava Sheva, Maharashtra, India")
	currentY += 5

	// Port of Loading
	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(leftMargin+2, currentY)
	pdf.Cell(60, 4, "PORT OF LOADING:")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+55, currentY)
	pdf.Cell(80, 4, "Nhava Sheva / Nhava Sheva (India) Gateway")
	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(150, currentY)
	pdf.Cell(20, 4, "ETA:")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(165, currentY)
	pdf.Cell(30, 4, "10 Mar 2025")
	currentY += 5

	// Vessel/Voyage
	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(leftMargin+2, currentY)
	pdf.Cell(60, 4, "INTENDED VESSEL/VOYAGE:")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+55, currentY)
	pdf.Cell(60, 4, "TSINGTAO EXPRESS 070 \\W")
	currentY += 5

	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(leftMargin+2, currentY)
	pdf.Cell(60, 4, "SERVICE CODE:")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+55, currentY)
	pdf.Cell(40, 4, "IP2")
	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(100, currentY)
	pdf.Cell(30, 4, "VESSEL FLAG:")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(125, currentY)
	pdf.Cell(25, 4, "Germany")
	currentY += 5

	// Port of Discharge
	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(leftMargin+2, currentY)
	pdf.Cell(60, 4, "PORT OF DISCHARGE:")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+55, currentY)
	pdf.Cell(80, 4, "Hamburg / HHLA Cntr-Trml Altenwerder GmbH")
	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(150, currentY)
	pdf.Cell(20, 4, "ETA:")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(165, currentY)
	pdf.Cell(30, 4, "29 Mar 2025")
	currentY += 5

	// Final Destination
	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(leftMargin+2, currentY)
	pdf.Cell(60, 4, "FINAL DESTINATION:")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+55, currentY)
	pdf.Cell(80, 4, "Hamburg, Hamburg, Germany")
	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(150, currentY)
	pdf.Cell(20, 4, "ETA:")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(165, currentY)
	pdf.Cell(30, 4, "29 Mar 2025")
	currentY += 5

	// Cargo Availability
	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(leftMargin+2, currentY)
	pdf.Cell(70, 4, "ESTIMATED CARGO AVAILABILITY AT DESTINATION HUB:")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+80, currentY)
	pdf.Cell(60, 4, "31 Mar 2025 06:00")
	currentY += 5

	// Cut-off dates
	renderInfoRow(pdf, "INTENDED CY CUT-OFF:", "10 Mar 2025 12:00", leftMargin, currentY)
	currentY += 5
	renderInfoRow(pdf, "INTENDED SI/eS CUT-OFF:", "", leftMargin, currentY)
	currentY += 5
	renderInfoRow(pdf, "INTENDED VGM CUT-OFF:", "09 Mar 2025 6:00:00 PM(IST)", leftMargin, currentY)
	currentY += 5

	// Warning message
	pdf.SetFont("Arial", "I", 7)
	pdf.SetXY(leftMargin+2, currentY)
	pdf.MultiCell(180, 3, "LATE AND/OR INCOMPLETE SHIPPING INSTRUCTION SUBMISSION MAY RESULT IN CONTAINER(S) SHORT SHIPMENT AND LATE SI SUBMISSION CHARGES", "", "L", false)
	currentY += 8

	// ==================== CARGO INFORMATION ====================
	renderSectionHeader(pdf, "CARGO INFORMATION", leftMargin, currentY, rightMargin-leftMargin, lightGray)
	currentY += 6
	renderInfoRow(pdf, "CARGO NATURE:", "General", leftMargin, currentY)
	currentY += 5
	renderInfoRow(pdf, "CARGO DESCRIPTION:", "GFL - POLYTETRA ETHYLEFLUORO", leftMargin, currentY)
	currentY += 7

	// Booking details
	renderInfoRow(pdf, "BOOKING QTY SIZE/TYPE:", "1 X 40' Hi-Cube Container", leftMargin, currentY)
	currentY += 5
	renderInfoRow(pdf, "CARGO WEIGHT:", "24000 KG", leftMargin, currentY)
	currentY += 5

	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(leftMargin+2, currentY)
	pdf.Cell(60, 4, "OUTBOUND DELIVERY MODE:")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+55, currentY)
	pdf.Cell(60, 4, "CY by transport mode Truck")
	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(130, currentY)
	pdf.Cell(30, 4, "TRAFFIC MODE:")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(160, currentY)
	pdf.Cell(30, 4, "FCL / FCL")
	currentY += 5

	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(leftMargin+2, currentY)
	pdf.Cell(60, 4, "INBOUND DELIVERY MODE:")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+55, currentY)
	pdf.Cell(60, 4, "CY by transport mode Truck")
	currentY += 5

	renderInfoRow(pdf, "EMPTY PICK UP DATE/TIME:", "09 Mar 2025 12:00", leftMargin, currentY)
	currentY += 7

	// ==================== PICKUP AND RETURN LOCATIONS ====================
	renderSectionHeader(pdf, "EMPTY PICK UP LOCATION", leftMargin, currentY, 90, lightGray)
	renderSectionHeader(pdf, "FULL RETURN LOCATION", leftMargin+95, currentY, 95, lightGray)
	currentY += 6

	// Empty pickup location
	pdf.SetFont("Arial", "", 7)
	pdf.SetXY(leftMargin+2, currentY)
	pdf.MultiCell(85, 3, "Bhavani Shipping (I) PVT LTD.\nOpposite Coldman Warehousing\nDist Raigad, Navi Mumbai\n412066\nINDIA", "", "L", false)

	// Full return location
	pdf.SetFont("Arial", "", 7)
	pdf.SetXY(leftMargin+97, currentY)
	pdf.MultiCell(88, 3, "Nhava Sheva (India) Gateway Trml\nNavi Mumbai, Maharashtra\n400707 India\n\nContact: Martin John", "", "L", false)

	currentY += 20

	// ==================== FOOTER ====================
	currentY = 275
	pdf.SetFont("Arial", "", 7)
	pdf.SetTextColor(mediumGray[0], mediumGray[1], mediumGray[2])
	pdf.SetXY(leftMargin, currentY)
	pdf.CellFormat(190, 3, "Generated with gofpdf - Go PDF Library", "T", 0, "C", false, 0, "")
	currentY += 4
	pdf.SetXY(leftMargin, currentY)
	pdf.Cell(190, 3, "https://github.com/looksocial/gofpdf")
	pdf.SetXY(190/2+leftMargin, currentY)
	pdf.CellFormat(0, 3, "", "", 0, "C", false, 0, "")

	// Save PDF
	outputPath := example.PdfFile("booking_acknowledgement_demo.pdf")
	err := pdf.OutputFileAndClose(outputPath)
	if err != nil {
		panic(err)
	}
}

// renderSectionHeader renders a section header with background color
func renderSectionHeader(pdf *gofpdf.Fpdf, title string, x, y, width float64, bgColor []int) {
	pdf.SetFillColor(bgColor[0], bgColor[1], bgColor[2])
	pdf.Rect(x, y, width, 6, "F")
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetXY(x+2, y+1)
	pdf.Cell(width-4, 4, title)
}

// renderInfoRow renders a label-value pair
func renderInfoRow(pdf *gofpdf.Fpdf, label, value string, x, y float64) {
	pdf.SetFont("Arial", "B", 8)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetXY(x+2, y)
	pdf.Cell(60, 4, label)
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(x+55, y)
	pdf.Cell(120, 4, value)
}
