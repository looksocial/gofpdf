package main

import (
	"github.com/looksocial/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Test 1: Set Arial and write text
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(100, 10, "Test 1: With Arial set")
	pdf.Ln(15)

	//Test 2: Set font with empty string (mimicking applyCellStyle)
	pdf.SetFont("", "", 10)
	pdf.Cell(100, 10, "Test 2: After empty font")
	pdf.Ln(15)

	// Test 3: Back to Arial
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(100, 10, "Test 3: Arial again")

	err := pdf.OutputFileAndClose("font_test.pdf")
	if err != nil {
		panic(err)
	}
	println("Test PDF created")
}
