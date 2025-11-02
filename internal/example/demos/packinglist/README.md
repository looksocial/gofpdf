# Packing List Demo

This demo shows how to create professional packing lists with repeatable headers and footers for multi-page documents.

## Structure

The code is organized into reusable components:

### Data Structures

- **PackingListData**: Main struct containing all document data (header, shipping, products, footer)
- **ProductItem**: Struct for individual product entries
- **Colors**: Color scheme for the document

### Rendering Functions

- **renderHeader()**: Renders the header section (logo, exporter info, document details, consignee, shipping details)
  - This function can be called on every page to repeat the header

- **renderBody()**: Renders the product table with data rows and totals
  - Renders only the products passed to it (for pagination)

- **renderFooter()**: Renders the footer section (additional info, signature)
  - This function can be called on every page to repeat the footer

## Usage

### Single Page Example

```bash
go run main.go
```

Generates: `pdf/packing_list.pdf`

### Multi-Page Example

To demonstrate repeating headers/footers across pages, the multi-page demo is embedded in `main.go`.

To create a multi-page packing list, simply:

1. Split your products array into chunks
2. Add a new page for each chunk
3. Call `renderHeader()`, `renderBody()`, and `renderFooter()` on each page

Example:
```go
for pageNum := 1; pageNum <= totalPages; pageNum++ {
    pdf.AddPage()

    // Update page number in data
    data.PageNumber = pageNum

    // Get products for this page
    pageProducts := getProductsForPage(data.Products, pageNum, productsPerPage)
    data.Products = pageProducts

    // Render sections (header and footer repeat on every page)
    headerHeight := renderHeader(pdf, data, colors)
    bodyStartY := headerHeight + 2
    bodyHeight := renderBody(pdf, data, colors, bodyStartY)
    footerY := bodyStartY + bodyHeight + 2
    renderFooter(pdf, data, colors, footerY)
}
```

## Customization

You can customize:

- **Company Information**: Update `ExporterName`, addresses, contact details
- **Product Data**: Modify the `Products` array with your items
- **Colors**: Change the `Colors` struct values
- **Layout**: Adjust column widths in `renderBody()` function
- **Styling**: Modify fonts, sizes, and spacing in each render function

## Features

✅ Clean separation of header, body, and footer
✅ Repeatable header and footer for multi-page documents
✅ Structured data types for easy customization
✅ Professional layout matching industry standards
✅ Multi-line table headers with proper formatting
✅ Automatic page numbering (X of Y)
✅ Totals calculation support

## Example Output

- Single page: All products fit on one page with complete header and footer
- Multi-page: Header and footer repeat on each page, products distributed across pages
