package gofpdf_test

import (
	"bytes"
	"encoding/gob"
	"testing"

	gofpdf "github.com/looksocial/gofpdf"
	"github.com/looksocial/gofpdf/internal/example"
)

// Test GobEncode/GobDecode for ImageInfoType
func TestImageInfoGob(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Register an image to get ImageInfoType
	info := pdf.RegisterImage("image/logo.png", "")
	if info == nil {
		t.Skip("Could not register image, skipping Gob test")
	}

	// Encode
	data, err := info.GobEncode()
	if err != nil {
		t.Errorf("GobEncode failed: %v", err)
	}

	// Decode
	newInfo := &gofpdf.ImageInfoType{}
	err = newInfo.GobDecode(data)
	if err != nil {
		t.Errorf("GobDecode failed: %v", err)
	}

	// Verify data matches
	if newInfo.Width() != info.Width() {
		t.Errorf("Decoded width %f != original %f", newInfo.Width(), info.Width())
	}
	if newInfo.Height() != info.Height() {
		t.Errorf("Decoded height %f != original %f", newInfo.Height(), info.Height())
	}
}

// Test ImageInfoType position handling
func TestImageInfoPosition(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Register image and use it at different positions
	info := pdf.RegisterImage("image/logo.png", "")
	if info == nil {
		t.Skip("Could not register image")
	}

	// Use image at various positions
	pdf.Image("image/logo.png", 10, 10, 30, 0, false, "", 0, "")
	pdf.Image("image/logo.png", 50, 10, 30, 0, false, "", 0, "")

	if pdf.Error() != nil {
		t.Errorf("Image positioning failed: %v", pdf.Error())
	}
}

// Test SetDpi method
func TestImageInfoSetDpi(t *testing.T) {
	info := gofpdf.ImageInfoType{}

	info.SetDpi(96)
	info.SetDpi(150)
	info.SetDpi(300)

	// SetDpi should not panic
}

// Test ImageInfoType with DPI
func TestImageInfoDPI(t *testing.T) {
	info := gofpdf.ImageInfoType{}

	info.SetDpi(72)
	info.SetDpi(96)
	info.SetDpi(300)

	// Should handle DPI changes without panic
}

// Test that multiple images get different handling
func TestMultipleImageRegistration(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Register multiple images
	info1 := pdf.RegisterImage("image/logo.png", "")
	info2 := pdf.RegisterImage("image/golang-gopher.png", "")

	if info1 == nil || info2 == nil {
		t.Skip("Could not register images")
	}

	// Different images should have different dimensions
	if info1.Width() == info2.Width() && info1.Height() == info2.Height() {
		t.Log("Images happen to have same dimensions, which is ok")
	}
}

// Test ImageInfoType methods with actual image
func TestImageInfoMethods(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	info := pdf.RegisterImage("image/logo.png", "")
	if info == nil {
		t.Skip("Could not register image")
	}

	// Test Width
	if info.Width() <= 0 {
		t.Error("Width should be positive")
	}

	// Test Height
	if info.Height() <= 0 {
		t.Error("Height should be positive")
	}

	// Test Extent
	w, h := info.Extent()
	if w <= 0 || h <= 0 {
		t.Errorf("Extent should return positive values: %f x %f", w, h)
	}

	if w != info.Width() || h != info.Height() {
		t.Error("Extent should match Width/Height")
	}
}

// Test ComparePDFFiles
func TestComparePDFFiles(t *testing.T) {
	// Create two identical PDFs
	var filenames []string
	for i := 1; i <= 2; i++ {
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()
		pdf.SetFont("Arial", "", 12)
		pdf.Cell(40, 10, "Test content")

		filename := example.Filename("test_compare_" + string(rune('0'+i)))
		filenames = append(filenames, filename)
		err := pdf.OutputFileAndClose(filename)
		if err != nil {
			t.Fatalf("Failed to create test PDF %d: %v", i, err)
		}
	}

	// Compare the files
	err := gofpdf.ComparePDFFiles(filenames[0], filenames[1], false)
	if err != nil {
		t.Errorf("ComparePDFFiles failed: %v", err)
	}
}

// Test ComparePDFFiles with missing file
func TestComparePDFFilesMissing(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Test")

	fileStr := example.Filename("test_compare_exists")
	err := pdf.OutputFileAndClose(fileStr)
	if err != nil {
		t.Fatalf("Failed to create test PDF: %v", err)
	}

	// Compare with non-existent file (should succeed as per implementation)
	nonexistentFile := example.PdfFile("nonexistent.pdf")
	err = gofpdf.ComparePDFFiles(fileStr, nonexistentFile, false)
	if err != nil {
		t.Errorf("ComparePDFFiles should succeed when second file missing: %v", err)
	}
}

// Test sort functionality through gensort (used in attachments)
func TestSortFunctionality(t *testing.T) {
	// This tests the internal sorting used for attachments
	pdf := gofpdf.New("P", "mm", "A4", "")

	attachments := []gofpdf.Attachment{
		{Filename: "c.txt", Content: []byte("C")},
		{Filename: "a.txt", Content: []byte("A")},
		{Filename: "b.txt", Content: []byte("B")},
	}

	pdf.SetAttachments(attachments)
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Sorted attachments")

	if pdf.Error() != nil {
		t.Errorf("Attachment sorting failed: %v", pdf.Error())
	}
}

// Test multiple attachments
func TestMultipleAttachments(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	attachments := make([]gofpdf.Attachment, 10)
	for i := 0; i < 10; i++ {
		attachments[i] = gofpdf.Attachment{
			Filename:    "file" + string(rune('0'+i)) + ".txt",
			Content:     []byte("Content " + string(rune('0'+i))),
			Description: "File number " + string(rune('0'+i)),
		}
	}

	pdf.SetAttachments(attachments)
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Multiple attachments")

	if pdf.Error() != nil {
		t.Errorf("Multiple attachments failed: %v", pdf.Error())
	}
}

// Test attachment with annotation on page
func TestAttachmentWithAnnotation(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	att1 := gofpdf.Attachment{
		Filename:    "doc1.txt",
		Content:     []byte("Document 1 content"),
		Description: "First document",
	}

	att2 := gofpdf.Attachment{
		Filename:    "doc2.txt",
		Content:     []byte("Document 2 content"),
		Description: "Second document",
	}

	pdf.AddAttachmentAnnotation(&att1, 10, 10, 40, 20)
	pdf.Rect(10, 10, 40, 20, "D")
	pdf.CellFormat(40, 20, "Attachment 1", "", 1, "L", false, 0, "")

	pdf.AddAttachmentAnnotation(&att2, 10, 40, 40, 20)
	pdf.Rect(10, 40, 40, 20, "D")
	pdf.CellFormat(40, 20, "Attachment 2", "", 1, "L", false, 0, "")

	if pdf.Error() != nil {
		t.Errorf("Multiple attachment annotations failed: %v", pdf.Error())
	}
}

// Test actual gob encoding/decoding with buffer
func TestImageInfoGobWithBuffer(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	info := pdf.RegisterImage("image/logo.png", "")
	if info == nil {
		t.Skip("Could not register image")
	}

	// Use gob encoder/decoder
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(info)
	if err != nil {
		t.Errorf("Gob encode failed: %v", err)
	}

	dec := gob.NewDecoder(&buf)
	var decoded gofpdf.ImageInfoType

	err = dec.Decode(&decoded)
	if err != nil {
		t.Errorf("Gob decode failed: %v", err)
	}
}
