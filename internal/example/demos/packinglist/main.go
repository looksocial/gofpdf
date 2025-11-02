//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/looksocial/gofpdf"
	"github.com/looksocial/gofpdf/table"
)

// PackingListData holds all the data for a packing list
type PackingListData struct {
	// Header data
	PageNumber    int
	TotalPages    int
	ExporterName  string
	ExporterAddr1 string
	ExporterAddr2 string
	ExporterAddr3 string
	ContactName   string
	TaxID         string
	Email         string

	// Document info
	InvoiceNumber  string
	InvoiceDate    string
	BillOfLading   string
	Reference      string
	BuyerReference string
	BuyerIfNotCons string

	// Consignee
	ConsigneeName string
	ConsigneeAddr string

	// Shipping details
	MethodOfDispatch string
	TypeOfShipment   string
	CountryOfOrigin  string
	CountryOfDest    string
	VesselAircraft   string
	VoyageNo         string
	PackingInfo      string
	PortOfLoading    string
	DateOfDeparture  string
	FinalDestination string

	// Products
	Products []ProductItem

	// Footer data
	SignatoryCompany  string
	AuthorizedName    string
	AuthorizedSurname string
}

// ProductItem represents a product in the packing list
type ProductItem struct {
	ProductCode  string
	Description  string
	UnitQuantity string
	KindNoOfPkg  string
	NetWeight    string
	GrossWeight  string
	Measurements string
}

// Colors for the document
type Colors struct {
	PrimaryRed []int
	LightGray  []int
	White      []int
	TextDark   []int
	BorderGray []int
}

