## Lesson 02 — Project Setup

### Dependency
Add the library to your project:
```bash
go get github.com/looksocial/gofpdf
```

### Folder tips
- Keep assets (images, fonts) in `assets/` and reference by path.
- For custom TTF/OTF fonts, consider embedding or using makefont (see later lesson).

### Minimal runnable file
```go
package main

import (
    gofpdf "github.com/looksocial/gofpdf"
)

func main() {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Helvetica", "", 14)
    pdf.Cell(0, 10, "Setup OK")
    _ = pdf.OutputFileAndClose("setup.pdf")
}
```

Run with `go run .` in the directory containing the file.


