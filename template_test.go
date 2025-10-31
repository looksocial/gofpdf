package gofpdf_test

import (
	"testing"

	gofpdf "github.com/looksocial/gofpdf"
	"github.com/looksocial/gofpdf/internal/example"
)

// Test template operations
func TestCreateTemplate(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	tpl := pdf.CreateTemplate(func(t *gofpdf.Tpl) {
		t.SetFont("Arial", "B", 16)
		t.Text(10, 10, "Template content")
	})

	if tpl == nil {
		t.Error("CreateTemplate should return non-nil template")
	}

	pdf.AddPage()
	pdf.UseTemplate(tpl)

	if pdf.Error() != nil {
		t.Errorf("Template operations failed: %v", pdf.Error())
	}
}

func TestTemplateSize(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	tpl := pdf.CreateTemplate(func(t *gofpdf.Tpl) {
		t.SetFont("Arial", "", 12)
		t.Text(10, 10, "Test")
	})

	_, size := tpl.Size()
	if size.Wd <= 0 || size.Ht <= 0 {
		t.Errorf("Template size should be positive: %f x %f", size.Wd, size.Ht)
	}
}

func TestUseTemplateScaled(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	tpl := pdf.CreateTemplate(func(t *gofpdf.Tpl) {
		t.Rect(0, 0, 50, 50, "D")
	})

	pdf.AddPage()

	// Use template with scaling
	corner := gofpdf.PointType{X: 10, Y: 10}
	size := gofpdf.SizeType{Wd: 30, Ht: 30}
	pdf.UseTemplateScaled(tpl, corner, size)

	if pdf.Error() != nil {
		t.Errorf("UseTemplateScaled failed: %v", pdf.Error())
	}
}

func TestTemplateNumPages(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	tpl := pdf.CreateTemplate(func(t *gofpdf.Tpl) {
		t.AddPage()
		t.AddPage()
		t.AddPage()
	})

	if tpl.NumPages() != 4 {
		t.Errorf("Template should have 4 pages, got %d", tpl.NumPages())
	}
}

func TestTemplateFromPage(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	tpl := pdf.CreateTemplate(func(t *gofpdf.Tpl) {
		t.AddPage()
		t.AddPage()
	})

	tpl2, err := tpl.FromPage(1)
	if err != nil {
		t.Errorf("FromPage failed: %v", err)
	}

	pdf.AddPage()
	pdf.UseTemplate(tpl2)

	if pdf.Error() != nil {
		t.Errorf("Template from page failed: %v", pdf.Error())
	}
}

func TestTemplateFromPages(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	tpl := pdf.CreateTemplate(func(t *gofpdf.Tpl) {
		t.AddPage()
		t.AddPage()
	})

	pages := tpl.FromPages()
	if len(pages) != 3 {
		t.Errorf("FromPages should return 3 templates, got %d", len(pages))
	}

	for _, page := range pages {
		pdf.AddPage()
		pdf.UseTemplate(page)
	}

	if pdf.Error() != nil {
		t.Errorf("Template from pages failed: %v", pdf.Error())
	}
}

func TestTemplateSerialize(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	tpl := pdf.CreateTemplate(func(t *gofpdf.Tpl) {
		t.SetFont("Arial", "", 12)
		t.Text(10, 10, "Serializable")
	})

	// Serialize
	bytes, err := tpl.Serialize()
	if err != nil {
		t.Errorf("Template serialize failed: %v", err)
	}

	if len(bytes) == 0 {
		t.Error("Serialized template should not be empty")
	}

	// Deserialize
	tpl2, err := gofpdf.DeserializeTemplate(bytes)
	if err != nil {
		t.Errorf("Template deserialize failed: %v", err)
	}

	pdf.AddPage()
	pdf.UseTemplate(tpl2)

	if pdf.Error() != nil {
		t.Errorf("Deserialized template use failed: %v", pdf.Error())
	}
}

func TestAddLayer(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	layer1 := pdf.AddLayer("Layer 1", true)
	layer2 := pdf.AddLayer("Layer 2", false)

	if layer1 == layer2 {
		t.Error("Different layers should have different IDs")
	}

	pdf.BeginLayer(layer1)
	pdf.Text(10, 10, "Layer 1 content")
	pdf.EndLayer()

	pdf.BeginLayer(layer2)
	pdf.Text(10, 20, "Layer 2 content")
	pdf.EndLayer()

	if pdf.Error() != nil {
		t.Errorf("Layer operations failed: %v", pdf.Error())
	}
}

