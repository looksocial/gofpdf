## Lesson 03 — Your First PDF

Create a page, write centered text, and save.

```go
package main

import (
    gofpdf "github.com/looksocial/gofpdf"
)

func main() {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()

    pdf.SetFont("Helvetica", "B", 24)
    pdf.CellFormat(0, 12, "gofpdf quickstart", "", 1, "C", false, 0, "")

    pdf.SetFont("Helvetica", "", 14)
    pdf.Ln(10)
    pdf.MultiCell(0, 7, "This is your first PDF generated with gofpdf.", "", "L", false)

    _ = pdf.OutputFileAndClose("lesson03.pdf")
}
```

Notes:
- `CellFormat` with width 0 spans to right margin; alignment `C` centers text.
- `Ln(n)` moves the cursor down by `n` units.