func main() {
	// Create new PDF
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Sample data for single page packing list
	data := PackingListData{
		PageNumber:       1,
		TotalPages:       1,
		ExporterName:     "ABC Exports",
		ExporterAddr1:    "123 California Blvd",
		ExporterAddr2:    "Longbeach, California, 90807",
		ExporterAddr3:    "United States",
		ContactName:      "Randy Clarke",
		TaxID:            "82377112",
		Email:            "randy@abcexport.com",
		InvoiceNumber:    "34567",
		InvoiceDate:      "30 Jul 2022",
		BillOfLading:     "LONSYD123456",
		Reference:        "34567",
		BuyerReference:   "PO223",
		BuyerIfNotCons:   "",
		ConsigneeName:    "XYZ Imports",
		ConsigneeAddr:    "456 Business Street, Brisbane, Queensland, 4814, Australia",
		MethodOfDispatch: "Sea",
		TypeOfShipment:   "FCL",
		CountryOfOrigin:  "United States",
		CountryOfDest:    "Australia",
		VesselAircraft:   "MAERSK",
		VoyageNo:         "V0015",
		PackingInfo:      "",
		PortOfLoading:    "Long Beach",
		DateOfDeparture:  "04 Jul 2022",
		FinalDestination: "Port of AUSTRALIA",
		Products: []ProductItem{
			{
				ProductCode:  "B-STOOL",
				Description:  "BAR STOOL ALUMINUM 500 X 100 X 100MM STAINLESS STEEL",
				UnitQuantity: "150",
				KindNoOfPkg:  "PALLET X II",
				NetWeight:    "1,500",
				GrossWeight:  "1,600",
				Measurements: "12",
			},
			{
				ProductCode:  "B-TABLE",
				Description:  "BAR TABLE ALUMINUM 1000 X 600 X 400MM STAINLESS STEEL",
				UnitQuantity: "75",
				KindNoOfPkg:  "PALLET X II",
				NetWeight:    "1,500",
				GrossWeight:  "1,575",
				Measurements: "15",
			},
		},
		SignatoryCompany:  "ABC Exports",
		AuthorizedName:    "Randy",
		AuthorizedSurname: "Clarke",
	}

	// Add a page
	pdf.AddPage()

	// Generate the packing list
	generatePackingList(pdf, data)

	// Ensure pdf directory exists
	pdfDir := filepath.Join("..", "..", "..", "..", "pdf")
	err := os.MkdirAll(pdfDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Save to file in pdf directory
	outputPath := filepath.Join(pdfDir, "packing_list.pdf")
	err = pdf.OutputFileAndClose(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	// Get absolute path for display
	absPath, _ := filepath.Abs(outputPath)
	fmt.Printf("✓ Packing List created successfully: %s\n", absPath)
}

func generatePackingList(pdf *gofpdf.Fpdf, data PackingListData) {
	colors := Colors{
		PrimaryRed: []int{230, 0, 0},
		LightGray:  []int{240, 240, 240},
		White:      []int{255, 255, 255},
		TextDark:   []int{0, 0, 0},
		BorderGray: []int{0, 0, 0},
	}

	// Render header
	headerHeight := renderHeader(pdf, data, colors)

	// Render body (with page break handling)
	bodyStartY := headerHeight + 2
	renderBody(pdf, data, colors, bodyStartY)

	// Check if footer fits on current page, if not, move to next page
	_, pageHeight := pdf.GetPageSize()
	_, _, _, bottomMargin := pdf.GetMargins()
	footerHeight := 30.0 // Approximate footer height
	currentY := pdf.GetY()
	availableSpace := pageHeight - bottomMargin - currentY

	if availableSpace < footerHeight {
		// Need new page for footer
		pdf.AddPage()
		// Re-render header on new page for footer
		data.PageNumber = data.PageNumber + 1
		renderHeader(pdf, data, colors)
		pdf.SetY(headerHeight + 2)
	}

	// Render footer
	renderFooter(pdf, data, colors, pdf.GetY())
}

func renderHeader(pdf *gofpdf.Fpdf, data PackingListData, colors Colors) float64 {
	pageWidth, _ := pdf.GetPageSize()
	leftMargin, topMargin, rightMargin, _ := pdf.GetMargins()

	currentY := topMargin

	// ===== HEADER SECTION =====
	// Logo
	pdf.SetFillColor(colors.PrimaryRed[0], colors.PrimaryRed[1], colors.PrimaryRed[2])
	pdf.Circle(leftMargin+8, currentY+8, 7, "F")
	pdf.SetTextColor(colors.White[0], colors.White[1], colors.White[2])
	pdf.SetFont("Arial", "B", 7)
	pdf.SetXY(leftMargin+4.5, currentY+6)
	pdf.Cell(7, 4, "ABC")

	// Title
	pdf.SetTextColor(colors.TextDark[0], colors.TextDark[1], colors.TextDark[2])
	pdf.SetFont("Arial", "B", 18)
	pdf.SetXY(leftMargin+60, currentY+2)
	pdf.Cell(70, 8, "PACKING LIST")

	// Page info
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(pageWidth-rightMargin-20, currentY)
	pdf.Cell(20, 4, "Pages")
	pdf.SetXY(pageWidth-rightMargin-20, currentY+4)
	pdf.Cell(20, 4, fmt.Sprintf("%d of %d", data.PageNumber, data.TotalPages))

	currentY += 18

	// ===== EXPORTER AND DOCUMENT INFO BOXES =====
	boxHeight := 42.0
	leftBoxWidth := 90.0
	rightBoxWidth := pageWidth - leftMargin - rightMargin - leftBoxWidth

	// Left box - Exporter
	pdf.SetDrawColor(colors.BorderGray[0], colors.BorderGray[1], colors.BorderGray[2])
	pdf.SetLineWidth(0.5)
	pdf.Rect(leftMargin, currentY, leftBoxWidth, boxHeight, "D")

	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(leftMargin+2, currentY+2)
	pdf.Cell(50, 4, "Exporter")

	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(leftMargin+2, currentY+7)
	pdf.Cell(80, 4, data.ExporterName)

	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+2, currentY+12)
	pdf.Cell(80, 3.5, data.ExporterAddr1)
	pdf.SetXY(leftMargin+2, currentY+16)
	pdf.Cell(80, 3.5, data.ExporterAddr2)
	pdf.SetXY(leftMargin+2, currentY+20)
	pdf.Cell(80, 3.5, data.ExporterAddr3)

	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(leftMargin+2, currentY+26)
	pdf.Cell(30, 3.5, data.ContactName)
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+2, currentY+30)
	pdf.Cell(50, 3.5, fmt.Sprintf("Company Tax ID: %s", data.TaxID))
	pdf.SetXY(leftMargin+2, currentY+34)
	pdf.Cell(50, 3.5, fmt.Sprintf("Email: %s", data.Email))

	// Right box - Document details
	rightBoxX := leftMargin + leftBoxWidth
	pdf.Rect(rightBoxX, currentY, rightBoxWidth, boxHeight, "D")

	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(rightBoxX+2, currentY+2)
	pdf.Cell(50, 3.5, "Export Invoice Number & Date")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(rightBoxX+55, currentY+2)
	pdf.Cell(30, 3.5, data.InvoiceNumber)

	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(rightBoxX+2, currentY+6)
	pdf.Cell(50, 3.5, "Bill of Lading Number")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(rightBoxX+55, currentY+6)
	pdf.Cell(30, 3.5, data.BillOfLading)

	pdf.SetXY(rightBoxX+55, currentY+10)
	pdf.Cell(30, 3.5, data.InvoiceDate)

	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(rightBoxX+2, currentY+15)
	pdf.Cell(50, 3.5, "Reference")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(rightBoxX+55, currentY+15)
	pdf.Cell(30, 3.5, "Buyer Reference")

	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(rightBoxX+2, currentY+19)
	pdf.Cell(50, 3.5, data.Reference)
	pdf.SetXY(rightBoxX+55, currentY+19)
	pdf.Cell(30, 3.5, data.BuyerReference)

	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(rightBoxX+2, currentY+25)
	pdf.Cell(50, 3.5, "Buyer (If not Consignee)")
	if data.BuyerIfNotCons != "" {
		pdf.SetFont("Arial", "", 8)
		pdf.SetXY(rightBoxX+2, currentY+29)
		pdf.Cell(50, 3.5, data.BuyerIfNotCons)
	}

	currentY += boxHeight + 2

	// ===== CONSIGNEE BOX =====
	consigneeHeight := 22.0
	pdf.Rect(leftMargin, currentY, pageWidth-leftMargin-rightMargin, consigneeHeight, "D")

	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(leftMargin+2, currentY+2)
	pdf.Cell(50, 4, "Consignee")

	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(leftMargin+2, currentY+7)
	pdf.Cell(80, 4, data.ConsigneeName)

	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+2, currentY+12)
	pdf.MultiCell(160, 3.5, data.ConsigneeAddr, "", "L", false)

	currentY += consigneeHeight + 2

	// ===== SHIPPING DETAILS TABLE =====
	currentY = renderShippingDetails(pdf, data, colors, currentY)

	return currentY
}