func TestOpenLayerPane(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	layer := pdf.AddLayer("Test Layer", true)
	pdf.OpenLayerPane()

	pdf.BeginLayer(layer)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Layer content")
	pdf.EndLayer()

	if pdf.Error() != nil {
		t.Errorf("OpenLayerPane failed: %v", pdf.Error())
	}
}

func TestAddPageFormat(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Add page with custom size
	pdf.AddPageFormat("L", gofpdf.SizeType{Wd: 200, Ht: 100})
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Custom size page")

	// Add another with different orientation
	pdf.AddPageFormat("P", gofpdf.SizeType{Wd: 100, Ht: 200})

	if pdf.Error() != nil {
		t.Errorf("AddPageFormat failed: %v", pdf.Error())
	}
}

func TestSetAcceptPageBreakFunc(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	pdf.SetAcceptPageBreakFunc(func() bool {
		return true
	})

	// This should not trigger page break
	pdf.Cell(40, 10, "Test")

	if pdf.Error() != nil {
		t.Errorf("SetAcceptPageBreakFunc failed: %v", pdf.Error())
	}
}

func TestSetHeaderFunc(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.SetHeaderFunc(func() {
		pdf.SetFont("Arial", "B", 12)
		pdf.Cell(0, 10, "Header")
	})

	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Content")

	if pdf.Error() != nil {
		t.Errorf("SetHeaderFunc failed: %v", pdf.Error())
	}
}

func TestSetHeaderFuncMode(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.SetHeaderFuncMode(func() {
		pdf.SetFont("Arial", "I", 10)
		pdf.Cell(0, 10, "Header Mode")
	}, true)

	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Content")

	if pdf.Error() != nil {
		t.Errorf("SetHeaderFuncMode failed: %v", pdf.Error())
	}
}

func TestSetFooterFunc(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 8)
		pdf.Cell(0, 10, "Footer")
	})

	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Content")

	if pdf.Error() != nil {
		t.Errorf("SetFooterFunc failed: %v", pdf.Error())
	}
}

func TestSetFooterFuncLpi(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.SetFooterFuncLpi(func(lastPage bool) {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 8)
		if lastPage {
			pdf.Cell(0, 10, "Last page")
		} else {
			pdf.Cell(0, 10, "Page footer")
		}
	})

	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Content")

	if pdf.Error() != nil {
		t.Errorf("SetFooterFuncLpi failed: %v", pdf.Error())
	}
}

func TestSetJavascript(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetJavascript("console.log('test');")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "JS test")

	if pdf.Error() != nil {
		t.Errorf("SetJavascript failed: %v", pdf.Error())
	}
}

func TestSetProtection(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.SetProtection(gofpdf.CnProtectPrint, "userpass", "ownerpass")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Protected")

	if pdf.Error() != nil {
		t.Errorf("SetProtection failed: %v", pdf.Error())
	}
}

func TestAddSpotColor(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.AddSpotColor("Custom Cyan", 100, 0, 0, 0)
	pdf.SetFillSpotColor("Custom Cyan", 100)
	pdf.Rect(10, 10, 50, 50, "F")

	if pdf.Error() != nil {
		t.Errorf("Spot color operations failed: %v", pdf.Error())
	}
}

func TestSetDrawSpotColor(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.AddSpotColor("Custom Magenta", 0, 100, 0, 0)
	pdf.SetDrawSpotColor("Custom Magenta", 80)
	pdf.Circle(50, 50, 20, "D")

	if pdf.Error() != nil {
		t.Errorf("SetDrawSpotColor failed: %v", pdf.Error())
	}
}

func TestSetTextSpotColor(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	pdf.AddSpotColor("Custom Yellow", 0, 0, 100, 0)
	pdf.SetTextSpotColor("Custom Yellow", 90)
	pdf.Text(10, 10, "Spot color text")

	if pdf.Error() != nil {
		t.Errorf("SetTextSpotColor failed: %v", pdf.Error())
	}
}

