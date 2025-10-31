package gofpdf_test

import (
	"testing"

	"github.com/looksocial/gofpdf"
)

// Test StateType functionality
func TestStateGetPut(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	// Set various properties
	pdf.SetDrawColor(255, 0, 0)
	pdf.SetFillColor(0, 255, 0)
	pdf.SetTextColor(0, 0, 255)
	pdf.SetLineWidth(2.5)
	pdf.SetFontUnitSize(14)
	pdf.SetAlpha(0.5, "Normal")
	pdf.SetCellMargin(3.0)
	
	// Get state
	state := gofpdf.StateGet(pdf)
	
	// Change properties
	pdf.SetDrawColor(128, 128, 128)
	pdf.SetFillColor(64, 64, 64)
	pdf.SetTextColor(200, 200, 200)
	pdf.SetLineWidth(1.0)
	
	// Restore state
	state.Put(pdf)
	
	// Verify colors restored
	r, g, b := pdf.GetDrawColor()
	if r != 255 || g != 0 || b != 0 {
		t.Errorf("DrawColor not restored correctly: got (%d,%d,%d)", r, g, b)
	}
	
	r, g, b = pdf.GetFillColor()
	if r != 0 || g != 255 || b != 0 {
		t.Errorf("FillColor not restored correctly: got (%d,%d,%d)", r, g, b)
	}
	
	r, g, b = pdf.GetTextColor()
	if r != 0 || g != 0 || b != 255 {
		t.Errorf("TextColor not restored correctly: got (%d,%d,%d)", r, g, b)
	}
}

func TestRGBType(t *testing.T) {
	rgb := gofpdf.RGBType{R: 255, G: 128, B: 64}
	
	if rgb.R != 255 {
		t.Errorf("R = %d; want 255", rgb.R)
	}
	if rgb.G != 128 {
		t.Errorf("G = %d; want 128", rgb.G)
	}
	if rgb.B != 64 {
		t.Errorf("B = %d; want 64", rgb.B)
	}
}

func TestRGBAType(t *testing.T) {
	rgba := gofpdf.RGBAType{R: 255, G: 128, B: 64, Alpha: 0.75}
	
	if rgba.R != 255 {
		t.Errorf("R = %d; want 255", rgba.R)
	}
	if rgba.G != 128 {
		t.Errorf("G = %d; want 128", rgba.G)
	}
	if rgba.B != 64 {
		t.Errorf("B = %d; want 64", rgba.B)
	}
	if rgba.Alpha != 0.75 {
		t.Errorf("Alpha = %f; want 0.75", rgba.Alpha)
	}
}

func TestSplitText(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	// Test simple text
	lines := pdf.SplitText("Hello World", 100)
	if len(lines) != 1 {
		t.Errorf("Simple text split into %d lines; want 1", len(lines))
	}
	
	// Test text with newlines
	lines = pdf.SplitText("Line 1\nLine 2\nLine 3", 100)
	if len(lines) != 3 {
		t.Errorf("Three-line text split into %d lines; want 3", len(lines))
	}
	
	// Test empty string
	lines = pdf.SplitText("", 100)
	if len(lines) != 0 {
		t.Errorf("Empty text split into %d lines; want 0", len(lines))
	}
	
	// Test trailing newline removal
	lines = pdf.SplitText("Text\n\n\n", 100)
	if len(lines) > 1 {
		t.Errorf("Text with trailing newlines split into %d lines; should trim trailing newlines", len(lines))
	}
}

func TestSplitLines(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	// Test simple text
	lines := pdf.SplitLines([]byte("Hello World"), 100)
	if len(lines) != 1 {
		t.Errorf("Simple text split into %d lines; want 1", len(lines))
	}
	
	// Test empty input
	lines = pdf.SplitLines([]byte(""), 100)
	if len(lines) != 0 {
		t.Errorf("Empty text split into %d lines; want 0", len(lines))
	}
	
	// Test with spaces
	lines = pdf.SplitLines([]byte("Word1 Word2 Word3"), 100)
	if len(lines) < 1 {
		t.Error("Text with spaces should split into at least 1 line")
	}
}

func TestGetStringWidth(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	// Test empty string
	width := pdf.GetStringWidth("")
	if width != 0 {
		t.Errorf("Empty string width = %f; want 0", width)
	}
	
	// Test non-empty string
	width = pdf.GetStringWidth("Hello")
	if width <= 0 {
		t.Error("Non-empty string should have positive width")
	}
	
	// Test that longer strings have greater width
	width1 := pdf.GetStringWidth("A")
	width2 := pdf.GetStringWidth("AA")
	if width2 <= width1 {
		t.Error("Longer string should have greater width")
	}
}

