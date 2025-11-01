package table

import (
	"github.com/looksocial/gofpdf"
)

// Column represents a table column definition
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

// CellStyle represents styling for individual cells
type CellStyle struct {
	FillColor []int  // RGB color for fill {R, G, B}
	TextColor []int  // RGB color for text {R, G, B}
	Border    string // Border style: "1", "LRT", "LR", etc.
	Bold      bool
	Italic    bool
	FontSize  float64 // 0 = use default
	Align     string  // "L", "C", "R" for alignment (overrides column alignment if set)
}

// RowStyle represents styling for entire rows
type RowStyle struct {
	Alternating bool   // Zebra striping
	FillColor   []int  // Fill color for the row
	Border      string // Default border for cells
}

// Table component
type Table struct {
	pdf            *gofpdf.Fpdf
	Columns        []Column
	HeaderStyle    CellStyle
	DataStyle      CellStyle
	RowStyle       RowStyle
	StartX         float64                  // Starting X position (0 = left margin)
	StartY         float64                  // Starting Y position (0 = current position)
	AutoWidth      bool                     // Auto-calculate column widths if true
	RowHeight      float64                  // Row height in mm (0 = auto)
	Spacing        float64                  // Space between rows
	RepeatHeader   bool                     // Repeat header on new pages (default: true)
	PageBreakMode  bool                     // Enable automatic page breaks (default: true)
	PageBreakMargin float64                 // Margin from bottom before page break (default: 20mm)
	rowSpanTracker map[string]int           // Tracks row spans: "colIndex-rowIndex" -> remaining rows
	storedRows     []map[string]interface{} // Stored rows for deferred rendering (used for nested tables)
}

