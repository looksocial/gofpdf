package gofpdf_test

import (
	"math"
	"testing"

	gofpdf "github.com/looksocial/gofpdf"
)

// Test transformation functions
func TestTransformTranslate(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	pdf.TransformBegin()
	pdf.TransformTranslate(10, 20)
	pdf.Text(0, 0, "Translated")
	pdf.TransformEnd()

	if pdf.Error() != nil {
		t.Errorf("TransformTranslate failed: %v", pdf.Error())
	}
}

func TestTransformTranslateXY(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.TransformBegin()
	pdf.TransformTranslateX(15)
	pdf.TransformEnd()

	pdf.TransformBegin()
	pdf.TransformTranslateY(25)
	pdf.TransformEnd()

	if pdf.Error() != nil {
		t.Errorf("TransformTranslateX/Y failed: %v", pdf.Error())
	}
}

func TestTransformScale(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	pdf.TransformBegin()
	pdf.TransformScale(150, 150, 50, 50)
	pdf.Text(50, 50, "Scaled")
	pdf.TransformEnd()

	if pdf.Error() != nil {
		t.Errorf("TransformScale failed: %v", pdf.Error())
	}
}

func TestTransformScaleXY(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.TransformBegin()
	pdf.TransformScaleX(200, 50, 50)
	pdf.TransformEnd()

	pdf.TransformBegin()
	pdf.TransformScaleY(200, 50, 50)
	pdf.TransformEnd()

	pdf.TransformBegin()
	pdf.TransformScaleXY(150, 50, 50)
	pdf.TransformEnd()

	if pdf.Error() != nil {
		t.Errorf("TransformScale X/Y/XY failed: %v", pdf.Error())
	}
}

func TestTransformRotate(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	pdf.TransformBegin()
	pdf.TransformRotate(45, 100, 100)
	pdf.Text(100, 100, "Rotated")
	pdf.TransformEnd()

	if pdf.Error() != nil {
		t.Errorf("TransformRotate failed: %v", pdf.Error())
	}
}

func TestTransformSkew(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	pdf.TransformBegin()
	pdf.TransformSkew(10, 5, 50, 50)
	pdf.Text(50, 50, "Skewed")
	pdf.TransformEnd()

	if pdf.Error() != nil {
		t.Errorf("TransformSkew failed: %v", pdf.Error())
	}
}

func TestTransformSkewXY(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.TransformBegin()
	pdf.TransformSkewX(15, 50, 50)
	pdf.TransformEnd()

	pdf.TransformBegin()
	pdf.TransformSkewY(15, 50, 50)
	pdf.TransformEnd()

	if pdf.Error() != nil {
		t.Errorf("TransformSkewX/Y failed: %v", pdf.Error())
	}
}

func TestTransformMirror(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	pdf.TransformBegin()
	pdf.TransformMirrorHorizontal(100)
	pdf.TransformEnd()

	pdf.TransformBegin()
	pdf.TransformMirrorVertical(100)
	pdf.TransformEnd()

	pdf.TransformBegin()
	pdf.TransformMirrorPoint(100, 100)
	pdf.TransformEnd()

	pdf.TransformBegin()
	pdf.TransformMirrorLine(45, 100, 100)
	pdf.TransformEnd()

	if pdf.Error() != nil {
		t.Errorf("Transform mirror operations failed: %v", pdf.Error())
	}
}

func TestAddLink(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	link := pdf.AddLink()
	if link <= 0 {
		t.Error("AddLink should return positive link ID")
	}

	pdf.SetLink(link, 0, 100)
	pdf.Write(10, "Click here")
	pdf.WriteLinkID(10, " for link", link)

	if pdf.Error() != nil {
		t.Errorf("Link operations failed: %v", pdf.Error())
	}
}

func TestWriteLinkString(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	pdf.WriteLinkString(10, "External link", "https://example.com")

	if pdf.Error() != nil {
		t.Errorf("WriteLinkString failed: %v", pdf.Error())
	}
}

func TestBookmark(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 15)

	pdf.Bookmark("Chapter 1", 0, 0)
	pdf.Bookmark("Section 1.1", 1, -1)
	pdf.Text(10, 20, "Content")

	if pdf.Error() != nil {
		t.Errorf("Bookmark failed: %v", pdf.Error())
	}
}

