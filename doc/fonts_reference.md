# Font Reference Guide

Complete reference for font handling in gofpdf.

## Table of Contents

1. [Quick Start](#quick-start)
2. [Font Methods](#font-methods)
3. [Embedded Fonts](#embedded-fonts)
4. [External Fonts](#external-fonts)
5. [Font Styles](#font-styles)
6. [Auto-Loading](#auto-loading)
7. [Thai Fonts](#thai-fonts)
8. [Advanced Usage](#advanced-usage)
9. [Troubleshooting](#troubleshooting)

---

## Quick Start

### Simplest Way (Embedded Fonts)

```go
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.UseEmbeddedFonts()
pdf.AddPage()
pdf.SetFont("Kanit", "", 14)
pdf.Cell(0, 10, "Hello สวัสดี")
pdf.OutputFileAndClose("output.pdf")
```

---

## Font Methods

### UseEmbeddedFonts()

Enable embedded fonts bundled with the package.

```go
pdf.UseEmbeddedFonts()
```

**Benefits:**
- No external files needed
- Works on any platform
- Thai font families included
- Fonts embedded in your compiled binary

---

### SetFont(family, style, size)

Set the current font.

```go
pdf.SetFont("Kanit", "B", 14)
```

**Parameters:**
- `family` (string) - Font family name (e.g., "Kanit", "Arial")
- `style` (string) - Style: "", "B", "I", "BI"
- `size` (float64) - Font size in points

**Auto-loads font** if not yet registered (embedded or from filesystem).

**Example:**
```go
pdf.SetFont("Kanit", "", 12)    // Regular
pdf.SetFont("Kanit", "B", 12)   // Bold
pdf.SetFont("Kanit", "I", 12)   // Italic
pdf.SetFont("Kanit", "BI", 12)  // Bold Italic
```

---

### SetFontSize(size)

Change font size without changing family/style.

```go
pdf.SetFontSize(18)
```

**Example:**
```go
pdf.SetFont("Kanit", "", 12)
pdf.Cell(0, 10, "Size 12")
pdf.SetFontSize(20)
pdf.Cell(0, 12, "Size 20 - still Kanit")
```

---

### SetFontLocation(path)

Set directory for external font files.

```go
pdf.SetFontLocation("font/th")
```

**Usage:**
```go
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.SetFontLocation("font/th")
pdf.SetFont("Kanit", "", 14)  // Loads from font/th/Kanit/
```

---

### AddUTF8Font(family, style, file)

Manually add a TrueType font.

```go
pdf.AddUTF8Font("MyFont", "", "path/to/font.ttf")
```

**Parameters:**
- `family` - Name to use with SetFont()
- `style` - "", "B", "I", or "BI"
- `file` - Path to TTF/OTF file (relative to font location)

**Example:**
```go
pdf.SetFontLocation("fonts")
pdf.AddUTF8Font("Kanit", "", "Kanit-Regular.ttf")
pdf.AddUTF8Font("Kanit", "B", "Kanit-Bold.ttf")
pdf.SetFont("Kanit", "", 14)
```

---

### AddUTF8FontFromBytes(family, style, fontBytes)

Add font from byte array (embedded in your code).

```go
pdf.AddUTF8FontFromBytes("MyFont", "", fontData)
```

**Use case:** When you embed font files using `go:embed`.

---

### GetEmbeddedFontFamilies()

List all available embedded font families.

```go
families := gofpdf.GetEmbeddedFontFamilies()
// Returns: ["Kanit", "Sarabun", "Prompt", ...]
```

---

### GetFontSize()

Get current font size.

```go
ptSize, unitSize := pdf.GetFontSize()
```

**Returns:**
- `ptSize` - Size in points
- `unitSize` - Size in document units (for line height)

---

## Embedded Fonts

### What Are Embedded Fonts?

Fonts bundled into the compiled binary using Go's `embed` directive. No external files needed!

### How to Use

```go
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.UseEmbeddedFonts()  // Enable embedded fonts
pdf.AddPage()

// Use any embedded font
pdf.SetFont("Kanit", "", 14)
pdf.Cell(0, 10, "Works anywhere!")
```

### List Available Fonts

```go
families := gofpdf.GetEmbeddedFontFamilies()
for _, family := range families {
    fmt.Println(family)
}
```

### Embedded Font Families

**Complete list of bundled fonts:**

| Family | Weights Available | Italic Support |
|--------|------------------|----------------|
| Kanit | Thin to Black (9 weights) | ✅ Yes |
| Sarabun | Thin to ExtraBold (8 weights) | ✅ Yes |
| Prompt | Thin to Black (9 weights) | ✅ Yes |
| Taviraj | Thin to Black (9 weights) | ✅ Yes |
| Trirong | Thin to Black (9 weights) | ✅ Yes |
| Maitree | ExtraLight to Bold | ❌ No |
| Mitr | ExtraLight to Bold | ❌ No |
| Pridi | ExtraLight to Bold | ❌ No |
| Athiti | ExtraLight to Bold | ❌ No |
| NotoSansThai | Thin to Black | ❌ No |
| NotoSerifThai | Thin to Black | ❌ No |
| NotoSansThaiLooped | Thin to Black | ❌ No |
| Tahoma | Regular, Bold | ❌ No |
| Chonburi | Regular | ❌ No |
| Itim | Regular | ❌ No |
| Pattaya | Regular | ❌ No |
| Sriracha | Regular | ❌ No |

---

## External Fonts

### Directory Structure

Organize fonts by family in subfolders:

```
project/
├── font/
│   └── th/
│       ├── Kanit/
│       │   ├── Kanit-Regular.ttf
│       │   ├── Kanit-Bold.ttf
│       │   ├── Kanit-Italic.ttf
│       │   └── Kanit-BoldItalic.ttf
│       ├── Sarabun/
│       │   ├── Sarabun-Regular.ttf
│       │   └── Sarabun-Bold.ttf
│       └── MyCustomFont/
│           └── MyCustomFont-Regular.ttf
└── main.go
```

### Using External Fonts

**Method 1: Auto-load (Recommended)**

```go
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.SetFontLocation("font/th")
pdf.AddPage()

// Auto-loads from font/th/Kanit/ subfolder
pdf.SetFont("Kanit", "", 14)
```

**Method 2: Manual loading**

```go
pdf := gofpdf.New("P", "mm", "A4", "font/th/Kanit")
pdf.AddPage()

pdf.AddUTF8Font("Kanit", "", "Kanit-Regular.ttf")
pdf.AddUTF8Font("Kanit", "B", "Kanit-Bold.ttf")
pdf.SetFont("Kanit", "", 14)
```

---

## Font Styles

### Style Codes

| Code | Style | Example |
|------|-------|---------|
| `""` | Regular | `pdf.SetFont("Kanit", "", 14)` |
| `"B"` | Bold | `pdf.SetFont("Kanit", "B", 14)` |
| `"I"` | Italic | `pdf.SetFont("Kanit", "I", 14)` |
| `"BI"` | Bold Italic | `pdf.SetFont("Kanit", "BI", 14)` |
| `"IB"` | Bold Italic | Same as "BI" |

### Style Example

```go
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.UseEmbeddedFonts()
pdf.AddPage()

pdf.SetFont("Kanit", "", 14)
pdf.Cell(0, 10, "Regular: ปกติ")
pdf.Ln(8)

pdf.SetFont("Kanit", "B", 14)
pdf.Cell(0, 10, "Bold: ตัวหนา")
pdf.Ln(8)

pdf.SetFont("Kanit", "I", 14)
pdf.Cell(0, 10, "Italic: ตัวเอียง")
pdf.Ln(8)

pdf.SetFont("Kanit", "BI", 14)
pdf.Cell(0, 10, "Bold Italic: หนาและเอียง")
```

---

## Auto-Loading

### How It Works

When you call `SetFont("Kanit", "B", 14)`:

1. Checks if font already loaded → use it
2. Checks if it's a core font (Arial, Times, etc.) → load from built-in
3. If embedded fonts enabled → search embedded FS
4. Otherwise → search filesystem

### Search Locations

For `SetFont("Kanit", "B", 14)` with `SetFontLocation("font/th")`:

**Searches for (in order):**
1. `font/th/Kanit-Bold.ttf`
2. `font/th/KanitBold.ttf`
3. `font/th/KanitB.ttf`
4. `font/th/Kanit/Kanit-Bold.ttf`
5. `font/th/Kanit/KanitBold.ttf`
6. `font/th/Kanit/KanitB.ttf`

**For regular style** (`""`), also tries:
- `Family-Regular.ttf`
- `FamilyRegular.ttf`
- `Family.ttf`

### Naming Patterns

Auto-load recognizes these patterns:

**Regular:**
- `Kanit.ttf`
- `Kanit-Regular.ttf`
- `KanitRegular.ttf`

**Bold:**
- `Kanit-Bold.ttf`
- `KanitBold.ttf`
- `Kanit-Bd.ttf`
- `KanitB.ttf`

**Italic:**
- `Kanit-Italic.ttf`
- `KanitItalic.ttf`
- `Kanit-It.ttf`
- `KanitI.ttf`
- `Kanit-Oblique.ttf`

**BoldItalic:**
- `Kanit-BoldItalic.ttf`
- `KanitBoldItalic.ttf`
- `Kanit-BI.ttf`
- `Kanit-Bold-Italic.ttf`

---

## Thai Fonts

### Character Support

All embedded Thai fonts support:
- Thai consonants: ก ข ค ง จ ฉ...
- Thai vowels: ะ า ำ เ แ โ ใ ไ...
- Thai tone marks: ่ ้ ๊ ๋
- Thai numerals: ๐ ๑ ๒ ๓ ๔ ๕ ๖ ๗ ๘ ๙
- Latin alphabet: A-Z, a-z
- Numbers: 0-9
- Common symbols

### Best Thai Fonts by Use Case

**Modern Documents:**
- **Kanit** - Clean, modern sans-serif
- **Sarabun** - Professional sans-serif
- **Prompt** - Friendly rounded sans-serif

**Traditional/Formal:**
- **Taviraj** - Classic serif
- **Trirong** - Looped serif style

**Web-safe:**
- **Tahoma** - Windows system font
- **NotoSansThai** - Google's universal font

### Thai Text Example

```go
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.UseEmbeddedFonts()
pdf.AddPage()

pdf.SetFont("Kanit", "B", 18)
pdf.Cell(0, 12, "หัวข้อเอกสาร")
pdf.Ln(10)

pdf.SetFont("Sarabun", "", 14)
text := `วรรค์หนึ่ง: นี่คือตัวอย่างข้อความภาษาไทย
ที่สามารถแสดงผลได้อย่างสวยงาม รองรับทั้ง
สระ วรรณยุกต์ และตัวเลขไทย ๑๒๓๔๕`

pdf.MultiCell(0, 8, text, "", "", false)
```

---

## Advanced Usage

### Mixing Multiple Fonts

```go
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.UseEmbeddedFonts()
pdf.AddPage()

// Title in Kanit
pdf.SetFont("Kanit", "B", 20)
pdf.Cell(0, 15, "หัวเรื่อง")
pdf.Ln(12)

// Body in Sarabun
pdf.SetFont("Sarabun", "", 14)
pdf.MultiCell(0, 8, "เนื้อหาเอกสาร...", "", "", false)
pdf.Ln(8)

// Footer in Prompt
pdf.SetFont("Prompt", "I", 10)
pdf.Cell(0, 6, "หมายเหตุ")
```

### Dynamic Font Selection

```go
fonts := []string{"Kanit", "Sarabun", "Prompt"}

for i, fontName := range fonts {
    pdf.SetFont(fontName, "", 14)
    pdf.Cell(0, 10, fmt.Sprintf("%d. Font: %s", i+1, fontName))
    pdf.Ln(8)
}
```

### Font with Color

```go
pdf.SetFont("Kanit", "B", 16)
pdf.SetTextColor(255, 0, 0)  // Red
pdf.Cell(0, 10, "Red text in Kanit Bold")

pdf.SetTextColor(0, 0, 255)  // Blue
pdf.Ln(8)
pdf.Cell(0, 10, "Blue text in Kanit Bold")

pdf.SetTextColor(0, 0, 0)    // Back to black
```

### Responsive Font Sizes

```go
// Page title
pdf.SetFont("Kanit", "B", 24)
pdf.Cell(0, 15, "Title")
pdf.Ln(12)

// Section heading
pdf.SetFont("Kanit", "B", 18)
pdf.Cell(0, 12, "Section")
pdf.Ln(10)

// Body text
pdf.SetFont("Sarabun", "", 14)
pdf.MultiCell(0, 8, "Body text...", "", "", false)
pdf.Ln(6)

// Footnote
pdf.SetFont("Sarabun", "I", 10)
pdf.Cell(0, 6, "* Footnote")
```

---

## Troubleshooting

### Font not found error

**Error:** `undefined font: kanit`

**Solutions:**

1. **Use embedded fonts:**
```go
pdf.UseEmbeddedFonts()  // Add this line
```

2. **Check font location:**
```go
pdf.SetFontLocation("font/th")  // Correct path
```

3. **Manually add font:**
```go
pdf.AddUTF8Font("Kanit", "", "font/th/Kanit/Kanit-Regular.ttf")
```

4. **Check font exists:**
```go
families := gofpdf.GetEmbeddedFontFamilies()
fmt.Println(families)  // See what's available
```

---

### Thai characters show as boxes

**Cause:** Using core fonts (Arial, Times) which don't support Thai.

**Solution:** Use Thai-compatible fonts:
```go
// ❌ Won't work for Thai
pdf.SetFont("Arial", "", 14)
pdf.Cell(0, 10, "สวัสดี")  // Shows boxes

// ✅ Works with Thai
pdf.UseEmbeddedFonts()
pdf.SetFont("Kanit", "", 14)
pdf.Cell(0, 10, "สวัสดี")  // Shows correctly
```

---

### Bold/Italic not working

**Cause:** Font file for that style doesn't exist.

**Check available styles:**
```go
pdf.SetFont("Kanit", "B", 14)
if pdf.Error() != nil {
    fmt.Println("Bold not available:", pdf.Error())
}
```

**Solution:** Use a font with that style, or use only available styles.

---

### Font file not found (external fonts)

**Error:** `font file not found: font/Kanit.ttf`

**Solutions:**

1. **Check path is correct:**
```go
// If running from project root:
pdf.SetFontLocation("font/th")

// If running from subdirectory:
pdf.SetFontLocation("../font/th")
```

2. **Use absolute path:**
```go
pdf.SetFontLocation("/full/path/to/fonts")
```

3. **Use embedded fonts instead:**
```go
pdf.UseEmbeddedFonts()  // No path issues!
```

---

### Large PDF file size

**Cause:** Multiple large fonts embedded.

**Solutions:**

1. **Use fewer fonts** - Each font adds ~200KB-800KB
2. **Font subsetting works automatically** - Only used characters embedded
3. **Pre-subset fonts** if needed for smaller size

---

## See Also

- **Lesson 05** - Fonts and UTF-8 tutorial
- **example_fonts.go** - Complete working examples
- **fonts_test.go** - Test cases showing all features

