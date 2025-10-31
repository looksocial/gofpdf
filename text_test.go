package gofpdf_test

import (
	"testing"

	"github.com/looksocial/gofpdf"
)

// Test text operations
func TestText(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	pdf.Text(10, 10, "Hello World")
	
	if pdf.Error() != nil {
		t.Errorf("Text failed: %v", pdf.Error())
	}
}

func TestWrite(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	pdf.Write(10, "This is a test")
	
	if pdf.Error() != nil {
		t.Errorf("Write failed: %v", pdf.Error())
	}
}

func TestWritef(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	pdf.Writef(10, "Number: %d, String: %s", 42, "test")
	
	if pdf.Error() != nil {
		t.Errorf("Writef failed: %v", pdf.Error())
	}
}

func TestWriteAligned(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	pdf.WriteAligned(100, 10, "Left aligned", "L")
	pdf.WriteAligned(100, 10, "Center aligned", "C")
	pdf.WriteAligned(100, 10, "Right aligned", "R")
	pdf.WriteAligned(100, 10, "Default", "")
	
	if pdf.Error() != nil {
		t.Errorf("WriteAligned failed: %v", pdf.Error())
	}
}

func TestCell(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	pdf.Cell(40, 10, "Simple cell")
	
	if pdf.Error() != nil {
		t.Errorf("Cell failed: %v", pdf.Error())
	}
}

func TestCellf(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	pdf.Cellf(40, 10, "Value: %d", 42)
	
	if pdf.Error() != nil {
		t.Errorf("Cellf failed: %v", pdf.Error())
	}
}

func TestCellFormat(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	// Test various alignments
	pdf.CellFormat(40, 10, "Left", "1", 0, "L", false, 0, "")
	pdf.CellFormat(40, 10, "Center", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, "Right", "1", 1, "R", false, 0, "")
	
	// Test with fill
	pdf.SetFillColor(220, 220, 220)
	pdf.CellFormat(40, 10, "Filled", "1", 1, "L", true, 0, "")
	
	if pdf.Error() != nil {
		t.Errorf("CellFormat failed: %v", pdf.Error())
	}
}

func TestMultiCell(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	longText := "This is a long text that should be wrapped across multiple lines when rendered in a MultiCell."
	
	pdf.MultiCell(80, 5, longText, "", "L", false)
	pdf.MultiCell(80, 5, longText, "1", "C", false)
	pdf.MultiCell(80, 5, longText, "", "R", false)
	pdf.MultiCell(80, 5, longText, "", "J", false)
	
	if pdf.Error() != nil {
		t.Errorf("MultiCell failed: %v", pdf.Error())
	}
}

func TestSetFont(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	// Test core fonts
	fonts := []string{"Arial", "Times", "Courier", "Helvetica"}
	styles := []string{"", "B", "I", "BI"}
	
	for _, font := range fonts {
		for _, style := range styles {
			pdf.SetFont(font, style, 12)
			if pdf.Error() != nil {
				t.Errorf("SetFont(%s, %s, 12) failed: %v", font, style, pdf.Error())
			}
		}
	}
}

func TestSetFontSize(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	pdf.SetFontSize(16)
	_, size := pdf.GetFontSize()
	
	if size <= 0 {
		t.Error("Font size should be positive after SetFontSize")
	}
}

func TestSetFontUnitSize(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	pdf.SetFontUnitSize(5.0)
	_, size := pdf.GetFontSize()
	
	if size != 5.0 {
		t.Errorf("SetFontUnitSize(5.0) resulted in size %f", size)
	}
}

func TestSetTextRenderingMode(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	// Test modes 0-3
	for mode := 0; mode <= 3; mode++ {
		pdf.SetTextRenderingMode(mode)
		pdf.Text(10, 10+float64(mode*10), "Mode test")
	}
	
	if pdf.Error() != nil {
		t.Errorf("SetTextRenderingMode failed: %v", pdf.Error())
	}
}

func TestSetUnderlineThickness(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "U", 12)
	
	pdf.SetUnderlineThickness(0.5)
	pdf.Cell(40, 10, "Thin underline")
	
	pdf.SetUnderlineThickness(2.0)
	pdf.Cell(40, 10, "Thick underline")
	
	if pdf.Error() != nil {
		t.Errorf("SetUnderlineThickness failed: %v", pdf.Error())
	}
}

func TestSetWordSpacing(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	pdf.SetWordSpacing(2.0)
	pdf.Write(10, "Words with spacing")
	
	pdf.SetWordSpacing(0)
	
	if pdf.Error() != nil {
		t.Errorf("SetWordSpacing failed: %v", pdf.Error())
	}
}

func TestSetAutoPageBreak(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	pdf.SetAutoPageBreak(true, 15)
	pdf.SetAutoPageBreak(false, 0)
	
	if pdf.Error() != nil {
		t.Errorf("SetAutoPageBreak failed: %v", pdf.Error())
	}
}

func TestSetDisplayMode(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	pdf.SetDisplayMode("fullpage", "single")
	pdf.SetDisplayMode("fullwidth", "continuous")
	pdf.SetDisplayMode("real", "two")
	pdf.SetDisplayMode("default", "default")
	
	if pdf.Error() != nil {
		t.Errorf("SetDisplayMode failed: %v", pdf.Error())
	}
}

func TestSetCompression(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetCompression(true)
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Compressed")
	
	if pdf.Error() != nil {
		t.Errorf("SetCompression failed: %v", pdf.Error())
	}
}

func TestMetadata(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	
	pdf.SetTitle("Test Title", false)
	pdf.SetSubject("Test Subject", false)
	pdf.SetAuthor("Test Author", false)
	pdf.SetKeywords("test, keywords", false)
	pdf.SetCreator("Test Creator", false)
	
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Metadata test")
	
	if pdf.Error() != nil {
		t.Errorf("Metadata operations failed: %v", pdf.Error())
	}
}

func TestMetadataUTF8(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	
	pdf.SetTitle("テスト", true)
	pdf.SetSubject("Тест", true)
	pdf.SetAuthor("Tëst", true)
	
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "UTF-8 metadata")
	
	if pdf.Error() != nil {
		t.Errorf("UTF-8 metadata failed: %v", pdf.Error())
	}
}

func TestImageTypeFromMime(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	
	tests := []struct {
		mime     string
		expected string
	}{
		{"image/jpeg", "jpg"},
		{"image/jpg", "jpg"},
		{"image/png", "png"},
		{"image/gif", "gif"},
	}
	
	for _, tt := range tests {
		result := pdf.ImageTypeFromMime(tt.mime)
		if result != tt.expected {
			t.Errorf("ImageTypeFromMime(%s) = %s; want %s", tt.mime, result, tt.expected)
		}
	}
}

func TestOk(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	
	if !pdf.Ok() {
		t.Error("New PDF should be Ok")
	}
	
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Test")
	
	if !pdf.Ok() {
		t.Error("PDF should still be Ok after basic operations")
	}
}

func TestError(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	
	if pdf.Error() != nil {
		t.Error("New PDF should have no error")
	}
	
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	if pdf.Error() != nil {
		t.Error("PDF should have no error after basic operations")
	}
}

func TestString(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	
	str := pdf.String()
	if len(str) == 0 {
		t.Error("String() should return non-empty string")
	}
}