func TestAliasNbPages(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	pdf.AliasNbPages("")
	pdf.AliasNbPages("{nb}")

	if pdf.Error() != nil {
		t.Errorf("AliasNbPages failed: %v", pdf.Error())
	}
}

func TestRegisterAlias(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	pdf.RegisterAlias("{custom}", "replacement")
	pdf.Write(10, "Test {custom} alias")

	if pdf.Error() != nil {
		t.Errorf("RegisterAlias failed: %v", pdf.Error())
	}
}

func TestLn(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	y1 := pdf.GetY()
	pdf.Ln(10)
	y2 := pdf.GetY()

	if y2 <= y1 {
		t.Error("Ln should increase Y position")
	}

	// Test negative Ln (uses line height)
	pdf.Ln(-1)

	if pdf.Error() != nil {
		t.Errorf("Ln failed: %v", pdf.Error())
	}
}

func TestSetDashPattern(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetDashPattern([]float64{5, 10}, 0)
	pdf.Line(10, 10, 100, 10)

	// Reset dash pattern
	pdf.SetDashPattern([]float64{}, 0)
	pdf.Line(10, 20, 100, 20)

	if pdf.Error() != nil {
		t.Errorf("SetDashPattern failed: %v", pdf.Error())
	}
}

func TestLineCapStyle(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetLineCapStyle("butt")
	pdf.SetLineCapStyle("round")
	pdf.SetLineCapStyle("square")

	if pdf.Error() != nil {
		t.Errorf("SetLineCapStyle failed: %v", pdf.Error())
	}
}

func TestLineJoinStyle(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetLineJoinStyle("miter")
	pdf.SetLineJoinStyle("round")
	pdf.SetLineJoinStyle("bevel")

	if pdf.Error() != nil {
		t.Errorf("SetLineJoinStyle failed: %v", pdf.Error())
	}
}

func TestDrawing(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetDrawColor(128, 128, 128)
	pdf.SetLineWidth(1.0)

	// Test various drawing functions
	pdf.Line(10, 10, 100, 10)
	pdf.Rect(10, 20, 50, 30, "D")
	pdf.Rect(70, 20, 50, 30, "F")
	pdf.Rect(130, 20, 50, 30, "FD")

	pdf.Circle(35, 80, 15, "D")
	pdf.Ellipse(95, 80, 25, 15, 0, "D")

	if pdf.Error() != nil {
		t.Errorf("Drawing operations failed: %v", pdf.Error())
	}
}

func TestPolygon(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetDrawColor(0, 0, 128)

	points := []gofpdf.PointType{
		{X: 50, Y: 20},
		{X: 80, Y: 50},
		{X: 50, Y: 80},
		{X: 20, Y: 50},
	}

	pdf.Polygon(points, "D")

	if pdf.Error() != nil {
		t.Errorf("Polygon failed: %v", pdf.Error())
	}
}

func TestCurve(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetDrawColor(128, 0, 0)

	pdf.Curve(10, 30, 50, 10, 90, 30, "D")

	if pdf.Error() != nil {
		t.Errorf("Curve failed: %v", pdf.Error())
	}
}

func TestCurveBezierCubic(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetDrawColor(0, 128, 0)

	pdf.CurveBezierCubic(10, 50, 30, 30, 70, 30, 90, 50, "D")

	if pdf.Error() != nil {
		t.Errorf("CurveBezierCubic failed: %v", pdf.Error())
	}
}

func TestArc(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetDrawColor(0, 0, 255)

	pdf.Arc(50, 50, 20, 20, 0, 0, 180, "D")

	if pdf.Error() != nil {
		t.Errorf("Arc failed: %v", pdf.Error())
	}
}

func TestMoveTo(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetDrawColor(200, 100, 0)

	pdf.MoveTo(20, 20)
	pdf.LineTo(80, 20)
	pdf.LineTo(50, 60)
	pdf.ClosePath()
	pdf.DrawPath("D")

	if pdf.Error() != nil {
		t.Errorf("Path operations failed: %v", pdf.Error())
	}
}

func TestCurveTo(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.MoveTo(20, 20)
	pdf.CurveTo(50, 10, 80, 20)
	pdf.DrawPath("D")

	if pdf.Error() != nil {
		t.Errorf("CurveTo failed: %v", pdf.Error())
	}
}

func TestCurveBezierCubicTo(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.MoveTo(20, 30)
	pdf.CurveBezierCubicTo(30, 10, 70, 10, 80, 30)
	pdf.DrawPath("D")

	if pdf.Error() != nil {
		t.Errorf("CurveBezierCubicTo failed: %v", pdf.Error())
	}
}

