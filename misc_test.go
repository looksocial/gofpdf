package gofpdf_test

import (
	"testing"

	gofpdf "github.com/looksocial/gofpdf"
	"github.com/looksocial/gofpdf/internal/example"
)

// Test attachment operations
func TestSetAttachments(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	attachments := []gofpdf.Attachment{
		{
			Filename: "test1.txt",
			Content:  []byte("Test content 1"),
		},
		{
			Filename: "test2.txt",
			Content:  []byte("Test content 2"),
		},
	}

	pdf.SetAttachments(attachments)
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "With attachments")

	if pdf.Error() != nil {
		t.Errorf("SetAttachments failed: %v", pdf.Error())
	}
}

func TestAddAttachmentAnnotation(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	attachment := gofpdf.Attachment{
		Filename:    "annotated.txt",
		Content:     []byte("Annotation content"),
		Description: "Test annotation",
	}

	pdf.AddAttachmentAnnotation(&attachment, 10, 10, 50, 20)
	pdf.Rect(10, 10, 50, 20, "D")
	pdf.Cell(50, 20, "Attachment area")

	if pdf.Error() != nil {
		t.Errorf("AddAttachmentAnnotation failed: %v", pdf.Error())
	}
}

func TestSVGBasicWrite(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Create a simple SVG shape - triangle
	svg := gofpdf.SVGBasicType{
		Wd: 100,
		Ht: 100,
	}

	// SVG path segments
	path := []gofpdf.SVGBasicSegmentType{
		{Cmd: 'M', Arg: [6]float64{10, 10, 0, 0, 0, 0}},
		{Cmd: 'L', Arg: [6]float64{90, 10, 0, 0, 0, 0}},
		{Cmd: 'L', Arg: [6]float64{50, 90, 0, 0, 0, 0}},
		{Cmd: 'Z', Arg: [6]float64{0, 0, 0, 0, 0, 0}},
	}

	svg.Segments = append(svg.Segments, path)

	pdf.SetLineWidth(0.5)
	pdf.SVGBasicWrite(&svg, 1.0)

	if pdf.Error() != nil {
		t.Errorf("SVGBasicWrite failed: %v", pdf.Error())
	}
}

func TestHTMLBasicNew(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	html := pdf.HTMLBasicNew()

	htmlContent := `This is <b>bold</b>, <i>italic</i>, and <u>underlined</u> text.<br>
	<center>Centered text</center>
	<right>Right aligned</right>
	Regular text with <a href="http://example.com">link</a>.`

	html.Write(5, htmlContent)

	if pdf.Error() != nil {
		t.Errorf("HTMLBasicNew failed: %v", pdf.Error())
	}
}

func TestSetLeftMargin(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	pdf.SetLeftMargin(20)
	left, _, _, _ := pdf.GetMargins()

	if left != 20 {
		t.Errorf("Left margin should be 20, got %f", left)
	}
}

func TestSetTopMargin(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.SetTopMargin(25)
	pdf.AddPage()

	y := pdf.GetY()
	if y != 25 {
		t.Errorf("Page should start at top margin %f, got %f", 25.0, y)
	}
}

func TestSetRightMargin(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetRightMargin(30)
	left, _, right, _ := pdf.GetMargins()

	if right != 30 {
		t.Errorf("Right margin should be 30, got %f", right)
	}

	if left >= right {
		t.Error("Left margin should be less than right margin")
	}
}

func TestAddPageWithDifferentFormats(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	w1, h1 := pdf.GetPageSize()

	pdf.AddPageFormat("L", gofpdf.SizeType{Wd: 100, Ht: 200})
	w2, h2 := pdf.GetPageSize()

	if w1 == w2 && h1 == h2 {
		t.Error("Different page formats should have different sizes")
	}
}

func TestOutput(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Test")

	fileStr := example.Filename("test_output_misc")
	err := pdf.OutputFileAndClose(fileStr)
	if err != nil {
		t.Errorf("Output failed: %v", err)
	}
}

func TestCellWithBorder(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	// Test different border styles
	pdf.CellFormat(40, 10, "No border", "", 1, "L", false, 0, "")
	pdf.CellFormat(40, 10, "Full border", "1", 1, "L", false, 0, "")
	pdf.CellFormat(40, 10, "LTRB", "LTRB", 1, "L", false, 0, "")
	pdf.CellFormat(40, 10, "LR", "LR", 1, "L", false, 0, "")

	if pdf.Error() != nil {
		t.Errorf("Cell with borders failed: %v", pdf.Error())
	}
}

