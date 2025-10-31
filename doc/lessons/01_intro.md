## Lesson 01 — Introduction to gofpdf

**Note**: This package (`gofpdf`) is a fork of [gofpdf](https://github.com/jung-kurt/gofpdf), renamed from `gofpdf` to `gofpdf` with **added Thai font support**. All the original functionality of gofpdf is preserved.

gofpdf is a pure Go PDF generator. It does not require external C libs and works cross‑platform. You build a `*gofpdf.Fpdf`, add pages, draw text/shapes/images, and output to file or bytes. This fork includes 18+ embedded Thai font families for seamless Thai language support.

### Install
```bash
go get github.com/looksocial/gofpdf
```

### Minimal example
```go
package main

import (
    gofpdf "github.com/looksocial/gofpdf"
)

func main() {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Helvetica", "", 16)
    pdf.Cell(40, 10, "Hello, PDF!")
    _ = pdf.OutputFileAndClose("hello.pdf")
}
```

### Key concepts
- Page units: mm, pt, in. Choose when calling `New`.
- Coordinate origin: top‑left of the page.
- State: font, color, line width persist until changed.
- Output: write to file with `OutputFileAndClose` or to `io.Writer` via `Output`.


