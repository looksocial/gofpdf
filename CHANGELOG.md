# Changelog

All notable changes to the gofpdf project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

#### Table Component - Major Feature Release

We're excited to introduce a powerful new **Table Component** for generating professional PDF tables with ease! This feature provides a high-level, type-safe API for creating complex tables with minimal code.

**Key Features:**

- ✅ **Simple API** - Define columns, pass data, and render with just a few lines of code
- ✅ **Fully Customizable** - Complete control over colors, fonts, borders, and alignment
- ✅ **Column Spanning** - Merge cells across multiple columns for complex layouts
- ✅ **Row Spanning** - Cells can span multiple rows for hierarchical data display
- ✅ **Nested Tables** - Tables within table cells for complex document structures
- ✅ **Alternating Rows** - Built-in zebra striping for better readability
- ✅ **Auto Layout** - Automatic column width calculation with manual override
- ✅ **Page Breaks** - Automatic page breaks with header repetition for long tables
- ✅ **Text Wrapping** - Automatic text wrapping in cells
- ✅ **Flexible Alignment** - Column-level, header-specific, style-level, and per-row alignment options

**Usage Example:**

```go
import (
    "github.com/looksocial/gofpdf"
    "github.com/looksocial/gofpdf/table"
)

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
```

**Advanced Features:**

- **Multi-Page Tables**: Automatic page breaks with configurable margins and header repetition
- **Nested Tables**: Embed tables within table cells for complex layouts
- **Row Spanning**: Create cells that span multiple rows for grouped data
- **Custom Styling**: Full control over header and data cell styles
- **Flexible Alignment**: Multiple levels of alignment control (column, header, style, per-row)

**Documentation & Examples:**

- Complete documentation: [`table/README.md`](table/README.md)
- Working examples in `internal/example/demos/tabledemo/`
- Real-world usage examples:
  - Invoice generation (`internal/example/demos/invoice/`)
  - Booking acknowledgements (`internal/example/demos/booking/`)
  - Multi-page tables (`internal/example/demos/pagebreak/`)
  - Nested table structures (`internal/example/demos/nested/`)

**Installation:**

The table component is part of the main gofpdf package. No additional dependencies are required.

```shell
go get github.com/looksocial/gofpdf
```

Import the table package:

```go
import "github.com/looksocial/gofpdf/table"
```

---

## Previous Releases

For changes prior to the table component release, please refer to the git history or the original gofpdf project.

