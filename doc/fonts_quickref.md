# Font Quick Reference

One-page cheat sheet for gofpdf fonts.

## Quick Start

```go
// Simplest way - embedded fonts
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.UseEmbeddedFonts()
pdf.AddPage()
pdf.SetFont("Kanit", "", 14)
pdf.Cell(0, 10, "Hello สวัสดี")
pdf.OutputFileAndClose("output.pdf")
```

## Core Methods

| Method | Usage | Description |
|--------|-------|-------------|
| `UseEmbeddedFonts()` | `pdf.UseEmbeddedFonts()` | Enable bundled fonts |
| `SetFont(family, style, size)` | `pdf.SetFont("Kanit", "B", 14)` | Set current font |
| `SetFontSize(size)` | `pdf.SetFontSize(18)` | Change size only |
| `SetFontLocation(path)` | `pdf.SetFontLocation("font/th")` | Set font directory |
| `AddUTF8Font(family, style, file)` | `pdf.AddUTF8Font("Kanit", "", "Kanit.ttf")` | Add TTF manually |
| `GetEmbeddedFontFamilies()` | `families := gofpdf.GetEmbeddedFontFamilies()` | List available fonts |

## Font Styles

| Code | Style | Example |
|------|-------|---------|
| `""` | Regular | `pdf.SetFont("Kanit", "", 14)` |
| `"B"` | Bold | `pdf.SetFont("Kanit", "B", 14)` |
| `"I"` | Italic | `pdf.SetFont("Kanit", "I", 14)` |
| `"BI"` | Bold Italic | `pdf.SetFont("Kanit", "BI", 14)` |

## Popular Thai Fonts

| Font | Best For | Weights | Italic |
|------|----------|---------|--------|
| **Kanit** | Modern docs | Thin → Black (9) | ✅ |
| **Sarabun** | Professional | Thin → ExtraBold (8) | ✅ |
| **Prompt** | Friendly | Thin → Black (9) | ✅ |
| **Taviraj** | Formal | Thin → Black (9) | ✅ |
| **Tahoma** | Windows-like | Regular, Bold | ❌ |
| **NotoSansThai** | Web-safe | Thin → Black | ❌ |

## Three Ways to Use Fonts

### 1. Embedded (Recommended) ⭐

```go
pdf.UseEmbeddedFonts()
pdf.SetFont("Kanit", "", 14)
```

### 2. Auto-Load from Filesystem

```go
pdf.SetFontLocation("font/th")
pdf.SetFont("Kanit", "", 14)  // Loads from font/th/Kanit/
```

### 3. Manual Loading

```go
pdf.AddUTF8Font("Kanit", "", "font/Kanit-Regular.ttf")
pdf.SetFont("Kanit", "", 14)
```

## Common Patterns

### Thai Document

```go
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.UseEmbeddedFonts()
pdf.AddPage()

pdf.SetFont("Kanit", "B", 20)
pdf.Cell(0, 15, "หัวข้อ")
pdf.Ln(12)

pdf.SetFont("Sarabun", "", 14)
pdf.MultiCell(0, 8, "เนื้อหา...", "", "", false)
```

### Mixed Languages

```go
pdf.SetFont("Sarabun", "", 14)
pdf.Cell(0, 10, "Thai: สวัสดี")
pdf.Ln(6)
pdf.Cell(0, 10, "English: Hello")
pdf.Ln(6)
pdf.Cell(0, 10, "Mixed: Hello สวัสดี 你好")
```

### Multiple Styles

```go
pdf.SetFont("Kanit", "", 14)
pdf.Cell(40, 10, "Regular ")

pdf.SetFont("Kanit", "B", 14)
pdf.Cell(40, 10, "Bold ")

pdf.SetFont("Kanit", "I", 14)
pdf.Cell(40, 10, "Italic")
```

### Font Sizes

```go
sizes := []float64{8, 10, 12, 14, 16, 18, 20, 24}
for _, size := range sizes {
    pdf.SetFont("Sarabun", "", size)
    pdf.Cell(0, size/2+2, fmt.Sprintf("Size %.0f", size))
    pdf.Ln(size/2 + 4)
}
```

## File Organization

```
project/
├── font/th/              ← Set with SetFontLocation()
│   ├── Kanit/
│   │   ├── Kanit-Regular.ttf
│   │   ├── Kanit-Bold.ttf
│   │   ├── Kanit-Italic.ttf
│   │   └── Kanit-BoldItalic.ttf
│   └── Sarabun/
│       └── Sarabun-Regular.ttf
└── main.go
```

## Troubleshooting

| Problem | Solution |
|---------|----------|
| Font not found | Add `pdf.UseEmbeddedFonts()` |
| Thai shows boxes | Use Thai font, not Arial/Times |
| Bold doesn't work | Check if bold file exists |
| Wrong path | Use `SetFontLocation()` or absolute path |

## Font Sizing

| Size (pt) | Use Case |
|-----------|----------|
| 8-10 | Footnotes, captions |
| 12-14 | Body text |
| 16-18 | Headings |
| 20-24 | Titles |
| 28+ | Display text |

## Performance Tips

- ✅ Fonts load once, reused automatically
- ✅ Only used characters embedded (subsetting)
- ✅ Embedded ≈ Same speed as filesystem
- ⚠️ Each font adds ~200-800KB to PDF
- 💡 Use fewer fonts for smaller PDFs

## Complete Example

```go
package main

import "github.com/looksocial/gofpdf"

func main() {
    // Setup
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.UseEmbeddedFonts()
    pdf.AddPage()
    
    // Title
    pdf.SetFont("Kanit", "B", 24)
    pdf.SetTextColor(0, 0, 128)
    pdf.Cell(0, 15, "เอกสารภาษาไทย")
    pdf.Ln(15)
    pdf.SetTextColor(0, 0, 0)
    
    // Subtitle
    pdf.SetFont("Sarabun", "B", 16)
    pdf.Cell(0, 10, "บทนำ")
    pdf.Ln(10)
    
    // Body
    pdf.SetFont("Sarabun", "", 14)
    pdf.MultiCell(0, 8, 
        "นี่คือตัวอย่างเอกสาร PDF ที่ใช้ฟอนต์ภาษาไทย "+
        "รองรับการแสดงผลภาษาไทยได้อย่างสมบูรณ์",
        "", "", false)
    pdf.Ln(8)
    
    // List
    items := []string{
        "รายการที่ 1",
        "รายการที่ 2",
        "รายการที่ 3",
    }
    
    for _, item := range items {
        pdf.Cell(10, 8, "•")
        pdf.Cell(0, 8, item)
        pdf.Ln(8)
    }
    
    // Save
    pdf.OutputFileAndClose("output.pdf")
}
```

## See Also

- `doc/lessons/05_fonts_utf8.md` - Step-by-step tutorial
- `doc/fonts_reference.md` - Complete API reference
- `example_fonts.go` - Working code examples
- `fonts_test.go` - Test cases