func renderShippingDetails(pdf *gofpdf.Fpdf, data PackingListData, colors Colors, startY float64) float64 {
	leftMargin, _, _, _ := pdf.GetMargins()

	// Calculate total width to match product table width
	// Product table columns: 25 + 60 + 22 + 22 + 22 + 22 + 22 = 195 mm
	productTableWidth := 25.0 + 60.0 + 22.0 + 22.0 + 22.0 + 22.0 + 22.0

	shipDetailY := startY
	colWidth := productTableWidth / 4 // 4 columns in shipping details table
	rowHeight := 7.0

	// Row 1
	pdf.SetFillColor(colors.LightGray[0], colors.LightGray[1], colors.LightGray[2])
	pdf.SetFont("Arial", "B", 8)
	shippingHeaders1 := []string{"Method of Dispatch", "Type of Shipment", "Country of Origin of Goods", "Country of Final Destination"}
	xPos := leftMargin
	for _, header := range shippingHeaders1 {
		pdf.SetXY(xPos, shipDetailY)
		pdf.CellFormat(colWidth, rowHeight, header, "1", 0, "L", true, 0, "")
		xPos += colWidth
	}

	shipDetailY += rowHeight
	pdf.SetFont("Arial", "", 8)
	shippingData1 := []string{data.MethodOfDispatch, data.TypeOfShipment, data.CountryOfOrigin, data.CountryOfDest}
	xPos = leftMargin
	for _, d := range shippingData1 {
		pdf.SetXY(xPos, shipDetailY)
		pdf.CellFormat(colWidth, rowHeight, d, "1", 0, "L", false, 0, "")
		xPos += colWidth
	}

	// Row 2
	shipDetailY += rowHeight
	pdf.SetFillColor(colors.LightGray[0], colors.LightGray[1], colors.LightGray[2])
	pdf.SetFont("Arial", "B", 8)
	shippingHeaders2 := []string{"Vessel/Aircraft", "Voyage No", "Packing Information", ""}
	xPos = leftMargin
	for _, header := range shippingHeaders2 {
		pdf.SetXY(xPos, shipDetailY)
		pdf.CellFormat(colWidth, rowHeight, header, "1", 0, "L", true, 0, "")
		xPos += colWidth
	}

	shipDetailY += rowHeight
	pdf.SetFont("Arial", "", 8)
	shippingData2 := []string{data.VesselAircraft, data.VoyageNo, data.PackingInfo, ""}
	xPos = leftMargin
	for _, d := range shippingData2 {
		pdf.SetXY(xPos, shipDetailY)
		pdf.CellFormat(colWidth, rowHeight, d, "1", 0, "L", false, 0, "")
		xPos += colWidth
	}

	// Row 3
	shipDetailY += rowHeight
	pdf.SetFillColor(colors.LightGray[0], colors.LightGray[1], colors.LightGray[2])
	pdf.SetFont("Arial", "B", 8)
	shippingHeaders3 := []string{"Port of Loading", "Date of Departure", "Final Destination", ""}
	xPos = leftMargin
	for _, header := range shippingHeaders3 {
		pdf.SetXY(xPos, shipDetailY)
		pdf.CellFormat(colWidth, rowHeight, header, "1", 0, "L", true, 0, "")
		xPos += colWidth
	}

	shipDetailY += rowHeight
	pdf.SetFont("Arial", "", 8)
	shippingData3 := []string{data.PortOfLoading, data.DateOfDeparture, data.FinalDestination, ""}
	xPos = leftMargin
	for _, d := range shippingData3 {
		pdf.SetXY(xPos, shipDetailY)
		pdf.CellFormat(colWidth, rowHeight, d, "1", 0, "L", false, 0, "")
		xPos += colWidth
	}

	return shipDetailY + rowHeight
}