func TestCellWithLink(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "U", 12)

	link := pdf.AddLink()
	pdf.SetLink(link, 0, -1)

	pdf.CellFormat(40, 10, "Internal link", "", 1, "L", false, link, "")
	pdf.CellFormat(40, 10, "External link", "", 1, "L", false, 0, "http://example.com")

	if pdf.Error() != nil {
		t.Errorf("Cell with links failed: %v", pdf.Error())
	}
}

func TestImage(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Test with actual image file
	pdf.Image("image/logo.png", 10, 10, 30, 0, false, "", 0, "")

	if pdf.Error() != nil {
		t.Errorf("Image failed: %v", pdf.Error())
	}
}

func TestImageOptions(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	var opt gofpdf.ImageOptions
	opt.ImageType = "png"
	opt.ReadDpi = true

	pdf.ImageOptions("image/logo.png", 10, 40, 30, 0, false, opt, 0, "")

	if pdf.Error() != nil {
		t.Errorf("ImageOptions failed: %v", pdf.Error())
	}
}

func TestGetImageInfo(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Register image first
	pdf.Image("image/logo.png", 10, 10, 30, 0, false, "", 0, "")

	info := pdf.GetImageInfo("image/logo.png")
	if info != nil {
		if info.Width() <= 0 {
			t.Error("Image width should be positive")
		}
		if info.Height() <= 0 {
			t.Error("Image height should be positive")
		}
	}
}

func TestRegisterImage(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	info := pdf.RegisterImage("image/logo.png", "")
	if info == nil && pdf.Error() == nil {
		t.Error("RegisterImage should return info or set error")
	}
}

func TestCircleFill(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFillColor(200, 200, 255)

	pdf.Circle(50, 50, 20, "F")
	pdf.Circle(100, 50, 20, "FD")

	if pdf.Error() != nil {
		t.Errorf("Circle fill failed: %v", pdf.Error())
	}
}

func TestEllipseFill(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFillColor(255, 200, 200)

	pdf.Ellipse(50, 80, 30, 20, 0, "F")
	pdf.Ellipse(100, 80, 30, 20, 45, "FD")

	if pdf.Error() != nil {
		t.Errorf("Ellipse fill failed: %v", pdf.Error())
	}
}

func TestPolygonFill(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFillColor(200, 255, 200)

	points := []gofpdf.PointType{
		{X: 50, Y: 20},
		{X: 80, Y: 40},
		{X: 50, Y: 60},
		{X: 20, Y: 40},
	}

	pdf.Polygon(points, "F")
	pdf.Polygon(points, "FD")

	if pdf.Error() != nil {
		t.Errorf("Polygon fill failed: %v", pdf.Error())
	}
}

func TestDrawPathStyles(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFillColor(220, 220, 255)

	pdf.MoveTo(20, 20)
	pdf.LineTo(40, 20)
	pdf.LineTo(30, 40)
	pdf.ClosePath()

	// Test different draw styles
	pdf.DrawPath("D")

	pdf.MoveTo(60, 20)
	pdf.LineTo(80, 20)
	pdf.LineTo(70, 40)
	pdf.ClosePath()
	pdf.DrawPath("F")

	pdf.MoveTo(100, 20)
	pdf.LineTo(120, 20)
	pdf.LineTo(110, 40)
	pdf.ClosePath()
	pdf.DrawPath("FD")

	if pdf.Error() != nil {
		t.Errorf("DrawPath styles failed: %v", pdf.Error())
	}
}

func TestSetXmpMetadata(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	xmp := []byte(`<?xml version="1.0"?>
<x:xmpmeta xmlns:x="adobe:ns:meta/">
  <rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#">
    <rdf:Description rdf:about="" xmlns:dc="http://purl.org/dc/elements/1.1/">
      <dc:title>Test Document</dc:title>
    </rdf:Description>
  </rdf:RDF>
</x:xmpmeta>`)

	pdf.SetXmpMetadata(xmp)
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "XMP metadata")

	if pdf.Error() != nil {
		t.Errorf("SetXmpMetadata failed: %v", pdf.Error())
	}
}

