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
	buildPdfData(&pdfData, resp)
	return pdfData
}

func buildPdfData(pdfData *PdfData, resp *http.Response) {
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
			addTextToPdfData(pdfData, &startCounter, tknData, &counter)
		case html.StartTagToken:
			addImageToPdfData(pdfData, htmlTokens)
		}
	}
}

func addTextToPdfData(pdfData *PdfData, startCounter *bool, text string, c *int ) {
	if strings.Contains(text, "Point Forecast:") {
		pdfData.pointForecast = text
		*startCounter = true
	}
	if *startCounter {
		switch *c {
		case 1:
			pdfData.city = text
		case 2:
			pdfData.coords = text
			*startCounter = false
		}
		*c++
	}
}

func addImageToPdfData(pdfData *PdfData, hTkn *html.Tokenizer) {
	t := hTkn.Token()
	if t.Data == "img" {
		if t.Attr[len(t.Attr)-1].Val == "#MouseVal" {
			pdfData.ImageUrl = t.Attr[0].Val
		}
	}
}
