package components

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/looksocial/gofpdf"
	"github.com/looksocial/gofpdf/table"
)

// renderCarrierHeaders renders the Carrier Information section headers
func renderCarrierHeaders(pdf *gofpdf.Fpdf, leftMargin, currentY, contentWidth float64, black, white []int) float64 {
	// Black title bar
	pdf.SetFillColor(black[0], black[1], black[2])
	pdf.Rect(leftMargin, currentY, contentWidth, 6, "F")
	pdf.SetTextColor(white[0], white[1], white[2])
	pdf.SetFont("Arial", "B", 10)
	pdf.SetXY(leftMargin, currentY+1)
	pdf.CellFormat(contentWidth, 5, "CARRIER INFORMATION", "", 0, "C", false, 0, "")

	pdf.SetTextColor(black[0], black[1], black[2])
	currentY += 6

	// Carrier information table headers
	pdf.SetFillColor(white[0], white[1], white[2])
	pdf.SetFont("Arial", "B", 7)

	// HANDLING header spanning two sub-columns
	pdf.SetXY(leftMargin, currentY)
	pdf.CellFormat(30, 10, "HANDLING", "1", 0, "C", false, 0, "")
	// PACKAGE header spanning two sub-columns
	pdf.SetXY(leftMargin+30, currentY)
	pdf.CellFormat(30, 10, "PACKAGE", "1", 0, "C", false, 0, "")
	// WEIGHT header
	pdf.SetXY(leftMargin+60, currentY)
	pdf.CellFormat(20, 10, "WEIGHT", "1", 0, "C", false, 0, "")

	// H.M. (X) column with smaller text
	pdf.SetXY(leftMargin+80, currentY)
	pdf.CellFormat(10, 10, "", "1", 0, "C", false, 0, "")
	pdf.SetFont("Arial", "B", 6)
	pdf.SetXY(leftMargin+80, currentY+2)
	pdf.CellFormat(10, 3, "H.M.", "", 0, "C", false, 0, "")
	pdf.SetXY(leftMargin+80, currentY+5)
	pdf.CellFormat(10, 3, "(X)", "", 0, "C", false, 0, "")
	pdf.SetFont("Arial", "B", 7)

	// COMMODITY DESCRIPTION with smaller font to fit text
	pdf.SetXY(leftMargin+90, currentY)
	pdf.CellFormat(70, 10, "", "1", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 6)
	pdf.SetXY(leftMargin+91, currentY+0.5)
	pdf.CellFormat(68, 2, "COMMODITY DESCRIPTION", "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "", 4.5)
	pdf.SetXY(leftMargin+91, currentY+2.5)
	pdf.MultiCell(68, 1.3, "Commodities requiring special or additional care or attention in handling or stowing must be so marked and packaged as to ensure safe transportation with ordinary care. See Section 2(e) of NMFC Item 360", "", "L", false)

	// LTL ONLY header spanning two sub-columns
	pdf.SetFont("Arial", "B", 7)
	pdf.SetXY(leftMargin+160, currentY)
	pdf.CellFormat(30, 10, "LTL ONLY", "1", 0, "C", false, 0, "")

	currentY += 10

	// Sub-headers for HANDLING, PACKAGE, and LTL ONLY sections
	pdf.SetFont("Arial", "B", 7)
	pdf.SetXY(leftMargin, currentY)
	pdf.CellFormat(7, 5, "QTY", "1", 0, "C", false, 0, "")
	pdf.SetXY(leftMargin+7, currentY)
	pdf.CellFormat(23, 5, "TYPE", "1", 0, "C", false, 0, "")
	pdf.SetXY(leftMargin+30, currentY)
	pdf.CellFormat(7, 5, "QTY", "1", 0, "C", false, 0, "")
	pdf.SetXY(leftMargin+37, currentY)
	pdf.CellFormat(23, 5, "TYPE", "1", 0, "C", false, 0, "")

	// Empty cells for WEIGHT, H.M., and COMMODITY DESCRIPTION (no sub-headers needed)
	pdf.SetXY(leftMargin+60, currentY)
	pdf.CellFormat(20, 5, "", "1", 0, "C", false, 0, "")
	pdf.SetXY(leftMargin+80, currentY)
	pdf.CellFormat(10, 5, "", "1", 0, "C", false, 0, "")
	pdf.SetXY(leftMargin+90, currentY)
	pdf.CellFormat(70, 5, "", "1", 0, "C", false, 0, "")

	// LTL ONLY sub-headers
	pdf.SetXY(leftMargin+160, currentY)
	pdf.CellFormat(15, 5, "NMFC #", "1", 0, "C", false, 0, "")
	pdf.SetXY(leftMargin+175, currentY)
	pdf.CellFormat(15, 5, "CLASS", "1", 0, "C", false, 0, "")

	currentY += 5

	return currentY
}

// RenderCustomerOrderInformation renders the Customer Order Information table
func RenderCustomerOrderInformation(pdf *gofpdf.Fpdf, leftMargin, currentY, contentWidth float64, black, white, lightGray []int) float64 {
	// Header
	pdf.SetFillColor(black[0], black[1], black[2])
	pdf.Rect(leftMargin, currentY, contentWidth, 6, "F")
	pdf.SetTextColor(white[0], white[1], white[2])
	pdf.SetFont("Arial", "B", 10)
	pdf.SetXY(leftMargin, currentY+1)
	pdf.CellFormat(contentWidth, 5, "CUSTOMER ORDER INFORMATION", "", 0, "C", false, 0, "")

	pdf.SetTextColor(black[0], black[1], black[2])
	currentY += 6

	// Create Customer Order Information table using table component
	pdf.SetFont("Arial", "", 7)
	orderColumns := []table.Column{
		{Key: "order_number", Label: "CUSTOMER ORDER\nNUMBER", Width: 48, Align: "L"},
		{Key: "pkgs", Label: "# PKGS", Width: 20, Align: "C"},
		{Key: "weight", Label: "WEIGHT", Width: 24, Align: "C"},
		{Key: "pallet", Label: "PALLET/SLIP", Width: 28, Align: "C"},
		{Key: "info", Label: "ADDITIONAL SHIPPER INFO", Width: 70, Align: "L"},
	}

	orderTable := table.NewTable(pdf, orderColumns).
		WithStartPosition(leftMargin, currentY).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FontSize:  7,
			FillColor: lightGray,
			TextColor: black,
		}).
		WithDataStyle(table.CellStyle{
			Border:   "1",
			FontSize: 7,
		}).
		WithRowHeight(6).
		WithPageBreakMode(true).
		WithRepeatHeader(true).
		WithPageBreakMargin(55) // Reserve space for footer section

	orderTable.AddHeader()

	// Add data rows
	orderData := []map[string]interface{}{
		{"order_number": "GOFPDF-2024-001", "pkgs": "4", "weight": "120", "pallet": "", "info": "DC-GOFPDF-001"},
		{"order_number": "GOFPDF-2024-002", "pkgs": "6", "weight": "180", "pallet": "", "info": "DC-GOFPDF-002"},
		{"order_number": "GOFPDF-2024-003", "pkgs": "2", "weight": "60", "pallet": "", "info": "DC-GOFPDF-003"},
		{"order_number": "GOFPDF-2024-004", "pkgs": "1", "weight": "30", "pallet": "X", "info": "DC-GOFPDF-004"},
		{"order_number": "GOFPDF-2024-005", "pkgs": "", "weight": "", "pallet": "X", "info": "DC-GOFPDF-005"},
	}

	for _, row := range orderData {
		orderTable.AddRow(row)
	}

	// Add grand total row
	orderTable.AddTotalRow("GRAND TOTAL", map[string]interface{}{
		"pkgs":   "13",
		"weight": "390",
		"pallet": "",
		"info":   "",
	}, table.CellStyle{
		Border:    "1",
		Bold:      true,
		FontSize:  7,
		FillColor: lightGray,
	})

	return pdf.GetY()
}

