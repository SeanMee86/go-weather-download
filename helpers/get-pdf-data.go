package helpers

import (
	"log"
	"net/http"
	"strings"
	"golang.org/x/net/html"
)

type PdfData struct {
	ImageUrl		string
	siteUrl			string
	pointForecast	string
	city			string
	coords			string
}

func GetPdfData(siteUrl string) PdfData {
	var pdfData PdfData
	pdfData.siteUrl = siteUrl
	resp, err := http.Get(siteUrl)
	if err != nil {
		log.Fatal("Errors: ", err)
	}
	defer resp.Body.Close()
	htmlTokens := html.NewTokenizer(resp.Body)
	counter := 0
	startCounter := false
	loop: for {
		tt := htmlTokens.Next()
		switch tt {
		case html.ErrorToken:
			break loop
		case html.TextToken:
			tknData := htmlTokens.Token().Data
			if strings.Contains(tknData, "Point Forecast:") {
				pdfData.pointForecast = tknData
				startCounter = true
			}
			if startCounter {
				switch counter {
				case 1:
					pdfData.city =  tknData
				case 2:
					pdfData.coords = tknData
					startCounter = false
				}
				counter++
			}
		case html.StartTagToken:
			t := htmlTokens.Token()
			if t.Data == "img" {
				if t.Attr[len(t.Attr)-1].Val == "#MouseVal" {
					pdfData.ImageUrl = t.Attr[0].Val
				}
			}
		}
	}
	return pdfData
}
