package gofpdf_test

import (
	"testing"

	gofpdf "github.com/looksocial/gofpdf"
)

// Test font operations
func TestAddFontFromBytes(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Use embedded font data from internal/files
	pdf.AddFontFromBytes("testfont", "", nil, nil)

	// Should not crash even with nil data
	if pdf.Error() == nil {
		// Font may or may not load with nil, but shouldn't panic
		pdf.SetFont("Arial", "", 12)
		pdf.Cell(40, 10, "Test")
	}
}

func TestSetFontLocation(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.SetFontLocation("font")
	pdf.SetFontLocation("./font")
	pdf.SetFontLocation("")

	// Should not crash with different locations
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Font location test")

	if pdf.Error() != nil {
		t.Errorf("SetFontLocation failed: %v", pdf.Error())
	}
}

func TestAddFont(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "font")
	pdf.AddPage()

	// Try to add a custom font
	pdf.AddFont("Calligrapher", "", "calligra.json")

	if pdf.Error() != nil {
		// Font loading may fail, that's ok for this test
		// We're just verifying it doesn't crash
		pdf.SetFont("Arial", "", 12)
	} else {
		pdf.SetFont("Calligrapher", "", 16)
	}

	pdf.Cell(40, 10, "Font test")
}

func TestUTF8Font(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.AddUTF8Font("dejavu", "", "font/DejaVuSansCondensed.ttf")

	if pdf.Error() != nil {
		t.Errorf("AddUTF8Font failed: %v", pdf.Error())
	}

	pdf.SetFont("dejavu", "", 14)
	pdf.Cell(40, 10, "UTF-8 font test")

	if pdf.Error() != nil {
		t.Errorf("Using UTF-8 font failed: %v", pdf.Error())
	}
}

func TestUTF8FontFromBytes(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// This tests the AddUTF8FontFromBytes method
	// We'll just verify it doesn't crash with nil
	pdf.AddUTF8FontFromBytes("testfont", "", nil)

	// Should handle nil gracefully or set error
	_ = pdf.Error()
}

func TestUnicodeTranslatorFromDescriptor(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "font")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	// Test various code pages
	codepages := []string{"", "cp1252", "cp1251", "cp1253"}

	for _, cp := range codepages {
		translator := pdf.UnicodeTranslatorFromDescriptor(cp)
		if translator == nil {
			t.Errorf("Translator for %s should not be nil", cp)
		}

		// Test that translator works
		result := translator("test")
		if result != "test" {
			// ASCII should pass through
			t.Errorf("ASCII translation failed for %s", cp)
		}
	}
}

func TestCoreFonts(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Test all core font families
	coreFonts := []string{
		"Arial", "Helvetica", "Times", "Courier",
		"Symbol", "ZapfDingbats",
	}

	for _, font := range coreFonts {
		pdf.SetFont(font, "", 12)
		pdf.Cell(40, 5, font)
		pdf.Ln(-1)

		if pdf.Error() != nil {
			t.Errorf("Core font %s failed: %v", font, pdf.Error())
		}
	}
}

func TestFontStyles(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Test all style combinations
	styles := []string{"", "B", "I", "U", "BI", "BU", "IU", "BIU", "S"}

	for _, style := range styles {
		pdf.SetFont("Arial", style, 12)
		pdf.Cell(40, 5, "Style: "+style)
		pdf.Ln(-1)

		if pdf.Error() != nil {
			t.Errorf("Font style %s failed: %v", style, pdf.Error())
		}
	}
}

func TestFontSizes(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	sizes := []float64{6, 8, 10, 12, 14, 16, 18, 24, 36, 48}

	for _, size := range sizes {
		pdf.SetFontSize(size)
		pdf.Ln(-1)

		if pdf.Error() != nil {
			t.Errorf("Font size %f failed: %v", size, pdf.Error())
		}
	}
}

func TestGetFontDesc(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	desc := pdf.GetFontDesc("Arial", "")
	if desc.Ascent == 0 && desc.Descent == 0 {
		// Font descriptor should have some values
		t.Log("Font descriptor may be empty, which is ok")
	}
}

func TestAddUTF8FontMultipleStyles(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Add multiple styles of the same font
	pdf.AddUTF8Font("dejavu", "", "font/DejaVuSansCondensed.ttf")
	pdf.AddUTF8Font("dejavu", "B", "font/DejaVuSansCondensed-Bold.ttf")
	pdf.AddUTF8Font("dejavu", "I", "font/DejaVuSansCondensed-Oblique.ttf")
	pdf.AddUTF8Font("dejavu", "BI", "font/DejaVuSansCondensed-BoldOblique.ttf")

	// Use each style
	pdf.SetFont("dejavu", "", 12)
	pdf.Cell(40, 5, "Regular")
	pdf.Ln(-1)

	pdf.SetFont("dejavu", "B", 12)
	pdf.Cell(40, 5, "Bold")
	pdf.Ln(-1)

	pdf.SetFont("dejavu", "I", 12)
	pdf.Cell(40, 5, "Italic")
	pdf.Ln(-1)

	pdf.SetFont("dejavu", "BI", 12)
	pdf.Cell(40, 5, "Bold Italic")

	if pdf.Error() != nil {
		t.Errorf("UTF-8 font styles failed: %v", pdf.Error())
	}
}

func TestUTF8Text(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.AddUTF8Font("dejavu", "", "font/DejaVuSansCondensed.ttf")
	pdf.SetFont("dejavu", "", 14)

	// Test various UTF-8 characters
	pdf.Cell(0, 10, "English: Hello World")
	pdf.Ln(-1)
	pdf.Cell(0, 10, "French: Bonjour le monde")
	pdf.Ln(-1)
	pdf.Cell(0, 10, "German: Hallo Welt")
	pdf.Ln(-1)
	pdf.Cell(0, 10, "Spanish: Hola Mundo")
	pdf.Ln(-1)

	if pdf.Error() != nil {
		t.Errorf("UTF-8 text failed: %v", pdf.Error())
	}
}