func TestOrientations(t *testing.T) {
	tests := []struct {
		name        string
		orientation string
	}{
		{"Portrait", "P"},
		{"Landscape", "L"},
		{"Portrait explicit", "Portrait"},
		{"Landscape explicit", "Landscape"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pdf := gofpdf.New(tt.orientation, "mm", "A4", "")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 12)
			pdf.Cell(40, 10, tt.name)

			if pdf.Error() != nil {
				t.Errorf("Orientation %s failed: %v", tt.orientation, pdf.Error())
			}
		})
	}
}

func TestPageFormats(t *testing.T) {
	formats := []string{"A3", "A4", "A5", "Letter", "Legal"}

	for _, format := range formats {
		t.Run(format, func(t *testing.T) {
			pdf := gofpdf.New("P", "mm", format, "")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 12)
			pdf.Cell(40, 10, format)

			w, h := pdf.GetPageSize()
			if w <= 0 || h <= 0 {
				t.Errorf("Format %s has invalid size: %f x %f", format, w, h)
			}

			if pdf.Error() != nil {
				t.Errorf("Format %s failed: %v", format, pdf.Error())
			}
		})
	}
}

func TestUnits(t *testing.T) {
	units := []string{"mm", "cm", "in", "pt"}

	for _, unit := range units {
		t.Run(unit, func(t *testing.T) {
			pdf := gofpdf.New("P", unit, "A4", "")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 12)
			pdf.Cell(40, 10, unit)

			if pdf.Error() != nil {
				t.Errorf("Unit %s failed: %v", unit, pdf.Error())
			}
		})
	}
}

func TestNewCustom(t *testing.T) {
	init := gofpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "mm",
		Size:           gofpdf.SizeType{Wd: 100, Ht: 150},
		FontDirStr:     "",
	}

	pdf := gofpdf.NewCustom(&init)
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Custom init")

	w, h := pdf.GetPageSize()
	if w != 100 || h != 150 {
		t.Errorf("Custom size should be 100x150, got %fx%f", w, h)
	}

	if pdf.Error() != nil {
		t.Errorf("NewCustom failed: %v", pdf.Error())
	}
}

func TestSetPageBox(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetPageBox("crop", 10, 10, 190, 277)
	pdf.SetPageBox("trim", 10, 10, 190, 277)
	pdf.SetPageBox("bleed", 5, 5, 200, 287)
	pdf.SetPageBox("art", 15, 15, 180, 267)

	if pdf.Error() != nil {
		t.Errorf("SetPageBox failed: %v", pdf.Error())
	}
}

func TestOutputAndClose(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Test output")

	fileStr := example.Filename("test_output")
	err := pdf.OutputFileAndClose(fileStr)
	if err != nil {
		t.Errorf("OutputFileAndClose failed: %v", err)
	}
}

func TestRoundedRectExt(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFillColor(200, 220, 255)

	// Different radius for each corner
	pdf.RoundedRectExt(10, 10, 80, 40, 5, 10, 15, 20, "FD")

	if pdf.Error() != nil {
		t.Errorf("RoundedRectExt failed: %v", pdf.Error())
	}
}

func TestClipRoundedRect(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.ClipRoundedRect(10, 10, 80, 40, 5, false)
	pdf.SetFillColor(255, 200, 200)
	pdf.Rect(0, 0, 100, 60, "F")
	pdf.ClipEnd()

	if pdf.Error() != nil {
		t.Errorf("ClipRoundedRect failed: %v", pdf.Error())
	}
}

func TestClipRoundedRectExt(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.ClipRoundedRectExt(10, 10, 80, 40, 5, 10, 5, 10, false)
	pdf.SetFillColor(200, 255, 200)
	pdf.Rect(0, 0, 100, 60, "F")
	pdf.ClipEnd()

	if pdf.Error() != nil {
		t.Errorf("ClipRoundedRectExt failed: %v", pdf.Error())
	}
}

func TestClipPolygon(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	points := []gofpdf.PointType{
		{X: 50, Y: 20},
		{X: 80, Y: 50},
		{X: 50, Y: 80},
		{X: 20, Y: 50},
	}

	pdf.ClipPolygon(points, false)
	pdf.SetFillColor(255, 255, 200)
	pdf.Rect(0, 0, 100, 100, "F")
	pdf.ClipEnd()

	if pdf.Error() != nil {
		t.Errorf("ClipPolygon failed: %v", pdf.Error())
	}
}
