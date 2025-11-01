package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/looksocial/gofpdf"
	"github.com/looksocial/gofpdf/table"
)

func main() {
	// Create new PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)

	// Example 1: Simple table with many rows (auto page break with header repeat)
	simpleMultiPageExample(pdf)

	pdf.AddPage()

	// Example 2: Table without header repetition
	noHeaderRepeatExample(pdf)

	pdf.AddPage()

	// Example 3: Custom page break margin
	customMarginExample(pdf)

	// Save PDF
	outputPath := filepath.Join("pdf", "pagebreak_demo.pdf")

	// Ensure the pdf directory exists
	pdfDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(pdfDir, 0755); err != nil {
		panic(err)
	}

	err := pdf.OutputFileAndClose(outputPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("PDF generated successfully:", outputPath)
}

// Example 1: Simple table with many rows (auto page break with header repeat)
func simpleMultiPageExample(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "Example 1: Multi-Page Table with Header Repetition")
	pdf.Ln(15)
	pdf.SetFont("Arial", "", 10)

	// Define columns
	columns := []table.Column{
		{Key: "id", Label: "ID", Width: 20, Align: "C"},
		{Key: "name", Label: "Product Name", Width: 70, Align: "L"},
		{Key: "category", Label: "Category", Width: 50, Align: "L"},
		{Key: "price", Label: "Price", Width: 30, Align: "R"},
		{Key: "stock", Label: "Stock", Width: 20, Align: "C"},
	}

	// Create table with page break enabled and header repetition
	tbl := table.NewTable(pdf, columns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{100, 150, 200},
			TextColor: []int{255, 255, 255},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(8).
		WithRepeatHeader(true).       // Enable header repetition (default)
		WithPageBreakMode(true).      // Enable page breaks (default)
		WithPageBreakMargin(20)       // 20mm margin from bottom

	// Add header
	tbl.AddHeader()

	// Generate many rows with realistic product data
	products := []struct {
		name     string
		category string
		price    float64
	}{
		{"Dell XPS 15 Laptop", "Electronics", 1299.99},
		{"Samsung 4K Monitor 27\"", "Electronics", 449.99},
		{"Wireless Gaming Mouse", "Electronics", 79.99},
		{"Mechanical Keyboard RGB", "Electronics", 159.99},
		{"USB-C Hub 7-in-1", "Electronics", 49.99},
		{"Men's Cotton T-Shirt", "Clothing", 24.99},
		{"Women's Denim Jeans", "Clothing", 59.99},
		{"Winter Jacket Waterproof", "Clothing", 129.99},
		{"Running Shoes Athletic", "Clothing", 89.99},
		{"Baseball Cap Classic", "Clothing", 19.99},
		{"Organic Coffee Beans 1kg", "Food", 18.99},
		{"Premium Green Tea 100pk", "Food", 12.99},
		{"Protein Bar Variety 12pk", "Food", 24.99},
		{"Dark Chocolate 85% Cacao", "Food", 4.99},
		{"Almond Butter Organic", "Food", 14.99},
		{"Clean Code by Robert Martin", "Books", 42.99},
		{"The Pragmatic Programmer", "Books", 39.99},
		{"Design Patterns (GoF)", "Books", 54.99},
		{"Learning Go Programming", "Books", 44.99},
		{"Refactoring 2nd Edition", "Books", 47.99},
		{"LEGO Star Wars Millennium Falcon", "Toys", 169.99},
		{"Remote Control Drone 4K", "Toys", 299.99},
		{"Board Game Settlers of Catan", "Toys", 44.99},
		{"Puzzle 1000 Pieces Landscape", "Toys", 19.99},
		{"Action Figure Marvel Avengers", "Toys", 29.99},
		{"Yoga Mat Non-Slip 6mm", "Sports", 34.99},
		{"Dumbbell Set 20kg Adjustable", "Sports", 89.99},
		{"Tennis Racket Professional", "Sports", 199.99},
		{"Basketball Official Size", "Sports", 39.99},
		{"Resistance Bands Set of 5", "Sports", 24.99},
		{"Bluetooth Headphones Noise Canceling", "Electronics", 279.99},
		{"Smart Watch Fitness Tracker", "Electronics", 199.99},
		{"Portable SSD 1TB External", "Electronics", 129.99},
		{"Webcam 1080p HD", "Electronics", 69.99},
		{"Power Bank 20000mAh", "Electronics", 39.99},
		{"Hoodie Cotton Fleece", "Clothing", 49.99},
		{"Leather Belt Men's", "Clothing", 34.99},
		{"Sunglasses Polarized UV400", "Clothing", 79.99},
		{"Backpack Travel 40L", "Clothing", 69.99},
		{"Socks Pack of 6 Athletic", "Clothing", 16.99},
		{"Olive Oil Extra Virgin 500ml", "Food", 22.99},
		{"Pasta Whole Wheat 500g", "Food", 3.99},
		{"Honey Raw Organic 500g", "Food", 16.99},
		{"Granola Bars 12 Pack", "Food", 9.99},
		{"Mixed Nuts Roasted 500g", "Food", 14.99},
		{"Python Crash Course 3rd Ed", "Books", 39.99},
		{"JavaScript: The Good Parts", "Books", 32.99},
		{"System Design Interview", "Books", 44.99},
		{"Domain-Driven Design", "Books", 54.99},
		{"Building Microservices", "Books", 49.99},
	}

	for i := 0; i < 50; i++ {
		product := products[i%len(products)]
		// Add some variation to stock levels
		stock := (i*7 + 15) % 100
		if stock < 10 {
			stock += 20
		}

		tbl.AddRow(map[string]interface{}{
			"id":       fmt.Sprintf("%d", i+1),
			"name":     product.name,
			"category": product.category,
			"price":    fmt.Sprintf("$%.2f", product.price),
			"stock":    fmt.Sprintf("%d", stock),
		})
	}
}

// Example 2: Table without header repetition
func noHeaderRepeatExample(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "Example 2: Multi-Page Table WITHOUT Header Repetition")
	pdf.Ln(15)
	pdf.SetFont("Arial", "", 10)

	columns := []table.Column{
		{Key: "order", Label: "Order #", Width: 30, Align: "C"},
		{Key: "customer", Label: "Customer", Width: 80, Align: "L"},
		{Key: "date", Label: "Date", Width: 40, Align: "C"},
		{Key: "total", Label: "Total", Width: 40, Align: "R"},
	}

	tbl := table.NewTable(pdf, columns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{200, 100, 100},
			TextColor: []int{255, 255, 255},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(8).
		WithRepeatHeader(false).      // Disable header repetition
		WithPageBreakMode(true)

	tbl.AddHeader()

	// Generate rows with realistic customer data
	customers := []string{
		"John Smith", "Emma Johnson", "Michael Williams", "Sophia Brown",
		"James Jones", "Olivia Garcia", "Robert Miller", "Ava Davis",
		"David Rodriguez", "Isabella Martinez", "William Anderson", "Mia Taylor",
		"Richard Thomas", "Charlotte Jackson", "Joseph White", "Amelia Harris",
		"Charles Martin", "Harper Thompson", "Christopher Lee", "Evelyn Walker",
		"Daniel Hall", "Abigail Allen", "Matthew Young", "Emily King",
		"Anthony Wright", "Elizabeth Lopez", "Mark Hill", "Sofia Scott",
		"Donald Green", "Avery Adams", "Steven Baker", "Ella Nelson",
		"Paul Carter", "Scarlett Mitchell", "Andrew Perez", "Grace Roberts",
		"Joshua Turner", "Chloe Phillips", "Kenneth Campbell", "Victoria Parker",
	}

	for i := 1; i <= 40; i++ {
		customer := customers[(i-1)%len(customers)]
		// Generate varied order totals
		baseTotal := float64(i%20+1) * 15.99
		total := baseTotal + float64(i%10)*12.50

		// Generate realistic dates across November
		day := (i%28) + 1

		tbl.AddRow(map[string]interface{}{
			"order":    fmt.Sprintf("ORD-%04d", 1000+i),
			"customer": customer,
			"date":     fmt.Sprintf("2024-11-%02d", day),
			"total":    fmt.Sprintf("$%.2f", total),
		})
	}
}

// Example 3: Custom page break margin
func customMarginExample(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "Example 3: Custom Page Break Margin (40mm)")
	pdf.Ln(15)
	pdf.SetFont("Arial", "", 10)

	columns := []table.Column{
		{Key: "item", Label: "Item", Width: 100, Align: "L"},
		{Key: "description", Label: "Description", Width: 90, Align: "L"},
	}

	tbl := table.NewTable(pdf, columns).
		WithHeaderStyle(table.CellStyle{
			Border:    "1",
			Bold:      true,
			FillColor: []int{150, 200, 100},
		}).
		WithDataStyle(table.CellStyle{
			Border: "1",
		}).
		WithRowHeight(10).
		WithRepeatHeader(true).
		WithPageBreakMargin(40)       // Larger margin - breaks earlier

	tbl.AddHeader()

	// Realistic inventory items with detailed descriptions
	items := []struct {
		name        string
		description string
	}{
		{"Network Router Enterprise", "High-performance dual-band router with 1.8 GHz quad-core processor, supporting up to 50 devices simultaneously"},
		{"CAT6 Ethernet Cable", "Professional grade shielded twisted pair cable, 23 AWG copper, rated for 10 Gbps up to 55 meters"},
		{"Managed Network Switch", "24-port gigabit switch with VLAN support, QoS capabilities, and web-based management interface"},
		{"Wireless Access Point", "Dual-band 802.11ax access point with MU-MIMO technology and seamless roaming support"},
		{"UPS Battery Backup", "1500VA uninterruptible power supply with pure sine wave output and automatic voltage regulation"},
		{"Server Rack Cabinet", "42U professional rack enclosure with lockable front door, cable management, and ventilation fans"},
		{"Patch Panel 48-Port", "Cat6 certified patch panel with color-coded ports and integrated cable management bar"},
		{"Fiber Optic Cable", "Single-mode LC-LC duplex fiber cable, 9/125um core, suitable for long-distance connections"},
		{"KVM Switch 8-Port", "Rack-mountable KVM switch supporting 4K resolution and USB peripheral sharing across 8 servers"},
		{"Network Attached Storage", "4-bay NAS with RAID 5 support, dual gigabit NICs, and cloud backup integration"},
		{"IP Security Camera", "4MP outdoor PoE camera with night vision, motion detection, and H.265 video compression"},
		{"PoE Injector", "Single-port 802.3af/at power over ethernet injector, 30W max output, surge protection included"},
		{"Cable Tester Professional", "Advanced network cable tester for Cat5e/6/6a with wiremap, length, and PoE testing"},
		{"Wall Mount Rack", "6U hinged wall mount bracket with 18-inch depth, suitable for small network installations"},
		{"Cooling Fan Tray", "Thermostat-controlled 2-fan unit for rack cabinets with adjustable temperature settings"},
		{"Surge Protector Rack", "1U horizontal PDU with 8 outlets, 15A capacity, and LED power indicator"},
		{"Console Server", "16-port serial console server for remote out-of-band management of network devices"},
		{"Media Converter", "Ethernet to fiber media converter supporting auto-negotiation and link fault pass-through"},
		{"Wireless Controller", "Centralized WLAN controller managing up to 100 access points with cloud-based analytics"},
		{"Network Monitoring Tool", "Software license for real-time network performance monitoring and traffic analysis"},
		{"Rack Shelf 2U", "Heavy-duty vented shelf supporting up to 50 lbs, suitable for non-rack-mountable equipment"},
		{"Cable Management Arm", "Articulating cable management arm for server rack, preventing cable strain during maintenance"},
		{"Blank Panel Set", "10-pack of 1U blank panels for improving airflow and aesthetics in server racks"},
		{"PDU Smart Metered", "Intelligent power distribution unit with per-outlet monitoring and remote reboot capability"},
		{"Grounding Kit", "Professional rack grounding kit with copper braid and bonding hardware for electrical safety"},
		{"Labeling System", "Thermal label printer with adhesive labels for cable and port identification"},
		{"Tool Kit Network", "Professional networking tool set including crimper, stripper, punch-down tool, and tester"},
		{"Screws & Cage Nuts", "Assortment of rack screws, cage nuts, and washers for standard 19-inch equipment mounting"},
		{"Velcro Cable Ties", "100-pack of reusable hook-and-loop cable straps in assorted colors for organization"},
		{"Documentation Holder", "Magnetic document holder attaching to rack for storing network diagrams and procedures"},
	}

	for i := 0; i < 30; i++ {
		item := items[i%len(items)]
		tbl.AddRow(map[string]interface{}{
			"item":        item.name,
			"description": item.description,
		})
	}
}

