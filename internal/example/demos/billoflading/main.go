//go:build ignore
// +build ignore

/*
Bill of Lading Demo

A professional Bill of Lading document generator using gofpdf.
This demo showcases:
- Code128 barcode generation
- Table component for structured data
- Modular component architecture
- Professional PDF layout

Usage:
  go run main.go

The demo uses a modular component structure where all rendering functions
are imported from the components package.
*/
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/looksocial/gofpdf"
	"github.com/looksocial/gofpdf/internal/example/demos/billoflading/components"
)

func main() {
	// Create new PDF
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Add a page
	pdf.AddPage()

	// Generate the Bill of Lading
	generateBillOfLading(pdf)

	// Ensure pdf directory exists
	pdfDir := filepath.Join("..", "..", "..", "..", "pdf")
	err := os.MkdirAll(pdfDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Save to file in pdf directory
	outputPath := filepath.Join(pdfDir, "bill_of_lading.pdf")
	err = pdf.OutputFileAndClose(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	// Get absolute path for display
	absPath, _ := filepath.Abs(outputPath)
	fmt.Printf("Bill of Lading created successfully: %s\n", absPath)
}

func generateBillOfLading(pdf *gofpdf.Fpdf) {
	// Colors
	black := []int{0, 0, 0}
	white := []int{255, 255, 255}
	lightGray := []int{230, 230, 230}

	// Page dimensions
	pageWidth, _ := pdf.GetPageSize()
	leftMargin := 10.0
	topMargin := 10.0
	rightMargin := 10.0
	contentWidth := pageWidth - leftMargin - rightMargin

	// Render header section
	currentY := components.RenderHeader(pdf, leftMargin, topMargin, contentWidth)

	// Render shipping sections
	currentY = components.RenderShipFrom(pdf, leftMargin, currentY, contentWidth, black, white)
	currentY = components.RenderShipTo(pdf, leftMargin, currentY, contentWidth, black, white)
	currentY = components.RenderThirdParty(pdf, leftMargin, currentY, contentWidth, black, white)
	currentY = components.RenderSpecialInstructions(pdf, leftMargin, currentY, contentWidth)

	// Render tables
	currentY = components.RenderCustomerOrderInformation(pdf, leftMargin, currentY, contentWidth, black, white, lightGray)
	currentY = components.RenderCarrierInformation(pdf, leftMargin, currentY, contentWidth, black, white, lightGray)

	// Render footer
	components.RenderFooter(pdf, leftMargin, currentY)
}
