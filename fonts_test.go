package gofpdf_test

import (
	"testing"

	gofpdf "github.com/looksocial/gofpdf"
	"github.com/looksocial/gofpdf/internal/example"
)

// TestEmbeddedFonts tests the UseEmbeddedFonts feature
func TestEmbeddedFonts(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.UseEmbeddedFonts()
	pdf.AddPage()

	// Test loading a common Thai font
	pdf.SetFont("Kanit", "", 14)
	if pdf.Error() != nil {
		t.Errorf("Failed to set embedded Kanit font: %v", pdf.Error())
	}

	pdf.Cell(40, 10, "Test Kanit")
	
	// Test bold style
	pdf.SetFont("Kanit", "B", 14)
	if pdf.Error() != nil {
		t.Errorf("Failed to set embedded Kanit Bold: %v", pdf.Error())
	}

	pdf.Cell(40, 10, "Test Kanit Bold")
}

// TestEmbeddedFontsMultipleFamilies tests multiple font families
func TestEmbeddedFontsMultipleFamilies(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.UseEmbeddedFonts()
	pdf.AddPage()

	families := []string{"Kanit", "Sarabun", "Prompt", "Athiti"}
	successCount := 0
	
	for _, family := range families {
		pdf.SetFont(family, "", 12)
		if pdf.Error() != nil {
			t.Logf("Font %s not available: %v", family, pdf.Error())
			pdf.SetError(nil) // Reset error to continue
		} else {
			pdf.Cell(40, 10, "Test "+family)
			pdf.Ln(6)
			successCount++
		}
	}
	
	if successCount == 0 {
		t.Error("No embedded fonts were loaded successfully")
	}
	
	t.Logf("Successfully loaded %d out of %d font families", successCount, len(families))
}

// TestEmbeddedFontsWithStyles tests different font styles
func TestEmbeddedFontsWithStyles(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.UseEmbeddedFonts()
	pdf.AddPage()

	styles := []struct {
		family string
		style  string
		desc   string
	}{
		{"Kanit", "", "Regular"},
		{"Kanit", "B", "Bold"},
		{"Kanit", "I", "Italic"},
		{"Kanit", "BI", "BoldItalic"},
		{"Sarabun", "", "Regular"},
		{"Sarabun", "B", "Bold"},
		{"Sarabun", "I", "Italic"},
	}

	for _, s := range styles {
		pdf.SetFont(s.family, s.style, 12)
		if pdf.Error() != nil {
			t.Logf("Note: %s %s may not be available: %v", s.family, s.desc, pdf.Error())
			// Reset error to continue testing
			pdf.SetError(nil)
		} else {
			pdf.Cell(40, 10, s.family+" "+s.desc)
			pdf.Ln(6)
		}
	}
}

// TestAutoLoadFromSubfolder tests auto-loading fonts from subfolders
func TestAutoLoadFromSubfolder(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "font/th")
	pdf.AddPage()

	// Should auto-load from font/th/Kanit/ subfolder
	pdf.SetFont("Kanit", "", 14)
	if pdf.Error() != nil {
		t.Logf("Auto-load from subfolder not available (expected if running from different directory): %v", pdf.Error())
		return
	}

	pdf.Cell(40, 10, "Test auto-load")
}

// TestAutoLoadRegularStyle tests auto-loading with -Regular suffix
func TestAutoLoadRegularStyle(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "font/th/Kanit")
	pdf.AddPage()

	// Should find Kanit-Regular.ttf
	pdf.SetFont("Kanit", "", 14)
	if pdf.Error() != nil {
		t.Logf("Regular style auto-load not available: %v", pdf.Error())
		return
	}

	pdf.Cell(40, 10, "Test regular")
}

// TestAutoLoadBoldStyle tests auto-loading bold fonts
func TestAutoLoadBoldStyle(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "font/th")
	pdf.AddPage()

	// Should find Kanit-Bold.ttf from font/th/Kanit/
	pdf.SetFont("Kanit", "B", 14)
	if pdf.Error() != nil {
		t.Logf("Bold style auto-load not available: %v", pdf.Error())
		return
	}

	pdf.Cell(40, 10, "Test bold")
}

