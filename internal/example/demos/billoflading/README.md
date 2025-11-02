# Bill of Lading Demo

A professional Bill of Lading document generator using gofpdf, demonstrating advanced PDF features including barcodes, tables, and modular component architecture.

## Features

- **Professional Layout** - Complete Bill of Lading document with all standard sections
- **Code128 Barcode** - Scannable barcode using the gofpdf barcode library
- **Table Component** - Clean, maintainable tables for Customer Order and Carrier Information
- **Modular Architecture** - Proper Go package structure with importable components
- **Responsive Design** - Proper font sizing and spacing to prevent text overflow
- **Single File Execution** - Run only `main.go` - all components are imported automatically

## Quick Start

### ✅ How to Run (It's Simple!)

**From the billoflading directory:**
```bash
cd internal/example/demos/billoflading
go run main.go
```

**From the gofpdf root directory:**
```bash
go run ./internal/example/demos/billoflading/main.go
```

**That's it!** Just run `main.go` - all components are imported automatically. No need to specify multiple files.

### 📄 Output

The PDF will be created at:
```
pdf/bill_of_lading.pdf
```

## File Structure

```
billoflading/
├── main.go                    # Entry point - imports all components (Just run this!)
├── components/                # Component package
│   ├── header.go             # Date and title rendering
│   ├── shipping.go           # Ship From/To, Third Party, Special Instructions
│   ├── tables.go             # Customer Order and Carrier Information tables
│   └── footer.go             # Footer, signatures, and checkboxes
└── README.md                  # This file
```

**Benefits:**
- ✅ Run only `main.go`
- ✅ No wildcards needed (`*.go`)
- ✅ Standard Go practice
- ✅ Works in any IDE
- ✅ Components are reusable

## Architecture

### Proper Go Package Structure

The demo uses a **proper Go package architecture**:

1. **`main.go`** - Entry point that imports the components package
2. **`components/`** - Separate package with all rendering functions
3. **Exported Functions** - All component functions are capitalized and exported
4. **Clean Imports** - Main file simply imports and calls component functions

### Component Overview

#### **main.go** - Entry Point
```go
import "github.com/looksocial/gofpdf/internal/example/demos/billoflading/components"

func generateBillOfLading(pdf *gofpdf.Fpdf) {
    currentY := components.RenderHeader(pdf, leftMargin, topMargin, contentWidth)
    currentY = components.RenderShipFrom(pdf, leftMargin, currentY, contentWidth, black, white)
    // ...
}
```

#### **components/header.go**
- `RenderHeader()` - Date and title section

#### **components/shipping.go**
- `RenderShipFrom()` - Ship From section with Code128 barcode
- `RenderShipTo()` - Ship To and Carrier info
- `RenderThirdParty()` - Third Party Freight Charges
- `RenderSpecialInstructions()` - Special Instructions with checkboxes

#### **components/tables.go**
- `RenderCustomerOrderInformation()` - Customer order table with totals
- `RenderCarrierInformation()` - Carrier info table with complex headers

#### **components/footer.go**
- `RenderFooter()` - Complete footer with all checkboxes and signatures

## Usage

### Running the Demo

**From billoflading directory:**
```bash
cd internal/example/demos/billoflading
go run main.go
```

**From root directory:**
```bash
go run ./internal/example/demos/billoflading/main.go
```

**Output:** `pdf/bill_of_lading.pdf`

### Using Components in Your Own Code

**Quick Example:**

```go
import "github.com/looksocial/gofpdf/internal/example/demos/billoflading/components"

currentY := components.RenderHeader(pdf, 10, 10, 190)
currentY = components.RenderShipFrom(pdf, 10, currentY, 190, black, white)
```

**Full Example:**

