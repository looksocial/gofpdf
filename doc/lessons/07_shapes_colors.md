## Lesson 07 — Shapes and Colors

Draw lines, rectangles, circles, and polygons. Set draw/fill colors.

```go
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.AddPage()

// Colors (RGB 0..255)
pdf.SetDrawColor(0, 0, 0)
pdf.SetFillColor(220, 235, 255)
pdf.SetLineWidth(0.5)

// Rectangle: style "D" draw, "F" fill, "DF" draw+fill
pdf.Rect(10, 10, 60, 30, "DF")

// Line
pdf.Line(10, 50, 100, 50)

// Circle
pdf.Circle(40, 80, 15, "D")

// Polygon
pts := []gofpdf.PointType{{10, 110}, {40, 130}, {15, 150}}
pdf.Polygon(pts, "DF")

_ = pdf.OutputFileAndClose("lesson07.pdf")
```

Tip: Use `TransformBegin/End` for rotations/scales (see advanced lesson).


