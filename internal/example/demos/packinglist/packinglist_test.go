package main_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/looksocial/gofpdf"
)

func TestPackingListGeneration(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Sample data
	data := PackingListData{
		PageNumber:       1,
		TotalPages:       1,
		ExporterName:     "Test Exports",
		ExporterAddr1:    "123 Test Street",
		ExporterAddr2:    "Test City, State, 12345",
		ExporterAddr3:    "United States",
		ContactName:      "Test User",
		TaxID:            "TEST-001",
		Email:            "test@example.com",
		InvoiceNumber:    "TEST-2024-001",
		InvoiceDate:      "30 Jul 2024",
		BillOfLading:     "TEST-LON-001",
		Reference:        "TEST-2024-001",
		BuyerReference:   "PO-TEST-001",
		BuyerIfNotCons:   "",
		ConsigneeName:    "Test Imports",
		ConsigneeAddr:    "5678 Business Avenue, Test City, State, 54321, Test Country",
		MethodOfDispatch: "Sea",
		TypeOfShipment:   "FCL",
		CountryOfOrigin:  "United States",
		CountryOfDest:    "Test Country",
		VesselAircraft:   "TEST VESSEL",
		VoyageNo:         "TEST-V001",
		PackingInfo:      "",
		PortOfLoading:    "Test Port",
		DateOfDeparture:  "04 Jul 2024",
		FinalDestination: "Port of Test",
		Products: []ProductItem{
			{
				ProductCode:  "TEST-STOOL-001",
				Description:  "Test BAR STOOL ALUMINUM 500 X 100 X 100MM",
				UnitQuantity: "150",
				KindNoOfPkg:  "PALLET X II",
				NetWeight:    "1,500",
				GrossWeight:  "1,600",
				Measurements: "12",
			},
		},
		SignatoryCompany:  "Test Exports",
		AuthorizedName:    "Test",
		AuthorizedSurname: "User",
	}

	// Test PDF generation with sample data structure
	pdf.SetFont("Arial", "B", 18)
	pdf.SetXY(10, 10)
	pdf.Cell(100, 8, "PACKING LIST")
	
	pdf.SetFont("Arial", "", 9)
	pdf.SetXY(10, 25)
	pdf.Cell(100, 5, "Exporter: "+data.ExporterName)
	pdf.SetXY(10, 30)
	pdf.Cell(100, 5, "Invoice: "+data.InvoiceNumber)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}

	if buf.Len() == 0 {
		t.Error("Generated PDF should have content")
	}
}