func renderBody(pdf *gofpdf.Fpdf, data PackingListData, colors Colors, startY float64) float64 {
	leftMargin, _, _, _ := pdf.GetMargins()

	currentY := startY
	tableRowHeight := 10.0

	// Define columns for the product table
	// Total width: 25 + 60 + 22 + 22 + 22 + 22 + 22 = 195 mm
	columns := []table.Column{
		{Key: "product_code", Label: "Product Code", Width: 25, Align: "L"},
		{Key: "description", Label: "Description of Goods", Width: 60, Align: "L"},
		{Key: "unit_quantity", Label: "Unit\nQuantity", Width: 22, Align: "C"},
		{Key: "kind_no_pkg", Label: "Kind & No of\nPackages", Width: 22, Align: "C"},
		{Key: "net_weight", Label: "Net Weight\n(Kg)", Width: 22, Align: "R"},
		{Key: "gross_weight", Label: "Gross Weight\n(Kg)", Width: 22, Align: "R"},
		{Key: "measurements", Label: "Measurements\n(m3)", Width: 22, Align: "R"},
	}

	// Get page height to calculate available space for footer
	_, pageHeight := pdf.GetPageSize()
	_, _, _, bottomMargin := pdf.GetMargins()
	footerHeight := 30.0 // Space needed for footer

	// Calculate margin to ensure footer fits on page
	pageBreakMargin := footerHeight + 5 // Extra margin to ensure footer fits

	// Create table with styling and proper page break handling
	productTable := table.NewTable(pdf, columns).
		WithStartPosition(leftMargin, currentY).
		WithRowHeight(tableRowHeight).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FontSize:  6.5,
			FillColor: colors.LightGray,
			TextColor: colors.TextDark,
		}).
		WithDataStyle(table.CellStyle{
			Border:   "1",
			FontSize: 7,
		}).
		WithPageBreakMode(true).             // Enable automatic page breaks
		WithRepeatHeader(false).             // Disable auto header (we'll render manually for multi-line support)
		WithPageBreakMargin(pageBreakMargin) // Set margin to leave space for footer

	// Manually render header with multi-line support
	headerRowHeight := 10.0 // Height for multi-line headers
	xPos := leftMargin

	// Set header style
	pdf.SetFillColor(colors.LightGray[0], colors.LightGray[1], colors.LightGray[2])
	pdf.SetTextColor(colors.TextDark[0], colors.TextDark[1], colors.TextDark[2])
	pdf.SetFont("Arial", "B", 6.5)

	// Helper function to get alignment string
	getAlignStr := func(align string) string {
		switch align {
		case "C", "Center":
			return "C"
		case "R", "Right":
			return "R"
		default:
			return "L"
		}
	}

	for _, col := range columns {
		// Draw cell border and background
		pdf.Rect(xPos, currentY, col.Width, headerRowHeight, "D")
		pdf.Rect(xPos, currentY, col.Width, headerRowHeight, "F")

		// Save current position before MultiCell (which changes X position)
		savedX := xPos

		// Use MultiCell for multi-line headers with proper alignment
		// For centered text, set padding to ensure proper centering
		padding := 0.3
		pdf.SetXY(xPos+padding, currentY+padding)
		pdf.MultiCell(col.Width-2*padding, 3.5, col.Label, "", getAlignStr(col.Align), false)

		// Reset X position for next cell (MultiCell moves X to left margin)
		xPos = savedX + col.Width
		pdf.SetX(xPos)
	}

	// Move to next line after header
	pdf.SetXY(leftMargin, currentY+headerRowHeight)
	currentY = pdf.GetY()

	// Prepare and add data rows
	for _, product := range data.Products {
		rowData := map[string]interface{}{
			"product_code":  product.ProductCode,
			"description":   product.Description,
			"unit_quantity": product.UnitQuantity,
			"kind_no_pkg":   product.KindNoOfPkg,
			"net_weight":    product.NetWeight,
			"gross_weight":  product.GrossWeight,
			"measurements":  product.Measurements,
		}

		productTable.AddRow(rowData)
		currentY = pdf.GetY()
	}

	// Add empty rows for spacing (reduce if needed to fit on page)
	// Check available space before adding empty rows
	pageHeight, _ = pdf.GetPageSize()
	_, _, _, bottomMargin = pdf.GetMargins()
	footerHeight = 30.0
	remainingSpace := pageHeight - bottomMargin - pdf.GetY() - footerHeight
	rowHeight := tableRowHeight

	// Calculate how many empty rows can fit
	maxEmptyRows := int((remainingSpace - (tableRowHeight * 3)) / rowHeight) // Reserve space for 2 total rows
	emptyRows := 5
	if maxEmptyRows < emptyRows {
		emptyRows = maxEmptyRows
		if emptyRows < 0 {
			emptyRows = 0
		}
	}

	for i := 0; i < emptyRows; i++ {
		emptyRow := map[string]interface{}{
			"product_code":  "",
			"description":   "",
			"unit_quantity": "",
			"kind_no_pkg":   "",
			"net_weight":    "",
			"gross_weight":  "",
			"measurements":  "",
		}
		productTable.AddRow(emptyRow)
		currentY = pdf.GetY()
	}

	// Calculate totals from products (simplified - you may want to enhance this)
	// For demo, use the static totals
	totals := map[string]interface{}{
		"unit_quantity": "225",
		"kind_no_pkg":   "16",
		"net_weight":    "3,000",
		"gross_weight":  "3,225",
		"measurements":  "27",
	}

	// Add "Total This Page" row (label spans first 2 columns: Product Code + Description)
	productTable.AddSummaryRow("Total This Page", 2, totals, table.CellStyle{
		Border:    "1",
		Bold:      true,
		FontSize:  7,
		FillColor: []int{255, 255, 255}, // White background for total row
	})
	currentY = pdf.GetY()

	// Check if there's space for the last total row before adding
	pageHeight, _ = pdf.GetPageSize()
	_, _, _, bottomMargin = pdf.GetMargins()
	footerHeight = 30.0
	availableSpace := pageHeight - bottomMargin - pdf.GetY() - footerHeight

	// Only add consignment total if there's enough space
	if availableSpace >= tableRowHeight {
		// Add "Consignment Total" row (label spans first 2 columns)
		productTable.AddSummaryRow("Consignment Total", 2, totals, table.CellStyle{
			Border:    "1",
			Bold:      true,
			FontSize:  7,
			FillColor: []int{255, 255, 255}, // White background for total row
		})
		currentY = pdf.GetY()
	}

	return pdf.GetY() - startY
}