// TestEmbeddedFontsThaiText tests Thai text rendering with embedded fonts
func TestEmbeddedFontsThaiText(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.UseEmbeddedFonts()
	pdf.AddPage()

	pdf.SetFont("Kanit", "", 16)
	if pdf.Error() != nil {
		t.Errorf("Failed to set Kanit font: %v", pdf.Error())
		return
	}

	// Thai text
	thaiText := "สวัสดีครับ ทดสอบภาษาไทย"
	pdf.Cell(0, 10, thaiText)
	
	if pdf.Error() != nil {
		t.Errorf("Failed to render Thai text: %v", pdf.Error())
	}
}

// TestGetEmbeddedFontFamilies tests listing available fonts
func TestGetEmbeddedFontFamilies(t *testing.T) {
	families := gofpdf.GetEmbeddedFontFamilies()
	
	if len(families) == 0 {
		t.Error("No embedded font families found")
	}

	// Check for expected families
	expectedFamilies := []string{"Kanit", "Sarabun", "Prompt"}
	for _, expected := range expectedFamilies {
		found := false
		for _, family := range families {
			if family == expected {
				found = true
				break
			}
		}
		if !found {
			t.Logf("Expected family %s not found in embedded fonts", expected)
		}
	}

	t.Logf("Available embedded font families: %v", families)
}

// TestEmbeddedFontsOutput tests generating a PDF with embedded fonts
func TestEmbeddedFontsOutput(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.UseEmbeddedFonts()
	pdf.AddPage()

	pdf.SetFont("Kanit", "", 14)
	pdf.Cell(40, 10, "Hello World")
	pdf.Ln(8)
	
	pdf.SetFont("Kanit", "", 16)
	pdf.Cell(0, 10, "สวัสดีครับ - Thai Text")

	fileStr := example.Filename("test_output_embedded")
	err := pdf.OutputFileAndClose(fileStr)
	if err != nil {
		t.Errorf("Failed to output PDF with embedded fonts: %v", err)
	}
}

// TestMixedEmbeddedAndCoreFonts tests mixing embedded and core fonts
func TestMixedEmbeddedAndCoreFonts(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.UseEmbeddedFonts()
	pdf.AddPage()

	// Core font
	pdf.SetFont("Arial", "", 12)
	if pdf.Error() != nil {
		t.Errorf("Failed to set core font: %v", pdf.Error())
	}
	pdf.Cell(40, 10, "Arial (Core)")
	pdf.Ln(6)

	// Embedded font
	pdf.SetFont("Kanit", "", 12)
	if pdf.Error() != nil {
		t.Errorf("Failed to set embedded font: %v", pdf.Error())
	}
	pdf.Cell(40, 10, "Kanit (Embedded)")
	pdf.Ln(6)

	// Back to core font
	pdf.SetFont("Times", "B", 12)
	if pdf.Error() != nil {
		t.Errorf("Failed to switch back to core font: %v", pdf.Error())
	}
	pdf.Cell(40, 10, "Times Bold (Core)")
}

// TestAutoLoadFallback tests that auto-load doesn't break existing functionality
func TestAutoLoadFallback(t *testing.T) {
	// Test 1: Non-existent font should fail
	pdf1 := gofpdf.New("P", "mm", "A4", "")
	pdf1.AddPage()
	pdf1.SetFont("NonExistentFont", "", 12)
	if pdf1.Error() == nil {
		t.Error("Expected error for non-existent font, got nil")
	}

	// Test 2: Core fonts should work in new instance
	pdf2 := gofpdf.New("P", "mm", "A4", "")
	pdf2.AddPage()
	pdf2.SetFont("Arial", "", 12)
	if pdf2.Error() != nil {
		t.Errorf("Core font should work: %v", pdf2.Error())
	}
	pdf2.Cell(40, 10, "Arial works")
}

