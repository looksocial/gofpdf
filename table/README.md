# Table Component for gofpdf

An easy-to-use, customizable table component for generating professional PDF tables with gofpdf.

## Features

✅ **Simple API** - Define columns, pass data, and render
✅ **Fully Customizable** - Colors, fonts, borders, alignment
✅ **Column Spanning** - Merge cells across multiple columns
✅ **Row Spanning** - Cells can span multiple rows
✅ **Nested Tables** - Tables within table cells
✅ **Alternating Rows** - Built-in zebra striping
✅ **Auto Layout** - Automatic column width calculation
✅ **Page Breaks** - Automatic page breaks with header repetition
✅ **Text Wrapping** - Automatic text wrapping in cells
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

The table component is used in various real-world scenarios. Below are complete working examples with code and visual outputs:

### Basic Table Examples

**Location:** `internal/example/demos/tabledemo/main.go`

This demo includes:
1. **Simple Table** - Basic table with headers and data
2. **Styled Table** - Custom colors, fonts, and borders
3. **Alternating Rows** - Zebra-striped tables with built-in alternating colors
4. **Header Alignment** - Different header and data alignment using `HeaderAlign`
5. **Style-Level Alignment** - Override alignment at the style level
6. **Per-Row Alignment** - Custom alignment for specific cells using `_align` map
7. **Mixed Alignments** - Combining all alignment features
8. **Column Spanning** - Headers spanning multiple columns
9. **Auto-Width Columns** - Automatic column width calculation
10. **Custom Positioning** - Positioning tables at specific coordinates

Run the demo:
```bash
cd internal/example/demos/tabledemo
go run main.go
```

### Nested Tables

**Location:** `internal/example/demos/nested/main.go`

Demonstrates advanced table features:
- **Simple Nested Tables** - Tables within table cells
- **Row Spanning** - Cells spanning multiple rows using `_rowspan`
- **Summary/Total Rows** - Adding summary and total rows with `AddSummaryRow()` and `AddTotalRow()`
- **Complex Nested Structures** - Nested tables combined with row spans
- **Text Wrapping** - Automatic text wrapping in nested cells

Run the demo:
```bash
cd internal/example/demos/nested
go run main.go
```

### Multi-Page Tables with Page Breaks

**Location:** `internal/example/demos/pagebreak/main.go`

Shows automatic page break handling:
- **Header Repetition** - Headers automatically repeated on each new page
- **Custom Page Break Margins** - Adjustable margin from bottom before page break
- **Large Datasets** - Handling hundreds of rows across multiple pages

Run the demo:
```bash
cd internal/example/demos/pagebreak
go run main.go
```

The demo generates a multi-page PDF demonstrating automatic page breaks with header repetition.

### Invoice Examples

These examples demonstrate professional invoice layouts using the table component:

#### Invoice Pattern 1 - Modern Blue Style

**Location:** `internal/example/demos/invoice/invoice_pattern1_demo.go`

**Features:**
- Modern blue header with company branding
- Alternating row colors for better readability
- Itemized product table with quantity, price, and totals
- Summary section with subtotal, tax, and grand total

**Visual Output:**
![Invoice Pattern 1](image/demo/invoice_pattern1.jpg)

Run the demo:
```bash
cd internal/example/demos/invoice
go run invoice_pattern1_demo.go
```

#### Invoice Pattern 2 - Detailed Layout

**Location:** `internal/example/demos/invoice/invoice_pattern2_demo.go`

**Features:**
- Comprehensive invoice with shipment information
- Multiple table sections (items, summary, terms)
- Extended item details with tax calculations
- QR code integration for tracking

**Visual Output:**
![Invoice Pattern 2](image/demo/invoice_pattern2.jpg)

Run the demo:
```bash
cd internal/example/demos/invoice
go run invoice_pattern2_demo.go
```

#### Thai Invoice (ใบแจ้งหนี้)

**Location:** `internal/example/demos/invoice/invoice_demo.go`

**Features:**
- Full Thai language support using embedded fonts
- Professional Thai invoice layout
- Billing information in Thai
- Tax calculation display

**Visual Output:**
![Thai Invoice](image/demo/invoice.jpg)

Run the demo:
```bash
cd internal/example/demos/invoice
go run invoice_demo.go
```

#### Thai Quotation (ใบเสนอราคา)

**Location:** `internal/example/demos/invoice/quotation_demo.go`

**Features:**
- Thai quotation document format
- Similar layout to invoice but for quotations
- Signature sections for approval workflow

**Visual Output:**
![Thai Quotation](image/demo/quotation_thai.jpg)

Run the demo:
```bash
cd internal/example/demos/invoice
go run quotation_demo.go
```

### Booking Acknowledgement

**Location:** `internal/example/demos/booking/main.go`

Demonstrates table usage in logistics/shipping documents:
- Booking details table
- Route information display
- Cargo information formatting
- Professional document layout

**Visual Output:**
![Booking Demo](image/demo/booking_demo.jpg)

Run the demo:
```bash
cd internal/example/demos/booking
go run main.go
```

