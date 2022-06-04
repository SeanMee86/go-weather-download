package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/jung-kurt/gofpdf"
)

func generatePdf(graph *os.File, date string, dir string) {
	u := "https://forecast.weather.gov/MapClick.php?w0=t&w3u=1&w5=pop&w7 =rain&w14u=1&w15u=1&AheadHour=0&Submit=Submit&FcstType=graphical& textField1=37.4353&textField2=-122.0712&site=all&unit=0&dd=&bw="
	ua := strings.Split(u, " ")
	fu := strings.Join(ua, "")
	fd := formatDate(date)
	var opt gofpdf.ImageOptions
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetAutoPageBreak(true, 10)
	pdf.SetFont("Arial", "", 16)
	opt.ImageType = "png"
	// Add Header
	headerDir := dir+"\\header_img.png"
	pdf.ImageOptions(headerDir, 25, 10, 0, 0, false, opt, 0, "")
	pdf.Ln(25)
	// Add Title
	printPdfLn(pdf, "Weather Forecast Information")
	printPdfLn(pdf, fd)
	pdf.Ln(5)
	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(0, 0, 255)
	// Add Graph URL
	for i := 0; i < len(ua); i++ {
		pdf.CellFormat(200, 5, ua[i], "", 1, gofpdf.AlignLeft, false, 0, fu)
	}
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(5)
	// Add Graph
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(0, 5, "Point Forecast: Moffett Nas/Mtn Vie CA", "", 1, gofpdf.AlignLeft, false, 0, "")
	pdf.CellFormat(0, 5, "37.44N 122.07W (Elev. 3 ft)", "", 1, gofpdf.AlignLeft, false, 0, "")
	pdf.ImageOptions(graph.Name(), 11, 10, 189, 0, true, opt, 0, "")
	err := pdf.OutputFileAndClose(dir+"\\weather-forecast-"+date+".pdf")
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