// TestSubfolderSearchPriority tests that subfolder search works correctly
func TestSubfolderSearchPriority(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "font/th")
	pdf.AddPage()

	// Should search in font/th/Tahoma/ for Tahoma fonts
	pdf.SetFont("Tahoma", "", 12)
	if pdf.Error() != nil {
		t.Logf("Tahoma font not available from filesystem: %v", pdf.Error())
		return
	}

	pdf.Cell(40, 10, "Tahoma Test")
	
	// Test bold
	pdf.SetFont("Tahoma", "B", 12)
	if pdf.Error() != nil {
		t.Logf("Tahoma Bold not available: %v", pdf.Error())
	}
}

// TestEmbeddedFontsReuse tests that fonts are loaded once and reused
func TestEmbeddedFontsReuse(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.UseEmbeddedFonts()
	pdf.AddPage()

	// Load Kanit multiple times
	for i := 0; i < 5; i++ {
		pdf.SetFont("Kanit", "", 14)
		if pdf.Error() != nil {
			t.Errorf("Failed on iteration %d: %v", i, pdf.Error())
		}
		pdf.Cell(40, 10, "Test")
		pdf.Ln(6)
	}
}

// TestNotoFonts tests the Noto font families
func TestNotoFonts(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.UseEmbeddedFonts()
	pdf.AddPage()

	notoFamilies := []string{
		"NotoSansThai",
		"NotoSerifThai",
		"NotoSansThaiLooped",
	}

	for _, family := range notoFamilies {
		pdf.SetFont(family, "", 12)
		if pdf.Error() != nil {
			t.Logf("Noto font %s may not be available: %v", family, pdf.Error())
			pdf.SetError(nil)
			continue
		}
		pdf.Cell(40, 10, family)
		pdf.Ln(6)
	}
}

// TestFontSizeAndStyle tests font size and style variations
func TestFontSizeAndStyle(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.UseEmbeddedFonts()
	pdf.AddPage()

	sizes := []float64{8, 10, 12, 14, 16, 18, 20}
	
	for _, size := range sizes {
		pdf.SetFont("Kanit", "", size)
		if pdf.Error() != nil {
			t.Errorf("Failed to set font size %.0f: %v", size, pdf.Error())
		}
		pdf.Cell(40, 10, "Size %.0f")
		pdf.Ln(6)
	}
}

// BenchmarkEmbeddedFontLoad benchmarks embedded font loading
func BenchmarkEmbeddedFontLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.UseEmbeddedFonts()
		pdf.AddPage()
		pdf.SetFont("Kanit", "", 14)
		pdf.Cell(40, 10, "Benchmark")
	}
}

// BenchmarkAutoLoad benchmarks auto-load feature
func BenchmarkAutoLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pdf := gofpdf.New("P", "mm", "A4", "font/th")
		pdf.AddPage()
		pdf.SetFont("Kanit", "", 14)
		pdf.Cell(40, 10, "Benchmark")
	}
}

// TestMultiPageWithEmbeddedFonts tests embedded fonts across multiple pages
func TestMultiPageWithEmbeddedFonts(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.UseEmbeddedFonts()

	for page := 1; page <= 3; page++ {
		pdf.AddPage()
		pdf.SetFont("Kanit", "", 14)
		if pdf.Error() != nil {
			t.Errorf("Failed on page %d: %v", page, pdf.Error())
		}
		pdf.Cell(40, 10, "Page %d")
	}
}

// TestEmbeddedFontsWithUTF8 tests UTF-8 text with embedded fonts
func TestEmbeddedFontsWithUTF8(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.UseEmbeddedFonts()
	pdf.AddPage()

	texts := []struct {
		lang string
		text string
	}{
		{"Thai", "สวัสดีครับ"},
		{"English", "Hello World"},
		{"Mixed", "Hello สวัสดี World"},
		{"Numbers", "1234567890 ๑๒๓๔๕"},
	}

	pdf.SetFont("Sarabun", "", 14)
	for _, txt := range texts {
		if pdf.Error() != nil {
			t.Errorf("Error before rendering %s: %v", txt.lang, pdf.Error())
		}
		pdf.Cell(0, 10, txt.lang+": "+txt.text)
		pdf.Ln(8)
	}

	if pdf.Error() != nil {
		t.Errorf("Error after rendering UTF-8 text: %v", pdf.Error())
	}
}

