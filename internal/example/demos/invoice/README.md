# Thai Invoice Demo (ใบแจ้งหนี้)

This demo showcases how to create a professional invoice document in Thai using the gofpdf library with embedded Thai fonts.

## Features

- ✨ Full Thai language support using embedded fonts
- 📄 Professional invoice layout matching business standards
- 🎨 Custom styling with colors and formatting
- 📊 Table component for itemized billing
- 💰 Automatic tax and total calculations
- 📝 Signature sections for approval workflow

## Generated Invoice Includes

1. **Header Section**
   - Company branding with blue background
   - Invoice title in both English and Thai

2. **Billing Information**
   - Customer details (ชื่อลูกค้า, ที่อยู่, เลขผู้เสียภาษี, etc.)
   - Invoice metadata (เลขที่, วันที่, ครบกำหนด, อ้างอิง)
   - Issuer details (ผู้ออก)

3. **Items Table**
   - Line items with description, quantity, price, and total
   - Supports Thai product names and descriptions
   - Clean table formatting with proper alignment

4. **Summary Section**
   - Subtotal (ราคารวม)
   - VAT calculation (ภาษีมูลค่าเพิ่ม 7%)
   - Discounts (ส่วนลด)
   - Grand total with Thai text representation

5. **Payment & Signatures**
   - Payment method information
   - Approval and recipient signature sections
   - Terms and conditions

## Usage

### Run the demo:

```bash
cd internal/example/demos/invoice
go run invoice_demo.go
```

This will generate `invoice_thai.pdf` in the `pdf` directory at the root of the project:
```
C:\Users\akkaraponph\Documents\workspace\codespace\gofpdf\pdf\invoice_thai.pdf
```

## Thai Fonts

The demo uses embedded Thai fonts which are automatically loaded:

- **Sarabun** - For body text and most content
- **Kanit** - For headers and titles

No external font files are required as the fonts are embedded in the library.

## Code Structure

### Main Components

1. **PDF Setup**
   ```go
   pdf := gofpdf.New("P", "mm", "A4", "")
   pdf.UseEmbeddedFonts() // Enable embedded Thai fonts
   ```

2. **Thai Text**
   ```go
   pdf.SetFont("Sarabun", "", 9)
   pdf.Cell(40, 5, "ชื่อลูกค้า") // Thai text renders correctly
   ```

3. **Table Drawing**
   ```go
   // Manual table drawing for precise border control
   colWidths := []float64{20, 80, 25, 35, 35}
   colHeaders := []string{"ลำดับ", "รายการสินค้า", "จำนวน", "ราคา/หน่วย", "ราคารวม"}

   // Draw header with gray background
   pdf.SetFillColor(242, 242, 242)
   pdf.SetFont("Sarabun", "B", 10)
   // ... render header cells

   // Draw borders to avoid overlapping
   pdf.SetLineWidth(0.2)
   pdf.Rect(tableStartX, tableStartY, totalTableWidth, totalHeight, "D")
   ```

4. **Border Control**
   ```go
   // Clean borders without overlaps
   // Draw outer rectangle once
   pdf.Rect(x, y, width, height, "D")

   // Draw internal grid lines separately
   pdf.Line(x1, y1, x2, y2)
   ```

## Customization

You can easily customize:

- **Colors**: Modify the RGB values in the color definitions
- **Fonts**: Change between Sarabun, Kanit, Prompt, or other Thai fonts
- **Layout**: Adjust positioning using X/Y coordinates
- **Data**: Update the invoice data, items, and calculations
- **Company Info**: Replace with your actual company details

## Example Data

The demo uses sample data:

```go
rows := []map[string]interface{}{
    {
        "no":          "1",
        "description": "อาคาร A",
        "quantity":    "1",
        "price":       "1,200,000.00",
        "total":       "1,200,000.00",
    },
    // ... more items
}
```

Replace this with your actual invoice data.

## Dependencies

- `github.com/looksocial/gofpdf` - Core PDF generation library

## Notes

- All Thai text is UTF-8 encoded
- Font sizes are in points (pt)
- Measurements are in millimeters (mm)
- Page size is A4 (210mm x 297mm)
- Portrait orientation is used
- Table uses manual drawing to avoid border overlaps and ensure clean lines
- Line width is set to 0.2mm for professional appearance

## Output

The generated PDF is saved to `pdf/invoice_thai.pdf` and includes:
- Professional layout suitable for business use
- Properly rendered Thai characters
- Print-ready format
- All sections properly aligned and formatted

## Further Customization

To create your own invoice based on this template:

1. Update company information (name, address, logo)
2. Modify the color scheme to match your branding
3. Adjust the table columns to match your needs
4. Update calculations based on your tax rules
5. Customize terms and conditions
6. Add your company logo (replace text-based logo with image)

For more examples and documentation, visit:
- [gofpdf Examples](../../)
- [Table Documentation](../../../../table/)
