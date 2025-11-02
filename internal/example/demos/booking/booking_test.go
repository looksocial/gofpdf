package main_test

import (
	"bytes"
	"testing"

	"github.com/looksocial/gofpdf"
)

func TestBookingDemoGeneration(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetAutoPageBreak(false, 0)

	// Colors
	darkBlue := []int{0, 51, 102}
	black := []int{0, 0, 0}

	// Page setup
	leftMargin := 10.0
	currentY := 10.0

	// Header Section
	pdf.SetFillColor(darkBlue[0], darkBlue[1], darkBlue[2])
	pdf.Rect(leftMargin, currentY, 50, 20, "F")

	pdf.SetFont("Arial", "B", 24)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetXY(leftMargin+5, currentY+6)
	pdf.Cell(40, 8, "gofpdf")

	currentY += 22

	// Generate a simple booking document structure
	pdf.SetFont("Arial", "B", 14)
	pdf.SetTextColor(black[0], black[1], black[2])
	pdf.SetXY(leftMargin, currentY)
	pdf.Cell(100, 6, "Booking Acknowledgement")

	currentY += 10

	pdf.SetFont("Arial", "", 9)
	pdf.SetXY(leftMargin, currentY)
	pdf.Cell(100, 4, "This is a test booking document")

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}

	if buf.Len() == 0 {
		t.Error("Generated PDF should have content")
	}
}

func TestBookingHelperFunctions(t *testing.T) {
	// Test renderSectionHeader function
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	renderSectionHeader(pdf, "TEST SECTION", 10, 20, 180, []int{240, 240, 240})

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestBookingInfoRow(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	renderInfoRow(pdf, "TEST LABEL:", "TEST VALUE", 10, 20)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestBookingMultipleSections(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetAutoPageBreak(false, 0)

	lightGray := []int{240, 240, 240}

	leftMargin := 10.0
	currentY := 10.0

	// Render multiple sections
	renderSectionHeader(pdf, "SECTION 1", leftMargin, currentY, 180, lightGray)
	currentY += 15
	renderInfoRow(pdf, "Label 1:", "Value 1", leftMargin, currentY)
	currentY += 10
	renderInfoRow(pdf, "Label 2:", "Value 2", leftMargin, currentY)
	currentY += 15

	renderSectionHeader(pdf, "SECTION 2", leftMargin, currentY, 180, lightGray)
	currentY += 15
	renderInfoRow(pdf, "Label 3:", "Value 3", leftMargin, currentY)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

// Helper function for testing (extracted from main.go)
func renderSectionHeader(pdf *gofpdf.Fpdf, title string, x, y, width float64, bgColor []int) {
	pdf.SetFillColor(bgColor[0], bgColor[1], bgColor[2])
	pdf.Rect(x, y, width, 6, "F")
	pdf.SetFont("Arial", "B", 9)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetXY(x+2, y+1)
	pdf.Cell(width-4, 4, title)
}

func renderInfoRow(pdf *gofpdf.Fpdf, label, value string, x, y float64) {
	pdf.SetFont("Arial", "B", 8)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetXY(x+2, y)
	pdf.Cell(60, 4, label)
	pdf.SetFont("Arial", "", 8)
	pdf.SetXY(x+55, y)
	pdf.Cell(120, 4, value)
}
