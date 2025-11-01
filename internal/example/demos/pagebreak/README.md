# Page Break Demo

This demo demonstrates the automatic page break functionality of the gofpdf table package.

## Features Demonstrated

### 1. Automatic Page Breaks
The table automatically detects when there isn't enough space on the current page and creates a new page.

### 2. Header Repetition
When enabled (default), table headers are automatically repeated on each new page for better readability.

### 3. Configurable Options

- **RepeatHeader**: Control whether headers repeat on new pages (default: true)
- **PageBreakMode**: Enable/disable automatic page breaks (default: true)
- **PageBreakMargin**: Set the margin from the bottom of the page before triggering a page break (default: 20mm)

## Examples

### Example 1: Basic Multi-Page Table
- 50 rows demonstrating automatic page breaks
- Headers repeat on each page
- Default 20mm page break margin

### Example 2: No Header Repetition
- Same as Example 1 but with `WithRepeatHeader(false)`
- Headers only appear on the first page

### Example 3: Custom Page Break Margin
- Uses 40mm margin from bottom
- Pages break earlier due to larger margin
- Useful when you need more space at the bottom

## Usage

```go
tbl := table.NewTable(pdf, columns).
    WithRepeatHeader(true).       // Repeat headers on new pages
    WithPageBreakMode(true).      // Enable auto page breaks
    WithPageBreakMargin(20)       // 20mm from bottom

tbl.AddHeader()
// Add many rows - page breaks happen automatically
for i := 1; i <= 100; i++ {
    tbl.AddRow(data)
}
```

## Running the Demo

```bash
go run main.go
```

The PDF will be generated at `pdf/pagebreak_demo.pdf`.