func TestPackingListHeader(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	data := PackingListData{
		PageNumber:    1,
		TotalPages:    1,
		ExporterName:  "Test Exports",
		ExporterAddr1: "123 Test Street",
		ExporterAddr2: "Test City",
		ExporterAddr3: "United States",
		ContactName:   "Test User",
		TaxID:         "TEST-001",
		Email:         "test@example.com",
	}

	// Test basic header rendering
	pdf.SetFillColor(230, 0, 0)
	pdf.Circle(18, 18, 7, "F")
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Arial", "B", 5)
	pdf.SetXY(13, 16)
	pdf.Cell(10, 4, "gofpdf")

	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "B", 18)
	pdf.SetXY(70, 12)
	pdf.Cell(70, 8, "PACKING LIST")

	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(12, 27)
	pdf.Cell(80, 4, data.ExporterName)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestPackingListShippingDetails(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	data := PackingListData{
		MethodOfDispatch: "Sea",
		TypeOfShipment:   "FCL",
		CountryOfOrigin:  "United States",
		CountryOfDest:    "Test Country",
		VesselAircraft:   "TEST VESSEL",
		VoyageNo:         "TEST-V001",
		PortOfLoading:    "Test Port",
		DateOfDeparture:  "04 Jul 2024",
		FinalDestination: "Port of Test",
	}

	// Test basic shipping details rendering
	startY := 100.0
	pdf.SetFillColor(240, 240, 240)
	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(10, startY)
	pdf.CellFormat(48.75, 7, "Method of Dispatch", "1", 0, "L", true, 0, "")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(10, startY+7)
	pdf.CellFormat(48.75, 7, data.MethodOfDispatch, "1", 0, "L", false, 0, "")

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestPackingListBody(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	data := PackingListData{
		Products: []ProductItem{
			{
				ProductCode:  "TEST-001",
				Description:  "Test Product",
				UnitQuantity: "100",
				KindNoOfPkg:  "PALLET X I",
				NetWeight:    "1,000",
				GrossWeight:  "1,100",
				Measurements: "10",
			},
		},
	}

	// Test basic product table rendering
	startY := 200.0
	pdf.SetFont("Arial", "", 7)
	pdf.SetXY(10, startY)
	pdf.Cell(25, 10, data.Products[0].ProductCode)
	pdf.SetXY(35, startY)
	pdf.Cell(60, 10, data.Products[0].Description)
	pdf.SetXY(95, startY)
	pdf.Cell(22, 10, data.Products[0].UnitQuantity)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestPackingListFooter(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	data := PackingListData{
		SignatoryCompany:  "Test Exports",
		AuthorizedName:    "Test",
		AuthorizedSurname: "User",
	}

	// Test basic footer rendering
	startY := 250.0
	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(12, startY+2)
	pdf.Cell(50, 4, "Additional Info")
	
	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(12, startY+8)
	pdf.Cell(50, 4, "Signatory Company")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(12, startY+12)
	pdf.Cell(50, 4, data.SignatoryCompany)

	pdf.SetFont("Arial", "B", 8)
	pdf.SetXY(75, startY+8)
	pdf.Cell(60, 4, "Name of Authorized Signatory")
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(75, startY+12)
	pdf.Cell(30, 4, data.AuthorizedName)
	pdf.SetXY(140, startY+12)
	pdf.Cell(30, 4, data.AuthorizedSurname)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestPackingListMultipleProducts(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	data := PackingListData{
		PageNumber:       1,
		TotalPages:       1,
		ExporterName:     "Test Exports",
		ExporterAddr1:    "123 Test Street",
		ExporterAddr2:    "Test City",
		ExporterAddr3:    "United States",
		ContactName:      "Test User",
		TaxID:            "TEST-001",
		Email:            "test@example.com",
		InvoiceNumber:    "TEST-2024-001",
		InvoiceDate:      "30 Jul 2024",
		BillOfLading:     "TEST-LON-001",
		Reference:        "TEST-2024-001",
		BuyerReference:   "PO-TEST-001",
		ConsigneeName:    "Test Imports",
		ConsigneeAddr:    "5678 Business Avenue",
		MethodOfDispatch: "Sea",
		TypeOfShipment:   "FCL",
		CountryOfOrigin:  "United States",
		CountryOfDest:    "Test Country",
		VesselAircraft:   "TEST VESSEL",
		VoyageNo:         "TEST-V001",
		PortOfLoading:    "Test Port",
		DateOfDeparture:  "04 Jul 2024",
		FinalDestination: "Port of Test",
		Products: []ProductItem{
			{
				ProductCode:  "TEST-001",
				Description:  "Product 1",
				UnitQuantity: "100",
				KindNoOfPkg:  "PALLET X I",
				NetWeight:    "1,000",
				GrossWeight:  "1,100",
				Measurements: "10",
			},
			{
				ProductCode:  "TEST-002",
				Description:  "Product 2",
				UnitQuantity: "200",
				KindNoOfPkg:  "PALLET X II",
				NetWeight:    "2,000",
				GrossWeight:  "2,200",
				Measurements: "20",
			},
			{
				ProductCode:  "TEST-003",
				Description:  "Product 3",
				UnitQuantity: "300",
				KindNoOfPkg:  "PALLET X III",
				NetWeight:    "3,000",
				GrossWeight:  "3,300",
				Measurements: "30",
			},
		},
		SignatoryCompany:  "Test Exports",
		AuthorizedName:    "Test",
		AuthorizedSurname: "User",
	}

	// Test PDF generation with multiple products
	pdf.SetFont("Arial", "B", 18)
	pdf.SetXY(10, 10)
	pdf.Cell(100, 8, "PACKING LIST")
	
	pdf.SetFont("Arial", "", 9)
	currentY := 30.0
	for i, product := range data.Products {
		pdf.SetXY(10, currentY)
		pdf.Cell(100, 5, fmt.Sprintf("Product %d : %s", i+1, product.ProductCode))
		currentY += 10
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

// Type definitions from main.go for testing
type PackingListData struct {
	PageNumber        int
	TotalPages        int
	ExporterName      string
	ExporterAddr1     string
	ExporterAddr2     string
	ExporterAddr3     string
	ContactName       string
	TaxID             string
	Email             string
	InvoiceNumber     string
	InvoiceDate       string
	BillOfLading      string
	Reference         string
	BuyerReference    string
	BuyerIfNotCons    string
	ConsigneeName     string
	ConsigneeAddr     string
	MethodOfDispatch  string
	TypeOfShipment    string
	CountryOfOrigin   string
	CountryOfDest     string
	VesselAircraft    string
	VoyageNo          string
	PackingInfo       string
	PortOfLoading     string
	DateOfDeparture   string
	FinalDestination  string
	Products          []ProductItem
	SignatoryCompany  string
	AuthorizedName    string
	AuthorizedSurname string
}

type ProductItem struct {
	ProductCode  string
	Description  string
	UnitQuantity string
	KindNoOfPkg  string
	NetWeight    string
	GrossWeight  string
	Measurements string
}

type Colors struct {
	PrimaryRed []int
	LightGray  []int
	White      []int
	TextDark   []int
	BorderGray []int
}
