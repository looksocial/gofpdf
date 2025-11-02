package table

import (
	"github.com/looksocial/gofpdf"
)

// Column represents a table column definition that specifies how a column
// should be rendered, including its appearance, alignment, and behavior.
//
// Fields:
//   - Key: The field key used to access data from row maps. This should match
//     the keys in the data maps passed to AddRow or Render.
//   - Label: The text displayed in the header row for this column.
//   - Width: Column width in millimeters. Set to 0 for automatic width calculation
//     based on available space.
//   - MaxWidth: Maximum width constraint for text wrapping in this column (in mm).
//     Text exceeding this width will be wrapped to multiple lines.
//   - Align: Alignment for data cells. Valid values: "L" (left), "C" (center), "R" (right).
//   - HeaderAlign: Alignment specifically for the header cell. Overrides the Align
//     field for header rendering. If empty, falls back to Align, then to style-level alignment.
//   - ColSpan: Number of columns this column should span (default: 1).
//     Used to merge cells horizontally in headers or data rows.
//   - RowSpan: Number of rows this column's cells should span (default: 1).
//     Used to merge cells vertically. Only applies to data cells, not headers.
//   - MergeCell: Deprecated field. Use ColSpan instead for column spanning.
type Column struct {
	Key         string  // Field key for data access
	Label       string  // Header label
	Width       float64 // Column width in mm (0 = auto)
	MaxWidth    float64 // Maximum width for text wrapping
	Align       string  // "L", "C", "R" for alignment (applies to data cells)
	HeaderAlign string  // "L", "C", "R" for header alignment (overrides Align for headers if set)
	ColSpan     int     // Number of columns to span (default 1)
	RowSpan     int     // Number of rows to span (default 1, for data cells only)
	MergeCell   bool    // Deprecated: use ColSpan instead
}

// CellStyle represents styling options for individual table cells, controlling
// their visual appearance including colors, fonts, borders, and alignment.
//
// Fields:
//   - FillColor: RGB color values [R, G, B] for the cell background fill.
//     Each value should be between 0-255. If nil or empty, no fill is applied.
//   - TextColor: RGB color values [R, G, B] for the text color.
//     Each value should be between 0-255. Defaults to black (0, 0, 0) if not specified.
//   - Border: Border style string. Valid values include:
//     "1" (all sides), "L" (left), "R" (right), "T" (top), "B" (bottom),
//     "LR" (left and right), "TB" (top and bottom), "LRT" (left, right, top), etc.
//     Empty string means no border. "0" also means no border.
//   - Bold: If true, renders text in bold font weight.
//   - Italic: If true, renders text in italic font style.
//   - FontSize: Font size in points. Set to 0 to use the current PDF font size.
//   - Align: Text alignment override. Valid values: "L" (left), "C" (center), "R" (right).
//     This overrides column-level alignment when set. Priority: per-row alignment >
//     CellStyle.Align > Column.Align for data cells, or HeaderAlign > CellStyle.Align >
//     Column.Align for header cells.
type CellStyle struct {
	FillColor []int  // RGB color for fill {R, G, B}
	TextColor []int  // RGB color for text {R, G, B}
	Border    string // Border style: "1", "LRT", "LR", etc.
	Bold      bool
	Italic    bool
	FontSize  float64 // 0 = use default
	Align     string  // "L", "C", "R" for alignment (overrides column alignment if set)
}

// RowStyle represents styling options that apply to entire table rows,
// allowing for row-level visual customization.
//
// Fields:
//   - Alternating: Enables zebra striping (alternating row colors).
//     When true, even-numbered rows (by position) will have fill color applied.
//   - FillColor: RGB color values [R, G, B] for row background fill when
//     Alternating is enabled. Each value should be between 0-255.
//   - Border: Default border style for cells in rows styled with this RowStyle.
//     Uses the same format as CellStyle.Border. This serves as a fallback when
//     cell-specific borders are not specified.
type RowStyle struct {
	Alternating bool   // Zebra striping
	FillColor   []int  // Fill color for the row
	Border      string // Default border for cells
}

// Table represents a table component that can be rendered on a PDF document.
// It provides a flexible API for creating styled tables with support for
// column/row spanning, nested tables, page breaks, and extensive customization.
//
// Fields:
//   - pdf: The gofpdf Fpdf instance used for rendering. Set automatically by NewTable.
//   - Columns: Slice of Column definitions specifying the table structure.
//   - HeaderStyle: CellStyle applied to header row cells. Default includes
//     borders, bold text, and gray fill (RGB: 200, 200, 200).
//   - DataStyle: CellStyle applied to data row cells. Default includes simple borders.
//   - RowStyle: RowStyle configuration for row-level styling, including zebra striping.
//   - StartX: Starting X position in mm for table rendering. 0 means use current PDF X position
//     (typically left margin). Used to align all rows to a consistent left edge.
//   - StartY: Starting Y position in mm for table rendering. 0 means use current PDF Y position.
//     Only applied to the first row; subsequent rows use updated positions from Ln().
//   - AutoWidth: If true, automatically calculates column widths. Currently not fully implemented;
//     column widths are always auto-calculated for columns with Width = 0.
//   - RowHeight: Height of each row in mm. 0 means auto-calculate based on font size.
//     Auto-calculated height is approximately 40% of font size in points.
//   - Spacing: Space between rows in mm. Applied after each row's Ln() call.
//   - RepeatHeader: If true, header row is automatically repeated when a page break occurs.
//     Default: true.
//   - PageBreakMode: If true, automatically inserts page breaks when there's insufficient
//     space for the next row. Default: true.
//   - PageBreakMargin: Minimum margin in mm from the bottom of the page before triggering
//     a page break. Default: 20mm.
//   - rowSpanTracker: Internal map tracking active row spans. Key format: "colIndex-rowIndex".
//   - storedRows: Internal slice storing rows for deferred rendering, primarily used
//     for nested tables that need to be rendered within parent table cells.
//   - currentRow: Internal counter tracking the logical row index for zebra striping.
//     Incremented each time AddRow is called, resets only when a new table is created.
type Table struct {
	pdf                *gofpdf.Fpdf
	Columns            []Column
	HeaderStyle        CellStyle
	DataStyle          CellStyle
	RowStyle           RowStyle
	StartX             float64                  // Starting X position (0 = left margin)
	StartY             float64                  // Starting Y position (0 = current position)
	AutoWidth          bool                     // Auto-calculate column widths if true
	RowHeight          float64                  // Row height in mm (0 = auto)
	Spacing            float64                  // Space between rows
	RepeatHeader       bool                     // Repeat header on new pages (default: true)
	PageBreakMode      bool                     // Enable automatic page breaks (default: true)
	PageBreakMargin    float64                  // Margin from bottom before page break (default: 20mm)
	CustomRepeatHeader func() float64           // Custom header function to call on page break (returns new Y position)
	rowSpanTracker     map[string]int           // Tracks row spans: "colIndex-rowIndex" -> remaining rows
	storedRows         []map[string]interface{} // Stored rows for deferred rendering (used for nested tables)
	currentRow         int                      // Logical row index for zebra striping (0-indexed)
}