func TestArcTo(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.MoveTo(20, 20)
	pdf.ArcTo(40, 40, 20, 20, 0, 90, 0)
	pdf.DrawPath("D")

	if pdf.Error() != nil {
		t.Errorf("ArcTo failed: %v", pdf.Error())
	}
}

func TestClipping(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 24)

	pdf.ClipRect(10, 10, 80, 30, false)
	pdf.SetFillColor(200, 200, 220)
	pdf.Rect(0, 0, 100, 50, "F")
	pdf.ClipEnd()

	pdf.ClipCircle(50, 80, 20, false)
	pdf.SetFillColor(220, 200, 200)
	pdf.Rect(20, 50, 60, 60, "F")
	pdf.ClipEnd()

	pdf.ClipEllipse(50, 140, 30, 20, false)
	pdf.SetFillColor(200, 220, 200)
	pdf.Rect(10, 110, 80, 60, "F")
	pdf.ClipEnd()

	if pdf.Error() != nil {
		t.Errorf("Clipping operations failed: %v", pdf.Error())
	}
}

func TestClipText(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 48)

	pdf.ClipText(10, 50, "CLIPPED", false)
	pdf.LinearGradient(10, 20, 100, 40, 200, 200, 255, 100, 100, 255, 0, 0, 0, 1)
	pdf.ClipEnd()

	if pdf.Error() != nil {
		t.Errorf("ClipText failed: %v", pdf.Error())
	}
}

func TestRoundedRect(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetDrawColor(128, 128, 128)
	pdf.SetFillColor(200, 200, 255)

	pdf.RoundedRect(10, 10, 80, 40, 5, "1234", "FD")
	pdf.RoundedRectExt(10, 60, 80, 40, 5, 10, 5, 10, "FD")

	if pdf.Error() != nil {
		t.Errorf("RoundedRect failed: %v", pdf.Error())
	}
}

func TestGradients(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.LinearGradient(10, 10, 80, 40, 255, 0, 0, 0, 0, 255, 0, 0, 1, 1)
	pdf.Rect(10, 10, 80, 40, "F")

	pdf.RadialGradient(50, 80, 40, 40, 255, 255, 0, 255, 0, 255, 0.25, 0.75, 0.25, 0.75, 1)
	pdf.Circle(50, 80, 20, "F")

	if pdf.Error() != nil {
		t.Errorf("Gradient operations failed: %v", pdf.Error())
	}
}

func TestBeziergon(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetDrawColor(64, 64, 128)

	points := []gofpdf.PointType{
		{X: 50, Y: 20},
		{X: 55, Y: 15}, {X: 75, Y: 15}, {X: 80, Y: 20},
		{X: 85, Y: 25}, {X: 85, Y: 45}, {X: 80, Y: 50},
		{X: 75, Y: 55}, {X: 55, Y: 55}, {X: 50, Y: 50},
		{X: 45, Y: 45}, {X: 45, Y: 25}, {X: 50, Y: 20},
	}

	pdf.Beziergon(points, "D")

	if pdf.Error() != nil {
		t.Errorf("Beziergon failed: %v", pdf.Error())
	}
}

func TestSetPage(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.AddPage()
	pdf.AddPage()

	if pdf.PageNo() != 3 {
		t.Errorf("Expected 3 pages, got %d", pdf.PageNo())
	}

	pdf.SetPage(1)
	pdf.SetPage(2)
	pdf.SetPage(3)

	if pdf.Error() != nil {
		t.Errorf("SetPage failed: %v", pdf.Error())
	}
}

func TestPageSize(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	w, h, unit := pdf.PageSize(0)
	if w <= 0 || h <= 0 {
		t.Errorf("PageSize returned invalid dimensions: %f x %f", w, h)
	}
	if unit != "mm" {
		t.Errorf("PageSize returned wrong unit: %s", unit)
	}
}

func TestPointConvertRoundTrip(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Test round-trip conversion
	unit := 10.0
	pt := pdf.UnitToPointConvert(unit)
	backToUnit := pdf.PointToUnitConvert(pt)

	if math.Abs(unit-backToUnit) > 0.01 {
		t.Errorf("Point conversion round-trip failed: %f -> %f -> %f", unit, pt, backToUnit)
	}
}