func renderFooter(pdf *gofpdf.Fpdf, data PackingListData, colors Colors, startY float64) float64 {
	pageWidth, _ := pdf.GetPageSize()
	leftMargin, _, rightMargin, _ := pdf.GetMargins()

	currentY := startY

	// ===== ADDITIONAL INFO SECTION =====
	additionalBoxHeight := 28.0
	pdf.SetDrawColor(colors.BorderGray[0], colors.BorderGray[1], colors.BorderGray[2])
	pdf.SetLineWidth(0.5)
	pdf.Rect(leftMargin, currentY, pageWidth-leftMargin-rightMargin, additionalBoxHeight, "D")

	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(leftMargin+2, currentY+2)
	pdf.Cell(50, 4, "Additional Info")

	// Three columns layout
	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(leftMargin+2, currentY+8)
	pdf.Cell(50, 4, "Signatory Company")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+2, currentY+12)
	pdf.Cell(50, 4, data.SignatoryCompany)

	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(leftMargin+65, currentY+8)
	pdf.Cell(60, 4, "Name of Authorized Signatory")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(leftMargin+65, currentY+12)
	pdf.Cell(30, 4, data.AuthorizedName)

	pdf.SetXY(leftMargin+130, currentY+12)
	pdf.Cell(30, 4, data.AuthorizedSurname)

	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(leftMargin+2, currentY+18)
	pdf.Cell(50, 4, "Signature")

	// Signature (handwritten style)
	pdf.SetFont("Times", "BI", 14)
	pdf.SetXY(leftMargin+100, currentY+20)
	pdf.Cell(60, 5, fmt.Sprintf("%s %s", data.AuthorizedName, data.AuthorizedSurname))

	return currentY + additionalBoxHeight
}

