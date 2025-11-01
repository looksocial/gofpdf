# Table Component for gofpdf

An easy-to-use, customizable table component for generating professional PDF tables with gofpdf.

## Features

✅ **Simple API** - Define columns, pass data, and render  
✅ **Fully Customizable** - Colors, fonts, borders, alignment  
✅ **Column Spanning** - Merge cells across multiple columns  
✅ **Alternating Rows** - Built-in zebra striping  
✅ **Auto Layout** - Automatic column width calculation  
✅ **Type-Safe** - Works with any data type  

## Quick Start

```go
package main

import (
    "github.com/looksocial/gofpdf"
    "github.com/looksocial/gofpdf/table"
)

func main() {
    // Create PDF
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "", 12)

    // Define columns
    columns := []table.Column{
        {Key: "id", Label: "ID", Width: 20, Align: "L"},
        {Key: "name", Label: "Name", Width: 60, Align: "L"},
        {Key: "email", Label: "Email", Width: 110, Align: "L"},
    }

    // Create and render table
    tbl := table.NewTable(pdf, columns)
    
    data := []map[string]interface{}{
        {"id": "1", "name": "Alice", "email": "alice@example.com"},
        {"id": "2", "name": "Bob", "email": "bob@example.com"},
    }
    
    tbl.Render(true, data)
    pdf.OutputFileAndClose("output.pdf")
}
```

## API Reference

### Column Definition

```go
type Column struct {
    Key         string  // Field key for data access
    Label       string  // Header label
    Width       float64 // Width in mm (0 = auto)
    Align       string  // "L", "C", "R" for data cells
    HeaderAlign string  // "L", "C", "R" for header alignment (overrides Align for headers if set)
    ColSpan     int     // Number of columns to span
    MergeCell   bool    // Deprecated: use ColSpan
}
```

### Table Creation

```go
// Basic
tbl := table.NewTable(pdf, columns)

// With options
tbl := table.NewTable(pdf, columns).
    WithRowHeight(8).
    WithAlternatingRows(true).
    WithHeaderStyle(table.CellStyle{
        Border:    "1",
        Bold:      true,
        FillColor: []int{100, 100, 200},
        TextColor: []int{255, 255, 255},
    }).
    WithDataStyle(table.CellStyle{
        Border: "LR",
    })
```

### Styling

```go
// Cell style
cellStyle := table.CellStyle{
    Border:    "1",              // Border: "1", "LRT", "LR", etc.
    Bold:      true,             // Bold text
    Italic:    false,            // Italic text
    FontSize:  0,                // 0 = use default
    FillColor: []int{R, G, B},   // Background color
    TextColor: []int{R, G, B},   // Text color
    Align:     "C",              // "L", "C", "R" - overrides column alignment if set
}

// Apply to table
tbl.WithHeaderStyle(cellStyle)
tbl.WithDataStyle(cellStyle)
```

## Examples

See `examples/table_demo.go` for complete examples including:

1. **Simple Table** - Basic table with headers and data
2. **Styled Table** - Custom colors and fonts
3. **Alternating Rows** - Zebra-striped tables

## Advanced Features

### Column Spanning

```go
columns := []table.Column{
    {Key: "quarter", Label: "Q1 2024", Width: 40, ColSpan: 2},
    {Key: "q1", Label: "", Width: 0},  // Skipped
    {Key: "q2", Label: "Q2 2024", Width: 40, ColSpan: 2},
}
```

### Auto-Width Columns

```go
columns := []table.Column{
    {Key: "name", Label: "Name", Width: 0},      // Auto-calculated
    {Key: "email", Label: "Email", Width: 100},   // Fixed width
    {Key: "phone", Label: "Phone", Width: 0},     // Auto-calculated
}
```

### Positioning

```go
tbl := table.NewTable(pdf, columns).
    WithStartPosition(50, 100)  // X, Y position in mm
```

### Alignment

The table component supports flexible alignment options at multiple levels:

#### 1. Column-Level Alignment

```go
columns := []table.Column{
    {Key: "id", Label: "ID", Width: 30, Align: "L"},           // Data: Left
    {Key: "name", Label: "Name", Width: 50, Align: "C"},        // Data: Center
    {Key: "price", Label: "Price", Width: 40, Align: "R", HeaderAlign: "C"}, // Data: Right, Header: Center
}
```

#### 2. Header-Specific Alignment

Use `HeaderAlign` to set different alignment for headers:

```go
columns := []table.Column{
    {Key: "id", Label: "ID", Width: 30, Align: "L", HeaderAlign: "C"},
    {Key: "price", Label: "Price", Width: 40, Align: "R", HeaderAlign: "R"},
}
```

#### 3. Style-Level Alignment

Override alignment for all headers or all data cells:

```go
tbl := table.NewTable(pdf, columns).
    WithHeaderStyle(table.CellStyle{
        Align: "C",  // All headers centered
        Bold:  true,
    }).
    WithDataStyle(table.CellStyle{
        Align: "R",  // All data cells right-aligned (unless overridden per-row)
    })
```

#### 4. Per-Row Alignment

Override alignment for specific cells in a row using the `_align` map:

```go
data := []map[string]interface{}{
    {
        "id":    "1",
        "name":  "Product A",
        "price": "10.50",
        "_align": map[string]string{
            "price": "R",  // Right-align price for this row
            "name":  "C",  // Center-align name for this row
        },
    },
    {"id": "2", "name": "Product B", "price": "25.00"}, // Uses default alignment
}
```

**Alignment Priority (highest to lowest):**
- Per-row alignment (`_align` map in data)
- Style-level alignment (`CellStyle.Align` in `HeaderStyle` or `DataStyle`)
- Column-level alignment (`Column.Align` or `Column.HeaderAlign`)

## Chaining Methods

All configuration methods are chainable:

```go
tbl := table.NewTable(pdf, columns).
    WithHeaderStyle(headerStyle).
    WithDataStyle(dataStyle).
    WithRowHeight(8).
    WithAlternatingRows(true).
    WithRowSpacing(2)
```

## License

Same as gofpdf (MIT License).
