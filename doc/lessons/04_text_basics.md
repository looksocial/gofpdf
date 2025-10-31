## Lesson 04 — Text Basics

Common text APIs: `SetFont`, `Cell`, `CellFormat`, `MultiCell`, `Write`, `Ln`.

```go
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.AddPage()
pdf.SetFont("Helvetica", "", 12)

// Single line cell
pdf.Cell(40, 10, "Left label:")
pdf.Cell(0, 10, "Value")
pdf.Ln(12)

// B/I/U styles
pdf.SetFont("Helvetica", "B", 12)
pdf.Cell(0, 8, "Bold text")
pdf.Ln(8)
pdf.SetFont("Helvetica", "I", 12)
pdf.Cell(0, 8, "Italic text")
pdf.Ln(8)
pdf.SetFont("Helvetica", "U", 12)
pdf.Cell(0, 8, "Underlined text")
pdf.Ln(10)

// MultiCell wraps text
pdf.SetFont("Helvetica", "", 12)
long := "MultiCell wraps long text across lines and can justify or align left/center/right."
pdf.MultiCell(0, 6, long, "1", "J", false)

_ = pdf.OutputFileAndClose("lesson04.pdf")
```

Tips:
- Borders: pass `"1"` for box, or a combo like `"LRT"`.
- Alignment: `L`, `C`, `R`, `J` (justify).