func TestPointConvert(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	
	// Test conversion
	pt := pdf.PointConvert(10)
	if pt <= 0 {
		t.Error("PointConvert should return positive value")
	}
	
	// Test PointToUnitConvert
	unit := pdf.PointToUnitConvert(10)
	if unit <= 0 {
		t.Error("PointToUnitConvert should return positive value")
	}
	
	// Test UnitToPointConvert
	pt = pdf.UnitToPointConvert(10)
	if pt <= 0 {
		t.Error("UnitToPointConvert should return positive value")
	}
}

func TestGetFontSize(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	ptSize, unitSize := pdf.GetFontSize()
	if ptSize <= 0 {
		t.Error("Font size in points should be positive")
	}
	if unitSize <= 0 {
		t.Error("Font size in units should be positive")
	}
}

func TestGetPageSize(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	w, h := pdf.GetPageSize()
	if w <= 0 || h <= 0 {
		t.Errorf("Page size should be positive: got %f x %f", w, h)
	}
	
	// A4 portrait: width should be less than height
	if w >= h {
		t.Error("A4 portrait width should be less than height")
	}
}

func TestGetMargins(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(10, 20, 15)
	pdf.SetAutoPageBreak(true, 25)
	
	left, top, right, bottom := pdf.GetMargins()
	
	if left != 10 {
		t.Errorf("Left margin = %f; want 10", left)
	}
	if top != 20 {
		t.Errorf("Top margin = %f; want 20", top)
	}
	if right != 15 {
		t.Errorf("Right margin = %f; want 15", right)
	}
	if bottom != 25 {
		t.Errorf("Bottom margin = %f; want 25", bottom)
	}
}

func TestGetXY(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	pdf.SetXY(50, 75)
	x, y := pdf.GetXY()
	
	if x != 50 {
		t.Errorf("X = %f; want 50", x)
	}
	if y != 75 {
		t.Errorf("Y = %f; want 75", y)
	}
	
	// Test GetX and GetY separately
	if pdf.GetX() != 50 {
		t.Errorf("GetX() = %f; want 50", pdf.GetX())
	}
	if pdf.GetY() != 75 {
		t.Errorf("GetY() = %f; want 75", pdf.GetY())
	}
}

func TestPageNo(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	
	// No pages yet
	if pdf.PageNo() != 0 {
		t.Errorf("PageNo() = %d before AddPage; want 0", pdf.PageNo())
	}
	
	pdf.AddPage()
	if pdf.PageNo() != 1 {
		t.Errorf("PageNo() = %d after first AddPage; want 1", pdf.PageNo())
	}
	
	pdf.AddPage()
	if pdf.PageNo() != 2 {
		t.Errorf("PageNo() = %d after second AddPage; want 2", pdf.PageNo())
	}
}

func TestSetDrawColor(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	pdf.SetDrawColor(128, 64, 200)
	r, g, b := pdf.GetDrawColor()
	
	if r != 128 || g != 64 || b != 200 {
		t.Errorf("DrawColor = (%d,%d,%d); want (128,64,200)", r, g, b)
	}
}

func TestSetFillColor(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	pdf.SetFillColor(200, 150, 100)
	r, g, b := pdf.GetFillColor()
	
	if r != 200 || g != 150 || b != 100 {
		t.Errorf("FillColor = (%d,%d,%d); want (200,150,100)", r, g, b)
	}
}

func TestSetTextColor(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	pdf.SetTextColor(50, 100, 150)
	r, g, b := pdf.GetTextColor()
	
	if r != 50 || g != 100 || b != 150 {
		t.Errorf("TextColor = (%d,%d,%d); want (50,100,150)", r, g, b)
	}
}

func TestLineWidth(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	pdf.SetLineWidth(2.5)
	width := pdf.GetLineWidth()
	
	if width != 2.5 {
		t.Errorf("LineWidth = %f; want 2.5", width)
	}
}

func TestCellMargin(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	pdf.SetCellMargin(3.5)
	margin := pdf.GetCellMargin()
	
	if margin != 3.5 {
		t.Errorf("CellMargin = %f; want 3.5", margin)
	}
}

func TestAlpha(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	pdf.SetAlpha(0.7, "Multiply")
	alpha, blendMode := pdf.GetAlpha()
	
	if alpha != 0.7 {
		t.Errorf("Alpha = %f; want 0.7", alpha)
	}
	if blendMode != "Multiply" {
		t.Errorf("BlendMode = %q; want \"Multiply\"", blendMode)
	}
}

