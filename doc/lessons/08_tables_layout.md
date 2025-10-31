## Lesson 08 — Simple Tables and Layout

Use `CellFormat` and `MultiCell` to build tables. Manage X/Y positions.

```go
pdf := gofpdf.New("P", "mm", "A4", "")
pdf.AddPage()
pdf.SetFont("Helvetica", "", 11)

headers := []string{"ID", "Name", "Email"}
widths := []float64{20, 60, 100}

// Header row
for i, h := range headers {
    pdf.CellFormat(widths[i], 8, h, "1", 0, "C", true, 0, "")
}
pdf.Ln(-1)

rows := [][]string{{"1", "Alice", "alice@example.com"}, {"2", "Bob", "bob@example.com"}}
for _, r := range rows {
    for i, v := range r {
        pdf.CellFormat(widths[i], 8, v, "1", 0, "L", false, 0, "")
    }
    pdf.Ln(-1)
}

_ = pdf.OutputFileAndClose("lesson08.pdf")
```

Tips:
- `Ln(-1)` moves to next line using last cell height.
- Track `GetX()/GetY()` to create multi‑column layouts.


