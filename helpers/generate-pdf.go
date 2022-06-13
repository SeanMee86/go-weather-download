package helpers

import (
	"fmt"
	"os"
	"github.com/jung-kurt/gofpdf"
)

func GeneratePdf(graph *os.File, date string, pd PdfData, dir string) {
	var opt gofpdf.ImageOptions
	opt.ImageType = "png"
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetAutoPageBreak(true, 10)
	pdf.SetFont("Arial", "", 16)
	addContentToPdf(graph, dir, pdf, opt, date, pd)
	err := pdf.OutputFileAndClose(dir+"\\weather-forecast-"+date+".pdf")
	if err != nil {
		fmt.Println("Errors: ", err)
	}
}

func addContentToPdf(graph *os.File, dir string, pdf *gofpdf.Fpdf, opt gofpdf.ImageOptions, date string, pd PdfData) {
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	fd := formatDateToTitle(date)
	// Add Header
	headerDir := dir+"\\header_img.png"
	pdf.ImageOptions(headerDir, 25, 10, 0, 0, false, opt, 0, "")
	pdf.Ln(25)
	// Add Title
	pdf.CellFormat(0, 10, "Weather Forecast Information", "", 1, gofpdf.AlignCenter, false, 0, "")
	pdf.CellFormat(0, 10, fd, "", 1, gofpdf.AlignCenter, false, 0, "")
	pdf.Ln(5)
	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(0, 0, 255)
	// Add Graph URL
	bs := []byte(pd.siteUrl)
	l := len(bs)/3
	pdf.CellFormat(200, 5, string(bs[0:l]), "", 1, gofpdf.AlignLeft, false, 0, pd.siteUrl)
	pdf.CellFormat(200, 5, string(bs[l:l*2]), "", 1, gofpdf.AlignLeft, false, 0, pd.siteUrl)
	pdf.CellFormat(200, 5, string(bs[l*2:len(bs)-1]), "", 1, gofpdf.AlignLeft, false, 0, pd.siteUrl)
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(5)
	// Add Graph
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(0, 5, tr(pd.pointForecast + pd.city), "", 1, gofpdf.AlignLeft, false, 0, "")
	pdf.CellFormat(0, 5, tr(pd.coords), "", 1, gofpdf.AlignLeft, false, 0, "")
	pdf.ImageOptions(graph.Name(), 11, 10, 189, 0, true, opt, 0, "")
}
