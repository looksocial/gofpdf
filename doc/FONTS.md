# Font Documentation Index

Complete guide to using fonts in gofpdf.

## Documentation Files

### 📘 [Font Quick Reference](fonts_quickref.md)
**One-page cheat sheet** - Quick lookup for common font operations.
- All core methods in one place
- Common patterns and examples
- Popular Thai fonts table
- Troubleshooting guide

**Start here if you:** Need a quick reference while coding

---

### 📕 [Font Reference Guide](fonts_reference.md)
**Complete API documentation** - Detailed reference for all font features.
- Every font method explained
- Embedded fonts reference
- External font management
- Auto-loading mechanics
- Thai font catalog
- Advanced usage patterns
- Comprehensive troubleshooting

**Start here if you:** Want to understand all font features in depth

---

### 📗 [Lesson 05: Fonts and UTF-8](lessons/05_fonts_utf8.md)
**Step-by-step tutorial** - Learn fonts from the ground up.
- Three methods to use fonts
- Embedded fonts walkthrough
- Thai document examples
- Font styles and sizes
- Mixed language support

**Start here if you:** Are new to the package or prefer learning by example

---

## Quick Start

### Simplest Way (Embedded Fonts)

```go
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.UseEmbeddedFonts()  // ← Add this one line
pdf.AddPage()
pdf.SetFont("Kanit", "", 14)
pdf.Cell(0, 10, "สวัสดี - Hello")
pdf.OutputFileAndClose("output.pdf")
```

## Available Resources

### Code Examples
- **[example_fonts.go](../example_fonts.go)** - 5 complete working examples
  - Embedded fonts usage
  - External fonts with auto-load
  - Thai language document
  - Font gallery
  - Mixed core and Thai fonts

### Tests
- **[fonts_test.go](../fonts_test.go)** - Comprehensive test suite
  - 20+ test functions
  - All features covered
  - Performance benchmarks

### Main Documentation
- **[document.md](document.md)** - Main package documentation with font section

## Font Features at a Glance

### ✅ What's Included

| Feature | Description |
|---------|-------------|
| **18+ Thai Fonts** | Embedded in binary, zero setup |
| **Auto-loading** | Fonts load on first `SetFont()` call |
| **UTF-8 Support** | Full Unicode including Thai, Chinese, Arabic |
| **TTF/OTF Direct** | No preprocessing or conversion needed |
| **Subfolder Search** | Organize fonts by family |
| **Style Variants** | Regular, Bold, Italic, BoldItalic |
| **Zero Config** | Just call `UseEmbeddedFonts()` |

### 📦 Embedded Font Families

**Modern Sans-Serif:**
- Kanit (9 weights + italics)
- Sarabun (8 weights + italics)
- Prompt (9 weights + italics)

**Serif:**
- Taviraj (9 weights + italics)
- Trirong (9 weights + italics)

**Web-Safe:**
- Tahoma (Regular, Bold)
- NotoSansThai, NotoSerifThai

**And more:** Maitree, Mitr, Pridi, Athiti, Chonburi, Itim, Pattaya, Sriracha...

## Common Tasks

### Use Embedded Fonts
```go
pdf.UseEmbeddedFonts()
pdf.SetFont("Kanit", "", 14)
```
📖 See: [Quick Reference](fonts_quickref.md#quick-start)

### Use External Fonts
```go
pdf.SetFontLocation("font/th")
pdf.SetFont("Kanit", "", 14)
```
📖 See: [Reference Guide - External Fonts](fonts_reference.md#external-fonts)

### List Available Fonts
```go
families := gofpdf.GetEmbeddedFontFamilies()
```
📖 See: [Reference Guide - GetEmbeddedFontFamilies](fonts_reference.md#getembeddedfontfamilies)

### Thai Language Document
```go
pdf.UseEmbeddedFonts()
pdf.SetFont("Sarabun", "", 14)
pdf.MultiCell(0, 8, "ข้อความภาษาไทย", "", "", false)
```
📖 See: [Lesson 05 - Thai Document](lessons/05_fonts_utf8.md#complete-thai-document-example)

### Font Styles
```go
pdf.SetFont("Kanit", "", 14)   // Regular
pdf.SetFont("Kanit", "B", 14)  // Bold
pdf.SetFont("Kanit", "I", 14)  // Italic
pdf.SetFont("Kanit", "BI", 14) // Bold Italic
```
📖 See: [Quick Reference - Font Styles](fonts_quickref.md#font-styles)

## Learning Path

### 🚀 Beginner
1. Read [Lesson 05](lessons/05_fonts_utf8.md) - Learn the basics
2. Try [example_fonts.go](../example_fonts.go) - Run working examples
3. Keep [Quick Reference](fonts_quickref.md) handy while coding

### 🎯 Intermediate
1. Review [Reference Guide](fonts_reference.md) - Understand all features
2. Explore auto-loading and external fonts
3. Study [fonts_test.go](../fonts_test.go) for advanced patterns

### 🏆 Advanced
1. Master font organization and custom fonts
2. Optimize PDF sizes with selective font usage
3. Create multi-language documents

## Troubleshooting

### Font not found
```
Error: undefined font: kanit
```
**Solution:** Enable embedded fonts
```go
pdf.UseEmbeddedFonts()  // ← Add this
```
📖 See: [Reference Guide - Troubleshooting](fonts_reference.md#font-not-found-error)

### Thai shows as boxes
**Problem:** Using Arial/Times which don't support Thai

**Solution:** Use Thai-compatible font
```go
pdf.SetFont("Kanit", "", 14)  // ← Use Thai font
```
📖 See: [Reference Guide - Thai characters show as boxes](fonts_reference.md#thai-characters-show-as-boxes)

### Bold/Italic not working
**Problem:** Font file for that style doesn't exist

**Solution:** Check available styles or use embedded fonts
```go
pdf.UseEmbeddedFonts()  // Has all styles
```
📖 See: [Reference Guide - Bold/Italic not working](fonts_reference.md#bolditalic-not-working)

## API Quick Lookup

| Want to... | Method | Link |
|------------|--------|------|
| Enable embedded fonts | `UseEmbeddedFonts()` | [Ref](fonts_reference.md#useembeddedfonts) |
| Set current font | `SetFont(family, style, size)` | [Ref](fonts_reference.md#setfont) |
| Change size only | `SetFontSize(size)` | [Ref](fonts_reference.md#setfontsize) |
| Set font directory | `SetFontLocation(path)` | [Ref](fonts_reference.md#setfontlocation) |
| Add TTF manually | `AddUTF8Font(family, style, file)` | [Ref](fonts_reference.md#addutf8font) |
| List embedded fonts | `GetEmbeddedFontFamilies()` | [Ref](fonts_reference.md#getembeddedfontfamilies) |
| Get current size | `GetFontSize()` | [Ref](fonts_reference.md#getfontsize) |

## Need Help?

1. **Check the docs:**
   - [Quick Reference](fonts_quickref.md) for quick answers
   - [Reference Guide](fonts_reference.md) for detailed info
   - [Lesson 05](lessons/05_fonts_utf8.md) for tutorials

2. **Run examples:**
   ```bash
   go run example_fonts.go
   ```

3. **Run tests:**
   ```bash
   go test -v -run Font
   ```

## Summary

gofpdf makes using fonts **simple**:
- ✅ **Zero config** with embedded fonts
- ✅ **18+ Thai families** built-in
- ✅ **Auto-loads** on first use
- ✅ **UTF-8 ready** for any language
- ✅ **No preprocessing** needed

Just one line to get started:
```go
pdf.UseEmbeddedFonts()
```

Happy coding! 🎉

