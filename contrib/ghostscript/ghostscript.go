package main

// This command demonstrates the use of ghotscript to reduce the size
// of generated PDFs. This is based on a comment made by farkerhaiku:
// https://github.com/looksocial/gofpdf/issues/57#issuecomment-185843315

import (
	"fmt"
	"os"
	"os/exec"

	gofpdf "github.com/looksocial/gofpdf"
)

// report พิมพ์ผลการสร้างหรือบีบอัดไฟล์ PDF ไปยังมาตรฐานเอาต์พุต.
// ถ้า err เป็น nil จะตรวจสอบขนาดไฟล์และพิมพ์ "<file>: OK, size <size>" ถ้า os.Stat ล้มเหลวจะพิมพ์ "<file>: bad stat".
// ถ้า err ไม่เป็น nil จะพิมพ์ "<file>: <err>".
func report(fileStr string, err error) {
	if err == nil {
		var info os.FileInfo
		info, err = os.Stat(fileStr)
		if err == nil {
			fmt.Printf("%s: OK, size %d\n", fileStr, info.Size())
		} else {
			fmt.Printf("%s: bad stat\n", fileStr)
		}
	} else {
		fmt.Printf("%s: %s\n", fileStr, err)
	}
}

func newPdf() (pdf *gofpdf.Fpdf) {
	pdf = gofpdf.New("P", "mm", "A4", "../../font")
	pdf.SetCompression(false)
	pdf.AddFont("Calligrapher", "", "calligra.json")
	pdf.AddPage()
	pdf.SetFont("Calligrapher", "", 35)
	pdf.Cell(0, 10, "Enjoy new fonts with FPDF!")
	return
}

func full(name string) {
	report(name, newPdf().OutputFileAndClose(name))
}

func min(name string) {
	cmd := exec.Command("gs", "-sDEVICE=pdfwrite",
		"-dCompatibilityLevel=1.4",
		"-dPDFSETTINGS=/screen", "-dNOPAUSE", "-dQUIET",
		"-dBATCH", "-sOutputFile="+name, "-")
	inPipe, err := cmd.StdinPipe()
	if err == nil {
		errChan := make(chan error, 1)
		go func() {
			errChan <- cmd.Start()
		}()
		err = newPdf().Output(inPipe)
		if err == nil {
			report(name, <-errChan)
		} else {
			report(name, err)
		}
	} else {
		report(name, err)
	}
}

func main() {
	full("full.pdf")
	min("min.pdf")
}