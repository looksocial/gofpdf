package components_test

import (
	"bytes"
	"testing"

	"github.com/looksocial/gofpdf"
	"github.com/looksocial/gofpdf/internal/example/demos/billoflading/components"
)

func TestRenderHeader(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	leftMargin := 10.0
	topMargin := 10.0
	contentWidth := 190.0
	
	currentY := components.RenderHeader(pdf, leftMargin, topMargin, contentWidth)
	
	if currentY < topMargin {
		t.Errorf("RenderHeader should return Y position >= topMargin, got %f", currentY)
	}
	
	// Verify PDF was created successfully
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestRenderShipFrom(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	black := []int{0, 0, 0}
	white := []int{255, 255, 255}
	leftMargin := 10.0
	currentY := 10.0
	contentWidth := 190.0
	
	newY := components.RenderShipFrom(pdf, leftMargin, currentY, contentWidth, black, white)
	
	if newY <= currentY {
		t.Errorf("RenderShipFrom should return Y position > currentY, got %f", newY)
	}
	
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestRenderShipTo(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	black := []int{0, 0, 0}
	white := []int{255, 255, 255}
	leftMargin := 10.0
	currentY := 50.0
	contentWidth := 190.0
	
	newY := components.RenderShipTo(pdf, leftMargin, currentY, contentWidth, black, white)
	
	if newY <= currentY {
		t.Errorf("RenderShipTo should return Y position > currentY, got %f", newY)
	}
	
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestRenderThirdParty(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	black := []int{0, 0, 0}
	white := []int{255, 255, 255}
	leftMargin := 10.0
	currentY := 90.0
	contentWidth := 190.0
	
	newY := components.RenderThirdParty(pdf, leftMargin, currentY, contentWidth, black, white)
	
	if newY <= currentY {
		t.Errorf("RenderThirdParty should return Y position > currentY, got %f", newY)
	}
	
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestRenderSpecialInstructions(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	leftMargin := 10.0
	currentY := 120.0
	contentWidth := 190.0
	
	newY := components.RenderSpecialInstructions(pdf, leftMargin, currentY, contentWidth)
	
	if newY <= currentY {
		t.Errorf("RenderSpecialInstructions should return Y position > currentY, got %f", newY)
	}
	
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestRenderCustomerOrderInformation(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	black := []int{0, 0, 0}
	white := []int{255, 255, 255}
	lightGray := []int{230, 230, 230}
	leftMargin := 10.0
	currentY := 160.0
	contentWidth := 190.0
	
	newY := components.RenderCustomerOrderInformation(pdf, leftMargin, currentY, contentWidth, black, white, lightGray)
	
	if newY <= currentY {
		t.Errorf("RenderCustomerOrderInformation should return Y position > currentY, got %f", newY)
	}
	
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestRenderCarrierInformation(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	black := []int{0, 0, 0}
	white := []int{255, 255, 255}
	lightGray := []int{230, 230, 230}
	leftMargin := 10.0
	currentY := 100.0  // Use a reasonable starting position that allows content to fit
	contentWidth := 190.0
	
	newY := components.RenderCarrierInformation(pdf, leftMargin, currentY, contentWidth, black, white, lightGray)
	
	// The function returns pdf.GetY() which could be on the same page or a new page after page break
	// We just verify it returns a valid Y position (>= 0)
	if newY < 0 {
		t.Errorf("RenderCarrierInformation should return Y position >= 0, got %f", newY)
	}
	
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
	
	if buf.Len() == 0 {
		t.Error("Generated PDF should have content")
	}
}

func TestRenderFooter(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	leftMargin := 10.0
	currentY := 200.0
	
	// RenderFooter doesn't return a value, so we just verify it doesn't panic
	components.RenderFooter(pdf, leftMargin, currentY)
	
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}

func TestRenderAllComponentsIntegration(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	black := []int{0, 0, 0}
	white := []int{255, 255, 255}
	lightGray := []int{230, 230, 230}
	
	pageWidth, _ := pdf.GetPageSize()
	leftMargin := 10.0
	topMargin := 10.0
	rightMargin := 10.0
	contentWidth := pageWidth - leftMargin - rightMargin
	
	currentY := components.RenderHeader(pdf, leftMargin, topMargin, contentWidth)
	currentY = components.RenderShipFrom(pdf, leftMargin, currentY, contentWidth, black, white)
	currentY = components.RenderShipTo(pdf, leftMargin, currentY, contentWidth, black, white)
	currentY = components.RenderThirdParty(pdf, leftMargin, currentY, contentWidth, black, white)
	currentY = components.RenderSpecialInstructions(pdf, leftMargin, currentY, contentWidth)
	currentY = components.RenderCustomerOrderInformation(pdf, leftMargin, currentY, contentWidth, black, white, lightGray)
	currentY = components.RenderCarrierInformation(pdf, leftMargin, currentY, contentWidth, black, white, lightGray)
	components.RenderFooter(pdf, leftMargin, currentY)
	
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
	
	if buf.Len() == 0 {
		t.Error("Generated PDF should have content")
	}
}