```go
import (
    "github.com/looksocial/gofpdf"
    "github.com/looksocial/gofpdf/internal/example/demos/billoflading/components"
)

func main() {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()

    // Use individual components
    black := []int{0, 0, 0}
    white := []int{255, 255, 255}
    lightGray := []int{230, 230, 230}

    currentY := components.RenderHeader(pdf, 10, 10, 190)
    currentY = components.RenderShipFrom(pdf, 10, currentY, 190, black, white)
    // Add more sections as needed

    pdf.OutputFileAndClose("output.pdf")
}
```

### Adding New Sections

1. Create a new function in the appropriate component file
2. Export the function (capitalize first letter)
3. Return the updated Y position
4. Import and call in `main.go`

```go
// In components/shipping.go
func RenderNewSection(pdf *gofpdf.Fpdf, leftMargin, currentY, contentWidth float64) float64 {
    // Your rendering code here
    return currentY + sectionHeight
}

// In main.go
currentY = components.RenderNewSection(pdf, leftMargin, currentY, contentWidth)
```

## Key Technologies

- **[gofpdf](https://github.com/looksocial/gofpdf)** - PDF generation library
- **[gofpdf/contrib/barcode](https://github.com/looksocial/gofpdf/tree/main/contrib/barcode)** - Barcode generation
- **[gofpdf/table](https://github.com/looksocial/gofpdf/tree/main/table)** - Table component for structured data

## Design Decisions

### Why a Components Package?

**Benefits:**
1. **Single File Execution** - Run only `main.go`, no wildcards needed
2. **Proper Imports** - Standard Go package structure
3. **Reusability** - Components can be imported by other projects
4. **Maintainability** - Clear separation of concerns
5. **Testability** - Components can be unit tested
6. **IDE Support** - Better autocomplete and navigation

### Font Sizing
- Headers: 10pt bold for section titles
- Data: 7pt for table content to prevent overflow
- Notes: 5-6pt for fine print sections

### Color Scheme
- Black (#000000) for headers and borders
- White (#FFFFFF) for header text
- Light Gray (#E6E6E6) for table header backgrounds

### Code Organization
- **Modular** - Each component file handles one aspect
- **Exported Functions** - Capitalized for external use
- **Clear Naming** - Functions describe what they render
- **Consistent Parameters** - All functions use similar signatures

## Output

The generated PDF includes:
- Professional Bill of Lading layout
- Scannable Code128 barcode
- Customer order information table
- Carrier information with commodity description
- Complete footer with signature sections
- All required legal notices and checkboxes

## Tips

### Avoiding Text Overflow
- Use appropriate font sizes for the column width
- Test with the longest expected text
- Consider using MultiCell for wrapping text

### Position Tracking
- All render functions return the current Y position
- Chain functions to build the document top-to-bottom
- Track Y position to manage page breaks

### Styling Consistency
- Define colors once in main function
- Pass color arrays to components
- Use consistent spacing and margins

## Troubleshooting

### 🆘 Having Issues?

### Issue: "package components is not in GOROOT"

**Solution:** Make sure you're running from the correct directory or using the correct path.

**From billoflading directory:**
```bash
go run main.go
```

**From root directory:**
```bash
go run ./internal/example/demos/billoflading/main.go
```

### Issue: "cannot find package"

**Solution:** The components package is in a subdirectory. Run `go mod tidy` if needed, or ensure you're using Go modules.

## Comparison: Old vs New Structure

### 📁 What Changed?

The demo now uses **proper Go package structure**:

```
billoflading/
├── main.go                    # Just run this!
└── components/                # All components here
    ├── header.go
    ├── shipping.go
    ├── tables.go
    └── footer.go
```

### ❌ Old Structure (Multiple files in package main)
```bash
# Required running all files together
go run *.go
go run main.go header.go shipping.go tables.go footer.go
```

### ✅ New Structure (Proper package with imports)
```bash
# Just run main.go
go run main.go
```

**Benefits:**
- ✅ Run only `main.go`
- ✅ No wildcards needed (`*.go`)
- ✅ Standard Go practice
- ✅ Works in any IDE
- ✅ Easy to understand
- ✅ Components can be reused

## License

Same as gofpdf (MIT License)
