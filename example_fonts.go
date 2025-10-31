//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"

	"github.com/looksocial/gofpdf"
)

// main เรียกฟังก์ชันตัวอย่างการใช้งานฟอนต์ของแพ็กเกจเป็นลำดับ ได้แก่ การใช้งานฟอนต์ฝังตัว การโหลดฟอนต์ภายนอก เอกสารภาษาไทย และการแสดงหลายฟอนต์ในเอกสาร PDF.
func main() {
	// Example 1: Using Embedded Fonts (Simplest - Recommended)
	example1_EmbeddedFonts()

	// Example 2: Using External Font Files with Auto-Load
	example2_ExternalFontsAutoLoad()

	// Example 3: Thai Language Document
	example3_ThaiDocument()

	// Example 4: Multiple Font Families and Styles
	example4_MultipleFonts()
}

// example1_EmbeddedFonts สาธิตการสร้างเอกสาร PDF โดยใช้ฟอนต์ฝัง (embedded) โดยไม่ต้องมีไฟล์ฟอนต์ภายนอก
// ฟังก์ชันจะเขียนข้อความตัวอย่างด้วยฟอนต์ Kanit และ Sarabun ลงในไฟล์ example1_embedded.pdf และจะบันทึกไฟล์นี้ (เรียก log.Fatal หากการบันทึกล้มเหลว)
func example1_EmbeddedFonts() {
	fmt.Println("Example 1: Embedded Fonts")

	pdf := gofpdf.New("P", "mm", "A4", "")

	// Enable embedded fonts - that's it! No file paths needed
	pdf.UseEmbeddedFonts()

	pdf.AddPage()

	// Use any embedded Thai font directly
	pdf.SetFont("Kanit", "", 16)
	pdf.Cell(0, 10, "Hello with Kanit Font")
	pdf.Ln(8)

	pdf.SetFont("Kanit", "B", 16)
	pdf.Cell(0, 10, "Hello with Kanit Bold")
	pdf.Ln(8)

	pdf.SetFont("Sarabun", "", 14)
	pdf.Cell(0, 10, "Hello with Sarabun Font")

	err := pdf.OutputFileAndClose("example1_embedded.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✓ Created example1_embedded.pdf")
}

// example2_ExternalFontsAutoLoad สาธิตการใช้ไฟล์ฟอนต์ภายนอกโดยเปิดใช้งานการโหลดอัตโนมัติจากโฟลเดอร์ค้นหาฟอนต์
//
// ตั้งตำแหน่งค้นหาฟอนต์เป็น "font/th" เพื่อให้ไลบรารีสามารถโหลดไฟล์ฟอนต์ (เช่น Kanit และ Kanit-Bold) จากโฟลเดอร์ย่อยโดยอัตโนมัติ; หากฟอนต์ไม่พร้อมใช้งาน ฟังก์ชันจะพิมพ์หมายเหตุแล้วคืนค่าโดยไม่สร้างไฟล์ PDF. เมื่อฟอนต์ถูกโหลดสำเร็จ ฟังก์ชันจะสร้างไฟล์ example2_autoload.pdf ในไดเรกทอรีปัจจุบัน.
func example2_ExternalFontsAutoLoad() {
	fmt.Println("\nExample 2: External Fonts with Auto-Load")

	pdf := gofpdf.New("P", "mm", "A4", "")

	// Set font location to the Thai fonts directory
	// Auto-load will search in subdirectories
	pdf.SetFontLocation("font/th")

	pdf.AddPage()

	// Fonts will be auto-loaded from font/th/Kanit/ subfolder
	pdf.SetFont("Kanit", "", 16)
	if pdf.Error() != nil {
		fmt.Printf("Note: External fonts not available: %v\n", pdf.Error())
		return
	}

	pdf.Cell(0, 10, "Auto-loaded from font/th/Kanit/")
	pdf.Ln(8)

	// Auto-load bold from font/th/Kanit/Kanit-Bold.ttf
	pdf.SetFont("Kanit", "B", 16)
	pdf.Cell(0, 10, "Bold auto-loaded too")

	err := pdf.OutputFileAndClose("example2_autoload.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✓ Created example2_autoload.pdf")
}

// example3_ThaiDocument สร้างไฟล์ PDF ชื่อ "example3_thai.pdf" ที่แสดงตัวอย่างการใช้งานฟอนต์ภาษาไทยแบบฝัง
// ฟังก์ชันจะเพิ่มหน้ากระดาษและเขียนหัวเรื่อง ข้อความยาวแบบหลายบรรทัด ข้อความผสมไทย-อังกฤษ ตัวเลขไทย และตัวอย่างสไตล์ตัวเอียง/ตัวหนา-เอียง โดยใช้ฟอนต์ที่ฝังมาในไฟล์ PDF และปิดไฟล์เมื่อเสร็จสิ้น.
func example3_ThaiDocument() {
	fmt.Println("\nExample 3: Thai Language Document")

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.UseEmbeddedFonts()
	pdf.AddPage()

	// Title in Thai
	pdf.SetFont("Kanit", "B", 20)
	pdf.Cell(0, 15, "เอกสารภาษาไทย")
	pdf.Ln(12)

	// Body text
	pdf.SetFont("Sarabun", "", 14)
	pdf.MultiCell(0, 8, "นี่คือตัวอย่างการใช้งานฟอนต์ภาษาไทยในไฟล์ PDF "+
		"โดยใช้ไลบรารี gofpdf ซึ่งรองรับการใช้งาน UTF-8 "+
		"และมีฟอนต์ภาษาไทยแบบฝังในตัว", "", "", false)
	pdf.Ln(8)

	// Mixed Thai and English
	pdf.SetFont("Prompt", "", 12)
	pdf.Cell(0, 10, "Mixed: Hello สวัสดี World โลก")
	pdf.Ln(6)

	// Thai numerals
	pdf.Cell(0, 10, "Thai numbers: ๑๒๓๔๕๖๗๘๙๐")
	pdf.Ln(6)

	// Different styles
	pdf.SetFont("Kanit", "I", 12)
	pdf.Cell(0, 10, "Italic: ตัวเอียง")
	pdf.Ln(6)

	pdf.SetFont("Kanit", "BI", 12)
	pdf.Cell(0, 10, "Bold Italic: ตัวหนาและเอียง")

	err := pdf.OutputFileAndClose("example3_thai.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✓ Created example3_thai.pdf")
}

