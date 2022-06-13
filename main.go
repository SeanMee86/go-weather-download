package main

import (
	"log"
	"os"
	"path/filepath"
	"github.com/SeanMee86/go-weather-download/helpers"
)

func main() {
	ex, err := os.Executable()

	if err != nil {
		log.Fatal("Errors: ", err)
	}

	dir := filepath.Dir(ex)
	tmpFile := dir+"\\temp-weather.png"
	d := helpers.FormatDateForFilename()
	pd := helpers.GetPdfData(os.Args[1])
	absImgUrl := baseUrl + pd.ImageUrl
	helpers.DownloadFile(absImgUrl, tmpFile)

	f, err := os.Open(tmpFile)

	if err != nil {
		log.Fatal("Errors: ", err)
	}

	helpers.GeneratePdf(f, d, pd, dir)
	f.Close()
	os.Remove(tmpFile)
}

const baseUrl = "https://forecast.weather.gov/"
