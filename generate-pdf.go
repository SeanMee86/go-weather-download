package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

func generatePdf(graph *os.File, date string) {
	u := "https://forecast.weather.gov/MapClick.php?w0=t&w3u=1&w5=pop&w7=rain&w14u=1&w15u=1&AheadHour=0&Submit=Submit&FcstType=graphical&textField1=37.4353&textField2=-122.0712&site=all&unit=0&dd=&bw="
	fd := formatDate(date)
	var opt gofpdf.ImageOptions
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	opt.ImageType = "png"
	pdf.ImageOptions("header_img.png", 25, 10, 0, 0, false, opt, 0, "")
	pdf.Ln(40)
	printPdfLn(pdf, "Weather Forecast Information")
	printPdfLn(pdf, fd)
	pdf.Ln(30)
	pdf.CellFormat(0, 10, u, "", 1, gofpdf.AlignLeft, false, 0, u)

	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		fmt.Println("Errors: ", err)
	}
}

func printPdfLn(pdf *gofpdf.Fpdf, text string) {
	pdf.CellFormat(0, 10, text, "", 1, gofpdf.AlignCenter, false, 0, "")
}

func formatDate(date string) string {
	da := strings.Split(date, "-")
	return da[len(da)-1] + ", " + da[1] + " " + da[2] + ", " + da[0]
}