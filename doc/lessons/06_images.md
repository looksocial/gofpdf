## Lesson 06 — Images

Supported: PNG, JPEG, GIF, SVG (via helpers), and more via contrib packages.

```go
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.AddPage()

pdf.ImageOptions(
    "image/logo.png",
    10, 10, 50, 0, // x, y, width, height (0 keeps aspect)
    false,
    gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
    0, "",
)

pdf.ImageOptions(
    "image/logo.jpg",
    10, 70, 60, 0,
    false,
    gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
    0, "",
)

_ = pdf.OutputFileAndClose("lesson06.pdf")
```

Tips:
- Set height to 0 to keep aspect ratio.
- Use `ReadDpi: true` to respect embedded DPI.


