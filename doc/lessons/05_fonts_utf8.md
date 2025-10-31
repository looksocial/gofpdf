## Lesson 05 — Fonts and UTF‑8

### Core Fonts vs. Unicode Fonts

**Core fonts** (Helvetica, Times, Courier, ZapfDingbats) are built into PDF readers but only support Latin characters.

For **UTF‑8 text** (Thai, Chinese, Russian, etc.), you need Unicode TrueType fonts.

---

## Method 1: Embedded Fonts (Recommended) ⭐

The simplest way — fonts are bundled with the package!

```go
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.UseEmbeddedFonts()  // Enable embedded fonts
pdf.AddPage()

// Use Thai fonts directly - no setup needed!
pdf.SetFont("Kanit", "", 14)
pdf.Cell(0, 10, "Thai: สวัสดีครับ")
pdf.Ln(8)

pdf.SetFont("Sarabun", "B", 14)
pdf.Cell(0, 10, "Bold: ตัวหนา")
pdf.Ln(8)

pdf.SetFont("Prompt", "I", 14)
pdf.Cell(0, 10, "Italic: ตัวเอียง")

_ = pdf.OutputFileAndClose("embedded.pdf")
```

**Benefits:**
- ✅ No external files needed
- ✅ Works on any computer
- ✅ Thai font families included
- ✅ Auto-loads on first use

### Available Embedded Fonts

Call `GetEmbeddedFontFamilies()` to list all available fonts:

```go
families := gofpdf.GetEmbeddedFontFamilies()
// Returns: ["Kanit", "Sarabun", "Prompt", "Tahoma", ...]
```

**Thai Font Families:**
- **Kanit** - Modern sans-serif (Thin to Black + Italics)
- **Sarabun** - Clean sans-serif (Thin to ExtraBold + Italics)
- **Prompt** - Rounded sans-serif (Thin to Black + Italics)
- **Taviraj** - Serif font (Thin to Black + Italics)
- **Trirong** - Looped serif (Thin to Black + Italics)
- **Maitree**, **Mitr**, **Pridi** - Various weights
- **NotoSansThai**, **NotoSerifThai** - Google Noto fonts
- **Tahoma** - Windows font (Regular, Bold)
- And more...

---

## Method 2: External Font Files with Auto-Load

Load fonts from your filesystem automatically:

```go
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.SetFontLocation("font/th")  // Point to font directory
pdf.AddPage()

// Auto-loads from font/th/Kanit/ subfolder
pdf.SetFont("Kanit", "", 14)
pdf.Cell(0, 10, "Auto-loaded from filesystem")

_ = pdf.OutputFileAndClose("external.pdf")
```

**Directory structure:**
```
font/th/
  ├── Kanit/
  │   ├── Kanit-Regular.ttf
  │   ├── Kanit-Bold.ttf
  │   └── Kanit-Italic.ttf
  ├── Sarabun/
  │   └── Sarabun-Regular.ttf
  └── ...
```

---

## Method 3: Traditional Manual Loading

Explicitly add fonts before use:

```go
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.AddPage()

// Manually add TTF
pdf.AddUTF8Font("DejaVu", "", "font/DejaVuSansCondensed.ttf")
pdf.SetFont("DejaVu", "", 14)
pdf.Cell(0, 10, "English: Hello")
pdf.Ln(8)
pdf.Cell(0, 10, "ไทย: สวัสดี")
pdf.Ln(8)
pdf.Cell(0, 10, "Русский: Привет")

_ = pdf.OutputFileAndClose("manual.pdf")
```

---

## Font Styles

Use the second parameter of `SetFont()` for styles:

```go
pdf.SetFont("Kanit", "", 14)   // Regular
pdf.SetFont("Kanit", "B", 14)  // Bold
pdf.SetFont("Kanit", "I", 14)  // Italic
pdf.SetFont("Kanit", "BI", 14) // Bold + Italic
```

**Style codes:**
- `""` - Regular
- `"B"` - Bold
- `"I"` - Italic
- `"BI"` or `"IB"` - Bold Italic

---

## Complete Thai Document Example

```go
package main

import "github.com/looksocial/gofpdf"

func main() {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.UseEmbeddedFonts()
    pdf.AddPage()

    // Title
    pdf.SetFont("Kanit", "B", 20)
    pdf.Cell(0, 15, "เอกสารภาษาไทย")
    pdf.Ln(12)

    // Body
    pdf.SetFont("Sarabun", "", 14)
    pdf.MultiCell(0, 8, 
        "นี่คือตัวอย่างการสร้างเอกสาร PDF ภาษาไทย "+
        "โดยใช้ไลบรารี gofpdf ซึ่งรองรับ UTF-8 "+
        "และมีฟอนต์ไทยแบบฝังในตัว",
        "", "", false)
    pdf.Ln(8)

    // Mixed language
    pdf.SetFont("Prompt", "", 12)
    pdf.Cell(0, 10, "Mixed: Hello สวัสดี 你好")
    pdf.Ln(6)

    // Thai numerals
    pdf.Cell(0, 10, "Numbers: ๑๒๓๔๕ = 12345")

    _ = pdf.OutputFileAndClose("thai_doc.pdf")
}
```

---

## Font Sizes

Third parameter sets size in points:

```go
pdf.SetFont("Kanit", "", 8)   // Small
pdf.SetFont("Kanit", "", 12)  // Normal
pdf.SetFont("Kanit", "", 16)  // Large
pdf.SetFont("Kanit", "", 24)  // Extra large
```

Or use `SetFontSize()` to change size without changing family/style:

```go
pdf.SetFont("Kanit", "", 14)
pdf.Cell(0, 10, "Size 14")
pdf.Ln(8)

pdf.SetFontSize(20)
pdf.Cell(0, 12, "Size 20 - still Kanit")
```

---

## Mixing Core and Thai Fonts

```go
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.UseEmbeddedFonts()
pdf.AddPage()

// Core font (Latin only)
pdf.SetFont("Arial", "", 14)
pdf.Cell(0, 10, "Arial: Latin characters only")
pdf.Ln(8)

// Thai font (supports Thai + Latin)
pdf.SetFont("Kanit", "", 14)
pdf.Cell(0, 10, "Kanit: Supports both Latin and ไทย")
```

---

## Notes

- **No preprocessing needed** - TTF files load directly
- **No `.json` or `.z` files** required for TTF/OTF fonts
- **Font subsetting** - Only used characters are embedded in PDF
- **Case insensitive** - `"Kanit"` and `"kanit"` both work
- **Auto-load patterns** - Tries `Family-Bold.ttf`, `FamilyBold.ttf`, `FamilyB.ttf`, etc.

---

## Next Steps

- **Lesson 06:** Images
- **Lesson 07:** Shapes and Colors

For more font options, see:
- `GetEmbeddedFontFamilies()` - List all available fonts
- `SetFontLocation()` - Use custom font directories
- `AddUTF8Font()` - Add specific font files