// RenderCarrierInformation renders the Carrier Information table
func RenderCarrierInformation(pdf *gofpdf.Fpdf, leftMargin, currentY, contentWidth float64, black, white, lightGray []int) float64 {
	// Render header section using helper function
	currentY = renderCarrierHeaders(pdf, leftMargin, currentY, contentWidth, black, white)

	// Data rows using table component
	pdf.SetFont("Arial", "", 7)

	// Create table for carrier information data (without header, already drawn above)
	carrierColumns := []table.Column{
		{Key: "h_qty", Label: "", Width: 7, Align: "C"},
		{Key: "h_type", Label: "", Width: 23, Align: "C"},
		{Key: "p_qty", Label: "", Width: 7, Align: "C"},
		{Key: "p_type", Label: "", Width: 23, Align: "C"},
		{Key: "weight", Label: "", Width: 20, Align: "C"},
		{Key: "hm", Label: "", Width: 10, Align: "C"},
		{Key: "commodity", Label: "", Width: 70, Align: "L"},
		{Key: "nmfc", Label: "", Width: 15, Align: "C"},
		{Key: "class", Label: "", Width: 15, Align: "C"},
	}

	carrierTable := table.NewTable(pdf, carrierColumns).
		WithStartPosition(leftMargin, currentY).
		WithDataStyle(table.CellStyle{
			Border:   "1",
			FontSize: 7,
		}).
		WithRowHeight(8).
		WithPageBreakMode(true).
		WithRepeatHeader(false).
		WithPageBreakMargin(55). // Reserve space for footer section
		WithCustomRepeatHeader(func() float64 {
			// Re-render custom headers on page break
			return renderCarrierHeaders(pdf, leftMargin, pdf.GetY(), contentWidth, black, white)
		})

	// Add data rows
	carrierData := []map[string]interface{}{
		{"h_qty": "2", "h_type": "PLT", "p_qty": "200", "p_type": "CTN", "weight": "20,000", "hm": "", "commodity": "gofpdf Demo Electronics", "nmfc": "", "class": ""},
		{"h_qty": "3", "h_type": "PLT", "p_qty": "150", "p_type": "BOX", "weight": "15,000", "hm": "X", "commodity": "gofpdf Demo Machinery Parts", "nmfc": "GOFPDF-001", "class": "70"},
		{"h_qty": "1", "h_type": "DRM", "p_qty": "100", "p_type": "CTG", "weight": "10,000", "hm": "", "commodity": "gofpdf Demo Furniture", "nmfc": "GOFPDF-002", "class": "77.5"},
		{"h_qty": "4", "h_type": "PLT", "p_qty": "250", "p_type": "BOX", "weight": "25,000", "hm": "X", "commodity": "gofpdf Demo Supplies", "nmfc": "GOFPDF-003", "class": "85"},
		{"h_qty": "2", "h_type": "BND", "p_qty": "100", "p_type": "CTN", "weight": "16,000", "hm": "", "commodity": "gofpdf Demo Materials", "nmfc": "GOFPDF-004", "class": "60"},
		{"h_qty": "5", "h_type": "PLT", "p_qty": "300", "p_type": "BOX", "weight": "30,000", "hm": "X", "commodity": "gofpdf Demo Equipment", "nmfc": "GOFPDF-005", "class": "92.5"},
		{"h_qty": "1", "h_type": "PLT", "p_qty": "80", "p_type": "CTN", "weight": "8,000", "hm": "", "commodity": "gofpdf Demo Tools", "nmfc": "GOFPDF-006", "class": "55"},
		{"h_qty": "3", "h_type": "PLT", "p_qty": "220", "p_type": "BOX", "weight": "22,000", "hm": "X", "commodity": "gofpdf Demo Components", "nmfc": "GOFPDF-007", "class": "75"},
	}

	var totalHqty, totalPqty int
	var totalWeight int

	for _, row := range carrierData {
		carrierTable.AddRow(row)
		// Parse and accumulate totals
		if hqty := parseNumber(row["h_qty"].(string)); hqty > 0 {
			totalHqty += hqty
		}
		if pqty := parseNumber(row["p_qty"].(string)); pqty > 0 {
			totalPqty += pqty
		}
		if wt := parseNumber(row["weight"].(string)); wt > 0 {
			totalWeight += wt
		}
	}

	// Add grand total row
	carrierTable.AddTotalRow("GRAND TOTAL", map[string]interface{}{
		"h_qty":  fmt.Sprintf("%d", totalHqty),
		"h_type": "",
		"p_qty":  fmt.Sprintf("%d", totalPqty),
		"p_type": "",
		"weight": formatNumber(totalWeight),
		"hm":     "",
		"nmfc":   "",
		"class":  "",
	}, table.CellStyle{
		Border:    "1",
		Bold:      true,
		FontSize:  7,
		FillColor: lightGray,
	})

	return pdf.GetY()
}

// parseNumber parses a string number (may contain commas) to an integer
func parseNumber(s string) int {
	if s == "" {
		return 0
	}
	// Remove commas
	s = strings.ReplaceAll(s, ",", "")
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return num
}

// formatNumber formats an integer with comma separators
func formatNumber(n int) string {
	if n == 0 {
		return ""
	}
	str := strconv.Itoa(n)
	// Add comma separators
	result := ""
	for i, c := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result += ","
		}
		result += string(c)
	}
	return result
}