/*
=================================
MULTI-PAGE PACKING LIST EXAMPLE
=================================

To create a multi-page packing list with repeating headers and footers,
you can split products across multiple pages like this:

func generateMultiPagePackingList(pdf *gofpdf.Fpdf, data PackingListData) {
	colors := Colors{
		PrimaryRed: []int{230, 0, 0},
		LightGray:  []int{240, 240, 240},
		White:      []int{255, 255, 255},
		TextDark:   []int{0, 0, 0},
		BorderGray: []int{0, 0, 0},
	}

	// Configuration
	productsPerPage := 4  // Number of products to show per page
	totalProducts := len(data.Products)
	totalPages := (totalProducts + productsPerPage - 1) / productsPerPage

	// Update total pages
	data.TotalPages = totalPages

	// Generate each page
	for pageNum := 1; pageNum <= totalPages; pageNum++ {
		// Add a new page
		if pageNum > 1 {
			pdf.AddPage()
		}

		// Update page number
		data.PageNumber = pageNum

		// Get products for this page
		startIdx := (pageNum - 1) * productsPerPage
		endIdx := startIdx + productsPerPage
		if endIdx > totalProducts {
			endIdx = totalProducts
		}

		// Create page-specific data with subset of products
		pageData := data
		pageData.Products = data.Products[startIdx:endIdx]

		// Render header (repeats on every page)
		headerHeight := renderHeader(pdf, pageData, colors)

		// Render body (only products for this page)
		bodyStartY := headerHeight + 2
		bodyHeight := renderBody(pdf, pageData, colors, bodyStartY)

		// Render footer (repeats on every page)
		footerY := bodyStartY + bodyHeight + 2
		renderFooter(pdf, pageData, colors, footerY)
	}
}

Usage:
	data := PackingListData{
		// ... your data with many products ...
		Products: []ProductItem{
			{...}, // Product 1
			{...}, // Product 2
			{...}, // Product 3
			{...}, // Product 4
			{...}, // Product 5
			// ... more products ...
		},
	}

	generateMultiPagePackingList(pdf, data)

This will automatically:
- Calculate the number of pages needed
- Repeat the header on each page with correct page numbers
- Distribute products across pages
- Repeat the footer on each page
*/
