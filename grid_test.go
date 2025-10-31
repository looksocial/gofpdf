package gofpdf_test

import (
	"math"
	"testing"

	"github.com/looksocial/gofpdf"
	"github.com/looksocial/gofpdf/internal/example"
)

func TestTickmarksAdvanced(t *testing.T) {
	tests := []struct {
		name     string
		min      float64
		max      float64
		validate func([]float64, int, *testing.T)
	}{
		{
			"Simple range",
			0, 100,
			func(ticks []float64, prec int, t *testing.T) {
				if len(ticks) == 0 {
					t.Error("Tickmarks should return at least one tick")
				}
				if ticks[0] > 0 {
					t.Error("First tick should be <= 0")
				}
				if ticks[len(ticks)-1] < 100 {
					t.Error("Last tick should be >= 100")
				}
			},
		},
		{
			"Negative to positive",
			-50, 50,
			func(ticks []float64, prec int, t *testing.T) {
				if len(ticks) < 2 {
					t.Error("Should have at least 2 ticks")
				}
				if ticks[0] > -50 {
					t.Error("First tick should be <= -50")
				}
				if ticks[len(ticks)-1] < 50 {
					t.Error("Last tick should be >= 50")
				}
				// Should span across zero
			},
		},
		{
			"Same values",
			100, 100,
			func(ticks []float64, prec int, t *testing.T) {
				if len(ticks) != 0 {
					t.Error("Same min/max should return empty slice")
				}
			},
		},
		{
			"Reversed",
			100, 0,
			func(ticks []float64, prec int, t *testing.T) {
				if len(ticks) != 0 {
					t.Error("Reversed range should return empty slice")
				}
			},
		},
		{
			"Small range",
			1, 2,
			func(ticks []float64, prec int, t *testing.T) {
				if len(ticks) < 2 {
					t.Error("Should have reasonable number of ticks")
				}
			},
		},
		{
			"Large range",
			0, 1000000,
			func(ticks []float64, prec int, t *testing.T) {
				if len(ticks) < 2 {
					t.Error("Should have ticks for large range")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ticks, prec := gofpdf.Tickmarks(tt.min, tt.max)
			if prec < 0 || prec > 10 {
				t.Errorf("Precision %d seems out of reasonable range", prec)
			}
			tt.validate(ticks, prec, t)
		})
	}
}

func TestNewGrid(t *testing.T) {
	grid := gofpdf.NewGrid(10, 20, 100, 200)
	
	if grid.X(0) != 10 {
		t.Errorf("Grid X(0) should be at x position")
	}
	if grid.Y(0) != 220 {
		t.Errorf("Grid Y(0) should be at y+h position")
	}
}

func TestGridTickmarksContain(t *testing.T) {
	grid := gofpdf.NewGrid(0, 0, 100, 100)
	
	grid.TickmarksContainX(0, 100)
	grid.TickmarksContainY(0, 100)
	
	minX, maxX := grid.XRange()
	if minX > 0 {
		t.Error("XRange min should be <= 0")
	}
	if maxX < 100 {
		t.Error("XRange max should be >= 100")
	}
	
	minY, maxY := grid.YRange()
	if minY > 0 {
		t.Error("YRange min should be <= 0")
	}
	if maxY < 100 {
		t.Error("YRange max should be >= 100")
	}
}

func TestGridTickmarksExtentAdvanced(t *testing.T) {
	grid := gofpdf.NewGrid(0, 0, 100, 100)
	
	grid.TickmarksExtentX(0, 10, 10)
	grid.TickmarksExtentY(0, 10, 10)
	
	minX, maxX := grid.XRange()
	if minX != 0 {
		t.Errorf("TickmarksExtentX min = %f; want 0", minX)
	}
	if maxX != 100 {
		t.Errorf("TickmarksExtentX max = %f; want 100", maxX)
	}
	
	minY, maxY := grid.YRange()
	if minY != 0 {
		t.Errorf("TickmarksExtentY min = %f; want 0", minY)
	}
	if maxY != 100 {
		t.Errorf("TickmarksExtentY max = %f; want 100", maxY)
	}
}

func TestGridConversion(t *testing.T) {
	grid := gofpdf.NewGrid(10, 20, 100, 200)
	grid.TickmarksExtentX(0, 10, 10)
	grid.TickmarksExtentY(0, 20, 10)
	
	// Test X conversion
	x := grid.X(50)
	if x < 10 || x > 110 {
		t.Errorf("X(50) = %f; should be in range [10, 110]", x)
	}
	
	// Test Y conversion (note: Y is flipped)
	y := grid.Y(50)
	if y < 20 || y > 220 {
		t.Errorf("Y(50) = %f; should be in range [20, 220]", y)
	}
	
	// Test width and height
	wd := grid.Wd(10)
	wdAbs := grid.WdAbs(10)
	if wdAbs != math.Abs(wd) {
		t.Error("WdAbs should return absolute value of Wd")
	}
	
	ht := grid.Ht(10)
	htAbs := grid.HtAbs(10)
	if htAbs != math.Abs(ht) {
		t.Error("HtAbs should return absolute value of Ht")
	}
}

func TestGridPosAdvanced(t *testing.T) {
	grid := gofpdf.NewGrid(10, 20, 100, 200)
	
	// Test corners and center
	x, y := grid.Pos(0, 0) // Bottom left
	if math.Abs(x-10) > 1e-6 {
		t.Errorf("Pos(0, 0) X = %f; want 10", x)
	}
	if math.Abs(y-220) > 1e-6 {
		t.Errorf("Pos(0, 0) Y = %f; want 220", y)
	}
	
	x, y = grid.Pos(1, 0) // Bottom right
	if math.Abs(x-110) > 1e-6 {
		t.Errorf("Pos(1, 0) X = %f; want 110", x)
	}
	
	x, y = grid.Pos(0, 1) // Top left
	if math.Abs(y-20) > 1e-6 {
		t.Errorf("Pos(0, 1) Y = %f; want 20", y)
	}
	
	x, y = grid.Pos(1, 1) // Top right
	if math.Abs(x-110) > 1e-6 {
		t.Errorf("Pos(1, 1) X = %f; want 110", x)
	}
	if math.Abs(y-20) > 1e-6 {
		t.Errorf("Pos(1, 1) Y = %f; want 20", y)
	}
	
	x, y = grid.Pos(0.5, 0.5) // Center
	if math.Abs(x-60) > 1e-6 {
		t.Errorf("Pos(0.5, 0.5) X = %f; want 60", x)
	}
	if math.Abs(y-120) > 1e-6 {
		t.Errorf("Pos(0.5, 0.5) Y = %f; want 120", y)
	}
}

func TestGridRanges(t *testing.T) {
	grid := gofpdf.NewGrid(0, 0, 100, 100)
	
	grid.TickmarksExtentX(0, 10, 10)
	grid.TickmarksExtentY(0, 10, 10)
	
	minX, maxX := grid.XRange()
	if minX != 0 {
		t.Errorf("XRange min = %f; want 0", minX)
	}
	if maxX != 100 {
		t.Errorf("XRange max = %f; want 100", maxX)
	}
	
	minY, maxY := grid.YRange()
	if minY != 0 {
		t.Errorf("YRange min = %f; want 0", minY)
	}
	if maxY != 100 {
		t.Errorf("YRange max = %f; want 100", maxY)
	}
}

// Test that Grid actually generates output without errors
func TestGridGeneration(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	grid := gofpdf.NewGrid(20, 30, 170, 230)
	grid.TickmarksContainX(0, 10)
	grid.TickmarksContainY(0, 10)
	grid.Grid(pdf)
	
	fileStr := example.Filename("test_grid")
	err := pdf.OutputFileAndClose(fileStr)
	if err != nil {
		t.Errorf("Failed to generate grid PDF: %v", err)
	}
}

func TestGridPlot(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	
	grid := gofpdf.NewGrid(20, 30, 170, 230)
	grid.TickmarksContainX(0, 10)
	grid.TickmarksContainY(0, 10)
	grid.Grid(pdf)
	
	// Plot a simple function
	pdf.SetDrawColor(255, 0, 0)
	grid.Plot(pdf, 0, 10, 100, func(x float64) float64 {
		return x * x / 10 // Simple parabola
	})
	
	fileStr := example.Filename("test_grid_plot")
	err := pdf.OutputFileAndClose(fileStr)
	if err != nil {
		t.Errorf("Failed to generate grid plot PDF: %v", err)
	}
}

func TestGridPrecision(t *testing.T) {
	grid := gofpdf.NewGrid(0, 0, 100, 100)
	
	// Test that precision is reasonable for various tickmarks
	grid.TickmarksExtentX(0, 1, 100)
	if precision := gofpdf.TickmarkPrecision(1); precision != 0 {
		t.Errorf("Precision for 1 should be 0, got %d", precision)
	}
	
	grid.TickmarksExtentX(0, 0.1, 100)
	if precision := gofpdf.TickmarkPrecision(0.1); precision != 1 {
		t.Errorf("Precision for 0.1 should be 1, got %d", precision)
	}
}