### Code Examples from Demos

#### Example: Invoice Items Table

```go
// From invoice_pattern1_demo.go
colWidths := []float64{80, 25, 35, 35}
colHeaders := []string{"DESCRIPTION", "QTY", "UNIT PRICE", "TOTAL"}
colAligns := []string{"L", "C", "R", "R"}

// Using the table component (alternative to manual drawing)
columns := []table.Column{
    {Key: "description", Label: "DESCRIPTION", Width: 80, Align: "L"},
    {Key: "qty", Label: "QTY", Width: 25, Align: "C"},
    {Key: "unit_price", Label: "UNIT PRICE", Width: 35, Align: "R"},
    {Key: "total", Label: "TOTAL", Width: 35, Align: "R"},
}

tbl := table.NewTable(pdf, columns).
    WithHeaderStyle(table.CellStyle{
        Border:    "1",
        Bold:      true,
        FillColor: []int{30, 136, 229},  // Modern blue
        TextColor: []int{255, 255, 255},
    }).
    WithDataStyle(table.CellStyle{
        Border: "1",
    }).
    WithAlternatingRows(true).
    WithRowHeight(6)

data := []map[string]interface{}{
    {"description": "Professional PDF Library License", "qty": "1", "unit_price": "299.00", "total": "299.00"},
    {"description": "Advanced Table Module", "qty": "2", "unit_price": "149.00", "total": "298.00"},
}

tbl.Render(true, data)
```

#### Example: Nested Table

```go
// From nested/main.go
mainColumns := []table.Column{
    {Key: "col1", Label: "Main Column 1", Width: 60, Align: "L"},
    {Key: "col2", Label: "Main Column 2", Width: 80, Align: "L"},
}

mainTbl := table.NewTable(pdf, mainColumns)

// Create nested table
nestedColumns := []table.Column{
    {Key: "ncol1", Label: "Inner Col 1", Width: 30, Align: "L"},
    {Key: "ncol2", Label: "Inner Col 2", Width: 35, Align: "L"},
}

nestedTbl := table.NewTable(pdf, nestedColumns)
nestedTbl.AddRows([]map[string]interface{}{
    {"ncol1": "Inner Row 1 Col 1", "ncol2": "Inner Row 1 Col 2"},
    {"ncol1": "Inner Row 2 Col 1", "ncol2": "Inner Row 2 Col 2"},
})

// Add row with nested table
mainTbl.AddRow(map[string]interface{}{
    "col1":         "Main Row 1 Col 1",
    "col2":         "",
    "_nested_col2": nestedTbl,  // Nested table in col2
})
```

#### Example: Multi-Page Table

```go
// From pagebreak/main.go
columns := []table.Column{
    {Key: "id", Label: "ID", Width: 20, Align: "C"},
    {Key: "name", Label: "Product Name", Width: 70, Align: "L"},
    {Key: "price", Label: "Price", Width: 30, Align: "R"},
}

tbl := table.NewTable(pdf, columns).
    WithRepeatHeader(true).       // Repeat headers on new pages
    WithPageBreakMode(true).      // Enable auto page breaks
    WithPageBreakMargin(20)       // 20mm margin from bottom

tbl.AddHeader()

// Add many rows - page breaks happen automatically
for i := 1; i <= 100; i++ {
    tbl.AddRow(map[string]interface{}{
        "id":    fmt.Sprintf("%d", i),
        "name":  fmt.Sprintf("Product %d", i),
        "price": fmt.Sprintf("$%.2f", float64(i)*10.99),
    })
}
```

## Advanced Features

### Multi-Page Tables with Page Breaks

The table component automatically handles page breaks for tables that span multiple pages. See the [Multi-Page Tables](#multi-page-tables-with-page-breaks) example section above for visual demos.

**Quick Reference:**

```go
tbl := table.NewTable(pdf, columns).
    WithRepeatHeader(true).       // Repeat headers on new pages (default: true)
    WithPageBreakMode(true).      // Enable auto page breaks (default: true)
    WithPageBreakMargin(20)       // 20mm margin from bottom (default: 20)

tbl.AddHeader()

// Add many rows - page breaks happen automatically
for i := 1; i <= 100; i++ {
    tbl.AddRow(data)
}
```

**Configuration Options:**

- **`RepeatHeader`**: When `true`, table headers are automatically repeated on each new page (default: `true`)
- **`PageBreakMode`**: When `true`, enables automatic page break detection (default: `true`)
- **`PageBreakMargin`**: Distance from bottom of page (in mm) before triggering a page break (default: `20.0`)

**Example: Disable Header Repetition**

```go
tbl := table.NewTable(pdf, columns).
    WithRepeatHeader(false)  // Headers only on first page
```

**Example: Custom Page Break Margin**

```go
tbl := table.NewTable(pdf, columns).
    WithPageBreakMargin(40)  // Break 40mm from bottom (earlier breaks)
```

For complete working examples with visual outputs, see the [Multi-Page Tables with Page Breaks](#multi-page-tables-with-page-breaks) section above or `internal/example/demos/pagebreak/main.go`.

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