// example4_MultipleFonts สร้างเอกสาร PDF ที่แสดงตัวอย่างฟอนต์หลายครอบครัว สไตล์ และขนาด
// ฟังก์ชันจะนำฟอนต์ที่ฝังมาใช้ แสดงรายการฟอนต์ที่มี ลองโหลดฟอนต์หลายแบบแล้วข้ามฟอนต์ที่ไม่พร้อมใช้งาน
// และบันทึกไฟล์เป็น example4_gallery.pdf.
func example4_MultipleFonts() {
	fmt.Println("\nExample 4: Multiple Fonts and Styles")

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.UseEmbeddedFonts()
	pdf.AddPage()

	// Get available font families
	families := gofpdf.GetEmbeddedFontFamilies()
	fmt.Printf("Available fonts: %v\n", families)

	pdf.SetFont("Kanit", "B", 18)
	pdf.Cell(0, 10, "Font Gallery")
	pdf.Ln(12)

	// Showcase different families
	fontExamples := []struct {
		family string
		style  string
		size   float64
		text   string
	}{
		{"Kanit", "", 14, "Kanit Regular - กานต์ ปกติ"},
		{"Kanit", "B", 14, "Kanit Bold - กานต์ หนา"},
		{"Kanit", "I", 14, "Kanit Italic - กานต์ เอียง"},
		{"Sarabun", "", 14, "Sarabun Regular - สารบุญ ปกติ"},
		{"Sarabun", "B", 14, "Sarabun Bold - สารบุญ หนา"},
		{"Prompt", "", 14, "Prompt Regular - พร้อมพ์ ปกติ"},
		{"Tahoma", "", 14, "Tahoma - ทาโฮมา"},
	}

	for _, font := range fontExamples {
		pdf.SetFont(font.family, font.style, font.size)
		if pdf.Error() != nil {
			// Skip if font not available
			fmt.Printf("Skipping %s %s\n", font.family, font.style)
			pdf.SetError(nil)
			continue
		}
		pdf.Cell(0, 10, font.text)
		pdf.Ln(8)
	}

	// Size variations
	pdf.Ln(5)
	pdf.SetFont("Kanit", "B", 14)
	pdf.Cell(0, 10, "Font Sizes:")
	pdf.Ln(10)

	sizes := []float64{8, 10, 12, 14, 16, 18, 20, 24}
	for _, size := range sizes {
		pdf.SetFont("Sarabun", "", size)
		pdf.Cell(0, size/2+2, fmt.Sprintf("Size %.0fpt - ขนาด %.0f", size, size))
		pdf.Ln(size/2 + 4)
	}

	err := pdf.OutputFileAndClose("example4_gallery.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✓ Created example4_gallery.pdf")
}

// example5_MixedFonts สาธิตการใช้ฟอนต์ core (ฝังใน PDF reader) ร่วมกับฟอนต์ embedded ในเอกสาร PDF เดียวกัน.
// สร้างหน้า PDF ที่แสดงข้อความตัวอย่างด้วยฟอนต์ core (เช่น Arial, Times) และฟอนต์ embedded (เช่น Kanit, Sarabun) แล้วบันทึกเป็นไฟล์ example5_mixed.pdf.
func example5_MixedFonts() {
	fmt.Println("\nExample 5: Mixed Core and Embedded Fonts")

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.UseEmbeddedFonts()
	pdf.AddPage()

	// Core PDF font (built into all PDF readers)
	pdf.SetFont("Arial", "", 14)
	pdf.Cell(0, 10, "Arial (Core PDF Font) - Latin only")
	pdf.Ln(8)

	// Embedded Thai font
	pdf.SetFont("Kanit", "", 14)
	pdf.Cell(0, 10, "Kanit (Embedded) - supports Thai: สวัสดี")
	pdf.Ln(8)

	// Core font again
	pdf.SetFont("Times", "B", 14)
	pdf.Cell(0, 10, "Times Bold (Core)")
	pdf.Ln(8)

	// Embedded font
	pdf.SetFont("Sarabun", "I", 14)
	pdf.Cell(0, 10, "Sarabun Italic (Embedded) - ทดสอบ")

	err := pdf.OutputFileAndClose("example5_mixed.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✓ Created example5_mixed.pdf")
}