func TestDrawPathFillModes(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFillColor(200, 200, 255)
	pdf.SetDrawColor(100, 100, 200)

	// Test different fill modes
	styles := []string{"D", "F", "FD", "DF", "B", "B*", "f", "f*"}

	x := 10.0
	for _, style := range styles {
		pdf.MoveTo(x, 20)
		pdf.LineTo(x+20, 20)
		pdf.LineTo(x+10, 40)
		pdf.ClosePath()
		pdf.DrawPath(style)
		x += 25
	}

	if pdf.Error() != nil {
		t.Errorf("DrawPath fill modes failed: %v", pdf.Error())
	}
}

func TestArcVariations(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetDrawColor(128, 0, 128)

	// Different arc angles and degeneracy
	pdf.Arc(30, 30, 20, 15, 0, 0, 90, "D")
	pdf.Arc(70, 30, 20, 15, 0, 90, 180, "D")
	pdf.Arc(110, 30, 20, 15, 0, 0, 360, "D")
	pdf.Arc(150, 30, 20, 15, 45, 45, 135, "D")

	if pdf.Error() != nil {
		t.Errorf("Arc variations failed: %v", pdf.Error())
	}
}

func TestCurveVariations(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetDrawColor(0, 128, 128)

	// Different curve styles
	pdf.Curve(10, 60, 30, 40, 50, 60, "D")
	pdf.Curve(60, 60, 80, 40, 100, 60, "F")
	pdf.Curve(110, 60, 130, 40, 150, 60, "FD")

	if pdf.Error() != nil {
		t.Errorf("Curve variations failed: %v", pdf.Error())
	}
}

func TestCurveBezierCubicVariations(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetDrawColor(128, 128, 0)
	pdf.SetFillColor(255, 255, 200)

	// Different styles
	pdf.CurveBezierCubic(10, 90, 20, 70, 40, 70, 50, 90, "D")
	pdf.CurveBezierCubic(60, 90, 70, 70, 90, 70, 100, 90, "F")
	pdf.CurveBezierCubic(110, 90, 120, 70, 140, 70, 150, 90, "FD")

	if pdf.Error() != nil {
		t.Errorf("CurveBezierCubic variations failed: %v", pdf.Error())
	}
}

func TestEllipseRotation(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetDrawColor(255, 100, 0)

	// Test different rotation angles
	angles := []float64{0, 30, 45, 60, 90}
	x := 30.0
	for _, angle := range angles {
		pdf.Ellipse(x, 120, 15, 10, angle, "D")
		x += 35
	}

	if pdf.Error() != nil {
		t.Errorf("Ellipse rotation failed: %v", pdf.Error())
	}
}

func TestRectStyles(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetDrawColor(100, 100, 100)
	pdf.SetFillColor(220, 220, 255)

	// Test all rect styles
	pdf.Rect(10, 150, 30, 20, "")
	pdf.Rect(50, 150, 30, 20, "D")
	pdf.Rect(90, 150, 30, 20, "F")
	pdf.Rect(130, 150, 30, 20, "FD")
	pdf.Rect(170, 150, 30, 20, "DF")

	if pdf.Error() != nil {
		t.Errorf("Rect styles failed: %v", pdf.Error())
	}
}

func TestWriteWithNewlines(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	pdf.Write(5, "Line 1\nLine 2\nLine 3")

	if pdf.Error() != nil {
		t.Errorf("Write with newlines failed: %v", pdf.Error())
	}
}

func TestMultiCellBorder(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	text := "This is a multi-line cell with border"

	pdf.MultiCell(60, 5, text, "1", "L", false)
	pdf.MultiCell(60, 5, text, "LTRB", "C", false)
	pdf.MultiCell(60, 5, text, "", "R", false)

	if pdf.Error() != nil {
		t.Errorf("MultiCell with border failed: %v", pdf.Error())
	}
}

func TestMultiCellFill(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.SetFillColor(230, 230, 255)

	text := "This is a filled multi-line cell"

	pdf.MultiCell(60, 5, text, "1", "L", true)

	if pdf.Error() != nil {
		t.Errorf("MultiCell with fill failed: %v", pdf.Error())
	}
}

func TestSubWrite(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	pdf.Write(10, "H")
	pdf.SubWrite(10, "2", 8, -3, 0, "")
	pdf.Write(10, "O")

	pdf.Ln(15)

	pdf.Write(10, "E=mc")
	pdf.SubWrite(10, "2", 8, 4, 0, "")

	if pdf.Error() != nil {
		t.Errorf("SubWrite failed: %v", pdf.Error())
	}
}
